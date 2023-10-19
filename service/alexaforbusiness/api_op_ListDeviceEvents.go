// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the device event history, including device connection status, for up to
// 30 days.
//
// Deprecated: Alexa For Business is no longer supported
func (c *Client) ListDeviceEvents(ctx context.Context, params *ListDeviceEventsInput, optFns ...func(*Options)) (*ListDeviceEventsOutput, error) {
	if params == nil {
		params = &ListDeviceEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeviceEvents", params, optFns, c.addOperationListDeviceEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeviceEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeviceEventsInput struct {

	// The ARN of a device.
	//
	// This member is required.
	DeviceArn *string

	// The event type to filter device events. If EventType isn't specified, this
	// returns a list of all device events in reverse chronological order. If EventType
	// is specified, this returns a list of device events for that EventType in reverse
	// chronological order.
	EventType types.DeviceEventType

	// The maximum number of results to include in the response. The default value is
	// 50. If more results exist than the specified MaxResults value, a token is
	// included in the response so that the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response only
	// includes results beyond the token, up to the value specified by MaxResults. When
	// the end of results is reached, the response has a value of null.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDeviceEventsOutput struct {

	// The device events requested for the device ARN.
	DeviceEvents []types.DeviceEvent

	// The token returned to indicate that there is more data available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeviceEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDeviceEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeviceEvents{}, middleware.After)
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
	if err = addOpListDeviceEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeviceEvents(options.Region), middleware.Before); err != nil {
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
		operation: "ListDeviceEvents",
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

// ListDeviceEventsAPIClient is a client that implements the ListDeviceEvents
// operation.
type ListDeviceEventsAPIClient interface {
	ListDeviceEvents(context.Context, *ListDeviceEventsInput, ...func(*Options)) (*ListDeviceEventsOutput, error)
}

var _ ListDeviceEventsAPIClient = (*Client)(nil)

// ListDeviceEventsPaginatorOptions is the paginator options for ListDeviceEvents
type ListDeviceEventsPaginatorOptions struct {
	// The maximum number of results to include in the response. The default value is
	// 50. If more results exist than the specified MaxResults value, a token is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeviceEventsPaginator is a paginator for ListDeviceEvents
type ListDeviceEventsPaginator struct {
	options   ListDeviceEventsPaginatorOptions
	client    ListDeviceEventsAPIClient
	params    *ListDeviceEventsInput
	nextToken *string
	firstPage bool
}

// NewListDeviceEventsPaginator returns a new ListDeviceEventsPaginator
func NewListDeviceEventsPaginator(client ListDeviceEventsAPIClient, params *ListDeviceEventsInput, optFns ...func(*ListDeviceEventsPaginatorOptions)) *ListDeviceEventsPaginator {
	if params == nil {
		params = &ListDeviceEventsInput{}
	}

	options := ListDeviceEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeviceEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeviceEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeviceEvents page.
func (p *ListDeviceEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeviceEventsOutput, error) {
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

	result, err := p.client.ListDeviceEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeviceEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListDeviceEvents",
	}
}
