// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of code signing configurations (https://docs.aws.amazon.com/lambda/latest/dg/configuring-codesigning.html)
// . A request returns up to 10,000 configurations per call. You can use the
// MaxItems parameter to return fewer configurations per call.
func (c *Client) ListCodeSigningConfigs(ctx context.Context, params *ListCodeSigningConfigsInput, optFns ...func(*Options)) (*ListCodeSigningConfigsOutput, error) {
	if params == nil {
		params = &ListCodeSigningConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCodeSigningConfigs", params, optFns, c.addOperationListCodeSigningConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCodeSigningConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCodeSigningConfigsInput struct {

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// Maximum number of items to return.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListCodeSigningConfigsOutput struct {

	// The code signing configurations
	CodeSigningConfigs []types.CodeSigningConfig

	// The pagination token that's included if more results are available.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCodeSigningConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCodeSigningConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCodeSigningConfigs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCodeSigningConfigs(options.Region), middleware.Before); err != nil {
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

// ListCodeSigningConfigsAPIClient is a client that implements the
// ListCodeSigningConfigs operation.
type ListCodeSigningConfigsAPIClient interface {
	ListCodeSigningConfigs(context.Context, *ListCodeSigningConfigsInput, ...func(*Options)) (*ListCodeSigningConfigsOutput, error)
}

var _ ListCodeSigningConfigsAPIClient = (*Client)(nil)

// ListCodeSigningConfigsPaginatorOptions is the paginator options for
// ListCodeSigningConfigs
type ListCodeSigningConfigsPaginatorOptions struct {
	// Maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCodeSigningConfigsPaginator is a paginator for ListCodeSigningConfigs
type ListCodeSigningConfigsPaginator struct {
	options   ListCodeSigningConfigsPaginatorOptions
	client    ListCodeSigningConfigsAPIClient
	params    *ListCodeSigningConfigsInput
	nextToken *string
	firstPage bool
}

// NewListCodeSigningConfigsPaginator returns a new ListCodeSigningConfigsPaginator
func NewListCodeSigningConfigsPaginator(client ListCodeSigningConfigsAPIClient, params *ListCodeSigningConfigsInput, optFns ...func(*ListCodeSigningConfigsPaginatorOptions)) *ListCodeSigningConfigsPaginator {
	if params == nil {
		params = &ListCodeSigningConfigsInput{}
	}

	options := ListCodeSigningConfigsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCodeSigningConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCodeSigningConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCodeSigningConfigs page.
func (p *ListCodeSigningConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCodeSigningConfigsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListCodeSigningConfigs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListCodeSigningConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListCodeSigningConfigs",
	}
}
