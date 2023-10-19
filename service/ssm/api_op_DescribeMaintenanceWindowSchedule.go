// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about upcoming executions of a maintenance window.
func (c *Client) DescribeMaintenanceWindowSchedule(ctx context.Context, params *DescribeMaintenanceWindowScheduleInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowScheduleOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowSchedule", params, optFns, c.addOperationDescribeMaintenanceWindowScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowScheduleInput struct {

	// Filters used to limit the range of results. For example, you can limit
	// maintenance window executions to only those scheduled before or after a certain
	// date and time.
	Filters []types.PatchOrchestratorFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The type of resource you want to retrieve information about. For example,
	// INSTANCE .
	ResourceType types.MaintenanceWindowResourceType

	// The managed node ID or key-value pair to retrieve information about.
	Targets []types.Target

	// The ID of the maintenance window to retrieve information about.
	WindowId *string

	noSmithyDocumentSerde
}

type DescribeMaintenanceWindowScheduleOutput struct {

	// The token for the next set of items to return. (You use this token in the next
	// call.)
	NextToken *string

	// Information about maintenance window executions scheduled for the specified
	// time range.
	ScheduledWindowExecutions []types.ScheduledWindowExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMaintenanceWindowScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowSchedule{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowSchedule(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeMaintenanceWindowSchedule",
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

// DescribeMaintenanceWindowScheduleAPIClient is a client that implements the
// DescribeMaintenanceWindowSchedule operation.
type DescribeMaintenanceWindowScheduleAPIClient interface {
	DescribeMaintenanceWindowSchedule(context.Context, *DescribeMaintenanceWindowScheduleInput, ...func(*Options)) (*DescribeMaintenanceWindowScheduleOutput, error)
}

var _ DescribeMaintenanceWindowScheduleAPIClient = (*Client)(nil)

// DescribeMaintenanceWindowSchedulePaginatorOptions is the paginator options for
// DescribeMaintenanceWindowSchedule
type DescribeMaintenanceWindowSchedulePaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowSchedulePaginator is a paginator for
// DescribeMaintenanceWindowSchedule
type DescribeMaintenanceWindowSchedulePaginator struct {
	options   DescribeMaintenanceWindowSchedulePaginatorOptions
	client    DescribeMaintenanceWindowScheduleAPIClient
	params    *DescribeMaintenanceWindowScheduleInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowSchedulePaginator returns a new
// DescribeMaintenanceWindowSchedulePaginator
func NewDescribeMaintenanceWindowSchedulePaginator(client DescribeMaintenanceWindowScheduleAPIClient, params *DescribeMaintenanceWindowScheduleInput, optFns ...func(*DescribeMaintenanceWindowSchedulePaginatorOptions)) *DescribeMaintenanceWindowSchedulePaginator {
	if params == nil {
		params = &DescribeMaintenanceWindowScheduleInput{}
	}

	options := DescribeMaintenanceWindowSchedulePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMaintenanceWindowSchedulePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowSchedulePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMaintenanceWindowSchedule page.
func (p *DescribeMaintenanceWindowSchedulePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowScheduleOutput, error) {
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

	result, err := p.client.DescribeMaintenanceWindowSchedule(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowSchedule",
	}
}
