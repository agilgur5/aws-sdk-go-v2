// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a WebACL per the specifications provided. A web ACL defines a
// collection of rules to use to inspect and control web requests. Each rule has a
// statement that defines what to look for in web requests and an action that WAF
// applies to requests that match the statement. In the web ACL, you assign a
// default action to take (allow, block) for any request that does not match any of
// the rules. The rules in a web ACL can be a combination of the types Rule ,
// RuleGroup , and managed rule group. You can associate a web ACL with one or more
// Amazon Web Services resources to protect. The resources can be an Amazon
// CloudFront distribution, an Amazon API Gateway REST API, an Application Load
// Balancer, an AppSync GraphQL API, an Amazon Cognito user pool, an App Runner
// service, or an Amazon Web Services Verified Access instance.
func (c *Client) CreateWebACL(ctx context.Context, params *CreateWebACLInput, optFns ...func(*Options)) (*CreateWebACLOutput, error) {
	if params == nil {
		params = &CreateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWebACL", params, optFns, c.addOperationCreateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWebACLInput struct {

	// The action to perform if none of the Rules contained in the WebACL match.
	//
	// This member is required.
	DefaultAction *types.DefaultAction

	// The name of the web ACL. You cannot change the name of a web ACL after you
	// create it.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, an App Runner service, or an Amazon Web Services Verified Access
	// instance. To work with CloudFront, you must also specify the Region US East (N.
	// Virginia) as follows:
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	//
	// This member is required.
	VisibilityConfig *types.VisibilityConfig

	// Specifies custom configurations for the associations between the web ACL and
	// protected resources. Use this to customize the maximum size of the request body
	// that your protected CloudFront distributions forward to WAF for inspection. The
	// default is 16 KB (16,384 bytes). You are charged additional fees when your
	// protected resources forward body sizes that are larger than the default. For
	// more information, see WAF Pricing (http://aws.amazon.com/waf/pricing/) .
	AssociationConfig *types.AssociationConfig

	// Specifies how WAF should handle CAPTCHA evaluations for rules that don't have
	// their own CaptchaConfig settings. If you don't specify this, WAF uses its
	// default settings for CaptchaConfig .
	CaptchaConfig *types.CaptchaConfig

	// Specifies how WAF should handle challenge evaluations for rules that don't have
	// their own ChallengeConfig settings. If you don't specify this, WAF uses its
	// default settings for ChallengeConfig .
	ChallengeConfig *types.ChallengeConfig

	// A map of custom response keys and content bodies. When you create a rule with a
	// block action, you can send a custom response to the web request. You define
	// these for the web ACL, and then use them in the rules and default actions that
	// you define in the web ACL. For information about customizing web requests and
	// responses, see Customizing web requests and responses in WAF (https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html)
	// in the WAF Developer Guide. For information about the limits on count and size
	// for custom request and response settings, see WAF quotas (https://docs.aws.amazon.com/waf/latest/developerguide/limits.html)
	// in the WAF Developer Guide.
	CustomResponseBodies map[string]types.CustomResponseBody

	// A description of the web ACL that helps with identification.
	Description *string

	// The Rule statements used to identify the web requests that you want to manage.
	// Each rule includes one top-level statement that WAF uses to identify matching
	// web requests, and parameters that govern how WAF handles them.
	Rules []types.Rule

	// An array of key:value pairs to associate with the resource.
	Tags []types.Tag

	// Specifies the domains that WAF should accept in a web request token. This
	// enables the use of tokens across multiple protected websites. When WAF provides
	// a token, it uses the domain of the Amazon Web Services resource that the web ACL
	// is protecting. If you don't specify a list of token domains, WAF accepts tokens
	// only for the domain of the protected resource. With a token domain list, WAF
	// accepts the resource's host domain plus all domains in the token domain list,
	// including their prefixed subdomains. Example JSON: "TokenDomains": {
	// "mywebsite.com", "myotherwebsite.com" } Public suffixes aren't allowed. For
	// example, you can't use usa.gov or co.uk as token domains.
	TokenDomains []string

	noSmithyDocumentSerde
}

type CreateWebACLOutput struct {

	// High-level information about a WebACL , returned by operations like create and
	// list. This provides information like the ID, that you can use to retrieve and
	// manage a WebACL , and the ARN, that you provide to operations like
	// AssociateWebACL .
	Summary *types.WebACLSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWebACL(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "CreateWebACL",
		options:   options,
	}, "ResolveEndpointV2", middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&signRequestMiddleware{}, middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&getIdentityMiddleware{
		options: options,
	}, middleware.Before)
	if err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "CreateWebACL",
	}
}
