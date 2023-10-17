// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakeredge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use to check if a device is registered with SageMaker Edge Manager.
func (c *Client) GetDeviceRegistration(ctx context.Context, params *GetDeviceRegistrationInput, optFns ...func(*Options)) (*GetDeviceRegistrationOutput, error) {
	if params == nil {
		params = &GetDeviceRegistrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeviceRegistration", params, optFns, c.addOperationGetDeviceRegistrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeviceRegistrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeviceRegistrationInput struct {

	// The name of the fleet that the device belongs to.
	//
	// This member is required.
	DeviceFleetName *string

	// The unique name of the device you want to get the registration status from.
	//
	// This member is required.
	DeviceName *string

	noSmithyDocumentSerde
}

type GetDeviceRegistrationOutput struct {

	// The amount of time, in seconds, that the registration status is stored on the
	// device’s cache before it is refreshed.
	CacheTTL *string

	// Describes if the device is currently registered with SageMaker Edge Manager.
	DeviceRegistration *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeviceRegistrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDeviceRegistration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeviceRegistration{}, middleware.After)
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
	if err = addOpGetDeviceRegistrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeviceRegistration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeviceRegistration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "GetDeviceRegistration",
	}
}
