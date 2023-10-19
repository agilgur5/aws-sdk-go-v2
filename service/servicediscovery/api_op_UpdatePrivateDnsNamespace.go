// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a private DNS namespace.
func (c *Client) UpdatePrivateDnsNamespace(ctx context.Context, params *UpdatePrivateDnsNamespaceInput, optFns ...func(*Options)) (*UpdatePrivateDnsNamespaceOutput, error) {
	if params == nil {
		params = &UpdatePrivateDnsNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePrivateDnsNamespace", params, optFns, c.addOperationUpdatePrivateDnsNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePrivateDnsNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePrivateDnsNamespaceInput struct {

	// The ID of the namespace that you want to update.
	//
	// This member is required.
	Id *string

	// Updated properties for the private DNS namespace.
	//
	// This member is required.
	Namespace *types.PrivateDnsNamespaceChange

	// A unique string that identifies the request and that allows failed
	// UpdatePrivateDnsNamespace requests to be retried without the risk of running the
	// operation twice. UpdaterRequestId can be any unique string (for example, a
	// date/timestamp).
	UpdaterRequestId *string

	noSmithyDocumentSerde
}

type UpdatePrivateDnsNamespaceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see GetOperation (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html)
	// .
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePrivateDnsNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePrivateDnsNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePrivateDnsNamespace{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdatePrivateDnsNamespaceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdatePrivateDnsNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePrivateDnsNamespace(options.Region), middleware.Before); err != nil {
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
		operation: "UpdatePrivateDnsNamespace",
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

type idempotencyToken_initializeOpUpdatePrivateDnsNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdatePrivateDnsNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdatePrivateDnsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdatePrivateDnsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdatePrivateDnsNamespaceInput ")
	}

	if input.UpdaterRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.UpdaterRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdatePrivateDnsNamespaceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdatePrivateDnsNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdatePrivateDnsNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "UpdatePrivateDnsNamespace",
	}
}
