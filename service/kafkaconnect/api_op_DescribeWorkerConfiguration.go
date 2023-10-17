// Code generated by smithy-go-codegen DO NOT EDIT.

package kafkaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a worker configuration.
func (c *Client) DescribeWorkerConfiguration(ctx context.Context, params *DescribeWorkerConfigurationInput, optFns ...func(*Options)) (*DescribeWorkerConfigurationOutput, error) {
	if params == nil {
		params = &DescribeWorkerConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkerConfiguration", params, optFns, c.addOperationDescribeWorkerConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkerConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkerConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the worker configuration that you want to get
	// information about.
	//
	// This member is required.
	WorkerConfigurationArn *string

	noSmithyDocumentSerde
}

type DescribeWorkerConfigurationOutput struct {

	// The time that the worker configuration was created.
	CreationTime *time.Time

	// The description of the worker configuration.
	Description *string

	// The latest revision of the custom configuration.
	LatestRevision *types.WorkerConfigurationRevisionDescription

	// The name of the worker configuration.
	Name *string

	// The Amazon Resource Name (ARN) of the custom configuration.
	WorkerConfigurationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWorkerConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeWorkerConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeWorkerConfiguration{}, middleware.After)
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
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeWorkerConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkerConfiguration(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDescribeWorkerConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafkaconnect",
		OperationName: "DescribeWorkerConfiguration",
	}
}
