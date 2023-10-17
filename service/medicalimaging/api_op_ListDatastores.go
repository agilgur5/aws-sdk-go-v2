// Code generated by smithy-go-codegen DO NOT EDIT.

package medicalimaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medicalimaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List data stores created by this AWS account.
func (c *Client) ListDatastores(ctx context.Context, params *ListDatastoresInput, optFns ...func(*Options)) (*ListDatastoresOutput, error) {
	if params == nil {
		params = &ListDatastoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatastores", params, optFns, c.addOperationListDatastoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatastoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatastoresInput struct {

	// The data store status.
	DatastoreStatus types.DatastoreStatus

	// Valid Range: Minimum value of 1. Maximum value of 50.
	MaxResults *int32

	// The pagination token used to request the list of data stores on the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDatastoresOutput struct {

	// The list of summaries of data stores.
	DatastoreSummaries []types.DatastoreSummary

	// The pagination token used to retrieve the list of data stores on the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatastoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDatastores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDatastores{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatastores(options.Region), middleware.Before); err != nil {
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

// ListDatastoresAPIClient is a client that implements the ListDatastores
// operation.
type ListDatastoresAPIClient interface {
	ListDatastores(context.Context, *ListDatastoresInput, ...func(*Options)) (*ListDatastoresOutput, error)
}

var _ ListDatastoresAPIClient = (*Client)(nil)

// ListDatastoresPaginatorOptions is the paginator options for ListDatastores
type ListDatastoresPaginatorOptions struct {
	// Valid Range: Minimum value of 1. Maximum value of 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatastoresPaginator is a paginator for ListDatastores
type ListDatastoresPaginator struct {
	options   ListDatastoresPaginatorOptions
	client    ListDatastoresAPIClient
	params    *ListDatastoresInput
	nextToken *string
	firstPage bool
}

// NewListDatastoresPaginator returns a new ListDatastoresPaginator
func NewListDatastoresPaginator(client ListDatastoresAPIClient, params *ListDatastoresInput, optFns ...func(*ListDatastoresPaginatorOptions)) *ListDatastoresPaginator {
	if params == nil {
		params = &ListDatastoresInput{}
	}

	options := ListDatastoresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatastoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatastoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDatastores page.
func (p *ListDatastoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatastoresOutput, error) {
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

	result, err := p.client.ListDatastores(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDatastores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medical-imaging",
		OperationName: "ListDatastores",
	}
}
