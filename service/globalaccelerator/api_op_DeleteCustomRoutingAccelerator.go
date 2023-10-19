// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a custom routing accelerator. Before you can delete an accelerator, you
// must disable it and remove all dependent resources (listeners and endpoint
// groups). To disable the accelerator, update the accelerator to set Enabled to
// false. When you create a custom routing accelerator, by default, Global
// Accelerator provides you with a set of two static IP addresses. The IP addresses
// are assigned to your accelerator for as long as it exists, even if you disable
// the accelerator and it no longer accepts or routes traffic. However, when you
// delete an accelerator, you lose the static IP addresses that are assigned to the
// accelerator, so you can no longer route traffic by using them. As a best
// practice, ensure that you have permissions in place to avoid inadvertently
// deleting accelerators. You can use IAM policies with Global Accelerator to limit
// the users who have permissions to delete an accelerator. For more information,
// see Identity and access management (https://docs.aws.amazon.com/global-accelerator/latest/dg/auth-and-access-control.html)
// in the Global Accelerator Developer Guide.
func (c *Client) DeleteCustomRoutingAccelerator(ctx context.Context, params *DeleteCustomRoutingAcceleratorInput, optFns ...func(*Options)) (*DeleteCustomRoutingAcceleratorOutput, error) {
	if params == nil {
		params = &DeleteCustomRoutingAcceleratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomRoutingAccelerator", params, optFns, c.addOperationDeleteCustomRoutingAcceleratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomRoutingAcceleratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCustomRoutingAcceleratorInput struct {

	// The Amazon Resource Name (ARN) of the custom routing accelerator to delete.
	//
	// This member is required.
	AcceleratorArn *string

	noSmithyDocumentSerde
}

type DeleteCustomRoutingAcceleratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCustomRoutingAcceleratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCustomRoutingAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCustomRoutingAccelerator{}, middleware.After)
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
	if err = addOpDeleteCustomRoutingAcceleratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomRoutingAccelerator(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteCustomRoutingAccelerator",
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

func newServiceMetadataMiddleware_opDeleteCustomRoutingAccelerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "DeleteCustomRoutingAccelerator",
	}
}
