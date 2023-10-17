// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an association between an Amazon Connect Wisdom assistant and another
// resource. Currently, the only supported association is with a knowledge base. An
// assistant can have only a single association.
func (c *Client) CreateAssistantAssociation(ctx context.Context, params *CreateAssistantAssociationInput, optFns ...func(*Options)) (*CreateAssistantAssociationOutput, error) {
	if params == nil {
		params = &CreateAssistantAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssistantAssociation", params, optFns, c.addOperationCreateAssistantAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssistantAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssistantAssociationInput struct {

	// The identifier of the Wisdom assistant. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	AssistantId *string

	// The identifier of the associated resource.
	//
	// This member is required.
	Association types.AssistantAssociationInputData

	// The type of association.
	//
	// This member is required.
	AssociationType types.AssociationType

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAssistantAssociationOutput struct {

	// The assistant association.
	AssistantAssociation *types.AssistantAssociationData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssistantAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssistantAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssistantAssociation{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAssistantAssociationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAssistantAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssistantAssociation(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAssistantAssociation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAssistantAssociation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAssistantAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAssistantAssociationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAssistantAssociationInput ")
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
func addIdempotencyToken_opCreateAssistantAssociationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAssistantAssociation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAssistantAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wisdom",
		OperationName: "CreateAssistantAssociation",
	}
}
