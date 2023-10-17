// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Records a WorkflowExecutionCancelRequested event in the currently running
// workflow execution identified by the given domain, workflowId, and runId. This
// logically requests the cancellation of the workflow execution as a whole. It is
// up to the decider to take appropriate actions when it receives an execution
// history with this event. If the runId isn't specified, the
// WorkflowExecutionCancelRequested event is recorded in the history of the current
// open workflow execution with the specified workflowId in the domain. Because
// this action allows the workflow to properly clean up and gracefully close, it
// should be used instead of TerminateWorkflowExecution when possible. Access
// Control You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//   - Use an Action element to allow or deny permission to call this action.
//   - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) RequestCancelWorkflowExecution(ctx context.Context, params *RequestCancelWorkflowExecutionInput, optFns ...func(*Options)) (*RequestCancelWorkflowExecutionOutput, error) {
	if params == nil {
		params = &RequestCancelWorkflowExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RequestCancelWorkflowExecution", params, optFns, c.addOperationRequestCancelWorkflowExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RequestCancelWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestCancelWorkflowExecutionInput struct {

	// The name of the domain containing the workflow execution to cancel.
	//
	// This member is required.
	Domain *string

	// The workflowId of the workflow execution to cancel.
	//
	// This member is required.
	WorkflowId *string

	// The runId of the workflow execution to cancel.
	RunId *string

	noSmithyDocumentSerde
}

type RequestCancelWorkflowExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRequestCancelWorkflowExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRequestCancelWorkflowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRequestCancelWorkflowExecution{}, middleware.After)
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
	if err = addOpRequestCancelWorkflowExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRequestCancelWorkflowExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRequestCancelWorkflowExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "RequestCancelWorkflowExecution",
	}
}
