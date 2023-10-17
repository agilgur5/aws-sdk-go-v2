// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of read set import jobs.
func (c *Client) ListReadSetImportJobs(ctx context.Context, params *ListReadSetImportJobsInput, optFns ...func(*Options)) (*ListReadSetImportJobsOutput, error) {
	if params == nil {
		params = &ListReadSetImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReadSetImportJobs", params, optFns, c.addOperationListReadSetImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReadSetImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReadSetImportJobsInput struct {

	// The jobs' sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// A filter to apply to the list.
	Filter *types.ImportReadSetFilter

	// The maximum number of jobs to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReadSetImportJobsOutput struct {

	// A list of jobs.
	ImportJobs []types.ImportReadSetJobItem

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReadSetImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReadSetImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReadSetImportJobs{}, middleware.After)
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
	if err = addEndpointPrefix_opListReadSetImportJobsMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpListReadSetImportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReadSetImportJobs(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListReadSetImportJobsMiddleware struct {
}

func (*endpointPrefix_opListReadSetImportJobsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListReadSetImportJobsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListReadSetImportJobsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListReadSetImportJobsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListReadSetImportJobsAPIClient is a client that implements the
// ListReadSetImportJobs operation.
type ListReadSetImportJobsAPIClient interface {
	ListReadSetImportJobs(context.Context, *ListReadSetImportJobsInput, ...func(*Options)) (*ListReadSetImportJobsOutput, error)
}

var _ ListReadSetImportJobsAPIClient = (*Client)(nil)

// ListReadSetImportJobsPaginatorOptions is the paginator options for
// ListReadSetImportJobs
type ListReadSetImportJobsPaginatorOptions struct {
	// The maximum number of jobs to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReadSetImportJobsPaginator is a paginator for ListReadSetImportJobs
type ListReadSetImportJobsPaginator struct {
	options   ListReadSetImportJobsPaginatorOptions
	client    ListReadSetImportJobsAPIClient
	params    *ListReadSetImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListReadSetImportJobsPaginator returns a new ListReadSetImportJobsPaginator
func NewListReadSetImportJobsPaginator(client ListReadSetImportJobsAPIClient, params *ListReadSetImportJobsInput, optFns ...func(*ListReadSetImportJobsPaginatorOptions)) *ListReadSetImportJobsPaginator {
	if params == nil {
		params = &ListReadSetImportJobsInput{}
	}

	options := ListReadSetImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReadSetImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReadSetImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReadSetImportJobs page.
func (p *ListReadSetImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReadSetImportJobsOutput, error) {
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

	result, err := p.client.ListReadSetImportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReadSetImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "ListReadSetImportJobs",
	}
}
