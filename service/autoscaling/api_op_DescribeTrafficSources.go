// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the traffic sources for the specified Auto Scaling
// group. You can optionally provide a traffic source type. If you provide a
// traffic source type, then the results only include that traffic source type. If
// you do not provide a traffic source type, then the results include all the
// traffic sources for the specified Auto Scaling group.
func (c *Client) DescribeTrafficSources(ctx context.Context, params *DescribeTrafficSourcesInput, optFns ...func(*Options)) (*DescribeTrafficSourcesOutput, error) {
	if params == nil {
		params = &DescribeTrafficSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrafficSources", params, optFns, c.addOperationDescribeTrafficSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrafficSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrafficSourcesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The maximum number of items to return with this call. The maximum value is 50 .
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The traffic source type that you want to describe. The following lists the
	// valid values:
	//   - elb if the traffic source is a Classic Load Balancer.
	//   - elbv2 if the traffic source is a Application Load Balancer, Gateway Load
	//   Balancer, or Network Load Balancer.
	//   - vpc-lattice if the traffic source is VPC Lattice.
	TrafficSourceType *string

	noSmithyDocumentSerde
}

type DescribeTrafficSourcesOutput struct {

	// This string indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Information about the traffic sources.
	TrafficSources []types.TrafficSourceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrafficSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTrafficSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTrafficSources{}, middleware.After)
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
	if err = addOpDescribeTrafficSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrafficSources(options.Region), middleware.Before); err != nil {
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

// DescribeTrafficSourcesAPIClient is a client that implements the
// DescribeTrafficSources operation.
type DescribeTrafficSourcesAPIClient interface {
	DescribeTrafficSources(context.Context, *DescribeTrafficSourcesInput, ...func(*Options)) (*DescribeTrafficSourcesOutput, error)
}

var _ DescribeTrafficSourcesAPIClient = (*Client)(nil)

// DescribeTrafficSourcesPaginatorOptions is the paginator options for
// DescribeTrafficSources
type DescribeTrafficSourcesPaginatorOptions struct {
	// The maximum number of items to return with this call. The maximum value is 50 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTrafficSourcesPaginator is a paginator for DescribeTrafficSources
type DescribeTrafficSourcesPaginator struct {
	options   DescribeTrafficSourcesPaginatorOptions
	client    DescribeTrafficSourcesAPIClient
	params    *DescribeTrafficSourcesInput
	nextToken *string
	firstPage bool
}

// NewDescribeTrafficSourcesPaginator returns a new DescribeTrafficSourcesPaginator
func NewDescribeTrafficSourcesPaginator(client DescribeTrafficSourcesAPIClient, params *DescribeTrafficSourcesInput, optFns ...func(*DescribeTrafficSourcesPaginatorOptions)) *DescribeTrafficSourcesPaginator {
	if params == nil {
		params = &DescribeTrafficSourcesInput{}
	}

	options := DescribeTrafficSourcesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTrafficSourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTrafficSourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTrafficSources page.
func (p *DescribeTrafficSourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTrafficSourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeTrafficSources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTrafficSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeTrafficSources",
	}
}
