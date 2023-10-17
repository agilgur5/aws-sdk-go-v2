// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels the execution of a job for a given thing. Requires permission to access
// the CancelJobExecution (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CancelJobExecution(ctx context.Context, params *CancelJobExecutionInput, optFns ...func(*Options)) (*CancelJobExecutionOutput, error) {
	if params == nil {
		params = &CancelJobExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelJobExecution", params, optFns, c.addOperationCancelJobExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelJobExecutionInput struct {

	// The ID of the job to be canceled.
	//
	// This member is required.
	JobId *string

	// The name of the thing whose execution of the job will be canceled.
	//
	// This member is required.
	ThingName *string

	// (Optional) The expected current version of the job execution. Each time you
	// update the job execution, its version is incremented. If the version of the job
	// execution stored in Jobs does not match, the update is rejected with a
	// VersionMismatch error, and an ErrorResponse that contains the current job
	// execution status data is returned. (This makes it unnecessary to perform a
	// separate DescribeJobExecution request in order to obtain the job execution
	// status data.)
	ExpectedVersion *int64

	// (Optional) If true the job execution will be canceled if it has status
	// IN_PROGRESS or QUEUED, otherwise the job execution will be canceled only if it
	// has status QUEUED. If you attempt to cancel a job execution that is IN_PROGRESS,
	// and you do not set force to true , then an InvalidStateTransitionException will
	// be thrown. The default is false . Canceling a job execution which is
	// "IN_PROGRESS", will cause the device to be unable to update the job execution
	// status. Use caution and ensure that the device is able to recover to a valid
	// state.
	Force bool

	// A collection of name/value pairs that describe the status of the job execution.
	// If not specified, the statusDetails are unchanged. You can specify at most 10
	// name/value pairs.
	StatusDetails map[string]string

	noSmithyDocumentSerde
}

type CancelJobExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelJobExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelJobExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJobExecution{}, middleware.After)
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
	if err = addOpCancelJobExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJobExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelJobExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "CancelJobExecution",
	}
}
