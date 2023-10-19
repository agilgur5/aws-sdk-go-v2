// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the compute resources in an Amazon GameLift fleet. You can request
// information for either managed EC2 fleets or Anywhere fleets. To request a list
// of computes, specify the fleet ID. You can filter the result set by location.
// Use the pagination parameters to retrieve results in a set of sequential pages.
// If successful, this operation returns the compute resource for the requested
// fleet. For managed EC2 fleets, it returns a list of EC2 instances. For Anywhere
// fleets, it returns a list of registered compute names.
func (c *Client) ListCompute(ctx context.Context, params *ListComputeInput, optFns ...func(*Options)) (*ListComputeOutput, error) {
	if params == nil {
		params = &ListComputeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCompute", params, optFns, c.addOperationListComputeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComputeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComputeInput struct {

	// A unique identifier for the fleet to retrieve compute resources for.
	//
	// This member is required.
	FleetId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// The name of a location to retrieve compute resources for.
	Location *string

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComputeOutput struct {

	// A list of compute resources in the specified fleet.
	ComputeList []types.Compute

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComputeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCompute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCompute{}, middleware.After)
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
	if err = addOpListComputeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCompute(options.Region), middleware.Before); err != nil {
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
		operation: "ListCompute",
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

// ListComputeAPIClient is a client that implements the ListCompute operation.
type ListComputeAPIClient interface {
	ListCompute(context.Context, *ListComputeInput, ...func(*Options)) (*ListComputeOutput, error)
}

var _ ListComputeAPIClient = (*Client)(nil)

// ListComputePaginatorOptions is the paginator options for ListCompute
type ListComputePaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComputePaginator is a paginator for ListCompute
type ListComputePaginator struct {
	options   ListComputePaginatorOptions
	client    ListComputeAPIClient
	params    *ListComputeInput
	nextToken *string
	firstPage bool
}

// NewListComputePaginator returns a new ListComputePaginator
func NewListComputePaginator(client ListComputeAPIClient, params *ListComputeInput, optFns ...func(*ListComputePaginatorOptions)) *ListComputePaginator {
	if params == nil {
		params = &ListComputeInput{}
	}

	options := ListComputePaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComputePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComputePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCompute page.
func (p *ListComputePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComputeOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListCompute(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCompute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListCompute",
	}
}
