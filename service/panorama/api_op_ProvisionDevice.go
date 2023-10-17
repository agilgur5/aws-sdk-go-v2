// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a device and returns a configuration archive. The configuration archive
// is a ZIP file that contains a provisioning certificate that is valid for 5
// minutes. Name the configuration archive certificates-omni_device-name.zip and
// transfer it to the device within 5 minutes. Use the included USB storage device
// and connect it to the USB 3.0 port next to the HDMI output.
func (c *Client) ProvisionDevice(ctx context.Context, params *ProvisionDeviceInput, optFns ...func(*Options)) (*ProvisionDeviceOutput, error) {
	if params == nil {
		params = &ProvisionDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ProvisionDevice", params, optFns, c.addOperationProvisionDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ProvisionDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ProvisionDeviceInput struct {

	// A name for the device.
	//
	// This member is required.
	Name *string

	// A description for the device.
	Description *string

	// A networking configuration for the device.
	NetworkingConfiguration *types.NetworkPayload

	// Tags for the device.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ProvisionDeviceOutput struct {

	// The device's ARN.
	//
	// This member is required.
	Arn *string

	// The device's status.
	//
	// This member is required.
	Status types.DeviceStatus

	// The device's configuration bundle.
	Certificates []byte

	// The device's ID.
	DeviceId *string

	// The device's IoT thing name.
	IotThingName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationProvisionDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpProvisionDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpProvisionDevice{}, middleware.After)
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
	if err = addOpProvisionDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opProvisionDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opProvisionDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "panorama",
		OperationName: "ProvisionDevice",
	}
}
