// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified the load balancers. If no load balancers are specified,
// the call describes all of your load balancers.
func (c *Client) DescribeLoadBalancers(ctx context.Context, params *DescribeLoadBalancersInput, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) {
	if params == nil {
		params = &DescribeLoadBalancersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoadBalancers", params, optFns, c.addOperationDescribeLoadBalancersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeLoadBalancers.
type DescribeLoadBalancersInput struct {

	// The names of the load balancers.
	LoadBalancerNames []string

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The maximum number of results to return with this call (a number from 1 to
	// 400). The default is 400.
	PageSize *int32

	noSmithyDocumentSerde
}

// Contains the parameters for DescribeLoadBalancers.
type DescribeLoadBalancersOutput struct {

	// Information about the load balancers.
	LoadBalancerDescriptions []types.LoadBalancerDescription

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoadBalancersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancers(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeLoadBalancers",
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

// DescribeLoadBalancersAPIClient is a client that implements the
// DescribeLoadBalancers operation.
type DescribeLoadBalancersAPIClient interface {
	DescribeLoadBalancers(context.Context, *DescribeLoadBalancersInput, ...func(*Options)) (*DescribeLoadBalancersOutput, error)
}

var _ DescribeLoadBalancersAPIClient = (*Client)(nil)

// DescribeLoadBalancersPaginatorOptions is the paginator options for
// DescribeLoadBalancers
type DescribeLoadBalancersPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLoadBalancersPaginator is a paginator for DescribeLoadBalancers
type DescribeLoadBalancersPaginator struct {
	options   DescribeLoadBalancersPaginatorOptions
	client    DescribeLoadBalancersAPIClient
	params    *DescribeLoadBalancersInput
	nextToken *string
	firstPage bool
}

// NewDescribeLoadBalancersPaginator returns a new DescribeLoadBalancersPaginator
func NewDescribeLoadBalancersPaginator(client DescribeLoadBalancersAPIClient, params *DescribeLoadBalancersInput, optFns ...func(*DescribeLoadBalancersPaginatorOptions)) *DescribeLoadBalancersPaginator {
	if params == nil {
		params = &DescribeLoadBalancersInput{}
	}

	options := DescribeLoadBalancersPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLoadBalancersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLoadBalancersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLoadBalancers page.
func (p *DescribeLoadBalancersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.DescribeLoadBalancers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeLoadBalancers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeLoadBalancers",
	}
}
