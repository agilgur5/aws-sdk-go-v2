// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the pricing rules that are associated with a pricing plan.
func (c *Client) ListPricingRulesAssociatedToPricingPlan(ctx context.Context, params *ListPricingRulesAssociatedToPricingPlanInput, optFns ...func(*Options)) (*ListPricingRulesAssociatedToPricingPlanOutput, error) {
	if params == nil {
		params = &ListPricingRulesAssociatedToPricingPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPricingRulesAssociatedToPricingPlan", params, optFns, c.addOperationListPricingRulesAssociatedToPricingPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPricingRulesAssociatedToPricingPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPricingRulesAssociatedToPricingPlanInput struct {

	// The Amazon Resource Name (ARN) of the pricing plan for which associations are
	// to be listed.
	//
	// This member is required.
	PricingPlanArn *string

	// The billing period for which the pricing rule associations are to be listed.
	BillingPeriod *string

	// The optional maximum number of pricing rule associations to retrieve.
	MaxResults *int32

	// The optional pagination token returned by a previous call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPricingRulesAssociatedToPricingPlanOutput struct {

	// The billing period for which the pricing rule associations are listed.
	BillingPeriod *string

	// The pagination token to be used on subsequent calls.
	NextToken *string

	// The Amazon Resource Name (ARN) of the pricing plan for which associations are
	// listed.
	PricingPlanArn *string

	// A list containing pricing rules that are associated with the requested pricing
	// plan.
	PricingRuleArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPricingRulesAssociatedToPricingPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPricingRulesAssociatedToPricingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPricingRulesAssociatedToPricingPlan{}, middleware.After)
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
	if err = addOpListPricingRulesAssociatedToPricingPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPricingRulesAssociatedToPricingPlan(options.Region), middleware.Before); err != nil {
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

// ListPricingRulesAssociatedToPricingPlanAPIClient is a client that implements
// the ListPricingRulesAssociatedToPricingPlan operation.
type ListPricingRulesAssociatedToPricingPlanAPIClient interface {
	ListPricingRulesAssociatedToPricingPlan(context.Context, *ListPricingRulesAssociatedToPricingPlanInput, ...func(*Options)) (*ListPricingRulesAssociatedToPricingPlanOutput, error)
}

var _ ListPricingRulesAssociatedToPricingPlanAPIClient = (*Client)(nil)

// ListPricingRulesAssociatedToPricingPlanPaginatorOptions is the paginator
// options for ListPricingRulesAssociatedToPricingPlan
type ListPricingRulesAssociatedToPricingPlanPaginatorOptions struct {
	// The optional maximum number of pricing rule associations to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPricingRulesAssociatedToPricingPlanPaginator is a paginator for
// ListPricingRulesAssociatedToPricingPlan
type ListPricingRulesAssociatedToPricingPlanPaginator struct {
	options   ListPricingRulesAssociatedToPricingPlanPaginatorOptions
	client    ListPricingRulesAssociatedToPricingPlanAPIClient
	params    *ListPricingRulesAssociatedToPricingPlanInput
	nextToken *string
	firstPage bool
}

// NewListPricingRulesAssociatedToPricingPlanPaginator returns a new
// ListPricingRulesAssociatedToPricingPlanPaginator
func NewListPricingRulesAssociatedToPricingPlanPaginator(client ListPricingRulesAssociatedToPricingPlanAPIClient, params *ListPricingRulesAssociatedToPricingPlanInput, optFns ...func(*ListPricingRulesAssociatedToPricingPlanPaginatorOptions)) *ListPricingRulesAssociatedToPricingPlanPaginator {
	if params == nil {
		params = &ListPricingRulesAssociatedToPricingPlanInput{}
	}

	options := ListPricingRulesAssociatedToPricingPlanPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPricingRulesAssociatedToPricingPlanPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPricingRulesAssociatedToPricingPlanPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPricingRulesAssociatedToPricingPlan page.
func (p *ListPricingRulesAssociatedToPricingPlanPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPricingRulesAssociatedToPricingPlanOutput, error) {
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

	result, err := p.client.ListPricingRulesAssociatedToPricingPlan(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPricingRulesAssociatedToPricingPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "ListPricingRulesAssociatedToPricingPlan",
	}
}
