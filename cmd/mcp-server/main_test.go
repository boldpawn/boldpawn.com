package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandlerInitializeSupportsTrailingSlashAndBase64Body(t *testing.T) {
	body := `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18"}}`
	resp, err := handler(context.Background(), events.APIGatewayV2HTTPRequest{
		RawPath:         "/mcp/",
		Body:            base64.StdEncoding.EncodeToString([]byte(body)),
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Origin": "https://github.com",
		},
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: http.MethodPost,
			},
		},
	})
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d: %s", resp.StatusCode, http.StatusOK, resp.Body)
	}

	var payload struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  struct {
			ProtocolVersion string `json:"protocolVersion"`
		} `json:"result"`
	}
	if err := json.Unmarshal([]byte(resp.Body), &payload); err != nil {
		t.Fatalf("response is not JSON: %v", err)
	}
	if payload.JSONRPC != "2.0" || payload.ID != 1 || payload.Result.ProtocolVersion != protocolVersion {
		t.Fatalf("unexpected response payload: %#v", payload)
	}
}

func TestHTMLToTextDropsScriptsAndTags(t *testing.T) {
	text := htmlToText(`<html><head><script>alert("x")</script></head><body><h1>Title</h1><p>Hello <strong>Boldpawn</strong>.</p></body></html>`)
	if strings.Contains(text, "alert") || strings.Contains(text, "<strong>") {
		t.Fatalf("htmlToText leaked script or tag content: %q", text)
	}
	if !strings.Contains(text, "Title") || !strings.Contains(text, "Hello Boldpawn.") {
		t.Fatalf("htmlToText lost expected text: %q", text)
	}
}

func TestRankDocumentsFindsSnippets(t *testing.T) {
	results := rankDocuments([]document{
		{Key: "docs/deployment.md", Text: "GitHub Actions deploys the CDK stack to AWS."},
		{Key: "index.html", Text: "A static site home page."},
	}, "GitHub Actions deploy role", 5)
	if len(results) != 1 {
		t.Fatalf("len(results) = %d, want 1", len(results))
	}
	if results[0].Doc.Key != "docs/deployment.md" {
		t.Fatalf("first result = %q, want docs/deployment.md", results[0].Doc.Key)
	}
	if !strings.Contains(results[0].Excerpt, "GitHub Actions") {
		t.Fatalf("excerpt does not include match: %q", results[0].Excerpt)
	}
}
