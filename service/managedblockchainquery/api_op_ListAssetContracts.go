// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the contracts for a given contract type deployed by an address
// (either a contract address or a wallet address). The Bitcoin blockchain networks
// do not support this operation.
func (c *Client) ListAssetContracts(ctx context.Context, params *ListAssetContractsInput, optFns ...func(*Options)) (*ListAssetContractsOutput, error) {
	if params == nil {
		params = &ListAssetContractsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssetContracts", params, optFns, c.addOperationListAssetContractsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetContractsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetContractsInput struct {

	// Contains the filter parameter for the request.
	//
	// This member is required.
	ContractFilter *types.ContractFilter

	// The maximum number of contracts to list.
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssetContractsOutput struct {

	// An array of contract objects that contain the properties for each contract.
	//
	// This member is required.
	Contracts []types.AssetContract

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssetContractsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssetContracts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssetContracts{}, middleware.After)
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
	if err = addOpListAssetContractsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssetContracts(options.Region), middleware.Before); err != nil {
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

// ListAssetContractsAPIClient is a client that implements the ListAssetContracts
// operation.
type ListAssetContractsAPIClient interface {
	ListAssetContracts(context.Context, *ListAssetContractsInput, ...func(*Options)) (*ListAssetContractsOutput, error)
}

var _ ListAssetContractsAPIClient = (*Client)(nil)

// ListAssetContractsPaginatorOptions is the paginator options for
// ListAssetContracts
type ListAssetContractsPaginatorOptions struct {
	// The maximum number of contracts to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetContractsPaginator is a paginator for ListAssetContracts
type ListAssetContractsPaginator struct {
	options   ListAssetContractsPaginatorOptions
	client    ListAssetContractsAPIClient
	params    *ListAssetContractsInput
	nextToken *string
	firstPage bool
}

// NewListAssetContractsPaginator returns a new ListAssetContractsPaginator
func NewListAssetContractsPaginator(client ListAssetContractsAPIClient, params *ListAssetContractsInput, optFns ...func(*ListAssetContractsPaginatorOptions)) *ListAssetContractsPaginator {
	if params == nil {
		params = &ListAssetContractsInput{}
	}

	options := ListAssetContractsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssetContractsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetContractsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssetContracts page.
func (p *ListAssetContractsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetContractsOutput, error) {
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

	result, err := p.client.ListAssetContracts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssetContracts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain-query",
		OperationName: "ListAssetContracts",
	}
}
