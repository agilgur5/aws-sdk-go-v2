// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticinference

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticinference/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes information over a provided set of accelerators belonging to an
// account. February 15, 2023: Starting April 15, 2023, AWS will not onboard new
// customers to Amazon Elastic Inference (EI), and will help current customers
// migrate their workloads to options that offer better price and performance.
// After April 15, 2023, new customers will not be able to launch instances with
// Amazon EI accelerators in Amazon SageMaker, Amazon ECS, or Amazon EC2. However,
// customers who have used Amazon EI at least once during the past 30-day period
// are considered current customers and will be able to continue using the service.
func (c *Client) DescribeAccelerators(ctx context.Context, params *DescribeAcceleratorsInput, optFns ...func(*Options)) (*DescribeAcceleratorsOutput, error) {
	if params == nil {
		params = &DescribeAcceleratorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccelerators", params, optFns, c.addOperationDescribeAcceleratorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAcceleratorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAcceleratorsInput struct {

	// The IDs of the accelerators to describe.
	AcceleratorIds []string

	// One or more filters. Filter names and values are case-sensitive. Valid filter
	// names are: accelerator-types: can provide a list of accelerator type names to
	// filter for. instance-id: can provide a list of EC2 instance ids to filter for.
	Filters []types.Filter

	// The total number of items to return in the command's output. If the total
	// number of items available is more than the value specified, a NextToken is
	// provided in the command's output. To resume pagination, provide the NextToken
	// value in the starting-token argument of a subsequent command. Do not use the
	// NextToken response element directly outside of the AWS CLI.
	MaxResults int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAcceleratorsOutput struct {

	// The details of the Elastic Inference Accelerators.
	AcceleratorSet []types.ElasticInferenceAccelerator

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAcceleratorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccelerators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccelerators{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccelerators(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeAccelerators",
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

// DescribeAcceleratorsAPIClient is a client that implements the
// DescribeAccelerators operation.
type DescribeAcceleratorsAPIClient interface {
	DescribeAccelerators(context.Context, *DescribeAcceleratorsInput, ...func(*Options)) (*DescribeAcceleratorsOutput, error)
}

var _ DescribeAcceleratorsAPIClient = (*Client)(nil)

// DescribeAcceleratorsPaginatorOptions is the paginator options for
// DescribeAccelerators
type DescribeAcceleratorsPaginatorOptions struct {
	// The total number of items to return in the command's output. If the total
	// number of items available is more than the value specified, a NextToken is
	// provided in the command's output. To resume pagination, provide the NextToken
	// value in the starting-token argument of a subsequent command. Do not use the
	// NextToken response element directly outside of the AWS CLI.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAcceleratorsPaginator is a paginator for DescribeAccelerators
type DescribeAcceleratorsPaginator struct {
	options   DescribeAcceleratorsPaginatorOptions
	client    DescribeAcceleratorsAPIClient
	params    *DescribeAcceleratorsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAcceleratorsPaginator returns a new DescribeAcceleratorsPaginator
func NewDescribeAcceleratorsPaginator(client DescribeAcceleratorsAPIClient, params *DescribeAcceleratorsInput, optFns ...func(*DescribeAcceleratorsPaginatorOptions)) *DescribeAcceleratorsPaginator {
	if params == nil {
		params = &DescribeAcceleratorsInput{}
	}

	options := DescribeAcceleratorsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAcceleratorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAcceleratorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAccelerators page.
func (p *DescribeAcceleratorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAcceleratorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeAccelerators(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAccelerators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastic-inference",
		OperationName: "DescribeAccelerators",
	}
}
