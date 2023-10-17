// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the summaries of all insights in the specified group matching the
// provided filter values.
func (c *Client) GetInsightSummaries(ctx context.Context, params *GetInsightSummariesInput, optFns ...func(*Options)) (*GetInsightSummariesOutput, error) {
	if params == nil {
		params = &GetInsightSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInsightSummaries", params, optFns, c.addOperationGetInsightSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInsightSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightSummariesInput struct {

	// The end of the time frame in which the insights ended. The end time can't be
	// more than 30 days old.
	//
	// This member is required.
	EndTime *time.Time

	// The beginning of the time frame in which the insights started. The start time
	// can't be more than 30 days old.
	//
	// This member is required.
	StartTime *time.Time

	// The Amazon Resource Name (ARN) of the group. Required if the GroupName isn't
	// provided.
	GroupARN *string

	// The name of the group. Required if the GroupARN isn't provided.
	GroupName *string

	// The maximum number of results to display.
	MaxResults *int32

	// Pagination token.
	NextToken *string

	// The list of insight states.
	States []types.InsightState

	noSmithyDocumentSerde
}

type GetInsightSummariesOutput struct {

	// The summary of each insight within the group matching the provided filters. The
	// summary contains the InsightID, start and end time, the root cause service, the
	// root cause and client impact statistics, the top anomalous services, and the
	// status of the insight.
	InsightSummaries []types.InsightSummary

	// Pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInsightSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInsightSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInsightSummaries{}, middleware.After)
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
	if err = addOpGetInsightSummariesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsightSummaries(options.Region), middleware.Before); err != nil {
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

// GetInsightSummariesAPIClient is a client that implements the
// GetInsightSummaries operation.
type GetInsightSummariesAPIClient interface {
	GetInsightSummaries(context.Context, *GetInsightSummariesInput, ...func(*Options)) (*GetInsightSummariesOutput, error)
}

var _ GetInsightSummariesAPIClient = (*Client)(nil)

// GetInsightSummariesPaginatorOptions is the paginator options for
// GetInsightSummaries
type GetInsightSummariesPaginatorOptions struct {
	// The maximum number of results to display.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetInsightSummariesPaginator is a paginator for GetInsightSummaries
type GetInsightSummariesPaginator struct {
	options   GetInsightSummariesPaginatorOptions
	client    GetInsightSummariesAPIClient
	params    *GetInsightSummariesInput
	nextToken *string
	firstPage bool
}

// NewGetInsightSummariesPaginator returns a new GetInsightSummariesPaginator
func NewGetInsightSummariesPaginator(client GetInsightSummariesAPIClient, params *GetInsightSummariesInput, optFns ...func(*GetInsightSummariesPaginatorOptions)) *GetInsightSummariesPaginator {
	if params == nil {
		params = &GetInsightSummariesInput{}
	}

	options := GetInsightSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetInsightSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetInsightSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetInsightSummaries page.
func (p *GetInsightSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetInsightSummariesOutput, error) {
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

	result, err := p.client.GetInsightSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetInsightSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetInsightSummaries",
	}
}
