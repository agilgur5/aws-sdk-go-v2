// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists statistics about the data collected for each of the sensors that have
// been successfully ingested in the particular dataset. Can also be used to
// retreive Sensor Statistics for a previous ingestion job.
func (c *Client) ListSensorStatistics(ctx context.Context, params *ListSensorStatisticsInput, optFns ...func(*Options)) (*ListSensorStatisticsOutput, error) {
	if params == nil {
		params = &ListSensorStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSensorStatistics", params, optFns, c.addOperationListSensorStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSensorStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSensorStatisticsInput struct {

	// The name of the dataset associated with the list of Sensor Statistics.
	//
	// This member is required.
	DatasetName *string

	// The ingestion job id associated with the list of Sensor Statistics. To get
	// sensor statistics for a particular ingestion job id, both dataset name and
	// ingestion job id must be submitted as inputs.
	IngestionJobId *string

	// Specifies the maximum number of sensors for which to retrieve statistics.
	MaxResults *int32

	// An opaque pagination token indicating where to continue the listing of sensor
	// statistics.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSensorStatisticsOutput struct {

	// An opaque pagination token indicating where to continue the listing of sensor
	// statistics.
	NextToken *string

	// Provides ingestion-based statistics regarding the specified sensor with respect
	// to various validation types, such as whether data exists, the number and
	// percentage of missing values, and the number and percentage of duplicate
	// timestamps.
	SensorStatisticsSummaries []types.SensorStatisticsSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSensorStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListSensorStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListSensorStatistics{}, middleware.After)
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
	if err = addOpListSensorStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSensorStatistics(options.Region), middleware.Before); err != nil {
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

// ListSensorStatisticsAPIClient is a client that implements the
// ListSensorStatistics operation.
type ListSensorStatisticsAPIClient interface {
	ListSensorStatistics(context.Context, *ListSensorStatisticsInput, ...func(*Options)) (*ListSensorStatisticsOutput, error)
}

var _ ListSensorStatisticsAPIClient = (*Client)(nil)

// ListSensorStatisticsPaginatorOptions is the paginator options for
// ListSensorStatistics
type ListSensorStatisticsPaginatorOptions struct {
	// Specifies the maximum number of sensors for which to retrieve statistics.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSensorStatisticsPaginator is a paginator for ListSensorStatistics
type ListSensorStatisticsPaginator struct {
	options   ListSensorStatisticsPaginatorOptions
	client    ListSensorStatisticsAPIClient
	params    *ListSensorStatisticsInput
	nextToken *string
	firstPage bool
}

// NewListSensorStatisticsPaginator returns a new ListSensorStatisticsPaginator
func NewListSensorStatisticsPaginator(client ListSensorStatisticsAPIClient, params *ListSensorStatisticsInput, optFns ...func(*ListSensorStatisticsPaginatorOptions)) *ListSensorStatisticsPaginator {
	if params == nil {
		params = &ListSensorStatisticsInput{}
	}

	options := ListSensorStatisticsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSensorStatisticsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSensorStatisticsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSensorStatistics page.
func (p *ListSensorStatisticsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSensorStatisticsOutput, error) {
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

	result, err := p.client.ListSensorStatistics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSensorStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "ListSensorStatistics",
	}
}
