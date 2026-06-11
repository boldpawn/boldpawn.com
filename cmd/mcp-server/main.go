package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	protocolVersion  = "2025-06-18"
	serverName       = "boldpawn-content-mcp"
	serverVersion    = "0.1.0"
	fetchConcurrency = 16
	maxObjectBytes   = int64(1_000_000)
	maxReadChars     = 20_000
)

var (
	bucketName     = os.Getenv("CONTENT_BUCKET")
	contentPrefix  = strings.TrimPrefix(os.Getenv("CONTENT_PREFIX"), "/")
	publicBaseURL  = strings.TrimRight(os.Getenv("PUBLIC_BASE_URL"), "/")
	requiredBearer = os.Getenv("MCP_BEARER_TOKEN")

	tagPattern              = regexp.MustCompile(`(?s)<[^>]+>`)
	scriptPattern           = regexp.MustCompile(`(?is)<script[^>]*>.*?</script>|<style[^>]*>.*?</style>|<noscript[^>]*>.*?</noscript>`)
	spacePattern            = regexp.MustCompile(`\s+`)
	punctuationSpacePattern = regexp.MustCompile(`\s+([.,;:!?])`)

	s3Client *s3.Client
	cache    = documentCache{ttl: 15 * time.Minute}
)

type documentCache struct {
	mu        sync.Mutex
	ttl       time.Duration
	expiresAt time.Time
	docs      []document
}

type document struct {
	Key      string
	Name     string
	MimeType string
	Text     string
	URL      string
}

type rpcRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type rpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  interface{}     `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type toolCallResult struct {
	Content []toolContent `json:"content"`
	IsError bool          `json:"isError,omitempty"`
}

type toolContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	s3Client = s3.NewFromConfig(cfg)
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if req.RawPath == "/health" {
		return jsonHTTP(http.StatusOK, map[string]string{"status": "ok", "server": serverName}, req.Headers), nil
	}
	if req.RawPath != "/mcp" && req.RawPath != "/mcp/" {
		return textHTTP(http.StatusNotFound, "not found", req.Headers), nil
	}
	if req.RequestContext.HTTP.Method == http.MethodOptions {
		return emptyHTTP(http.StatusNoContent, req.Headers), nil
	}
	if !originAllowed(req.Headers) {
		return textHTTP(http.StatusForbidden, "origin is not allowed", req.Headers), nil
	}
	if !authorized(req.Headers) {
		return textHTTP(http.StatusUnauthorized, "missing or invalid bearer token", req.Headers), nil
	}
	if req.RequestContext.HTTP.Method == http.MethodGet {
		return textHTTP(http.StatusMethodNotAllowed, "SSE streams are not enabled for this MCP server; use POST", req.Headers), nil
	}
	if req.RequestContext.HTTP.Method != http.MethodPost {
		return textHTTP(http.StatusMethodNotAllowed, "method not allowed", req.Headers), nil
	}

	body, err := requestBody(req)
	if err != nil {
		return rpcHTTP(http.StatusBadRequest, rpcResponse{
			JSONRPC: "2.0",
			ID:      json.RawMessage("null"),
			Error:   &rpcError{Code: -32700, Message: "invalid JSON-RPC payload"},
		}, req.Headers), nil
	}

	var call rpcRequest
	if err := json.Unmarshal(body, &call); err != nil {
		return rpcHTTP(http.StatusBadRequest, rpcResponse{
			JSONRPC: "2.0",
			ID:      json.RawMessage("null"),
			Error:   &rpcError{Code: -32700, Message: "invalid JSON-RPC payload"},
		}, req.Headers), nil
	}
	if call.ID == nil || len(call.ID) == 0 {
		return emptyHTTP(http.StatusAccepted, req.Headers), nil
	}

	result, err := dispatch(ctx, call)
	if err != nil {
		return rpcHTTP(http.StatusOK, rpcResponse{
			JSONRPC: "2.0",
			ID:      call.ID,
			Error:   &rpcError{Code: -32603, Message: err.Error()},
		}, req.Headers), nil
	}

	return rpcHTTP(http.StatusOK, rpcResponse{
		JSONRPC: "2.0",
		ID:      call.ID,
		Result:  result,
	}, req.Headers), nil
}

func requestBody(req events.APIGatewayV2HTTPRequest) ([]byte, error) {
	if !req.IsBase64Encoded {
		return []byte(req.Body), nil
	}
	return base64.StdEncoding.DecodeString(req.Body)
}

func dispatch(ctx context.Context, call rpcRequest) (interface{}, error) {
	switch call.Method {
	case "initialize":
		return initializeResult(), nil
	case "ping":
		return map[string]interface{}{}, nil
	case "tools/list":
		return map[string]interface{}{"tools": tools()}, nil
	case "tools/call":
		return callTool(ctx, call.Params)
	case "resources/list":
		return listResources(ctx)
	case "resources/read":
		return readResource(ctx, call.Params)
	default:
		return nil, fmt.Errorf("unsupported method %q", call.Method)
	}
}

func initializeResult() map[string]interface{} {
	return map[string]interface{}{
		"protocolVersion": protocolVersion,
		"serverInfo": map[string]string{
			"name":    serverName,
			"version": serverVersion,
		},
		"capabilities": map[string]interface{}{
			"tools": map[string]bool{
				"listChanged": false,
			},
			"resources": map[string]bool{
				"subscribe":   false,
				"listChanged": false,
			},
		},
		"instructions": "Use this server to search and read the public Boldpawn content deployed to S3. Prefer search_content first, then read_content for exact source text.",
	}
}

func tools() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "search_content",
			"description": "Search the deployed Boldpawn S3 content and return matching snippets with source URLs.",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"query": map[string]interface{}{
						"type":        "string",
						"description": "Question or keywords to search for.",
					},
					"limit": map[string]interface{}{
						"type":        "integer",
						"description": "Maximum number of results to return.",
						"minimum":     1,
						"maximum":     10,
					},
				},
				"required": []string{"query"},
			},
		},
		{
			"name":        "read_content",
			"description": "Read one deployed content file by S3 key.",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"key": map[string]interface{}{
						"type":        "string",
						"description": "S3 object key to read, for example docs/deployment.md or index.html.",
					},
				},
				"required": []string{"key"},
			},
		},
		{
			"name":        "list_content",
			"description": "List readable deployed content files.",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"prefix": map[string]interface{}{
						"type":        "string",
						"description": "Optional S3 prefix to filter by.",
					},
				},
			},
		},
	}
}

func callTool(ctx context.Context, raw json.RawMessage) (toolCallResult, error) {
	var params struct {
		Name      string                 `json:"name"`
		Arguments map[string]interface{} `json:"arguments"`
	}
	if err := json.Unmarshal(raw, &params); err != nil {
		return toolError("invalid tools/call parameters"), nil
	}
	if params.Arguments == nil {
		params.Arguments = map[string]interface{}{}
	}

	switch params.Name {
	case "search_content":
		return searchContent(ctx, params.Arguments)
	case "read_content":
		return readContent(ctx, params.Arguments)
	case "list_content":
		return listContent(ctx, params.Arguments)
	default:
		return toolError(fmt.Sprintf("unknown tool %q", params.Name)), nil
	}
}

func searchContent(ctx context.Context, args map[string]interface{}) (toolCallResult, error) {
	query := strings.TrimSpace(stringArg(args, "query"))
	if query == "" {
		return toolError("query is required"), nil
	}
	limit := intArg(args, "limit", 5)
	if limit < 1 {
		limit = 1
	}
	if limit > 10 {
		limit = 10
	}

	docs, err := loadDocuments(ctx)
	if err != nil {
		return toolCallResult{}, err
	}
	results := rankDocuments(docs, query, limit)
	if len(results) == 0 {
		return textTool("No matching deployed content was found."), nil
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Found %d matching deployed content file(s) for %q.\n\n", len(results), query)
	for i, result := range results {
		fmt.Fprintf(&b, "%d. %s\n", i+1, result.Doc.Key)
		if result.Doc.URL != "" {
			fmt.Fprintf(&b, "   URL: %s\n", result.Doc.URL)
		}
		fmt.Fprintf(&b, "   Score: %.1f\n", result.Score)
		fmt.Fprintf(&b, "   Excerpt: %s\n\n", result.Excerpt)
	}
	return textTool(strings.TrimSpace(b.String())), nil
}

func readContent(ctx context.Context, args map[string]interface{}) (toolCallResult, error) {
	key := cleanKey(stringArg(args, "key"))
	if key == "" {
		return toolError("key is required"), nil
	}
	docs, err := loadDocuments(ctx)
	if err != nil {
		return toolCallResult{}, err
	}
	for _, doc := range docs {
		if doc.Key == key {
			text := doc.Text
			if len([]rune(text)) > maxReadChars {
				text = string([]rune(text)[:maxReadChars]) + "\n\n[truncated]"
			}
			var b strings.Builder
			fmt.Fprintf(&b, "Key: %s\n", doc.Key)
			if doc.URL != "" {
				fmt.Fprintf(&b, "URL: %s\n", doc.URL)
			}
			fmt.Fprintf(&b, "MIME type: %s\n\n%s", doc.MimeType, text)
			return textTool(b.String()), nil
		}
	}
	return toolError(fmt.Sprintf("content key %q was not found or is not readable", key)), nil
}

func listContent(ctx context.Context, args map[string]interface{}) (toolCallResult, error) {
	prefix := cleanKey(stringArg(args, "prefix"))
	docs, err := loadDocuments(ctx)
	if err != nil {
		return toolCallResult{}, err
	}
	var b strings.Builder
	count := 0
	for _, doc := range docs {
		if prefix != "" && !strings.HasPrefix(doc.Key, prefix) {
			continue
		}
		count++
		fmt.Fprintf(&b, "- %s", doc.Key)
		if doc.URL != "" {
			fmt.Fprintf(&b, " (%s)", doc.URL)
		}
		b.WriteByte('\n')
		if count >= 100 {
			b.WriteString("- [truncated]\n")
			break
		}
	}
	if count == 0 {
		return textTool("No readable deployed content files found."), nil
	}
	return textTool(strings.TrimSpace(b.String())), nil
}

type searchResult struct {
	Doc     document
	Score   float64
	Excerpt string
}

func rankDocuments(docs []document, query string, limit int) []searchResult {
	terms := tokenize(query)
	if len(terms) == 0 {
		return nil
	}
	lowerQuery := strings.ToLower(query)
	results := make([]searchResult, 0, len(docs))
	for _, doc := range docs {
		text := strings.ToLower(doc.Text)
		key := strings.ToLower(doc.Key)
		score := 0.0
		for _, term := range terms {
			textCount := strings.Count(text, term)
			keyCount := strings.Count(key, term)
			score += math.Min(float64(textCount), 12)
			score += float64(keyCount) * 4
		}
		if strings.Contains(text, lowerQuery) {
			score += 15
		}
		if score <= 0 {
			continue
		}
		results = append(results, searchResult{
			Doc:     doc,
			Score:   score,
			Excerpt: excerpt(doc.Text, terms),
		})
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return results[i].Doc.Key < results[j].Doc.Key
		}
		return results[i].Score > results[j].Score
	})
	if len(results) > limit {
		results = results[:limit]
	}
	return results
}

func tokenize(input string) []string {
	seen := map[string]bool{}
	var terms []string
	for _, part := range strings.FieldsFunc(strings.ToLower(input), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}) {
		if len(part) < 2 || seen[part] {
			continue
		}
		seen[part] = true
		terms = append(terms, part)
	}
	return terms
}

func excerpt(text string, terms []string) string {
	clean := strings.TrimSpace(spacePattern.ReplaceAllString(text, " "))
	lower := strings.ToLower(clean)
	pos := -1
	for _, term := range terms {
		if idx := strings.Index(lower, term); idx >= 0 && (pos == -1 || idx < pos) {
			pos = idx
		}
	}
	if pos == -1 {
		pos = 0
	}
	start := pos - 140
	if start < 0 {
		start = 0
	}
	end := start + 360
	if end > len(clean) {
		end = len(clean)
	}
	out := strings.TrimSpace(clean[start:end])
	if start > 0 {
		out = "... " + out
	}
	if end < len(clean) {
		out += " ..."
	}
	return out
}

func listResources(ctx context.Context) (map[string]interface{}, error) {
	docs, err := loadDocuments(ctx)
	if err != nil {
		return nil, err
	}
	resources := make([]map[string]interface{}, 0, len(docs))
	for _, doc := range docs {
		resources = append(resources, map[string]interface{}{
			"uri":         resourceURI(doc.Key),
			"name":        doc.Key,
			"description": "Deployed Boldpawn content file",
			"mimeType":    doc.MimeType,
		})
	}
	return map[string]interface{}{"resources": resources}, nil
}

func readResource(ctx context.Context, raw json.RawMessage) (map[string]interface{}, error) {
	var params struct {
		URI string `json:"uri"`
	}
	if err := json.Unmarshal(raw, &params); err != nil {
		return nil, errors.New("invalid resources/read parameters")
	}
	key, err := keyFromResourceURI(params.URI)
	if err != nil {
		return nil, err
	}
	docs, err := loadDocuments(ctx)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		if doc.Key == key {
			return map[string]interface{}{
				"contents": []map[string]interface{}{
					{
						"uri":      resourceURI(doc.Key),
						"mimeType": doc.MimeType,
						"text":     doc.Text,
					},
				},
			}, nil
		}
	}
	return nil, fmt.Errorf("resource %q was not found", params.URI)
}

func loadDocuments(ctx context.Context) ([]document, error) {
	cache.mu.Lock()
	if time.Now().Before(cache.expiresAt) && cache.docs != nil {
		docs := cache.docs
		cache.mu.Unlock()
		return docs, nil
	}
	cache.mu.Unlock()

	if bucketName == "" {
		return nil, errors.New("CONTENT_BUCKET is not configured")
	}

	var keys []string
	paginator := s3.NewListObjectsV2Paginator(s3Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(contentPrefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.Contents {
			key := aws.ToString(item.Key)
			if key == "" || strings.HasSuffix(key, "/") || !readableKey(key) || aws.ToInt64(item.Size) > maxObjectBytes {
				continue
			}
			keys = append(keys, key)
		}
	}

	docs, err := fetchDocuments(ctx, keys)
	if err != nil {
		return nil, err
	}
	sort.Slice(docs, func(i, j int) bool { return docs[i].Key < docs[j].Key })

	cache.mu.Lock()
	cache.docs = docs
	cache.expiresAt = time.Now().Add(cache.ttl)
	cache.mu.Unlock()

	return docs, nil
}

func fetchDocuments(ctx context.Context, keys []string) ([]document, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	workerCount := fetchConcurrency
	if len(keys) < workerCount {
		workerCount = len(keys)
	}

	jobs := make(chan string)
	results := make(chan document, len(keys))
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for key := range jobs {
				doc, err := fetchDocument(ctx, key)
				if err != nil {
					continue
				}
				select {
				case results <- doc:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	for _, key := range keys {
		select {
		case jobs <- key:
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			close(results)
			return nil, ctx.Err()
		}
	}
	close(jobs)
	wg.Wait()
	close(results)

	docs := make([]document, 0, len(results))
	for doc := range results {
		docs = append(docs, doc)
	}
	return docs, nil
}

func fetchDocument(ctx context.Context, key string) (document, error) {
	obj, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return document{}, err
	}
	defer obj.Body.Close()

	body, err := io.ReadAll(io.LimitReader(obj.Body, maxObjectBytes+1))
	if err != nil {
		return document{}, err
	}
	mimeType := mimeTypeForKey(key)
	text := string(body)
	if strings.Contains(mimeType, "html") {
		text = htmlToText(text)
	}

	return document{
		Key:      key,
		Name:     path.Base(key),
		MimeType: mimeType,
		Text:     strings.TrimSpace(text),
		URL:      publicURL(key),
	}, nil
}

func readableKey(key string) bool {
	switch strings.ToLower(path.Ext(key)) {
	case ".md", ".markdown", ".txt", ".html", ".htm":
		return true
	default:
		return false
	}
}

func mimeTypeForKey(key string) string {
	switch strings.ToLower(path.Ext(key)) {
	case ".md", ".markdown":
		return "text/markdown"
	case ".html", ".htm":
		return "text/html"
	default:
		return "text/plain"
	}
}

func htmlToText(input string) string {
	input = scriptPattern.ReplaceAllString(input, " ")
	input = strings.ReplaceAll(input, "</h1>", "\n")
	input = strings.ReplaceAll(input, "</h2>", "\n")
	input = strings.ReplaceAll(input, "</h3>", "\n")
	input = strings.ReplaceAll(input, "</p>", "\n")
	input = strings.ReplaceAll(input, "</li>", "\n")
	input = tagPattern.ReplaceAllString(input, " ")
	input = html.UnescapeString(input)
	input = spacePattern.ReplaceAllString(input, " ")
	input = punctuationSpacePattern.ReplaceAllString(input, "$1")
	return strings.TrimSpace(input)
}

func publicURL(key string) string {
	if publicBaseURL == "" {
		return ""
	}
	return publicBaseURL + "/" + strings.TrimPrefix(key, "/")
}

func resourceURI(key string) string {
	return "mcp://boldpawn/content/" + url.PathEscape(key)
}

func keyFromResourceURI(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if u.Scheme != "mcp" || u.Host != "boldpawn" || !strings.HasPrefix(u.Path, "/content/") {
		return "", fmt.Errorf("unsupported resource URI %q", raw)
	}
	key, err := url.PathUnescape(strings.TrimPrefix(u.Path, "/content/"))
	if err != nil {
		return "", err
	}
	return cleanKey(key), nil
}

func cleanKey(key string) string {
	key = strings.TrimSpace(strings.TrimPrefix(key, "/"))
	cleaned := path.Clean("/" + key)
	return strings.TrimPrefix(cleaned, "/")
}

func textTool(text string) toolCallResult {
	return toolCallResult{Content: []toolContent{{Type: "text", Text: text}}}
}

func toolError(text string) toolCallResult {
	return toolCallResult{Content: []toolContent{{Type: "text", Text: text}}, IsError: true}
}

func stringArg(args map[string]interface{}, name string) string {
	if value, ok := args[name]; ok {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

func intArg(args map[string]interface{}, name string, fallback int) int {
	value, ok := args[name]
	if !ok {
		return fallback
	}
	switch typed := value.(type) {
	case float64:
		return int(typed)
	case int:
		return typed
	case string:
		parsed, err := strconv.Atoi(typed)
		if err == nil {
			return parsed
		}
	}
	return fallback
}

func originAllowed(headers map[string]string) bool {
	origin := header(headers, "origin")
	if origin == "" {
		return true
	}
	for _, allowed := range allowedOrigins() {
		if origin == allowed {
			return true
		}
	}
	return false
}

func allowedOrigins() []string {
	raw := os.Getenv("MCP_ALLOWED_ORIGINS")
	if raw == "" {
		raw = "https://github.com,https://github.dev,https://vscode.dev,https://mcp.boldpawn.com"
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func authorized(headers map[string]string) bool {
	if requiredBearer == "" {
		return true
	}
	return header(headers, "authorization") == "Bearer "+requiredBearer
}

func header(headers map[string]string, name string) string {
	for key, value := range headers {
		if strings.EqualFold(key, name) {
			return value
		}
	}
	return ""
}

func rpcHTTP(status int, payload rpcResponse, requestHeaders map[string]string) events.APIGatewayV2HTTPResponse {
	body, _ := json.Marshal(payload)
	headers := responseHeaders(requestHeaders)
	headers["Content-Type"] = "application/json"
	return events.APIGatewayV2HTTPResponse{StatusCode: status, Headers: headers, Body: string(body)}
}

func jsonHTTP(status int, payload interface{}, requestHeaders map[string]string) events.APIGatewayV2HTTPResponse {
	body, _ := json.Marshal(payload)
	headers := responseHeaders(requestHeaders)
	headers["Content-Type"] = "application/json"
	return events.APIGatewayV2HTTPResponse{StatusCode: status, Headers: headers, Body: string(body)}
}

func textHTTP(status int, text string, requestHeaders map[string]string) events.APIGatewayV2HTTPResponse {
	headers := responseHeaders(requestHeaders)
	headers["Content-Type"] = "text/plain; charset=utf-8"
	return events.APIGatewayV2HTTPResponse{StatusCode: status, Headers: headers, Body: text}
}

func emptyHTTP(status int, requestHeaders map[string]string) events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{StatusCode: status, Headers: responseHeaders(requestHeaders)}
}

func responseHeaders(requestHeaders map[string]string) map[string]string {
	origin := header(requestHeaders, "origin")
	allowOrigin := "https://mcp.boldpawn.com"
	if originAllowed(requestHeaders) && origin != "" {
		allowOrigin = origin
	}
	return map[string]string{
		"Access-Control-Allow-Headers": "authorization,content-type,mcp-protocol-version,mcp-session-id",
		"Access-Control-Allow-Methods": "GET,POST,OPTIONS",
		"Access-Control-Allow-Origin":  allowOrigin,
		"Vary":                         "Origin",
	}
}
