// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the data set imports for the specified application.
func (c *Client) ListDataSetImportHistory(ctx context.Context, params *ListDataSetImportHistoryInput, optFns ...func(*Options)) (*ListDataSetImportHistoryOutput, error) {
	if params == nil {
		params = &ListDataSetImportHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSetImportHistory", params, optFns, c.addOperationListDataSetImportHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSetImportHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSetImportHistoryInput struct {

	// The unique identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The maximum number of objects to return.
	MaxResults *int32

	// A pagination token returned from a previous call to this operation. This
	// specifies the next item to return. To return to the beginning of the list,
	// exclude this parameter.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataSetImportHistoryOutput struct {

	// The data set import tasks.
	//
	// This member is required.
	DataSetImportTasks []types.DataSetImportTask

	// If there are more items to return, this contains a token that is passed to a
	// subsequent call to this operation to retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSetImportHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSetImportHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSetImportHistory{}, middleware.After)
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
	if err = addOpListDataSetImportHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSetImportHistory(options.Region), middleware.Before); err != nil {
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

// ListDataSetImportHistoryAPIClient is a client that implements the
// ListDataSetImportHistory operation.
type ListDataSetImportHistoryAPIClient interface {
	ListDataSetImportHistory(context.Context, *ListDataSetImportHistoryInput, ...func(*Options)) (*ListDataSetImportHistoryOutput, error)
}

var _ ListDataSetImportHistoryAPIClient = (*Client)(nil)

// ListDataSetImportHistoryPaginatorOptions is the paginator options for
// ListDataSetImportHistory
type ListDataSetImportHistoryPaginatorOptions struct {
	// The maximum number of objects to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSetImportHistoryPaginator is a paginator for ListDataSetImportHistory
type ListDataSetImportHistoryPaginator struct {
	options   ListDataSetImportHistoryPaginatorOptions
	client    ListDataSetImportHistoryAPIClient
	params    *ListDataSetImportHistoryInput
	nextToken *string
	firstPage bool
}

// NewListDataSetImportHistoryPaginator returns a new
// ListDataSetImportHistoryPaginator
func NewListDataSetImportHistoryPaginator(client ListDataSetImportHistoryAPIClient, params *ListDataSetImportHistoryInput, optFns ...func(*ListDataSetImportHistoryPaginatorOptions)) *ListDataSetImportHistoryPaginator {
	if params == nil {
		params = &ListDataSetImportHistoryInput{}
	}

	options := ListDataSetImportHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSetImportHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSetImportHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataSetImportHistory page.
func (p *ListDataSetImportHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSetImportHistoryOutput, error) {
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

	result, err := p.client.ListDataSetImportHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataSetImportHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "ListDataSetImportHistory",
	}
}
