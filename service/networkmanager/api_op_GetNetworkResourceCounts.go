// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the count of network resources, by resource type, for the specified global
// network.
func (c *Client) GetNetworkResourceCounts(ctx context.Context, params *GetNetworkResourceCountsInput, optFns ...func(*Options)) (*GetNetworkResourceCountsOutput, error) {
	if params == nil {
		params = &GetNetworkResourceCountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetNetworkResourceCounts", params, optFns, c.addOperationGetNetworkResourceCountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetNetworkResourceCountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetNetworkResourceCountsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The resource type. The following are the supported resource types for Direct
	// Connect:
	//   - dxcon
	//   - dx-gateway
	//   - dx-vif
	// The following are the supported resource types for Network Manager:
	//   - connection
	//   - device
	//   - link
	//   - site
	// The following are the supported resource types for Amazon VPC:
	//   - customer-gateway
	//   - transit-gateway
	//   - transit-gateway-attachment
	//   - transit-gateway-connect-peer
	//   - transit-gateway-route-table
	//   - vpn-connection
	ResourceType *string

	noSmithyDocumentSerde
}

type GetNetworkResourceCountsOutput struct {

	// The count of resources.
	NetworkResourceCounts []types.NetworkResourceCount

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetNetworkResourceCountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetNetworkResourceCounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetNetworkResourceCounts{}, middleware.After)
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
	if err = addOpGetNetworkResourceCountsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetNetworkResourceCounts(options.Region), middleware.Before); err != nil {
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
		operation: "GetNetworkResourceCounts",
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

// GetNetworkResourceCountsAPIClient is a client that implements the
// GetNetworkResourceCounts operation.
type GetNetworkResourceCountsAPIClient interface {
	GetNetworkResourceCounts(context.Context, *GetNetworkResourceCountsInput, ...func(*Options)) (*GetNetworkResourceCountsOutput, error)
}

var _ GetNetworkResourceCountsAPIClient = (*Client)(nil)

// GetNetworkResourceCountsPaginatorOptions is the paginator options for
// GetNetworkResourceCounts
type GetNetworkResourceCountsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetNetworkResourceCountsPaginator is a paginator for GetNetworkResourceCounts
type GetNetworkResourceCountsPaginator struct {
	options   GetNetworkResourceCountsPaginatorOptions
	client    GetNetworkResourceCountsAPIClient
	params    *GetNetworkResourceCountsInput
	nextToken *string
	firstPage bool
}

// NewGetNetworkResourceCountsPaginator returns a new
// GetNetworkResourceCountsPaginator
func NewGetNetworkResourceCountsPaginator(client GetNetworkResourceCountsAPIClient, params *GetNetworkResourceCountsInput, optFns ...func(*GetNetworkResourceCountsPaginatorOptions)) *GetNetworkResourceCountsPaginator {
	if params == nil {
		params = &GetNetworkResourceCountsInput{}
	}

	options := GetNetworkResourceCountsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetNetworkResourceCountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetNetworkResourceCountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetNetworkResourceCounts page.
func (p *GetNetworkResourceCountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetNetworkResourceCountsOutput, error) {
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

	result, err := p.client.GetNetworkResourceCounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetNetworkResourceCounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetNetworkResourceCounts",
	}
}
