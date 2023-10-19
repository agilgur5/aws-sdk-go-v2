// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists managed endpoints based on a set of parameters. A managed endpoint is a
// gateway that connects Amazon EMR Studio to Amazon EMR on EKS so that Amazon EMR
// Studio can communicate with your virtual cluster.
func (c *Client) ListManagedEndpoints(ctx context.Context, params *ListManagedEndpointsInput, optFns ...func(*Options)) (*ListManagedEndpointsOutput, error) {
	if params == nil {
		params = &ListManagedEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListManagedEndpoints", params, optFns, c.addOperationListManagedEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListManagedEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListManagedEndpointsInput struct {

	// The ID of the virtual cluster.
	//
	// This member is required.
	VirtualClusterId *string

	// The date and time after which the endpoints are created.
	CreatedAfter *time.Time

	// The date and time before which the endpoints are created.
	CreatedBefore *time.Time

	// The maximum number of managed endpoints that can be listed.
	MaxResults *int32

	// The token for the next set of managed endpoints to return.
	NextToken *string

	// The states of the managed endpoints.
	States []types.EndpointState

	// The types of the managed endpoints.
	Types []string

	noSmithyDocumentSerde
}

type ListManagedEndpointsOutput struct {

	// The managed endpoints to be listed.
	Endpoints []types.Endpoint

	// The token for the next set of endpoints to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListManagedEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListManagedEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListManagedEndpoints{}, middleware.After)
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
	if err = addOpListManagedEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListManagedEndpoints(options.Region), middleware.Before); err != nil {
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
		operation: "ListManagedEndpoints",
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

// ListManagedEndpointsAPIClient is a client that implements the
// ListManagedEndpoints operation.
type ListManagedEndpointsAPIClient interface {
	ListManagedEndpoints(context.Context, *ListManagedEndpointsInput, ...func(*Options)) (*ListManagedEndpointsOutput, error)
}

var _ ListManagedEndpointsAPIClient = (*Client)(nil)

// ListManagedEndpointsPaginatorOptions is the paginator options for
// ListManagedEndpoints
type ListManagedEndpointsPaginatorOptions struct {
	// The maximum number of managed endpoints that can be listed.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListManagedEndpointsPaginator is a paginator for ListManagedEndpoints
type ListManagedEndpointsPaginator struct {
	options   ListManagedEndpointsPaginatorOptions
	client    ListManagedEndpointsAPIClient
	params    *ListManagedEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListManagedEndpointsPaginator returns a new ListManagedEndpointsPaginator
func NewListManagedEndpointsPaginator(client ListManagedEndpointsAPIClient, params *ListManagedEndpointsInput, optFns ...func(*ListManagedEndpointsPaginatorOptions)) *ListManagedEndpointsPaginator {
	if params == nil {
		params = &ListManagedEndpointsInput{}
	}

	options := ListManagedEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListManagedEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListManagedEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListManagedEndpoints page.
func (p *ListManagedEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListManagedEndpointsOutput, error) {
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

	result, err := p.client.ListManagedEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListManagedEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-containers",
		OperationName: "ListManagedEndpoints",
	}
}
