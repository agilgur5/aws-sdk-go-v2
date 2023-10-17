// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of source repositories in a project.
func (c *Client) ListSourceRepositories(ctx context.Context, params *ListSourceRepositoriesInput, optFns ...func(*Options)) (*ListSourceRepositoriesOutput, error) {
	if params == nil {
		params = &ListSourceRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSourceRepositories", params, optFns, c.addOperationListSourceRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSourceRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSourceRepositoriesInput struct {

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	MaxResults *int32

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSourceRepositoriesOutput struct {

	// Information about the source repositories.
	Items []types.ListSourceRepositoriesItem

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSourceRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSourceRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSourceRepositories{}, middleware.After)
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
	if err = addBearerAuthSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListSourceRepositoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSourceRepositories(options.Region), middleware.Before); err != nil {
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

// ListSourceRepositoriesAPIClient is a client that implements the
// ListSourceRepositories operation.
type ListSourceRepositoriesAPIClient interface {
	ListSourceRepositories(context.Context, *ListSourceRepositoriesInput, ...func(*Options)) (*ListSourceRepositoriesOutput, error)
}

var _ ListSourceRepositoriesAPIClient = (*Client)(nil)

// ListSourceRepositoriesPaginatorOptions is the paginator options for
// ListSourceRepositories
type ListSourceRepositoriesPaginatorOptions struct {
	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSourceRepositoriesPaginator is a paginator for ListSourceRepositories
type ListSourceRepositoriesPaginator struct {
	options   ListSourceRepositoriesPaginatorOptions
	client    ListSourceRepositoriesAPIClient
	params    *ListSourceRepositoriesInput
	nextToken *string
	firstPage bool
}

// NewListSourceRepositoriesPaginator returns a new ListSourceRepositoriesPaginator
func NewListSourceRepositoriesPaginator(client ListSourceRepositoriesAPIClient, params *ListSourceRepositoriesInput, optFns ...func(*ListSourceRepositoriesPaginatorOptions)) *ListSourceRepositoriesPaginator {
	if params == nil {
		params = &ListSourceRepositoriesInput{}
	}

	options := ListSourceRepositoriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSourceRepositoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSourceRepositoriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSourceRepositories page.
func (p *ListSourceRepositoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSourceRepositoriesOutput, error) {
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

	result, err := p.client.ListSourceRepositories(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSourceRepositories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSourceRepositories",
	}
}
