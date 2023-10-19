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

// Lists information about function packages. A function package is a .zip file in
// CSAR (Cloud Service Archive) format that contains a network function (an ETSI
// standard telecommunication application) and function package descriptor that
// uses the TOSCA standard to describe how the network functions should run on your
// network.
func (c *Client) ListSolFunctionPackages(ctx context.Context, params *ListSolFunctionPackagesInput, optFns ...func(*Options)) (*ListSolFunctionPackagesOutput, error) {
	if params == nil {
		params = &ListSolFunctionPackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSolFunctionPackages", params, optFns, c.addOperationListSolFunctionPackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSolFunctionPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolFunctionPackagesInput struct {

	// The maximum number of results to include in the response.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSolFunctionPackagesOutput struct {

	// Function packages. A function package is a .zip file in CSAR (Cloud Service
	// Archive) format that contains a network function (an ETSI standard
	// telecommunication application) and function package descriptor that uses the
	// TOSCA standard to describe how the network functions should run on your network.
	//
	// This member is required.
	FunctionPackages []types.ListSolFunctionPackageInfo

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSolFunctionPackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSolFunctionPackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSolFunctionPackages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSolFunctionPackages(options.Region), middleware.Before); err != nil {
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
		operation: "ListSolFunctionPackages",
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

// ListSolFunctionPackagesAPIClient is a client that implements the
// ListSolFunctionPackages operation.
type ListSolFunctionPackagesAPIClient interface {
	ListSolFunctionPackages(context.Context, *ListSolFunctionPackagesInput, ...func(*Options)) (*ListSolFunctionPackagesOutput, error)
}

var _ ListSolFunctionPackagesAPIClient = (*Client)(nil)

// ListSolFunctionPackagesPaginatorOptions is the paginator options for
// ListSolFunctionPackages
type ListSolFunctionPackagesPaginatorOptions struct {
	// The maximum number of results to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSolFunctionPackagesPaginator is a paginator for ListSolFunctionPackages
type ListSolFunctionPackagesPaginator struct {
	options   ListSolFunctionPackagesPaginatorOptions
	client    ListSolFunctionPackagesAPIClient
	params    *ListSolFunctionPackagesInput
	nextToken *string
	firstPage bool
}

// NewListSolFunctionPackagesPaginator returns a new
// ListSolFunctionPackagesPaginator
func NewListSolFunctionPackagesPaginator(client ListSolFunctionPackagesAPIClient, params *ListSolFunctionPackagesInput, optFns ...func(*ListSolFunctionPackagesPaginatorOptions)) *ListSolFunctionPackagesPaginator {
	if params == nil {
		params = &ListSolFunctionPackagesInput{}
	}

	options := ListSolFunctionPackagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSolFunctionPackagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSolFunctionPackagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSolFunctionPackages page.
func (p *ListSolFunctionPackagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSolFunctionPackagesOutput, error) {
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

	result, err := p.client.ListSolFunctionPackages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSolFunctionPackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "ListSolFunctionPackages",
	}
}
