// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists network packages. A network package is a .zip file in CSAR (Cloud Service
// Archive) format defines the function packages you want to deploy and the Amazon
// Web Services infrastructure you want to deploy them on.
func (c *Client) ListSolNetworkPackages(ctx context.Context, params *ListSolNetworkPackagesInput, optFns ...func(*Options)) (*ListSolNetworkPackagesOutput, error) {
	if params == nil {
		params = &ListSolNetworkPackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSolNetworkPackages", params, optFns, c.addOperationListSolNetworkPackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSolNetworkPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolNetworkPackagesInput struct {

	// The maximum number of results to include in the response.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSolNetworkPackagesOutput struct {

	// Network packages. A network package is a .zip file in CSAR (Cloud Service
	// Archive) format defines the function packages you want to deploy and the Amazon
	// Web Services infrastructure you want to deploy them on.
	//
	// This member is required.
	NetworkPackages []types.ListSolNetworkPackageInfo

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSolNetworkPackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSolNetworkPackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSolNetworkPackages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSolNetworkPackages(options.Region), middleware.Before); err != nil {
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

// ListSolNetworkPackagesAPIClient is a client that implements the
// ListSolNetworkPackages operation.
type ListSolNetworkPackagesAPIClient interface {
	ListSolNetworkPackages(context.Context, *ListSolNetworkPackagesInput, ...func(*Options)) (*ListSolNetworkPackagesOutput, error)
}

var _ ListSolNetworkPackagesAPIClient = (*Client)(nil)

// ListSolNetworkPackagesPaginatorOptions is the paginator options for
// ListSolNetworkPackages
type ListSolNetworkPackagesPaginatorOptions struct {
	// The maximum number of results to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSolNetworkPackagesPaginator is a paginator for ListSolNetworkPackages
type ListSolNetworkPackagesPaginator struct {
	options   ListSolNetworkPackagesPaginatorOptions
	client    ListSolNetworkPackagesAPIClient
	params    *ListSolNetworkPackagesInput
	nextToken *string
	firstPage bool
}

// NewListSolNetworkPackagesPaginator returns a new ListSolNetworkPackagesPaginator
func NewListSolNetworkPackagesPaginator(client ListSolNetworkPackagesAPIClient, params *ListSolNetworkPackagesInput, optFns ...func(*ListSolNetworkPackagesPaginatorOptions)) *ListSolNetworkPackagesPaginator {
	if params == nil {
		params = &ListSolNetworkPackagesInput{}
	}

	options := ListSolNetworkPackagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSolNetworkPackagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSolNetworkPackagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSolNetworkPackages page.
func (p *ListSolNetworkPackagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSolNetworkPackagesOutput, error) {
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

	result, err := p.client.ListSolNetworkPackages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSolNetworkPackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "ListSolNetworkPackages",
	}
}
