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

// Lists models created with the CreateModel API.
func (c *Client) ListModels(ctx context.Context, params *ListModelsInput, optFns ...func(*Options)) (*ListModelsOutput, error) {
	if params == nil {
		params = &ListModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModels", params, optFns, c.addOperationListModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelsInput struct {

	// A filter that returns only models with a creation time greater than or equal to
	// the specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only models created before the specified time (timestamp).
	CreationTimeBefore *time.Time

	// The maximum number of models to return in the response.
	MaxResults *int32

	// A string in the model name. This filter returns only models whose name contains
	// the specified string.
	NameContains *string

	// If the response to a previous ListModels request was truncated, the response
	// includes a NextToken . To retrieve the next set of models, use the token in the
	// next request.
	NextToken *string

	// Sorts the list of results. The default is CreationTime .
	SortBy types.ModelSortKey

	// The sort order for results. The default is Descending .
	SortOrder types.OrderKey

	noSmithyDocumentSerde
}

type ListModelsOutput struct {

	// An array of ModelSummary objects, each of which lists a model.
	//
	// This member is required.
	Models []types.ModelSummary

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of models, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModels(options.Region), middleware.Before); err != nil {
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
		operation: "ListModels",
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

// ListModelsAPIClient is a client that implements the ListModels operation.
type ListModelsAPIClient interface {
	ListModels(context.Context, *ListModelsInput, ...func(*Options)) (*ListModelsOutput, error)
}

var _ ListModelsAPIClient = (*Client)(nil)

// ListModelsPaginatorOptions is the paginator options for ListModels
type ListModelsPaginatorOptions struct {
	// The maximum number of models to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelsPaginator is a paginator for ListModels
type ListModelsPaginator struct {
	options   ListModelsPaginatorOptions
	client    ListModelsAPIClient
	params    *ListModelsInput
	nextToken *string
	firstPage bool
}

// NewListModelsPaginator returns a new ListModelsPaginator
func NewListModelsPaginator(client ListModelsAPIClient, params *ListModelsInput, optFns ...func(*ListModelsPaginatorOptions)) *ListModelsPaginator {
	if params == nil {
		params = &ListModelsInput{}
	}

	options := ListModelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModels page.
func (p *ListModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelsOutput, error) {
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

	result, err := p.client.ListModels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModels",
	}
}
