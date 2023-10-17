// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops transcription for the specified meetingId . For more information, refer to
// Using Amazon Chime SDK live transcription  (https://docs.aws.amazon.com/chime-sdk/latest/dg/meeting-transcription.html)
// in the Amazon Chime SDK Developer Guide. Amazon Chime SDK live transcription is
// powered by Amazon Transcribe. Use of Amazon Transcribe is subject to the AWS
// Service Terms (https://aws.amazon.com/service-terms/) , including the terms
// specific to the AWS Machine Learning and Artificial Intelligence Services.
func (c *Client) StopMeetingTranscription(ctx context.Context, params *StopMeetingTranscriptionInput, optFns ...func(*Options)) (*StopMeetingTranscriptionOutput, error) {
	if params == nil {
		params = &StopMeetingTranscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopMeetingTranscription", params, optFns, c.addOperationStopMeetingTranscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopMeetingTranscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopMeetingTranscriptionInput struct {

	// The unique ID of the meeting for which you stop transcription.
	//
	// This member is required.
	MeetingId *string

	noSmithyDocumentSerde
}

type StopMeetingTranscriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopMeetingTranscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopMeetingTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopMeetingTranscription{}, middleware.After)
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
	if err = addOpStopMeetingTranscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopMeetingTranscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopMeetingTranscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "StopMeetingTranscription",
	}
}
