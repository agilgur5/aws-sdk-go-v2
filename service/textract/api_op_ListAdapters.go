// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all adapters that match the specified filtration criteria.
func (c *Client) ListAdapters(ctx context.Context, params *ListAdaptersInput, optFns ...func(*Options)) (*ListAdaptersOutput, error) {
	if params == nil {
		params = &ListAdaptersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAdapters", params, optFns, c.addOperationListAdaptersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAdaptersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAdaptersInput struct {

	// Specifies the lower bound for the ListAdapters operation. Ensures ListAdapters
	// returns only adapters created after the specified creation time.
	AfterCreationTime *time.Time

	// Specifies the upper bound for the ListAdapters operation. Ensures ListAdapters
	// returns only adapters created before the specified creation time.
	BeforeCreationTime *time.Time

	// The maximum number of results to return when listing adapters.
	MaxResults *int32

	// Identifies the next page of results to return when listing adapters.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAdaptersOutput struct {

	// A list of adapters that matches the filtering criteria specified when calling
	// ListAdapters.
	Adapters []types.AdapterOverview

	// Identifies the next page of results to return when listing adapters.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAdaptersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAdapters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAdapters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAdapters(options.Region), middleware.Before); err != nil {
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
		operation: "ListAdapters",
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

// ListAdaptersAPIClient is a client that implements the ListAdapters operation.
type ListAdaptersAPIClient interface {
	ListAdapters(context.Context, *ListAdaptersInput, ...func(*Options)) (*ListAdaptersOutput, error)
}

var _ ListAdaptersAPIClient = (*Client)(nil)

// ListAdaptersPaginatorOptions is the paginator options for ListAdapters
type ListAdaptersPaginatorOptions struct {
	// The maximum number of results to return when listing adapters.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAdaptersPaginator is a paginator for ListAdapters
type ListAdaptersPaginator struct {
	options   ListAdaptersPaginatorOptions
	client    ListAdaptersAPIClient
	params    *ListAdaptersInput
	nextToken *string
	firstPage bool
}

// NewListAdaptersPaginator returns a new ListAdaptersPaginator
func NewListAdaptersPaginator(client ListAdaptersAPIClient, params *ListAdaptersInput, optFns ...func(*ListAdaptersPaginatorOptions)) *ListAdaptersPaginator {
	if params == nil {
		params = &ListAdaptersInput{}
	}

	options := ListAdaptersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAdaptersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAdaptersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAdapters page.
func (p *ListAdaptersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAdaptersOutput, error) {
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

	result, err := p.client.ListAdapters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAdapters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "ListAdapters",
	}
}
