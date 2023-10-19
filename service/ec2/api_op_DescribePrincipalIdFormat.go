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
)

// Describes the ID format settings for the root user and all IAM roles and IAM
// users that have explicitly specified a longer ID (17-character ID) preference.
// By default, all IAM roles and IAM users default to the same ID settings as the
// root user, unless they explicitly override the settings. This request is useful
// for identifying those IAM users and IAM roles that have overridden the default
// ID settings. The following resource types support longer IDs: bundle |
// conversion-task | customer-gateway | dhcp-options | elastic-ip-allocation |
// elastic-ip-association | export-task | flow-log | image | import-task | instance
// | internet-gateway | network-acl | network-acl-association | network-interface
// | network-interface-attachment | prefix-list | reservation | route-table |
// route-table-association | security-group | snapshot | subnet |
// subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association |
// vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway .
func (c *Client) DescribePrincipalIdFormat(ctx context.Context, params *DescribePrincipalIdFormatInput, optFns ...func(*Options)) (*DescribePrincipalIdFormatOutput, error) {
	if params == nil {
		params = &DescribePrincipalIdFormatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePrincipalIdFormat", params, optFns, c.addOperationDescribePrincipalIdFormatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePrincipalIdFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePrincipalIdFormatInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The type of resource: bundle | conversion-task | customer-gateway | dhcp-options
	// | elastic-ip-allocation | elastic-ip-association | export-task | flow-log |
	// image | import-task | instance | internet-gateway | network-acl |
	// network-acl-association | network-interface | network-interface-attachment |
	// prefix-list | reservation | route-table | route-table-association |
	// security-group | snapshot | subnet | subnet-cidr-block-association | volume |
	// vpc | vpc-cidr-block-association | vpc-endpoint | vpc-peering-connection |
	// vpn-connection | vpn-gateway
	Resources []string

	noSmithyDocumentSerde
}

type DescribePrincipalIdFormatOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the ID format settings for the ARN.
	Principals []types.PrincipalIdFormat

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePrincipalIdFormatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribePrincipalIdFormat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribePrincipalIdFormat{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePrincipalIdFormat(options.Region), middleware.Before); err != nil {
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
		operation: "DescribePrincipalIdFormat",
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

// DescribePrincipalIdFormatAPIClient is a client that implements the
// DescribePrincipalIdFormat operation.
type DescribePrincipalIdFormatAPIClient interface {
	DescribePrincipalIdFormat(context.Context, *DescribePrincipalIdFormatInput, ...func(*Options)) (*DescribePrincipalIdFormatOutput, error)
}

var _ DescribePrincipalIdFormatAPIClient = (*Client)(nil)

// DescribePrincipalIdFormatPaginatorOptions is the paginator options for
// DescribePrincipalIdFormat
type DescribePrincipalIdFormatPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePrincipalIdFormatPaginator is a paginator for DescribePrincipalIdFormat
type DescribePrincipalIdFormatPaginator struct {
	options   DescribePrincipalIdFormatPaginatorOptions
	client    DescribePrincipalIdFormatAPIClient
	params    *DescribePrincipalIdFormatInput
	nextToken *string
	firstPage bool
}

// NewDescribePrincipalIdFormatPaginator returns a new
// DescribePrincipalIdFormatPaginator
func NewDescribePrincipalIdFormatPaginator(client DescribePrincipalIdFormatAPIClient, params *DescribePrincipalIdFormatInput, optFns ...func(*DescribePrincipalIdFormatPaginatorOptions)) *DescribePrincipalIdFormatPaginator {
	if params == nil {
		params = &DescribePrincipalIdFormatInput{}
	}

	options := DescribePrincipalIdFormatPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePrincipalIdFormatPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePrincipalIdFormatPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePrincipalIdFormat page.
func (p *DescribePrincipalIdFormatPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePrincipalIdFormatOutput, error) {
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

	result, err := p.client.DescribePrincipalIdFormat(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePrincipalIdFormat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribePrincipalIdFormat",
	}
}
