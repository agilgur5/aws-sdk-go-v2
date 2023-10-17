// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a resiliency policy for an application.
func (c *Client) CreateResiliencyPolicy(ctx context.Context, params *CreateResiliencyPolicyInput, optFns ...func(*Options)) (*CreateResiliencyPolicyOutput, error) {
	if params == nil {
		params = &CreateResiliencyPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResiliencyPolicy", params, optFns, c.addOperationCreateResiliencyPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResiliencyPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResiliencyPolicyInput struct {

	// The type of resiliency policy to be created, including the recovery time
	// objective (RTO) and recovery point objective (RPO) in seconds.
	//
	// This member is required.
	Policy map[string]types.FailurePolicy

	// The name of the policy
	//
	// This member is required.
	PolicyName *string

	// The tier for this resiliency policy, ranging from the highest severity (
	// MissionCritical ) to lowest ( NonCritical ).
	//
	// This member is required.
	Tier types.ResiliencyPolicyTier

	// Used for an idempotency token. A client token is a unique, case-sensitive
	// string of up to 64 ASCII characters. You should not reuse the same client token
	// for other API requests.
	ClientToken *string

	// Specifies a high-level geographical location constraint for where your
	// resilience policy data can be stored.
	DataLocationConstraint types.DataLocationConstraint

	// The description for the policy.
	PolicyDescription *string

	// Tags assigned to the resource. A tag is a label that you assign to an Amazon
	// Web Services resource. Each tag consists of a key/value pair.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateResiliencyPolicyOutput struct {

	// The type of resiliency policy that was created, including the recovery time
	// objective (RTO) and recovery point objective (RPO) in seconds.
	//
	// This member is required.
	Policy *types.ResiliencyPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResiliencyPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateResiliencyPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResiliencyPolicy{}, middleware.After)
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
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateResiliencyPolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateResiliencyPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResiliencyPolicy(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateResiliencyPolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateResiliencyPolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateResiliencyPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateResiliencyPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateResiliencyPolicyInput ")
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
func addIdempotencyToken_opCreateResiliencyPolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateResiliencyPolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateResiliencyPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "CreateResiliencyPolicy",
	}
}
