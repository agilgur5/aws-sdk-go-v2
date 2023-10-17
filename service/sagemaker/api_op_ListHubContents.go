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

// List the contents of a hub. Hub APIs are only callable through SageMaker Studio.
func (c *Client) ListHubContents(ctx context.Context, params *ListHubContentsInput, optFns ...func(*Options)) (*ListHubContentsOutput, error) {
	if params == nil {
		params = &ListHubContentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHubContents", params, optFns, c.addOperationListHubContentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHubContentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHubContentsInput struct {

	// The type of hub content to list.
	//
	// This member is required.
	HubContentType types.HubContentType

	// The name of the hub to list the contents of.
	//
	// This member is required.
	HubName *string

	// Only list hub content that was created after the time specified.
	CreationTimeAfter *time.Time

	// Only list hub content that was created before the time specified.
	CreationTimeBefore *time.Time

	// The maximum amount of hub content to list.
	MaxResults *int32

	// The upper bound of the hub content schema verion.
	MaxSchemaVersion *string

	// Only list hub content if the name contains the specified string.
	NameContains *string

	// If the response to a previous ListHubContents request was truncated, the
	// response includes a NextToken . To retrieve the next set of hub content, use the
	// token in the next request.
	NextToken *string

	// Sort hub content versions by either name or creation time.
	SortBy types.HubContentSortBy

	// Sort hubs by ascending or descending order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListHubContentsOutput struct {

	// The summaries of the listed hub content.
	//
	// This member is required.
	HubContentSummaries []types.HubContentInfo

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of hub content, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHubContentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHubContents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHubContents{}, middleware.After)
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
	if err = addOpListHubContentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHubContents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListHubContents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListHubContents",
	}
}
