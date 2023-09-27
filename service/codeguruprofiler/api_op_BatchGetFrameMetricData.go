// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the time series of values for a requested list of frame metrics from a
// time period.
func (c *Client) BatchGetFrameMetricData(ctx context.Context, params *BatchGetFrameMetricDataInput, optFns ...func(*Options)) (*BatchGetFrameMetricDataOutput, error) {
	if params == nil {
		params = &BatchGetFrameMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetFrameMetricData", params, optFns, c.addOperationBatchGetFrameMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetFrameMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the BatchGetFrameMetricDataRequest.
type BatchGetFrameMetricDataInput struct {

	// The name of the profiling group associated with the the frame metrics used to
	// return the time series values.
	//
	// This member is required.
	ProfilingGroupName *string

	// The end time of the time period for the returned time series values. This is
	// specified using the ISO 8601 format. For example, 2020-06-01T13:15:02.001Z
	// represents 1 millisecond past June 1, 2020 1:15:02 PM UTC.
	EndTime *time.Time

	// The details of the metrics that are used to request a time series of values.
	// The metric includes the name of the frame, the aggregation type to calculate the
	// metric value for the frame, and the thread states to use to get the count for
	// the metric value of the frame.
	FrameMetrics []types.FrameMetric

	// The duration of the frame metrics used to return the time series values.
	// Specify using the ISO 8601 format. The maximum period duration is one day ( PT24H
	// or P1D ).
	Period *string

	// The start time of the time period for the frame metrics used to return the time
	// series values. This is specified using the ISO 8601 format. For example,
	// 2020-06-01T13:15:02.001Z represents 1 millisecond past June 1, 2020 1:15:02 PM
	// UTC.
	StartTime *time.Time

	// The requested resolution of time steps for the returned time series of values.
	// If the requested target resolution is not available due to data not being
	// retained we provide a best effort result by falling back to the most granular
	// available resolution after the target resolution. There are 3 valid values.
	//   - P1D — 1 day
	//   - PT1H — 1 hour
	//   - PT5M — 5 minutes
	TargetResolution types.AggregationPeriod

	noSmithyDocumentSerde
}

func (*BatchGetFrameMetricDataInput) operationName() string {
	return "BatchGetFrameMetricData"
}

// The structure representing the BatchGetFrameMetricDataResponse.
type BatchGetFrameMetricDataOutput struct {

	// The end time of the time period for the returned time series values. This is
	// specified using the ISO 8601 format. For example, 2020-06-01T13:15:02.001Z
	// represents 1 millisecond past June 1, 2020 1:15:02 PM UTC.
	//
	// This member is required.
	EndTime *time.Time

	// List of instances, or time steps, in the time series. For example, if the period
	// is one day ( PT24H) ), and the resolution is five minutes ( PT5M ), then there
	// are 288 endTimes in the list that are each five minutes appart.
	//
	// This member is required.
	EndTimes []types.TimestampStructure

	// Details of the metrics to request a time series of values. The metric includes
	// the name of the frame, the aggregation type to calculate the metric value for
	// the frame, and the thread states to use to get the count for the metric value of
	// the frame.
	//
	// This member is required.
	FrameMetricData []types.FrameMetricDatum

	// Resolution or granularity of the profile data used to generate the time series.
	// This is the value used to jump through time steps in a time series. There are 3
	// valid values.
	//   - P1D — 1 day
	//   - PT1H — 1 hour
	//   - PT5M — 5 minutes
	//
	// This member is required.
	Resolution types.AggregationPeriod

	// The start time of the time period for the returned time series values. This is
	// specified using the ISO 8601 format. For example, 2020-06-01T13:15:02.001Z
	// represents 1 millisecond past June 1, 2020 1:15:02 PM UTC.
	//
	// This member is required.
	StartTime *time.Time

	// List of instances which remained unprocessed. This will create a missing time
	// step in the list of end times.
	//
	// This member is required.
	UnprocessedEndTimes map[string][]types.TimestampStructure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetFrameMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetFrameMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetFrameMetricData{}, middleware.After)
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
	if err = addBatchGetFrameMetricDataResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpBatchGetFrameMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetFrameMetricData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetFrameMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "BatchGetFrameMetricData",
	}
}

type opBatchGetFrameMetricDataResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opBatchGetFrameMetricDataResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opBatchGetFrameMetricDataResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "codeguru-profiler"
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
				signingName = "codeguru-profiler"
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
				v4aScheme.SigningName = aws.String("codeguru-profiler")
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

func addBatchGetFrameMetricDataResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opBatchGetFrameMetricDataResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
