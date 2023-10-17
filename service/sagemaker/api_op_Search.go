// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Finds SageMaker resources that match a search query. Matching resources are
// returned as a list of SearchRecord objects in the response. You can sort the
// search results by any resource property in a ascending or descending order. You
// can query against the following value types: numeric, text, Boolean, and
// timestamp. The Search API may provide access to otherwise restricted data. See
// Amazon SageMaker API Permissions: Actions, Permissions, and Resources Reference (https://docs.aws.amazon.com/sagemaker/latest/dg/api-permissions-reference.html)
// for more information.
func (c *Client) Search(ctx context.Context, params *SearchInput, optFns ...func(*Options)) (*SearchOutput, error) {
	if params == nil {
		params = &SearchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Search", params, optFns, c.addOperationSearchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchInput struct {

	// The name of the SageMaker resource to search for.
	//
	// This member is required.
	Resource types.ResourceType

	// A cross account filter option. When the value is "CrossAccount" the search
	// results will only include resources made discoverable to you from other
	// accounts. When the value is "SameAccount" or null the search results will only
	// include resources from your account. Default is null . For more information on
	// searching for resources made discoverable to your account, see Search
	// discoverable resources (https://docs.aws.amazon.com/sagemaker/latest/dg/feature-store-cross-account-discoverability-use.html)
	// in the SageMaker Developer Guide. The maximum number of ResourceCatalog s
	// viewable is 1000.
	CrossAccountFilterOption types.CrossAccountFilterOption

	// The maximum number of results to return.
	MaxResults *int32

	// If more than MaxResults resources match the specified SearchExpression , the
	// response includes a NextToken . The NextToken can be passed to the next
	// SearchRequest to continue retrieving results.
	NextToken *string

	// A Boolean conditional statement. Resources must satisfy this condition to be
	// included in search results. You must provide at least one subexpression, filter,
	// or nested filter. The maximum number of recursive SubExpressions , NestedFilters
	// , and Filters that can be included in a SearchExpression object is 50.
	SearchExpression *types.SearchExpression

	// The name of the resource property used to sort the SearchResults . The default
	// is LastModifiedTime .
	SortBy *string

	// How SearchResults are ordered. Valid values are Ascending or Descending . The
	// default is Descending .
	SortOrder types.SearchSortOrder

	noSmithyDocumentSerde
}

type SearchOutput struct {

	// If the result of the previous Search request was truncated, the response
	// includes a NextToken. To retrieve the next set of results, use the token in the
	// next request.
	NextToken *string

	// A list of SearchRecord objects.
	Results []types.SearchRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearch{}, middleware.After)
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
	if err = addOpSearchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearch(options.Region), middleware.Before); err != nil {
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

// SearchAPIClient is a client that implements the Search operation.
type SearchAPIClient interface {
	Search(context.Context, *SearchInput, ...func(*Options)) (*SearchOutput, error)
}

var _ SearchAPIClient = (*Client)(nil)

// SearchPaginatorOptions is the paginator options for Search
type SearchPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchPaginator is a paginator for Search
type SearchPaginator struct {
	options   SearchPaginatorOptions
	client    SearchAPIClient
	params    *SearchInput
	nextToken *string
	firstPage bool
}

// NewSearchPaginator returns a new SearchPaginator
func NewSearchPaginator(client SearchAPIClient, params *SearchInput, optFns ...func(*SearchPaginatorOptions)) *SearchPaginator {
	if params == nil {
		params = &SearchInput{}
	}

	options := SearchPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next Search page.
func (p *SearchPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchOutput, error) {
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

	result, err := p.client.Search(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "Search",
	}
}
