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
	"time"
)

// Returns list of all monitoring schedules.
func (c *Client) ListMonitoringSchedules(ctx context.Context, params *ListMonitoringSchedulesInput, optFns ...func(*Options)) (*ListMonitoringSchedulesOutput, error) {
	if params == nil {
		params = &ListMonitoringSchedulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMonitoringSchedules", params, optFns, c.addOperationListMonitoringSchedulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMonitoringSchedulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMonitoringSchedulesInput struct {

	// A filter that returns only monitoring schedules created after a specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only monitoring schedules created before a specified time.
	CreationTimeBefore *time.Time

	// Name of a specific endpoint to fetch schedules for.
	EndpointName *string

	// A filter that returns only monitoring schedules modified after a specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only monitoring schedules modified before a specified
	// time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of jobs to return in the response. The default value is 10.
	MaxResults *int32

	// Gets a list of the monitoring schedules for the specified monitoring job
	// definition.
	MonitoringJobDefinitionName *string

	// A filter that returns only the monitoring schedules for the specified
	// monitoring type.
	MonitoringTypeEquals types.MonitoringType

	// Filter for monitoring schedules whose name contains a specified string.
	NameContains *string

	// The token returned if the response is truncated. To retrieve the next set of
	// job executions, use it in the next request.
	NextToken *string

	// Whether to sort the results by the Status , CreationTime , or ScheduledTime
	// field. The default is CreationTime .
	SortBy types.MonitoringScheduleSortKey

	// Whether to sort the results in Ascending or Descending order. The default is
	// Descending .
	SortOrder types.SortOrder

	// A filter that returns only monitoring schedules modified before a specified
	// time.
	StatusEquals types.ScheduleStatus

	noSmithyDocumentSerde
}

type ListMonitoringSchedulesOutput struct {

	// A JSON array in which each element is a summary for a monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleSummaries []types.MonitoringScheduleSummary

	// The token returned if the response is truncated. To retrieve the next set of
	// job executions, use it in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMonitoringSchedulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMonitoringSchedules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMonitoringSchedules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMonitoringSchedules(options.Region), middleware.Before); err != nil {
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
		operation: "ListMonitoringSchedules",
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

// ListMonitoringSchedulesAPIClient is a client that implements the
// ListMonitoringSchedules operation.
type ListMonitoringSchedulesAPIClient interface {
	ListMonitoringSchedules(context.Context, *ListMonitoringSchedulesInput, ...func(*Options)) (*ListMonitoringSchedulesOutput, error)
}

var _ ListMonitoringSchedulesAPIClient = (*Client)(nil)

// ListMonitoringSchedulesPaginatorOptions is the paginator options for
// ListMonitoringSchedules
type ListMonitoringSchedulesPaginatorOptions struct {
	// The maximum number of jobs to return in the response. The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMonitoringSchedulesPaginator is a paginator for ListMonitoringSchedules
type ListMonitoringSchedulesPaginator struct {
	options   ListMonitoringSchedulesPaginatorOptions
	client    ListMonitoringSchedulesAPIClient
	params    *ListMonitoringSchedulesInput
	nextToken *string
	firstPage bool
}

// NewListMonitoringSchedulesPaginator returns a new
// ListMonitoringSchedulesPaginator
func NewListMonitoringSchedulesPaginator(client ListMonitoringSchedulesAPIClient, params *ListMonitoringSchedulesInput, optFns ...func(*ListMonitoringSchedulesPaginatorOptions)) *ListMonitoringSchedulesPaginator {
	if params == nil {
		params = &ListMonitoringSchedulesInput{}
	}

	options := ListMonitoringSchedulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMonitoringSchedulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMonitoringSchedulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMonitoringSchedules page.
func (p *ListMonitoringSchedulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMonitoringSchedulesOutput, error) {
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

	result, err := p.client.ListMonitoringSchedules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMonitoringSchedules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListMonitoringSchedules",
	}
}
