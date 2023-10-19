// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a vehicle, which is an instance of a vehicle model (model manifest).
// Vehicles created from the same vehicle model consist of the same signals
// inherited from the vehicle model. If you have an existing Amazon Web Services
// IoT thing, you can use Amazon Web Services IoT FleetWise to create a vehicle and
// collect data from your thing. For more information, see Create a vehicle (AWS
// CLI) (https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/create-vehicle-cli.html)
// in the Amazon Web Services IoT FleetWise Developer Guide.
func (c *Client) CreateVehicle(ctx context.Context, params *CreateVehicleInput, optFns ...func(*Options)) (*CreateVehicleOutput, error) {
	if params == nil {
		params = &CreateVehicleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVehicle", params, optFns, c.addOperationCreateVehicleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVehicleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVehicleInput struct {

	// The ARN of a decoder manifest.
	//
	// This member is required.
	DecoderManifestArn *string

	// The Amazon Resource Name ARN of a vehicle model.
	//
	// This member is required.
	ModelManifestArn *string

	// The unique ID of the vehicle to create.
	//
	// This member is required.
	VehicleName *string

	// An option to create a new Amazon Web Services IoT thing when creating a
	// vehicle, or to validate an existing Amazon Web Services IoT thing as a vehicle.
	// Default:
	AssociationBehavior types.VehicleAssociationBehavior

	// Static information about a vehicle in a key-value pair. For example:
	// "engineType" : "1.3 L R2" A campaign must include the keys (attribute names) in
	// dataExtraDimensions for them to display in Amazon Timestream.
	Attributes map[string]string

	// Metadata that can be used to manage the vehicle.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateVehicleOutput struct {

	// The ARN of the created vehicle.
	Arn *string

	// The ARN of a created or validated Amazon Web Services IoT thing.
	ThingArn *string

	// The unique ID of the created vehicle.
	VehicleName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVehicleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateVehicle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateVehicle{}, middleware.After)
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
	if err = addOpCreateVehicleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVehicle(options.Region), middleware.Before); err != nil {
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
		operation: "CreateVehicle",
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

func newServiceMetadataMiddleware_opCreateVehicle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "CreateVehicle",
	}
}
