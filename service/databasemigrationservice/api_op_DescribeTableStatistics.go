// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns table statistics on the database migration task, including table name,
// rows inserted, rows updated, and rows deleted. Note that the "last updated"
// column the DMS console only indicates the time that DMS last updated the table
// statistics record for a table. It does not indicate the time of the last update
// to the table.
func (c *Client) DescribeTableStatistics(ctx context.Context, params *DescribeTableStatisticsInput, optFns ...func(*Options)) (*DescribeTableStatisticsOutput, error) {
	if params == nil {
		params = &DescribeTableStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTableStatistics", params, optFns, c.addOperationDescribeTableStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTableStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTableStatisticsInput struct {

	// The Amazon Resource Name (ARN) of the replication task.
	//
	// This member is required.
	ReplicationTaskArn *string

	// Filters applied to table statistics. Valid filter names: schema-name |
	// table-name | table-state A combination of filters creates an AND condition where
	// each record matches all specified filters.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 500.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeTableStatisticsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The Amazon Resource Name (ARN) of the replication task.
	ReplicationTaskArn *string

	// The table statistics.
	TableStatistics []types.TableStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTableStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTableStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTableStatistics{}, middleware.After)
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
	if err = addOpDescribeTableStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTableStatistics(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeTableStatistics",
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

// DescribeTableStatisticsAPIClient is a client that implements the
// DescribeTableStatistics operation.
type DescribeTableStatisticsAPIClient interface {
	DescribeTableStatistics(context.Context, *DescribeTableStatisticsInput, ...func(*Options)) (*DescribeTableStatisticsOutput, error)
}

var _ DescribeTableStatisticsAPIClient = (*Client)(nil)

// DescribeTableStatisticsPaginatorOptions is the paginator options for
// DescribeTableStatistics
type DescribeTableStatisticsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTableStatisticsPaginator is a paginator for DescribeTableStatistics
type DescribeTableStatisticsPaginator struct {
	options   DescribeTableStatisticsPaginatorOptions
	client    DescribeTableStatisticsAPIClient
	params    *DescribeTableStatisticsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTableStatisticsPaginator returns a new
// DescribeTableStatisticsPaginator
func NewDescribeTableStatisticsPaginator(client DescribeTableStatisticsAPIClient, params *DescribeTableStatisticsInput, optFns ...func(*DescribeTableStatisticsPaginatorOptions)) *DescribeTableStatisticsPaginator {
	if params == nil {
		params = &DescribeTableStatisticsInput{}
	}

	options := DescribeTableStatisticsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTableStatisticsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTableStatisticsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTableStatistics page.
func (p *DescribeTableStatisticsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTableStatisticsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeTableStatistics(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeTableStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeTableStatistics",
	}
}
