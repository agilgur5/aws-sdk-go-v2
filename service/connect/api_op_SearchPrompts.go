// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches prompts in an Amazon Connect instance, with optional filtering.
func (c *Client) SearchPrompts(ctx context.Context, params *SearchPromptsInput, optFns ...func(*Options)) (*SearchPromptsOutput, error) {
	if params == nil {
		params = &SearchPromptsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchPrompts", params, optFns, c.addOperationSearchPromptsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchPromptsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchPromptsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The search criteria to be used to return prompts.
	SearchCriteria *types.PromptSearchCriteria

	// Filters to be applied to search results.
	SearchFilter *types.PromptSearchFilter

	noSmithyDocumentSerde
}

type SearchPromptsOutput struct {

	// The total number of quick connects which matched your search query.
	ApproximateTotalCount *int64

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the prompts.
	Prompts []types.Prompt

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchPromptsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchPrompts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchPrompts{}, middleware.After)
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
	if err = addOpSearchPromptsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchPrompts(options.Region), middleware.Before); err != nil {
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

// SearchPromptsAPIClient is a client that implements the SearchPrompts operation.
type SearchPromptsAPIClient interface {
	SearchPrompts(context.Context, *SearchPromptsInput, ...func(*Options)) (*SearchPromptsOutput, error)
}

var _ SearchPromptsAPIClient = (*Client)(nil)

// SearchPromptsPaginatorOptions is the paginator options for SearchPrompts
type SearchPromptsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchPromptsPaginator is a paginator for SearchPrompts
type SearchPromptsPaginator struct {
	options   SearchPromptsPaginatorOptions
	client    SearchPromptsAPIClient
	params    *SearchPromptsInput
	nextToken *string
	firstPage bool
}

// NewSearchPromptsPaginator returns a new SearchPromptsPaginator
func NewSearchPromptsPaginator(client SearchPromptsAPIClient, params *SearchPromptsInput, optFns ...func(*SearchPromptsPaginatorOptions)) *SearchPromptsPaginator {
	if params == nil {
		params = &SearchPromptsInput{}
	}

	options := SearchPromptsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchPromptsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchPromptsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchPrompts page.
func (p *SearchPromptsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchPromptsOutput, error) {
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

	result, err := p.client.SearchPrompts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchPrompts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "SearchPrompts",
	}
}
