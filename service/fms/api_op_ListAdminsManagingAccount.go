// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the accounts that are managing the specified Organizations member
// account. This is useful for any member account so that they can view the
// accounts who are managing their account. This operation only returns the
// managing administrators that have the requested account within their AdminScope .
func (c *Client) ListAdminsManagingAccount(ctx context.Context, params *ListAdminsManagingAccountInput, optFns ...func(*Options)) (*ListAdminsManagingAccountOutput, error) {
	if params == nil {
		params = &ListAdminsManagingAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAdminsManagingAccount", params, optFns, c.addOperationListAdminsManagingAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAdminsManagingAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAdminsManagingAccountInput struct {

	// The maximum number of objects that you want Firewall Manager to return for this
	// request. If more objects are available, in the response, Firewall Manager
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	MaxResults *int32

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Firewall Manager returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAdminsManagingAccountOutput struct {

	// The list of accounts who manage member accounts within their AdminScope .
	AdminAccounts []string

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Firewall Manager returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAdminsManagingAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAdminsManagingAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAdminsManagingAccount{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAdminsManagingAccount(options.Region), middleware.Before); err != nil {
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
		operation: "ListAdminsManagingAccount",
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

// ListAdminsManagingAccountAPIClient is a client that implements the
// ListAdminsManagingAccount operation.
type ListAdminsManagingAccountAPIClient interface {
	ListAdminsManagingAccount(context.Context, *ListAdminsManagingAccountInput, ...func(*Options)) (*ListAdminsManagingAccountOutput, error)
}

var _ ListAdminsManagingAccountAPIClient = (*Client)(nil)

// ListAdminsManagingAccountPaginatorOptions is the paginator options for
// ListAdminsManagingAccount
type ListAdminsManagingAccountPaginatorOptions struct {
	// The maximum number of objects that you want Firewall Manager to return for this
	// request. If more objects are available, in the response, Firewall Manager
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAdminsManagingAccountPaginator is a paginator for ListAdminsManagingAccount
type ListAdminsManagingAccountPaginator struct {
	options   ListAdminsManagingAccountPaginatorOptions
	client    ListAdminsManagingAccountAPIClient
	params    *ListAdminsManagingAccountInput
	nextToken *string
	firstPage bool
}

// NewListAdminsManagingAccountPaginator returns a new
// ListAdminsManagingAccountPaginator
func NewListAdminsManagingAccountPaginator(client ListAdminsManagingAccountAPIClient, params *ListAdminsManagingAccountInput, optFns ...func(*ListAdminsManagingAccountPaginatorOptions)) *ListAdminsManagingAccountPaginator {
	if params == nil {
		params = &ListAdminsManagingAccountInput{}
	}

	options := ListAdminsManagingAccountPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAdminsManagingAccountPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAdminsManagingAccountPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAdminsManagingAccount page.
func (p *ListAdminsManagingAccountPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAdminsManagingAccountOutput, error) {
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

	result, err := p.client.ListAdminsManagingAccount(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAdminsManagingAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListAdminsManagingAccount",
	}
}
