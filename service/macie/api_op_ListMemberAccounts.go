// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// (Discontinued) Lists all Amazon Macie Classic member accounts for the current
// Macie Classic administrator account.
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

	// (Discontinued) Use this parameter to indicate the maximum number of items that
	// you want in the response. The default value is 250.
	MaxResults *int32

	// (Discontinued) Use this parameter when paginating results. Set the value of
	// this parameter to null on your first call to the ListMemberAccounts action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// nextToken from the previous response to continue listing data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMemberAccountsOutput struct {

	// (Discontinued) A list of the Amazon Macie Classic member accounts returned by
	// the action. The current Macie Classic administrator account is also included in
	// this list.
	MemberAccounts []types.MemberAccount

	// (Discontinued) When a response is generated, if there is more data to be
	// listed, this parameter is present in the response and contains the value to use
	// for the nextToken parameter in a subsequent pagination request. If there is no
	// more data to be listed, this parameter is set to null.
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
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "ListMemberAccounts",
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

// ListMemberAccountsAPIClient is a client that implements the ListMemberAccounts
// operation.
type ListMemberAccountsAPIClient interface {
	ListMemberAccounts(context.Context, *ListMemberAccountsInput, ...func(*Options)) (*ListMemberAccountsOutput, error)
}

var _ ListMemberAccountsAPIClient = (*Client)(nil)

// ListMemberAccountsPaginatorOptions is the paginator options for
// ListMemberAccounts
type ListMemberAccountsPaginatorOptions struct {
	// (Discontinued) Use this parameter to indicate the maximum number of items that
	// you want in the response. The default value is 250.
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
		SigningName:   "macie",
		OperationName: "ListMemberAccounts",
	}
}
