// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A list of locations from which the customer can choose to pickup a device.
func (c *Client) ListPickupLocations(ctx context.Context, params *ListPickupLocationsInput, optFns ...func(*Options)) (*ListPickupLocationsOutput, error) {
	if params == nil {
		params = &ListPickupLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPickupLocations", params, optFns, c.addOperationListPickupLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPickupLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPickupLocationsInput struct {

	// The maximum number of locations to list per page.
	MaxResults *int32

	// HTTP requests are stateless. To identify what object comes "next" in the list
	// of ListPickupLocationsRequest objects, you have the option of specifying
	// NextToken as the starting point for your returned list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPickupLocationsOutput struct {

	// Information about the address of pickup locations.
	Addresses []types.Address

	// HTTP requests are stateless. To identify what object comes "next" in the list
	// of ListPickupLocationsResult objects, you have the option of specifying
	// NextToken as the starting point for your returned list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPickupLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPickupLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPickupLocations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPickupLocations(options.Region), middleware.Before); err != nil {
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

// ListPickupLocationsAPIClient is a client that implements the
// ListPickupLocations operation.
type ListPickupLocationsAPIClient interface {
	ListPickupLocations(context.Context, *ListPickupLocationsInput, ...func(*Options)) (*ListPickupLocationsOutput, error)
}

var _ ListPickupLocationsAPIClient = (*Client)(nil)

// ListPickupLocationsPaginatorOptions is the paginator options for
// ListPickupLocations
type ListPickupLocationsPaginatorOptions struct {
	// The maximum number of locations to list per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPickupLocationsPaginator is a paginator for ListPickupLocations
type ListPickupLocationsPaginator struct {
	options   ListPickupLocationsPaginatorOptions
	client    ListPickupLocationsAPIClient
	params    *ListPickupLocationsInput
	nextToken *string
	firstPage bool
}

// NewListPickupLocationsPaginator returns a new ListPickupLocationsPaginator
func NewListPickupLocationsPaginator(client ListPickupLocationsAPIClient, params *ListPickupLocationsInput, optFns ...func(*ListPickupLocationsPaginatorOptions)) *ListPickupLocationsPaginator {
	if params == nil {
		params = &ListPickupLocationsInput{}
	}

	options := ListPickupLocationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPickupLocationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPickupLocationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPickupLocations page.
func (p *ListPickupLocationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPickupLocationsOutput, error) {
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

	result, err := p.client.ListPickupLocations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPickupLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "ListPickupLocations",
	}
}
