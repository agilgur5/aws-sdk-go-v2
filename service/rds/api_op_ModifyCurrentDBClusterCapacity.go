// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set the capacity of an Aurora Serverless v1 DB cluster to a specific value.
// Aurora Serverless v1 scales seamlessly based on the workload on the DB cluster.
// In some cases, the capacity might not scale fast enough to meet a sudden change
// in workload, such as a large number of new transactions. Call
// ModifyCurrentDBClusterCapacity to set the capacity explicitly. After this call
// sets the DB cluster capacity, Aurora Serverless v1 can automatically scale the
// DB cluster based on the cooldown period for scaling up and the cooldown period
// for scaling down. For more information about Aurora Serverless v1, see Using
// Amazon Aurora Serverless v1 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html)
// in the Amazon Aurora User Guide. If you call ModifyCurrentDBClusterCapacity
// with the default TimeoutAction , connections that prevent Aurora Serverless v1
// from finding a scaling point might be dropped. For more information about
// scaling points, see Autoscaling for Aurora Serverless v1 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling)
// in the Amazon Aurora User Guide. This action only applies to Aurora Serverless
// v1 DB clusters.
func (c *Client) ModifyCurrentDBClusterCapacity(ctx context.Context, params *ModifyCurrentDBClusterCapacityInput, optFns ...func(*Options)) (*ModifyCurrentDBClusterCapacityOutput, error) {
	if params == nil {
		params = &ModifyCurrentDBClusterCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCurrentDBClusterCapacity", params, optFns, c.addOperationModifyCurrentDBClusterCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCurrentDBClusterCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCurrentDBClusterCapacityInput struct {

	// The DB cluster identifier for the cluster being modified. This parameter isn't
	// case-sensitive. Constraints:
	//   - Must match the identifier of an existing DB cluster.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The DB cluster capacity. When you change the capacity of a paused Aurora
	// Serverless v1 DB cluster, it automatically resumes. Constraints:
	//   - For Aurora MySQL, valid capacity values are 1 , 2 , 4 , 8 , 16 , 32 , 64 ,
	//   128 , and 256 .
	//   - For Aurora PostgreSQL, valid capacity values are 2 , 4 , 8 , 16 , 32 , 64 ,
	//   192 , and 384 .
	Capacity *int32

	// The amount of time, in seconds, that Aurora Serverless v1 tries to find a
	// scaling point to perform seamless scaling before enforcing the timeout action.
	// The default is 300. Specify a value between 10 and 600 seconds.
	SecondsBeforeTimeout *int32

	// The action to take when the timeout is reached, either ForceApplyCapacityChange
	// or RollbackCapacityChange . ForceApplyCapacityChange , the default, sets the
	// capacity to the specified value as soon as possible. RollbackCapacityChange
	// ignores the capacity change if a scaling point isn't found in the timeout
	// period.
	TimeoutAction *string

	noSmithyDocumentSerde
}

func (*ModifyCurrentDBClusterCapacityInput) operationName() string {
	return "ModifyCurrentDBClusterCapacity"
}

type ModifyCurrentDBClusterCapacityOutput struct {

	// The current capacity of the DB cluster.
	CurrentCapacity *int32

	// A user-supplied DB cluster identifier. This identifier is the unique key that
	// identifies a DB cluster.
	DBClusterIdentifier *string

	// A value that specifies the capacity that the DB cluster scales to next.
	PendingCapacity *int32

	// The number of seconds before a call to ModifyCurrentDBClusterCapacity times out.
	SecondsBeforeTimeout *int32

	// The timeout action of a call to ModifyCurrentDBClusterCapacity , either
	// ForceApplyCapacityChange or RollbackCapacityChange .
	TimeoutAction *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyCurrentDBClusterCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCurrentDBClusterCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCurrentDBClusterCapacity{}, middleware.After)
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
	if err = addModifyCurrentDBClusterCapacityResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpModifyCurrentDBClusterCapacityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCurrentDBClusterCapacity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCurrentDBClusterCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyCurrentDBClusterCapacity",
	}
}

type opModifyCurrentDBClusterCapacityResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opModifyCurrentDBClusterCapacityResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opModifyCurrentDBClusterCapacityResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "rds"
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
				signingName = "rds"
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
				v4aScheme.SigningName = aws.String("rds")
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

func addModifyCurrentDBClusterCapacityResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opModifyCurrentDBClusterCapacityResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
