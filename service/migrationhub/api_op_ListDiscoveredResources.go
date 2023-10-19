// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists discovered resources associated with the given MigrationTask .
func (c *Client) ListDiscoveredResources(ctx context.Context, params *ListDiscoveredResourcesInput, optFns ...func(*Options)) (*ListDiscoveredResourcesOutput, error) {
	if params == nil {
		params = &ListDiscoveredResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDiscoveredResources", params, optFns, c.addOperationListDiscoveredResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDiscoveredResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDiscoveredResourcesInput struct {

	// The name of the MigrationTask. Do not store personal data in this field.
	//
	// This member is required.
	MigrationTaskName *string

	// The name of the ProgressUpdateStream.
	//
	// This member is required.
	ProgressUpdateStream *string

	// The maximum number of results returned per page.
	MaxResults *int32

	// If a NextToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in NextToken .
	NextToken *string

	noSmithyDocumentSerde
}

type ListDiscoveredResourcesOutput struct {

	// Returned list of discovered resources associated with the given MigrationTask.
	DiscoveredResourceList []types.DiscoveredResource

	// If there are more discovered resources than the max result, return the next
	// token to be passed to the next call as a bookmark of where to start from.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDiscoveredResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDiscoveredResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDiscoveredResources{}, middleware.After)
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
	if err = addOpListDiscoveredResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDiscoveredResources(options.Region), middleware.Before); err != nil {
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
		operation: "ListDiscoveredResources",
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

// ListDiscoveredResourcesAPIClient is a client that implements the
// ListDiscoveredResources operation.
type ListDiscoveredResourcesAPIClient interface {
	ListDiscoveredResources(context.Context, *ListDiscoveredResourcesInput, ...func(*Options)) (*ListDiscoveredResourcesOutput, error)
}

var _ ListDiscoveredResourcesAPIClient = (*Client)(nil)

// ListDiscoveredResourcesPaginatorOptions is the paginator options for
// ListDiscoveredResources
type ListDiscoveredResourcesPaginatorOptions struct {
	// The maximum number of results returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDiscoveredResourcesPaginator is a paginator for ListDiscoveredResources
type ListDiscoveredResourcesPaginator struct {
	options   ListDiscoveredResourcesPaginatorOptions
	client    ListDiscoveredResourcesAPIClient
	params    *ListDiscoveredResourcesInput
	nextToken *string
	firstPage bool
}

// NewListDiscoveredResourcesPaginator returns a new
// ListDiscoveredResourcesPaginator
func NewListDiscoveredResourcesPaginator(client ListDiscoveredResourcesAPIClient, params *ListDiscoveredResourcesInput, optFns ...func(*ListDiscoveredResourcesPaginatorOptions)) *ListDiscoveredResourcesPaginator {
	if params == nil {
		params = &ListDiscoveredResourcesInput{}
	}

	options := ListDiscoveredResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDiscoveredResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDiscoveredResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDiscoveredResources page.
func (p *ListDiscoveredResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDiscoveredResourcesOutput, error) {
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

	result, err := p.client.ListDiscoveredResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDiscoveredResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "ListDiscoveredResources",
	}
}
