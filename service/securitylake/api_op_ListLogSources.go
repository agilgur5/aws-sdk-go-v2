// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the log sources in the current Amazon Web Services Region.
func (c *Client) ListLogSources(ctx context.Context, params *ListLogSourcesInput, optFns ...func(*Options)) (*ListLogSourcesOutput, error) {
	if params == nil {
		params = &ListLogSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLogSources", params, optFns, c.addOperationListLogSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLogSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLogSourcesInput struct {

	// The list of Amazon Web Services accounts for which log sources are displayed.
	Accounts []string

	// The maximum number of accounts for which the log sources are displayed.
	MaxResults *int32

	// If nextToken is returned, there are more results available. You can repeat the
	// call using the returned token to retrieve the next page.
	NextToken *string

	// The list of regions for which log sources are displayed.
	Regions []string

	// The list of sources for which log sources are displayed.
	Sources []types.LogSourceResource

	noSmithyDocumentSerde
}

type ListLogSourcesOutput struct {

	// If nextToken is returned, there are more results available. You can repeat the
	// call using the returned token to retrieve the next page.
	NextToken *string

	// The list of log sources in your organization that send data to the data lake.
	Sources []types.LogSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLogSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLogSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLogSources{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLogSources(options.Region), middleware.Before); err != nil {
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
		operation: "ListLogSources",
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

// ListLogSourcesAPIClient is a client that implements the ListLogSources
// operation.
type ListLogSourcesAPIClient interface {
	ListLogSources(context.Context, *ListLogSourcesInput, ...func(*Options)) (*ListLogSourcesOutput, error)
}

var _ ListLogSourcesAPIClient = (*Client)(nil)

// ListLogSourcesPaginatorOptions is the paginator options for ListLogSources
type ListLogSourcesPaginatorOptions struct {
	// The maximum number of accounts for which the log sources are displayed.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLogSourcesPaginator is a paginator for ListLogSources
type ListLogSourcesPaginator struct {
	options   ListLogSourcesPaginatorOptions
	client    ListLogSourcesAPIClient
	params    *ListLogSourcesInput
	nextToken *string
	firstPage bool
}

// NewListLogSourcesPaginator returns a new ListLogSourcesPaginator
func NewListLogSourcesPaginator(client ListLogSourcesAPIClient, params *ListLogSourcesInput, optFns ...func(*ListLogSourcesPaginatorOptions)) *ListLogSourcesPaginator {
	if params == nil {
		params = &ListLogSourcesInput{}
	}

	options := ListLogSourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLogSourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLogSourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLogSources page.
func (p *ListLogSourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLogSourcesOutput, error) {
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

	result, err := p.client.ListLogSources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLogSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "ListLogSources",
	}
}
