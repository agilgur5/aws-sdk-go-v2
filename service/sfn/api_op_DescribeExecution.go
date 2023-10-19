// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a state machine execution, such as the state machine
// associated with the execution, the execution input and output, and relevant
// execution metadata. Use this API action to return the Map Run Amazon Resource
// Name (ARN) if the execution was dispatched by a Map Run. If you specify a
// version or alias ARN when you call the StartExecution API action,
// DescribeExecution returns that ARN. This operation is eventually consistent. The
// results are best effort and may not reflect very recent updates and changes.
// Executions of an EXPRESS state machinearen't supported by DescribeExecution
// unless a Map Run dispatched them.
func (c *Client) DescribeExecution(ctx context.Context, params *DescribeExecutionInput, optFns ...func(*Options)) (*DescribeExecutionOutput, error) {
	if params == nil {
		params = &DescribeExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExecution", params, optFns, c.addOperationDescribeExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExecutionInput struct {

	// The Amazon Resource Name (ARN) of the execution to describe.
	//
	// This member is required.
	ExecutionArn *string

	noSmithyDocumentSerde
}

type DescribeExecutionOutput struct {

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// This member is required.
	ExecutionArn *string

	// The date the execution is started.
	//
	// This member is required.
	StartDate *time.Time

	// The Amazon Resource Name (ARN) of the executed stated machine.
	//
	// This member is required.
	StateMachineArn *string

	// The current status of the execution.
	//
	// This member is required.
	Status types.ExecutionStatus

	// The cause string if the state machine execution failed.
	Cause *string

	// The error string if the state machine execution failed.
	Error *string

	// The string that contains the JSON input data of the execution. Length
	// constraints apply to the payload size, and are expressed as bytes in UTF-8
	// encoding.
	Input *string

	// Provides details about execution input or output.
	InputDetails *types.CloudWatchEventsExecutionDataDetails

	// The Amazon Resource Name (ARN) that identifies a Map Run, which dispatched this
	// execution.
	MapRunArn *string

	// The name of the execution. A name must not contain:
	//   - white space
	//   - brackets < > { } [ ]
	//   - wildcard characters ? *
	//   - special characters " # % \ ^ | ~ ` $ & , ; : /
	//   - control characters ( U+0000-001F , U+007F-009F )
	// To enable logging with CloudWatch Logs, the name should only contain 0-9, A-Z,
	// a-z, - and _.
	Name *string

	// The JSON output data of the execution. Length constraints apply to the payload
	// size, and are expressed as bytes in UTF-8 encoding. This field is set only if
	// the execution succeeds. If the execution fails, this field is null.
	Output *string

	// Provides details about execution input or output.
	OutputDetails *types.CloudWatchEventsExecutionDataDetails

	// The Amazon Resource Name (ARN) of the state machine alias associated with the
	// execution. The alias ARN is a combination of state machine ARN and the alias
	// name separated by a colon (:). For example, stateMachineARN:PROD . If you start
	// an execution from a StartExecution request with a state machine version ARN,
	// this field will be null.
	StateMachineAliasArn *string

	// The Amazon Resource Name (ARN) of the state machine version associated with the
	// execution. The version ARN is a combination of state machine ARN and the version
	// number separated by a colon (:). For example, stateMachineARN:1 . If you start
	// an execution from a StartExecution request without specifying a state machine
	// version or alias ARN, Step Functions returns a null value.
	StateMachineVersionArn *string

	// If the execution ended, the date the execution stopped.
	StopDate *time.Time

	// The X-Ray trace header that was passed to the execution.
	TraceHeader *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeExecution{}, middleware.After)
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
	if err = addOpDescribeExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExecution(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeExecution",
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

func newServiceMetadataMiddleware_opDescribeExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DescribeExecution",
	}
}
