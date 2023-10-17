// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a job in an Batch job queue. Jobs that are in the SUBMITTED or PENDING
// are canceled. A job in RUNNABLE remains in RUNNABLE until it reaches the head
// of the job queue. Then the job status is updated to FAILED . A PENDING job is
// canceled after all dependency jobs are completed. Therefore, it may take longer
// than expected to cancel a job in PENDING status. When you try to cancel an
// array parent job in PENDING , Batch attempts to cancel all child jobs. The array
// parent job is canceled when all child jobs are completed. Jobs that progressed
// to the STARTING or RUNNING state aren't canceled. However, the API operation
// still succeeds, even if no job is canceled. These jobs must be terminated with
// the TerminateJob operation.
func (c *Client) CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) {
	if params == nil {
		params = &CancelJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelJob", params, optFns, c.addOperationCancelJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CancelJob .
type CancelJobInput struct {

	// The Batch job ID of the job to cancel.
	//
	// This member is required.
	JobId *string

	// A message to attach to the job that explains the reason for canceling it. This
	// message is returned by future DescribeJobs operations on the job. This message
	// is also recorded in the Batch activity logs.
	//
	// This member is required.
	Reason *string

	noSmithyDocumentSerde
}

type CancelJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJob{}, middleware.After)
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
	if err = addOpCancelJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CancelJob",
	}
}
