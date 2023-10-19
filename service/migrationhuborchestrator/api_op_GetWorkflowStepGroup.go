// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get the step group of a migration workflow.
func (c *Client) GetWorkflowStepGroup(ctx context.Context, params *GetWorkflowStepGroupInput, optFns ...func(*Options)) (*GetWorkflowStepGroupOutput, error) {
	if params == nil {
		params = &GetWorkflowStepGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowStepGroup", params, optFns, c.addOperationGetWorkflowStepGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowStepGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowStepGroupInput struct {

	// The ID of the step group.
	//
	// This member is required.
	Id *string

	// The ID of the migration workflow.
	//
	// This member is required.
	WorkflowId *string

	noSmithyDocumentSerde
}

type GetWorkflowStepGroupOutput struct {

	// The time at which the step group was created.
	CreationTime *time.Time

	// The description of the step group.
	Description *string

	// The time at which the step group ended.
	EndTime *time.Time

	// The ID of the step group.
	Id *string

	// The time at which the step group was last modified.
	LastModifiedTime *time.Time

	// The name of the step group.
	Name *string

	// The next step group.
	Next []string

	// The owner of the step group.
	Owner types.Owner

	// The previous step group.
	Previous []string

	// The status of the step group.
	Status types.StepGroupStatus

	// List of AWS services utilized in a migration workflow.
	Tools []types.Tool

	// The ID of the migration workflow.
	WorkflowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowStepGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkflowStepGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkflowStepGroup{}, middleware.After)
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
	if err = addOpGetWorkflowStepGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowStepGroup(options.Region), middleware.Before); err != nil {
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
		operation: "GetWorkflowStepGroup",
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

func newServiceMetadataMiddleware_opGetWorkflowStepGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-orchestrator",
		OperationName: "GetWorkflowStepGroup",
	}
}
