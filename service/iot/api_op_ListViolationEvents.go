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

// Lists the Device Defender security profile violations discovered during the
// given time period. You can use filters to limit the results to those alerts
// issued for a particular security profile, behavior, or thing (device). Requires
// permission to access the ListViolationEvents (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListViolationEvents(ctx context.Context, params *ListViolationEventsInput, optFns ...func(*Options)) (*ListViolationEventsOutput, error) {
	if params == nil {
		params = &ListViolationEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListViolationEvents", params, optFns, c.addOperationListViolationEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListViolationEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListViolationEventsInput struct {

	// The end time for the alerts to be listed.
	//
	// This member is required.
	EndTime *time.Time

	// The start time for the alerts to be listed.
	//
	// This member is required.
	StartTime *time.Time

	// The criteria for a behavior.
	BehaviorCriteriaType types.BehaviorCriteriaType

	// A list of all suppressed alerts.
	ListSuppressedAlerts *bool

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// A filter to limit results to those alerts generated by the specified security
	// profile.
	SecurityProfileName *string

	// A filter to limit results to those alerts caused by the specified thing.
	ThingName *string

	// The verification state of the violation (detect alarm).
	VerificationState types.VerificationState

	noSmithyDocumentSerde
}

type ListViolationEventsOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// The security profile violation alerts issued for this account during the given
	// time period, potentially filtered by security profile, behavior violated, or
	// thing (device) violating.
	ViolationEvents []types.ViolationEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListViolationEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListViolationEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListViolationEvents{}, middleware.After)
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
	if err = addOpListViolationEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListViolationEvents(options.Region), middleware.Before); err != nil {
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

// ListViolationEventsAPIClient is a client that implements the
// ListViolationEvents operation.
type ListViolationEventsAPIClient interface {
	ListViolationEvents(context.Context, *ListViolationEventsInput, ...func(*Options)) (*ListViolationEventsOutput, error)
}

var _ ListViolationEventsAPIClient = (*Client)(nil)

// ListViolationEventsPaginatorOptions is the paginator options for
// ListViolationEvents
type ListViolationEventsPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListViolationEventsPaginator is a paginator for ListViolationEvents
type ListViolationEventsPaginator struct {
	options   ListViolationEventsPaginatorOptions
	client    ListViolationEventsAPIClient
	params    *ListViolationEventsInput
	nextToken *string
	firstPage bool
}

// NewListViolationEventsPaginator returns a new ListViolationEventsPaginator
func NewListViolationEventsPaginator(client ListViolationEventsAPIClient, params *ListViolationEventsInput, optFns ...func(*ListViolationEventsPaginatorOptions)) *ListViolationEventsPaginator {
	if params == nil {
		params = &ListViolationEventsInput{}
	}

	options := ListViolationEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListViolationEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListViolationEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListViolationEvents page.
func (p *ListViolationEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListViolationEventsOutput, error) {
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

	result, err := p.client.ListViolationEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListViolationEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "ListViolationEvents",
	}
}
