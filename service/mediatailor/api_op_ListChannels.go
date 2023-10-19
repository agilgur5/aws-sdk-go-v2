// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the channels that are associated with the current
// AWS account.
func (c *Client) ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) {
	if params == nil {
		params = &ListChannelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChannels", params, optFns, c.addOperationListChannelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChannelsInput struct {

	// The maximum number of channels that you want MediaTailor to return in response
	// to the current request. If there are more than MaxResults channels, use the
	// value of NextToken in the response to get the next page of results.
	MaxResults int32

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListChannelsOutput struct {

	// A list of channels that are associated with this account.
	Items []types.Channel

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChannelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChannels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChannels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChannels(options.Region), middleware.Before); err != nil {
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
		operation: "ListChannels",
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

// ListChannelsAPIClient is a client that implements the ListChannels operation.
type ListChannelsAPIClient interface {
	ListChannels(context.Context, *ListChannelsInput, ...func(*Options)) (*ListChannelsOutput, error)
}

var _ ListChannelsAPIClient = (*Client)(nil)

// ListChannelsPaginatorOptions is the paginator options for ListChannels
type ListChannelsPaginatorOptions struct {
	// The maximum number of channels that you want MediaTailor to return in response
	// to the current request. If there are more than MaxResults channels, use the
	// value of NextToken in the response to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChannelsPaginator is a paginator for ListChannels
type ListChannelsPaginator struct {
	options   ListChannelsPaginatorOptions
	client    ListChannelsAPIClient
	params    *ListChannelsInput
	nextToken *string
	firstPage bool
}

// NewListChannelsPaginator returns a new ListChannelsPaginator
func NewListChannelsPaginator(client ListChannelsAPIClient, params *ListChannelsInput, optFns ...func(*ListChannelsPaginatorOptions)) *ListChannelsPaginator {
	if params == nil {
		params = &ListChannelsInput{}
	}

	options := ListChannelsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChannelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChannelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChannels page.
func (p *ListChannelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChannelsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListChannels(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListChannels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "ListChannels",
	}
}
