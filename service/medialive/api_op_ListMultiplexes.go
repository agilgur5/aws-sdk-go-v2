// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a list of the existing multiplexes.
func (c *Client) ListMultiplexes(ctx context.Context, params *ListMultiplexesInput, optFns ...func(*Options)) (*ListMultiplexesOutput, error) {
	if params == nil {
		params = &ListMultiplexesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMultiplexes", params, optFns, c.addOperationListMultiplexesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMultiplexesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListMultiplexesRequest
type ListMultiplexesInput struct {

	// The maximum number of items to return.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListMultiplexesResponse
type ListMultiplexesOutput struct {

	// List of multiplexes.
	Multiplexes []types.MultiplexSummary

	// Token for the next ListMultiplexes request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMultiplexesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMultiplexes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMultiplexes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMultiplexes(options.Region), middleware.Before); err != nil {
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
		operation: "ListMultiplexes",
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

// ListMultiplexesAPIClient is a client that implements the ListMultiplexes
// operation.
type ListMultiplexesAPIClient interface {
	ListMultiplexes(context.Context, *ListMultiplexesInput, ...func(*Options)) (*ListMultiplexesOutput, error)
}

var _ ListMultiplexesAPIClient = (*Client)(nil)

// ListMultiplexesPaginatorOptions is the paginator options for ListMultiplexes
type ListMultiplexesPaginatorOptions struct {
	// The maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMultiplexesPaginator is a paginator for ListMultiplexes
type ListMultiplexesPaginator struct {
	options   ListMultiplexesPaginatorOptions
	client    ListMultiplexesAPIClient
	params    *ListMultiplexesInput
	nextToken *string
	firstPage bool
}

// NewListMultiplexesPaginator returns a new ListMultiplexesPaginator
func NewListMultiplexesPaginator(client ListMultiplexesAPIClient, params *ListMultiplexesInput, optFns ...func(*ListMultiplexesPaginatorOptions)) *ListMultiplexesPaginator {
	if params == nil {
		params = &ListMultiplexesInput{}
	}

	options := ListMultiplexesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMultiplexesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMultiplexesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMultiplexes page.
func (p *ListMultiplexesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMultiplexesOutput, error) {
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

	result, err := p.client.ListMultiplexes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMultiplexes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "ListMultiplexes",
	}
}
