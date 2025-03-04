// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a route calculator resource in your Amazon Web Services account. You
// can send requests to a route calculator resource to estimate travel time,
// distance, and get directions. A route calculator sources traffic and road
// network data from your chosen data provider. If your application is tracking or
// routing assets you use in your business, such as delivery vehicles or employees,
// you must not use Esri as your geolocation provider. See section 82 of the
// Amazon Web Services service terms (http://aws.amazon.com/service-terms) for more
// details.
func (c *Client) CreateRouteCalculator(ctx context.Context, params *CreateRouteCalculatorInput, optFns ...func(*Options)) (*CreateRouteCalculatorOutput, error) {
	if params == nil {
		params = &CreateRouteCalculatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRouteCalculator", params, optFns, c.addOperationCreateRouteCalculatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRouteCalculatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRouteCalculatorInput struct {

	// The name of the route calculator resource. Requirements:
	//   - Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.),
	//   and underscores (_).
	//   - Must be a unique Route calculator resource name.
	//   - No spaces allowed. For example, ExampleRouteCalculator .
	//
	// This member is required.
	CalculatorName *string

	// Specifies the data provider of traffic and road network data. This field is
	// case-sensitive. Enter the valid values as shown. For example, entering HERE
	// returns an error. Valid values include:
	//   - Esri – For additional information about Esri (https://docs.aws.amazon.com/location/latest/developerguide/esri.html)
	//   's coverage in your region of interest, see Esri details on street networks
	//   and traffic coverage (https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm)
	//   . Route calculators that use Esri as a data source only calculate routes that
	//   are shorter than 400 km.
	//   - Grab – Grab provides routing functionality for Southeast Asia. For
	//   additional information about GrabMaps (https://docs.aws.amazon.com/location/latest/developerguide/grab.html)
	//   ' coverage, see GrabMaps countries and areas covered (https://docs.aws.amazon.com/location/latest/developerguide/grab.html#grab-coverage-area)
	//   .
	//   - Here – For additional information about HERE Technologies (https://docs.aws.amazon.com/location/latest/developerguide/HERE.html)
	//   ' coverage in your region of interest, see HERE car routing coverage (https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html)
	//   and HERE truck routing coverage (https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html)
	//   .
	// For additional information , see Data providers (https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html)
	// on the Amazon Location Service Developer Guide.
	//
	// This member is required.
	DataSource *string

	// The optional description for the route calculator resource.
	Description *string

	// No longer used. If included, the only allowed value is RequestBasedUsage .
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	// Applies one or more tags to the route calculator resource. A tag is a key-value
	// pair helps manage, identify, search, and filter your resources by labelling
	// them.
	//   - For example: { "tag1" : "value1" , "tag2" : "value2" }
	// Format: "key" : "value" Restrictions:
	//   - Maximum 50 tags per resource
	//   - Each resource tag must be unique with a maximum of one value.
	//   - Maximum key length: 128 Unicode characters in UTF-8
	//   - Maximum value length: 256 Unicode characters in UTF-8
	//   - Can use alphanumeric characters (A–Z, a–z, 0–9), and the following
	//   characters: + - = . _ : / @.
	//   - Cannot use "aws:" as a prefix for a key.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateRouteCalculatorOutput struct {

	// The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN
	// when you specify a resource across all Amazon Web Services.
	//   - Format example:
	//   arn:aws:geo:region:account-id:route-calculator/ExampleCalculator
	//
	// This member is required.
	CalculatorArn *string

	// The name of the route calculator resource.
	//   - For example, ExampleRouteCalculator .
	//
	// This member is required.
	CalculatorName *string

	// The timestamp when the route calculator resource was created in ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html)
	// format: YYYY-MM-DDThh:mm:ss.sssZ .
	//   - For example, 2020–07-2T12:15:20.000Z+01:00
	//
	// This member is required.
	CreateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRouteCalculatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRouteCalculator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRouteCalculator{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateRouteCalculatorMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateRouteCalculatorResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRouteCalculatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRouteCalculator(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateRouteCalculatorMiddleware struct {
}

func (*endpointPrefix_opCreateRouteCalculatorMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateRouteCalculatorMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.routes." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateRouteCalculatorMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateRouteCalculatorMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRouteCalculator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "CreateRouteCalculator",
	}
}

type opCreateRouteCalculatorResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateRouteCalculatorResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateRouteCalculatorResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "geo"
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
				signingName = "geo"
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
				v4aScheme.SigningName = aws.String("geo")
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

func addCreateRouteCalculatorResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateRouteCalculatorResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
