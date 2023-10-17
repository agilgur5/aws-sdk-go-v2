// Code generated by smithy-go-codegen DO NOT EDIT.

package simspaceweaver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/simspaceweaver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the SimSpace Weaver simulations in the Amazon Web Services account used
// to make the API call.
func (c *Client) ListSimulations(ctx context.Context, params *ListSimulationsInput, optFns ...func(*Options)) (*ListSimulationsOutput, error) {
	if params == nil {
		params = &ListSimulationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSimulations", params, optFns, c.addOperationListSimulationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSimulationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSimulationsInput struct {

	// The maximum number of simulations to list.
	MaxResults *int32

	// If SimSpace Weaver returns nextToken , then there are more results available.
	// The value of nextToken is a unique pagination token for each page. To retrieve
	// the next page, call the operation again using the returned token. Keep all other
	// arguments unchanged. If no results remain, then nextToken is set to null . Each
	// pagination token expires after 24 hours. If you provide a token that isn't
	// valid, then you receive an HTTP 400 ValidationException error.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSimulationsOutput struct {

	// If SimSpace Weaver returns nextToken , then there are more results available.
	// The value of nextToken is a unique pagination token for each page. To retrieve
	// the next page, call the operation again using the returned token. Keep all other
	// arguments unchanged. If no results remain, then nextToken is set to null . Each
	// pagination token expires after 24 hours. If you provide a token that isn't
	// valid, then you receive an HTTP 400 ValidationException error.
	NextToken *string

	// The list of simulations.
	Simulations []types.SimulationMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSimulationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSimulations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSimulations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSimulations(options.Region), middleware.Before); err != nil {
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

// ListSimulationsAPIClient is a client that implements the ListSimulations
// operation.
type ListSimulationsAPIClient interface {
	ListSimulations(context.Context, *ListSimulationsInput, ...func(*Options)) (*ListSimulationsOutput, error)
}

var _ ListSimulationsAPIClient = (*Client)(nil)

// ListSimulationsPaginatorOptions is the paginator options for ListSimulations
type ListSimulationsPaginatorOptions struct {
	// The maximum number of simulations to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSimulationsPaginator is a paginator for ListSimulations
type ListSimulationsPaginator struct {
	options   ListSimulationsPaginatorOptions
	client    ListSimulationsAPIClient
	params    *ListSimulationsInput
	nextToken *string
	firstPage bool
}

// NewListSimulationsPaginator returns a new ListSimulationsPaginator
func NewListSimulationsPaginator(client ListSimulationsAPIClient, params *ListSimulationsInput, optFns ...func(*ListSimulationsPaginatorOptions)) *ListSimulationsPaginator {
	if params == nil {
		params = &ListSimulationsInput{}
	}

	options := ListSimulationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSimulationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSimulationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSimulations page.
func (p *ListSimulationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSimulationsOutput, error) {
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

	result, err := p.client.ListSimulations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSimulations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "simspaceweaver",
		OperationName: "ListSimulations",
	}
}
