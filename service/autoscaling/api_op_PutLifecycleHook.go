// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

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

// Creates or updates a lifecycle hook for the specified Auto Scaling group.
// Lifecycle hooks let you create solutions that are aware of events in the Auto
// Scaling instance lifecycle, and then perform a custom action on instances when
// the corresponding lifecycle event occurs. This step is a part of the procedure
// for adding a lifecycle hook to an Auto Scaling group:
//   - (Optional) Create a launch template or launch configuration with a user
//     data script that runs while an instance is in a wait state due to a lifecycle
//     hook.
//   - (Optional) Create a Lambda function and a rule that allows Amazon
//     EventBridge to invoke your Lambda function when an instance is put into a wait
//     state due to a lifecycle hook.
//   - (Optional) Create a notification target and an IAM role. The target can be
//     either an Amazon SQS queue or an Amazon SNS topic. The role allows Amazon EC2
//     Auto Scaling to publish lifecycle notifications to the target.
//   - Create the lifecycle hook. Specify whether the hook is used when the
//     instances launch or terminate.
//   - If you need more time, record the lifecycle action heartbeat to keep the
//     instance in a wait state using the RecordLifecycleActionHeartbeat API call.
//   - If you finish before the timeout period ends, send a callback by using the
//     CompleteLifecycleAction API call.
//
// For more information, see Amazon EC2 Auto Scaling lifecycle hooks (https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html)
// in the Amazon EC2 Auto Scaling User Guide. If you exceed your maximum limit of
// lifecycle hooks, which by default is 50 per Auto Scaling group, the call fails.
// You can view the lifecycle hooks for an Auto Scaling group using the
// DescribeLifecycleHooks API call. If you are no longer using a lifecycle hook,
// you can delete it by calling the DeleteLifecycleHook API.
func (c *Client) PutLifecycleHook(ctx context.Context, params *PutLifecycleHookInput, optFns ...func(*Options)) (*PutLifecycleHookOutput, error) {
	if params == nil {
		params = &PutLifecycleHookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecycleHook", params, optFns, c.addOperationPutLifecycleHookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecycleHookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleHookInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The name of the lifecycle hook.
	//
	// This member is required.
	LifecycleHookName *string

	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses
	// or if an unexpected failure occurs. The default value is ABANDON . Valid values:
	// CONTINUE | ABANDON
	DefaultResult *string

	// The maximum time, in seconds, that can elapse before the lifecycle hook times
	// out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1
	// hour).
	HeartbeatTimeout *int32

	// The lifecycle transition. For Auto Scaling groups, there are two major
	// lifecycle transitions.
	//   - To create a lifecycle hook for scale-out events, specify
	//   autoscaling:EC2_INSTANCE_LAUNCHING .
	//   - To create a lifecycle hook for scale-in events, specify
	//   autoscaling:EC2_INSTANCE_TERMINATING .
	// Required for new lifecycle hooks, but optional when updating existing hooks.
	LifecycleTransition *string

	// Additional information that you want to include any time Amazon EC2 Auto
	// Scaling sends a message to the notification target.
	NotificationMetadata *string

	// The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto
	// Scaling uses to notify you when an instance is in a wait state for the lifecycle
	// hook. You can specify either an Amazon SNS topic or an Amazon SQS queue. If you
	// specify an empty string, this overrides the current ARN. This operation uses the
	// JSON format when sending notifications to an Amazon SQS queue, and an email
	// key-value pair format when sending notifications to an Amazon SNS topic. When
	// you specify a notification target, Amazon EC2 Auto Scaling sends it a test
	// message. Test messages contain the following additional key-value pair:
	// "Event": "autoscaling:TEST_NOTIFICATION" .
	NotificationTargetARN *string

	// The ARN of the IAM role that allows the Auto Scaling group to publish to the
	// specified notification target. Valid only if the notification target is an
	// Amazon SNS topic or an Amazon SQS queue. Required for new lifecycle hooks, but
	// optional when updating existing hooks.
	RoleARN *string

	noSmithyDocumentSerde
}

func (*PutLifecycleHookInput) operationName() string {
	return "PutLifecycleHook"
}

type PutLifecycleHookOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLifecycleHookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutLifecycleHook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutLifecycleHook{}, middleware.After)
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
	if err = addPutLifecycleHookResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutLifecycleHookValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleHook(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLifecycleHook(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutLifecycleHook",
	}
}

type opPutLifecycleHookResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutLifecycleHookResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutLifecycleHookResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "autoscaling"
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
				signingName = "autoscaling"
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
				v4aScheme.SigningName = aws.String("autoscaling")
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

func addPutLifecycleHookResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutLifecycleHookResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
