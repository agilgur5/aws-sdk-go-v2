// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an accelerator. An accelerator includes one or more listeners that
// process inbound connections and direct traffic to one or more endpoint groups,
// each of which includes endpoints, such as Network Load Balancers. Global
// Accelerator is a global service that supports endpoints in multiple Amazon Web
// Services Regions but you must specify the US West (Oregon) Region to create,
// update, or otherwise work with accelerators. That is, for example, specify
// --region us-west-2 on Amazon Web Services CLI commands.
func (c *Client) CreateAccelerator(ctx context.Context, params *CreateAcceleratorInput, optFns ...func(*Options)) (*CreateAcceleratorOutput, error) {
	if params == nil {
		params = &CreateAcceleratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccelerator", params, optFns, c.addOperationCreateAcceleratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAcceleratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAcceleratorInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the
	// idempotency—that is, the uniqueness—of an accelerator.
	//
	// This member is required.
	IdempotencyToken *string

	// The name of the accelerator. The name can have a maximum of 64 characters, must
	// contain only alphanumeric characters, periods (.), or hyphens (-), and must not
	// begin or end with a hyphen or period.
	//
	// This member is required.
	Name *string

	// Indicates whether an accelerator is enabled. The value is true or false. The
	// default value is true. If the value is set to true, an accelerator cannot be
	// deleted. If set to false, the accelerator can be deleted.
	Enabled *bool

	// The IP address type that an accelerator supports. For a standard accelerator,
	// the value can be IPV4 or DUAL_STACK.
	IpAddressType types.IpAddressType

	// Optionally, if you've added your own IP address pool to Global Accelerator
	// (BYOIP), you can choose an IPv4 address from your own pool to use for the
	// accelerator's static IPv4 address when you create an accelerator. After you
	// bring an address range to Amazon Web Services, it appears in your account as an
	// address pool. When you create an accelerator, you can assign one IPv4 address
	// from your range to it. Global Accelerator assigns you a second static IPv4
	// address from an Amazon IP address range. If you bring two IPv4 address ranges to
	// Amazon Web Services, you can assign one IPv4 address from each range to your
	// accelerator. This restriction is because Global Accelerator assigns each address
	// range to a different network zone, for high availability. You can specify one or
	// two addresses, separated by a space. Do not include the /32 suffix. Note that
	// you can't update IP addresses for an existing accelerator. To change them, you
	// must create a new accelerator with the new addresses. For more information, see
	// Bring your own IP addresses (BYOIP) (https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html)
	// in the Global Accelerator Developer Guide.
	IpAddresses []string

	// Create tags for an accelerator. For more information, see Tagging in Global
	// Accelerator (https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html)
	// in the Global Accelerator Developer Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAcceleratorOutput struct {

	// The accelerator that is created by specifying a listener and the supported IP
	// address types.
	Accelerator *types.Accelerator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAcceleratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAccelerator{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAcceleratorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAcceleratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccelerator(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAccelerator struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAccelerator) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAccelerator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAcceleratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAcceleratorInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAcceleratorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAccelerator{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAccelerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "CreateAccelerator",
	}
}
