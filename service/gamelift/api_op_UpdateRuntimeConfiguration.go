// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the current runtime configuration for the specified fleet, which tells
// Amazon GameLift how to launch server processes on all instances in the fleet.
// You can update a fleet's runtime configuration at any time after the fleet is
// created; it does not need to be in ACTIVE status. To update runtime
// configuration, specify the fleet ID and provide a RuntimeConfiguration with an
// updated set of server process configurations. If successful, the fleet's runtime
// configuration settings are updated. Each instance in the fleet regularly checks
// for and retrieves updated runtime configurations. Instances immediately begin
// complying with the new configuration by launching new server processes or not
// replacing existing processes when they shut down. Updating a fleet's runtime
// configuration never affects existing server processes. Learn more Setting up
// Amazon GameLift fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
func (c *Client) UpdateRuntimeConfiguration(ctx context.Context, params *UpdateRuntimeConfigurationInput, optFns ...func(*Options)) (*UpdateRuntimeConfigurationOutput, error) {
	if params == nil {
		params = &UpdateRuntimeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRuntimeConfiguration", params, optFns, c.addOperationUpdateRuntimeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRuntimeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRuntimeConfigurationInput struct {

	// A unique identifier for the fleet to update runtime configuration for. You can
	// use either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// Instructions for launching server processes on each instance in the fleet.
	// Server processes run either a custom game build executable or a Realtime Servers
	// script. The runtime configuration lists the types of server processes to run on
	// an instance, how to launch them, and the number of processes to run
	// concurrently.
	//
	// This member is required.
	RuntimeConfiguration *types.RuntimeConfiguration

	noSmithyDocumentSerde
}

type UpdateRuntimeConfigurationOutput struct {

	// The runtime configuration currently in use by all instances in the fleet. If
	// the update was successful, all property changes are shown.
	RuntimeConfiguration *types.RuntimeConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRuntimeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRuntimeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRuntimeConfiguration{}, middleware.After)
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
	if err = addOpUpdateRuntimeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRuntimeConfiguration(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateRuntimeConfiguration",
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

func newServiceMetadataMiddleware_opUpdateRuntimeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateRuntimeConfiguration",
	}
}
