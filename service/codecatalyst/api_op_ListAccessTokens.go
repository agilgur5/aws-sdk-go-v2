// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all personal access tokens (PATs) associated with the user who calls the
// API. You can only list PATs associated with your Amazon Web Services Builder ID.
func (c *Client) ListAccessTokens(ctx context.Context, params *ListAccessTokensInput, optFns ...func(*Options)) (*ListAccessTokensOutput, error) {
	if params == nil {
		params = &ListAccessTokensInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessTokens", params, optFns, c.addOperationListAccessTokensMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessTokensOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessTokensInput struct {

	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	MaxResults *int32

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccessTokensOutput struct {

	// A list of personal access tokens (PATs) associated with the calling user
	// identity.
	//
	// This member is required.
	Items []types.AccessTokenSummary

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessTokensMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAccessTokens{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccessTokens{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessTokens(options.Region), middleware.Before); err != nil {
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
		operation: "ListAccessTokens",
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

// ListAccessTokensAPIClient is a client that implements the ListAccessTokens
// operation.
type ListAccessTokensAPIClient interface {
	ListAccessTokens(context.Context, *ListAccessTokensInput, ...func(*Options)) (*ListAccessTokensOutput, error)
}

var _ ListAccessTokensAPIClient = (*Client)(nil)

// ListAccessTokensPaginatorOptions is the paginator options for ListAccessTokens
type ListAccessTokensPaginatorOptions struct {
	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessTokensPaginator is a paginator for ListAccessTokens
type ListAccessTokensPaginator struct {
	options   ListAccessTokensPaginatorOptions
	client    ListAccessTokensAPIClient
	params    *ListAccessTokensInput
	nextToken *string
	firstPage bool
}

// NewListAccessTokensPaginator returns a new ListAccessTokensPaginator
func NewListAccessTokensPaginator(client ListAccessTokensAPIClient, params *ListAccessTokensInput, optFns ...func(*ListAccessTokensPaginatorOptions)) *ListAccessTokensPaginator {
	if params == nil {
		params = &ListAccessTokensInput{}
	}

	options := ListAccessTokensPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessTokensPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessTokensPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessTokens page.
func (p *ListAccessTokensPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessTokensOutput, error) {
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

	result, err := p.client.ListAccessTokens(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccessTokens(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccessTokens",
	}
}
