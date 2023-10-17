// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List all existing hubs. Hub APIs are only callable through SageMaker Studio.
func (c *Client) ListHubs(ctx context.Context, params *ListHubsInput, optFns ...func(*Options)) (*ListHubsOutput, error) {
	if params == nil {
		params = &ListHubsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHubs", params, optFns, c.addOperationListHubsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHubsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHubsInput struct {

	// Only list hubs that were created after the time specified.
	CreationTimeAfter *time.Time

	// Only list hubs that were created before the time specified.
	CreationTimeBefore *time.Time

	// Only list hubs that were last modified after the time specified.
	LastModifiedTimeAfter *time.Time

	// Only list hubs that were last modified before the time specified.
	LastModifiedTimeBefore *time.Time

	// The maximum number of hubs to list.
	MaxResults *int32

	// Only list hubs with names that contain the specified string.
	NameContains *string

	// If the response to a previous ListHubs request was truncated, the response
	// includes a NextToken . To retrieve the next set of hubs, use the token in the
	// next request.
	NextToken *string

	// Sort hubs by either name or creation time.
	SortBy types.HubSortBy

	// Sort hubs by ascending or descending order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListHubsOutput struct {

	// The summaries of the listed hubs.
	//
	// This member is required.
	HubSummaries []types.HubInfo

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of hubs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHubsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHubs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHubs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHubs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListHubs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListHubs",
	}
}
