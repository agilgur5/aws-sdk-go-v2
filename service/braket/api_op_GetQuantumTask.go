// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the specified quantum task.
func (c *Client) GetQuantumTask(ctx context.Context, params *GetQuantumTaskInput, optFns ...func(*Options)) (*GetQuantumTaskOutput, error) {
	if params == nil {
		params = &GetQuantumTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQuantumTask", params, optFns, c.addOperationGetQuantumTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQuantumTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQuantumTaskInput struct {

	// the ARN of the task to retrieve.
	//
	// This member is required.
	QuantumTaskArn *string

	// A list of attributes to return information for.
	AdditionalAttributeNames []types.QuantumTaskAdditionalAttributeName

	noSmithyDocumentSerde
}

type GetQuantumTaskOutput struct {

	// The time at which the task was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The ARN of the device the task was run on.
	//
	// This member is required.
	DeviceArn *string

	// The parameters for the device on which the task ran.
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	DeviceParameters *string

	// The S3 bucket where task results are stored.
	//
	// This member is required.
	OutputS3Bucket *string

	// The folder in the S3 bucket where task results are stored.
	//
	// This member is required.
	OutputS3Directory *string

	// The ARN of the task.
	//
	// This member is required.
	QuantumTaskArn *string

	// The number of shots used in the task.
	//
	// This member is required.
	Shots *int64

	// The status of the task.
	//
	// This member is required.
	Status types.QuantumTaskStatus

	// The time at which the task ended.
	EndedAt *time.Time

	// The reason that a task failed.
	FailureReason *string

	// The ARN of the Amazon Braket job associated with the quantum task.
	JobArn *string

	// Queue information for the requested quantum task. Only returned if QueueInfo is
	// specified in the additionalAttributeNames" field in the GetQuantumTask API
	// request.
	QueueInfo *types.QuantumTaskQueueInfo

	// The tags that belong to this task.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQuantumTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetQuantumTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetQuantumTask{}, middleware.After)
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
	if err = addOpGetQuantumTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQuantumTask(options.Region), middleware.Before); err != nil {
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
		operation: "GetQuantumTask",
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

func newServiceMetadataMiddleware_opGetQuantumTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "braket",
		OperationName: "GetQuantumTask",
	}
}
