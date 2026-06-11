# boldpawn.com

Static site infrastructure for `boldpawn.com`, deployed with AWS CDK in Go.

## What It Creates

- Private S3 bucket for static content from `site/`
- CloudFront distribution with Origin Access Control to the S3 bucket
- ACM certificate for `boldpawn.com`, DNS-validated in Route 53
- Route 53 A and AAAA alias records for `boldpawn.com`
- Remote MCP server at `https://mcp.boldpawn.com/mcp` for read-only access to deployed S3 content
- GitHub Actions OIDC deploy role restricted to `boldpawn/boldpawn.com` on `main`, with support for the `prod` environment subject
- GitHub Actions workflow that runs `cdk deploy`

The application stack deploys to `eu-west-1`. CloudFront viewer certificates for custom domains must live in `us-east-1`, so this app also creates a small certificate stack there and references it from the `eu-west-1` stack.

## First Deployment

This repository is configured for AWS account `139702123223`, region `eu-west-1`, hosted zone `Z062991533VQXW91DRKHX`, and the existing GitHub Actions OIDC provider in that account. It also uses the existing CDK bootstrap qualifier `cloudfront`.

Prerequisites:

- `boldpawn.com` is hosted in Route 53 in the default AWS account
- `boldpawn.com` is publicly registered and delegated to the Route 53 hosted zone name servers:
  - `ns-364.awsdns-45.com`
  - `ns-1897.awsdns-45.co.uk`
  - `ns-1101.awsdns-09.org`
  - `ns-985.awsdns-59.net`
- AWS credentials are configured locally for that account
- The account is CDK bootstrapped in `eu-west-1` and `us-east-1`

Bootstrap once if needed:

```sh
npm install
npx cdk bootstrap --qualifier cloudfront aws://$(aws sts get-caller-identity --query Account --output text)/eu-west-1 aws://$(aws sts get-caller-identity --query Account --output text)/us-east-1
```

Deploy locally once to create the site and the GitHub deploy role:

```sh
npm run deploy
```

Set the `GitHubDeployRoleArn` stack output as a GitHub repository variable named `AWS_ROLE_TO_ASSUME`. Also set `AWS_REGION` to `eu-west-1`. After that, pushes to `main` deploy through GitHub Actions.

If you move this to another AWS account, update the CDK context in `cdk.json`. If that account does not already have the GitHub Actions OIDC provider, remove `githubOidcProviderArn` for the first deploy so CDK creates it.

The first deploy attempted on June 10, 2026 was stopped because public DNS returned `NXDOMAIN` for `boldpawn.com`, so ACM could not validate the certificate. Register or delegate the domain before deploying again.

## MCP Server

The stack deploys a stateless remote MCP server backed by AWS Lambda and API Gateway:

```text
https://mcp.boldpawn.com/mcp
```

The MCP server reads the same private S3 bucket that CloudFront serves publicly. It exposes Markdown, text, and HTML files from `site/` as searchable content for Copilot.

Available tools:

- `search_content`: search deployed files and return snippets with source URLs
- `read_content`: read one deployed file by S3 key
- `list_content`: list readable deployed files

Users can add the server to their Copilot MCP configuration:

```json
{
  "servers": {
    "boldpawn": {
      "url": "https://mcp.boldpawn.com/mcp"
    }
  }
}
```

The server is public and read-only because the source content is already public on `boldpawn.com`. The Lambda also supports bearer-token protection with `MCP_BEARER_TOKEN` if private content is added later.

## Useful Commands

```sh
npm run build:mcp
npm run synth
npm run deploy
GOCACHE=$PWD/.gocache go test . ./cmd/mcp-server
```

You can override defaults with CDK context:

```sh
npm run synth -- -c domainName=boldpawn.com -c hostedZoneName=boldpawn.com -c githubRepo=boldpawn/boldpawn.com -c githubBranch=main
```

Override the application region or CloudFront certificate region with:

```sh
npm run synth -- -c region=eu-west-1 -c certificateRegion=us-east-1
```
