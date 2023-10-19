// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the proxy sessions for the specified Amazon Chime SDK Voice Connector.
func (c *Client) ListProxySessions(ctx context.Context, params *ListProxySessionsInput, optFns ...func(*Options)) (*ListProxySessionsOutput, error) {
	if params == nil {
		params = &ListProxySessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProxySessions", params, optFns, c.addOperationListProxySessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProxySessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProxySessionsInput struct {

	// The Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token used to retrieve the next page of results.
	NextToken *string

	// The proxy session status.
	Status types.ProxySessionStatus

	noSmithyDocumentSerde
}

type ListProxySessionsOutput struct {

	// The token used to retrieve the next page of results.
	NextToken *string

	// The proxy sessions' details.
	ProxySessions []types.ProxySession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProxySessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProxySessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProxySessions{}, middleware.After)
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
	if err = addOpListProxySessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProxySessions(options.Region), middleware.Before); err != nil {
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
		operation: "ListProxySessions",
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

// ListProxySessionsAPIClient is a client that implements the ListProxySessions
// operation.
type ListProxySessionsAPIClient interface {
	ListProxySessions(context.Context, *ListProxySessionsInput, ...func(*Options)) (*ListProxySessionsOutput, error)
}

var _ ListProxySessionsAPIClient = (*Client)(nil)

// ListProxySessionsPaginatorOptions is the paginator options for ListProxySessions
type ListProxySessionsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProxySessionsPaginator is a paginator for ListProxySessions
type ListProxySessionsPaginator struct {
	options   ListProxySessionsPaginatorOptions
	client    ListProxySessionsAPIClient
	params    *ListProxySessionsInput
	nextToken *string
	firstPage bool
}

// NewListProxySessionsPaginator returns a new ListProxySessionsPaginator
func NewListProxySessionsPaginator(client ListProxySessionsAPIClient, params *ListProxySessionsInput, optFns ...func(*ListProxySessionsPaginatorOptions)) *ListProxySessionsPaginator {
	if params == nil {
		params = &ListProxySessionsInput{}
	}

	options := ListProxySessionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProxySessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProxySessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProxySessions page.
func (p *ListProxySessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProxySessionsOutput, error) {
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

	result, err := p.client.ListProxySessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProxySessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListProxySessions",
	}
}
