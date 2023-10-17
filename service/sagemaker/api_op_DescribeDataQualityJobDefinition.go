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

// Gets the details of a data quality monitoring job definition.
func (c *Client) DescribeDataQualityJobDefinition(ctx context.Context, params *DescribeDataQualityJobDefinitionInput, optFns ...func(*Options)) (*DescribeDataQualityJobDefinitionOutput, error) {
	if params == nil {
		params = &DescribeDataQualityJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataQualityJobDefinition", params, optFns, c.addOperationDescribeDataQualityJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataQualityJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataQualityJobDefinitionInput struct {

	// The name of the data quality monitoring job definition to describe.
	//
	// This member is required.
	JobDefinitionName *string

	noSmithyDocumentSerde
}

type DescribeDataQualityJobDefinitionOutput struct {

	// The time that the data quality monitoring job definition was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Information about the container that runs the data quality monitoring job.
	//
	// This member is required.
	DataQualityAppSpecification *types.DataQualityAppSpecification

	// The list of inputs for the data quality monitoring job. Currently endpoints are
	// supported.
	//
	// This member is required.
	DataQualityJobInput *types.DataQualityJobInput

	// The output configuration for monitoring jobs.
	//
	// This member is required.
	DataQualityJobOutputConfig *types.MonitoringOutputConfig

	// The Amazon Resource Name (ARN) of the data quality monitoring job definition.
	//
	// This member is required.
	JobDefinitionArn *string

	// The name of the data quality monitoring job definition.
	//
	// This member is required.
	JobDefinitionName *string

	// Identifies the resources to deploy for a monitoring job.
	//
	// This member is required.
	JobResources *types.MonitoringResources

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	//
	// This member is required.
	RoleArn *string

	// The constraints and baselines for the data quality monitoring job definition.
	DataQualityBaselineConfig *types.DataQualityBaselineConfig

	// The networking configuration for the data quality monitoring job.
	NetworkConfig *types.MonitoringNetworkConfig

	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition *types.MonitoringStoppingCondition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDataQualityJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDataQualityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDataQualityJobDefinition{}, middleware.After)
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
	if err = addOpDescribeDataQualityJobDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataQualityJobDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDataQualityJobDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeDataQualityJobDefinition",
	}
}
