// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve the dimensions that can be queried for each specified metric type on a
// specified DB instance.
func (c *Client) ListAvailableResourceDimensions(ctx context.Context, params *ListAvailableResourceDimensionsInput, optFns ...func(*Options)) (*ListAvailableResourceDimensionsOutput, error) {
	if params == nil {
		params = &ListAvailableResourceDimensionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAvailableResourceDimensions", params, optFns, c.addOperationListAvailableResourceDimensionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAvailableResourceDimensionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAvailableResourceDimensionsInput struct {

	// An immutable identifier for a data source that is unique within an Amazon Web
	// Services Region. Performance Insights gathers metrics from this data source. To
	// use an Amazon RDS DB instance as a data source, specify its DbiResourceId
	// value. For example, specify db-ABCDEFGHIJKLMNOPQRSTU1VWZ .
	//
	// This member is required.
	Identifier *string

	// The types of metrics for which to retrieve dimensions. Valid values include
	// db.load .
	//
	// This member is required.
	Metrics []string

	// The Amazon Web Services service for which Performance Insights returns metrics.
	//
	// This member is required.
	ServiceType types.ServiceType

	// The maximum number of items to return in the response. If more items exist than
	// the specified MaxRecords value, a pagination token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int32

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to the
	// value specified by MaxRecords .
	NextToken *string

	noSmithyDocumentSerde
}

func (*ListAvailableResourceDimensionsInput) operationName() string {
	return "ListAvailableResourceDimensions"
}

type ListAvailableResourceDimensionsOutput struct {

	// The dimension information returned for requested metric types.
	MetricDimensions []types.MetricDimensionGroups

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to the
	// value specified by MaxRecords .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAvailableResourceDimensionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAvailableResourceDimensions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAvailableResourceDimensions{}, middleware.After)
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
	if err = addListAvailableResourceDimensionsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListAvailableResourceDimensionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAvailableResourceDimensions(options.Region), middleware.Before); err != nil {
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

// ListAvailableResourceDimensionsAPIClient is a client that implements the
// ListAvailableResourceDimensions operation.
type ListAvailableResourceDimensionsAPIClient interface {
	ListAvailableResourceDimensions(context.Context, *ListAvailableResourceDimensionsInput, ...func(*Options)) (*ListAvailableResourceDimensionsOutput, error)
}

var _ ListAvailableResourceDimensionsAPIClient = (*Client)(nil)

// ListAvailableResourceDimensionsPaginatorOptions is the paginator options for
// ListAvailableResourceDimensions
type ListAvailableResourceDimensionsPaginatorOptions struct {
	// The maximum number of items to return in the response. If more items exist than
	// the specified MaxRecords value, a pagination token is included in the response
	// so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAvailableResourceDimensionsPaginator is a paginator for
// ListAvailableResourceDimensions
type ListAvailableResourceDimensionsPaginator struct {
	options   ListAvailableResourceDimensionsPaginatorOptions
	client    ListAvailableResourceDimensionsAPIClient
	params    *ListAvailableResourceDimensionsInput
	nextToken *string
	firstPage bool
}

// NewListAvailableResourceDimensionsPaginator returns a new
// ListAvailableResourceDimensionsPaginator
func NewListAvailableResourceDimensionsPaginator(client ListAvailableResourceDimensionsAPIClient, params *ListAvailableResourceDimensionsInput, optFns ...func(*ListAvailableResourceDimensionsPaginatorOptions)) *ListAvailableResourceDimensionsPaginator {
	if params == nil {
		params = &ListAvailableResourceDimensionsInput{}
	}

	options := ListAvailableResourceDimensionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAvailableResourceDimensionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAvailableResourceDimensionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAvailableResourceDimensions page.
func (p *ListAvailableResourceDimensionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAvailableResourceDimensionsOutput, error) {
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

	result, err := p.client.ListAvailableResourceDimensions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAvailableResourceDimensions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pi",
		OperationName: "ListAvailableResourceDimensions",
	}
}

type opListAvailableResourceDimensionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListAvailableResourceDimensionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListAvailableResourceDimensionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "pi"
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
				signingName = "pi"
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
				v4aScheme.SigningName = aws.String("pi")
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

func addListAvailableResourceDimensionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListAvailableResourceDimensionsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
