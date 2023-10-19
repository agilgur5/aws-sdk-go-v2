// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about all event data stores in the account, in the current
// Region.
func (c *Client) ListEventDataStores(ctx context.Context, params *ListEventDataStoresInput, optFns ...func(*Options)) (*ListEventDataStoresOutput, error) {
	if params == nil {
		params = &ListEventDataStoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventDataStores", params, optFns, c.addOperationListEventDataStoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventDataStoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventDataStoresInput struct {

	// The maximum number of event data stores to display on a single page.
	MaxResults *int32

	// A token you can use to get the next page of event data store results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEventDataStoresOutput struct {

	// Contains information about event data stores in the account, in the current
	// Region.
	EventDataStores []types.EventDataStore

	// A token you can use to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventDataStoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEventDataStores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEventDataStores{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventDataStores(options.Region), middleware.Before); err != nil {
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
		operation: "ListEventDataStores",
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

// ListEventDataStoresAPIClient is a client that implements the
// ListEventDataStores operation.
type ListEventDataStoresAPIClient interface {
	ListEventDataStores(context.Context, *ListEventDataStoresInput, ...func(*Options)) (*ListEventDataStoresOutput, error)
}

var _ ListEventDataStoresAPIClient = (*Client)(nil)

// ListEventDataStoresPaginatorOptions is the paginator options for
// ListEventDataStores
type ListEventDataStoresPaginatorOptions struct {
	// The maximum number of event data stores to display on a single page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventDataStoresPaginator is a paginator for ListEventDataStores
type ListEventDataStoresPaginator struct {
	options   ListEventDataStoresPaginatorOptions
	client    ListEventDataStoresAPIClient
	params    *ListEventDataStoresInput
	nextToken *string
	firstPage bool
}

// NewListEventDataStoresPaginator returns a new ListEventDataStoresPaginator
func NewListEventDataStoresPaginator(client ListEventDataStoresAPIClient, params *ListEventDataStoresInput, optFns ...func(*ListEventDataStoresPaginatorOptions)) *ListEventDataStoresPaginator {
	if params == nil {
		params = &ListEventDataStoresInput{}
	}

	options := ListEventDataStoresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventDataStoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventDataStoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventDataStores page.
func (p *ListEventDataStoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventDataStoresOutput, error) {
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

	result, err := p.client.ListEventDataStores(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEventDataStores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "ListEventDataStores",
	}
}
