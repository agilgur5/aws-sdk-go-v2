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

// Starts a hyperparameter tuning job. A hyperparameter tuning job finds the best
// version of a model by running many training jobs on your dataset using the
// algorithm you choose and values for hyperparameters within ranges that you
// specify. It then chooses the hyperparameter values that result in a model that
// performs the best, as measured by an objective metric that you choose. A
// hyperparameter tuning job automatically creates Amazon SageMaker experiments,
// trials, and trial components for each training job that it runs. You can view
// these entities in Amazon SageMaker Studio. For more information, see View
// Experiments, Trials, and Trial Components (https://docs.aws.amazon.com/sagemaker/latest/dg/experiments-view-compare.html#experiments-view)
// . Do not include any security-sensitive information including account access
// IDs, secrets or tokens in any hyperparameter field. If the use of
// security-sensitive credentials are detected, SageMaker will reject your training
// job request and return an exception error.
func (c *Client) CreateHyperParameterTuningJob(ctx context.Context, params *CreateHyperParameterTuningJobInput, optFns ...func(*Options)) (*CreateHyperParameterTuningJobOutput, error) {
	if params == nil {
		params = &CreateHyperParameterTuningJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHyperParameterTuningJob", params, optFns, c.addOperationCreateHyperParameterTuningJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHyperParameterTuningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHyperParameterTuningJobInput struct {

	// The HyperParameterTuningJobConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html)
	// object that describes the tuning job, including the search strategy, the
	// objective metric used to evaluate training jobs, ranges of parameters to search,
	// and resource limits for the tuning job. For more information, see How
	// Hyperparameter Tuning Works (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html)
	// .
	//
	// This member is required.
	HyperParameterTuningJobConfig *types.HyperParameterTuningJobConfig

	// The name of the tuning job. This name is the prefix for the names of all
	// training jobs that this tuning job launches. The name must be unique within the
	// same Amazon Web Services account and Amazon Web Services Region. The name must
	// have 1 to 32 characters. Valid characters are a-z, A-Z, 0-9, and : + = @ _ % -
	// (hyphen). The name is not case sensitive.
	//
	// This member is required.
	HyperParameterTuningJobName *string

	// Configures SageMaker Automatic model tuning (AMT) to automatically find optimal
	// parameters for the following fields:
	//   - ParameterRanges (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html#sagemaker-Type-HyperParameterTuningJobConfig-ParameterRanges)
	//   : The names and ranges of parameters that a hyperparameter tuning job can
	//   optimize.
	//   - ResourceLimits (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ResourceLimits.html)
	//   : The maximum resources that can be used for a training job. These resources
	//   include the maximum number of training jobs, the maximum runtime of a tuning
	//   job, and the maximum number of training jobs to run at the same time.
	//   - TrainingJobEarlyStoppingType (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html#sagemaker-Type-HyperParameterTuningJobConfig-TrainingJobEarlyStoppingType)
	//   : A flag that specifies whether or not to use early stopping for training jobs
	//   launched by a hyperparameter tuning job.
	//   - RetryStrategy (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html#sagemaker-Type-HyperParameterTrainingJobDefinition-RetryStrategy)
	//   : The number of times to retry a training job.
	//   - Strategy (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html)
	//   : Specifies how hyperparameter tuning chooses the combinations of hyperparameter
	//   values to use for the training jobs that it launches.
	//   - ConvergenceDetected (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ConvergenceDetected.html)
	//   : A flag to indicate that Automatic model tuning (AMT) has detected model
	//   convergence.
	Autotune *types.Autotune

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// . Tags that you specify for the tuning job are also added to all training jobs
	// that the tuning job launches.
	Tags []types.Tag

	// The HyperParameterTrainingJobDefinition (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html)
	// object that describes the training jobs that this tuning job launches, including
	// static hyperparameters, input data configuration, output data configuration,
	// resource configuration, and stopping condition.
	TrainingJobDefinition *types.HyperParameterTrainingJobDefinition

	// A list of the HyperParameterTrainingJobDefinition (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html)
	// objects launched for this tuning job.
	TrainingJobDefinitions []types.HyperParameterTrainingJobDefinition

	// Specifies the configuration for starting the hyperparameter tuning job using
	// one or more previous tuning jobs as a starting point. The results of previous
	// tuning jobs are used to inform which combinations of hyperparameters to search
	// over in the new tuning job. All training jobs launched by the new hyperparameter
	// tuning job are evaluated by using the objective metric. If you specify
	// IDENTICAL_DATA_AND_ALGORITHM as the WarmStartType value for the warm start
	// configuration, the training job that performs the best in the new tuning job is
	// compared to the best training jobs from the parent tuning jobs. From these, the
	// training job that performs the best as measured by the objective metric is
	// returned as the overall best training job. All training jobs launched by parent
	// hyperparameter tuning jobs and the new hyperparameter tuning jobs count against
	// the limit of training jobs for the tuning job.
	WarmStartConfig *types.HyperParameterTuningJobWarmStartConfig

	noSmithyDocumentSerde
}

type CreateHyperParameterTuningJobOutput struct {

	// The Amazon Resource Name (ARN) of the tuning job. SageMaker assigns an ARN to a
	// hyperparameter tuning job when you create it.
	//
	// This member is required.
	HyperParameterTuningJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHyperParameterTuningJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHyperParameterTuningJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHyperParameterTuningJob{}, middleware.After)
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
	if err = addOpCreateHyperParameterTuningJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHyperParameterTuningJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHyperParameterTuningJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateHyperParameterTuningJob",
	}
}
