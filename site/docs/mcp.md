---
title: Boldpawn MCP Server
description: How Copilot can query the deployed Boldpawn S3 content through MCP.
tags:
  - mcp
  - copilot
  - s3
---

# Boldpawn MCP Server

The Boldpawn MCP server is available at:

```text
https://mcp.boldpawn.com/mcp
```

It exposes read-only tools over the content deployed to the S3 bucket.

## Tools

`search_content` searches deployed Markdown, text, and HTML files. It returns matching snippets and public URLs.

`read_content` reads one deployed content file by S3 key, such as `docs/deployment.md`.

`list_content` lists readable deployed content files.

## Copilot Configuration

Users can add the remote MCP server to their Copilot MCP configuration:

```json
{
  "servers": {
    "boldpawn": {
      "url": "https://mcp.boldpawn.com/mcp"
    }
  }
}
```

After enabling the server, users can ask Copilot questions about the deployed Boldpawn documentation and site content.

