// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the principals associated with the specified policy. Note: This action is
// deprecated and works as expected for backward compatibility, but we won't add
// enhancements. Use ListTargetsForPolicy instead. Requires permission to access
// the ListPolicyPrincipals (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
//
// Deprecated: This operation has been deprecated.
func (c *Client) ListPolicyPrincipals(ctx context.Context, params *ListPolicyPrincipalsInput, optFns ...func(*Options)) (*ListPolicyPrincipalsOutput, error) {
	if params == nil {
		params = &ListPolicyPrincipalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicyPrincipals", params, optFns, c.addOperationListPolicyPrincipalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPolicyPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPolicyPrincipals operation.
type ListPolicyPrincipalsInput struct {

	// The policy name.
	//
	// This member is required.
	PolicyName *string

	// Specifies the order for results. If true, the results are returned in ascending
	// creation order.
	AscendingOrder bool

	// The marker for the next set of results.
	Marker *string

	// The result page size.
	PageSize *int32

	noSmithyDocumentSerde
}

func (*ListPolicyPrincipalsInput) operationName() string {
	return "ListPolicyPrincipals"
}

// The output from the ListPolicyPrincipals operation.
type ListPolicyPrincipalsOutput struct {

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string

	// The descriptions of the principals.
	Principals []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPolicyPrincipalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPolicyPrincipals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPolicyPrincipals{}, middleware.After)
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
	if err = addListPolicyPrincipalsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListPolicyPrincipalsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyPrincipals(options.Region), middleware.Before); err != nil {
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

// ListPolicyPrincipalsAPIClient is a client that implements the
// ListPolicyPrincipals operation.
type ListPolicyPrincipalsAPIClient interface {
	ListPolicyPrincipals(context.Context, *ListPolicyPrincipalsInput, ...func(*Options)) (*ListPolicyPrincipalsOutput, error)
}

var _ ListPolicyPrincipalsAPIClient = (*Client)(nil)

// ListPolicyPrincipalsPaginatorOptions is the paginator options for
// ListPolicyPrincipals
type ListPolicyPrincipalsPaginatorOptions struct {
	// The result page size.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPolicyPrincipalsPaginator is a paginator for ListPolicyPrincipals
type ListPolicyPrincipalsPaginator struct {
	options   ListPolicyPrincipalsPaginatorOptions
	client    ListPolicyPrincipalsAPIClient
	params    *ListPolicyPrincipalsInput
	nextToken *string
	firstPage bool
}

// NewListPolicyPrincipalsPaginator returns a new ListPolicyPrincipalsPaginator
func NewListPolicyPrincipalsPaginator(client ListPolicyPrincipalsAPIClient, params *ListPolicyPrincipalsInput, optFns ...func(*ListPolicyPrincipalsPaginatorOptions)) *ListPolicyPrincipalsPaginator {
	if params == nil {
		params = &ListPolicyPrincipalsInput{}
	}

	options := ListPolicyPrincipalsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPolicyPrincipalsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPolicyPrincipalsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPolicyPrincipals page.
func (p *ListPolicyPrincipalsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPolicyPrincipalsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListPolicyPrincipals(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPolicyPrincipals(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "ListPolicyPrincipals",
	}
}

type opListPolicyPrincipalsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListPolicyPrincipalsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListPolicyPrincipalsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "iot"
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
				signingName = "iot"
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
				v4aScheme.SigningName = aws.String("iot")
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

func addListPolicyPrincipalsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListPolicyPrincipalsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
