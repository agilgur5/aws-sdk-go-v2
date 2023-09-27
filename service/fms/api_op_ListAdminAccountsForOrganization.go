// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a AdminAccounts object that lists the Firewall Manager administrators
// within the organization that are onboarded to Firewall Manager by
// AssociateAdminAccount . This operation can be called only from the
// organization's management account.
func (c *Client) ListAdminAccountsForOrganization(ctx context.Context, params *ListAdminAccountsForOrganizationInput, optFns ...func(*Options)) (*ListAdminAccountsForOrganizationOutput, error) {
	if params == nil {
		params = &ListAdminAccountsForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAdminAccountsForOrganization", params, optFns, c.addOperationListAdminAccountsForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAdminAccountsForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAdminAccountsForOrganizationInput struct {

	// The maximum number of objects that you want Firewall Manager to return for this
	// request. If more objects are available, in the response, Firewall Manager
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	MaxResults *int32

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Firewall Manager returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

func (*ListAdminAccountsForOrganizationInput) operationName() string {
	return "ListAdminAccountsForOrganization"
}

type ListAdminAccountsForOrganizationOutput struct {

	// A list of Firewall Manager administrator accounts within the organization that
	// were onboarded as administrators by AssociateAdminAccount or PutAdminAccount .
	AdminAccounts []types.AdminAccountSummary

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Firewall Manager returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAdminAccountsForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAdminAccountsForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAdminAccountsForOrganization{}, middleware.After)
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
	if err = addListAdminAccountsForOrganizationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAdminAccountsForOrganization(options.Region), middleware.Before); err != nil {
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

// ListAdminAccountsForOrganizationAPIClient is a client that implements the
// ListAdminAccountsForOrganization operation.
type ListAdminAccountsForOrganizationAPIClient interface {
	ListAdminAccountsForOrganization(context.Context, *ListAdminAccountsForOrganizationInput, ...func(*Options)) (*ListAdminAccountsForOrganizationOutput, error)
}

var _ ListAdminAccountsForOrganizationAPIClient = (*Client)(nil)

// ListAdminAccountsForOrganizationPaginatorOptions is the paginator options for
// ListAdminAccountsForOrganization
type ListAdminAccountsForOrganizationPaginatorOptions struct {
	// The maximum number of objects that you want Firewall Manager to return for this
	// request. If more objects are available, in the response, Firewall Manager
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAdminAccountsForOrganizationPaginator is a paginator for
// ListAdminAccountsForOrganization
type ListAdminAccountsForOrganizationPaginator struct {
	options   ListAdminAccountsForOrganizationPaginatorOptions
	client    ListAdminAccountsForOrganizationAPIClient
	params    *ListAdminAccountsForOrganizationInput
	nextToken *string
	firstPage bool
}

// NewListAdminAccountsForOrganizationPaginator returns a new
// ListAdminAccountsForOrganizationPaginator
func NewListAdminAccountsForOrganizationPaginator(client ListAdminAccountsForOrganizationAPIClient, params *ListAdminAccountsForOrganizationInput, optFns ...func(*ListAdminAccountsForOrganizationPaginatorOptions)) *ListAdminAccountsForOrganizationPaginator {
	if params == nil {
		params = &ListAdminAccountsForOrganizationInput{}
	}

	options := ListAdminAccountsForOrganizationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAdminAccountsForOrganizationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAdminAccountsForOrganizationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAdminAccountsForOrganization page.
func (p *ListAdminAccountsForOrganizationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAdminAccountsForOrganizationOutput, error) {
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

	result, err := p.client.ListAdminAccountsForOrganization(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAdminAccountsForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListAdminAccountsForOrganization",
	}
}

type opListAdminAccountsForOrganizationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListAdminAccountsForOrganizationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListAdminAccountsForOrganizationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "fms"
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
				signingName = "fms"
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
				v4aScheme.SigningName = aws.String("fms")
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

func addListAdminAccountsForOrganizationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListAdminAccountsForOrganizationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
