// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a change request for Change Manager. The Automation runbooks specified
// in the change request run only after all required approvals for the change
// request have been received.
func (c *Client) StartChangeRequestExecution(ctx context.Context, params *StartChangeRequestExecutionInput, optFns ...func(*Options)) (*StartChangeRequestExecutionOutput, error) {
	if params == nil {
		params = &StartChangeRequestExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartChangeRequestExecution", params, optFns, c.addOperationStartChangeRequestExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartChangeRequestExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartChangeRequestExecutionInput struct {

	// The name of the change template document to run during the runbook workflow.
	//
	// This member is required.
	DocumentName *string

	// Information about the Automation runbooks that are run during the runbook
	// workflow. The Automation runbooks specified for the runbook workflow can't run
	// until all required approvals for the change request have been received.
	//
	// This member is required.
	Runbooks []types.Runbook

	// Indicates whether the change request can be approved automatically without the
	// need for manual approvals. If AutoApprovable is enabled in a change template,
	// then setting AutoApprove to true in StartChangeRequestExecution creates a
	// change request that bypasses approver review. Change Calendar restrictions are
	// not bypassed in this scenario. If the state of an associated calendar is CLOSED
	// , change freeze approvers must still grant permission for this change request to
	// run. If they don't, the change won't be processed until the calendar state is
	// again OPEN .
	AutoApprove bool

	// User-provided details about the change. If no details are provided, content
	// specified in the Template information section of the associated change template
	// is added.
	ChangeDetails *string

	// The name of the change request associated with the runbook workflow to be run.
	ChangeRequestName *string

	// The user-provided idempotency token. The token must be unique, is case
	// insensitive, enforces the UUID format, and can't be reused.
	ClientToken *string

	// The version of the change template document to run during the runbook workflow.
	DocumentVersion *string

	// A key-value map of parameters that match the declared parameters in the change
	// template document.
	Parameters map[string][]string

	// The time that the requester expects the runbook workflow related to the change
	// request to complete. The time is an estimate only that the requester provides
	// for reviewers.
	ScheduledEndTime *time.Time

	// The date and time specified in the change request to run the Automation
	// runbooks. The Automation runbooks specified for the runbook workflow can't run
	// until all required approvals for the change request have been received.
	ScheduledTime *time.Time

	// Optional metadata that you assign to a resource. You can specify a maximum of
	// five tags for a change request. Tags enable you to categorize a resource in
	// different ways, such as by purpose, owner, or environment. For example, you
	// might want to tag a change request to identify an environment or target Amazon
	// Web Services Region. In this case, you could specify the following key-value
	// pairs:
	//   - Key=Environment,Value=Production
	//   - Key=Region,Value=us-east-2
	Tags []types.Tag

	noSmithyDocumentSerde
}

func (*StartChangeRequestExecutionInput) operationName() string {
	return "StartChangeRequestExecution"
}

type StartChangeRequestExecutionOutput struct {

	// The unique ID of a runbook workflow operation. (A runbook workflow is a type of
	// Automation operation.)
	AutomationExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartChangeRequestExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartChangeRequestExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartChangeRequestExecution{}, middleware.After)
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
	if err = addStartChangeRequestExecutionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartChangeRequestExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartChangeRequestExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartChangeRequestExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "StartChangeRequestExecution",
	}
}

type opStartChangeRequestExecutionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartChangeRequestExecutionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartChangeRequestExecutionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ssm"
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
				signingName = "ssm"
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
				v4aScheme.SigningName = aws.String("ssm")
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

func addStartChangeRequestExecutionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartChangeRequestExecutionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
