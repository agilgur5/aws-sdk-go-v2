// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the runtime management configuration for a function's version. If the
// runtime update mode is Manual, this includes the ARN of the runtime version and
// the runtime update mode. If the runtime update mode is Auto or Function update,
// this includes the runtime update mode and null is returned for the ARN. For
// more information, see Runtime updates (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html)
// .
func (c *Client) GetRuntimeManagementConfig(ctx context.Context, params *GetRuntimeManagementConfigInput, optFns ...func(*Options)) (*GetRuntimeManagementConfigOutput, error) {
	if params == nil {
		params = &GetRuntimeManagementConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRuntimeManagementConfig", params, optFns, c.addOperationGetRuntimeManagementConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRuntimeManagementConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRuntimeManagementConfigInput struct {

	// The name of the Lambda function. Name formats
	//   - Function name – my-function .
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//   - Partial ARN – 123456789012:function:my-function .
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify a version of the function. This can be $LATEST or a published version
	// number. If no value is specified, the configuration for the $LATEST version is
	// returned.
	Qualifier *string

	noSmithyDocumentSerde
}

type GetRuntimeManagementConfigOutput struct {

	// The Amazon Resource Name (ARN) of your function.
	FunctionArn *string

	// The ARN of the runtime the function is configured to use. If the runtime update
	// mode is Manual, the ARN is returned, otherwise null is returned.
	RuntimeVersionArn *string

	// The current runtime update mode of the function.
	UpdateRuntimeOn types.UpdateRuntimeOn

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRuntimeManagementConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRuntimeManagementConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRuntimeManagementConfig{}, middleware.After)
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
	if err = addOpGetRuntimeManagementConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRuntimeManagementConfig(options.Region), middleware.Before); err != nil {
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
		operation: "GetRuntimeManagementConfig",
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

func newServiceMetadataMiddleware_opGetRuntimeManagementConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetRuntimeManagementConfig",
	}
}
