// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resources associated with the specified TagOption.
func (c *Client) ListResourcesForTagOption(ctx context.Context, params *ListResourcesForTagOptionInput, optFns ...func(*Options)) (*ListResourcesForTagOptionOutput, error) {
	if params == nil {
		params = &ListResourcesForTagOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourcesForTagOption", params, optFns, c.addOperationListResourcesForTagOptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesForTagOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesForTagOptionInput struct {

	// The TagOption identifier.
	//
	// This member is required.
	TagOptionId *string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The resource type.
	//   - Portfolio
	//   - Product
	ResourceType *string

	noSmithyDocumentSerde
}

type ListResourcesForTagOptionOutput struct {

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// Information about the resources.
	ResourceDetails []types.ResourceDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourcesForTagOptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourcesForTagOption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourcesForTagOption{}, middleware.After)
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
	if err = addOpListResourcesForTagOptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourcesForTagOption(options.Region), middleware.Before); err != nil {
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

// ListResourcesForTagOptionAPIClient is a client that implements the
// ListResourcesForTagOption operation.
type ListResourcesForTagOptionAPIClient interface {
	ListResourcesForTagOption(context.Context, *ListResourcesForTagOptionInput, ...func(*Options)) (*ListResourcesForTagOptionOutput, error)
}

var _ ListResourcesForTagOptionAPIClient = (*Client)(nil)

// ListResourcesForTagOptionPaginatorOptions is the paginator options for
// ListResourcesForTagOption
type ListResourcesForTagOptionPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourcesForTagOptionPaginator is a paginator for ListResourcesForTagOption
type ListResourcesForTagOptionPaginator struct {
	options   ListResourcesForTagOptionPaginatorOptions
	client    ListResourcesForTagOptionAPIClient
	params    *ListResourcesForTagOptionInput
	nextToken *string
	firstPage bool
}

// NewListResourcesForTagOptionPaginator returns a new
// ListResourcesForTagOptionPaginator
func NewListResourcesForTagOptionPaginator(client ListResourcesForTagOptionAPIClient, params *ListResourcesForTagOptionInput, optFns ...func(*ListResourcesForTagOptionPaginatorOptions)) *ListResourcesForTagOptionPaginator {
	if params == nil {
		params = &ListResourcesForTagOptionInput{}
	}

	options := ListResourcesForTagOptionPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourcesForTagOptionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourcesForTagOptionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourcesForTagOption page.
func (p *ListResourcesForTagOptionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourcesForTagOptionOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListResourcesForTagOption(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListResourcesForTagOption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListResourcesForTagOption",
	}
}
