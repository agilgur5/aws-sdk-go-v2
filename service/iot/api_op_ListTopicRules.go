// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the rules for the specific topic. Requires permission to access the
// ListTopicRules (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListTopicRules(ctx context.Context, params *ListTopicRulesInput, optFns ...func(*Options)) (*ListTopicRulesOutput, error) {
	if params == nil {
		params = &ListTopicRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTopicRules", params, optFns, c.addOperationListTopicRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTopicRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListTopicRules operation.
type ListTopicRulesInput struct {

	// The maximum number of results to return.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// Specifies whether the rule is disabled.
	RuleDisabled *bool

	// The topic.
	Topic *string

	noSmithyDocumentSerde
}

// The output from the ListTopicRules operation.
type ListTopicRulesOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The rules.
	Rules []types.TopicRuleListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTopicRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTopicRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTopicRules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTopicRules(options.Region), middleware.Before); err != nil {
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
		operation: "ListTopicRules",
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

// ListTopicRulesAPIClient is a client that implements the ListTopicRules
// operation.
type ListTopicRulesAPIClient interface {
	ListTopicRules(context.Context, *ListTopicRulesInput, ...func(*Options)) (*ListTopicRulesOutput, error)
}

var _ ListTopicRulesAPIClient = (*Client)(nil)

// ListTopicRulesPaginatorOptions is the paginator options for ListTopicRules
type ListTopicRulesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTopicRulesPaginator is a paginator for ListTopicRules
type ListTopicRulesPaginator struct {
	options   ListTopicRulesPaginatorOptions
	client    ListTopicRulesAPIClient
	params    *ListTopicRulesInput
	nextToken *string
	firstPage bool
}

// NewListTopicRulesPaginator returns a new ListTopicRulesPaginator
func NewListTopicRulesPaginator(client ListTopicRulesAPIClient, params *ListTopicRulesInput, optFns ...func(*ListTopicRulesPaginatorOptions)) *ListTopicRulesPaginator {
	if params == nil {
		params = &ListTopicRulesInput{}
	}

	options := ListTopicRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTopicRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTopicRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTopicRules page.
func (p *ListTopicRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTopicRulesOutput, error) {
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

	result, err := p.client.ListTopicRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTopicRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "ListTopicRules",
	}
}
