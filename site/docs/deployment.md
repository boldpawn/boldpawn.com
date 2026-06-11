---
title: Boldpawn Deployment
description: How boldpawn.com is deployed with AWS CDK, S3, CloudFront, Route 53, and GitHub Actions.
tags:
  - aws
  - cdk
  - cloudfront
  - github-actions
---

# Boldpawn Deployment

The `boldpawn.com` site is deployed from this repository with AWS CDK in Go.

The stack creates a private S3 bucket for static content, a CloudFront distribution that reads from the bucket with Origin Access Control, ACM certificates, Route 53 alias records, and a GitHub Actions deploy role.

## Static Site

Human-facing files are deployed from `site/` into the private content bucket. CloudFront is the only public reader of that bucket.

## DNS And Certificates

The stack manages `boldpawn.com` and `mcp.boldpawn.com` in the Route 53 hosted zone `Z062991533VQXW91DRKHX`.

CloudFront custom certificates must be issued in `us-east-1`, so the stack runs in `us-east-1`.

## GitHub Actions

GitHub Actions authenticates to AWS through OIDC. The deploy role ARN is exposed as the `GitHubDeployRoleArn` stack output and should be stored in the repository variable `AWS_ROLE_ARN`.

## Deployment Command

Run:

```sh
npm run deploy
```

The deploy script builds the MCP Lambda package and then runs `cdk deploy --all --require-approval never`.

