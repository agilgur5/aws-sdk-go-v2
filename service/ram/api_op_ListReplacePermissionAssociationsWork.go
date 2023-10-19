// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the current status of the asynchronous tasks performed by RAM when
// you perform the ReplacePermissionAssociationsWork operation.
func (c *Client) ListReplacePermissionAssociationsWork(ctx context.Context, params *ListReplacePermissionAssociationsWorkInput, optFns ...func(*Options)) (*ListReplacePermissionAssociationsWorkOutput, error) {
	if params == nil {
		params = &ListReplacePermissionAssociationsWorkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReplacePermissionAssociationsWork", params, optFns, c.addOperationListReplacePermissionAssociationsWorkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReplacePermissionAssociationsWorkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReplacePermissionAssociationsWorkInput struct {

	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	// Specifies that you want to see only the details about requests with a status
	// that matches this value.
	Status types.ReplacePermissionAssociationsWorkStatus

	// A list of IDs. These values come from the id field of the
	// replacePermissionAssociationsWork structure returned by the
	// ReplacePermissionAssociations operation.
	WorkIds []string

	noSmithyDocumentSerde
}

type ListReplacePermissionAssociationsWorkOutput struct {

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// An array of data structures that provide details of the matching work IDs.
	ReplacePermissionAssociationsWorks []types.ReplacePermissionAssociationsWork

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReplacePermissionAssociationsWorkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReplacePermissionAssociationsWork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReplacePermissionAssociationsWork{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReplacePermissionAssociationsWork(options.Region), middleware.Before); err != nil {
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
		operation: "ListReplacePermissionAssociationsWork",
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

// ListReplacePermissionAssociationsWorkAPIClient is a client that implements the
// ListReplacePermissionAssociationsWork operation.
type ListReplacePermissionAssociationsWorkAPIClient interface {
	ListReplacePermissionAssociationsWork(context.Context, *ListReplacePermissionAssociationsWorkInput, ...func(*Options)) (*ListReplacePermissionAssociationsWorkOutput, error)
}

var _ ListReplacePermissionAssociationsWorkAPIClient = (*Client)(nil)

// ListReplacePermissionAssociationsWorkPaginatorOptions is the paginator options
// for ListReplacePermissionAssociationsWork
type ListReplacePermissionAssociationsWorkPaginatorOptions struct {
	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReplacePermissionAssociationsWorkPaginator is a paginator for
// ListReplacePermissionAssociationsWork
type ListReplacePermissionAssociationsWorkPaginator struct {
	options   ListReplacePermissionAssociationsWorkPaginatorOptions
	client    ListReplacePermissionAssociationsWorkAPIClient
	params    *ListReplacePermissionAssociationsWorkInput
	nextToken *string
	firstPage bool
}

// NewListReplacePermissionAssociationsWorkPaginator returns a new
// ListReplacePermissionAssociationsWorkPaginator
func NewListReplacePermissionAssociationsWorkPaginator(client ListReplacePermissionAssociationsWorkAPIClient, params *ListReplacePermissionAssociationsWorkInput, optFns ...func(*ListReplacePermissionAssociationsWorkPaginatorOptions)) *ListReplacePermissionAssociationsWorkPaginator {
	if params == nil {
		params = &ListReplacePermissionAssociationsWorkInput{}
	}

	options := ListReplacePermissionAssociationsWorkPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReplacePermissionAssociationsWorkPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReplacePermissionAssociationsWorkPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReplacePermissionAssociationsWork page.
func (p *ListReplacePermissionAssociationsWorkPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReplacePermissionAssociationsWorkOutput, error) {
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

	result, err := p.client.ListReplacePermissionAssociationsWork(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReplacePermissionAssociationsWork(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "ListReplacePermissionAssociationsWork",
	}
}
