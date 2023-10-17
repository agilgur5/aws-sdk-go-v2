// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of all the installed collectors.
func (c *Client) ListCollectors(ctx context.Context, params *ListCollectorsInput, optFns ...func(*Options)) (*ListCollectorsOutput, error) {
	if params == nil {
		params = &ListCollectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCollectors", params, optFns, c.addOperationListCollectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCollectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCollectorsInput struct {

	// The maximum number of items to include in the response. The maximum value is
	// 100.
	MaxResults *int32

	// The token from a previous call that you use to retrieve the next set of
	// results. For example, if a previous call to this action returned 100 items, but
	// you set maxResults to 10. You'll receive a set of 10 results along with a
	// token. You then use the returned token to retrieve the next set of 10.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCollectorsOutput struct {

	// The list of all the installed collectors.
	Collectors []types.Collector

	// The token you use to retrieve the next set of results, or null if there are no
	// more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCollectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCollectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCollectors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCollectors(options.Region), middleware.Before); err != nil {
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

// ListCollectorsAPIClient is a client that implements the ListCollectors
// operation.
type ListCollectorsAPIClient interface {
	ListCollectors(context.Context, *ListCollectorsInput, ...func(*Options)) (*ListCollectorsOutput, error)
}

var _ ListCollectorsAPIClient = (*Client)(nil)

// ListCollectorsPaginatorOptions is the paginator options for ListCollectors
type ListCollectorsPaginatorOptions struct {
	// The maximum number of items to include in the response. The maximum value is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCollectorsPaginator is a paginator for ListCollectors
type ListCollectorsPaginator struct {
	options   ListCollectorsPaginatorOptions
	client    ListCollectorsAPIClient
	params    *ListCollectorsInput
	nextToken *string
	firstPage bool
}

// NewListCollectorsPaginator returns a new ListCollectorsPaginator
func NewListCollectorsPaginator(client ListCollectorsAPIClient, params *ListCollectorsInput, optFns ...func(*ListCollectorsPaginatorOptions)) *ListCollectorsPaginator {
	if params == nil {
		params = &ListCollectorsInput{}
	}

	options := ListCollectorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCollectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCollectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCollectors page.
func (p *ListCollectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCollectorsOutput, error) {
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

	result, err := p.client.ListCollectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCollectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-strategy",
		OperationName: "ListCollectors",
	}
}
