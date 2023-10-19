// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about reserved nodes for this account, or about a specified
// reserved node.
func (c *Client) DescribeReservedNodes(ctx context.Context, params *DescribeReservedNodesInput, optFns ...func(*Options)) (*DescribeReservedNodesOutput, error) {
	if params == nil {
		params = &DescribeReservedNodesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedNodes", params, optFns, c.addOperationDescribeReservedNodesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReservedNodesInput struct {

	// The duration filter value, specified in years or seconds. Use this parameter to
	// show only reservations for this duration.
	Duration *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	MaxResults *int32

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	NextToken *string

	// The node type filter value. Use this parameter to show only those reservations
	// matching the specified node type. For more information, see Supported node types (https://docs.aws.amazon.com/memorydb/latest/devguide/nodes.reserved.html#reserved-nodes-supported)
	// .
	NodeType *string

	// The offering type filter value. Use this parameter to show only the available
	// offerings matching the specified offering type. Valid values: "All
	// Upfront"|"Partial Upfront"| "No Upfront"
	OfferingType *string

	// The reserved node identifier filter value. Use this parameter to show only the
	// reservation that matches the specified reservation ID.
	ReservationId *string

	// The offering identifier filter value. Use this parameter to show only purchased
	// reservations matching the specified offering identifier.
	ReservedNodesOfferingId *string

	noSmithyDocumentSerde
}

type DescribeReservedNodesOutput struct {

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	NextToken *string

	// Returns information about reserved nodes for this account, or about a specified
	// reserved node.
	ReservedNodes []types.ReservedNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedNodesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReservedNodes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReservedNodes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedNodes(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeReservedNodes",
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

// DescribeReservedNodesAPIClient is a client that implements the
// DescribeReservedNodes operation.
type DescribeReservedNodesAPIClient interface {
	DescribeReservedNodes(context.Context, *DescribeReservedNodesInput, ...func(*Options)) (*DescribeReservedNodesOutput, error)
}

var _ DescribeReservedNodesAPIClient = (*Client)(nil)

// DescribeReservedNodesPaginatorOptions is the paginator options for
// DescribeReservedNodes
type DescribeReservedNodesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedNodesPaginator is a paginator for DescribeReservedNodes
type DescribeReservedNodesPaginator struct {
	options   DescribeReservedNodesPaginatorOptions
	client    DescribeReservedNodesAPIClient
	params    *DescribeReservedNodesInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedNodesPaginator returns a new DescribeReservedNodesPaginator
func NewDescribeReservedNodesPaginator(client DescribeReservedNodesAPIClient, params *DescribeReservedNodesInput, optFns ...func(*DescribeReservedNodesPaginatorOptions)) *DescribeReservedNodesPaginator {
	if params == nil {
		params = &DescribeReservedNodesInput{}
	}

	options := DescribeReservedNodesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedNodesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedNodesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedNodes page.
func (p *DescribeReservedNodesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedNodesOutput, error) {
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

	result, err := p.client.DescribeReservedNodes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReservedNodes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "DescribeReservedNodes",
	}
}
