// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Neptune ML model training job. See Model training using the
// modeltraining command (https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-modeltraining.html)
// . When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the neptune-db:StartMLModelTrainingJob (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startmlmodeltrainingjob)
// IAM action in that cluster.
func (c *Client) StartMLModelTrainingJob(ctx context.Context, params *StartMLModelTrainingJobInput, optFns ...func(*Options)) (*StartMLModelTrainingJobOutput, error) {
	if params == nil {
		params = &StartMLModelTrainingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMLModelTrainingJob", params, optFns, c.addOperationStartMLModelTrainingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMLModelTrainingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMLModelTrainingJobInput struct {

	// The job ID of the completed data-processing job that has created the data that
	// the training will work with.
	//
	// This member is required.
	DataProcessingJobId *string

	// The location in Amazon S3 where the model artifacts are to be stored.
	//
	// This member is required.
	TrainModelS3Location *string

	// The type of ML instance used in preparing and managing training of ML models.
	// This is a CPU instance chosen based on memory requirements for processing the
	// training data and model.
	BaseProcessingInstanceType *string

	// The configuration for custom model training. This is a JSON object.
	CustomModelTrainingParameters *types.CustomModelTrainingParameters

	// Optimizes the cost of training machine-learning models by using Amazon Elastic
	// Compute Cloud spot instances. The default is False .
	EnableManagedSpotTraining *bool

	// A unique identifier for the new job. The default is An autogenerated UUID.
	Id *string

	// Maximum total number of training jobs to start for the hyperparameter tuning
	// job. The default is 2. Neptune ML automatically tunes the hyperparameters of the
	// machine learning model. To obtain a model that performs well, use at least 10
	// jobs (in other words, set maxHPONumberOfTrainingJobs to 10). In general, the
	// more tuning runs, the better the results.
	MaxHPONumberOfTrainingJobs *int32

	// Maximum number of parallel training jobs to start for the hyperparameter tuning
	// job. The default is 2. The number of parallel jobs you can run is limited by the
	// available resources on your training instance.
	MaxHPOParallelTrainingJobs *int32

	// The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3
	// resources. This must be listed in your DB cluster parameter group or an error
	// will occur.
	NeptuneIamRoleArn *string

	// The job ID of a completed model-training job that you want to update
	// incrementally based on updated data.
	PreviousModelTrainingJobId *string

	// The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt the
	// output of the processing job. The default is none.
	S3OutputEncryptionKMSKey *string

	// The ARN of an IAM role for SageMaker execution.This must be listed in your DB
	// cluster parameter group or an error will occur.
	SagemakerIamRoleArn *string

	// The VPC security group IDs. The default is None.
	SecurityGroupIds []string

	// The IDs of the subnets in the Neptune VPC. The default is None.
	Subnets []string

	// The type of ML instance used for model training. All Neptune ML models support
	// CPU, GPU, and multiGPU training. The default is ml.p3.2xlarge . Choosing the
	// right instance type for training depends on the task type, graph size, and your
	// budget.
	TrainingInstanceType *string

	// The disk volume size of the training instance. Both input data and the output
	// model are stored on disk, so the volume size must be large enough to hold both
	// data sets. The default is 0. If not specified or 0, Neptune ML selects a disk
	// volume size based on the recommendation generated in the data processing step.
	TrainingInstanceVolumeSizeInGB *int32

	// Timeout in seconds for the training job. The default is 86,400 (1 day).
	TrainingTimeOutInSeconds *int32

	// The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt data
	// on the storage volume attached to the ML compute instances that run the training
	// job. The default is None.
	VolumeEncryptionKMSKey *string

	noSmithyDocumentSerde
}

type StartMLModelTrainingJobOutput struct {

	// The ARN of the new model training job.
	Arn *string

	// The model training job creation time, in milliseconds.
	CreationTimeInMillis *int64

	// The unique ID of the new model training job.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMLModelTrainingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMLModelTrainingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMLModelTrainingJob{}, middleware.After)
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
	if err = addStartMLModelTrainingJobResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartMLModelTrainingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMLModelTrainingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMLModelTrainingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "neptune-db",
		OperationName: "StartMLModelTrainingJob",
	}
}

type opStartMLModelTrainingJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartMLModelTrainingJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartMLModelTrainingJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "neptune-db"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "neptune-db"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("neptune-db")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addStartMLModelTrainingJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartMLModelTrainingJobResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
