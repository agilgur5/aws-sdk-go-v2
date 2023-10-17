// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of monitors created with the CreateMonitor operation and
// CreateAutoPredictor operation. For each monitor resource, this operation returns
// of a summary of its properties, including its Amazon Resource Name (ARN). You
// can retrieve a complete set of properties of a monitor resource by specify the
// monitor's ARN in the DescribeMonitor operation.
func (c *Client) ListMonitors(ctx context.Context, params *ListMonitorsInput, optFns ...func(*Options)) (*ListMonitorsOutput, error) {
	if params == nil {
		params = &ListMonitorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMonitors", params, optFns, c.addOperationListMonitorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMonitorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMonitorsInput struct {

	// An array of filters. For each filter, provide a condition and a match
	// statement. The condition is either IS or IS_NOT , which specifies whether to
	// include or exclude the resources that match the statement from the list. The
	// match statement consists of a key and a value. Filter properties
	//   - Condition - The condition to apply. Valid values are IS and IS_NOT .
	//   - Key - The name of the parameter to filter on. The only valid value is Status
	//   .
	//   - Value - The value to match.
	// For example, to list all monitors who's status is ACTIVE, you would specify:
	// "Filters": [ { "Condition": "IS", "Key": "Status", "Value": "ACTIVE" } ]
	Filters []types.Filter

	// The maximum number of monitors to include in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken . To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMonitorsOutput struct {

	// An array of objects that summarize each monitor's properties.
	Monitors []types.MonitorSummary

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMonitorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMonitors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMonitors{}, middleware.After)
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
	if err = addOpListMonitorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMonitors(options.Region), middleware.Before); err != nil {
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

// ListMonitorsAPIClient is a client that implements the ListMonitors operation.
type ListMonitorsAPIClient interface {
	ListMonitors(context.Context, *ListMonitorsInput, ...func(*Options)) (*ListMonitorsOutput, error)
}

var _ ListMonitorsAPIClient = (*Client)(nil)

// ListMonitorsPaginatorOptions is the paginator options for ListMonitors
type ListMonitorsPaginatorOptions struct {
	// The maximum number of monitors to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMonitorsPaginator is a paginator for ListMonitors
type ListMonitorsPaginator struct {
	options   ListMonitorsPaginatorOptions
	client    ListMonitorsAPIClient
	params    *ListMonitorsInput
	nextToken *string
	firstPage bool
}

// NewListMonitorsPaginator returns a new ListMonitorsPaginator
func NewListMonitorsPaginator(client ListMonitorsAPIClient, params *ListMonitorsInput, optFns ...func(*ListMonitorsPaginatorOptions)) *ListMonitorsPaginator {
	if params == nil {
		params = &ListMonitorsInput{}
	}

	options := ListMonitorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMonitorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMonitorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMonitors page.
func (p *ListMonitorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMonitorsOutput, error) {
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

	result, err := p.client.ListMonitors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMonitors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListMonitors",
	}
}
