// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of vehicle models (model manifests). This API operation uses
// pagination. Specify the nextToken parameter in the request to return more
// results.
func (c *Client) ListModelManifests(ctx context.Context, params *ListModelManifestsInput, optFns ...func(*Options)) (*ListModelManifestsOutput, error) {
	if params == nil {
		params = &ListModelManifestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelManifests", params, optFns, c.addOperationListModelManifestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelManifestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelManifestsInput struct {

	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	// A pagination token for the next set of results. If the results of a search are
	// large, only a portion of the results are returned, and a nextToken pagination
	// token is returned in the response. To retrieve the next set of results, reissue
	// the search request and include the returned token. When all results have been
	// returned, the response does not contain a pagination token value.
	NextToken *string

	// The ARN of a signal catalog. If you specify a signal catalog, only the vehicle
	// models associated with it are returned.
	SignalCatalogArn *string

	noSmithyDocumentSerde
}

type ListModelManifestsOutput struct {

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// A list of information about vehicle models.
	Summaries []types.ModelManifestSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelManifestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListModelManifests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListModelManifests{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelManifests(options.Region), middleware.Before); err != nil {
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

// ListModelManifestsAPIClient is a client that implements the ListModelManifests
// operation.
type ListModelManifestsAPIClient interface {
	ListModelManifests(context.Context, *ListModelManifestsInput, ...func(*Options)) (*ListModelManifestsOutput, error)
}

var _ ListModelManifestsAPIClient = (*Client)(nil)

// ListModelManifestsPaginatorOptions is the paginator options for
// ListModelManifests
type ListModelManifestsPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelManifestsPaginator is a paginator for ListModelManifests
type ListModelManifestsPaginator struct {
	options   ListModelManifestsPaginatorOptions
	client    ListModelManifestsAPIClient
	params    *ListModelManifestsInput
	nextToken *string
	firstPage bool
}

// NewListModelManifestsPaginator returns a new ListModelManifestsPaginator
func NewListModelManifestsPaginator(client ListModelManifestsAPIClient, params *ListModelManifestsInput, optFns ...func(*ListModelManifestsPaginatorOptions)) *ListModelManifestsPaginator {
	if params == nil {
		params = &ListModelManifestsInput{}
	}

	options := ListModelManifestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelManifestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelManifestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelManifests page.
func (p *ListModelManifestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelManifestsOutput, error) {
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

	result, err := p.client.ListModelManifests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModelManifests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "ListModelManifests",
	}
}
