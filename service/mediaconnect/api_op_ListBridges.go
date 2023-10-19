// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays a list of bridges that are associated with this account and an
// optionally specified Arn. This request returns a paginated result.
func (c *Client) ListBridges(ctx context.Context, params *ListBridgesInput, optFns ...func(*Options)) (*ListBridgesOutput, error) {
	if params == nil {
		params = &ListBridgesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBridges", params, optFns, c.addOperationListBridgesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBridgesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBridgesInput struct {

	// Filter the list results to display only the bridges associated with the
	// selected Amazon Resource Name (ARN).
	FilterArn *string

	// The maximum number of results to return per API request. For example, you
	// submit a ListBridges request with MaxResults set at 5. Although 20 items match
	// your request, the service returns no more than the first 5 items. (The service
	// also returns a NextToken value that you can use to fetch the next batch of
	// results.) The service might return fewer results than the MaxResults value. If
	// MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 10 results per page.
	MaxResults *int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListBridges request with MaxResults set at 5. The service
	// returns the first batch of results (up to 5) and a NextToken value. To see the
	// next batch of results, you can submit the ListBridges request a second time and
	// specify the NextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBridgesOutput struct {

	// A list of bridge summaries.
	Bridges []types.ListedBridge

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListBridges request with MaxResults set at 5. The service
	// returns the first batch of results (up to 5) and a NextToken value. To see the
	// next batch of results, you can submit the ListBridges request a second time and
	// specify the NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBridgesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBridges{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBridges{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBridges(options.Region), middleware.Before); err != nil {
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
		operation: "ListBridges",
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

// ListBridgesAPIClient is a client that implements the ListBridges operation.
type ListBridgesAPIClient interface {
	ListBridges(context.Context, *ListBridgesInput, ...func(*Options)) (*ListBridgesOutput, error)
}

var _ ListBridgesAPIClient = (*Client)(nil)

// ListBridgesPaginatorOptions is the paginator options for ListBridges
type ListBridgesPaginatorOptions struct {
	// The maximum number of results to return per API request. For example, you
	// submit a ListBridges request with MaxResults set at 5. Although 20 items match
	// your request, the service returns no more than the first 5 items. (The service
	// also returns a NextToken value that you can use to fetch the next batch of
	// results.) The service might return fewer results than the MaxResults value. If
	// MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 10 results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBridgesPaginator is a paginator for ListBridges
type ListBridgesPaginator struct {
	options   ListBridgesPaginatorOptions
	client    ListBridgesAPIClient
	params    *ListBridgesInput
	nextToken *string
	firstPage bool
}

// NewListBridgesPaginator returns a new ListBridgesPaginator
func NewListBridgesPaginator(client ListBridgesAPIClient, params *ListBridgesInput, optFns ...func(*ListBridgesPaginatorOptions)) *ListBridgesPaginator {
	if params == nil {
		params = &ListBridgesInput{}
	}

	options := ListBridgesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBridgesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBridgesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBridges page.
func (p *ListBridgesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBridgesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListBridges(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBridges(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "ListBridges",
	}
}
