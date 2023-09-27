// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes repositories that are in a public registry.
func (c *Client) DescribeRepositories(ctx context.Context, params *DescribeRepositoriesInput, optFns ...func(*Options)) (*DescribeRepositoriesOutput, error) {
	if params == nil {
		params = &DescribeRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRepositories", params, optFns, c.addOperationDescribeRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRepositoriesInput struct {

	// The maximum number of repository results that's returned by DescribeRepositories
	// in paginated output. When this parameter is used, DescribeRepositories only
	// returns maxResults results in a single page along with a nextToken response
	// element. You can see the remaining results of the initial request by sending
	// another DescribeRepositories request with the returned nextToken value. This
	// value can be between 1 and 1000. If this parameter isn't used, then
	// DescribeRepositories returns up to 100 results and a nextToken value, if
	// applicable. If you specify repositories with repositoryNames , you can't use
	// this option.
	MaxResults *int32

	// The nextToken value that's returned from a previous paginated
	// DescribeRepositories request where maxResults was used and the results exceeded
	// the value of that parameter. Pagination continues from the end of the previous
	// results that returned the nextToken value. If there are no more results to
	// return, this value is null . If you specify repositories with repositoryNames ,
	// you can't use this option. This token should be treated as an opaque identifier
	// that is only used to retrieve the next items in a list and not for other
	// programmatic purposes.
	NextToken *string

	// The Amazon Web Services account ID that's associated with the registry that
	// contains the repositories to be described. If you do not specify a registry, the
	// default public registry is assumed.
	RegistryId *string

	// A list of repositories to describe. If this parameter is omitted, then all
	// repositories in a registry are described.
	RepositoryNames []string

	noSmithyDocumentSerde
}

func (*DescribeRepositoriesInput) operationName() string {
	return "DescribeRepositories"
}

type DescribeRepositoriesOutput struct {

	// The nextToken value to include in a future DescribeRepositories request. When
	// the results of a DescribeRepositories request exceed maxResults , this value can
	// be used to retrieve the next page of results. If there are no more results to
	// return, this value is null .
	NextToken *string

	// A list of repository objects corresponding to valid repositories.
	Repositories []types.Repository

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRepositories{}, middleware.After)
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
	if err = addDescribeRepositoriesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRepositories(options.Region), middleware.Before); err != nil {
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

// DescribeRepositoriesAPIClient is a client that implements the
// DescribeRepositories operation.
type DescribeRepositoriesAPIClient interface {
	DescribeRepositories(context.Context, *DescribeRepositoriesInput, ...func(*Options)) (*DescribeRepositoriesOutput, error)
}

var _ DescribeRepositoriesAPIClient = (*Client)(nil)

// DescribeRepositoriesPaginatorOptions is the paginator options for
// DescribeRepositories
type DescribeRepositoriesPaginatorOptions struct {
	// The maximum number of repository results that's returned by DescribeRepositories
	// in paginated output. When this parameter is used, DescribeRepositories only
	// returns maxResults results in a single page along with a nextToken response
	// element. You can see the remaining results of the initial request by sending
	// another DescribeRepositories request with the returned nextToken value. This
	// value can be between 1 and 1000. If this parameter isn't used, then
	// DescribeRepositories returns up to 100 results and a nextToken value, if
	// applicable. If you specify repositories with repositoryNames , you can't use
	// this option.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRepositoriesPaginator is a paginator for DescribeRepositories
type DescribeRepositoriesPaginator struct {
	options   DescribeRepositoriesPaginatorOptions
	client    DescribeRepositoriesAPIClient
	params    *DescribeRepositoriesInput
	nextToken *string
	firstPage bool
}

// NewDescribeRepositoriesPaginator returns a new DescribeRepositoriesPaginator
func NewDescribeRepositoriesPaginator(client DescribeRepositoriesAPIClient, params *DescribeRepositoriesInput, optFns ...func(*DescribeRepositoriesPaginatorOptions)) *DescribeRepositoriesPaginator {
	if params == nil {
		params = &DescribeRepositoriesInput{}
	}

	options := DescribeRepositoriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRepositoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRepositoriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRepositories page.
func (p *DescribeRepositoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRepositoriesOutput, error) {
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

	result, err := p.client.DescribeRepositories(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRepositories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr-public",
		OperationName: "DescribeRepositories",
	}
}

type opDescribeRepositoriesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeRepositoriesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeRepositoriesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ecr-public"
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
				signingName = "ecr-public"
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
				v4aScheme.SigningName = aws.String("ecr-public")
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

func addDescribeRepositoriesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeRepositoriesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
