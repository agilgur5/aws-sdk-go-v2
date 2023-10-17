// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a world export job.
func (c *Client) CreateWorldExportJob(ctx context.Context, params *CreateWorldExportJobInput, optFns ...func(*Options)) (*CreateWorldExportJobOutput, error) {
	if params == nil {
		params = &CreateWorldExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorldExportJob", params, optFns, c.addOperationCreateWorldExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorldExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorldExportJobInput struct {

	// The IAM role that the world export process uses to access the Amazon S3 bucket
	// and put the export.
	//
	// This member is required.
	IamRole *string

	// The output location.
	//
	// This member is required.
	OutputLocation *types.OutputLocation

	// A list of Amazon Resource Names (arns) that correspond to worlds to export.
	//
	// This member is required.
	Worlds []string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// A map that contains tag keys and tag values that are attached to the world
	// export job.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateWorldExportJobOutput struct {

	// The Amazon Resource Name (ARN) of the world export job.
	Arn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The time, in milliseconds since the epoch, when the world export job was
	// created.
	CreatedAt *time.Time

	// The failure code of the world export job if it failed: InternalServiceError
	// Internal service error. LimitExceeded The requested resource exceeds the maximum
	// number allowed, or the number of concurrent stream requests exceeds the maximum
	// number allowed. ResourceNotFound The specified resource could not be found.
	// RequestThrottled The request was throttled. InvalidInput An input parameter in
	// the request is not valid. AllWorldGenerationFailed All of the worlds in the
	// world generation job failed. This can happen if your worldCount is greater than
	// 50 or less than 1. For more information about troubleshooting WorldForge, see
	// Troubleshooting Simulation WorldForge (https://docs.aws.amazon.com/robomaker/latest/dg/troubleshooting-worldforge.html)
	// .
	FailureCode types.WorldExportJobErrorCode

	// The IAM role that the world export process uses to access the Amazon S3 bucket
	// and put the export.
	IamRole *string

	// The output location.
	OutputLocation *types.OutputLocation

	// The status of the world export job. Pending The world export job request is
	// pending. Running The world export job is running. Completed The world export job
	// completed. Failed The world export job failed. See failureCode for more
	// information. Canceled The world export job was cancelled. Canceling The world
	// export job is being cancelled.
	Status types.WorldExportJobStatus

	// A map that contains tag keys and tag values that are attached to the world
	// export job.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorldExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorldExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorldExportJob{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateWorldExportJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWorldExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorldExportJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWorldExportJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWorldExportJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWorldExportJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWorldExportJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWorldExportJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateWorldExportJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWorldExportJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWorldExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateWorldExportJob",
	}
}
