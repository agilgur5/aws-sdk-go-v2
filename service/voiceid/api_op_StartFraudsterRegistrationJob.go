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

// Starts a new batch fraudster registration job using provided details.
func (c *Client) StartFraudsterRegistrationJob(ctx context.Context, params *StartFraudsterRegistrationJobInput, optFns ...func(*Options)) (*StartFraudsterRegistrationJobOutput, error) {
	if params == nil {
		params = &StartFraudsterRegistrationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartFraudsterRegistrationJob", params, optFns, c.addOperationStartFraudsterRegistrationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartFraudsterRegistrationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFraudsterRegistrationJobInput struct {

	// The IAM role Amazon Resource Name (ARN) that grants Voice ID permissions to
	// access customer's buckets to read the input manifest file and write the Job
	// output file. Refer to the Create and edit a fraudster watchlist (https://docs.aws.amazon.com/connect/latest/adminguide/voiceid-fraudster-watchlist.html)
	// documentation for the permissions needed in this role.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The identifier of the domain that contains the fraudster registration job and
	// in which the fraudsters are registered.
	//
	// This member is required.
	DomainId *string

	// The input data config containing an S3 URI for the input manifest file that
	// contains the list of fraudster registration requests.
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

	// The name of the new fraudster registration job.
	JobName *string

	// The registration config containing details such as the action to take when a
	// duplicate fraudster is detected, and the similarity threshold to use for
	// detecting a duplicate fraudster.
	RegistrationConfig *types.RegistrationConfig

	noSmithyDocumentSerde
}

type StartFraudsterRegistrationJobOutput struct {

	// Details about the started fraudster registration job.
	Job *types.FraudsterRegistrationJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartFraudsterRegistrationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartFraudsterRegistrationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartFraudsterRegistrationJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartFraudsterRegistrationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartFraudsterRegistrationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartFraudsterRegistrationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartFraudsterRegistrationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartFraudsterRegistrationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartFraudsterRegistrationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartFraudsterRegistrationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartFraudsterRegistrationJobInput ")
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
func addIdempotencyToken_opStartFraudsterRegistrationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartFraudsterRegistrationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartFraudsterRegistrationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "voiceid",
		OperationName: "StartFraudsterRegistrationJob",
	}
}
