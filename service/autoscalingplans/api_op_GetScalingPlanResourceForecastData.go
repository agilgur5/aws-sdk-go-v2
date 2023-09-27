// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the forecast data for a scalable resource. Capacity forecasts are
// represented as predicted values, or data points, that are calculated using
// historical data points from a specified CloudWatch load metric. Data points are
// available for up to 56 days.
func (c *Client) GetScalingPlanResourceForecastData(ctx context.Context, params *GetScalingPlanResourceForecastDataInput, optFns ...func(*Options)) (*GetScalingPlanResourceForecastDataOutput, error) {
	if params == nil {
		params = &GetScalingPlanResourceForecastDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetScalingPlanResourceForecastData", params, optFns, c.addOperationGetScalingPlanResourceForecastDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetScalingPlanResourceForecastDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetScalingPlanResourceForecastDataInput struct {

	// The exclusive end time of the time range for the forecast data to get. The
	// maximum time duration between the start and end time is seven days. Although
	// this parameter can accept a date and time that is more than two days in the
	// future, the availability of forecast data has limits. AWS Auto Scaling only
	// issues forecasts for periods of two days in advance.
	//
	// This member is required.
	EndTime *time.Time

	// The type of forecast data to get.
	//   - LoadForecast : The load metric forecast.
	//   - CapacityForecast : The capacity forecast.
	//   - ScheduledActionMinCapacity : The minimum capacity for each scheduled scaling
	//   action. This data is calculated as the larger of two values: the capacity
	//   forecast or the minimum capacity in the scaling instruction.
	//   - ScheduledActionMaxCapacity : The maximum capacity for each scheduled scaling
	//   action. The calculation used is determined by the predictive scaling maximum
	//   capacity behavior setting in the scaling instruction.
	//
	// This member is required.
	ForecastDataType types.ForecastDataType

	// The ID of the resource. This string consists of a prefix ( autoScalingGroup )
	// followed by the name of a specified Auto Scaling group ( my-asg ). Example:
	// autoScalingGroup/my-asg .
	//
	// This member is required.
	ResourceId *string

	// The scalable dimension for the resource. The only valid value is
	// autoscaling:autoScalingGroup:DesiredCapacity .
	//
	// This member is required.
	ScalableDimension types.ScalableDimension

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan. Currently, the only valid value is 1 .
	//
	// This member is required.
	ScalingPlanVersion *int64

	// The namespace of the AWS service. The only valid value is autoscaling .
	//
	// This member is required.
	ServiceNamespace types.ServiceNamespace

	// The inclusive start time of the time range for the forecast data to get. The
	// date and time can be at most 56 days before the current date and time.
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

func (*GetScalingPlanResourceForecastDataInput) operationName() string {
	return "GetScalingPlanResourceForecastData"
}

type GetScalingPlanResourceForecastDataOutput struct {

	// The data points to return.
	//
	// This member is required.
	Datapoints []types.Datapoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetScalingPlanResourceForecastDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetScalingPlanResourceForecastData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetScalingPlanResourceForecastData{}, middleware.After)
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
	if err = addGetScalingPlanResourceForecastDataResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetScalingPlanResourceForecastDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetScalingPlanResourceForecastData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetScalingPlanResourceForecastData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling-plans",
		OperationName: "GetScalingPlanResourceForecastData",
	}
}

type opGetScalingPlanResourceForecastDataResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetScalingPlanResourceForecastDataResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetScalingPlanResourceForecastDataResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "autoscaling-plans"
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
				signingName = "autoscaling-plans"
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
				v4aScheme.SigningName = aws.String("autoscaling-plans")
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

func addGetScalingPlanResourceForecastDataResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetScalingPlanResourceForecastDataResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
