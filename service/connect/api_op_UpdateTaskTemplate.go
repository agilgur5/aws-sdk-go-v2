// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates details about a specific task template in the specified Amazon Connect
// instance. This operation does not support partial updates. Instead it does a
// full update of template content.
func (c *Client) UpdateTaskTemplate(ctx context.Context, params *UpdateTaskTemplateInput, optFns ...func(*Options)) (*UpdateTaskTemplateOutput, error) {
	if params == nil {
		params = &UpdateTaskTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTaskTemplate", params, optFns, c.addOperationUpdateTaskTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTaskTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTaskTemplateInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// A unique identifier for the task template.
	//
	// This member is required.
	TaskTemplateId *string

	// Constraints that are applicable to the fields listed.
	Constraints *types.TaskTemplateConstraints

	// The identifier of the flow that runs by default when a task is created by
	// referencing this template.
	ContactFlowId *string

	// The default values for fields when a task is created by referencing this
	// template.
	Defaults *types.TaskTemplateDefaults

	// The description of the task template.
	Description *string

	// Fields that are part of the template.
	Fields []types.TaskTemplateField

	// The name of the task template.
	Name *string

	// Marks a template as ACTIVE or INACTIVE for a task to refer to it. Tasks can
	// only be created from ACTIVE templates. If a template is marked as INACTIVE ,
	// then a task that refers to this template cannot be created.
	Status types.TaskTemplateStatus

	noSmithyDocumentSerde
}

type UpdateTaskTemplateOutput struct {

	// The Amazon Resource Name (ARN) for the task template resource.
	Arn *string

	// Constraints that are applicable to the fields listed.
	Constraints *types.TaskTemplateConstraints

	// The identifier of the flow that runs by default when a task is created by
	// referencing this template.
	ContactFlowId *string

	// The timestamp when the task template was created.
	CreatedTime *time.Time

	// The default values for fields when a task is created by referencing this
	// template.
	Defaults *types.TaskTemplateDefaults

	// The description of the task template.
	Description *string

	// Fields that are part of the template.
	Fields []types.TaskTemplateField

	// The identifier of the task template resource.
	Id *string

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	InstanceId *string

	// The timestamp when the task template was last modified.
	LastModifiedTime *time.Time

	// The name of the task template.
	Name *string

	// Marks a template as ACTIVE or INACTIVE for a task to refer to it. Tasks can
	// only be created from ACTIVE templates. If a template is marked as INACTIVE ,
	// then a task that refers to this template cannot be created.
	Status types.TaskTemplateStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTaskTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTaskTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTaskTemplate{}, middleware.After)
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
	if err = addOpUpdateTaskTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTaskTemplate(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateTaskTemplate",
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

func newServiceMetadataMiddleware_opUpdateTaskTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateTaskTemplate",
	}
}
