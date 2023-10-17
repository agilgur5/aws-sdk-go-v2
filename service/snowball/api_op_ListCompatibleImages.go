// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action returns a list of the different Amazon EC2-compatible Amazon
// Machine Images (AMIs) that are owned by your Amazon Web Services accountthat
// would be supported for use on a Snow device. Currently, supported AMIs are based
// on the Amazon Linux-2, Ubuntu 20.04 LTS - Focal, or Ubuntu 22.04 LTS - Jammy
// images, available on the Amazon Web Services Marketplace. Ubuntu 16.04 LTS -
// Xenial (HVM) images are no longer supported in the Market, but still supported
// for use on devices through Amazon EC2 VM Import/Export and running locally in
// AMIs.
func (c *Client) ListCompatibleImages(ctx context.Context, params *ListCompatibleImagesInput, optFns ...func(*Options)) (*ListCompatibleImagesOutput, error) {
	if params == nil {
		params = &ListCompatibleImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCompatibleImages", params, optFns, c.addOperationListCompatibleImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCompatibleImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCompatibleImagesInput struct {

	// The maximum number of results for the list of compatible images. Currently, a
	// Snowball Edge device can store 10 AMIs.
	MaxResults *int32

	// HTTP requests are stateless. To identify what object comes "next" in the list
	// of compatible images, you can specify a value for NextToken as the starting
	// point for your list of returned images.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCompatibleImagesOutput struct {

	// A JSON-formatted object that describes a compatible AMI, including the ID and
	// name for a Snow device AMI.
	CompatibleImages []types.CompatibleImage

	// Because HTTP requests are stateless, this is the starting point for your next
	// list of returned images.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCompatibleImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCompatibleImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCompatibleImages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCompatibleImages(options.Region), middleware.Before); err != nil {
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

// ListCompatibleImagesAPIClient is a client that implements the
// ListCompatibleImages operation.
type ListCompatibleImagesAPIClient interface {
	ListCompatibleImages(context.Context, *ListCompatibleImagesInput, ...func(*Options)) (*ListCompatibleImagesOutput, error)
}

var _ ListCompatibleImagesAPIClient = (*Client)(nil)

// ListCompatibleImagesPaginatorOptions is the paginator options for
// ListCompatibleImages
type ListCompatibleImagesPaginatorOptions struct {
	// The maximum number of results for the list of compatible images. Currently, a
	// Snowball Edge device can store 10 AMIs.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCompatibleImagesPaginator is a paginator for ListCompatibleImages
type ListCompatibleImagesPaginator struct {
	options   ListCompatibleImagesPaginatorOptions
	client    ListCompatibleImagesAPIClient
	params    *ListCompatibleImagesInput
	nextToken *string
	firstPage bool
}

// NewListCompatibleImagesPaginator returns a new ListCompatibleImagesPaginator
func NewListCompatibleImagesPaginator(client ListCompatibleImagesAPIClient, params *ListCompatibleImagesInput, optFns ...func(*ListCompatibleImagesPaginatorOptions)) *ListCompatibleImagesPaginator {
	if params == nil {
		params = &ListCompatibleImagesInput{}
	}

	options := ListCompatibleImagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCompatibleImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCompatibleImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCompatibleImages page.
func (p *ListCompatibleImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCompatibleImagesOutput, error) {
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

	result, err := p.client.ListCompatibleImages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCompatibleImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "ListCompatibleImages",
	}
}
