// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests whether the given impersonation role can impersonate a target user.
func (c *Client) GetImpersonationRoleEffect(ctx context.Context, params *GetImpersonationRoleEffectInput, optFns ...func(*Options)) (*GetImpersonationRoleEffectOutput, error) {
	if params == nil {
		params = &GetImpersonationRoleEffectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetImpersonationRoleEffect", params, optFns, c.addOperationGetImpersonationRoleEffectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetImpersonationRoleEffectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImpersonationRoleEffectInput struct {

	// The impersonation role ID to test.
	//
	// This member is required.
	ImpersonationRoleId *string

	// The WorkMail organization where the impersonation role is defined.
	//
	// This member is required.
	OrganizationId *string

	// The WorkMail organization user chosen to test the impersonation role. The
	// following identity formats are available:
	//   - User ID: 12345678-1234-1234-1234-123456789012 or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//   - Email address: user@domain.tld
	//   - User name: user
	//
	// This member is required.
	TargetUser *string

	noSmithyDocumentSerde
}

type GetImpersonationRoleEffectOutput struct {

	// Effect of the impersonation role on the target user based on its rules.
	// Available effects are ALLOW or DENY .
	Effect types.AccessEffect

	// A list of the rules that match the input and produce the configured effect.
	MatchedRules []types.ImpersonationMatchedRule

	// The impersonation role type.
	Type types.ImpersonationRoleType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetImpersonationRoleEffectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetImpersonationRoleEffect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetImpersonationRoleEffect{}, middleware.After)
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
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetImpersonationRoleEffectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetImpersonationRoleEffect(options.Region), middleware.Before); err != nil {
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
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "GetImpersonationRoleEffect",
		options:   options,
	}, "ResolveEndpointV2", middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&signRequestMiddleware{}, middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&getIdentityMiddleware{
		options: options,
	}, middleware.Before)
	if err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetImpersonationRoleEffect(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "GetImpersonationRoleEffect",
	}
}
