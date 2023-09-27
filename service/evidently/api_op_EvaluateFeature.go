// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation assigns a feature variation to one given user session. You pass
// in an entityID that represents the user. Evidently then checks the evaluation
// rules and assigns the variation. The first rules that are evaluated are the
// override rules. If the user's entityID matches an override rule, the user is
// served the variation specified by that rule. If there is a current launch with
// this feature that uses segment overrides, and if the user session's
// evaluationContext matches a segment rule defined in a segment override, the
// configuration in the segment overrides is used. For more information about
// segments, see CreateSegment (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_CreateSegment.html)
// and Use segments to focus your audience (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html)
// . If there is a launch with no segment overrides, the user might be assigned to
// a variation in the launch. The chance of this depends on the percentage of users
// that are allocated to that launch. If the user is enrolled in the launch, the
// variation they are served depends on the allocation of the various feature
// variations used for the launch. If the user is not assigned to a launch, and
// there is an ongoing experiment for this feature, the user might be assigned to a
// variation in the experiment. The chance of this depends on the percentage of
// users that are allocated to that experiment. If the experiment uses a segment,
// then only user sessions with evaluationContext values that match the segment
// rule are used in the experiment. If the user is enrolled in the experiment, the
// variation they are served depends on the allocation of the various feature
// variations used for the experiment. If the user is not assigned to a launch or
// experiment, they are served the default variation.
func (c *Client) EvaluateFeature(ctx context.Context, params *EvaluateFeatureInput, optFns ...func(*Options)) (*EvaluateFeatureOutput, error) {
	if params == nil {
		params = &EvaluateFeatureInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EvaluateFeature", params, optFns, c.addOperationEvaluateFeatureMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EvaluateFeatureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EvaluateFeatureInput struct {

	// An internal ID that represents a unique user of the application. This entityID
	// is checked against any override rules assigned for this feature.
	//
	// This member is required.
	EntityId *string

	// The name of the feature being evaluated.
	//
	// This member is required.
	Feature *string

	// The name or ARN of the project that contains this feature.
	//
	// This member is required.
	Project *string

	// A JSON object of attributes that you can optionally pass in as part of the
	// evaluation event sent to Evidently from the user session. Evidently can use this
	// value to match user sessions with defined audience segments. For more
	// information, see Use segments to focus your audience (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html)
	// . If you include this parameter, the value must be a JSON object. A JSON array
	// is not supported.
	//
	// This value conforms to the media type: application/json
	EvaluationContext *string

	noSmithyDocumentSerde
}

func (*EvaluateFeatureInput) operationName() string {
	return "EvaluateFeature"
}

type EvaluateFeatureOutput struct {

	// If this user was assigned to a launch or experiment, this field lists the
	// launch or experiment name.
	//
	// This value conforms to the media type: application/json
	Details *string

	// Specifies the reason that the user session was assigned this variation.
	// Possible values include DEFAULT , meaning the user was served the default
	// variation; LAUNCH_RULE_MATCH , if the user session was enrolled in a launch;
	// EXPERIMENT_RULE_MATCH , if the user session was enrolled in an experiment; or
	// ENTITY_OVERRIDES_MATCH , if the user's entityId matches an override rule.
	Reason *string

	// The value assigned to this variation to differentiate it from the other
	// variations of this feature.
	Value types.VariableValue

	// The name of the variation that was served to the user session.
	Variation *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEvaluateFeatureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEvaluateFeature{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEvaluateFeature{}, middleware.After)
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
	if err = addEndpointPrefix_opEvaluateFeatureMiddleware(stack); err != nil {
		return err
	}
	if err = addEvaluateFeatureResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpEvaluateFeatureValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEvaluateFeature(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opEvaluateFeatureMiddleware struct {
}

func (*endpointPrefix_opEvaluateFeatureMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opEvaluateFeatureMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "dataplane." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opEvaluateFeatureMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opEvaluateFeatureMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opEvaluateFeature(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "EvaluateFeature",
	}
}

type opEvaluateFeatureResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opEvaluateFeatureResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opEvaluateFeatureResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "evidently"
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
				signingName = "evidently"
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
				v4aScheme.SigningName = aws.String("evidently")
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

func addEvaluateFeatureResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opEvaluateFeatureResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
