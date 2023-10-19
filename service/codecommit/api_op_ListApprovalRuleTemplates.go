// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all approval rule templates in the specified Amazon Web Services Region
// in your Amazon Web Services account. If an Amazon Web Services Region is not
// specified, the Amazon Web Services Region where you are signed in is used.
func (c *Client) ListApprovalRuleTemplates(ctx context.Context, params *ListApprovalRuleTemplatesInput, optFns ...func(*Options)) (*ListApprovalRuleTemplatesOutput, error) {
	if params == nil {
		params = &ListApprovalRuleTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApprovalRuleTemplates", params, optFns, c.addOperationListApprovalRuleTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApprovalRuleTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApprovalRuleTemplatesInput struct {

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListApprovalRuleTemplatesOutput struct {

	// The names of all the approval rule templates found in the Amazon Web Services
	// Region for your Amazon Web Services account.
	ApprovalRuleTemplateNames []string

	// An enumeration token that allows the operation to batch the next results of the
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApprovalRuleTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApprovalRuleTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApprovalRuleTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApprovalRuleTemplates(options.Region), middleware.Before); err != nil {
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
		operation: "ListApprovalRuleTemplates",
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

// ListApprovalRuleTemplatesAPIClient is a client that implements the
// ListApprovalRuleTemplates operation.
type ListApprovalRuleTemplatesAPIClient interface {
	ListApprovalRuleTemplates(context.Context, *ListApprovalRuleTemplatesInput, ...func(*Options)) (*ListApprovalRuleTemplatesOutput, error)
}

var _ ListApprovalRuleTemplatesAPIClient = (*Client)(nil)

// ListApprovalRuleTemplatesPaginatorOptions is the paginator options for
// ListApprovalRuleTemplates
type ListApprovalRuleTemplatesPaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApprovalRuleTemplatesPaginator is a paginator for ListApprovalRuleTemplates
type ListApprovalRuleTemplatesPaginator struct {
	options   ListApprovalRuleTemplatesPaginatorOptions
	client    ListApprovalRuleTemplatesAPIClient
	params    *ListApprovalRuleTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListApprovalRuleTemplatesPaginator returns a new
// ListApprovalRuleTemplatesPaginator
func NewListApprovalRuleTemplatesPaginator(client ListApprovalRuleTemplatesAPIClient, params *ListApprovalRuleTemplatesInput, optFns ...func(*ListApprovalRuleTemplatesPaginatorOptions)) *ListApprovalRuleTemplatesPaginator {
	if params == nil {
		params = &ListApprovalRuleTemplatesInput{}
	}

	options := ListApprovalRuleTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApprovalRuleTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApprovalRuleTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApprovalRuleTemplates page.
func (p *ListApprovalRuleTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApprovalRuleTemplatesOutput, error) {
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

	result, err := p.client.ListApprovalRuleTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApprovalRuleTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "ListApprovalRuleTemplates",
	}
}
