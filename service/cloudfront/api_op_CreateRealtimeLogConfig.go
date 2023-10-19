// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a real-time log configuration. After you create a real-time log
// configuration, you can attach it to one or more cache behaviors to send
// real-time log data to the specified Amazon Kinesis data stream. For more
// information about real-time log configurations, see Real-time logs (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateRealtimeLogConfig(ctx context.Context, params *CreateRealtimeLogConfigInput, optFns ...func(*Options)) (*CreateRealtimeLogConfigOutput, error) {
	if params == nil {
		params = &CreateRealtimeLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRealtimeLogConfig", params, optFns, c.addOperationCreateRealtimeLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRealtimeLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRealtimeLogConfigInput struct {

	// Contains information about the Amazon Kinesis data stream where you are sending
	// real-time log data.
	//
	// This member is required.
	EndPoints []types.EndPoint

	// A list of fields to include in each real-time log record. For more information
	// about fields, see Real-time log configuration fields (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields)
	// in the Amazon CloudFront Developer Guide.
	//
	// This member is required.
	Fields []string

	// A unique name to identify this real-time log configuration.
	//
	// This member is required.
	Name *string

	// The sampling rate for this real-time log configuration. The sampling rate
	// determines the percentage of viewer requests that are represented in the
	// real-time log data. You must provide an integer between 1 and 100, inclusive.
	//
	// This member is required.
	SamplingRate *int64

	noSmithyDocumentSerde
}

type CreateRealtimeLogConfigOutput struct {

	// A real-time log configuration.
	RealtimeLogConfig *types.RealtimeLogConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRealtimeLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateRealtimeLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateRealtimeLogConfig{}, middleware.After)
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
	if err = addOpCreateRealtimeLogConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRealtimeLogConfig(options.Region), middleware.Before); err != nil {
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
		operation: "CreateRealtimeLogConfig",
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

func newServiceMetadataMiddleware_opCreateRealtimeLogConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateRealtimeLogConfig",
	}
}
