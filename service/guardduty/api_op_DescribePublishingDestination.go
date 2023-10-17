// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the publishing destination specified by the provided
// destinationId .
func (c *Client) DescribePublishingDestination(ctx context.Context, params *DescribePublishingDestinationInput, optFns ...func(*Options)) (*DescribePublishingDestinationOutput, error) {
	if params == nil {
		params = &DescribePublishingDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePublishingDestination", params, optFns, c.addOperationDescribePublishingDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePublishingDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePublishingDestinationInput struct {

	// The ID of the publishing destination to retrieve.
	//
	// This member is required.
	DestinationId *string

	// The unique ID of the detector associated with the publishing destination to
	// retrieve.
	//
	// This member is required.
	DetectorId *string

	noSmithyDocumentSerde
}

type DescribePublishingDestinationOutput struct {

	// The ID of the publishing destination.
	//
	// This member is required.
	DestinationId *string

	// A DestinationProperties object that includes the DestinationArn and KmsKeyArn
	// of the publishing destination.
	//
	// This member is required.
	DestinationProperties *types.DestinationProperties

	// The type of publishing destination. Currently, only Amazon S3 buckets are
	// supported.
	//
	// This member is required.
	DestinationType types.DestinationType

	// The time, in epoch millisecond format, at which GuardDuty was first unable to
	// publish findings to the destination.
	//
	// This member is required.
	PublishingFailureStartTimestamp int64

	// The status of the publishing destination.
	//
	// This member is required.
	Status types.PublishingStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePublishingDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePublishingDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePublishingDestination{}, middleware.After)
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
	if err = addOpDescribePublishingDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePublishingDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePublishingDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "DescribePublishingDestination",
	}
}
