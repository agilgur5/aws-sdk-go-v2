// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a model-specific endpoint for synchronous inference for a previously
// trained custom model For information about endpoints, see Managing endpoints (https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html)
// .
func (c *Client) CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) {
	if params == nil {
		params = &CreateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpoint", params, optFns, c.addOperationCreateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointInput struct {

	// The desired number of inference units to be used by the model using this
	// endpoint. Each inference unit represents of a throughput of 100 characters per
	// second.
	//
	// This member is required.
	DesiredInferenceUnits *int32

	// This is the descriptive suffix that becomes part of the EndpointArn used for
	// all subsequent requests to this resource.
	//
	// This member is required.
	EndpointName *string

	// An idempotency token provided by the customer. If this token matches a previous
	// endpoint creation request, Amazon Comprehend will not return a
	// ResourceInUseException .
	ClientRequestToken *string

	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend
	// read access to trained custom models encrypted with a customer managed key
	// (ModelKmsKeyId).
	DataAccessRoleArn *string

	// The Amazon Resource Number (ARN) of the flywheel to which the endpoint will be
	// attached.
	FlywheelArn *string

	// The Amazon Resource Number (ARN) of the model to which the endpoint will be
	// attached.
	ModelArn *string

	// Tags to associate with the endpoint. A tag is a key-value pair that adds
	// metadata to the endpoint. For example, a tag with "Sales" as the key might be
	// added to an endpoint to indicate its use by the sales department.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEndpointOutput struct {

	// The Amazon Resource Number (ARN) of the endpoint being created.
	EndpointArn *string

	// The Amazon Resource Number (ARN) of the model to which the endpoint is attached.
	ModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpoint{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpoint(options.Region), middleware.Before); err != nil {
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
		operation: "CreateEndpoint",
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

type idempotencyToken_initializeOpCreateEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEndpointInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "CreateEndpoint",
	}
}
