// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Composes an email message to multiple destinations.
func (c *Client) SendBulkEmail(ctx context.Context, params *SendBulkEmailInput, optFns ...func(*Options)) (*SendBulkEmailOutput, error) {
	if params == nil {
		params = &SendBulkEmailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendBulkEmail", params, optFns, c.addOperationSendBulkEmailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendBulkEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send email messages to multiple destinations using
// Amazon SES. For more information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html)
// .
type SendBulkEmailInput struct {

	// The list of bulk email entry objects.
	//
	// This member is required.
	BulkEmailEntries []types.BulkEmailEntry

	// An object that contains the body of the message. You can specify a template
	// message.
	//
	// This member is required.
	DefaultContent *types.BulkEmailContent

	// The name of the configuration set to use when sending the email.
	ConfigurationSetName *string

	// A list of tags, in the form of name/value pairs, to apply to an email that you
	// send using the SendEmail operation. Tags correspond to characteristics of the
	// email that you define, so that you can publish email sending events.
	DefaultEmailTags []types.MessageTag

	// The address that you want bounce and complaint notifications to be sent to.
	FeedbackForwardingEmailAddress *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the FeedbackForwardingEmailAddress
	// parameter. For example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com) attaches a policy to it
	// that authorizes you to use feedback@example.com, then you would specify the
	// FeedbackForwardingEmailAddressIdentityArn to be
	// arn:aws:ses:us-east-1:123456789012:identity/example.com, and the
	// FeedbackForwardingEmailAddress to be feedback@example.com. For more information
	// about sending authorization, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html)
	// .
	FeedbackForwardingEmailAddressIdentityArn *string

	// The email address to use as the "From" address for the email. The address that
	// you specify has to be verified.
	FromEmailAddress *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the FromEmailAddress parameter. For
	// example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com) attaches a policy to it
	// that authorizes you to use sender@example.com, then you would specify the
	// FromEmailAddressIdentityArn to be
	// arn:aws:ses:us-east-1:123456789012:identity/example.com, and the
	// FromEmailAddress to be sender@example.com. For more information about sending
	// authorization, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html)
	// .
	FromEmailAddressIdentityArn *string

	// The "Reply-to" email addresses for the message. When the recipient replies to
	// the message, each Reply-to address receives the reply.
	ReplyToAddresses []string

	noSmithyDocumentSerde
}

// The following data is returned in JSON format by the service.
type SendBulkEmailOutput struct {

	// One object per intended recipient. Check each response object and retry any
	// messages with a failure status.
	//
	// This member is required.
	BulkEmailEntryResults []types.BulkEmailEntryResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendBulkEmailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendBulkEmail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendBulkEmail{}, middleware.After)
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
	if err = addOpSendBulkEmailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendBulkEmail(options.Region), middleware.Before); err != nil {
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
		operation: "SendBulkEmail",
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

func newServiceMetadataMiddleware_opSendBulkEmail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendBulkEmail",
	}
}
