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

// Starts a new simulation job batch. The batch is defined using one or more
// SimulationJobRequest objects.
func (c *Client) StartSimulationJobBatch(ctx context.Context, params *StartSimulationJobBatchInput, optFns ...func(*Options)) (*StartSimulationJobBatchOutput, error) {
	if params == nil {
		params = &StartSimulationJobBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSimulationJobBatch", params, optFns, c.addOperationStartSimulationJobBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSimulationJobBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSimulationJobBatchInput struct {

	// A list of simulation job requests to create in the batch.
	//
	// This member is required.
	CreateSimulationJobRequests []types.SimulationJobRequest

	// The batch policy.
	BatchPolicy *types.BatchPolicy

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// A map that contains tag keys and tag values that are attached to the deployment
	// job batch.
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartSimulationJobBatchOutput struct {

	// The Amazon Resource Name (arn) of the batch.
	Arn *string

	// The batch policy.
	BatchPolicy *types.BatchPolicy

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The time, in milliseconds since the epoch, when the simulation job batch was
	// created.
	CreatedAt *time.Time

	// A list of created simulation job request summaries.
	CreatedRequests []types.SimulationJobSummary

	// A list of failed simulation job requests. The request failed to be created into
	// a simulation job. Failed requests do not have a simulation job ID.
	FailedRequests []types.FailedCreateSimulationJobRequest

	// The failure code if the simulation job batch failed.
	FailureCode types.SimulationJobBatchErrorCode

	// The reason the simulation job batch failed.
	FailureReason *string

	// A list of pending simulation job requests. These requests have not yet been
	// created into simulation jobs.
	PendingRequests []types.SimulationJobRequest

	// The status of the simulation job batch. Pending The simulation job batch
	// request is pending. InProgress The simulation job batch is in progress. Failed
	// The simulation job batch failed. One or more simulation job requests could not
	// be completed due to an internal failure (like InternalServiceError ). See
	// failureCode and failureReason for more information. Completed The simulation
	// batch job completed. A batch is complete when (1) there are no pending
	// simulation job requests in the batch and none of the failed simulation job
	// requests are due to InternalServiceError and (2) when all created simulation
	// jobs have reached a terminal state (for example, Completed or Failed ). Canceled
	// The simulation batch job was cancelled. Canceling The simulation batch job is
	// being cancelled. Completing The simulation batch job is completing. TimingOut
	// The simulation job batch is timing out. If a batch timing out, and there are
	// pending requests that were failing due to an internal failure (like
	// InternalServiceError ), the batch status will be Failed . If there are no such
	// failing request, the batch status will be TimedOut . TimedOut The simulation
	// batch job timed out.
	Status types.SimulationJobBatchStatus

	// A map that contains tag keys and tag values that are attached to the deployment
	// job batch.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSimulationJobBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSimulationJobBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSimulationJobBatch{}, middleware.After)
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
	if err = addIdempotencyToken_opStartSimulationJobBatchMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSimulationJobBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSimulationJobBatch(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartSimulationJobBatch struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSimulationJobBatch) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSimulationJobBatch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSimulationJobBatchInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSimulationJobBatchInput ")
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
func addIdempotencyToken_opStartSimulationJobBatchMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartSimulationJobBatch{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSimulationJobBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "StartSimulationJobBatch",
	}
}
