// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of what-if forecast exports created using the
// CreateWhatIfForecastExport operation. For each what-if forecast export, this
// operation returns a summary of its properties, including its Amazon Resource
// Name (ARN). You can retrieve the complete set of properties by using the what-if
// forecast export ARN with the DescribeWhatIfForecastExport operation.
func (c *Client) ListWhatIfForecastExports(ctx context.Context, params *ListWhatIfForecastExportsInput, optFns ...func(*Options)) (*ListWhatIfForecastExportsOutput, error) {
	if params == nil {
		params = &ListWhatIfForecastExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWhatIfForecastExports", params, optFns, c.addOperationListWhatIfForecastExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWhatIfForecastExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWhatIfForecastExportsInput struct {

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT , which specifies whether to
	// include or exclude the what-if forecast export jobs that match the statement
	// from the list, respectively. The match statement consists of a key and a value.
	// Filter properties
	//   - Condition - The condition to apply. Valid values are IS and IS_NOT . To
	//   include the forecast export jobs that match the statement, specify IS . To
	//   exclude matching forecast export jobs, specify IS_NOT .
	//   - Key - The name of the parameter to filter on. Valid values are
	//   WhatIfForecastExportArn and Status .
	//   - Value - The value to match.
	// For example, to list all jobs that export a forecast named
	// electricityWIFExport, specify the following filter: "Filters": [ { "Condition":
	// "IS", "Key": "WhatIfForecastExportArn", "Value":
	// "arn:aws:forecast:us-west-2::forecast/electricityWIFExport" } ]
	Filters []types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken . To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

func (*ListWhatIfForecastExportsInput) operationName() string {
	return "ListWhatIfForecastExports"
}

type ListWhatIfForecastExportsOutput struct {

	// If the response is truncated, Forecast returns this token. To retrieve the next
	// set of results, use the token in the next request.
	NextToken *string

	// An array of WhatIfForecastExports objects that describe the matched forecast
	// exports.
	WhatIfForecastExports []types.WhatIfForecastExportSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWhatIfForecastExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWhatIfForecastExports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWhatIfForecastExports{}, middleware.After)
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
	if err = addListWhatIfForecastExportsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListWhatIfForecastExportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWhatIfForecastExports(options.Region), middleware.Before); err != nil {
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

// ListWhatIfForecastExportsAPIClient is a client that implements the
// ListWhatIfForecastExports operation.
type ListWhatIfForecastExportsAPIClient interface {
	ListWhatIfForecastExports(context.Context, *ListWhatIfForecastExportsInput, ...func(*Options)) (*ListWhatIfForecastExportsOutput, error)
}

var _ ListWhatIfForecastExportsAPIClient = (*Client)(nil)

// ListWhatIfForecastExportsPaginatorOptions is the paginator options for
// ListWhatIfForecastExports
type ListWhatIfForecastExportsPaginatorOptions struct {
	// The number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWhatIfForecastExportsPaginator is a paginator for ListWhatIfForecastExports
type ListWhatIfForecastExportsPaginator struct {
	options   ListWhatIfForecastExportsPaginatorOptions
	client    ListWhatIfForecastExportsAPIClient
	params    *ListWhatIfForecastExportsInput
	nextToken *string
	firstPage bool
}

// NewListWhatIfForecastExportsPaginator returns a new
// ListWhatIfForecastExportsPaginator
func NewListWhatIfForecastExportsPaginator(client ListWhatIfForecastExportsAPIClient, params *ListWhatIfForecastExportsInput, optFns ...func(*ListWhatIfForecastExportsPaginatorOptions)) *ListWhatIfForecastExportsPaginator {
	if params == nil {
		params = &ListWhatIfForecastExportsInput{}
	}

	options := ListWhatIfForecastExportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWhatIfForecastExportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWhatIfForecastExportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWhatIfForecastExports page.
func (p *ListWhatIfForecastExportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWhatIfForecastExportsOutput, error) {
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

	result, err := p.client.ListWhatIfForecastExports(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWhatIfForecastExports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListWhatIfForecastExports",
	}
}

type opListWhatIfForecastExportsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListWhatIfForecastExportsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListWhatIfForecastExportsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "forecast"
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
				signingName = "forecast"
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
				v4aScheme.SigningName = aws.String("forecast")
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

func addListWhatIfForecastExportsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListWhatIfForecastExportsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
