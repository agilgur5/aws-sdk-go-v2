// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a monitor in Amazon CloudWatch Internet Monitor. A monitor is built
// based on information from the application resources that you add: VPCs, Network
// Load Balancers (NLBs), Amazon CloudFront distributions, and Amazon WorkSpaces
// directories. Internet Monitor then publishes internet measurements from Amazon
// Web Services that are specific to the city-networks. That is, the locations and
// ASNs (typically internet service providers or ISPs), where clients access your
// application. For more information, see Using Amazon CloudWatch Internet Monitor (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-InternetMonitor.html)
// in the Amazon CloudWatch User Guide. When you create a monitor, you choose the
// percentage of traffic that you want to monitor. You can also set a maximum limit
// for the number of city-networks where client traffic is monitored, that caps the
// total traffic that Internet Monitor monitors. A city-network maximum is the
// limit of city-networks, but you only pay for the number of city-networks that
// are actually monitored. You can update your monitor at any time to change the
// percentage of traffic to monitor or the city-networks maximum. For more
// information, see Choosing a city-network maximum value (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html)
// in the Amazon CloudWatch User Guide.
func (c *Client) CreateMonitor(ctx context.Context, params *CreateMonitorInput, optFns ...func(*Options)) (*CreateMonitorOutput, error) {
	if params == nil {
		params = &CreateMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMonitor", params, optFns, c.addOperationCreateMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMonitorInput struct {

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	// A unique, case-sensitive string of up to 64 ASCII characters that you specify
	// to make an idempotent API request. Don't reuse the same client token for other
	// API requests.
	ClientToken *string

	// Defines the threshold percentages and other configuration information for when
	// Amazon CloudWatch Internet Monitor creates a health event. Internet Monitor
	// creates a health event when an internet issue that affects your application end
	// users has a health score percentage that is at or below a specific threshold,
	// and, sometimes, when other criteria are met. If you don't set a health event
	// threshold, the default value is 95%. For more information, see Change health
	// event thresholds (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-overview.html#IMUpdateThresholdFromOverview)
	// in the Internet Monitor section of the CloudWatch User Guide.
	HealthEventsConfig *types.HealthEventsConfig

	// Publish internet measurements for Internet Monitor to an Amazon S3 bucket in
	// addition to CloudWatch Logs.
	InternetMeasurementsLogDelivery *types.InternetMeasurementsLogDelivery

	// The maximum number of city-networks to monitor for your resources. A
	// city-network is the location (city) where clients access your application
	// resources from and the ASN or network provider, such as an internet service
	// provider (ISP), that clients access the resources through. Setting this limit
	// can help control billing costs. To learn more, see Choosing a city-network
	// maximum value  (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	MaxCityNetworksToMonitor int32

	// The resources to include in a monitor, which you provide as a set of Amazon
	// Resource Names (ARNs). Resources can be VPCs, NLBs, Amazon CloudFront
	// distributions, or Amazon WorkSpaces directories. You can add a combination of
	// VPCs and CloudFront distributions, or you can add WorkSpaces directories, or you
	// can add NLBs. You can't add NLBs or WorkSpaces directories together with any
	// other resources. If you add only Amazon VPC resources, at least one VPC must
	// have an Internet Gateway attached to it, to make sure that it has internet
	// connectivity.
	Resources []string

	// The tags for a monitor. You can add a maximum of 50 tags in Internet Monitor.
	Tags map[string]string

	// The percentage of the internet-facing traffic for your application that you
	// want to monitor with this monitor. If you set a city-networks maximum, that
	// limit overrides the traffic percentage that you set. To learn more, see
	// Choosing an application traffic percentage to monitor  (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMTrafficPercentage.html)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	TrafficPercentageToMonitor int32

	noSmithyDocumentSerde
}

func (*CreateMonitorInput) operationName() string {
	return "CreateMonitor"
}

type CreateMonitorOutput struct {

	// The Amazon Resource Name (ARN) of the monitor.
	//
	// This member is required.
	Arn *string

	// The status of a monitor.
	//
	// This member is required.
	Status types.MonitorConfigState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMonitor{}, middleware.After)
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
	if err = addCreateMonitorResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateMonitorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMonitor(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMonitor struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMonitor) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMonitorInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMonitorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMonitor{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "internetmonitor",
		OperationName: "CreateMonitor",
	}
}

type opCreateMonitorResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateMonitorResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateMonitorResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "internetmonitor"
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
				signingName = "internetmonitor"
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
				v4aScheme.SigningName = aws.String("internetmonitor")
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

func addCreateMonitorResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateMonitorResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
