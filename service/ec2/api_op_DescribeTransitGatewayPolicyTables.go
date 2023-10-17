// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more transit gateway route policy tables.
func (c *Client) DescribeTransitGatewayPolicyTables(ctx context.Context, params *DescribeTransitGatewayPolicyTablesInput, optFns ...func(*Options)) (*DescribeTransitGatewayPolicyTablesOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayPolicyTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayPolicyTables", params, optFns, c.addOperationDescribeTransitGatewayPolicyTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayPolicyTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayPolicyTablesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters associated with the transit gateway policy table.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the transit gateway policy tables.
	TransitGatewayPolicyTableIds []string

	noSmithyDocumentSerde
}

type DescribeTransitGatewayPolicyTablesOutput struct {

	// The token for the next page of results.
	NextToken *string

	// Describes the transit gateway policy tables.
	TransitGatewayPolicyTables []types.TransitGatewayPolicyTable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTransitGatewayPolicyTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayPolicyTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayPolicyTables{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayPolicyTables(options.Region), middleware.Before); err != nil {
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

// DescribeTransitGatewayPolicyTablesAPIClient is a client that implements the
// DescribeTransitGatewayPolicyTables operation.
type DescribeTransitGatewayPolicyTablesAPIClient interface {
	DescribeTransitGatewayPolicyTables(context.Context, *DescribeTransitGatewayPolicyTablesInput, ...func(*Options)) (*DescribeTransitGatewayPolicyTablesOutput, error)
}

var _ DescribeTransitGatewayPolicyTablesAPIClient = (*Client)(nil)

// DescribeTransitGatewayPolicyTablesPaginatorOptions is the paginator options for
// DescribeTransitGatewayPolicyTables
type DescribeTransitGatewayPolicyTablesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewayPolicyTablesPaginator is a paginator for
// DescribeTransitGatewayPolicyTables
type DescribeTransitGatewayPolicyTablesPaginator struct {
	options   DescribeTransitGatewayPolicyTablesPaginatorOptions
	client    DescribeTransitGatewayPolicyTablesAPIClient
	params    *DescribeTransitGatewayPolicyTablesInput
	nextToken *string
	firstPage bool
}

// NewDescribeTransitGatewayPolicyTablesPaginator returns a new
// DescribeTransitGatewayPolicyTablesPaginator
func NewDescribeTransitGatewayPolicyTablesPaginator(client DescribeTransitGatewayPolicyTablesAPIClient, params *DescribeTransitGatewayPolicyTablesInput, optFns ...func(*DescribeTransitGatewayPolicyTablesPaginatorOptions)) *DescribeTransitGatewayPolicyTablesPaginator {
	if params == nil {
		params = &DescribeTransitGatewayPolicyTablesInput{}
	}

	options := DescribeTransitGatewayPolicyTablesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTransitGatewayPolicyTablesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewayPolicyTablesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTransitGatewayPolicyTables page.
func (p *DescribeTransitGatewayPolicyTablesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewayPolicyTablesOutput, error) {
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

	result, err := p.client.DescribeTransitGatewayPolicyTables(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTransitGatewayPolicyTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayPolicyTables",
	}
}
