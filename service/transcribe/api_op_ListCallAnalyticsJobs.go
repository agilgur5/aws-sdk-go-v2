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

// Provides a list of Call Analytics jobs that match the specified criteria. If no
// criteria are specified, all Call Analytics jobs are returned. To get detailed
// information about a specific Call Analytics job, use the operation.
func (c *Client) ListCallAnalyticsJobs(ctx context.Context, params *ListCallAnalyticsJobsInput, optFns ...func(*Options)) (*ListCallAnalyticsJobsOutput, error) {
	if params == nil {
		params = &ListCallAnalyticsJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCallAnalyticsJobs", params, optFns, c.addOperationListCallAnalyticsJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCallAnalyticsJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCallAnalyticsJobsInput struct {

	// Returns only the Call Analytics jobs that contain the specified string. The
	// search is not case sensitive.
	JobNameContains *string

	// The maximum number of Call Analytics jobs to return in each page of results. If
	// there are fewer results than the value that you specify, only the actual results
	// are returned. If you don't specify a value, a default of 5 is used.
	MaxResults *int32

	// If your ListCallAnalyticsJobs request returns more results than can be
	// displayed, NextToken is displayed in the response with an associated string. To
	// get the next page of results, copy this string and repeat your request,
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	// Returns only Call Analytics jobs with the specified status. Jobs are ordered by
	// creation date, with the newest job first. If you don't include Status , all Call
	// Analytics jobs are returned.
	Status types.CallAnalyticsJobStatus

	noSmithyDocumentSerde
}

type ListCallAnalyticsJobsOutput struct {

	// Provides a summary of information about each result.
	CallAnalyticsJobSummaries []types.CallAnalyticsJobSummary

	// If NextToken is present in your response, it indicates that not all results are
	// displayed. To view the next set of results, copy the string associated with the
	// NextToken parameter in your results output, then run your request again
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	// Lists all Call Analytics jobs that have the status specified in your request.
	// Jobs are ordered by creation date, with the newest job first.
	Status types.CallAnalyticsJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCallAnalyticsJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCallAnalyticsJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCallAnalyticsJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCallAnalyticsJobs(options.Region), middleware.Before); err != nil {
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
		operation: "ListCallAnalyticsJobs",
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

// ListCallAnalyticsJobsAPIClient is a client that implements the
// ListCallAnalyticsJobs operation.
type ListCallAnalyticsJobsAPIClient interface {
	ListCallAnalyticsJobs(context.Context, *ListCallAnalyticsJobsInput, ...func(*Options)) (*ListCallAnalyticsJobsOutput, error)
}

var _ ListCallAnalyticsJobsAPIClient = (*Client)(nil)

// ListCallAnalyticsJobsPaginatorOptions is the paginator options for
// ListCallAnalyticsJobs
type ListCallAnalyticsJobsPaginatorOptions struct {
	// The maximum number of Call Analytics jobs to return in each page of results. If
	// there are fewer results than the value that you specify, only the actual results
	// are returned. If you don't specify a value, a default of 5 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCallAnalyticsJobsPaginator is a paginator for ListCallAnalyticsJobs
type ListCallAnalyticsJobsPaginator struct {
	options   ListCallAnalyticsJobsPaginatorOptions
	client    ListCallAnalyticsJobsAPIClient
	params    *ListCallAnalyticsJobsInput
	nextToken *string
	firstPage bool
}

// NewListCallAnalyticsJobsPaginator returns a new ListCallAnalyticsJobsPaginator
func NewListCallAnalyticsJobsPaginator(client ListCallAnalyticsJobsAPIClient, params *ListCallAnalyticsJobsInput, optFns ...func(*ListCallAnalyticsJobsPaginatorOptions)) *ListCallAnalyticsJobsPaginator {
	if params == nil {
		params = &ListCallAnalyticsJobsInput{}
	}

	options := ListCallAnalyticsJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCallAnalyticsJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCallAnalyticsJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCallAnalyticsJobs page.
func (p *ListCallAnalyticsJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCallAnalyticsJobsOutput, error) {
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

	result, err := p.client.ListCallAnalyticsJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCallAnalyticsJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListCallAnalyticsJobs",
	}
}
