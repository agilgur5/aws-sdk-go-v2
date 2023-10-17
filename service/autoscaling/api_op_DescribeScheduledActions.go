// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about the scheduled actions that haven't run or that have not
// reached their end time. To describe the scaling activities for scheduled actions
// that have already run, call the DescribeScalingActivities API.
func (c *Client) DescribeScheduledActions(ctx context.Context, params *DescribeScheduledActionsInput, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) {
	if params == nil {
		params = &DescribeScheduledActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScheduledActions", params, optFns, c.addOperationDescribeScheduledActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScheduledActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScheduledActionsInput struct {

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string

	// The latest scheduled start time to return. If scheduled action names are
	// provided, this property is ignored.
	EndTime *time.Time

	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100 .
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The names of one or more scheduled actions. If you omit this property, all
	// scheduled actions are described. If you specify an unknown scheduled action, it
	// is ignored with no error. Array Members: Maximum number of 50 actions.
	ScheduledActionNames []string

	// The earliest scheduled start time to return. If scheduled action names are
	// provided, this property is ignored.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type DescribeScheduledActionsOutput struct {

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// The scheduled actions.
	ScheduledUpdateGroupActions []types.ScheduledUpdateGroupAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScheduledActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeScheduledActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeScheduledActions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScheduledActions(options.Region), middleware.Before); err != nil {
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

// DescribeScheduledActionsAPIClient is a client that implements the
// DescribeScheduledActions operation.
type DescribeScheduledActionsAPIClient interface {
	DescribeScheduledActions(context.Context, *DescribeScheduledActionsInput, ...func(*Options)) (*DescribeScheduledActionsOutput, error)
}

var _ DescribeScheduledActionsAPIClient = (*Client)(nil)

// DescribeScheduledActionsPaginatorOptions is the paginator options for
// DescribeScheduledActions
type DescribeScheduledActionsPaginatorOptions struct {
	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeScheduledActionsPaginator is a paginator for DescribeScheduledActions
type DescribeScheduledActionsPaginator struct {
	options   DescribeScheduledActionsPaginatorOptions
	client    DescribeScheduledActionsAPIClient
	params    *DescribeScheduledActionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeScheduledActionsPaginator returns a new
// DescribeScheduledActionsPaginator
func NewDescribeScheduledActionsPaginator(client DescribeScheduledActionsAPIClient, params *DescribeScheduledActionsInput, optFns ...func(*DescribeScheduledActionsPaginatorOptions)) *DescribeScheduledActionsPaginator {
	if params == nil {
		params = &DescribeScheduledActionsInput{}
	}

	options := DescribeScheduledActionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeScheduledActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeScheduledActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeScheduledActions page.
func (p *DescribeScheduledActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeScheduledActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeScheduledActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeScheduledActions",
	}
}
