// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the network profile.
func (c *Client) UpdateNetworkProfile(ctx context.Context, params *UpdateNetworkProfileInput, optFns ...func(*Options)) (*UpdateNetworkProfileOutput, error) {
	if params == nil {
		params = &UpdateNetworkProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNetworkProfile", params, optFns, c.addOperationUpdateNetworkProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNetworkProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNetworkProfileInput struct {

	// The Amazon Resource Name (ARN) of the project for which you want to update
	// network profile settings.
	//
	// This member is required.
	Arn *string

	// The description of the network profile about which you are returning
	// information.
	Description *string

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	DownlinkBandwidthBits *int64

	// Delay time for all packets to destination in milliseconds as an integer from 0
	// to 2000.
	DownlinkDelayMs *int64

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	DownlinkJitterMs *int64

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent int32

	// The name of the network profile about which you are returning information.
	Name *string

	// The type of network profile to return information about. Valid values are
	// listed here.
	Type types.NetworkProfileType

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	UplinkBandwidthBits *int64

	// Delay time for all packets to destination in milliseconds as an integer from 0
	// to 2000.
	UplinkDelayMs *int64

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	UplinkJitterMs *int64

	// Proportion of transmitted packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent int32

	noSmithyDocumentSerde
}

type UpdateNetworkProfileOutput struct {

	// A list of the available network profiles.
	NetworkProfile *types.NetworkProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNetworkProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNetworkProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNetworkProfile{}, middleware.After)
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
	if err = addOpUpdateNetworkProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNetworkProfile(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateNetworkProfile",
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

func newServiceMetadataMiddleware_opUpdateNetworkProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "UpdateNetworkProfile",
	}
}
