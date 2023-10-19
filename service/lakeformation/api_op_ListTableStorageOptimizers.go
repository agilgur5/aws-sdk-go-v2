// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the configuration of all storage optimizers associated with a specified
// table.
func (c *Client) ListTableStorageOptimizers(ctx context.Context, params *ListTableStorageOptimizersInput, optFns ...func(*Options)) (*ListTableStorageOptimizersOutput, error) {
	if params == nil {
		params = &ListTableStorageOptimizersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTableStorageOptimizers", params, optFns, c.addOperationListTableStorageOptimizersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTableStorageOptimizersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTableStorageOptimizersInput struct {

	// Name of the database where the table is present.
	//
	// This member is required.
	DatabaseName *string

	// Name of the table.
	//
	// This member is required.
	TableName *string

	// The Catalog ID of the table.
	CatalogId *string

	// The number of storage optimizers to return on each call.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string

	// The specific type of storage optimizers to list. The supported value is
	// compaction .
	StorageOptimizerType types.OptimizerType

	noSmithyDocumentSerde
}

type ListTableStorageOptimizersOutput struct {

	// A continuation token for paginating the returned list of tokens, returned if
	// the current segment of the list is not the last.
	NextToken *string

	// A list of the storage optimizers associated with a table.
	StorageOptimizerList []types.StorageOptimizer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTableStorageOptimizersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTableStorageOptimizers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTableStorageOptimizers{}, middleware.After)
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
	if err = addOpListTableStorageOptimizersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTableStorageOptimizers(options.Region), middleware.Before); err != nil {
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
		operation: "ListTableStorageOptimizers",
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

// ListTableStorageOptimizersAPIClient is a client that implements the
// ListTableStorageOptimizers operation.
type ListTableStorageOptimizersAPIClient interface {
	ListTableStorageOptimizers(context.Context, *ListTableStorageOptimizersInput, ...func(*Options)) (*ListTableStorageOptimizersOutput, error)
}

var _ ListTableStorageOptimizersAPIClient = (*Client)(nil)

// ListTableStorageOptimizersPaginatorOptions is the paginator options for
// ListTableStorageOptimizers
type ListTableStorageOptimizersPaginatorOptions struct {
	// The number of storage optimizers to return on each call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTableStorageOptimizersPaginator is a paginator for
// ListTableStorageOptimizers
type ListTableStorageOptimizersPaginator struct {
	options   ListTableStorageOptimizersPaginatorOptions
	client    ListTableStorageOptimizersAPIClient
	params    *ListTableStorageOptimizersInput
	nextToken *string
	firstPage bool
}

// NewListTableStorageOptimizersPaginator returns a new
// ListTableStorageOptimizersPaginator
func NewListTableStorageOptimizersPaginator(client ListTableStorageOptimizersAPIClient, params *ListTableStorageOptimizersInput, optFns ...func(*ListTableStorageOptimizersPaginatorOptions)) *ListTableStorageOptimizersPaginator {
	if params == nil {
		params = &ListTableStorageOptimizersInput{}
	}

	options := ListTableStorageOptimizersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTableStorageOptimizersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTableStorageOptimizersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTableStorageOptimizers page.
func (p *ListTableStorageOptimizersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTableStorageOptimizersOutput, error) {
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

	result, err := p.client.ListTableStorageOptimizers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTableStorageOptimizers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "ListTableStorageOptimizers",
	}
}
