// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists objects attached to the specified index.
func (c *Client) ListIndex(ctx context.Context, params *ListIndexInput, optFns ...func(*Options)) (*ListIndexOutput, error) {
	if params == nil {
		params = &ListIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIndex", params, optFns, c.addOperationListIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIndexInput struct {

	// The ARN of the directory that the index exists in.
	//
	// This member is required.
	DirectoryArn *string

	// The reference to the index to list.
	//
	// This member is required.
	IndexReference *types.ObjectReference

	// The consistency level to execute the request at.
	ConsistencyLevel types.ConsistencyLevel

	// The maximum number of objects in a single page to retrieve from the index
	// during a request. For more information, see Amazon Cloud Directory Limits (http://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html)
	// .
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// Specifies the ranges of indexed values that you want to query.
	RangesOnIndexedValues []types.ObjectAttributeRange

	noSmithyDocumentSerde
}

type ListIndexOutput struct {

	// The objects and indexed values attached to the index.
	IndexAttachments []types.IndexAttachment

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIndex{}, middleware.After)
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
	if err = addOpListIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIndex(options.Region), middleware.Before); err != nil {
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

// ListIndexAPIClient is a client that implements the ListIndex operation.
type ListIndexAPIClient interface {
	ListIndex(context.Context, *ListIndexInput, ...func(*Options)) (*ListIndexOutput, error)
}

var _ ListIndexAPIClient = (*Client)(nil)

// ListIndexPaginatorOptions is the paginator options for ListIndex
type ListIndexPaginatorOptions struct {
	// The maximum number of objects in a single page to retrieve from the index
	// during a request. For more information, see Amazon Cloud Directory Limits (http://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html)
	// .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIndexPaginator is a paginator for ListIndex
type ListIndexPaginator struct {
	options   ListIndexPaginatorOptions
	client    ListIndexAPIClient
	params    *ListIndexInput
	nextToken *string
	firstPage bool
}

// NewListIndexPaginator returns a new ListIndexPaginator
func NewListIndexPaginator(client ListIndexAPIClient, params *ListIndexInput, optFns ...func(*ListIndexPaginatorOptions)) *ListIndexPaginator {
	if params == nil {
		params = &ListIndexInput{}
	}

	options := ListIndexPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIndexPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIndexPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIndex page.
func (p *ListIndexPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIndexOutput, error) {
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

	result, err := p.client.ListIndex(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListIndex",
	}
}
