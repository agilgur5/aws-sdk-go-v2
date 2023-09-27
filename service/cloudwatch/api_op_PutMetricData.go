// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Publishes metric data points to Amazon CloudWatch. CloudWatch associates the
// data points with the specified metric. If the specified metric does not exist,
// CloudWatch creates the metric. When CloudWatch creates a metric, it can take up
// to fifteen minutes for the metric to appear in calls to ListMetrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html)
// . You can publish either individual data points in the Value field, or arrays
// of values and the number of times each value occurred during the period by using
// the Values and Counts fields in the MetricDatum structure. Using the Values and
// Counts method enables you to publish up to 150 values per metric with one
// PutMetricData request, and supports retrieving percentile statistics on this
// data. Each PutMetricData request is limited to 1 MB in size for HTTP POST
// requests. You can send a payload compressed by gzip. Each request is also
// limited to no more than 1000 different metrics. Although the Value parameter
// accepts numbers of type Double , CloudWatch rejects values that are either too
// small or too large. Values must be in the range of -2^360 to 2^360. In addition,
// special values (for example, NaN, +Infinity, -Infinity) are not supported. You
// can use up to 30 dimensions per metric to further clarify what data the metric
// collects. Each dimension consists of a Name and Value pair. For more information
// about specifying dimensions, see Publishing Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)
// in the Amazon CloudWatch User Guide. You specify the time stamp to be associated
// with each data point. You can specify time stamps that are as much as two weeks
// before the current date, and as much as 2 hours after the current day and time.
// Data points with time stamps from 24 hours ago or longer can take at least 48
// hours to become available for GetMetricData (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html)
// from the time they are submitted. Data points with time stamps between 3 and 24
// hours ago can take as much as 2 hours to become available for for GetMetricData (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html)
// . CloudWatch needs raw data points to calculate percentile statistics. If you
// publish data using a statistic set instead, you can only retrieve percentile
// statistics for this data if one of the following conditions is true:
//   - The SampleCount value of the statistic set is 1 and Min , Max , and Sum are
//     all equal.
//   - The Min and Max are equal, and Sum is equal to Min multiplied by SampleCount
//     .
func (c *Client) PutMetricData(ctx context.Context, params *PutMetricDataInput, optFns ...func(*Options)) (*PutMetricDataOutput, error) {
	if params == nil {
		params = &PutMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMetricData", params, optFns, c.addOperationPutMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricDataInput struct {

	// The data for the metric. The array can include no more than 1000 metrics per
	// call.
	//
	// This member is required.
	MetricData []types.MetricDatum

	// The namespace for the metric data. You can use ASCII characters for the
	// namespace, except for control characters which are not supported. To avoid
	// conflicts with Amazon Web Services service namespaces, you should not specify a
	// namespace that begins with AWS/
	//
	// This member is required.
	Namespace *string

	noSmithyDocumentSerde
}

func (*PutMetricDataInput) operationName() string {
	return "PutMetricData"
}

type PutMetricDataOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutMetricData{}, middleware.After)
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
	if err = addPutMetricDataResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "PutMetricData",
	}
}

type opPutMetricDataResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutMetricDataResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutMetricDataResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "monitoring"
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
				signingName = "monitoring"
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
				v4aScheme.SigningName = aws.String("monitoring")
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

func addPutMetricDataResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutMetricDataResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
