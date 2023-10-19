// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of transcription jobs that match the specified criteria. If no
// criteria are specified, all transcription jobs are returned. To get detailed
// information about a specific transcription job, use the operation.
func (c *Client) ListTranscriptionJobs(ctx context.Context, params *ListTranscriptionJobsInput, optFns ...func(*Options)) (*ListTranscriptionJobsOutput, error) {
	if params == nil {
		params = &ListTranscriptionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTranscriptionJobs", params, optFns, c.addOperationListTranscriptionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTranscriptionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTranscriptionJobsInput struct {

	// Returns only the transcription jobs that contain the specified string. The
	// search is not case sensitive.
	JobNameContains *string

	// The maximum number of transcription jobs to return in each page of results. If
	// there are fewer results than the value that you specify, only the actual results
	// are returned. If you don't specify a value, a default of 5 is used.
	MaxResults *int32

	// If your ListTranscriptionJobs request returns more results than can be
	// displayed, NextToken is displayed in the response with an associated string. To
	// get the next page of results, copy this string and repeat your request,
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	// Returns only transcription jobs with the specified status. Jobs are ordered by
	// creation date, with the newest job first. If you don't include Status , all
	// transcription jobs are returned.
	Status types.TranscriptionJobStatus

	noSmithyDocumentSerde
}

type ListTranscriptionJobsOutput struct {

	// If NextToken is present in your response, it indicates that not all results are
	// displayed. To view the next set of results, copy the string associated with the
	// NextToken parameter in your results output, then run your request again
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	// Lists all transcription jobs that have the status specified in your request.
	// Jobs are ordered by creation date, with the newest job first.
	Status types.TranscriptionJobStatus

	// Provides a summary of information about each result.
	TranscriptionJobSummaries []types.TranscriptionJobSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTranscriptionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTranscriptionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTranscriptionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTranscriptionJobs(options.Region), middleware.Before); err != nil {
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
		operation: "ListTranscriptionJobs",
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

// ListTranscriptionJobsAPIClient is a client that implements the
// ListTranscriptionJobs operation.
type ListTranscriptionJobsAPIClient interface {
	ListTranscriptionJobs(context.Context, *ListTranscriptionJobsInput, ...func(*Options)) (*ListTranscriptionJobsOutput, error)
}

var _ ListTranscriptionJobsAPIClient = (*Client)(nil)

// ListTranscriptionJobsPaginatorOptions is the paginator options for
// ListTranscriptionJobs
type ListTranscriptionJobsPaginatorOptions struct {
	// The maximum number of transcription jobs to return in each page of results. If
	// there are fewer results than the value that you specify, only the actual results
	// are returned. If you don't specify a value, a default of 5 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTranscriptionJobsPaginator is a paginator for ListTranscriptionJobs
type ListTranscriptionJobsPaginator struct {
	options   ListTranscriptionJobsPaginatorOptions
	client    ListTranscriptionJobsAPIClient
	params    *ListTranscriptionJobsInput
	nextToken *string
	firstPage bool
}

// NewListTranscriptionJobsPaginator returns a new ListTranscriptionJobsPaginator
func NewListTranscriptionJobsPaginator(client ListTranscriptionJobsAPIClient, params *ListTranscriptionJobsInput, optFns ...func(*ListTranscriptionJobsPaginatorOptions)) *ListTranscriptionJobsPaginator {
	if params == nil {
		params = &ListTranscriptionJobsInput{}
	}

	options := ListTranscriptionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTranscriptionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTranscriptionJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTranscriptionJobs page.
func (p *ListTranscriptionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTranscriptionJobsOutput, error) {
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

	result, err := p.client.ListTranscriptionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTranscriptionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListTranscriptionJobs",
	}
}
