// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detect drift on a stack set. When CloudFormation performs drift detection on a
// stack set, it performs drift detection on the stack associated with each stack
// instance in the stack set. For more information, see How CloudFormation
// performs drift detection on a stack set (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-drift.html)
// . DetectStackSetDrift returns the OperationId of the stack set drift detection
// operation. Use this operation id with DescribeStackSetOperation to monitor the
// progress of the drift detection operation. The drift detection operation may
// take some time, depending on the number of stack instances included in the stack
// set, in addition to the number of resources included in each stack. Once the
// operation has completed, use the following actions to return drift information:
//   - Use DescribeStackSet to return detailed information about the stack set,
//     including detailed information about the last completed drift operation
//     performed on the stack set. (Information about drift operations that are in
//     progress isn't included.)
//   - Use ListStackInstances to return a list of stack instances belonging to the
//     stack set, including the drift status and last drift time checked of each
//     instance.
//   - Use DescribeStackInstance to return detailed information about a specific
//     stack instance, including its drift status and last drift time checked.
//
// For more information about performing a drift detection operation on a stack
// set, see Detecting unmanaged changes in stack sets (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-drift.html)
// . You can only run a single drift detection operation on a given stack set at
// one time. To stop a drift detection stack set operation, use
// StopStackSetOperation .
func (c *Client) DetectStackSetDrift(ctx context.Context, params *DetectStackSetDriftInput, optFns ...func(*Options)) (*DetectStackSetDriftOutput, error) {
	if params == nil {
		params = &DetectStackSetDriftInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectStackSetDrift", params, optFns, c.addOperationDetectStackSetDriftMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectStackSetDriftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectStackSetDriftInput struct {

	// The name of the stack set on which to perform the drift detection operation.
	//
	// This member is required.
	StackSetName *string

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account. By default, SELF is specified. Use SELF for
	// stack sets with self-managed permissions.
	//   - If you are signed in to the management account, specify SELF .
	//   - If you are signed in to a delegated administrator account, specify
	//   DELEGATED_ADMIN . Your Amazon Web Services account must be registered as a
	//   delegated administrator in the management account. For more information, see
	//   Register a delegated administrator (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	//   in the CloudFormation User Guide.
	CallAs types.CallAs

	// The ID of the stack set operation.
	OperationId *string

	// The user-specified preferences for how CloudFormation performs a stack set
	// operation. For more information about maximum concurrent accounts and failure
	// tolerance, see Stack set operation options (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html#stackset-ops-options)
	// .
	OperationPreferences *types.StackSetOperationPreferences

	noSmithyDocumentSerde
}

type DetectStackSetDriftOutput struct {

	// The ID of the drift detection stack set operation. You can use this operation
	// ID with DescribeStackSetOperation to monitor the progress of the drift
	// detection operation.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectStackSetDriftMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDetectStackSetDrift{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDetectStackSetDrift{}, middleware.After)
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
	if err = addIdempotencyToken_opDetectStackSetDriftMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDetectStackSetDriftValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectStackSetDrift(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDetectStackSetDrift struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDetectStackSetDrift) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDetectStackSetDrift) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DetectStackSetDriftInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DetectStackSetDriftInput ")
	}

	if input.OperationId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.OperationId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDetectStackSetDriftMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDetectStackSetDrift{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDetectStackSetDrift(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DetectStackSetDrift",
	}
}
