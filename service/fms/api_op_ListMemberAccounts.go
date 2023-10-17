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

// Returns a MemberAccounts object that lists the member accounts in the
// administrator's Amazon Web Services organization. Either an Firewall Manager
// administrator or the organization's management account can make this request.
func (c *Client) ListMemberAccounts(ctx context.Context, params *ListMemberAccountsInput, optFns ...func(*Options)) (*ListMemberAccountsOutput, error) {
	if params == nil {
		params = &ListMemberAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMemberAccounts", params, optFns, c.addOperationListMemberAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMemberAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMemberAccountsInput struct {

	// Specifies the number of member account IDs that you want Firewall Manager to
	// return for this request. If you have more IDs than the number that you specify
	// for MaxResults , the response includes a NextToken value that you can use to
	// get another batch of member account IDs.
	MaxResults *int32

	// If you specify a value for MaxResults and you have more account IDs than the
	// number that you specify for MaxResults , Firewall Manager returns a NextToken
	// value in the response that allows you to list another group of IDs. For the
	// second and subsequent ListMemberAccountsRequest requests, specify the value of
	// NextToken from the previous response to get information about another batch of
	// member account IDs.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMemberAccountsOutput struct {

	// An array of account IDs.
	MemberAccounts []string

	// If you have more member account IDs than the number that you specified for
	// MaxResults in the request, the response includes a NextToken value. To list
	// more IDs, submit another ListMemberAccounts request, and specify the NextToken
	// value from the response in the NextToken value in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMemberAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMemberAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMemberAccounts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMemberAccounts(options.Region), middleware.Before); err != nil {
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

// ListMemberAccountsAPIClient is a client that implements the ListMemberAccounts
// operation.
type ListMemberAccountsAPIClient interface {
	ListMemberAccounts(context.Context, *ListMemberAccountsInput, ...func(*Options)) (*ListMemberAccountsOutput, error)
}

var _ ListMemberAccountsAPIClient = (*Client)(nil)

// ListMemberAccountsPaginatorOptions is the paginator options for
// ListMemberAccounts
type ListMemberAccountsPaginatorOptions struct {
	// Specifies the number of member account IDs that you want Firewall Manager to
	// return for this request. If you have more IDs than the number that you specify
	// for MaxResults , the response includes a NextToken value that you can use to
	// get another batch of member account IDs.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMemberAccountsPaginator is a paginator for ListMemberAccounts
type ListMemberAccountsPaginator struct {
	options   ListMemberAccountsPaginatorOptions
	client    ListMemberAccountsAPIClient
	params    *ListMemberAccountsInput
	nextToken *string
	firstPage bool
}

// NewListMemberAccountsPaginator returns a new ListMemberAccountsPaginator
func NewListMemberAccountsPaginator(client ListMemberAccountsAPIClient, params *ListMemberAccountsInput, optFns ...func(*ListMemberAccountsPaginatorOptions)) *ListMemberAccountsPaginator {
	if params == nil {
		params = &ListMemberAccountsInput{}
	}

	options := ListMemberAccountsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMemberAccountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMemberAccountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMemberAccounts page.
func (p *ListMemberAccountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMemberAccountsOutput, error) {
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

	result, err := p.client.ListMemberAccounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMemberAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListMemberAccounts",
	}
}
