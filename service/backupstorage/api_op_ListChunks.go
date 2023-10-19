// Code generated by smithy-go-codegen DO NOT EDIT.

package backupstorage

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backupstorage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List chunks in a given Object
func (c *Client) ListChunks(ctx context.Context, params *ListChunksInput, optFns ...func(*Options)) (*ListChunksOutput, error) {
	if params == nil {
		params = &ListChunksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChunks", params, optFns, c.addOperationListChunksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChunksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChunksInput struct {

	// Object token
	//
	// This member is required.
	ObjectToken *string

	// Storage job id
	//
	// This member is required.
	StorageJobId *string

	// Maximum number of chunks
	MaxResults int32

	// Pagination token
	NextToken *string

	noSmithyDocumentSerde
}

type ListChunksOutput struct {

	// List of chunks
	//
	// This member is required.
	ChunkList []types.Chunk

	// Pagination token
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChunksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChunks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChunks{}, middleware.After)
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
	if err = addOpListChunksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChunks(options.Region), middleware.Before); err != nil {
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
		operation: "ListChunks",
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

// ListChunksAPIClient is a client that implements the ListChunks operation.
type ListChunksAPIClient interface {
	ListChunks(context.Context, *ListChunksInput, ...func(*Options)) (*ListChunksOutput, error)
}

var _ ListChunksAPIClient = (*Client)(nil)

// ListChunksPaginatorOptions is the paginator options for ListChunks
type ListChunksPaginatorOptions struct {
	// Maximum number of chunks
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChunksPaginator is a paginator for ListChunks
type ListChunksPaginator struct {
	options   ListChunksPaginatorOptions
	client    ListChunksAPIClient
	params    *ListChunksInput
	nextToken *string
	firstPage bool
}

// NewListChunksPaginator returns a new ListChunksPaginator
func NewListChunksPaginator(client ListChunksAPIClient, params *ListChunksInput, optFns ...func(*ListChunksPaginatorOptions)) *ListChunksPaginator {
	if params == nil {
		params = &ListChunksInput{}
	}

	options := ListChunksPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChunksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChunksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChunks page.
func (p *ListChunksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChunksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListChunks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListChunks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-storage",
		OperationName: "ListChunks",
	}
}
