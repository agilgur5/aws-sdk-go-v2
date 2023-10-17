// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the anomalies that belong to an insight that you specify
// using its ID.
func (c *Client) ListAnomaliesForInsight(ctx context.Context, params *ListAnomaliesForInsightInput, optFns ...func(*Options)) (*ListAnomaliesForInsightOutput, error) {
	if params == nil {
		params = &ListAnomaliesForInsightInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnomaliesForInsight", params, optFns, c.addOperationListAnomaliesForInsightMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnomaliesForInsightOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnomaliesForInsightInput struct {

	// The ID of the insight. The returned anomalies belong to this insight.
	//
	// This member is required.
	InsightId *string

	// The ID of the Amazon Web Services account.
	AccountId *string

	// Specifies one or more service names that are used to list anomalies.
	Filters *types.ListAnomaliesForInsightFilters

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// A time range used to specify when the requested anomalies started. All returned
	// anomalies started during this time range.
	StartTimeRange *types.StartTimeRange

	noSmithyDocumentSerde
}

type ListAnomaliesForInsightOutput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// An array of ProactiveAnomalySummary objects that represent the requested
	// anomalies
	ProactiveAnomalies []types.ProactiveAnomalySummary

	// An array of ReactiveAnomalySummary objects that represent the requested
	// anomalies
	ReactiveAnomalies []types.ReactiveAnomalySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnomaliesForInsightMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnomaliesForInsight{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnomaliesForInsight{}, middleware.After)
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
	if err = addOpListAnomaliesForInsightValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnomaliesForInsight(options.Region), middleware.Before); err != nil {
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

// ListAnomaliesForInsightAPIClient is a client that implements the
// ListAnomaliesForInsight operation.
type ListAnomaliesForInsightAPIClient interface {
	ListAnomaliesForInsight(context.Context, *ListAnomaliesForInsightInput, ...func(*Options)) (*ListAnomaliesForInsightOutput, error)
}

var _ ListAnomaliesForInsightAPIClient = (*Client)(nil)

// ListAnomaliesForInsightPaginatorOptions is the paginator options for
// ListAnomaliesForInsight
type ListAnomaliesForInsightPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnomaliesForInsightPaginator is a paginator for ListAnomaliesForInsight
type ListAnomaliesForInsightPaginator struct {
	options   ListAnomaliesForInsightPaginatorOptions
	client    ListAnomaliesForInsightAPIClient
	params    *ListAnomaliesForInsightInput
	nextToken *string
	firstPage bool
}

// NewListAnomaliesForInsightPaginator returns a new
// ListAnomaliesForInsightPaginator
func NewListAnomaliesForInsightPaginator(client ListAnomaliesForInsightAPIClient, params *ListAnomaliesForInsightInput, optFns ...func(*ListAnomaliesForInsightPaginatorOptions)) *ListAnomaliesForInsightPaginator {
	if params == nil {
		params = &ListAnomaliesForInsightInput{}
	}

	options := ListAnomaliesForInsightPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnomaliesForInsightPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnomaliesForInsightPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnomaliesForInsight page.
func (p *ListAnomaliesForInsightPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnomaliesForInsightOutput, error) {
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

	result, err := p.client.ListAnomaliesForInsight(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAnomaliesForInsight(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "ListAnomaliesForInsight",
	}
}
