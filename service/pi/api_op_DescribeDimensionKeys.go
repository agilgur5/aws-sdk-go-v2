// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// For a specific time period, retrieve the top N dimension keys for a metric.
// Each response element returns a maximum of 500 bytes. For larger elements, such
// as SQL statements, only the first 500 bytes are returned.
func (c *Client) DescribeDimensionKeys(ctx context.Context, params *DescribeDimensionKeysInput, optFns ...func(*Options)) (*DescribeDimensionKeysOutput, error) {
	if params == nil {
		params = &DescribeDimensionKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDimensionKeys", params, optFns, c.addOperationDescribeDimensionKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDimensionKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDimensionKeysInput struct {

	// The date and time specifying the end of the requested time series data. The
	// value specified is exclusive, which means that data points less than (but not
	// equal to) EndTime are returned. The value for EndTime must be later than the
	// value for StartTime .
	//
	// This member is required.
	EndTime *time.Time

	// A specification for how to aggregate the data points from a query result. You
	// must specify a valid dimension group. Performance Insights returns all
	// dimensions within this group, unless you provide the names of specific
	// dimensions within this group. You can also request that Performance Insights
	// return a limited number of values for a dimension.
	//
	// This member is required.
	GroupBy *types.DimensionGroup

	// An immutable, Amazon Web Services Region-unique identifier for a data source.
	// Performance Insights gathers metrics from this data source. To use an Amazon RDS
	// instance as a data source, you specify its DbiResourceId value. For example,
	// specify db-FAIHNTYBKTGAUSUZQYPDS2GW4A .
	//
	// This member is required.
	Identifier *string

	// The name of a Performance Insights metric to be measured. Valid values for
	// Metric are:
	//   - db.load.avg - A scaled representation of the number of active sessions for
	//   the database engine.
	//   - db.sampledload.avg - The raw number of active sessions for the database
	//   engine.
	// If the number of active sessions is less than an internal Performance Insights
	// threshold, db.load.avg and db.sampledload.avg are the same value. If the number
	// of active sessions is greater than the internal threshold, Performance Insights
	// samples the active sessions, with db.load.avg showing the scaled values,
	// db.sampledload.avg showing the raw values, and db.sampledload.avg less than
	// db.load.avg . For most use cases, you can query db.load.avg only.
	//
	// This member is required.
	Metric *string

	// The Amazon Web Services service for which Performance Insights will return
	// metrics. Valid values are as follows:
	//   - RDS
	//   - DOCDB
	//
	// This member is required.
	ServiceType types.ServiceType

	// The date and time specifying the beginning of the requested time series data.
	// You must specify a StartTime within the past 7 days. The value specified is
	// inclusive, which means that data points equal to or greater than StartTime are
	// returned. The value for StartTime must be earlier than the value for EndTime .
	//
	// This member is required.
	StartTime *time.Time

	// Additional metrics for the top N dimension keys. If the specified dimension
	// group in the GroupBy parameter is db.sql_tokenized , you can specify per-SQL
	// metrics to get the values for the top N SQL digests. The response syntax is as
	// follows: "AdditionalMetrics" : { "string" : "string" } .
	AdditionalMetrics []string

	// One or more filters to apply in the request. Restrictions:
	//   - Any number of filters by the same dimension, as specified in the GroupBy or
	//   Partition parameters.
	//   - A single filter for any other dimension in this dimension group.
	Filter map[string]string

	// The maximum number of items to return in the response. If more items exist than
	// the specified MaxRecords value, a pagination token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int32

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to the
	// value specified by MaxRecords .
	NextToken *string

	// For each dimension specified in GroupBy , specify a secondary dimension to
	// further subdivide the partition keys in the response.
	PartitionBy *types.DimensionGroup

	// The granularity, in seconds, of the data points returned from Performance
	// Insights. A period can be as short as one second, or as long as one day (86400
	// seconds). Valid values are:
	//   - 1 (one second)
	//   - 60 (one minute)
	//   - 300 (five minutes)
	//   - 3600 (one hour)
	//   - 86400 (twenty-four hours)
	// If you don't specify PeriodInSeconds , then Performance Insights chooses a value
	// for you, with a goal of returning roughly 100-200 data points in the response.
	PeriodInSeconds *int32

	noSmithyDocumentSerde
}

type DescribeDimensionKeysOutput struct {

	// The end time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds ). AlignedEndTime will be greater
	// than or equal to the value of the user-specified Endtime .
	AlignedEndTime *time.Time

	// The start time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds ). AlignedStartTime will be less than
	// or equal to the value of the user-specified StartTime .
	AlignedStartTime *time.Time

	// The dimension keys that were requested.
	Keys []types.DimensionKeyDescription

	// A pagination token that indicates the response didn’t return all available
	// records because MaxRecords was specified in the previous request. To get the
	// remaining records, specify NextToken in a separate request with this value.
	NextToken *string

	// If PartitionBy was present in the request, PartitionKeys contains the breakdown
	// of dimension keys by the specified partitions.
	PartitionKeys []types.ResponsePartitionKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDimensionKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDimensionKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDimensionKeys{}, middleware.After)
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
	if err = addOpDescribeDimensionKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDimensionKeys(options.Region), middleware.Before); err != nil {
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

// DescribeDimensionKeysAPIClient is a client that implements the
// DescribeDimensionKeys operation.
type DescribeDimensionKeysAPIClient interface {
	DescribeDimensionKeys(context.Context, *DescribeDimensionKeysInput, ...func(*Options)) (*DescribeDimensionKeysOutput, error)
}

var _ DescribeDimensionKeysAPIClient = (*Client)(nil)

// DescribeDimensionKeysPaginatorOptions is the paginator options for
// DescribeDimensionKeys
type DescribeDimensionKeysPaginatorOptions struct {
	// The maximum number of items to return in the response. If more items exist than
	// the specified MaxRecords value, a pagination token is included in the response
	// so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDimensionKeysPaginator is a paginator for DescribeDimensionKeys
type DescribeDimensionKeysPaginator struct {
	options   DescribeDimensionKeysPaginatorOptions
	client    DescribeDimensionKeysAPIClient
	params    *DescribeDimensionKeysInput
	nextToken *string
	firstPage bool
}

// NewDescribeDimensionKeysPaginator returns a new DescribeDimensionKeysPaginator
func NewDescribeDimensionKeysPaginator(client DescribeDimensionKeysAPIClient, params *DescribeDimensionKeysInput, optFns ...func(*DescribeDimensionKeysPaginatorOptions)) *DescribeDimensionKeysPaginator {
	if params == nil {
		params = &DescribeDimensionKeysInput{}
	}

	options := DescribeDimensionKeysPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDimensionKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDimensionKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDimensionKeys page.
func (p *DescribeDimensionKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDimensionKeysOutput, error) {
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

	result, err := p.client.DescribeDimensionKeys(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDimensionKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pi",
		OperationName: "DescribeDimensionKeys",
	}
}
