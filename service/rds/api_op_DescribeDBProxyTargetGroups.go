// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about DB proxy target groups, represented by
// DBProxyTargetGroup data structures.
func (c *Client) DescribeDBProxyTargetGroups(ctx context.Context, params *DescribeDBProxyTargetGroupsInput, optFns ...func(*Options)) (*DescribeDBProxyTargetGroupsOutput, error) {
	if params == nil {
		params = &DescribeDBProxyTargetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBProxyTargetGroups", params, optFns, c.addOperationDescribeDBProxyTargetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBProxyTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBProxyTargetGroupsInput struct {

	// The identifier of the DBProxy associated with the target group.
	//
	// This member is required.
	DBProxyName *string

	// This parameter is not currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The identifier of the DBProxyTargetGroup to describe.
	TargetGroupName *string

	noSmithyDocumentSerde
}

type DescribeDBProxyTargetGroupsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// An arbitrary number of DBProxyTargetGroup objects, containing details of the
	// corresponding target groups.
	TargetGroups []types.DBProxyTargetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBProxyTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBProxyTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBProxyTargetGroups{}, middleware.After)
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
	if err = addOpDescribeDBProxyTargetGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBProxyTargetGroups(options.Region), middleware.Before); err != nil {
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

// DescribeDBProxyTargetGroupsAPIClient is a client that implements the
// DescribeDBProxyTargetGroups operation.
type DescribeDBProxyTargetGroupsAPIClient interface {
	DescribeDBProxyTargetGroups(context.Context, *DescribeDBProxyTargetGroupsInput, ...func(*Options)) (*DescribeDBProxyTargetGroupsOutput, error)
}

var _ DescribeDBProxyTargetGroupsAPIClient = (*Client)(nil)

// DescribeDBProxyTargetGroupsPaginatorOptions is the paginator options for
// DescribeDBProxyTargetGroups
type DescribeDBProxyTargetGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBProxyTargetGroupsPaginator is a paginator for
// DescribeDBProxyTargetGroups
type DescribeDBProxyTargetGroupsPaginator struct {
	options   DescribeDBProxyTargetGroupsPaginatorOptions
	client    DescribeDBProxyTargetGroupsAPIClient
	params    *DescribeDBProxyTargetGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBProxyTargetGroupsPaginator returns a new
// DescribeDBProxyTargetGroupsPaginator
func NewDescribeDBProxyTargetGroupsPaginator(client DescribeDBProxyTargetGroupsAPIClient, params *DescribeDBProxyTargetGroupsInput, optFns ...func(*DescribeDBProxyTargetGroupsPaginatorOptions)) *DescribeDBProxyTargetGroupsPaginator {
	if params == nil {
		params = &DescribeDBProxyTargetGroupsInput{}
	}

	options := DescribeDBProxyTargetGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBProxyTargetGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBProxyTargetGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBProxyTargetGroups page.
func (p *DescribeDBProxyTargetGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBProxyTargetGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBProxyTargetGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBProxyTargetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBProxyTargetGroups",
	}
}
