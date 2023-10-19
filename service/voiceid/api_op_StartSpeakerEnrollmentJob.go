// Code generated by smithy-go-codegen DO NOT EDIT.

package voiceid

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/voiceid/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a new batch speaker enrollment job using specified details.
func (c *Client) StartSpeakerEnrollmentJob(ctx context.Context, params *StartSpeakerEnrollmentJobInput, optFns ...func(*Options)) (*StartSpeakerEnrollmentJobOutput, error) {
	if params == nil {
		params = &StartSpeakerEnrollmentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSpeakerEnrollmentJob", params, optFns, c.addOperationStartSpeakerEnrollmentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSpeakerEnrollmentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSpeakerEnrollmentJobInput struct {

	// The IAM role Amazon Resource Name (ARN) that grants Voice ID permissions to
	// access customer's buckets to read the input manifest file and write the job
	// output file. Refer to Batch enrollment using audio data from prior calls (https://docs.aws.amazon.com/connect/latest/adminguide/voiceid-batch-enrollment.html)
	// for the permissions needed in this role.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The identifier of the domain that contains the speaker enrollment job and in
	// which the speakers are enrolled.
	//
	// This member is required.
	DomainId *string

	// The input data config containing the S3 location for the input manifest file
	// that contains the list of speaker enrollment requests.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The output data config containing the S3 location where Voice ID writes the job
	// output file; you must also include a KMS key ID to encrypt the file.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// The enrollment config that contains details such as the action to take when a
	// speaker is already enrolled in Voice ID or when a speaker is identified as a
	// fraudster.
	EnrollmentConfig *types.EnrollmentConfig

	// A name for your speaker enrollment job.
	JobName *string

	noSmithyDocumentSerde
}

type StartSpeakerEnrollmentJobOutput struct {

	// Details about the started speaker enrollment job.
	Job *types.SpeakerEnrollmentJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSpeakerEnrollmentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartSpeakerEnrollmentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartSpeakerEnrollmentJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartSpeakerEnrollmentJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSpeakerEnrollmentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSpeakerEnrollmentJob(options.Region), middleware.Before); err != nil {
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
		operation: "StartSpeakerEnrollmentJob",
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

type idempotencyToken_initializeOpStartSpeakerEnrollmentJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSpeakerEnrollmentJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSpeakerEnrollmentJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSpeakerEnrollmentJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSpeakerEnrollmentJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartSpeakerEnrollmentJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartSpeakerEnrollmentJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSpeakerEnrollmentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "voiceid",
		OperationName: "StartSpeakerEnrollmentJob",
	}
}
