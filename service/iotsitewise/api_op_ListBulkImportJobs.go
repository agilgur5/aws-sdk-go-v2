// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of bulk import job requests. For more information,
// see List bulk import jobs (CLI) (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/ListBulkImportJobs.html)
// in the IoT SiteWise User Guide.
func (c *Client) ListBulkImportJobs(ctx context.Context, params *ListBulkImportJobsInput, optFns ...func(*Options)) (*ListBulkImportJobsOutput, error) {
	if params == nil {
		params = &ListBulkImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBulkImportJobs", params, optFns, c.addOperationListBulkImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBulkImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBulkImportJobsInput struct {

	// You can use a filter to select the bulk import jobs that you want to retrieve.
	Filter types.ListBulkImportJobsFilter

	// The maximum number of results to return for each paginated request.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBulkImportJobsOutput struct {

	// One or more job summaries to list.
	//
	// This member is required.
	JobSummaries []types.JobSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBulkImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBulkImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBulkImportJobs{}, middleware.After)
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
	if err = addEndpointPrefix_opListBulkImportJobsMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBulkImportJobs(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListBulkImportJobsMiddleware struct {
}

func (*endpointPrefix_opListBulkImportJobsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListBulkImportJobsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListBulkImportJobsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListBulkImportJobsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListBulkImportJobsAPIClient is a client that implements the ListBulkImportJobs
// operation.
type ListBulkImportJobsAPIClient interface {
	ListBulkImportJobs(context.Context, *ListBulkImportJobsInput, ...func(*Options)) (*ListBulkImportJobsOutput, error)
}

var _ ListBulkImportJobsAPIClient = (*Client)(nil)

// ListBulkImportJobsPaginatorOptions is the paginator options for
// ListBulkImportJobs
type ListBulkImportJobsPaginatorOptions struct {
	// The maximum number of results to return for each paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBulkImportJobsPaginator is a paginator for ListBulkImportJobs
type ListBulkImportJobsPaginator struct {
	options   ListBulkImportJobsPaginatorOptions
	client    ListBulkImportJobsAPIClient
	params    *ListBulkImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListBulkImportJobsPaginator returns a new ListBulkImportJobsPaginator
func NewListBulkImportJobsPaginator(client ListBulkImportJobsAPIClient, params *ListBulkImportJobsInput, optFns ...func(*ListBulkImportJobsPaginatorOptions)) *ListBulkImportJobsPaginator {
	if params == nil {
		params = &ListBulkImportJobsInput{}
	}

	options := ListBulkImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBulkImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBulkImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBulkImportJobs page.
func (p *ListBulkImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBulkImportJobsOutput, error) {
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

	result, err := p.client.ListBulkImportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBulkImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListBulkImportJobs",
	}
}
