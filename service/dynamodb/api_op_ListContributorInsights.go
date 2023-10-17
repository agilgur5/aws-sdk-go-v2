// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of ContributorInsightsSummary for a table and all its global
// secondary indexes.
func (c *Client) ListContributorInsights(ctx context.Context, params *ListContributorInsightsInput, optFns ...func(*Options)) (*ListContributorInsightsOutput, error) {
	if params == nil {
		params = &ListContributorInsightsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContributorInsights", params, optFns, c.addOperationListContributorInsightsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContributorInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContributorInsightsInput struct {

	// Maximum number of results to return per page.
	MaxResults int32

	// A token to for the desired page, if there is one.
	NextToken *string

	// The name of the table.
	TableName *string

	noSmithyDocumentSerde
}

type ListContributorInsightsOutput struct {

	// A list of ContributorInsightsSummary.
	ContributorInsightsSummaries []types.ContributorInsightsSummary

	// A token to go to the next page if there is one.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContributorInsightsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListContributorInsights{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListContributorInsights{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContributorInsights(options.Region), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
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

// ListContributorInsightsAPIClient is a client that implements the
// ListContributorInsights operation.
type ListContributorInsightsAPIClient interface {
	ListContributorInsights(context.Context, *ListContributorInsightsInput, ...func(*Options)) (*ListContributorInsightsOutput, error)
}

var _ ListContributorInsightsAPIClient = (*Client)(nil)

// ListContributorInsightsPaginatorOptions is the paginator options for
// ListContributorInsights
type ListContributorInsightsPaginatorOptions struct {
	// Maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContributorInsightsPaginator is a paginator for ListContributorInsights
type ListContributorInsightsPaginator struct {
	options   ListContributorInsightsPaginatorOptions
	client    ListContributorInsightsAPIClient
	params    *ListContributorInsightsInput
	nextToken *string
	firstPage bool
}

// NewListContributorInsightsPaginator returns a new
// ListContributorInsightsPaginator
func NewListContributorInsightsPaginator(client ListContributorInsightsAPIClient, params *ListContributorInsightsInput, optFns ...func(*ListContributorInsightsPaginatorOptions)) *ListContributorInsightsPaginator {
	if params == nil {
		params = &ListContributorInsightsInput{}
	}

	options := ListContributorInsightsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContributorInsightsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContributorInsightsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContributorInsights page.
func (p *ListContributorInsightsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContributorInsightsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListContributorInsights(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContributorInsights(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListContributorInsights",
	}
}
