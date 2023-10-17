// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of change sets owned by the account being used to make the
// call. You can filter this list by providing any combination of entityId ,
// ChangeSetName , and status. If you provide more than one filter, the API
// operation applies a logical AND between the filters. You can describe a change
// during the 60-day request history retention period for API calls.
func (c *Client) ListChangeSets(ctx context.Context, params *ListChangeSetsInput, optFns ...func(*Options)) (*ListChangeSetsOutput, error) {
	if params == nil {
		params = &ListChangeSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChangeSets", params, optFns, c.addOperationListChangeSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChangeSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChangeSetsInput struct {

	// The catalog related to the request. Fixed value: AWSMarketplace
	//
	// This member is required.
	Catalog *string

	// An array of filter objects.
	FilterList []types.Filter

	// The maximum number of results returned by a single call. This value must be
	// provided in the next call to retrieve the next set of results. By default, this
	// value is 20.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// An object that contains two attributes, SortBy and SortOrder .
	Sort *types.Sort

	noSmithyDocumentSerde
}

type ListChangeSetsOutput struct {

	// Array of ChangeSetSummaryListItem objects.
	ChangeSetSummaryList []types.ChangeSetSummaryListItem

	// The value of the next token, if it exists. Null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChangeSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChangeSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChangeSets{}, middleware.After)
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
	if err = addOpListChangeSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChangeSets(options.Region), middleware.Before); err != nil {
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

// ListChangeSetsAPIClient is a client that implements the ListChangeSets
// operation.
type ListChangeSetsAPIClient interface {
	ListChangeSets(context.Context, *ListChangeSetsInput, ...func(*Options)) (*ListChangeSetsOutput, error)
}

var _ ListChangeSetsAPIClient = (*Client)(nil)

// ListChangeSetsPaginatorOptions is the paginator options for ListChangeSets
type ListChangeSetsPaginatorOptions struct {
	// The maximum number of results returned by a single call. This value must be
	// provided in the next call to retrieve the next set of results. By default, this
	// value is 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChangeSetsPaginator is a paginator for ListChangeSets
type ListChangeSetsPaginator struct {
	options   ListChangeSetsPaginatorOptions
	client    ListChangeSetsAPIClient
	params    *ListChangeSetsInput
	nextToken *string
	firstPage bool
}

// NewListChangeSetsPaginator returns a new ListChangeSetsPaginator
func NewListChangeSetsPaginator(client ListChangeSetsAPIClient, params *ListChangeSetsInput, optFns ...func(*ListChangeSetsPaginatorOptions)) *ListChangeSetsPaginator {
	if params == nil {
		params = &ListChangeSetsInput{}
	}

	options := ListChangeSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChangeSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChangeSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChangeSets page.
func (p *ListChangeSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChangeSetsOutput, error) {
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

	result, err := p.client.ListChangeSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListChangeSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "ListChangeSets",
	}
}
