// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all journal export jobs for a specified ledger. This action returns a
// maximum of MaxResults items, and is paginated so that you can retrieve all the
// items by calling ListJournalS3ExportsForLedger multiple times. This action does
// not return any expired export jobs. For more information, see Export job
// expiration (https://docs.aws.amazon.com/qldb/latest/developerguide/export-journal.request.html#export-journal.request.expiration)
// in the Amazon QLDB Developer Guide.
func (c *Client) ListJournalS3ExportsForLedger(ctx context.Context, params *ListJournalS3ExportsForLedgerInput, optFns ...func(*Options)) (*ListJournalS3ExportsForLedgerOutput, error) {
	if params == nil {
		params = &ListJournalS3ExportsForLedgerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJournalS3ExportsForLedger", params, optFns, c.addOperationListJournalS3ExportsForLedgerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJournalS3ExportsForLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJournalS3ExportsForLedgerInput struct {

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The maximum number of results to return in a single
	// ListJournalS3ExportsForLedger request. (The actual number of results returned
	// might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListJournalS3ExportsForLedger call, then you should use that value as input here.
	NextToken *string

	noSmithyDocumentSerde
}

type ListJournalS3ExportsForLedgerOutput struct {

	// The journal export jobs that are currently associated with the specified ledger.
	JournalS3Exports []types.JournalS3ExportDescription

	//   - If NextToken is empty, then the last page of results has been processed and
	//   there are no more results to be retrieved.
	//   - If NextToken is not empty, then there are more results available. To
	//   retrieve the next page of results, use the value of NextToken in a subsequent
	//   ListJournalS3ExportsForLedger call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJournalS3ExportsForLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJournalS3ExportsForLedger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJournalS3ExportsForLedger{}, middleware.After)
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
	if err = addOpListJournalS3ExportsForLedgerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJournalS3ExportsForLedger(options.Region), middleware.Before); err != nil {
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
		operation: "ListJournalS3ExportsForLedger",
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

// ListJournalS3ExportsForLedgerAPIClient is a client that implements the
// ListJournalS3ExportsForLedger operation.
type ListJournalS3ExportsForLedgerAPIClient interface {
	ListJournalS3ExportsForLedger(context.Context, *ListJournalS3ExportsForLedgerInput, ...func(*Options)) (*ListJournalS3ExportsForLedgerOutput, error)
}

var _ ListJournalS3ExportsForLedgerAPIClient = (*Client)(nil)

// ListJournalS3ExportsForLedgerPaginatorOptions is the paginator options for
// ListJournalS3ExportsForLedger
type ListJournalS3ExportsForLedgerPaginatorOptions struct {
	// The maximum number of results to return in a single
	// ListJournalS3ExportsForLedger request. (The actual number of results returned
	// might be fewer.)
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJournalS3ExportsForLedgerPaginator is a paginator for
// ListJournalS3ExportsForLedger
type ListJournalS3ExportsForLedgerPaginator struct {
	options   ListJournalS3ExportsForLedgerPaginatorOptions
	client    ListJournalS3ExportsForLedgerAPIClient
	params    *ListJournalS3ExportsForLedgerInput
	nextToken *string
	firstPage bool
}

// NewListJournalS3ExportsForLedgerPaginator returns a new
// ListJournalS3ExportsForLedgerPaginator
func NewListJournalS3ExportsForLedgerPaginator(client ListJournalS3ExportsForLedgerAPIClient, params *ListJournalS3ExportsForLedgerInput, optFns ...func(*ListJournalS3ExportsForLedgerPaginatorOptions)) *ListJournalS3ExportsForLedgerPaginator {
	if params == nil {
		params = &ListJournalS3ExportsForLedgerInput{}
	}

	options := ListJournalS3ExportsForLedgerPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJournalS3ExportsForLedgerPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJournalS3ExportsForLedgerPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJournalS3ExportsForLedger page.
func (p *ListJournalS3ExportsForLedgerPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJournalS3ExportsForLedgerOutput, error) {
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

	result, err := p.client.ListJournalS3ExportsForLedger(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJournalS3ExportsForLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListJournalS3ExportsForLedger",
	}
}
