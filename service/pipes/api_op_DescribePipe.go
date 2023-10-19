// Code generated by smithy-go-codegen DO NOT EDIT.

package pipes

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get the information about an existing pipe. For more information about pipes,
// see Amazon EventBridge Pipes (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html)
// in the Amazon EventBridge User Guide.
func (c *Client) DescribePipe(ctx context.Context, params *DescribePipeInput, optFns ...func(*Options)) (*DescribePipeOutput, error) {
	if params == nil {
		params = &DescribePipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePipe", params, optFns, c.addOperationDescribePipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePipeInput struct {

	// The name of the pipe.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribePipeOutput struct {

	// The ARN of the pipe.
	Arn *string

	// The time the pipe was created.
	CreationTime *time.Time

	// The state the pipe is in.
	CurrentState types.PipeState

	// A description of the pipe.
	Description *string

	// The state the pipe should be in.
	DesiredState types.RequestedPipeStateDescribeResponse

	// The ARN of the enrichment resource.
	Enrichment *string

	// The parameters required to set up enrichment on your pipe.
	EnrichmentParameters *types.PipeEnrichmentParameters

	// When the pipe was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModifiedTime *time.Time

	// The name of the pipe.
	Name *string

	// The ARN of the role that allows the pipe to send data to the target.
	RoleArn *string

	// The ARN of the source resource.
	Source *string

	// The parameters required to set up a source for your pipe.
	SourceParameters *types.PipeSourceParameters

	// The reason the pipe is in its current state.
	StateReason *string

	// The list of key-value pairs to associate with the pipe.
	Tags map[string]string

	// The ARN of the target resource.
	Target *string

	// The parameters required to set up a target for your pipe.
	TargetParameters *types.PipeTargetParameters

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePipe{}, middleware.After)
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
	if err = addOpDescribePipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePipe(options.Region), middleware.Before); err != nil {
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
		operation: "DescribePipe",
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

func newServiceMetadataMiddleware_opDescribePipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pipes",
		OperationName: "DescribePipe",
	}
}
