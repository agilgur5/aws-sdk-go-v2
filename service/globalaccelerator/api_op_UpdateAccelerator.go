// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an accelerator to make changes, such as the following:
//   - Change the name of the accelerator.
//   - Disable the accelerator so that it no longer accepts or routes traffic, or
//     so that you can delete it.
//   - Enable the accelerator, if it is disabled.
//   - Change the IP address type to dual-stack if it is IPv4, or change the IP
//     address type to IPv4 if it's dual-stack.
//
// Be aware that static IP addresses remain assigned to your accelerator for as
// long as it exists, even if you disable the accelerator and it no longer accepts
// or routes traffic. However, when you delete the accelerator, you lose the static
// IP addresses that are assigned to it, so you can no longer route traffic by
// using them. Global Accelerator is a global service that supports endpoints in
// multiple Amazon Web Services Regions but you must specify the US West (Oregon)
// Region to create, update, or otherwise work with accelerators. That is, for
// example, specify --region us-west-2 on Amazon Web Services CLI commands.
func (c *Client) UpdateAccelerator(ctx context.Context, params *UpdateAcceleratorInput, optFns ...func(*Options)) (*UpdateAcceleratorOutput, error) {
	if params == nil {
		params = &UpdateAcceleratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccelerator", params, optFns, c.addOperationUpdateAcceleratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAcceleratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAcceleratorInput struct {

	// The Amazon Resource Name (ARN) of the accelerator to update.
	//
	// This member is required.
	AcceleratorArn *string

	// Indicates whether an accelerator is enabled. The value is true or false. The
	// default value is true. If the value is set to true, the accelerator cannot be
	// deleted. If set to false, the accelerator can be deleted.
	Enabled *bool

	// The IP address type that an accelerator supports. For a standard accelerator,
	// the value can be IPV4 or DUAL_STACK.
	IpAddressType types.IpAddressType

	// The name of the accelerator. The name can have a maximum of 64 characters, must
	// contain only alphanumeric characters, periods (.), or hyphens (-), and must not
	// begin or end with a hyphen or period.
	Name *string

	noSmithyDocumentSerde
}

type UpdateAcceleratorOutput struct {

	// Information about the updated accelerator.
	Accelerator *types.Accelerator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAcceleratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAccelerator{}, middleware.After)
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
	if err = addOpUpdateAcceleratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccelerator(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateAccelerator",
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

func newServiceMetadataMiddleware_opUpdateAccelerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "UpdateAccelerator",
	}
}
