// Code generated by smithy-go-codegen DO NOT EDIT.

package scheduler

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified schedule. When you call UpdateSchedule , EventBridge
// Scheduler uses all values, including empty values, specified in the request and
// overrides the existing schedule. This is by design. This means that if you do
// not set an optional field in your request, that field will be set to its
// system-default value after the update. Before calling this operation, we
// recommend that you call the GetSchedule API operation and make a note of all
// optional parameters for your UpdateSchedule call.
func (c *Client) UpdateSchedule(ctx context.Context, params *UpdateScheduleInput, optFns ...func(*Options)) (*UpdateScheduleOutput, error) {
	if params == nil {
		params = &UpdateScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSchedule", params, optFns, c.addOperationUpdateScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateScheduleInput struct {

	// Allows you to configure a time window during which EventBridge Scheduler
	// invokes the schedule.
	//
	// This member is required.
	FlexibleTimeWindow *types.FlexibleTimeWindow

	// The name of the schedule that you are updating.
	//
	// This member is required.
	Name *string

	// The expression that defines when the schedule runs. The following formats are
	// supported.
	//   - at expression - at(yyyy-mm-ddThh:mm:ss)
	//   - rate expression - rate(value unit)
	//   - cron expression - cron(fields)
	// You can use at expressions to create one-time schedules that invoke a target
	// once, at the time and in the time zone, that you specify. You can use rate and
	// cron expressions to create recurring schedules. Rate-based schedules are useful
	// when you want to invoke a target at regular intervals, such as every 15 minutes
	// or every five days. Cron-based schedules are useful when you want to invoke a
	// target periodically at a specific time, such as at 8:00 am (UTC+0) every 1st day
	// of the month. A cron expression consists of six fields separated by white
	// spaces: (minutes hours day_of_month month day_of_week year) . A rate expression
	// consists of a value as a positive integer, and a unit with the following
	// options: minute | minutes | hour | hours | day | days For more information and
	// examples, see Schedule types on EventBridge Scheduler (https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html)
	// in the EventBridge Scheduler User Guide.
	//
	// This member is required.
	ScheduleExpression *string

	// The schedule target. You can use this operation to change the target that your
	// schedule invokes.
	//
	// This member is required.
	Target *types.Target

	// Specifies the action that EventBridge Scheduler applies to the schedule after
	// the schedule completes invoking the target.
	ActionAfterCompletion types.ActionAfterCompletion

	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. If you do not specify a client token, EventBridge Scheduler uses a
	// randomly generated token for the request to ensure idempotency.
	ClientToken *string

	// The description you specify for the schedule.
	Description *string

	// The date, in UTC, before which the schedule can invoke its target. Depending on
	// the schedule's recurrence expression, invocations might stop on, or before, the
	// EndDate you specify. EventBridge Scheduler ignores EndDate for one-time
	// schedules.
	EndDate *time.Time

	// The name of the schedule group with which the schedule is associated. You must
	// provide this value in order for EventBridge Scheduler to find the schedule you
	// want to update. If you omit this value, EventBridge Scheduler assumes the group
	// is associated to the default group.
	GroupName *string

	// The ARN for the customer managed KMS key that that you want EventBridge
	// Scheduler to use to encrypt and decrypt your data.
	KmsKeyArn *string

	// The timezone in which the scheduling expression is evaluated.
	ScheduleExpressionTimezone *string

	// The date, in UTC, after which the schedule can begin invoking its target.
	// Depending on the schedule's recurrence expression, invocations might occur on,
	// or after, the StartDate you specify. EventBridge Scheduler ignores StartDate
	// for one-time schedules.
	StartDate *time.Time

	// Specifies whether the schedule is enabled or disabled.
	State types.ScheduleState

	noSmithyDocumentSerde
}

func (*UpdateScheduleInput) operationName() string {
	return "UpdateSchedule"
}

type UpdateScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the schedule that you updated.
	//
	// This member is required.
	ScheduleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSchedule{}, middleware.After)
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
	if err = addUpdateScheduleResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opUpdateScheduleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSchedule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateSchedule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateSchedule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateSchedule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateScheduleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateScheduleInput ")
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
func addIdempotencyToken_opUpdateScheduleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateSchedule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "scheduler",
		OperationName: "UpdateSchedule",
	}
}

type opUpdateScheduleResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateScheduleResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateScheduleResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "scheduler"
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
				signingName = "scheduler"
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
				v4aScheme.SigningName = aws.String("scheduler")
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

func addUpdateScheduleResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateScheduleResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
