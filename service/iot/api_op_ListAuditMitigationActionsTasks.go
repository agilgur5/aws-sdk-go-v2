// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a list of audit mitigation action tasks that match the specified filters.
// Requires permission to access the ListAuditMitigationActionsTasks (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListAuditMitigationActionsTasks(ctx context.Context, params *ListAuditMitigationActionsTasksInput, optFns ...func(*Options)) (*ListAuditMitigationActionsTasksOutput, error) {
	if params == nil {
		params = &ListAuditMitigationActionsTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAuditMitigationActionsTasks", params, optFns, c.addOperationListAuditMitigationActionsTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAuditMitigationActionsTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAuditMitigationActionsTasksInput struct {

	// Specify this filter to limit results to tasks that were completed or canceled
	// on or before a specific date and time.
	//
	// This member is required.
	EndTime *time.Time

	// Specify this filter to limit results to tasks that began on or after a specific
	// date and time.
	//
	// This member is required.
	StartTime *time.Time

	// Specify this filter to limit results to tasks that were applied to results for
	// a specific audit.
	AuditTaskId *string

	// Specify this filter to limit results to tasks that were applied to a specific
	// audit finding.
	FindingId *string

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// Specify this filter to limit results to tasks that are in a specific state.
	TaskStatus types.AuditMitigationActionsTaskStatus

	noSmithyDocumentSerde
}

type ListAuditMitigationActionsTasksOutput struct {

	// The token for the next set of results.
	NextToken *string

	// The collection of audit mitigation tasks that matched the filter criteria.
	Tasks []types.AuditMitigationActionsTaskMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAuditMitigationActionsTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAuditMitigationActionsTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAuditMitigationActionsTasks{}, middleware.After)
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
	if err = addOpListAuditMitigationActionsTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAuditMitigationActionsTasks(options.Region), middleware.Before); err != nil {
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

// ListAuditMitigationActionsTasksAPIClient is a client that implements the
// ListAuditMitigationActionsTasks operation.
type ListAuditMitigationActionsTasksAPIClient interface {
	ListAuditMitigationActionsTasks(context.Context, *ListAuditMitigationActionsTasksInput, ...func(*Options)) (*ListAuditMitigationActionsTasksOutput, error)
}

var _ ListAuditMitigationActionsTasksAPIClient = (*Client)(nil)

// ListAuditMitigationActionsTasksPaginatorOptions is the paginator options for
// ListAuditMitigationActionsTasks
type ListAuditMitigationActionsTasksPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAuditMitigationActionsTasksPaginator is a paginator for
// ListAuditMitigationActionsTasks
type ListAuditMitigationActionsTasksPaginator struct {
	options   ListAuditMitigationActionsTasksPaginatorOptions
	client    ListAuditMitigationActionsTasksAPIClient
	params    *ListAuditMitigationActionsTasksInput
	nextToken *string
	firstPage bool
}

// NewListAuditMitigationActionsTasksPaginator returns a new
// ListAuditMitigationActionsTasksPaginator
func NewListAuditMitigationActionsTasksPaginator(client ListAuditMitigationActionsTasksAPIClient, params *ListAuditMitigationActionsTasksInput, optFns ...func(*ListAuditMitigationActionsTasksPaginatorOptions)) *ListAuditMitigationActionsTasksPaginator {
	if params == nil {
		params = &ListAuditMitigationActionsTasksInput{}
	}

	options := ListAuditMitigationActionsTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAuditMitigationActionsTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAuditMitigationActionsTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAuditMitigationActionsTasks page.
func (p *ListAuditMitigationActionsTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAuditMitigationActionsTasksOutput, error) {
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

	result, err := p.client.ListAuditMitigationActionsTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAuditMitigationActionsTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "ListAuditMitigationActionsTasks",
	}
}
