// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all custom and Amazon Web Services authored AppConfig extensions in the
// account. For more information about extensions, see Working with AppConfig
// extensions (https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html)
// in the AppConfig User Guide.
func (c *Client) ListExtensions(ctx context.Context, params *ListExtensionsInput, optFns ...func(*Options)) (*ListExtensionsOutput, error) {
	if params == nil {
		params = &ListExtensionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExtensions", params, optFns, c.addOperationListExtensionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExtensionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExtensionsInput struct {

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The extension name.
	Name *string

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExtensionsOutput struct {

	// The list of available extensions. The list includes Amazon Web Services
	// authored and user-created extensions.
	Items []types.ExtensionSummary

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExtensionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExtensions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExtensions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExtensions(options.Region), middleware.Before); err != nil {
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

// ListExtensionsAPIClient is a client that implements the ListExtensions
// operation.
type ListExtensionsAPIClient interface {
	ListExtensions(context.Context, *ListExtensionsInput, ...func(*Options)) (*ListExtensionsOutput, error)
}

var _ ListExtensionsAPIClient = (*Client)(nil)

// ListExtensionsPaginatorOptions is the paginator options for ListExtensions
type ListExtensionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExtensionsPaginator is a paginator for ListExtensions
type ListExtensionsPaginator struct {
	options   ListExtensionsPaginatorOptions
	client    ListExtensionsAPIClient
	params    *ListExtensionsInput
	nextToken *string
	firstPage bool
}

// NewListExtensionsPaginator returns a new ListExtensionsPaginator
func NewListExtensionsPaginator(client ListExtensionsAPIClient, params *ListExtensionsInput, optFns ...func(*ListExtensionsPaginatorOptions)) *ListExtensionsPaginator {
	if params == nil {
		params = &ListExtensionsInput{}
	}

	options := ListExtensionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExtensionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExtensionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExtensions page.
func (p *ListExtensionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExtensionsOutput, error) {
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

	result, err := p.client.ListExtensions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExtensions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "ListExtensions",
	}
}
