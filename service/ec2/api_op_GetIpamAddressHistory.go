// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieve historical information about a CIDR within an IPAM scope. For more
// information, see View the history of IP addresses (https://docs.aws.amazon.com/vpc/latest/ipam/view-history-cidr-ipam.html)
// in the Amazon VPC IPAM User Guide.
func (c *Client) GetIpamAddressHistory(ctx context.Context, params *GetIpamAddressHistoryInput, optFns ...func(*Options)) (*GetIpamAddressHistoryOutput, error) {
	if params == nil {
		params = &GetIpamAddressHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIpamAddressHistory", params, optFns, c.addOperationGetIpamAddressHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIpamAddressHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIpamAddressHistoryInput struct {

	// The CIDR you want the history of. The CIDR can be an IPv4 or IPv6 IP address
	// range. If you enter a /16 IPv4 CIDR, you will get records that match it exactly.
	// You will not get records for any subnets within the /16 CIDR.
	//
	// This member is required.
	Cidr *string

	// The ID of the IPAM scope that the CIDR is in.
	//
	// This member is required.
	IpamScopeId *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The end of the time period for which you are looking for history. If you omit
	// this option, it will default to the current time.
	EndTime *time.Time

	// The maximum number of historical results you would like returned per page.
	// Defaults to 100.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The start of the time period for which you are looking for history. If you omit
	// this option, it will default to the value of EndTime.
	StartTime *time.Time

	// The ID of the VPC you want your history records filtered by.
	VpcId *string

	noSmithyDocumentSerde
}

type GetIpamAddressHistoryOutput struct {

	// A historical record for a CIDR within an IPAM scope. If the CIDR is associated
	// with an EC2 instance, you will see an object in the response for the instance
	// and one for the network interface.
	HistoryRecords []types.IpamAddressHistoryRecord

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIpamAddressHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetIpamAddressHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetIpamAddressHistory{}, middleware.After)
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
	if err = addOpGetIpamAddressHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIpamAddressHistory(options.Region), middleware.Before); err != nil {
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

// GetIpamAddressHistoryAPIClient is a client that implements the
// GetIpamAddressHistory operation.
type GetIpamAddressHistoryAPIClient interface {
	GetIpamAddressHistory(context.Context, *GetIpamAddressHistoryInput, ...func(*Options)) (*GetIpamAddressHistoryOutput, error)
}

var _ GetIpamAddressHistoryAPIClient = (*Client)(nil)

// GetIpamAddressHistoryPaginatorOptions is the paginator options for
// GetIpamAddressHistory
type GetIpamAddressHistoryPaginatorOptions struct {
	// The maximum number of historical results you would like returned per page.
	// Defaults to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetIpamAddressHistoryPaginator is a paginator for GetIpamAddressHistory
type GetIpamAddressHistoryPaginator struct {
	options   GetIpamAddressHistoryPaginatorOptions
	client    GetIpamAddressHistoryAPIClient
	params    *GetIpamAddressHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetIpamAddressHistoryPaginator returns a new GetIpamAddressHistoryPaginator
func NewGetIpamAddressHistoryPaginator(client GetIpamAddressHistoryAPIClient, params *GetIpamAddressHistoryInput, optFns ...func(*GetIpamAddressHistoryPaginatorOptions)) *GetIpamAddressHistoryPaginator {
	if params == nil {
		params = &GetIpamAddressHistoryInput{}
	}

	options := GetIpamAddressHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetIpamAddressHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetIpamAddressHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetIpamAddressHistory page.
func (p *GetIpamAddressHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetIpamAddressHistoryOutput, error) {
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

	result, err := p.client.GetIpamAddressHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetIpamAddressHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetIpamAddressHistory",
	}
}
