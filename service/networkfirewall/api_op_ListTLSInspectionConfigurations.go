// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the metadata for the TLS inspection configurations that you have
// defined. Depending on your setting for max results and the number of TLS
// inspection configurations, a single call might not return the full list.
func (c *Client) ListTLSInspectionConfigurations(ctx context.Context, params *ListTLSInspectionConfigurationsInput, optFns ...func(*Options)) (*ListTLSInspectionConfigurationsOutput, error) {
	if params == nil {
		params = &ListTLSInspectionConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTLSInspectionConfigurations", params, optFns, c.addOperationListTLSInspectionConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTLSInspectionConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTLSInspectionConfigurationsInput struct {

	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	MaxResults *int32

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

func (*ListTLSInspectionConfigurationsInput) operationName() string {
	return "ListTLSInspectionConfigurations"
}

type ListTLSInspectionConfigurationsOutput struct {

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// The TLS inspection configuration metadata objects that you've defined.
	// Depending on your setting for max results and the number of TLS inspection
	// configurations, this might not be the full list.
	TLSInspectionConfigurations []types.TLSInspectionConfigurationMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTLSInspectionConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListTLSInspectionConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListTLSInspectionConfigurations{}, middleware.After)
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
	if err = addListTLSInspectionConfigurationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTLSInspectionConfigurations(options.Region), middleware.Before); err != nil {
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

// ListTLSInspectionConfigurationsAPIClient is a client that implements the
// ListTLSInspectionConfigurations operation.
type ListTLSInspectionConfigurationsAPIClient interface {
	ListTLSInspectionConfigurations(context.Context, *ListTLSInspectionConfigurationsInput, ...func(*Options)) (*ListTLSInspectionConfigurationsOutput, error)
}

var _ ListTLSInspectionConfigurationsAPIClient = (*Client)(nil)

// ListTLSInspectionConfigurationsPaginatorOptions is the paginator options for
// ListTLSInspectionConfigurations
type ListTLSInspectionConfigurationsPaginatorOptions struct {
	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTLSInspectionConfigurationsPaginator is a paginator for
// ListTLSInspectionConfigurations
type ListTLSInspectionConfigurationsPaginator struct {
	options   ListTLSInspectionConfigurationsPaginatorOptions
	client    ListTLSInspectionConfigurationsAPIClient
	params    *ListTLSInspectionConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListTLSInspectionConfigurationsPaginator returns a new
// ListTLSInspectionConfigurationsPaginator
func NewListTLSInspectionConfigurationsPaginator(client ListTLSInspectionConfigurationsAPIClient, params *ListTLSInspectionConfigurationsInput, optFns ...func(*ListTLSInspectionConfigurationsPaginatorOptions)) *ListTLSInspectionConfigurationsPaginator {
	if params == nil {
		params = &ListTLSInspectionConfigurationsInput{}
	}

	options := ListTLSInspectionConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTLSInspectionConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTLSInspectionConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTLSInspectionConfigurations page.
func (p *ListTLSInspectionConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTLSInspectionConfigurationsOutput, error) {
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

	result, err := p.client.ListTLSInspectionConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTLSInspectionConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "ListTLSInspectionConfigurations",
	}
}

type opListTLSInspectionConfigurationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListTLSInspectionConfigurationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListTLSInspectionConfigurationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "network-firewall"
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
				signingName = "network-firewall"
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
				v4aScheme.SigningName = aws.String("network-firewall")
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

func addListTLSInspectionConfigurationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListTLSInspectionConfigurationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
