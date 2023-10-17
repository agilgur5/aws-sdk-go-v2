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

// Returns information about a core network Connect peer associations.
func (c *Client) GetConnectPeerAssociations(ctx context.Context, params *GetConnectPeerAssociationsInput, optFns ...func(*Options)) (*GetConnectPeerAssociationsOutput, error) {
	if params == nil {
		params = &GetConnectPeerAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConnectPeerAssociations", params, optFns, c.addOperationGetConnectPeerAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConnectPeerAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConnectPeerAssociationsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The IDs of the Connect peers.
	ConnectPeerIds []string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetConnectPeerAssociationsOutput struct {

	// Displays a list of Connect peer associations.
	ConnectPeerAssociations []types.ConnectPeerAssociation

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConnectPeerAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConnectPeerAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConnectPeerAssociations{}, middleware.After)
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
	if err = addOpGetConnectPeerAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConnectPeerAssociations(options.Region), middleware.Before); err != nil {
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

// GetConnectPeerAssociationsAPIClient is a client that implements the
// GetConnectPeerAssociations operation.
type GetConnectPeerAssociationsAPIClient interface {
	GetConnectPeerAssociations(context.Context, *GetConnectPeerAssociationsInput, ...func(*Options)) (*GetConnectPeerAssociationsOutput, error)
}

var _ GetConnectPeerAssociationsAPIClient = (*Client)(nil)

// GetConnectPeerAssociationsPaginatorOptions is the paginator options for
// GetConnectPeerAssociations
type GetConnectPeerAssociationsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetConnectPeerAssociationsPaginator is a paginator for
// GetConnectPeerAssociations
type GetConnectPeerAssociationsPaginator struct {
	options   GetConnectPeerAssociationsPaginatorOptions
	client    GetConnectPeerAssociationsAPIClient
	params    *GetConnectPeerAssociationsInput
	nextToken *string
	firstPage bool
}

// NewGetConnectPeerAssociationsPaginator returns a new
// GetConnectPeerAssociationsPaginator
func NewGetConnectPeerAssociationsPaginator(client GetConnectPeerAssociationsAPIClient, params *GetConnectPeerAssociationsInput, optFns ...func(*GetConnectPeerAssociationsPaginatorOptions)) *GetConnectPeerAssociationsPaginator {
	if params == nil {
		params = &GetConnectPeerAssociationsInput{}
	}

	options := GetConnectPeerAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetConnectPeerAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetConnectPeerAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetConnectPeerAssociations page.
func (p *GetConnectPeerAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetConnectPeerAssociationsOutput, error) {
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

	result, err := p.client.GetConnectPeerAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetConnectPeerAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetConnectPeerAssociations",
	}
}
