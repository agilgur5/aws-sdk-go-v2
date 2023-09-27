// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a usage limit for a specified Amazon Redshift feature on a cluster. The
// usage limit is identified by the returned usage limit identifier.
func (c *Client) CreateUsageLimit(ctx context.Context, params *CreateUsageLimitInput, optFns ...func(*Options)) (*CreateUsageLimitOutput, error) {
	if params == nil {
		params = &CreateUsageLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUsageLimit", params, optFns, c.addOperationCreateUsageLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUsageLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUsageLimitInput struct {

	// The limit amount. If time-based, this amount is in minutes. If data-based, this
	// amount is in terabytes (TB). The value must be a positive number.
	//
	// This member is required.
	Amount int64

	// The identifier of the cluster that you want to limit usage.
	//
	// This member is required.
	ClusterIdentifier *string

	// The Amazon Redshift feature that you want to limit.
	//
	// This member is required.
	FeatureType types.UsageLimitFeatureType

	// The type of limit. Depending on the feature type, this can be based on a time
	// duration or data size. If FeatureType is spectrum , then LimitType must be
	// data-scanned . If FeatureType is concurrency-scaling , then LimitType must be
	// time . If FeatureType is cross-region-datasharing , then LimitType must be
	// data-scanned .
	//
	// This member is required.
	LimitType types.UsageLimitLimitType

	// The action that Amazon Redshift takes when the limit is reached. The default is
	// log. For more information about this parameter, see UsageLimit .
	BreachAction types.UsageLimitBreachAction

	// The time period that the amount applies to. A weekly period begins on Sunday.
	// The default is monthly .
	Period types.UsageLimitPeriod

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

func (*CreateUsageLimitInput) operationName() string {
	return "CreateUsageLimit"
}

// Describes a usage limit object for a cluster.
type CreateUsageLimitOutput struct {

	// The limit amount. If time-based, this amount is in minutes. If data-based, this
	// amount is in terabytes (TB).
	Amount int64

	// The action that Amazon Redshift takes when the limit is reached. Possible
	// values are:
	//   - log - To log an event in a system table. The default is log.
	//   - emit-metric - To emit CloudWatch metrics.
	//   - disable - To disable the feature until the next usage period begins.
	BreachAction types.UsageLimitBreachAction

	// The identifier of the cluster with a usage limit.
	ClusterIdentifier *string

	// The Amazon Redshift feature to which the limit applies.
	FeatureType types.UsageLimitFeatureType

	// The type of limit. Depending on the feature type, this can be based on a time
	// duration or data size.
	LimitType types.UsageLimitLimitType

	// The time period that the amount applies to. A weekly period begins on Sunday.
	// The default is monthly .
	Period types.UsageLimitPeriod

	// A list of tag instances.
	Tags []types.Tag

	// The identifier of the usage limit.
	UsageLimitId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUsageLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateUsageLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateUsageLimit{}, middleware.After)
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
	if err = addCreateUsageLimitResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateUsageLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUsageLimit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUsageLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateUsageLimit",
	}
}

type opCreateUsageLimitResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateUsageLimitResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateUsageLimitResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "redshift"
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
				signingName = "redshift"
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
				v4aScheme.SigningName = aws.String("redshift")
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

func addCreateUsageLimitResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateUsageLimitResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
