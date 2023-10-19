// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a Cedar policy and saves it in the specified policy store. You can
// create either a static policy or a policy linked to a policy template.
//   - To create a static policy, provide the Cedar policy text in the StaticPolicy
//     section of the PolicyDefinition .
//   - To create a policy that is dynamically linked to a policy template, specify
//     the policy template ID and the principal and resource to associate with this
//     policy in the templateLinked section of the PolicyDefinition . If the policy
//     template is ever updated, any policies linked to the policy template
//     automatically use the updated template.
//
// Creating a policy causes it to be validated against the schema in the policy
// store. If the policy doesn't pass validation, the operation fails and the policy
// isn't stored. Verified Permissions is eventually consistent (https://wikipedia.org/wiki/Eventual_consistency)
// . It can take a few seconds for a new or changed element to be propagate through
// the service and be visible in the results of other Verified Permissions
// operations.
func (c *Client) CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) {
	if params == nil {
		params = &CreatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicy", params, optFns, c.addOperationCreatePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePolicyInput struct {

	// A structure that specifies the policy type and content to use for the new
	// policy. You must include either a static or a templateLinked element. The policy
	// content must be written in the Cedar policy language.
	//
	// This member is required.
	Definition types.PolicyDefinition

	// Specifies the PolicyStoreId of the policy store you want to store the policy in.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies a unique, case-sensitive ID that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value. (https://wikipedia.org/wiki/Universally_unique_identifier)
	// . If you don't provide this value, then Amazon Web Services generates a random
	// one for you. If you retry the operation with the same ClientToken , but with
	// different parameters, the retry fails with an IdempotentParameterMismatch error.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreatePolicyOutput struct {

	// The date and time the policy was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time the policy was last updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The unique ID of the new policy.
	//
	// This member is required.
	PolicyId *string

	// The ID of the policy store that contains the new policy.
	//
	// This member is required.
	PolicyStoreId *string

	// The policy type of the new policy.
	//
	// This member is required.
	PolicyType types.PolicyType

	// The principal specified in the new policy's scope. This response element isn't
	// present when principal isn't specified in the policy content.
	Principal *types.EntityIdentifier

	// The resource specified in the new policy's scope. This response element isn't
	// present when the resource isn't specified in the policy content.
	Resource *types.EntityIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreatePolicy{}, middleware.After)
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
	if err = addIdempotencyToken_opCreatePolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicy(options.Region), middleware.Before); err != nil {
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
		operation: "CreatePolicy",
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

type idempotencyToken_initializeOpCreatePolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePolicyInput ")
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
func addIdempotencyToken_opCreatePolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "verifiedpermissions",
		OperationName: "CreatePolicy",
	}
}
