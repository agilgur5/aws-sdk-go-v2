// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of shifts generated by an existing rotation in the system.
func (c *Client) ListRotationShifts(ctx context.Context, params *ListRotationShiftsInput, optFns ...func(*Options)) (*ListRotationShiftsOutput, error) {
	if params == nil {
		params = &ListRotationShiftsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRotationShifts", params, optFns, c.addOperationListRotationShiftsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRotationShiftsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRotationShiftsInput struct {

	// The date and time for the end of the time range to list shifts for.
	//
	// This member is required.
	EndTime *time.Time

	// The Amazon Resource Name (ARN) of the rotation to retrieve shift information
	// about.
	//
	// This member is required.
	RotationId *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// The date and time for the beginning of the time range to list shifts for.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type ListRotationShiftsOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Information about shifts that meet the filter criteria.
	RotationShifts []types.RotationShift

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRotationShiftsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRotationShifts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRotationShifts{}, middleware.After)
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
	if err = addOpListRotationShiftsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRotationShifts(options.Region), middleware.Before); err != nil {
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

// ListRotationShiftsAPIClient is a client that implements the ListRotationShifts
// operation.
type ListRotationShiftsAPIClient interface {
	ListRotationShifts(context.Context, *ListRotationShiftsInput, ...func(*Options)) (*ListRotationShiftsOutput, error)
}

var _ ListRotationShiftsAPIClient = (*Client)(nil)

// ListRotationShiftsPaginatorOptions is the paginator options for
// ListRotationShifts
type ListRotationShiftsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRotationShiftsPaginator is a paginator for ListRotationShifts
type ListRotationShiftsPaginator struct {
	options   ListRotationShiftsPaginatorOptions
	client    ListRotationShiftsAPIClient
	params    *ListRotationShiftsInput
	nextToken *string
	firstPage bool
}

// NewListRotationShiftsPaginator returns a new ListRotationShiftsPaginator
func NewListRotationShiftsPaginator(client ListRotationShiftsAPIClient, params *ListRotationShiftsInput, optFns ...func(*ListRotationShiftsPaginatorOptions)) *ListRotationShiftsPaginator {
	if params == nil {
		params = &ListRotationShiftsInput{}
	}

	options := ListRotationShiftsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRotationShiftsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRotationShiftsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRotationShifts page.
func (p *ListRotationShiftsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRotationShiftsOutput, error) {
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

	result, err := p.client.ListRotationShifts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRotationShifts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-contacts",
		OperationName: "ListRotationShifts",
	}
}
