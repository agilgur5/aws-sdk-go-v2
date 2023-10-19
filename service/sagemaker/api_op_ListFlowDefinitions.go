// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the flow definitions in your account.
func (c *Client) ListFlowDefinitions(ctx context.Context, params *ListFlowDefinitionsInput, optFns ...func(*Options)) (*ListFlowDefinitionsOutput, error) {
	if params == nil {
		params = &ListFlowDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFlowDefinitions", params, optFns, c.addOperationListFlowDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFlowDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlowDefinitionsInput struct {

	// A filter that returns only flow definitions with a creation time greater than
	// or equal to the specified timestamp.
	CreationTimeAfter *time.Time

	// A filter that returns only flow definitions that were created before the
	// specified timestamp.
	CreationTimeBefore *time.Time

	// The total number of items to return. If the total number of available items is
	// more than the value specified in MaxResults , then a NextToken will be provided
	// in the output that you can use to resume pagination.
	MaxResults *int32

	// A token to resume pagination.
	NextToken *string

	// An optional value that specifies whether you want the results sorted in
	// Ascending or Descending order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListFlowDefinitionsOutput struct {

	// An array of objects describing the flow definitions.
	//
	// This member is required.
	FlowDefinitionSummaries []types.FlowDefinitionSummary

	// A token to resume pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFlowDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFlowDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFlowDefinitions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFlowDefinitions(options.Region), middleware.Before); err != nil {
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
		operation: "ListFlowDefinitions",
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

// ListFlowDefinitionsAPIClient is a client that implements the
// ListFlowDefinitions operation.
type ListFlowDefinitionsAPIClient interface {
	ListFlowDefinitions(context.Context, *ListFlowDefinitionsInput, ...func(*Options)) (*ListFlowDefinitionsOutput, error)
}

var _ ListFlowDefinitionsAPIClient = (*Client)(nil)

// ListFlowDefinitionsPaginatorOptions is the paginator options for
// ListFlowDefinitions
type ListFlowDefinitionsPaginatorOptions struct {
	// The total number of items to return. If the total number of available items is
	// more than the value specified in MaxResults , then a NextToken will be provided
	// in the output that you can use to resume pagination.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFlowDefinitionsPaginator is a paginator for ListFlowDefinitions
type ListFlowDefinitionsPaginator struct {
	options   ListFlowDefinitionsPaginatorOptions
	client    ListFlowDefinitionsAPIClient
	params    *ListFlowDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListFlowDefinitionsPaginator returns a new ListFlowDefinitionsPaginator
func NewListFlowDefinitionsPaginator(client ListFlowDefinitionsAPIClient, params *ListFlowDefinitionsInput, optFns ...func(*ListFlowDefinitionsPaginatorOptions)) *ListFlowDefinitionsPaginator {
	if params == nil {
		params = &ListFlowDefinitionsInput{}
	}

	options := ListFlowDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFlowDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFlowDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFlowDefinitions page.
func (p *ListFlowDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFlowDefinitionsOutput, error) {
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

	result, err := p.client.ListFlowDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFlowDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListFlowDefinitions",
	}
}
