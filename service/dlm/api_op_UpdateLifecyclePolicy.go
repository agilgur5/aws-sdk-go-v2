// Code generated by smithy-go-codegen DO NOT EDIT.

package dlm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dlm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified lifecycle policy. For more information about updating a
// policy, see Modify lifecycle policies (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/view-modify-delete.html#modify)
// .
func (c *Client) UpdateLifecyclePolicy(ctx context.Context, params *UpdateLifecyclePolicyInput, optFns ...func(*Options)) (*UpdateLifecyclePolicyOutput, error) {
	if params == nil {
		params = &UpdateLifecyclePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLifecyclePolicy", params, optFns, c.addOperationUpdateLifecyclePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLifecyclePolicyInput struct {

	// The identifier of the lifecycle policy.
	//
	// This member is required.
	PolicyId *string

	// A description of the lifecycle policy.
	Description *string

	// The Amazon Resource Name (ARN) of the IAM role used to run the operations
	// specified by the lifecycle policy.
	ExecutionRoleArn *string

	// The configuration of the lifecycle policy. You cannot update the policy type or
	// the resource type.
	PolicyDetails *types.PolicyDetails

	// The desired activation state of the lifecycle policy after creation.
	State types.SettablePolicyStateValues

	noSmithyDocumentSerde
}

type UpdateLifecyclePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLifecyclePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLifecyclePolicy{}, middleware.After)
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
	if err = addOpUpdateLifecyclePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLifecyclePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLifecyclePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dlm",
		OperationName: "UpdateLifecyclePolicy",
	}
}
