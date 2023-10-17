// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a destination. This operation is used only to create
// destinations for cross-account subscriptions. A destination encapsulates a
// physical resource (such as an Amazon Kinesis stream). With a destination, you
// can subscribe to a real-time stream of log events for a different account,
// ingested using PutLogEvents (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html)
// . Through an access policy, a destination controls what is written to it. By
// default, PutDestination does not set any access policy with the destination,
// which means a cross-account user cannot call PutSubscriptionFilter (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutSubscriptionFilter.html)
// against this destination. To enable this, the destination owner must call
// PutDestinationPolicy (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestinationPolicy.html)
// after PutDestination . To perform a PutDestination operation, you must also
// have the iam:PassRole permission.
func (c *Client) PutDestination(ctx context.Context, params *PutDestinationInput, optFns ...func(*Options)) (*PutDestinationOutput, error) {
	if params == nil {
		params = &PutDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDestination", params, optFns, c.addOperationPutDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDestinationInput struct {

	// A name for the destination.
	//
	// This member is required.
	DestinationName *string

	// The ARN of an IAM role that grants CloudWatch Logs permissions to call the
	// Amazon Kinesis PutRecord operation on the destination stream.
	//
	// This member is required.
	RoleArn *string

	// The ARN of an Amazon Kinesis stream to which to deliver matching log events.
	//
	// This member is required.
	TargetArn *string

	// An optional list of key-value pairs to associate with the resource. For more
	// information about tagging, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	Tags map[string]string

	noSmithyDocumentSerde
}

type PutDestinationOutput struct {

	// The destination.
	Destination *types.Destination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutDestination{}, middleware.After)
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
	if err = addOpPutDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutDestination",
	}
}
