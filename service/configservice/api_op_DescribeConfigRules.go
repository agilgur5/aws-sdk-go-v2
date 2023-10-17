// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details about your Config rules.
func (c *Client) DescribeConfigRules(ctx context.Context, params *DescribeConfigRulesInput, optFns ...func(*Options)) (*DescribeConfigRulesOutput, error) {
	if params == nil {
		params = &DescribeConfigRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigRules", params, optFns, c.addOperationDescribeConfigRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigRulesInput struct {

	// The names of the Config rules for which you want details. If you do not specify
	// any names, Config returns details for all your rules.
	ConfigRuleNames []string

	// Returns a list of Detective or Proactive Config rules. By default, this API
	// returns an unfiltered list. For more information on Detective or Proactive
	// Config rules, see Evaluation Mode  (https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config-rules.html)
	// in the Config Developer Guide.
	Filters *types.DescribeConfigRulesFilters

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeConfigRulesOutput struct {

	// The details about your Config rules.
	ConfigRules []types.ConfigRule

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigRules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigRules(options.Region), middleware.Before); err != nil {
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

// DescribeConfigRulesAPIClient is a client that implements the
// DescribeConfigRules operation.
type DescribeConfigRulesAPIClient interface {
	DescribeConfigRules(context.Context, *DescribeConfigRulesInput, ...func(*Options)) (*DescribeConfigRulesOutput, error)
}

var _ DescribeConfigRulesAPIClient = (*Client)(nil)

// DescribeConfigRulesPaginatorOptions is the paginator options for
// DescribeConfigRules
type DescribeConfigRulesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConfigRulesPaginator is a paginator for DescribeConfigRules
type DescribeConfigRulesPaginator struct {
	options   DescribeConfigRulesPaginatorOptions
	client    DescribeConfigRulesAPIClient
	params    *DescribeConfigRulesInput
	nextToken *string
	firstPage bool
}

// NewDescribeConfigRulesPaginator returns a new DescribeConfigRulesPaginator
func NewDescribeConfigRulesPaginator(client DescribeConfigRulesAPIClient, params *DescribeConfigRulesInput, optFns ...func(*DescribeConfigRulesPaginatorOptions)) *DescribeConfigRulesPaginator {
	if params == nil {
		params = &DescribeConfigRulesInput{}
	}

	options := DescribeConfigRulesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConfigRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConfigRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeConfigRules page.
func (p *DescribeConfigRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConfigRulesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeConfigRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeConfigRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigRules",
	}
}
