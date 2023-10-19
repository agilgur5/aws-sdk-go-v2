// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a job execution. Requires permission to access the DeleteJobExecution (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DeleteJobExecution(ctx context.Context, params *DeleteJobExecutionInput, optFns ...func(*Options)) (*DeleteJobExecutionOutput, error) {
	if params == nil {
		params = &DeleteJobExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteJobExecution", params, optFns, c.addOperationDeleteJobExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobExecutionInput struct {

	// The ID of the job execution to be deleted. The executionNumber refers to the
	// execution of a particular job on a particular device. Note that once a job
	// execution is deleted, the executionNumber may be reused by IoT, so be sure you
	// get and use the correct value here.
	//
	// This member is required.
	ExecutionNumber *int64

	// The ID of the job whose execution on a particular device will be deleted.
	//
	// This member is required.
	JobId *string

	// The name of the thing whose job execution will be deleted.
	//
	// This member is required.
	ThingName *string

	// (Optional) When true, you can delete a job execution which is "IN_PROGRESS".
	// Otherwise, you can only delete a job execution which is in a terminal state
	// ("SUCCEEDED", "FAILED", "REJECTED", "REMOVED" or "CANCELED") or an exception
	// will occur. The default is false. Deleting a job execution which is
	// "IN_PROGRESS", will cause the device to be unable to access job information or
	// update the job execution status. Use caution and ensure that the device is able
	// to recover to a valid state.
	Force bool

	// The namespace used to indicate that a job is a customer-managed job. When you
	// specify a value for this parameter, Amazon Web Services IoT Core sends jobs
	// notifications to MQTT topics that contain the value in the following format.
	// $aws/things/THING_NAME/jobs/JOB_ID/notify-namespace-NAMESPACE_ID/ The
	// namespaceId feature is in public preview.
	NamespaceId *string

	noSmithyDocumentSerde
}

type DeleteJobExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteJobExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteJobExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteJobExecution{}, middleware.After)
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
	if err = addOpDeleteJobExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJobExecution(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteJobExecution",
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

func newServiceMetadataMiddleware_opDeleteJobExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "DeleteJobExecution",
	}
}
