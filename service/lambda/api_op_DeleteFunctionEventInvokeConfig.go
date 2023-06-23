// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the configuration for asynchronous invocation for a function, version,
// or alias. To configure options for asynchronous invocation, use
// PutFunctionEventInvokeConfig .
func (c *Client) DeleteFunctionEventInvokeConfig(ctx context.Context, params *DeleteFunctionEventInvokeConfigInput, optFns ...func(*Options)) (*DeleteFunctionEventInvokeConfigOutput, error) {
	if params == nil {
		params = &DeleteFunctionEventInvokeConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFunctionEventInvokeConfig", params, optFns, c.addOperationDeleteFunctionEventInvokeConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFunctionEventInvokeConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFunctionEventInvokeConfigInput struct {

	// The name of the Lambda function, version, or alias. Name formats
	//   - Function name - my-function (name-only), my-function:v1 (with alias).
	//   - Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//   - Partial ARN - 123456789012:function:my-function .
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function name,
	// it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// A version number or alias name.
	Qualifier *string

	noSmithyDocumentSerde
}

type DeleteFunctionEventInvokeConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFunctionEventInvokeConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFunctionEventInvokeConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFunctionEventInvokeConfig{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
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
	if err = addDeleteFunctionEventInvokeConfigResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteFunctionEventInvokeConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFunctionEventInvokeConfig(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDeleteFunctionEventInvokeConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "DeleteFunctionEventInvokeConfig",
	}
}

type opDeleteFunctionEventInvokeConfigResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opDeleteFunctionEventInvokeConfigResolveEndpointMiddleware) ID() string {
	return "opDeleteFunctionEventInvokeConfigResolveEndpointMiddleware"
}

func (m *opDeleteFunctionEventInvokeConfigResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	return next.HandleSerialize(ctx, in)
}

func addDeleteFunctionEventInvokeConfigResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDeleteFunctionEventInvokeConfigResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
