// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a list of the health checks that are associated with the current
// Amazon Web Services account.
func (c *Client) ListHealthChecks(ctx context.Context, params *ListHealthChecksInput, optFns ...func(*Options)) (*ListHealthChecksOutput, error) {
	if params == nil {
		params = &ListHealthChecksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHealthChecks", params, optFns, c.addOperationListHealthChecksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHealthChecksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve a list of the health checks that are associated with the
// current Amazon Web Services account.
type ListHealthChecksInput struct {

	// If the value of IsTruncated in the previous response was true , you have more
	// health checks. To get another group, submit another ListHealthChecks request.
	// For the value of marker , specify the value of NextMarker from the previous
	// response, which is the ID of the first health check that Amazon Route 53 will
	// return if you submit another request. If the value of IsTruncated in the
	// previous response was false , there are no more health checks to get.
	Marker *string

	// The maximum number of health checks that you want ListHealthChecks to return in
	// response to the current request. Amazon Route 53 returns a maximum of 1000
	// items. If you set MaxItems to a value greater than 1000, Route 53 returns only
	// the first 1000 health checks.
	MaxItems *int32

	noSmithyDocumentSerde
}

// A complex type that contains the response to a ListHealthChecks request.
type ListHealthChecksOutput struct {

	// A complex type that contains one HealthCheck element for each health check that
	// is associated with the current Amazon Web Services account.
	//
	// This member is required.
	HealthChecks []types.HealthCheck

	// A flag that indicates whether there are more health checks to be listed. If the
	// response was truncated, you can get the next group of health checks by
	// submitting another ListHealthChecks request and specifying the value of
	// NextMarker in the marker parameter.
	//
	// This member is required.
	IsTruncated bool

	// For the second and subsequent calls to ListHealthChecks , Marker is the value
	// that you specified for the marker parameter in the previous request.
	//
	// This member is required.
	Marker *string

	// The value that you specified for the maxitems parameter in the call to
	// ListHealthChecks that produced the current response.
	//
	// This member is required.
	MaxItems *int32

	// If IsTruncated is true , the value of NextMarker identifies the first health
	// check that Amazon Route 53 returns if you submit another ListHealthChecks
	// request and specify the value of NextMarker in the marker parameter.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHealthChecksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListHealthChecks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListHealthChecks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHealthChecks(options.Region), middleware.Before); err != nil {
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

// ListHealthChecksAPIClient is a client that implements the ListHealthChecks
// operation.
type ListHealthChecksAPIClient interface {
	ListHealthChecks(context.Context, *ListHealthChecksInput, ...func(*Options)) (*ListHealthChecksOutput, error)
}

var _ ListHealthChecksAPIClient = (*Client)(nil)

// ListHealthChecksPaginatorOptions is the paginator options for ListHealthChecks
type ListHealthChecksPaginatorOptions struct {
	// The maximum number of health checks that you want ListHealthChecks to return in
	// response to the current request. Amazon Route 53 returns a maximum of 1000
	// items. If you set MaxItems to a value greater than 1000, Route 53 returns only
	// the first 1000 health checks.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHealthChecksPaginator is a paginator for ListHealthChecks
type ListHealthChecksPaginator struct {
	options   ListHealthChecksPaginatorOptions
	client    ListHealthChecksAPIClient
	params    *ListHealthChecksInput
	nextToken *string
	firstPage bool
}

// NewListHealthChecksPaginator returns a new ListHealthChecksPaginator
func NewListHealthChecksPaginator(client ListHealthChecksAPIClient, params *ListHealthChecksInput, optFns ...func(*ListHealthChecksPaginatorOptions)) *ListHealthChecksPaginator {
	if params == nil {
		params = &ListHealthChecksInput{}
	}

	options := ListHealthChecksPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHealthChecksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHealthChecksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHealthChecks page.
func (p *ListHealthChecksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHealthChecksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListHealthChecks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListHealthChecks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListHealthChecks",
	}
}
