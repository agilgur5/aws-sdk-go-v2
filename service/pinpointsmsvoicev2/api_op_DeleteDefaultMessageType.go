// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing default message type on a configuration set. A message type
// is a type of messages that you plan to send. If you send account-related
// messages or time-sensitive messages such as one-time passcodes, choose
// Transactional. If you plan to send messages that contain marketing material or
// other promotional content, choose Promotional. This setting applies to your
// entire Amazon Web Services account.
func (c *Client) DeleteDefaultMessageType(ctx context.Context, params *DeleteDefaultMessageTypeInput, optFns ...func(*Options)) (*DeleteDefaultMessageTypeOutput, error) {
	if params == nil {
		params = &DeleteDefaultMessageTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDefaultMessageType", params, optFns, c.addOperationDeleteDefaultMessageTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDefaultMessageTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDefaultMessageTypeInput struct {

	// The name of the configuration set or the configuration set Amazon Resource Name
	// (ARN) to delete the default message type from. The ConfigurationSetName and
	// ConfigurationSetArn can be found using the DescribeConfigurationSets action.
	//
	// This member is required.
	ConfigurationSetName *string

	noSmithyDocumentSerde
}

type DeleteDefaultMessageTypeOutput struct {

	// The Amazon Resource Name (ARN) of the configuration set.
	ConfigurationSetArn *string

	// The name of the configuration set.
	ConfigurationSetName *string

	// The current message type for the configuration set.
	MessageType types.MessageType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDefaultMessageTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteDefaultMessageType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteDefaultMessageType{}, middleware.After)
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
	if err = addOpDeleteDefaultMessageTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDefaultMessageType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDefaultMessageType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DeleteDefaultMessageType",
	}
}
