// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists route calculator resources in your Amazon Web Services account.
func (c *Client) ListRouteCalculators(ctx context.Context, params *ListRouteCalculatorsInput, optFns ...func(*Options)) (*ListRouteCalculatorsOutput, error) {
	if params == nil {
		params = &ListRouteCalculatorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRouteCalculators", params, optFns, c.addOperationListRouteCalculatorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRouteCalculatorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRouteCalculatorsInput struct {

	// An optional maximum number of results returned in a single call. Default Value:
	// 100
	MaxResults *int32

	// The pagination token specifying which page of results to return in the
	// response. If no token is provided, the default page is the first page. Default
	// Value: null
	NextToken *string

	noSmithyDocumentSerde
}

type ListRouteCalculatorsOutput struct {

	// Lists the route calculator resources that exist in your Amazon Web Services
	// account
	//
	// This member is required.
	Entries []types.ListRouteCalculatorsResponseEntry

	// A pagination token indicating there are additional pages available. You can use
	// the token in a subsequent request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRouteCalculatorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRouteCalculators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRouteCalculators{}, middleware.After)
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
	if err = addEndpointPrefix_opListRouteCalculatorsMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRouteCalculators(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListRouteCalculatorsMiddleware struct {
}

func (*endpointPrefix_opListRouteCalculatorsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListRouteCalculatorsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.routes." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListRouteCalculatorsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListRouteCalculatorsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListRouteCalculatorsAPIClient is a client that implements the
// ListRouteCalculators operation.
type ListRouteCalculatorsAPIClient interface {
	ListRouteCalculators(context.Context, *ListRouteCalculatorsInput, ...func(*Options)) (*ListRouteCalculatorsOutput, error)
}

var _ ListRouteCalculatorsAPIClient = (*Client)(nil)

// ListRouteCalculatorsPaginatorOptions is the paginator options for
// ListRouteCalculators
type ListRouteCalculatorsPaginatorOptions struct {
	// An optional maximum number of results returned in a single call. Default Value:
	// 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRouteCalculatorsPaginator is a paginator for ListRouteCalculators
type ListRouteCalculatorsPaginator struct {
	options   ListRouteCalculatorsPaginatorOptions
	client    ListRouteCalculatorsAPIClient
	params    *ListRouteCalculatorsInput
	nextToken *string
	firstPage bool
}

// NewListRouteCalculatorsPaginator returns a new ListRouteCalculatorsPaginator
func NewListRouteCalculatorsPaginator(client ListRouteCalculatorsAPIClient, params *ListRouteCalculatorsInput, optFns ...func(*ListRouteCalculatorsPaginatorOptions)) *ListRouteCalculatorsPaginator {
	if params == nil {
		params = &ListRouteCalculatorsInput{}
	}

	options := ListRouteCalculatorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRouteCalculatorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRouteCalculatorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRouteCalculators page.
func (p *ListRouteCalculatorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRouteCalculatorsOutput, error) {
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

	result, err := p.client.ListRouteCalculators(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRouteCalculators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "ListRouteCalculators",
	}
}
