// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of summaries for the entity recognizers that you have created.
func (c *Client) ListEntityRecognizerSummaries(ctx context.Context, params *ListEntityRecognizerSummariesInput, optFns ...func(*Options)) (*ListEntityRecognizerSummariesOutput, error) {
	if params == nil {
		params = &ListEntityRecognizerSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntityRecognizerSummaries", params, optFns, c.addOperationListEntityRecognizerSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntityRecognizerSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntityRecognizerSummariesInput struct {

	// The maximum number of results to return on each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEntityRecognizerSummariesOutput struct {

	// The list entity recognizer summaries.
	EntityRecognizerSummariesList []types.EntityRecognizerSummary

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEntityRecognizerSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEntityRecognizerSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntityRecognizerSummaries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntityRecognizerSummaries(options.Region), middleware.Before); err != nil {
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
		operation: "ListEntityRecognizerSummaries",
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

// ListEntityRecognizerSummariesAPIClient is a client that implements the
// ListEntityRecognizerSummaries operation.
type ListEntityRecognizerSummariesAPIClient interface {
	ListEntityRecognizerSummaries(context.Context, *ListEntityRecognizerSummariesInput, ...func(*Options)) (*ListEntityRecognizerSummariesOutput, error)
}

var _ ListEntityRecognizerSummariesAPIClient = (*Client)(nil)

// ListEntityRecognizerSummariesPaginatorOptions is the paginator options for
// ListEntityRecognizerSummaries
type ListEntityRecognizerSummariesPaginatorOptions struct {
	// The maximum number of results to return on each page. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEntityRecognizerSummariesPaginator is a paginator for
// ListEntityRecognizerSummaries
type ListEntityRecognizerSummariesPaginator struct {
	options   ListEntityRecognizerSummariesPaginatorOptions
	client    ListEntityRecognizerSummariesAPIClient
	params    *ListEntityRecognizerSummariesInput
	nextToken *string
	firstPage bool
}

// NewListEntityRecognizerSummariesPaginator returns a new
// ListEntityRecognizerSummariesPaginator
func NewListEntityRecognizerSummariesPaginator(client ListEntityRecognizerSummariesAPIClient, params *ListEntityRecognizerSummariesInput, optFns ...func(*ListEntityRecognizerSummariesPaginatorOptions)) *ListEntityRecognizerSummariesPaginator {
	if params == nil {
		params = &ListEntityRecognizerSummariesInput{}
	}

	options := ListEntityRecognizerSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEntityRecognizerSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEntityRecognizerSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEntityRecognizerSummaries page.
func (p *ListEntityRecognizerSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEntityRecognizerSummariesOutput, error) {
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

	result, err := p.client.ListEntityRecognizerSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEntityRecognizerSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListEntityRecognizerSummaries",
	}
}
