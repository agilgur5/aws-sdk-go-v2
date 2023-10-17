// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests updates to the configuration of the specified provisioned product. If
// there are tags associated with the object, they cannot be updated or added.
// Depending on the specific updates requested, this operation can update with no
// interruption, with some interruption, or replace the provisioned product
// entirely. You can check the status of this request using DescribeRecord .
func (c *Client) UpdateProvisionedProduct(ctx context.Context, params *UpdateProvisionedProductInput, optFns ...func(*Options)) (*UpdateProvisionedProductOutput, error) {
	if params == nil {
		params = &UpdateProvisionedProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProvisionedProduct", params, optFns, c.addOperationUpdateProvisionedProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProvisionedProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProvisionedProductInput struct {

	// The idempotency token that uniquely identifies the provisioning update request.
	//
	// This member is required.
	UpdateToken *string

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	// The path identifier. This value is optional if the product has a default path,
	// and required if the product has more than one path. You must provide the name or
	// ID, but not both.
	PathId *string

	// The name of the path. You must provide the name or ID, but not both.
	PathName *string

	// The identifier of the product. You must provide the name or ID, but not both.
	ProductId *string

	// The name of the product. You must provide the name or ID, but not both.
	ProductName *string

	// The identifier of the provisioned product. You must provide the name or ID, but
	// not both.
	ProvisionedProductId *string

	// The name of the provisioned product. You cannot specify both
	// ProvisionedProductName and ProvisionedProductId .
	ProvisionedProductName *string

	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string

	// The name of the provisioning artifact. You must provide the name or ID, but not
	// both.
	ProvisioningArtifactName *string

	// The new parameters.
	ProvisioningParameters []types.UpdateProvisioningParameter

	// An object that contains information about the provisioning preferences for a
	// stack set.
	ProvisioningPreferences *types.UpdateProvisioningPreferences

	// One or more tags. Requires the product to have RESOURCE_UPDATE constraint with
	// TagUpdatesOnProvisionedProduct set to ALLOWED to allow tag updates.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type UpdateProvisionedProductOutput struct {

	// Information about the result of the request.
	RecordDetail *types.RecordDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProvisionedProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProvisionedProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProvisionedProduct{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateProvisionedProductMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateProvisionedProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProvisionedProduct(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateProvisionedProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateProvisionedProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateProvisionedProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateProvisionedProductInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateProvisionedProductInput ")
	}

	if input.UpdateToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.UpdateToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateProvisionedProductMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateProvisionedProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateProvisionedProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "UpdateProvisionedProduct",
	}
}
