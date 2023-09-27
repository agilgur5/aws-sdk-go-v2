// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the associations between virtual interface groups and local gateway
// route tables.
func (c *Client) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx context.Context, params *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, optFns ...func(*Options)) (*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	if params == nil {
		params = &DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations", params, optFns, c.addOperationDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//   - local-gateway-id - The ID of a local gateway.
	//   - local-gateway-route-table-arn - The Amazon Resource Name (ARN) of the local
	//   gateway route table for the virtual interface group.
	//   - local-gateway-route-table-id - The ID of the local gateway route table.
	//   - local-gateway-route-table-virtual-interface-group-association-id - The ID of
	//   the association.
	//   - local-gateway-route-table-virtual-interface-group-id - The ID of the virtual
	//   interface group.
	//   - owner-id - The ID of the Amazon Web Services account that owns the local
	//   gateway virtual interface group association.
	//   - state - The state of the association.
	Filters []types.Filter

	// The IDs of the associations.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociationIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

func (*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) operationName() string {
	return "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations"
}

type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput struct {

	// Information about the associations.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociations []types.LocalGatewayRouteTableVirtualInterfaceGroupAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations{}, middleware.After)
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
	if err = addDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(options.Region), middleware.Before); err != nil {
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

// DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsAPIClient is a
// client that implements the
// DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations operation.
type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsAPIClient interface {
	DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(context.Context, *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, ...func(*Options)) (*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error)
}

var _ DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsAPIClient = (*Client)(nil)

// DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginatorOptions
// is the paginator options for
// DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations
type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator is a
// paginator for DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations
type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator struct {
	options   DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginatorOptions
	client    DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsAPIClient
	params    *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator
// returns a new
// DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator
func NewDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator(client DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsAPIClient, params *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, optFns ...func(*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginatorOptions)) *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator {
	if params == nil {
		params = &DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput{}
	}

	options := DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next
// DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations page.
func (p *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
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

	result, err := p.client.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
	}
}

type opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ec2"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ec2"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ec2")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
