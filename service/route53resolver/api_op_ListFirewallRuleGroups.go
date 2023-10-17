// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the minimal high-level information for the rule groups that you have
// defined. A single call might return only a partial list of the rule groups. For
// information, see MaxResults .
func (c *Client) ListFirewallRuleGroups(ctx context.Context, params *ListFirewallRuleGroupsInput, optFns ...func(*Options)) (*ListFirewallRuleGroupsOutput, error) {
	if params == nil {
		params = &ListFirewallRuleGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFirewallRuleGroups", params, optFns, c.addOperationListFirewallRuleGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFirewallRuleGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFirewallRuleGroupsInput struct {

	// The maximum number of objects that you want Resolver to return for this
	// request. If more objects are available, in the response, Resolver provides a
	// NextToken value that you can use in a subsequent call to get the next batch of
	// objects. If you don't specify a value for MaxResults , Resolver returns up to
	// 100 objects.
	MaxResults *int32

	// For the first call to this list request, omit this value. When you request a
	// list of objects, Resolver returns at most the number of objects specified in
	// MaxResults . If more objects are available for retrieval, Resolver returns a
	// NextToken value in the response. To retrieve the next batch of objects, use the
	// token that was returned for the prior request in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFirewallRuleGroupsOutput struct {

	// A list of your firewall rule groups. This might be a partial list of the rule
	// groups that you have defined. For information, see MaxResults .
	FirewallRuleGroups []types.FirewallRuleGroupMetadata

	// If objects are still available for retrieval, Resolver returns this token in
	// the response. To retrieve the next batch of objects, provide this token in your
	// next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFirewallRuleGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFirewallRuleGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFirewallRuleGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFirewallRuleGroups(options.Region), middleware.Before); err != nil {
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

// ListFirewallRuleGroupsAPIClient is a client that implements the
// ListFirewallRuleGroups operation.
type ListFirewallRuleGroupsAPIClient interface {
	ListFirewallRuleGroups(context.Context, *ListFirewallRuleGroupsInput, ...func(*Options)) (*ListFirewallRuleGroupsOutput, error)
}

var _ ListFirewallRuleGroupsAPIClient = (*Client)(nil)

// ListFirewallRuleGroupsPaginatorOptions is the paginator options for
// ListFirewallRuleGroups
type ListFirewallRuleGroupsPaginatorOptions struct {
	// The maximum number of objects that you want Resolver to return for this
	// request. If more objects are available, in the response, Resolver provides a
	// NextToken value that you can use in a subsequent call to get the next batch of
	// objects. If you don't specify a value for MaxResults , Resolver returns up to
	// 100 objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFirewallRuleGroupsPaginator is a paginator for ListFirewallRuleGroups
type ListFirewallRuleGroupsPaginator struct {
	options   ListFirewallRuleGroupsPaginatorOptions
	client    ListFirewallRuleGroupsAPIClient
	params    *ListFirewallRuleGroupsInput
	nextToken *string
	firstPage bool
}

// NewListFirewallRuleGroupsPaginator returns a new ListFirewallRuleGroupsPaginator
func NewListFirewallRuleGroupsPaginator(client ListFirewallRuleGroupsAPIClient, params *ListFirewallRuleGroupsInput, optFns ...func(*ListFirewallRuleGroupsPaginatorOptions)) *ListFirewallRuleGroupsPaginator {
	if params == nil {
		params = &ListFirewallRuleGroupsInput{}
	}

	options := ListFirewallRuleGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFirewallRuleGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFirewallRuleGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFirewallRuleGroups page.
func (p *ListFirewallRuleGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFirewallRuleGroupsOutput, error) {
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

	result, err := p.client.ListFirewallRuleGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFirewallRuleGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListFirewallRuleGroups",
	}
}
