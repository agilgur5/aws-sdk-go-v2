// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a delivery channel object to deliver configuration information and
// other compliance information to an Amazon S3 bucket and Amazon SNS topic. For
// more information, see Notifications that Config Sends to an Amazon SNS topic (https://docs.aws.amazon.com/config/latest/developerguide/notifications-for-AWS-Config.html)
// . Before you can create a delivery channel, you must create a configuration
// recorder. You can use this action to change the Amazon S3 bucket or an Amazon
// SNS topic of the existing delivery channel. To change the Amazon S3 bucket or an
// Amazon SNS topic, call this action and specify the changed values for the S3
// bucket and the SNS topic. If you specify a different value for either the S3
// bucket or the SNS topic, this action will keep the existing value for the
// parameter that is not changed. You can have only one delivery channel per region
// in your account.
func (c *Client) PutDeliveryChannel(ctx context.Context, params *PutDeliveryChannelInput, optFns ...func(*Options)) (*PutDeliveryChannelOutput, error) {
	if params == nil {
		params = &PutDeliveryChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDeliveryChannel", params, optFns, c.addOperationPutDeliveryChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDeliveryChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the PutDeliveryChannel action.
type PutDeliveryChannelInput struct {

	// The configuration delivery channel object that delivers the configuration
	// information to an Amazon S3 bucket and to an Amazon SNS topic.
	//
	// This member is required.
	DeliveryChannel *types.DeliveryChannel

	noSmithyDocumentSerde
}

type PutDeliveryChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDeliveryChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutDeliveryChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutDeliveryChannel{}, middleware.After)
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
	if err = addOpPutDeliveryChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDeliveryChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDeliveryChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutDeliveryChannel",
	}
}
