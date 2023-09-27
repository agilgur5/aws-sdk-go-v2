// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an existing state machine by modifying its definition , roleArn , or
// loggingConfiguration . Running executions will continue to use the previous
// definition and roleArn . You must include at least one of definition or roleArn
// or you will receive a MissingRequiredParameter error. A qualified state machine
// ARN refers to a Distributed Map state defined within a state machine. For
// example, the qualified state machine ARN
// arn:partition:states:region:account-id:stateMachine:stateMachineName/mapStateLabel
// refers to a Distributed Map state with a label mapStateLabel in the state
// machine named stateMachineName . A qualified state machine ARN can either refer
// to a Distributed Map state defined within a state machine, a version ARN, or an
// alias ARN. The following are some examples of qualified and unqualified state
// machine ARNs:
//   - The following qualified state machine ARN refers to a Distributed Map state
//     with a label mapStateLabel in a state machine named myStateMachine .
//     arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//     If you provide a qualified state machine ARN that refers to a Distributed Map
//     state, the request fails with ValidationException .
//   - The following qualified state machine ARN refers to an alias named PROD .
//     arn::states:::stateMachine: If you provide a qualified state machine ARN that
//     refers to a version ARN or an alias ARN, the request starts execution for that
//     version or alias.
//   - The following unqualified state machine ARN refers to a state machine named
//     myStateMachine . arn::states:::stateMachine:
//
// After you update your state machine, you can set the publish parameter to true
// in the same action to publish a new version (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html)
// . This way, you can opt-in to strict versioning of your state machine. Step
// Functions assigns monotonically increasing integers for state machine versions,
// starting at version number 1. All StartExecution calls within a few seconds use
// the updated definition and roleArn . Executions started immediately after you
// call UpdateStateMachine may use the previous state machine definition and
// roleArn .
func (c *Client) UpdateStateMachine(ctx context.Context, params *UpdateStateMachineInput, optFns ...func(*Options)) (*UpdateStateMachineOutput, error) {
	if params == nil {
		params = &UpdateStateMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStateMachine", params, optFns, c.addOperationUpdateStateMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStateMachineInput struct {

	// The Amazon Resource Name (ARN) of the state machine.
	//
	// This member is required.
	StateMachineArn *string

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html)
	// .
	Definition *string

	// Use the LoggingConfiguration data type to set CloudWatch Logs options.
	LoggingConfiguration *types.LoggingConfiguration

	// Specifies whether the state machine version is published. The default is false .
	// To publish a version after updating the state machine, set publish to true .
	Publish bool

	// The Amazon Resource Name (ARN) of the IAM role of the state machine.
	RoleArn *string

	// Selects whether X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration

	// An optional description of the state machine version to publish. You can only
	// specify the versionDescription parameter if you've set publish to true .
	VersionDescription *string

	noSmithyDocumentSerde
}

func (*UpdateStateMachineInput) operationName() string {
	return "UpdateStateMachine"
}

type UpdateStateMachineOutput struct {

	// The date and time the state machine was updated.
	//
	// This member is required.
	UpdateDate *time.Time

	// The revision identifier for the updated state machine.
	RevisionId *string

	// The Amazon Resource Name (ARN) of the published state machine version. If the
	// publish parameter isn't set to true , this field returns null.
	StateMachineVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStateMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateStateMachine{}, middleware.After)
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
	if err = addUpdateStateMachineResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateStateMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStateMachine(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStateMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "UpdateStateMachine",
	}
}

type opUpdateStateMachineResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateStateMachineResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateStateMachineResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "states"
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
				signingName = "states"
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
				v4aScheme.SigningName = aws.String("states")
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

func addUpdateStateMachineResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateStateMachineResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
