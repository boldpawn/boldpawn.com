package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	apigwv2 "github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	integrations "github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	acm "github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	cloudfront "github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	origins "github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	route53 "github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	targets "github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	s3 "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	s3deploy "github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type siteStackProps struct {
	awscdk.StackProps

	DomainName            string
	HostedZoneName        string
	HostedZoneId          string
	GitHubRepo            string
	GitHubBranch          string
	GitHubEnvironment     string
	GitHubOidcProviderArn string
	CloudFrontCertificate acm.ICertificate
}

type certificateStackProps struct {
	awscdk.StackProps

	DomainName     string
	HostedZoneName string
	HostedZoneId   string
}

func newCertificateStack(scope constructs.Construct, id string, props *certificateStackProps) (awscdk.Stack, acm.ICertificate) {
	stack := awscdk.NewStack(scope, jsii.String(id), &props.StackProps)
	hostedZone := resolveHostedZone(stack, props.HostedZoneName, props.HostedZoneId)

	certificate := acm.NewCertificate(stack, jsii.String("Certificate"), &acm.CertificateProps{
		DomainName: jsii.String(props.DomainName),
		Validation: acm.CertificateValidation_FromDns(hostedZone),
	})

	awscdk.NewCfnOutput(stack, jsii.String("CertificateArn"), &awscdk.CfnOutputProps{
		Value: certificate.CertificateArn(),
	})

	return stack, certificate
}

func newSiteStack(scope constructs.Construct, id string, props *siteStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, jsii.String(id), &props.StackProps)

	hostedZone := resolveHostedZone(stack, props.HostedZoneName, props.HostedZoneId)

	contentBucket := s3.NewBucket(stack, jsii.String("ContentBucket"), &s3.BucketProps{
		BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ALL(),
		Encryption:        s3.BucketEncryption_S3_MANAGED,
		EnforceSSL:        jsii.Bool(true),
		ObjectOwnership:   s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
		RemovalPolicy:     awscdk.RemovalPolicy_RETAIN,
		Versioned:         jsii.Bool(true),
	})

	distribution := cloudfront.NewDistribution(stack, jsii.String("Distribution"), &cloudfront.DistributionProps{
		Certificate:            props.CloudFrontCertificate,
		DefaultRootObject:      jsii.String("index.html"),
		DomainNames:            &[]*string{jsii.String(props.DomainName)},
		MinimumProtocolVersion: cloudfront.SecurityPolicyProtocol_TLS_V1_2_2025,
		DefaultBehavior: &cloudfront.BehaviorOptions{
			AllowedMethods:        cloudfront.AllowedMethods_ALLOW_GET_HEAD_OPTIONS(),
			CachedMethods:         cloudfront.CachedMethods_CACHE_GET_HEAD_OPTIONS(),
			CachePolicy:           cloudfront.CachePolicy_CACHING_OPTIMIZED(),
			Compress:              jsii.Bool(true),
			Origin:                origins.S3BucketOrigin_WithOriginAccessControl(contentBucket, nil),
			ResponseHeadersPolicy: cloudfront.ResponseHeadersPolicy_SECURITY_HEADERS(),
			ViewerProtocolPolicy:  cloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
		},
	})

	route53.NewARecord(stack, jsii.String("ARecord"), &route53.ARecordProps{
		RecordName: jsii.String(props.DomainName),
		Target:     route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
		Zone:       hostedZone,
	})

	route53.NewAaaaRecord(stack, jsii.String("AaaaRecord"), &route53.AaaaRecordProps{
		RecordName: jsii.String(props.DomainName),
		Target:     route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
		Zone:       hostedZone,
	})

	s3deploy.NewBucketDeployment(stack, jsii.String("DeploySite"), &s3deploy.BucketDeploymentProps{
		DestinationBucket: contentBucket,
		Distribution:      distribution,
		DistributionPaths: &[]*string{jsii.String("/*")},
		Sources: &[]s3deploy.ISource{
			s3deploy.Source_Asset(jsii.String("./site"), nil),
		},
	})

	mcpURL := newMCPServer(stack, props, hostedZone, contentBucket)
	deployRole := newGitHubDeployRole(stack, props)

	awscdk.NewCfnOutput(stack, jsii.String("SiteUrl"), &awscdk.CfnOutputProps{
		Value: jsii.String("https://" + props.DomainName),
	})
	awscdk.NewCfnOutput(stack, jsii.String("CloudFrontDomainName"), &awscdk.CfnOutputProps{
		Value: distribution.DistributionDomainName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("ContentBucketName"), &awscdk.CfnOutputProps{
		Value: contentBucket.BucketName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("GitHubDeployRoleArn"), &awscdk.CfnOutputProps{
		Value: deployRole.RoleArn(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("McpServerUrl"), &awscdk.CfnOutputProps{
		Value: jsii.String(mcpURL),
	})

	return stack
}

func newMCPServer(scope constructs.Construct, props *siteStackProps, hostedZone route53.IHostedZone, contentBucket s3.Bucket) string {
	mcpDomainName := "mcp." + props.DomainName
	mcpCertificate := acm.NewCertificate(scope, jsii.String("McpCertificate"), &acm.CertificateProps{
		DomainName: jsii.String(mcpDomainName),
		Validation: acm.CertificateValidation_FromDns(hostedZone),
	})

	mcpFunction := lambda.NewFunction(scope, jsii.String("McpFunction"), &lambda.FunctionProps{
		Architecture: lambda.Architecture_ARM_64(),
		Code:         lambda.Code_FromAsset(jsii.String("./build/mcp-server.zip"), nil),
		Description:  jsii.String("Remote MCP server that exposes deployed Boldpawn S3 content to Copilot."),
		Environment: &map[string]*string{
			"CONTENT_BUCKET":      contentBucket.BucketName(),
			"CONTENT_PREFIX":      jsii.String(""),
			"PUBLIC_BASE_URL":     jsii.String("https://" + props.DomainName),
			"MCP_ALLOWED_ORIGINS": jsii.String("https://github.com,https://github.dev,https://vscode.dev,https://mcp." + props.DomainName),
		},
		Handler:    jsii.String("bootstrap"),
		MemorySize: jsii.Number(256),
		Runtime:    lambda.Runtime_PROVIDED_AL2023(),
		Timeout:    awscdk.Duration_Seconds(jsii.Number(20)),
	})
	contentBucket.GrantRead(mcpFunction, nil)

	mcpDomain := apigwv2.NewDomainName(scope, jsii.String("McpDomainName"), &apigwv2.DomainNameProps{
		Certificate: mcpCertificate,
		DomainName:  jsii.String(mcpDomainName),
	})

	mcpIntegration := integrations.NewHttpLambdaIntegration(jsii.String("McpIntegration"), mcpFunction, nil)
	apigwv2.NewHttpApi(scope, jsii.String("McpHttpApi"), &apigwv2.HttpApiProps{
		ApiName:            jsii.String("boldpawn-content-mcp"),
		DefaultIntegration: mcpIntegration,
		DefaultDomainMapping: &apigwv2.DomainMappingOptions{
			DomainName: mcpDomain,
		},
		Description: jsii.String("MCP endpoint for read-only Boldpawn S3 content search."),
	})

	route53.NewARecord(scope, jsii.String("McpARecord"), &route53.ARecordProps{
		RecordName: jsii.String(mcpDomainName),
		Target: route53.RecordTarget_FromAlias(targets.NewApiGatewayv2DomainProperties(
			mcpDomain.RegionalDomainName(),
			mcpDomain.RegionalHostedZoneId(),
		)),
		Zone: hostedZone,
	})

	return "https://" + mcpDomainName + "/mcp"
}

func newGitHubDeployRole(scope constructs.Construct, props *siteStackProps) iam.Role {
	var provider iam.IOidcProvider
	if props.GitHubOidcProviderArn != "" {
		provider = iam.OidcProviderNative_FromOidcProviderArn(scope, jsii.String("GitHubOidcProvider"), jsii.String(props.GitHubOidcProviderArn))
	} else {
		provider = iam.NewOidcProviderNative(scope, jsii.String("GitHubOidcProvider"), &iam.OidcProviderNativeProps{
			ClientIds: &[]*string{jsii.String("sts.amazonaws.com")},
			Url:       jsii.String("https://token.actions.githubusercontent.com"),
		})
	}

	subjects := []string{
		fmt.Sprintf("repo:%s:environment:%s", props.GitHubRepo, props.GitHubEnvironment),
		fmt.Sprintf("repo:%s:ref:refs/heads/%s", props.GitHubRepo, props.GitHubBranch),
	}
	role := iam.NewRole(scope, jsii.String("GitHubDeployRole"), &iam.RoleProps{
		AssumedBy: iam.NewOpenIdConnectPrincipal(provider, &map[string]interface{}{
			"StringEquals": map[string]interface{}{
				"token.actions.githubusercontent.com:aud": "sts.amazonaws.com",
				"token.actions.githubusercontent.com:sub": subjects,
			},
		}),
		Description:        jsii.String("Allows GitHub Actions to deploy boldpawn.com with AWS CDK."),
		MaxSessionDuration: awscdk.Duration_Hours(jsii.Number(1)),
		RoleName:           jsii.String("boldpawn-com-github-deploy"),
	})

	role.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AdministratorAccess")))

	return role
}

func resolveHostedZone(scope constructs.Construct, hostedZoneName string, hostedZoneId string) route53.IHostedZone {
	if hostedZoneId != "" {
		return route53.HostedZone_FromHostedZoneAttributes(scope, jsii.String("HostedZone"), &route53.HostedZoneAttributes{
			HostedZoneId: jsii.String(hostedZoneId),
			ZoneName:     jsii.String(hostedZoneName),
		})
	}

	return route53.HostedZone_FromLookup(scope, jsii.String("HostedZone"), &route53.HostedZoneProviderProps{
		DomainName: jsii.String(hostedZoneName),
	})
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	account := firstNonEmpty(
		contextString(app, "accountId"),
		os.Getenv("CDK_DEFAULT_ACCOUNT"),
		os.Getenv("AWS_ACCOUNT_ID"),
	)
	region := contextOrDefault(app, "region", "eu-west-1")
	certificateRegion := contextOrDefault(app, "certificateRegion", "us-east-1")
	domainName := contextOrDefault(app, "domainName", "boldpawn.com")
	hostedZoneName := contextOrDefault(app, "hostedZoneName", "boldpawn.com")
	hostedZoneId := contextString(app, "hostedZoneId")

	_, cloudFrontCertificate := newCertificateStack(app, "BoldpawnCertificateStack", &certificateStackProps{
		DomainName:     domainName,
		HostedZoneId:   hostedZoneId,
		HostedZoneName: hostedZoneName,
		StackProps: awscdk.StackProps{
			CrossRegionReferences: jsii.Bool(true),
			Env: &awscdk.Environment{
				Account: optionalString(account),
				Region:  jsii.String(certificateRegion),
			},
		},
	})

	config := siteStackProps{
		DomainName:            domainName,
		CloudFrontCertificate: cloudFrontCertificate,
		GitHubBranch:          contextOrDefault(app, "githubBranch", "main"),
		GitHubEnvironment:     contextOrDefault(app, "githubEnvironment", "prod"),
		GitHubOidcProviderArn: contextOrDefault(app, "githubOidcProviderArn", os.Getenv("GITHUB_OIDC_PROVIDER_ARN")),
		GitHubRepo:            contextOrDefault(app, "githubRepo", "boldpawn/boldpawn.com"),
		HostedZoneId:          hostedZoneId,
		HostedZoneName:        hostedZoneName,
		StackProps: awscdk.StackProps{
			CrossRegionReferences: jsii.Bool(true),
			Env: &awscdk.Environment{
				Account: optionalString(account),
				Region:  jsii.String(region),
			},
		},
	}

	newSiteStack(app, "BoldpawnComStack", &config)

	app.Synth(nil)
}

func contextOrDefault(app awscdk.App, key string, fallback string) string {
	value := contextString(app, key)
	if value != "" {
		return value
	}
	return fallback
}

func contextString(app awscdk.App, key string) string {
	value := app.Node().TryGetContext(jsii.String(key))
	if value == nil {
		return ""
	}

	switch typed := value.(type) {
	case string:
		if typed != "" {
			return typed
		}
	case *string:
		if typed != nil && *typed != "" {
			return *typed
		}
	}

	return ""
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func optionalString(value string) *string {
	if value == "" {
		return nil
	}
	return jsii.String(value)
}
