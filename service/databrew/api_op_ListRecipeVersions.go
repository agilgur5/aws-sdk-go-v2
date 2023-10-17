// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the versions of a particular DataBrew recipe, except for LATEST_WORKING .
func (c *Client) ListRecipeVersions(ctx context.Context, params *ListRecipeVersionsInput, optFns ...func(*Options)) (*ListRecipeVersionsOutput, error) {
	if params == nil {
		params = &ListRecipeVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecipeVersions", params, optFns, c.addOperationListRecipeVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecipeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecipeVersionsInput struct {

	// The name of the recipe for which to return version information.
	//
	// This member is required.
	Name *string

	// The maximum number of results to return in this request.
	MaxResults *int32

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRecipeVersionsOutput struct {

	// A list of versions for the specified recipe.
	//
	// This member is required.
	Recipes []types.Recipe

	// A token that you can use in a subsequent call to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecipeVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecipeVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecipeVersions{}, middleware.After)
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
	if err = addOpListRecipeVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecipeVersions(options.Region), middleware.Before); err != nil {
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

// ListRecipeVersionsAPIClient is a client that implements the ListRecipeVersions
// operation.
type ListRecipeVersionsAPIClient interface {
	ListRecipeVersions(context.Context, *ListRecipeVersionsInput, ...func(*Options)) (*ListRecipeVersionsOutput, error)
}

var _ ListRecipeVersionsAPIClient = (*Client)(nil)

// ListRecipeVersionsPaginatorOptions is the paginator options for
// ListRecipeVersions
type ListRecipeVersionsPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecipeVersionsPaginator is a paginator for ListRecipeVersions
type ListRecipeVersionsPaginator struct {
	options   ListRecipeVersionsPaginatorOptions
	client    ListRecipeVersionsAPIClient
	params    *ListRecipeVersionsInput
	nextToken *string
	firstPage bool
}

// NewListRecipeVersionsPaginator returns a new ListRecipeVersionsPaginator
func NewListRecipeVersionsPaginator(client ListRecipeVersionsAPIClient, params *ListRecipeVersionsInput, optFns ...func(*ListRecipeVersionsPaginatorOptions)) *ListRecipeVersionsPaginator {
	if params == nil {
		params = &ListRecipeVersionsInput{}
	}

	options := ListRecipeVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecipeVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecipeVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecipeVersions page.
func (p *ListRecipeVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecipeVersionsOutput, error) {
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

	result, err := p.client.ListRecipeVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecipeVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "ListRecipeVersions",
	}
}
