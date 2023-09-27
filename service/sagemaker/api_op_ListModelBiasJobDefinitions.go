// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists model bias jobs definitions that satisfy various filters.
func (c *Client) ListModelBiasJobDefinitions(ctx context.Context, params *ListModelBiasJobDefinitionsInput, optFns ...func(*Options)) (*ListModelBiasJobDefinitionsOutput, error) {
	if params == nil {
		params = &ListModelBiasJobDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelBiasJobDefinitions", params, optFns, c.addOperationListModelBiasJobDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelBiasJobDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelBiasJobDefinitionsInput struct {

	// A filter that returns only model bias jobs created after a specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only model bias jobs created before a specified time.
	CreationTimeBefore *time.Time

	// Name of the endpoint to monitor for model bias.
	EndpointName *string

	// The maximum number of model bias jobs to return in the response. The default
	// value is 10.
	MaxResults *int32

	// Filter for model bias jobs whose name contains a specified string.
	NameContains *string

	// The token returned if the response is truncated. To retrieve the next set of
	// job executions, use it in the next request.
	NextToken *string

	// Whether to sort results by the Name or CreationTime field. The default is
	// CreationTime .
	SortBy types.MonitoringJobDefinitionSortKey

	// Whether to sort the results in Ascending or Descending order. The default is
	// Descending .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

func (*ListModelBiasJobDefinitionsInput) operationName() string {
	return "ListModelBiasJobDefinitions"
}

type ListModelBiasJobDefinitionsOutput struct {

	// A JSON array in which each element is a summary for a model bias jobs.
	//
	// This member is required.
	JobDefinitionSummaries []types.MonitoringJobDefinitionSummary

	// The token returned if the response is truncated. To retrieve the next set of
	// job executions, use it in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelBiasJobDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModelBiasJobDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModelBiasJobDefinitions{}, middleware.After)
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
	if err = addListModelBiasJobDefinitionsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelBiasJobDefinitions(options.Region), middleware.Before); err != nil {
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

// ListModelBiasJobDefinitionsAPIClient is a client that implements the
// ListModelBiasJobDefinitions operation.
type ListModelBiasJobDefinitionsAPIClient interface {
	ListModelBiasJobDefinitions(context.Context, *ListModelBiasJobDefinitionsInput, ...func(*Options)) (*ListModelBiasJobDefinitionsOutput, error)
}

var _ ListModelBiasJobDefinitionsAPIClient = (*Client)(nil)

// ListModelBiasJobDefinitionsPaginatorOptions is the paginator options for
// ListModelBiasJobDefinitions
type ListModelBiasJobDefinitionsPaginatorOptions struct {
	// The maximum number of model bias jobs to return in the response. The default
	// value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelBiasJobDefinitionsPaginator is a paginator for
// ListModelBiasJobDefinitions
type ListModelBiasJobDefinitionsPaginator struct {
	options   ListModelBiasJobDefinitionsPaginatorOptions
	client    ListModelBiasJobDefinitionsAPIClient
	params    *ListModelBiasJobDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListModelBiasJobDefinitionsPaginator returns a new
// ListModelBiasJobDefinitionsPaginator
func NewListModelBiasJobDefinitionsPaginator(client ListModelBiasJobDefinitionsAPIClient, params *ListModelBiasJobDefinitionsInput, optFns ...func(*ListModelBiasJobDefinitionsPaginatorOptions)) *ListModelBiasJobDefinitionsPaginator {
	if params == nil {
		params = &ListModelBiasJobDefinitionsInput{}
	}

	options := ListModelBiasJobDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelBiasJobDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelBiasJobDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelBiasJobDefinitions page.
func (p *ListModelBiasJobDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelBiasJobDefinitionsOutput, error) {
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

	result, err := p.client.ListModelBiasJobDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModelBiasJobDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModelBiasJobDefinitions",
	}
}

type opListModelBiasJobDefinitionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListModelBiasJobDefinitionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListModelBiasJobDefinitionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "sagemaker"
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
				signingName = "sagemaker"
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
				v4aScheme.SigningName = aws.String("sagemaker")
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

func addListModelBiasJobDefinitionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListModelBiasJobDefinitionsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
