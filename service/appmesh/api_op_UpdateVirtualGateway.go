// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing virtual gateway in a specified service mesh.
func (c *Client) UpdateVirtualGateway(ctx context.Context, params *UpdateVirtualGatewayInput, optFns ...func(*Options)) (*UpdateVirtualGatewayOutput, error) {
	if params == nil {
		params = &UpdateVirtualGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVirtualGateway", params, optFns, c.addOperationUpdateVirtualGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVirtualGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVirtualGatewayInput struct {

	// The name of the service mesh that the virtual gateway resides in.
	//
	// This member is required.
	MeshName *string

	// The new virtual gateway specification to apply. This overwrites the existing
	// data.
	//
	// This member is required.
	Spec *types.VirtualGatewaySpec

	// The name of the virtual gateway to update.
	//
	// This member is required.
	VirtualGatewayName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The Amazon Web Services IAM account ID of the service mesh owner. If the
	// account ID is not your own, then it's the ID of the account that shared the mesh
	// with your account. For more information about mesh sharing, see Working with
	// shared meshes (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html)
	// .
	MeshOwner *string

	noSmithyDocumentSerde
}

type UpdateVirtualGatewayOutput struct {

	// A full description of the virtual gateway that was updated.
	//
	// This member is required.
	VirtualGateway *types.VirtualGatewayData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVirtualGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVirtualGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVirtualGateway{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateVirtualGatewayMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateVirtualGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVirtualGateway(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateVirtualGateway struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateVirtualGateway) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateVirtualGateway) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateVirtualGatewayInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateVirtualGatewayInput ")
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
func addIdempotencyToken_opUpdateVirtualGatewayMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateVirtualGateway{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateVirtualGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appmesh",
		OperationName: "UpdateVirtualGateway",
	}
}
