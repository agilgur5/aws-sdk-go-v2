// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the associations between the service network and the service. You can
// filter the list either by service or service network. You must provide either
// the service network identifier or the service identifier. Every association in
// Amazon VPC Lattice is given a unique Amazon Resource Name (ARN), such as when a
// service network is associated with a VPC or when a service is associated with a
// service network. If the association is for a resource that is shared with
// another account, the association will include the local account ID as the prefix
// in the ARN for each account the resource is shared with.
func (c *Client) ListServiceNetworkServiceAssociations(ctx context.Context, params *ListServiceNetworkServiceAssociationsInput, optFns ...func(*Options)) (*ListServiceNetworkServiceAssociationsOutput, error) {
	if params == nil {
		params = &ListServiceNetworkServiceAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceNetworkServiceAssociations", params, optFns, c.addOperationListServiceNetworkServiceAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceNetworkServiceAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceNetworkServiceAssociationsInput struct {

	// The maximum number of results to return.
	MaxResults *int32

	// A pagination token for the next page of results.
	NextToken *string

	// The ID or Amazon Resource Name (ARN) of the service.
	ServiceIdentifier *string

	// The ID or Amazon Resource Name (ARN) of the service network.
	ServiceNetworkIdentifier *string

	noSmithyDocumentSerde
}

type ListServiceNetworkServiceAssociationsOutput struct {

	// Information about the associations.
	//
	// This member is required.
	Items []types.ServiceNetworkServiceAssociationSummary

	// If there are additional results, a pagination token for the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceNetworkServiceAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListServiceNetworkServiceAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListServiceNetworkServiceAssociations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceNetworkServiceAssociations(options.Region), middleware.Before); err != nil {
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
		operation: "ListServiceNetworkServiceAssociations",
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

// ListServiceNetworkServiceAssociationsAPIClient is a client that implements the
// ListServiceNetworkServiceAssociations operation.
type ListServiceNetworkServiceAssociationsAPIClient interface {
	ListServiceNetworkServiceAssociations(context.Context, *ListServiceNetworkServiceAssociationsInput, ...func(*Options)) (*ListServiceNetworkServiceAssociationsOutput, error)
}

var _ ListServiceNetworkServiceAssociationsAPIClient = (*Client)(nil)

// ListServiceNetworkServiceAssociationsPaginatorOptions is the paginator options
// for ListServiceNetworkServiceAssociations
type ListServiceNetworkServiceAssociationsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceNetworkServiceAssociationsPaginator is a paginator for
// ListServiceNetworkServiceAssociations
type ListServiceNetworkServiceAssociationsPaginator struct {
	options   ListServiceNetworkServiceAssociationsPaginatorOptions
	client    ListServiceNetworkServiceAssociationsAPIClient
	params    *ListServiceNetworkServiceAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListServiceNetworkServiceAssociationsPaginator returns a new
// ListServiceNetworkServiceAssociationsPaginator
func NewListServiceNetworkServiceAssociationsPaginator(client ListServiceNetworkServiceAssociationsAPIClient, params *ListServiceNetworkServiceAssociationsInput, optFns ...func(*ListServiceNetworkServiceAssociationsPaginatorOptions)) *ListServiceNetworkServiceAssociationsPaginator {
	if params == nil {
		params = &ListServiceNetworkServiceAssociationsInput{}
	}

	options := ListServiceNetworkServiceAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceNetworkServiceAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceNetworkServiceAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceNetworkServiceAssociations page.
func (p *ListServiceNetworkServiceAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceNetworkServiceAssociationsOutput, error) {
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

	result, err := p.client.ListServiceNetworkServiceAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceNetworkServiceAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "ListServiceNetworkServiceAssociations",
	}
}
