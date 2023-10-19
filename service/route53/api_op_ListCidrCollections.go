// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of CIDR collections in the Amazon Web Services account
// (metadata only).
func (c *Client) ListCidrCollections(ctx context.Context, params *ListCidrCollectionsInput, optFns ...func(*Options)) (*ListCidrCollectionsOutput, error) {
	if params == nil {
		params = &ListCidrCollectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCidrCollections", params, optFns, c.addOperationListCidrCollectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCidrCollectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCidrCollectionsInput struct {

	// The maximum number of CIDR collections to return in the response.
	MaxResults *int32

	// An opaque pagination token to indicate where the service is to begin
	// enumerating results. If no value is provided, the listing of results starts from
	// the beginning.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCidrCollectionsOutput struct {

	// A complex type with information about the CIDR collection.
	CidrCollections []types.CollectionSummary

	// An opaque pagination token to indicate where the service is to begin
	// enumerating results. If no value is provided, the listing of results starts from
	// the beginning.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCidrCollectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListCidrCollections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListCidrCollections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCidrCollections(options.Region), middleware.Before); err != nil {
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
		operation: "ListCidrCollections",
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

// ListCidrCollectionsAPIClient is a client that implements the
// ListCidrCollections operation.
type ListCidrCollectionsAPIClient interface {
	ListCidrCollections(context.Context, *ListCidrCollectionsInput, ...func(*Options)) (*ListCidrCollectionsOutput, error)
}

var _ ListCidrCollectionsAPIClient = (*Client)(nil)

// ListCidrCollectionsPaginatorOptions is the paginator options for
// ListCidrCollections
type ListCidrCollectionsPaginatorOptions struct {
	// The maximum number of CIDR collections to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCidrCollectionsPaginator is a paginator for ListCidrCollections
type ListCidrCollectionsPaginator struct {
	options   ListCidrCollectionsPaginatorOptions
	client    ListCidrCollectionsAPIClient
	params    *ListCidrCollectionsInput
	nextToken *string
	firstPage bool
}

// NewListCidrCollectionsPaginator returns a new ListCidrCollectionsPaginator
func NewListCidrCollectionsPaginator(client ListCidrCollectionsAPIClient, params *ListCidrCollectionsInput, optFns ...func(*ListCidrCollectionsPaginatorOptions)) *ListCidrCollectionsPaginator {
	if params == nil {
		params = &ListCidrCollectionsInput{}
	}

	options := ListCidrCollectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCidrCollectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCidrCollectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCidrCollections page.
func (p *ListCidrCollectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCidrCollectionsOutput, error) {
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

	result, err := p.client.ListCidrCollections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCidrCollections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListCidrCollections",
	}
}
