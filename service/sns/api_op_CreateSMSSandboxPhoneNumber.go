// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a destination phone number to an Amazon Web Services account in the SMS
// sandbox and sends a one-time password (OTP) to that phone number. When you start
// using Amazon SNS to send SMS messages, your Amazon Web Services account is in
// the SMS sandbox. The SMS sandbox provides a safe environment for you to try
// Amazon SNS features without risking your reputation as an SMS sender. While your
// Amazon Web Services account is in the SMS sandbox, you can use all of the
// features of Amazon SNS. However, you can send SMS messages only to verified
// destination phone numbers. For more information, including how to move out of
// the sandbox to send messages without restrictions, see SMS sandbox (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html)
// in the Amazon SNS Developer Guide.
func (c *Client) CreateSMSSandboxPhoneNumber(ctx context.Context, params *CreateSMSSandboxPhoneNumberInput, optFns ...func(*Options)) (*CreateSMSSandboxPhoneNumberOutput, error) {
	if params == nil {
		params = &CreateSMSSandboxPhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSMSSandboxPhoneNumber", params, optFns, c.addOperationCreateSMSSandboxPhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSMSSandboxPhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSMSSandboxPhoneNumberInput struct {

	// The destination phone number to verify. On verification, Amazon SNS adds this
	// phone number to the list of verified phone numbers that you can send SMS
	// messages to.
	//
	// This member is required.
	PhoneNumber *string

	// The language to use for sending the OTP. The default value is en-US .
	LanguageCode types.LanguageCodeString

	noSmithyDocumentSerde
}

type CreateSMSSandboxPhoneNumberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSMSSandboxPhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateSMSSandboxPhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateSMSSandboxPhoneNumber{}, middleware.After)
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
	if err = addOpCreateSMSSandboxPhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSMSSandboxPhoneNumber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSMSSandboxPhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "CreateSMSSandboxPhoneNumber",
	}
}
