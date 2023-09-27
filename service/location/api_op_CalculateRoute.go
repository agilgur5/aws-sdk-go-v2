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

// Calculates a route (https://docs.aws.amazon.com/location/latest/developerguide/calculate-route.html)
// given the following required parameters: DeparturePosition and
// DestinationPosition . Requires that you first create a route calculator resource (https://docs.aws.amazon.com/location-routes/latest/APIReference/API_CreateRouteCalculator.html)
// . By default, a request that doesn't specify a departure time uses the best time
// of day to travel with the best traffic conditions when calculating the route.
// Additional options include:
//   - Specifying a departure time (https://docs.aws.amazon.com/location/latest/developerguide/departure-time.html)
//     using either DepartureTime or DepartNow . This calculates a route based on
//     predictive traffic data at the given time. You can't specify both
//     DepartureTime and DepartNow in a single request. Specifying both parameters
//     returns a validation error.
//   - Specifying a travel mode (https://docs.aws.amazon.com/location/latest/developerguide/travel-mode.html)
//     using TravelMode sets the transportation mode used to calculate the routes. This
//     also lets you specify additional route preferences in CarModeOptions if
//     traveling by Car , or TruckModeOptions if traveling by Truck . If you specify
//     walking for the travel mode and your data provider is Esri, the start and
//     destination must be within 40km.
func (c *Client) CalculateRoute(ctx context.Context, params *CalculateRouteInput, optFns ...func(*Options)) (*CalculateRouteOutput, error) {
	if params == nil {
		params = &CalculateRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CalculateRoute", params, optFns, c.addOperationCalculateRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CalculateRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CalculateRouteInput struct {

	// The name of the route calculator resource that you want to use to calculate the
	// route.
	//
	// This member is required.
	CalculatorName *string

	// The start position for the route. Defined in World Geodetic System (WGS 84) (https://earth-info.nga.mil/index.php?dir=wgs84&action=wgs84)
	// format: [longitude, latitude] .
	//   - For example, [-123.115, 49.285]
	// If you specify a departure that's not located on a road, Amazon Location moves
	// the position to the nearest road (https://docs.aws.amazon.com/location/latest/developerguide/snap-to-nearby-road.html)
	// . If Esri is the provider for your route calculator, specifying a route that is
	// longer than 400 km returns a 400 RoutesValidationException error. Valid Values:
	// [-180 to 180,-90 to 90]
	//
	// This member is required.
	DeparturePosition []float64

	// The finish position for the route. Defined in World Geodetic System (WGS 84) (https://earth-info.nga.mil/index.php?dir=wgs84&action=wgs84)
	// format: [longitude, latitude] .
	//   - For example, [-122.339, 47.615]
	// If you specify a destination that's not located on a road, Amazon Location
	// moves the position to the nearest road (https://docs.aws.amazon.com/location/latest/developerguide/snap-to-nearby-road.html)
	// . Valid Values: [-180 to 180,-90 to 90]
	//
	// This member is required.
	DestinationPosition []float64

	// Specifies route preferences when traveling by Car , such as avoiding routes that
	// use ferries or tolls. Requirements: TravelMode must be specified as Car .
	CarModeOptions *types.CalculateRouteCarModeOptions

	// Sets the time of departure as the current time. Uses the current time to
	// calculate a route. Otherwise, the best time of day to travel with the best
	// traffic conditions is used to calculate the route. Default Value: false Valid
	// Values: false | true
	DepartNow *bool

	// Specifies the desired time of departure. Uses the given time to calculate the
	// route. Otherwise, the best time of day to travel with the best traffic
	// conditions is used to calculate the route. Setting a departure time in the past
	// returns a 400 ValidationException error.
	//   - In ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	//   YYYY-MM-DDThh:mm:ss.sssZ . For example, 2020–07-2T12:15:20.000Z+01:00
	DepartureTime *time.Time

	// Set the unit system to specify the distance. Default Value: Kilometers
	DistanceUnit types.DistanceUnit

	// Set to include the geometry details in the result for each path between a pair
	// of positions. Default Value: false Valid Values: false | true
	IncludeLegGeometry *bool

	// The optional API key (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
	// to authorize the request.
	Key *string

	// Specifies the mode of transport when calculating a route. Used in estimating
	// the speed of travel and road compatibility. You can choose Car , Truck , Walking
	// , Bicycle or Motorcycle as options for the TravelMode . Bicycle and Motorcycle
	// are only valid when using Grab as a data provider, and only within Southeast
	// Asia. Truck is not available for Grab. For more details on the using Grab for
	// routing, including areas of coverage, see GrabMaps (https://docs.aws.amazon.com/location/latest/developerguide/grab.html)
	// in the Amazon Location Service Developer Guide. The TravelMode you specify also
	// determines how you specify route preferences:
	//   - If traveling by Car use the CarModeOptions parameter.
	//   - If traveling by Truck use the TruckModeOptions parameter.
	// Default Value: Car
	TravelMode types.TravelMode

	// Specifies route preferences when traveling by Truck , such as avoiding routes
	// that use ferries or tolls, and truck specifications to consider when choosing an
	// optimal road. Requirements: TravelMode must be specified as Truck .
	TruckModeOptions *types.CalculateRouteTruckModeOptions

	// Specifies an ordered list of up to 23 intermediate positions to include along a
	// route between the departure position and destination position.
	//   - For example, from the DeparturePosition [-123.115, 49.285] , the route
	//   follows the order that the waypoint positions are given [[-122.757,
	//   49.0021],[-122.349, 47.620]]
	// If you specify a waypoint position that's not located on a road, Amazon
	// Location moves the position to the nearest road (https://docs.aws.amazon.com/location/latest/developerguide/snap-to-nearby-road.html)
	// . Specifying more than 23 waypoints returns a 400 ValidationException error. If
	// Esri is the provider for your route calculator, specifying a route that is
	// longer than 400 km returns a 400 RoutesValidationException error. Valid Values:
	// [-180 to 180,-90 to 90]
	WaypointPositions [][]float64

	noSmithyDocumentSerde
}

func (*CalculateRouteInput) operationName() string {
	return "CalculateRoute"
}

// Returns the result of the route calculation. Metadata includes legs and route
// summary.
type CalculateRouteOutput struct {

	// Contains details about each path between a pair of positions included along a
	// route such as: StartPosition , EndPosition , Distance , DurationSeconds ,
	// Geometry , and Steps . The number of legs returned corresponds to one fewer than
	// the total number of positions in the request. For example, a route with a
	// departure position and destination position returns one leg with the positions
	// snapped to a nearby road (https://docs.aws.amazon.com/location/latest/developerguide/snap-to-nearby-road.html)
	// :
	//   - The StartPosition is the departure position.
	//   - The EndPosition is the destination position.
	// A route with a waypoint between the departure and destination position returns
	// two legs with the positions snapped to a nearby road:
	//   - Leg 1: The StartPosition is the departure position . The EndPosition is the
	//   waypoint positon.
	//   - Leg 2: The StartPosition is the waypoint position. The EndPosition is the
	//   destination position.
	//
	// This member is required.
	Legs []types.Leg

	// Contains information about the whole route, such as: RouteBBox , DataSource ,
	// Distance , DistanceUnit , and DurationSeconds .
	//
	// This member is required.
	Summary *types.CalculateRouteSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCalculateRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCalculateRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCalculateRoute{}, middleware.After)
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
	if err = addEndpointPrefix_opCalculateRouteMiddleware(stack); err != nil {
		return err
	}
	if err = addCalculateRouteResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCalculateRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCalculateRoute(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCalculateRouteMiddleware struct {
}

func (*endpointPrefix_opCalculateRouteMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCalculateRouteMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "routes." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCalculateRouteMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCalculateRouteMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCalculateRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "CalculateRoute",
	}
}

type opCalculateRouteResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCalculateRouteResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCalculateRouteResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addCalculateRouteResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCalculateRouteResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
