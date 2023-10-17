// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the workload shares associated with the workload.
func (c *Client) ListWorkloadShares(ctx context.Context, params *ListWorkloadSharesInput, optFns ...func(*Options)) (*ListWorkloadSharesOutput, error) {
	if params == nil {
		params = &ListWorkloadSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkloadShares", params, optFns, c.addOperationListWorkloadSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkloadSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for List Workload Share
type ListWorkloadSharesInput struct {

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	// The maximum number of results to return for this request.
	MaxResults int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The Amazon Web Services account ID, organization ID, or organizational unit
	// (OU) ID with which the workload is shared.
	SharedWithPrefix *string

	// The status of the share request.
	Status types.ShareStatus

	noSmithyDocumentSerde
}

// Input for List Workload Share
type ListWorkloadSharesOutput struct {

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// A list of workload share summaries.
	WorkloadShareSummaries []types.WorkloadShareSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkloadSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkloadShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkloadShares{}, middleware.After)
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
	if err = addOpListWorkloadSharesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkloadShares(options.Region), middleware.Before); err != nil {
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

// ListWorkloadSharesAPIClient is a client that implements the ListWorkloadShares
// operation.
type ListWorkloadSharesAPIClient interface {
	ListWorkloadShares(context.Context, *ListWorkloadSharesInput, ...func(*Options)) (*ListWorkloadSharesOutput, error)
}

var _ ListWorkloadSharesAPIClient = (*Client)(nil)

// ListWorkloadSharesPaginatorOptions is the paginator options for
// ListWorkloadShares
type ListWorkloadSharesPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkloadSharesPaginator is a paginator for ListWorkloadShares
type ListWorkloadSharesPaginator struct {
	options   ListWorkloadSharesPaginatorOptions
	client    ListWorkloadSharesAPIClient
	params    *ListWorkloadSharesInput
	nextToken *string
	firstPage bool
}

// NewListWorkloadSharesPaginator returns a new ListWorkloadSharesPaginator
func NewListWorkloadSharesPaginator(client ListWorkloadSharesAPIClient, params *ListWorkloadSharesInput, optFns ...func(*ListWorkloadSharesPaginatorOptions)) *ListWorkloadSharesPaginator {
	if params == nil {
		params = &ListWorkloadSharesInput{}
	}

	options := ListWorkloadSharesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkloadSharesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkloadSharesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkloadShares page.
func (p *ListWorkloadSharesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkloadSharesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListWorkloadShares(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkloadShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ListWorkloadShares",
	}
}
