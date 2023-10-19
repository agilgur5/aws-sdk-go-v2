// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a collection of OriginEndpoint records.
func (c *Client) ListOriginEndpoints(ctx context.Context, params *ListOriginEndpointsInput, optFns ...func(*Options)) (*ListOriginEndpointsOutput, error) {
	if params == nil {
		params = &ListOriginEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOriginEndpoints", params, optFns, c.addOperationListOriginEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOriginEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOriginEndpointsInput struct {

	// When specified, the request will return only OriginEndpoints associated with
	// the given Channel ID.
	ChannelId *string

	// The upper bound on the number of records to return.
	MaxResults *int32

	// A token used to resume pagination from the end of a previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOriginEndpointsOutput struct {

	// A token that can be used to resume pagination from the end of the collection.
	NextToken *string

	// A list of OriginEndpoint records.
	OriginEndpoints []types.OriginEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOriginEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOriginEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOriginEndpoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOriginEndpoints(options.Region), middleware.Before); err != nil {
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
		operation: "ListOriginEndpoints",
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

// ListOriginEndpointsAPIClient is a client that implements the
// ListOriginEndpoints operation.
type ListOriginEndpointsAPIClient interface {
	ListOriginEndpoints(context.Context, *ListOriginEndpointsInput, ...func(*Options)) (*ListOriginEndpointsOutput, error)
}

var _ ListOriginEndpointsAPIClient = (*Client)(nil)

// ListOriginEndpointsPaginatorOptions is the paginator options for
// ListOriginEndpoints
type ListOriginEndpointsPaginatorOptions struct {
	// The upper bound on the number of records to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOriginEndpointsPaginator is a paginator for ListOriginEndpoints
type ListOriginEndpointsPaginator struct {
	options   ListOriginEndpointsPaginatorOptions
	client    ListOriginEndpointsAPIClient
	params    *ListOriginEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListOriginEndpointsPaginator returns a new ListOriginEndpointsPaginator
func NewListOriginEndpointsPaginator(client ListOriginEndpointsAPIClient, params *ListOriginEndpointsInput, optFns ...func(*ListOriginEndpointsPaginatorOptions)) *ListOriginEndpointsPaginator {
	if params == nil {
		params = &ListOriginEndpointsInput{}
	}

	options := ListOriginEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOriginEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOriginEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOriginEndpoints page.
func (p *ListOriginEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOriginEndpointsOutput, error) {
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

	result, err := p.client.ListOriginEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOriginEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "ListOriginEndpoints",
	}
}
