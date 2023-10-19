// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists traffic distribution group users.
func (c *Client) ListTrafficDistributionGroupUsers(ctx context.Context, params *ListTrafficDistributionGroupUsersInput, optFns ...func(*Options)) (*ListTrafficDistributionGroupUsersOutput, error) {
	if params == nil {
		params = &ListTrafficDistributionGroupUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrafficDistributionGroupUsers", params, optFns, c.addOperationListTrafficDistributionGroupUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrafficDistributionGroupUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrafficDistributionGroupUsersInput struct {

	// The identifier of the traffic distribution group. This can be the ID or the ARN
	// if the API is being called in the Region where the traffic distribution group
	// was created. The ARN must be provided if the call is from the replicated Region.
	//
	// This member is required.
	TrafficDistributionGroupId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTrafficDistributionGroupUsersOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// A list of traffic distribution group users.
	TrafficDistributionGroupUserSummaryList []types.TrafficDistributionGroupUserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrafficDistributionGroupUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTrafficDistributionGroupUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTrafficDistributionGroupUsers{}, middleware.After)
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
	if err = addOpListTrafficDistributionGroupUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficDistributionGroupUsers(options.Region), middleware.Before); err != nil {
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
		operation: "ListTrafficDistributionGroupUsers",
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

// ListTrafficDistributionGroupUsersAPIClient is a client that implements the
// ListTrafficDistributionGroupUsers operation.
type ListTrafficDistributionGroupUsersAPIClient interface {
	ListTrafficDistributionGroupUsers(context.Context, *ListTrafficDistributionGroupUsersInput, ...func(*Options)) (*ListTrafficDistributionGroupUsersOutput, error)
}

var _ ListTrafficDistributionGroupUsersAPIClient = (*Client)(nil)

// ListTrafficDistributionGroupUsersPaginatorOptions is the paginator options for
// ListTrafficDistributionGroupUsers
type ListTrafficDistributionGroupUsersPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrafficDistributionGroupUsersPaginator is a paginator for
// ListTrafficDistributionGroupUsers
type ListTrafficDistributionGroupUsersPaginator struct {
	options   ListTrafficDistributionGroupUsersPaginatorOptions
	client    ListTrafficDistributionGroupUsersAPIClient
	params    *ListTrafficDistributionGroupUsersInput
	nextToken *string
	firstPage bool
}

// NewListTrafficDistributionGroupUsersPaginator returns a new
// ListTrafficDistributionGroupUsersPaginator
func NewListTrafficDistributionGroupUsersPaginator(client ListTrafficDistributionGroupUsersAPIClient, params *ListTrafficDistributionGroupUsersInput, optFns ...func(*ListTrafficDistributionGroupUsersPaginatorOptions)) *ListTrafficDistributionGroupUsersPaginator {
	if params == nil {
		params = &ListTrafficDistributionGroupUsersInput{}
	}

	options := ListTrafficDistributionGroupUsersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrafficDistributionGroupUsersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrafficDistributionGroupUsersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrafficDistributionGroupUsers page.
func (p *ListTrafficDistributionGroupUsersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrafficDistributionGroupUsersOutput, error) {
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

	result, err := p.client.ListTrafficDistributionGroupUsers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTrafficDistributionGroupUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListTrafficDistributionGroupUsers",
	}
}
