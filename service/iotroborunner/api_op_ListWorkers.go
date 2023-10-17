// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotroborunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants permission to list workers
func (c *Client) ListWorkers(ctx context.Context, params *ListWorkersInput, optFns ...func(*Options)) (*ListWorkersOutput, error) {
	if params == nil {
		params = &ListWorkersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkers", params, optFns, c.addOperationListWorkersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkersInput struct {

	// Site ARN.
	//
	// This member is required.
	Site *string

	// Full ARN of the worker fleet.
	Fleet *string

	// Maximum number of results to retrieve in a single ListWorkers call.
	MaxResults *int32

	// Pagination token returned when another page of data exists. Provide it in your
	// next call to the API to receive the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkersOutput struct {

	// Pagination token returned when another page of data exists. Provide it in your
	// next call to the API to receive the next page.
	NextToken *string

	// List of workers.
	Workers []types.Worker

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkers{}, middleware.After)
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
	if err = addOpListWorkersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkers(options.Region), middleware.Before); err != nil {
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

// ListWorkersAPIClient is a client that implements the ListWorkers operation.
type ListWorkersAPIClient interface {
	ListWorkers(context.Context, *ListWorkersInput, ...func(*Options)) (*ListWorkersOutput, error)
}

var _ ListWorkersAPIClient = (*Client)(nil)

// ListWorkersPaginatorOptions is the paginator options for ListWorkers
type ListWorkersPaginatorOptions struct {
	// Maximum number of results to retrieve in a single ListWorkers call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkersPaginator is a paginator for ListWorkers
type ListWorkersPaginator struct {
	options   ListWorkersPaginatorOptions
	client    ListWorkersAPIClient
	params    *ListWorkersInput
	nextToken *string
	firstPage bool
}

// NewListWorkersPaginator returns a new ListWorkersPaginator
func NewListWorkersPaginator(client ListWorkersAPIClient, params *ListWorkersInput, optFns ...func(*ListWorkersPaginatorOptions)) *ListWorkersPaginator {
	if params == nil {
		params = &ListWorkersInput{}
	}

	options := ListWorkersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkers page.
func (p *ListWorkersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkersOutput, error) {
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

	result, err := p.client.ListWorkers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotroborunner",
		OperationName: "ListWorkers",
	}
}
