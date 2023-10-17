// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the revisions for the asset.
func (c *Client) ListAssetRevisions(ctx context.Context, params *ListAssetRevisionsInput, optFns ...func(*Options)) (*ListAssetRevisionsOutput, error) {
	if params == nil {
		params = &ListAssetRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssetRevisions", params, optFns, c.addOperationListAssetRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetRevisionsInput struct {

	// The identifier of the domain.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the asset.
	//
	// This member is required.
	Identifier *string

	// The maximum number of revisions to return in a single call to ListAssetRevisions
	// . When the number of revisions to be listed is greater than the value of
	// MaxResults , the response contains a NextToken value that you can use in a
	// subsequent call to ListAssetRevisions to list the next set of revisions.
	MaxResults *int32

	// When the number of revisions is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of revisions, the response includes a pagination token
	// named NextToken . You can specify this NextToken value in a subsequent call to
	// ListAssetRevisions to list the next set of revisions.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssetRevisionsOutput struct {

	// The results of the ListAssetRevisions action.
	Items []types.AssetRevision

	// When the number of revisions is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of revisions, the response includes a pagination token
	// named NextToken . You can specify this NextToken value in a subsequent call to
	// ListAssetRevisions to list the next set of revisions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssetRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssetRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssetRevisions{}, middleware.After)
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
	if err = addOpListAssetRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssetRevisions(options.Region), middleware.Before); err != nil {
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

// ListAssetRevisionsAPIClient is a client that implements the ListAssetRevisions
// operation.
type ListAssetRevisionsAPIClient interface {
	ListAssetRevisions(context.Context, *ListAssetRevisionsInput, ...func(*Options)) (*ListAssetRevisionsOutput, error)
}

var _ ListAssetRevisionsAPIClient = (*Client)(nil)

// ListAssetRevisionsPaginatorOptions is the paginator options for
// ListAssetRevisions
type ListAssetRevisionsPaginatorOptions struct {
	// The maximum number of revisions to return in a single call to ListAssetRevisions
	// . When the number of revisions to be listed is greater than the value of
	// MaxResults , the response contains a NextToken value that you can use in a
	// subsequent call to ListAssetRevisions to list the next set of revisions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetRevisionsPaginator is a paginator for ListAssetRevisions
type ListAssetRevisionsPaginator struct {
	options   ListAssetRevisionsPaginatorOptions
	client    ListAssetRevisionsAPIClient
	params    *ListAssetRevisionsInput
	nextToken *string
	firstPage bool
}

// NewListAssetRevisionsPaginator returns a new ListAssetRevisionsPaginator
func NewListAssetRevisionsPaginator(client ListAssetRevisionsAPIClient, params *ListAssetRevisionsInput, optFns ...func(*ListAssetRevisionsPaginatorOptions)) *ListAssetRevisionsPaginator {
	if params == nil {
		params = &ListAssetRevisionsInput{}
	}

	options := ListAssetRevisionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssetRevisionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetRevisionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssetRevisions page.
func (p *ListAssetRevisionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetRevisionsOutput, error) {
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

	result, err := p.client.ListAssetRevisions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssetRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "ListAssetRevisions",
	}
}
