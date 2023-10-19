// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the specified metric filters. You can list all of the metric filters or
// filter the results by log name, prefix, metric name, or metric namespace. The
// results are ASCII-sorted by filter name.
func (c *Client) DescribeMetricFilters(ctx context.Context, params *DescribeMetricFiltersInput, optFns ...func(*Options)) (*DescribeMetricFiltersOutput, error) {
	if params == nil {
		params = &DescribeMetricFiltersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMetricFilters", params, optFns, c.addOperationDescribeMetricFiltersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMetricFiltersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMetricFiltersInput struct {

	// The prefix to match. CloudWatch Logs uses the value that you set here only if
	// you also include the logGroupName parameter in your request.
	FilterNamePrefix *string

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int32

	// The name of the log group.
	LogGroupName *string

	// Filters results to include only those with the specified metric name. If you
	// include this parameter in your request, you must also include the
	// metricNamespace parameter.
	MetricName *string

	// Filters results to include only those in the specified namespace. If you
	// include this parameter in your request, you must also include the metricName
	// parameter.
	MetricNamespace *string

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeMetricFiltersOutput struct {

	// The metric filters.
	MetricFilters []types.MetricFilter

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMetricFiltersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMetricFilters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMetricFilters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMetricFilters(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeMetricFilters",
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

// DescribeMetricFiltersAPIClient is a client that implements the
// DescribeMetricFilters operation.
type DescribeMetricFiltersAPIClient interface {
	DescribeMetricFilters(context.Context, *DescribeMetricFiltersInput, ...func(*Options)) (*DescribeMetricFiltersOutput, error)
}

var _ DescribeMetricFiltersAPIClient = (*Client)(nil)

// DescribeMetricFiltersPaginatorOptions is the paginator options for
// DescribeMetricFilters
type DescribeMetricFiltersPaginatorOptions struct {
	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMetricFiltersPaginator is a paginator for DescribeMetricFilters
type DescribeMetricFiltersPaginator struct {
	options   DescribeMetricFiltersPaginatorOptions
	client    DescribeMetricFiltersAPIClient
	params    *DescribeMetricFiltersInput
	nextToken *string
	firstPage bool
}

// NewDescribeMetricFiltersPaginator returns a new DescribeMetricFiltersPaginator
func NewDescribeMetricFiltersPaginator(client DescribeMetricFiltersAPIClient, params *DescribeMetricFiltersInput, optFns ...func(*DescribeMetricFiltersPaginatorOptions)) *DescribeMetricFiltersPaginator {
	if params == nil {
		params = &DescribeMetricFiltersInput{}
	}

	options := DescribeMetricFiltersPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMetricFiltersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMetricFiltersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMetricFilters page.
func (p *DescribeMetricFiltersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMetricFiltersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeMetricFilters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMetricFilters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeMetricFilters",
	}
}
