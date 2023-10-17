// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all batch prediction jobs or a specific job if you specify a job ID. This
// is a paginated API. If you provide a null maxResults, this action retrieves a
// maximum of 50 records per page. If you provide a maxResults, the value must be
// between 1 and 50. To get the next page results, provide the pagination token
// from the GetBatchPredictionJobsResponse as part of your request. A null
// pagination token fetches the records from the beginning.
func (c *Client) GetBatchPredictionJobs(ctx context.Context, params *GetBatchPredictionJobsInput, optFns ...func(*Options)) (*GetBatchPredictionJobsOutput, error) {
	if params == nil {
		params = &GetBatchPredictionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBatchPredictionJobs", params, optFns, c.addOperationGetBatchPredictionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBatchPredictionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBatchPredictionJobsInput struct {

	// The batch prediction job for which to get the details.
	JobId *string

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The next token from the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetBatchPredictionJobsOutput struct {

	// An array containing the details of each batch prediction job.
	BatchPredictions []types.BatchPrediction

	// The next token for the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBatchPredictionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetBatchPredictionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBatchPredictionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBatchPredictionJobs(options.Region), middleware.Before); err != nil {
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

// GetBatchPredictionJobsAPIClient is a client that implements the
// GetBatchPredictionJobs operation.
type GetBatchPredictionJobsAPIClient interface {
	GetBatchPredictionJobs(context.Context, *GetBatchPredictionJobsInput, ...func(*Options)) (*GetBatchPredictionJobsOutput, error)
}

var _ GetBatchPredictionJobsAPIClient = (*Client)(nil)

// GetBatchPredictionJobsPaginatorOptions is the paginator options for
// GetBatchPredictionJobs
type GetBatchPredictionJobsPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBatchPredictionJobsPaginator is a paginator for GetBatchPredictionJobs
type GetBatchPredictionJobsPaginator struct {
	options   GetBatchPredictionJobsPaginatorOptions
	client    GetBatchPredictionJobsAPIClient
	params    *GetBatchPredictionJobsInput
	nextToken *string
	firstPage bool
}

// NewGetBatchPredictionJobsPaginator returns a new GetBatchPredictionJobsPaginator
func NewGetBatchPredictionJobsPaginator(client GetBatchPredictionJobsAPIClient, params *GetBatchPredictionJobsInput, optFns ...func(*GetBatchPredictionJobsPaginatorOptions)) *GetBatchPredictionJobsPaginator {
	if params == nil {
		params = &GetBatchPredictionJobsInput{}
	}

	options := GetBatchPredictionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBatchPredictionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBatchPredictionJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBatchPredictionJobs page.
func (p *GetBatchPredictionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBatchPredictionJobsOutput, error) {
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

	result, err := p.client.GetBatchPredictionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetBatchPredictionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetBatchPredictionJobs",
	}
}
