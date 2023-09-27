// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new recommendation preference or updates an existing recommendation
// preference, such as enhanced infrastructure metrics. For more information, see
// Activating enhanced infrastructure metrics (https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html)
// in the Compute Optimizer User Guide.
func (c *Client) PutRecommendationPreferences(ctx context.Context, params *PutRecommendationPreferencesInput, optFns ...func(*Options)) (*PutRecommendationPreferencesOutput, error) {
	if params == nil {
		params = &PutRecommendationPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRecommendationPreferences", params, optFns, c.addOperationPutRecommendationPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRecommendationPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRecommendationPreferencesInput struct {

	// The target resource type of the recommendation preference to create. The
	// Ec2Instance option encompasses standalone instances and instances that are part
	// of Auto Scaling groups. The AutoScalingGroup option encompasses only instances
	// that are part of an Auto Scaling group. The valid values for this parameter are
	// Ec2Instance and AutoScalingGroup .
	//
	// This member is required.
	ResourceType types.ResourceType

	// The status of the enhanced infrastructure metrics recommendation preference to
	// create or update. Specify the Active status to activate the preference, or
	// specify Inactive to deactivate the preference. For more information, see
	// Enhanced infrastructure metrics (https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html)
	// in the Compute Optimizer User Guide.
	EnhancedInfrastructureMetrics types.EnhancedInfrastructureMetrics

	// The provider of the external metrics recommendation preference to create or
	// update. Specify a valid provider in the source field to activate the
	// preference. To delete this preference, see the DeleteRecommendationPreferences
	// action. This preference can only be set for the Ec2Instance resource type. For
	// more information, see External metrics ingestion (https://docs.aws.amazon.com/compute-optimizer/latest/ug/external-metrics-ingestion.html)
	// in the Compute Optimizer User Guide.
	ExternalMetricsPreference *types.ExternalMetricsPreference

	// The status of the inferred workload types recommendation preference to create
	// or update. The inferred workload type feature is active by default. To
	// deactivate it, create a recommendation preference. Specify the Inactive status
	// to deactivate the feature, or specify Active to activate it. For more
	// information, see Inferred workload types (https://docs.aws.amazon.com/compute-optimizer/latest/ug/inferred-workload-types.html)
	// in the Compute Optimizer User Guide.
	InferredWorkloadTypes types.InferredWorkloadTypesPreference

	// An object that describes the scope of the recommendation preference to create.
	// You can create recommendation preferences at the organization level (for
	// management accounts of an organization only), account level, and resource level.
	// For more information, see Activating enhanced infrastructure metrics (https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html)
	// in the Compute Optimizer User Guide. You cannot create recommendation
	// preferences for Auto Scaling groups at the organization and account levels. You
	// can create recommendation preferences for Auto Scaling groups only at the
	// resource level by specifying a scope name of ResourceArn and a scope value of
	// the Auto Scaling group Amazon Resource Name (ARN). This will configure the
	// preference for all instances that are part of the specified Auto Scaling group.
	// You also cannot create recommendation preferences at the resource level for
	// instances that are part of an Auto Scaling group. You can create recommendation
	// preferences at the resource level only for standalone instances.
	Scope *types.Scope

	noSmithyDocumentSerde
}

func (*PutRecommendationPreferencesInput) operationName() string {
	return "PutRecommendationPreferences"
}

type PutRecommendationPreferencesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRecommendationPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutRecommendationPreferences{}, middleware.After)
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
	if err = addPutRecommendationPreferencesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutRecommendationPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecommendationPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRecommendationPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "PutRecommendationPreferences",
	}
}

type opPutRecommendationPreferencesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutRecommendationPreferencesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutRecommendationPreferencesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "compute-optimizer"
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
				signingName = "compute-optimizer"
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
				v4aScheme.SigningName = aws.String("compute-optimizer")
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

func addPutRecommendationPreferencesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutRecommendationPreferencesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
