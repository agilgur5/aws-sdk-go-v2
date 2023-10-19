// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of image scan aggregations for your account. You can filter by
// the type of key that Image Builder uses to group results. For example, if you
// want to get a list of findings by severity level for one of your pipelines, you
// might specify your pipeline with the imagePipelineArn filter. If you don't
// specify a filter, Image Builder returns an aggregation for your account. To
// streamline results, you can use the following filters in your request:
//   - accountId
//   - imageBuildVersionArn
//   - imagePipelineArn
//   - vulnerabilityId
func (c *Client) ListImageScanFindingAggregations(ctx context.Context, params *ListImageScanFindingAggregationsInput, optFns ...func(*Options)) (*ListImageScanFindingAggregationsOutput, error) {
	if params == nil {
		params = &ListImageScanFindingAggregationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImageScanFindingAggregations", params, optFns, c.addOperationListImageScanFindingAggregationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImageScanFindingAggregationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImageScanFindingAggregationsInput struct {

	// A filter name and value pair that is used to return a more specific list of
	// results from a list operation. Filters can be used to match a set of resources
	// by specific criteria, such as tags, attributes, or IDs.
	Filter *types.Filter

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListImageScanFindingAggregationsOutput struct {

	// The aggregation type specifies what type of key is used to group the image scan
	// findings. Image Builder returns results based on the request filter. If you
	// didn't specify a filter in the request, the type defaults to accountId .
	// Aggregation types
	//   - accountId
	//   - imageBuildVersionArn
	//   - imagePipelineArn
	//   - vulnerabilityId
	// Each aggregation includes counts by severity level for medium severity and
	// higher level findings, plus a total for all of the findings for each key value.
	AggregationType *string

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service has'ot included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// An array of image scan finding aggregations that match the filter criteria.
	Responses []types.ImageScanFindingAggregation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImageScanFindingAggregationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImageScanFindingAggregations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImageScanFindingAggregations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImageScanFindingAggregations(options.Region), middleware.Before); err != nil {
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
		operation: "ListImageScanFindingAggregations",
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

// ListImageScanFindingAggregationsAPIClient is a client that implements the
// ListImageScanFindingAggregations operation.
type ListImageScanFindingAggregationsAPIClient interface {
	ListImageScanFindingAggregations(context.Context, *ListImageScanFindingAggregationsInput, ...func(*Options)) (*ListImageScanFindingAggregationsOutput, error)
}

var _ ListImageScanFindingAggregationsAPIClient = (*Client)(nil)

// ListImageScanFindingAggregationsPaginatorOptions is the paginator options for
// ListImageScanFindingAggregations
type ListImageScanFindingAggregationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImageScanFindingAggregationsPaginator is a paginator for
// ListImageScanFindingAggregations
type ListImageScanFindingAggregationsPaginator struct {
	options   ListImageScanFindingAggregationsPaginatorOptions
	client    ListImageScanFindingAggregationsAPIClient
	params    *ListImageScanFindingAggregationsInput
	nextToken *string
	firstPage bool
}

// NewListImageScanFindingAggregationsPaginator returns a new
// ListImageScanFindingAggregationsPaginator
func NewListImageScanFindingAggregationsPaginator(client ListImageScanFindingAggregationsAPIClient, params *ListImageScanFindingAggregationsInput, optFns ...func(*ListImageScanFindingAggregationsPaginatorOptions)) *ListImageScanFindingAggregationsPaginator {
	if params == nil {
		params = &ListImageScanFindingAggregationsInput{}
	}

	options := ListImageScanFindingAggregationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImageScanFindingAggregationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImageScanFindingAggregationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImageScanFindingAggregations page.
func (p *ListImageScanFindingAggregationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImageScanFindingAggregationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListImageScanFindingAggregations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImageScanFindingAggregations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListImageScanFindingAggregations",
	}
}
