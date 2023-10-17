// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the export jobs.
func (c *Client) ListExportJobs(ctx context.Context, params *ListExportJobsInput, optFns ...func(*Options)) (*ListExportJobsOutput, error) {
	if params == nil {
		params = &ListExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExportJobs", params, optFns, c.addOperationListExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list all export jobs with filters.
type ListExportJobsInput struct {

	// A value used to list export jobs that have a certain ExportSourceType .
	ExportSourceType types.ExportSourceType

	// A value used to list export jobs that have a certain JobStatus .
	JobStatus types.JobStatus

	// The pagination token returned from a previous call to ListExportJobs to
	// indicate the position in the list of export jobs.
	NextToken *string

	// Maximum number of export jobs to return at once. Use this parameter to paginate
	// results. If additional export jobs exist beyond the specified limit, the
	// NextToken element is sent in the response. Use the NextToken value in
	// subsequent calls to ListExportJobs to retrieve additional export jobs.
	PageSize *int32

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type ListExportJobsOutput struct {

	// A list of the export job summaries.
	ExportJobs []types.ExportJobSummary

	// A string token indicating that there might be additional export jobs available
	// to be listed. Use this token to a subsequent call to ListExportJobs with the
	// same parameters to retrieve the next page of export jobs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExportJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExportJobs(options.Region), middleware.Before); err != nil {
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

// ListExportJobsAPIClient is a client that implements the ListExportJobs
// operation.
type ListExportJobsAPIClient interface {
	ListExportJobs(context.Context, *ListExportJobsInput, ...func(*Options)) (*ListExportJobsOutput, error)
}

var _ ListExportJobsAPIClient = (*Client)(nil)

// ListExportJobsPaginatorOptions is the paginator options for ListExportJobs
type ListExportJobsPaginatorOptions struct {
	// Maximum number of export jobs to return at once. Use this parameter to paginate
	// results. If additional export jobs exist beyond the specified limit, the
	// NextToken element is sent in the response. Use the NextToken value in
	// subsequent calls to ListExportJobs to retrieve additional export jobs.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExportJobsPaginator is a paginator for ListExportJobs
type ListExportJobsPaginator struct {
	options   ListExportJobsPaginatorOptions
	client    ListExportJobsAPIClient
	params    *ListExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListExportJobsPaginator returns a new ListExportJobsPaginator
func NewListExportJobsPaginator(client ListExportJobsAPIClient, params *ListExportJobsInput, optFns ...func(*ListExportJobsPaginatorOptions)) *ListExportJobsPaginator {
	if params == nil {
		params = &ListExportJobsInput{}
	}

	options := ListExportJobsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExportJobs page.
func (p *ListExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExportJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListExportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListExportJobs",
	}
}
