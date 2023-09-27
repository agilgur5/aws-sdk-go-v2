// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of all policy stores in the calling Amazon Web
// Services account.
func (c *Client) ListPolicyStores(ctx context.Context, params *ListPolicyStoresInput, optFns ...func(*Options)) (*ListPolicyStoresOutput, error) {
	if params == nil {
		params = &ListPolicyStoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicyStores", params, optFns, c.addOperationListPolicyStoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPolicyStoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPolicyStoresInput struct {

	// Specifies the total number of results that you want included in each response.
	// If additional items exist beyond the number you specify, the NextToken response
	// element is returned with a value (not null). Include the specified value as the
	// NextToken request parameter in the next call to the operation to get the next
	// set of results. Note that the service might return fewer results than the
	// maximum even when there are more results available. You should check NextToken
	// after every operation to ensure that you receive all of the results. If you do
	// not specify this parameter, the operation defaults to 10 policy stores per
	// response. You can specify a maximum of 50 policy stores per response.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

func (*ListPolicyStoresInput) operationName() string {
	return "ListPolicyStores"
}

type ListPolicyStoresOutput struct {

	// The list of policy stores in the account.
	//
	// This member is required.
	PolicyStores []types.PolicyStoreItem

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPolicyStoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListPolicyStores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListPolicyStores{}, middleware.After)
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
	if err = addListPolicyStoresResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyStores(options.Region), middleware.Before); err != nil {
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

// ListPolicyStoresAPIClient is a client that implements the ListPolicyStores
// operation.
type ListPolicyStoresAPIClient interface {
	ListPolicyStores(context.Context, *ListPolicyStoresInput, ...func(*Options)) (*ListPolicyStoresOutput, error)
}

var _ ListPolicyStoresAPIClient = (*Client)(nil)

// ListPolicyStoresPaginatorOptions is the paginator options for ListPolicyStores
type ListPolicyStoresPaginatorOptions struct {
	// Specifies the total number of results that you want included in each response.
	// If additional items exist beyond the number you specify, the NextToken response
	// element is returned with a value (not null). Include the specified value as the
	// NextToken request parameter in the next call to the operation to get the next
	// set of results. Note that the service might return fewer results than the
	// maximum even when there are more results available. You should check NextToken
	// after every operation to ensure that you receive all of the results. If you do
	// not specify this parameter, the operation defaults to 10 policy stores per
	// response. You can specify a maximum of 50 policy stores per response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPolicyStoresPaginator is a paginator for ListPolicyStores
type ListPolicyStoresPaginator struct {
	options   ListPolicyStoresPaginatorOptions
	client    ListPolicyStoresAPIClient
	params    *ListPolicyStoresInput
	nextToken *string
	firstPage bool
}

// NewListPolicyStoresPaginator returns a new ListPolicyStoresPaginator
func NewListPolicyStoresPaginator(client ListPolicyStoresAPIClient, params *ListPolicyStoresInput, optFns ...func(*ListPolicyStoresPaginatorOptions)) *ListPolicyStoresPaginator {
	if params == nil {
		params = &ListPolicyStoresInput{}
	}

	options := ListPolicyStoresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPolicyStoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPolicyStoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPolicyStores page.
func (p *ListPolicyStoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPolicyStoresOutput, error) {
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

	result, err := p.client.ListPolicyStores(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPolicyStores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "verifiedpermissions",
		OperationName: "ListPolicyStores",
	}
}

type opListPolicyStoresResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListPolicyStoresResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListPolicyStoresResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "verifiedpermissions"
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
				signingName = "verifiedpermissions"
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
				v4aScheme.SigningName = aws.String("verifiedpermissions")
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

func addListPolicyStoresResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListPolicyStoresResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
