// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new studio. When creating a studio, two IAM roles must be provided:
// the admin role and the user role. These roles are assumed by your users when
// they log in to the Nimble Studio portal. The user role must have the
// AmazonNimbleStudio-StudioUser managed policy attached for the portal to function
// properly. The admin role must have the AmazonNimbleStudio-StudioAdmin managed
// policy attached for the portal to function properly. You may optionally specify
// a KMS key in the StudioEncryptionConfiguration . In Nimble Studio, resource
// names, descriptions, initialization scripts, and other data you provide are
// always encrypted at rest using an KMS key. By default, this key is owned by
// Amazon Web Services and managed on your behalf. You may provide your own KMS key
// when calling CreateStudio to encrypt this data using a key you own and manage.
// When providing an KMS key during studio creation, Nimble Studio creates KMS
// grants in your account to provide your studio user and admin roles access to
// these KMS keys. If you delete this grant, the studio will no longer be
// accessible to your portal users. If you delete the studio KMS key, your studio
// will no longer be accessible.
func (c *Client) CreateStudio(ctx context.Context, params *CreateStudioInput, optFns ...func(*Options)) (*CreateStudioOutput, error) {
	if params == nil {
		params = &CreateStudioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStudio", params, optFns, c.addOperationCreateStudioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStudioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStudioInput struct {

	// The IAM role that studio admins will assume when logging in to the Nimble
	// Studio portal.
	//
	// This member is required.
	AdminRoleArn *string

	// A friendly name for the studio.
	//
	// This member is required.
	DisplayName *string

	// The studio name that is used in the URL of the Nimble Studio portal when
	// accessed by Nimble Studio users.
	//
	// This member is required.
	StudioName *string

	// The IAM role that studio users will assume when logging in to the Nimble Studio
	// portal.
	//
	// This member is required.
	UserRoleArn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don’t specify a client token, the Amazon Web Services SDK
	// automatically generates a client token and uses it for the request to ensure
	// idempotency.
	ClientToken *string

	// The studio encryption configuration.
	StudioEncryptionConfiguration *types.StudioEncryptionConfiguration

	// A collection of labels, in the form of key-value pairs, that apply to this
	// resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateStudioOutput struct {

	// Information about a studio.
	Studio *types.Studio

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStudioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateStudio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateStudio{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateStudioMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateStudioValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStudio(options.Region), middleware.Before); err != nil {
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
		operation: "CreateStudio",
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

type idempotencyToken_initializeOpCreateStudio struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateStudio) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateStudio) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateStudioInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateStudioInput ")
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
func addIdempotencyToken_opCreateStudioMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateStudio{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateStudio(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "CreateStudio",
	}
}
