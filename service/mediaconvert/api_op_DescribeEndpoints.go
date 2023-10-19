// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Send an request with an empty body to the regional API endpoint to get your
// account API endpoint.
func (c *Client) DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) {
	if params == nil {
		params = &DescribeEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpoints", params, optFns, c.addOperationDescribeEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeEndpointsRequest
type DescribeEndpointsInput struct {

	// Optional. Max number of endpoints, up to twenty, that will be returned at one
	// time.
	MaxResults *int32

	// Optional field, defaults to DEFAULT. Specify DEFAULT for this operation to
	// return your endpoints if any exist, or to create an endpoint for you and return
	// it if one doesn't already exist. Specify GET_ONLY to return your endpoints if
	// any exist, or an empty list if none exist.
	Mode types.DescribeEndpointsMode

	// Use this string, provided with the response to a previous request, to request
	// the next batch of endpoints.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeEndpointsOutput struct {

	// List of endpoints
	Endpoints []types.Endpoint

	// Use this string to request the next batch of endpoints.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeEndpoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpoints(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeEndpoints",
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

// DescribeEndpointsAPIClient is a client that implements the DescribeEndpoints
// operation.
type DescribeEndpointsAPIClient interface {
	DescribeEndpoints(context.Context, *DescribeEndpointsInput, ...func(*Options)) (*DescribeEndpointsOutput, error)
}

var _ DescribeEndpointsAPIClient = (*Client)(nil)

// DescribeEndpointsPaginatorOptions is the paginator options for DescribeEndpoints
type DescribeEndpointsPaginatorOptions struct {
	// Optional. Max number of endpoints, up to twenty, that will be returned at one
	// time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEndpointsPaginator is a paginator for DescribeEndpoints
type DescribeEndpointsPaginator struct {
	options   DescribeEndpointsPaginatorOptions
	client    DescribeEndpointsAPIClient
	params    *DescribeEndpointsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEndpointsPaginator returns a new DescribeEndpointsPaginator
func NewDescribeEndpointsPaginator(client DescribeEndpointsAPIClient, params *DescribeEndpointsInput, optFns ...func(*DescribeEndpointsPaginatorOptions)) *DescribeEndpointsPaginator {
	if params == nil {
		params = &DescribeEndpointsInput{}
	}

	options := DescribeEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEndpoints page.
func (p *DescribeEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) {
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

	result, err := p.client.DescribeEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "DescribeEndpoints",
	}
}
