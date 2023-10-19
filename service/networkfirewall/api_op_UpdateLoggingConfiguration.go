// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the logging configuration for the specified firewall. To change the
// logging configuration, retrieve the LoggingConfiguration by calling
// DescribeLoggingConfiguration , then change it and provide the modified object to
// this update call. You must change the logging configuration one
// LogDestinationConfig at a time inside the retrieved LoggingConfiguration
// object. You can perform only one of the following actions in any call to
// UpdateLoggingConfiguration :
//   - Create a new log destination object by adding a single LogDestinationConfig
//     array element to LogDestinationConfigs .
//   - Delete a log destination object by removing a single LogDestinationConfig
//     array element from LogDestinationConfigs .
//   - Change the LogDestination setting in a single LogDestinationConfig array
//     element.
//
// You can't change the LogDestinationType or LogType in a LogDestinationConfig .
// To change these settings, delete the existing LogDestinationConfig object and
// create a new one, using two separate calls to this update operation.
func (c *Client) UpdateLoggingConfiguration(ctx context.Context, params *UpdateLoggingConfigurationInput, optFns ...func(*Options)) (*UpdateLoggingConfigurationOutput, error) {
	if params == nil {
		params = &UpdateLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLoggingConfiguration", params, optFns, c.addOperationUpdateLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLoggingConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the firewall. You must specify the ARN or the
	// name, and you can specify both.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it. You must specify the ARN or the name, and you can specify
	// both.
	FirewallName *string

	// Defines how Network Firewall performs logging for a firewall. If you omit this
	// setting, Network Firewall disables logging for the firewall.
	LoggingConfiguration *types.LoggingConfiguration

	noSmithyDocumentSerde
}

type UpdateLoggingConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the firewall.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it.
	FirewallName *string

	// Defines how Network Firewall performs logging for a Firewall .
	LoggingConfiguration *types.LoggingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateLoggingConfiguration{}, middleware.After)
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
	if err = addOpUpdateLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLoggingConfiguration(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateLoggingConfiguration",
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

func newServiceMetadataMiddleware_opUpdateLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "UpdateLoggingConfiguration",
	}
}
