// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of profiling groups. The profiling groups are returned as
// ProfilingGroupDescription (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html)
// objects.
func (c *Client) ListProfilingGroups(ctx context.Context, params *ListProfilingGroupsInput, optFns ...func(*Options)) (*ListProfilingGroupsOutput, error) {
	if params == nil {
		params = &ListProfilingGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfilingGroups", params, optFns, c.addOperationListProfilingGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfilingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the listProfilingGroupsRequest.
type ListProfilingGroupsInput struct {

	// A Boolean value indicating whether to include a description. If true , then a
	// list of ProfilingGroupDescription (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html)
	// objects that contain detailed information about profiling groups is returned. If
	// false , then a list of profiling group names is returned.
	IncludeDescription *bool

	// The maximum number of profiling groups results returned by ListProfilingGroups
	// in paginated output. When this parameter is used, ListProfilingGroups only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListProfilingGroups request with the returned nextToken value.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListProfilingGroups
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This token should be treated as an opaque
	// identifier that is only used to retrieve the next items in a list and not for
	// other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

// The structure representing the listProfilingGroupsResponse.
type ListProfilingGroupsOutput struct {

	// A returned list of profiling group names. A list of the names is returned only
	// if includeDescription is false , otherwise a list of ProfilingGroupDescription (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html)
	// objects is returned.
	//
	// This member is required.
	ProfilingGroupNames []string

	// The nextToken value to include in a future ListProfilingGroups request. When
	// the results of a ListProfilingGroups request exceed maxResults , this value can
	// be used to retrieve the next page of results. This value is null when there are
	// no more results to return.
	NextToken *string

	// A returned list ProfilingGroupDescription (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html)
	// objects. A list of ProfilingGroupDescription (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html)
	// objects is returned only if includeDescription is true , otherwise a list of
	// profiling group names is returned.
	ProfilingGroups []types.ProfilingGroupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfilingGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProfilingGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfilingGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfilingGroups(options.Region), middleware.Before); err != nil {
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

// ListProfilingGroupsAPIClient is a client that implements the
// ListProfilingGroups operation.
type ListProfilingGroupsAPIClient interface {
	ListProfilingGroups(context.Context, *ListProfilingGroupsInput, ...func(*Options)) (*ListProfilingGroupsOutput, error)
}

var _ ListProfilingGroupsAPIClient = (*Client)(nil)

// ListProfilingGroupsPaginatorOptions is the paginator options for
// ListProfilingGroups
type ListProfilingGroupsPaginatorOptions struct {
	// The maximum number of profiling groups results returned by ListProfilingGroups
	// in paginated output. When this parameter is used, ListProfilingGroups only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListProfilingGroups request with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProfilingGroupsPaginator is a paginator for ListProfilingGroups
type ListProfilingGroupsPaginator struct {
	options   ListProfilingGroupsPaginatorOptions
	client    ListProfilingGroupsAPIClient
	params    *ListProfilingGroupsInput
	nextToken *string
	firstPage bool
}

// NewListProfilingGroupsPaginator returns a new ListProfilingGroupsPaginator
func NewListProfilingGroupsPaginator(client ListProfilingGroupsAPIClient, params *ListProfilingGroupsInput, optFns ...func(*ListProfilingGroupsPaginatorOptions)) *ListProfilingGroupsPaginator {
	if params == nil {
		params = &ListProfilingGroupsInput{}
	}

	options := ListProfilingGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProfilingGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProfilingGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProfilingGroups page.
func (p *ListProfilingGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProfilingGroupsOutput, error) {
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

	result, err := p.client.ListProfilingGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProfilingGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "ListProfilingGroups",
	}
}
