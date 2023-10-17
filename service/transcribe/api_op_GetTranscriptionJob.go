// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the specified transcription job. To view the status
// of the specified transcription job, check the TranscriptionJobStatus field. If
// the status is COMPLETED , the job is finished. You can find the results at the
// location specified in TranscriptFileUri . If the status is FAILED ,
// FailureReason provides details on why your transcription job failed. If you
// enabled content redaction, the redacted transcript can be found at the location
// specified in RedactedTranscriptFileUri . To get a list of your transcription
// jobs, use the operation.
func (c *Client) GetTranscriptionJob(ctx context.Context, params *GetTranscriptionJobInput, optFns ...func(*Options)) (*GetTranscriptionJobOutput, error) {
	if params == nil {
		params = &GetTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTranscriptionJob", params, optFns, c.addOperationGetTranscriptionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTranscriptionJobInput struct {

	// The name of the transcription job you want information about. Job names are
	// case sensitive.
	//
	// This member is required.
	TranscriptionJobName *string

	noSmithyDocumentSerde
}

type GetTranscriptionJobOutput struct {

	// Provides detailed information about the specified transcription job, including
	// job status and, if applicable, failure reason.
	TranscriptionJob *types.TranscriptionJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTranscriptionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTranscriptionJob{}, middleware.After)
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
	if err = addOpGetTranscriptionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTranscriptionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTranscriptionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetTranscriptionJob",
	}
}
