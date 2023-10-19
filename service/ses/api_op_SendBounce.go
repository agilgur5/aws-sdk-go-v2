// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates and sends a bounce message to the sender of an email you received
// through Amazon SES. You can only use this operation on an email up to 24 hours
// after you receive it. You cannot use this operation to send generic bounces for
// mail that was not received by Amazon SES. For information about receiving email
// through Amazon SES, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email.html)
// . You can execute this operation no more than once per second.
func (c *Client) SendBounce(ctx context.Context, params *SendBounceInput, optFns ...func(*Options)) (*SendBounceOutput, error) {
	if params == nil {
		params = &SendBounceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendBounce", params, optFns, c.addOperationSendBounceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendBounceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a bounce message to the sender of an email you
// received through Amazon SES.
type SendBounceInput struct {

	// The address to use in the "From" header of the bounce message. This must be an
	// identity that you have verified with Amazon SES.
	//
	// This member is required.
	BounceSender *string

	// A list of recipients of the bounced message, including the information required
	// to create the Delivery Status Notifications (DSNs) for the recipients. You must
	// specify at least one BouncedRecipientInfo in the list.
	//
	// This member is required.
	BouncedRecipientInfoList []types.BouncedRecipientInfo

	// The message ID of the message to be bounced.
	//
	// This member is required.
	OriginalMessageId *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the address in the "From" header of the bounce. For more information
	// about sending authorization, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html)
	// .
	BounceSenderArn *string

	// Human-readable text for the bounce message to explain the failure. If not
	// specified, the text is auto-generated based on the bounced recipient
	// information.
	Explanation *string

	// Message-related DSN fields. If not specified, Amazon SES chooses the values.
	MessageDsn *types.MessageDsn

	noSmithyDocumentSerde
}

// Represents a unique message ID.
type SendBounceOutput struct {

	// The message ID of the bounce message.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendBounceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSendBounce{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSendBounce{}, middleware.After)
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
	if err = addOpSendBounceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendBounce(options.Region), middleware.Before); err != nil {
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
		operation: "SendBounce",
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

func newServiceMetadataMiddleware_opSendBounce(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendBounce",
	}
}
