// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels active runs for a flow. You can cancel all of the active runs for a
// flow, or you can cancel specific runs by providing their IDs. You can cancel a
// flow run only when the run is in progress. You can't cancel a run that has
// already completed or failed. You also can't cancel a run that's scheduled to
// occur but hasn't started yet. To prevent a scheduled run, you can deactivate the
// flow with the StopFlow action. You cannot resume a run after you cancel it.
// When you send your request, the status for each run becomes CancelStarted . When
// the cancellation completes, the status becomes Canceled . When you cancel a run,
// you still incur charges for any data that the run already processed before the
// cancellation. If the run had already written some data to the flow destination,
// then that data remains in the destination. If you configured the flow to use a
// batch API (such as the Salesforce Bulk API 2.0), then the run will finish
// reading or writing its entire batch of data after the cancellation. For these
// operations, the data processing charges for Amazon AppFlow apply. For the
// pricing information, see Amazon AppFlow pricing (http://aws.amazon.com/appflow/pricing/)
// .
func (c *Client) CancelFlowExecutions(ctx context.Context, params *CancelFlowExecutionsInput, optFns ...func(*Options)) (*CancelFlowExecutionsOutput, error) {
	if params == nil {
		params = &CancelFlowExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelFlowExecutions", params, optFns, c.addOperationCancelFlowExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelFlowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelFlowExecutionsInput struct {

	// The name of a flow with active runs that you want to cancel.
	//
	// This member is required.
	FlowName *string

	// The ID of each active run to cancel. These runs must belong to the flow you
	// specify in your request. If you omit this parameter, your request ends all
	// active runs that belong to the flow.
	ExecutionIds []string

	noSmithyDocumentSerde
}

type CancelFlowExecutionsOutput struct {

	// The IDs of runs that Amazon AppFlow couldn't cancel. These runs might be
	// ineligible for canceling because they haven't started yet or have already
	// completed.
	InvalidExecutions []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelFlowExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelFlowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelFlowExecutions{}, middleware.After)
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
	if err = addOpCancelFlowExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelFlowExecutions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelFlowExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "CancelFlowExecutions",
	}
}
