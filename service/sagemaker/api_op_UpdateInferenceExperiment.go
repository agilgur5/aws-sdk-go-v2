// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an inference experiment that you created. The status of the inference
// experiment has to be either Created , Running . For more information on the
// status of an inference experiment, see DescribeInferenceExperiment (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeInferenceExperiment.html)
// .
func (c *Client) UpdateInferenceExperiment(ctx context.Context, params *UpdateInferenceExperimentInput, optFns ...func(*Options)) (*UpdateInferenceExperimentOutput, error) {
	if params == nil {
		params = &UpdateInferenceExperimentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInferenceExperiment", params, optFns, c.addOperationUpdateInferenceExperimentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInferenceExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInferenceExperimentInput struct {

	// The name of the inference experiment to be updated.
	//
	// This member is required.
	Name *string

	// The Amazon S3 location and configuration for storing inference request and
	// response data.
	DataStorageConfig *types.InferenceExperimentDataStorageConfig

	// The description of the inference experiment.
	Description *string

	// An array of ModelVariantConfig objects. There is one for each variant, whose
	// infrastructure configuration you want to update.
	ModelVariants []types.ModelVariantConfig

	// The duration for which the inference experiment will run. If the status of the
	// inference experiment is Created , then you can update both the start and end
	// dates. If the status of the inference experiment is Running , then you can
	// update only the end date.
	Schedule *types.InferenceExperimentSchedule

	// The configuration of ShadowMode inference experiment type. Use this field to
	// specify a production variant which takes all the inference requests, and a
	// shadow variant to which Amazon SageMaker replicates a percentage of the
	// inference requests. For the shadow variant also specify the percentage of
	// requests that Amazon SageMaker replicates.
	ShadowModeConfig *types.ShadowModeConfig

	noSmithyDocumentSerde
}

type UpdateInferenceExperimentOutput struct {

	// The ARN of the updated inference experiment.
	//
	// This member is required.
	InferenceExperimentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateInferenceExperimentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateInferenceExperiment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateInferenceExperiment{}, middleware.After)
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
	if err = addOpUpdateInferenceExperimentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInferenceExperiment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInferenceExperiment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateInferenceExperiment",
	}
}
