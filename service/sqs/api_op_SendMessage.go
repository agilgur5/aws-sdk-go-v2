// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delivers a message to the specified queue. A message can include only XML,
// JSON, and unformatted text. The following Unicode characters are allowed: #x9 |
// #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF Any
// characters not included in this list will be rejected. For more information, see
// the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets) .
func (c *Client) SendMessage(ctx context.Context, params *SendMessageInput, optFns ...func(*Options)) (*SendMessageOutput, error) {
	if params == nil {
		params = &SendMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendMessage", params, optFns, c.addOperationSendMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendMessageInput struct {

	// The message to send. The minimum size is one character. The maximum size is 256
	// KiB. A message can include only XML, JSON, and unformatted text. The following
	// Unicode characters are allowed: #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to
	// #xFFFD | #x10000 to #x10FFFF Any characters not included in this list will be
	// rejected. For more information, see the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets)
	// .
	//
	// This member is required.
	MessageBody *string

	// The URL of the Amazon SQS queue to which a message is sent. Queue URLs and
	// names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	// The length of time, in seconds, for which to delay a specific message. Valid
	// values: 0 to 900. Maximum: 15 minutes. Messages with a positive DelaySeconds
	// value become available for processing after the delay period is finished. If you
	// don't specify a value, the default value for the queue applies. When you set
	// FifoQueue , you can't set DelaySeconds per message. You can set this parameter
	// only on a queue level.
	DelaySeconds int32

	// Each message attribute consists of a Name , Type , and Value . For more
	// information, see Amazon SQS message attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-metadata.html#sqs-message-attributes)
	// in the Amazon SQS Developer Guide.
	MessageAttributes map[string]types.MessageAttributeValue

	// This parameter applies only to FIFO (first-in-first-out) queues. The token used
	// for deduplication of sent messages. If a message with a particular
	// MessageDeduplicationId is sent successfully, any messages sent with the same
	// MessageDeduplicationId are accepted successfully but aren't delivered during the
	// 5-minute deduplication interval. For more information, see Exactly-once
	// processing (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-exactly-once-processing.html)
	// in the Amazon SQS Developer Guide.
	//   - Every message must have a unique MessageDeduplicationId ,
	//   - You may provide a MessageDeduplicationId explicitly.
	//   - If you aren't able to provide a MessageDeduplicationId and you enable
	//   ContentBasedDeduplication for your queue, Amazon SQS uses a SHA-256 hash to
	//   generate the MessageDeduplicationId using the body of the message (but not the
	//   attributes of the message).
	//   - If you don't provide a MessageDeduplicationId and the queue doesn't have
	//   ContentBasedDeduplication set, the action fails with an error.
	//   - If the queue has ContentBasedDeduplication set, your MessageDeduplicationId
	//   overrides the generated one.
	//   - When ContentBasedDeduplication is in effect, messages with identical content
	//   sent within the deduplication interval are treated as duplicates and only one
	//   copy of the message is delivered.
	//   - If you send one message with ContentBasedDeduplication enabled and then
	//   another message with a MessageDeduplicationId that is the same as the one
	//   generated for the first MessageDeduplicationId , the two messages are treated
	//   as duplicates and only one copy of the message is delivered.
	// The MessageDeduplicationId is available to the consumer of the message (this
	// can be useful for troubleshooting delivery issues). If a message is sent
	// successfully but the acknowledgement is lost and the message is resent with the
	// same MessageDeduplicationId after the deduplication interval, Amazon SQS can't
	// detect duplicate messages. Amazon SQS continues to keep track of the message
	// deduplication ID even after the message is received and deleted. The maximum
	// length of MessageDeduplicationId is 128 characters. MessageDeduplicationId can
	// contain alphanumeric characters ( a-z , A-Z , 0-9 ) and punctuation (
	// !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~ ). For best practices of using
	// MessageDeduplicationId , see Using the MessageDeduplicationId Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagededuplicationid-property.html)
	// in the Amazon SQS Developer Guide.
	MessageDeduplicationId *string

	// This parameter applies only to FIFO (first-in-first-out) queues. The tag that
	// specifies that a message belongs to a specific message group. Messages that
	// belong to the same message group are processed in a FIFO manner (however,
	// messages in different message groups might be processed out of order). To
	// interleave multiple ordered streams within a single queue, use MessageGroupId
	// values (for example, session data for multiple users). In this scenario,
	// multiple consumers can process the queue, but the session data of each user is
	// processed in a FIFO fashion.
	//   - You must associate a non-empty MessageGroupId with a message. If you don't
	//   provide a MessageGroupId , the action fails.
	//   - ReceiveMessage might return messages with multiple MessageGroupId values.
	//   For each MessageGroupId , the messages are sorted by time sent. The caller
	//   can't specify a MessageGroupId .
	// The length of MessageGroupId is 128 characters. Valid values: alphanumeric
	// characters and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~) . For best
	// practices of using MessageGroupId , see Using the MessageGroupId Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagegroupid-property.html)
	// in the Amazon SQS Developer Guide. MessageGroupId is required for FIFO queues.
	// You can't use it for Standard queues.
	MessageGroupId *string

	// The message system attribute to send. Each message system attribute consists of
	// a Name , Type , and Value .
	//   - Currently, the only supported message system attribute is AWSTraceHeader .
	//   Its type must be String and its value must be a correctly formatted X-Ray
	//   trace header string.
	//   - The size of a message system attribute doesn't count towards the total size
	//   of a message.
	MessageSystemAttributes map[string]types.MessageSystemAttributeValue

	noSmithyDocumentSerde
}

func (*SendMessageInput) operationName() string {
	return "SendMessage"
}

// The MD5OfMessageBody and MessageId elements.
type SendMessageOutput struct {

	// An MD5 digest of the non-URL-encoded message attribute string. You can use this
	// attribute to verify that Amazon SQS received the message correctly. Amazon SQS
	// URL-decodes the message before creating the MD5 digest. For information about
	// MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt) .
	MD5OfMessageAttributes *string

	// An MD5 digest of the non-URL-encoded message body string. You can use this
	// attribute to verify that Amazon SQS received the message correctly. Amazon SQS
	// URL-decodes the message before creating the MD5 digest. For information about
	// MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt) .
	MD5OfMessageBody *string

	// An MD5 digest of the non-URL-encoded message system attribute string. You can
	// use this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest.
	MD5OfMessageSystemAttributes *string

	// An attribute containing the MessageId of the message sent to the queue. For
	// more information, see Queue and Message Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html)
	// in the Amazon SQS Developer Guide.
	MessageId *string

	// This parameter applies only to FIFO (first-in-first-out) queues. The large,
	// non-consecutive number that Amazon SQS assigns to each message. The length of
	// SequenceNumber is 128 bits. SequenceNumber continues to increase for a
	// particular MessageGroupId .
	SequenceNumber *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSendMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSendMessage{}, middleware.After)
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
	if err = addValidateSendMessageChecksum(stack, options); err != nil {
		return err
	}
	if err = addSendMessageResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendMessage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "SendMessage",
	}
}

type opSendMessageResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSendMessageResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSendMessageResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "sqs"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "sqs"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("sqs")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addSendMessageResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSendMessageResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
