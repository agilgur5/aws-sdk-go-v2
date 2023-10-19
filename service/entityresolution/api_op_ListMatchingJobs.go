// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all jobs for a given workflow.
func (c *Client) ListMatchingJobs(ctx context.Context, params *ListMatchingJobsInput, optFns ...func(*Options)) (*ListMatchingJobsOutput, error) {
	if params == nil {
		params = &ListMatchingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMatchingJobs", params, optFns, c.addOperationListMatchingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMatchingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMatchingJobsInput struct {

	// The name of the workflow to be retrieved.
	//
	// This member is required.
	WorkflowName *string

	// The maximum number of objects returned per page.
	MaxResults *int32

	// The pagination token from the previous API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMatchingJobsOutput struct {

	// A list of JobSummary objects, each of which contain the ID, status, start time,
	// and end time of a job.
	Jobs []types.JobSummary

	// The pagination token from the previous API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMatchingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMatchingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMatchingJobs{}, middleware.After)
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
	if err = addOpListMatchingJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMatchingJobs(options.Region), middleware.Before); err != nil {
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
		operation: "ListMatchingJobs",
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

// ListMatchingJobsAPIClient is a client that implements the ListMatchingJobs
// operation.
type ListMatchingJobsAPIClient interface {
	ListMatchingJobs(context.Context, *ListMatchingJobsInput, ...func(*Options)) (*ListMatchingJobsOutput, error)
}

var _ ListMatchingJobsAPIClient = (*Client)(nil)

// ListMatchingJobsPaginatorOptions is the paginator options for ListMatchingJobs
type ListMatchingJobsPaginatorOptions struct {
	// The maximum number of objects returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMatchingJobsPaginator is a paginator for ListMatchingJobs
type ListMatchingJobsPaginator struct {
	options   ListMatchingJobsPaginatorOptions
	client    ListMatchingJobsAPIClient
	params    *ListMatchingJobsInput
	nextToken *string
	firstPage bool
}

// NewListMatchingJobsPaginator returns a new ListMatchingJobsPaginator
func NewListMatchingJobsPaginator(client ListMatchingJobsAPIClient, params *ListMatchingJobsInput, optFns ...func(*ListMatchingJobsPaginatorOptions)) *ListMatchingJobsPaginator {
	if params == nil {
		params = &ListMatchingJobsInput{}
	}

	options := ListMatchingJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMatchingJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMatchingJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMatchingJobs page.
func (p *ListMatchingJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMatchingJobsOutput, error) {
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

	result, err := p.client.ListMatchingJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMatchingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "entityresolution",
		OperationName: "ListMatchingJobs",
	}
}
