// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the storage tier status of one or more Amazon EBS snapshots.
func (c *Client) DescribeSnapshotTierStatus(ctx context.Context, params *DescribeSnapshotTierStatusInput, optFns ...func(*Options)) (*DescribeSnapshotTierStatusOutput, error) {
	if params == nil {
		params = &DescribeSnapshotTierStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshotTierStatus", params, optFns, c.addOperationDescribeSnapshotTierStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSnapshotTierStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSnapshotTierStatusInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - snapshot-id - The snapshot ID.
	//   - volume-id - The ID of the volume the snapshot is for.
	//   - last-tiering-operation - The state of the last archive or restore action. (
	//   archival-in-progress | archival-completed | archival-failed |
	//   permanent-restore-in-progress | permanent-restore-completed |
	//   permanent-restore-failed | temporary-restore-in-progress |
	//   temporary-restore-completed | temporary-restore-failed )
	Filters []types.Filter

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeSnapshotTierStatusOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Information about the snapshot's storage tier.
	SnapshotTierStatuses []types.SnapshotTierStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSnapshotTierStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSnapshotTierStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSnapshotTierStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshotTierStatus(options.Region), middleware.Before); err != nil {
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

// DescribeSnapshotTierStatusAPIClient is a client that implements the
// DescribeSnapshotTierStatus operation.
type DescribeSnapshotTierStatusAPIClient interface {
	DescribeSnapshotTierStatus(context.Context, *DescribeSnapshotTierStatusInput, ...func(*Options)) (*DescribeSnapshotTierStatusOutput, error)
}

var _ DescribeSnapshotTierStatusAPIClient = (*Client)(nil)

// DescribeSnapshotTierStatusPaginatorOptions is the paginator options for
// DescribeSnapshotTierStatus
type DescribeSnapshotTierStatusPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSnapshotTierStatusPaginator is a paginator for
// DescribeSnapshotTierStatus
type DescribeSnapshotTierStatusPaginator struct {
	options   DescribeSnapshotTierStatusPaginatorOptions
	client    DescribeSnapshotTierStatusAPIClient
	params    *DescribeSnapshotTierStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeSnapshotTierStatusPaginator returns a new
// DescribeSnapshotTierStatusPaginator
func NewDescribeSnapshotTierStatusPaginator(client DescribeSnapshotTierStatusAPIClient, params *DescribeSnapshotTierStatusInput, optFns ...func(*DescribeSnapshotTierStatusPaginatorOptions)) *DescribeSnapshotTierStatusPaginator {
	if params == nil {
		params = &DescribeSnapshotTierStatusInput{}
	}

	options := DescribeSnapshotTierStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSnapshotTierStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSnapshotTierStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSnapshotTierStatus page.
func (p *DescribeSnapshotTierStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSnapshotTierStatusOutput, error) {
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

	result, err := p.client.DescribeSnapshotTierStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSnapshotTierStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSnapshotTierStatus",
	}
}
