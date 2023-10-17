// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the pull through cache rules for a registry.
func (c *Client) DescribePullThroughCacheRules(ctx context.Context, params *DescribePullThroughCacheRulesInput, optFns ...func(*Options)) (*DescribePullThroughCacheRulesOutput, error) {
	if params == nil {
		params = &DescribePullThroughCacheRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePullThroughCacheRules", params, optFns, c.addOperationDescribePullThroughCacheRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePullThroughCacheRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePullThroughCacheRulesInput struct {

	// The Amazon ECR repository prefixes associated with the pull through cache rules
	// to return. If no repository prefix value is specified, all pull through cache
	// rules are returned.
	EcrRepositoryPrefixes []string

	// The maximum number of pull through cache rules returned by
	// DescribePullThroughCacheRulesRequest in paginated output. When this parameter is
	// used, DescribePullThroughCacheRulesRequest only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another
	// DescribePullThroughCacheRulesRequest request with the returned nextToken value.
	// This value can be between 1 and 1000. If this parameter is not used, then
	// DescribePullThroughCacheRulesRequest returns up to 100 results and a nextToken
	// value, if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// DescribePullThroughCacheRulesRequest request where maxResults was used and the
	// results exceeded the value of that parameter. Pagination continues from the end
	// of the previous results that returned the nextToken value. This value is null
	// when there are no more results to return.
	NextToken *string

	// The Amazon Web Services account ID associated with the registry to return the
	// pull through cache rules for. If you do not specify a registry, the default
	// registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type DescribePullThroughCacheRulesOutput struct {

	// The nextToken value to include in a future DescribePullThroughCacheRulesRequest
	// request. When the results of a DescribePullThroughCacheRulesRequest request
	// exceed maxResults , this value can be used to retrieve the next page of results.
	// This value is null when there are no more results to return.
	NextToken *string

	// The details of the pull through cache rules.
	PullThroughCacheRules []types.PullThroughCacheRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePullThroughCacheRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePullThroughCacheRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePullThroughCacheRules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePullThroughCacheRules(options.Region), middleware.Before); err != nil {
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

// DescribePullThroughCacheRulesAPIClient is a client that implements the
// DescribePullThroughCacheRules operation.
type DescribePullThroughCacheRulesAPIClient interface {
	DescribePullThroughCacheRules(context.Context, *DescribePullThroughCacheRulesInput, ...func(*Options)) (*DescribePullThroughCacheRulesOutput, error)
}

var _ DescribePullThroughCacheRulesAPIClient = (*Client)(nil)

// DescribePullThroughCacheRulesPaginatorOptions is the paginator options for
// DescribePullThroughCacheRules
type DescribePullThroughCacheRulesPaginatorOptions struct {
	// The maximum number of pull through cache rules returned by
	// DescribePullThroughCacheRulesRequest in paginated output. When this parameter is
	// used, DescribePullThroughCacheRulesRequest only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another
	// DescribePullThroughCacheRulesRequest request with the returned nextToken value.
	// This value can be between 1 and 1000. If this parameter is not used, then
	// DescribePullThroughCacheRulesRequest returns up to 100 results and a nextToken
	// value, if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePullThroughCacheRulesPaginator is a paginator for
// DescribePullThroughCacheRules
type DescribePullThroughCacheRulesPaginator struct {
	options   DescribePullThroughCacheRulesPaginatorOptions
	client    DescribePullThroughCacheRulesAPIClient
	params    *DescribePullThroughCacheRulesInput
	nextToken *string
	firstPage bool
}

// NewDescribePullThroughCacheRulesPaginator returns a new
// DescribePullThroughCacheRulesPaginator
func NewDescribePullThroughCacheRulesPaginator(client DescribePullThroughCacheRulesAPIClient, params *DescribePullThroughCacheRulesInput, optFns ...func(*DescribePullThroughCacheRulesPaginatorOptions)) *DescribePullThroughCacheRulesPaginator {
	if params == nil {
		params = &DescribePullThroughCacheRulesInput{}
	}

	options := DescribePullThroughCacheRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePullThroughCacheRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePullThroughCacheRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePullThroughCacheRules page.
func (p *DescribePullThroughCacheRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePullThroughCacheRulesOutput, error) {
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

	result, err := p.client.DescribePullThroughCacheRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePullThroughCacheRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "DescribePullThroughCacheRules",
	}
}
