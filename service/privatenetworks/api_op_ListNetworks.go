// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/privatenetworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists networks. Add filters to your request to return a more specific list of
// results. Use filters to match the status of the network.
func (c *Client) ListNetworks(ctx context.Context, params *ListNetworksInput, optFns ...func(*Options)) (*ListNetworksOutput, error) {
	if params == nil {
		params = &ListNetworksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNetworks", params, optFns, c.addOperationListNetworksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNetworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNetworksInput struct {

	// The filters.
	//   - STATUS - The status ( AVAILABLE | CREATED | DELETED | DEPROVISIONING |
	//   PROVISIONING ).
	// Filter values are case sensitive. If you specify multiple values for a filter,
	// the values are joined with an OR , and the request returns all results that
	// match any of the specified values.
	Filters map[string][]string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	StartToken *string

	noSmithyDocumentSerde
}

type ListNetworksOutput struct {

	// The networks.
	Networks []types.Network

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNetworksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNetworks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNetworks(options.Region), middleware.Before); err != nil {
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

// ListNetworksAPIClient is a client that implements the ListNetworks operation.
type ListNetworksAPIClient interface {
	ListNetworks(context.Context, *ListNetworksInput, ...func(*Options)) (*ListNetworksOutput, error)
}

var _ ListNetworksAPIClient = (*Client)(nil)

// ListNetworksPaginatorOptions is the paginator options for ListNetworks
type ListNetworksPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNetworksPaginator is a paginator for ListNetworks
type ListNetworksPaginator struct {
	options   ListNetworksPaginatorOptions
	client    ListNetworksAPIClient
	params    *ListNetworksInput
	nextToken *string
	firstPage bool
}

// NewListNetworksPaginator returns a new ListNetworksPaginator
func NewListNetworksPaginator(client ListNetworksAPIClient, params *ListNetworksInput, optFns ...func(*ListNetworksPaginatorOptions)) *ListNetworksPaginator {
	if params == nil {
		params = &ListNetworksInput{}
	}

	options := ListNetworksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNetworksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.StartToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNetworksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNetworks page.
func (p *ListNetworksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNetworksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.StartToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListNetworks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNetworks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "private-networks",
		OperationName: "ListNetworks",
	}
}
