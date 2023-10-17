// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the tags for the specified delivery stream. This operation has a limit of
// five transactions per second per account.
func (c *Client) ListTagsForDeliveryStream(ctx context.Context, params *ListTagsForDeliveryStreamInput, optFns ...func(*Options)) (*ListTagsForDeliveryStreamOutput, error) {
	if params == nil {
		params = &ListTagsForDeliveryStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForDeliveryStream", params, optFns, c.addOperationListTagsForDeliveryStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsForDeliveryStreamInput struct {

	// The name of the delivery stream whose tags you want to list.
	//
	// This member is required.
	DeliveryStreamName *string

	// The key to use as the starting point for the list of tags. If you set this
	// parameter, ListTagsForDeliveryStream gets all tags that occur after
	// ExclusiveStartTagKey .
	ExclusiveStartTagKey *string

	// The number of tags to return. If this number is less than the total number of
	// tags associated with the delivery stream, HasMoreTags is set to true in the
	// response. To list additional tags, set ExclusiveStartTagKey to the last key in
	// the response.
	Limit *int32

	noSmithyDocumentSerde
}

type ListTagsForDeliveryStreamOutput struct {

	// If this is true in the response, more tags are available. To list the remaining
	// tags, set ExclusiveStartTagKey to the key of the last tag returned and call
	// ListTagsForDeliveryStream again.
	//
	// This member is required.
	HasMoreTags *bool

	// A list of tags associated with DeliveryStreamName , starting with the first tag
	// after ExclusiveStartTagKey and up to the specified Limit .
	//
	// This member is required.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTagsForDeliveryStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTagsForDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagsForDeliveryStream{}, middleware.After)
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
	if err = addOpListTagsForDeliveryStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForDeliveryStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTagsForDeliveryStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "ListTagsForDeliveryStream",
	}
}
