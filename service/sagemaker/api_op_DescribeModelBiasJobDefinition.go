// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a description of a model bias job definition.
func (c *Client) DescribeModelBiasJobDefinition(ctx context.Context, params *DescribeModelBiasJobDefinitionInput, optFns ...func(*Options)) (*DescribeModelBiasJobDefinitionOutput, error) {
	if params == nil {
		params = &DescribeModelBiasJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModelBiasJobDefinition", params, optFns, c.addOperationDescribeModelBiasJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelBiasJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelBiasJobDefinitionInput struct {

	// The name of the model bias job definition. The name must be unique within an
	// Amazon Web Services Region in the Amazon Web Services account.
	//
	// This member is required.
	JobDefinitionName *string

	noSmithyDocumentSerde
}

type DescribeModelBiasJobDefinitionOutput struct {

	// The time at which the model bias job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model bias job.
	//
	// This member is required.
	JobDefinitionArn *string

	// The name of the bias job definition. The name must be unique within an Amazon
	// Web Services Region in the Amazon Web Services account.
	//
	// This member is required.
	JobDefinitionName *string

	// Identifies the resources to deploy for a monitoring job.
	//
	// This member is required.
	JobResources *types.MonitoringResources

	// Configures the model bias job to run a specified Docker container image.
	//
	// This member is required.
	ModelBiasAppSpecification *types.ModelBiasAppSpecification

	// Inputs for the model bias job.
	//
	// This member is required.
	ModelBiasJobInput *types.ModelBiasJobInput

	// The output configuration for monitoring jobs.
	//
	// This member is required.
	ModelBiasJobOutputConfig *types.MonitoringOutputConfig

	// The Amazon Resource Name (ARN) of the IAM role that has read permission to the
	// input data location and write permission to the output data location in Amazon
	// S3.
	//
	// This member is required.
	RoleArn *string

	// The baseline configuration for a model bias job.
	ModelBiasBaselineConfig *types.ModelBiasBaselineConfig

	// Networking options for a model bias job.
	NetworkConfig *types.MonitoringNetworkConfig

	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition *types.MonitoringStoppingCondition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelBiasJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModelBiasJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModelBiasJobDefinition{}, middleware.After)
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
	if err = addOpDescribeModelBiasJobDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModelBiasJobDefinition(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeModelBiasJobDefinition",
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

func newServiceMetadataMiddleware_opDescribeModelBiasJobDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeModelBiasJobDefinition",
	}
}
