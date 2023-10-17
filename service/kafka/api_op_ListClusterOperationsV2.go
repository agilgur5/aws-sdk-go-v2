// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the operations that have been performed on the specified
// MSK cluster.
func (c *Client) ListClusterOperationsV2(ctx context.Context, params *ListClusterOperationsV2Input, optFns ...func(*Options)) (*ListClusterOperationsV2Output, error) {
	if params == nil {
		params = &ListClusterOperationsV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusterOperationsV2", params, optFns, c.addOperationListClusterOperationsV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClusterOperationsV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClusterOperationsV2Input struct {

	// The arn of the cluster whose operations are being requested.
	//
	// This member is required.
	ClusterArn *string

	// The maxResults of the query.
	MaxResults *int32

	// The nextToken of the query.
	NextToken *string

	noSmithyDocumentSerde
}

type ListClusterOperationsV2Output struct {

	// An array of cluster operation information objects.
	ClusterOperationInfoList []types.ClusterOperationV2Summary

	// If the response of ListClusterOperationsV2 is truncated, it returns a NextToken
	// in the response. This NextToken should be sent in the subsequent request to
	// ListClusterOperationsV2.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClusterOperationsV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListClusterOperationsV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListClusterOperationsV2{}, middleware.After)
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
	if err = addOpListClusterOperationsV2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusterOperationsV2(options.Region), middleware.Before); err != nil {
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

// ListClusterOperationsV2APIClient is a client that implements the
// ListClusterOperationsV2 operation.
type ListClusterOperationsV2APIClient interface {
	ListClusterOperationsV2(context.Context, *ListClusterOperationsV2Input, ...func(*Options)) (*ListClusterOperationsV2Output, error)
}

var _ ListClusterOperationsV2APIClient = (*Client)(nil)

// ListClusterOperationsV2PaginatorOptions is the paginator options for
// ListClusterOperationsV2
type ListClusterOperationsV2PaginatorOptions struct {
	// The maxResults of the query.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClusterOperationsV2Paginator is a paginator for ListClusterOperationsV2
type ListClusterOperationsV2Paginator struct {
	options   ListClusterOperationsV2PaginatorOptions
	client    ListClusterOperationsV2APIClient
	params    *ListClusterOperationsV2Input
	nextToken *string
	firstPage bool
}

// NewListClusterOperationsV2Paginator returns a new
// ListClusterOperationsV2Paginator
func NewListClusterOperationsV2Paginator(client ListClusterOperationsV2APIClient, params *ListClusterOperationsV2Input, optFns ...func(*ListClusterOperationsV2PaginatorOptions)) *ListClusterOperationsV2Paginator {
	if params == nil {
		params = &ListClusterOperationsV2Input{}
	}

	options := ListClusterOperationsV2PaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClusterOperationsV2Paginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClusterOperationsV2Paginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListClusterOperationsV2 page.
func (p *ListClusterOperationsV2Paginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClusterOperationsV2Output, error) {
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

	result, err := p.client.ListClusterOperationsV2(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListClusterOperationsV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "ListClusterOperationsV2",
	}
}
