// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides a JSON containing the overall information about a specific machine
// learning model, including model name and ARN, dataset, training and evaluation
// information, status, and so on.
func (c *Client) DescribeModel(ctx context.Context, params *DescribeModelInput, optFns ...func(*Options)) (*DescribeModelOutput, error) {
	if params == nil {
		params = &DescribeModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModel", params, optFns, c.addOperationDescribeModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelInput struct {

	// The name of the machine learning model to be described.
	//
	// This member is required.
	ModelName *string

	noSmithyDocumentSerde
}

func (*DescribeModelInput) operationName() string {
	return "DescribeModel"
}

type DescribeModelOutput struct {

	// Indicates the end time of the inference data that has been accumulated.
	AccumulatedInferenceDataEndTime *time.Time

	// Indicates the start time of the inference data that has been accumulated.
	AccumulatedInferenceDataStartTime *time.Time

	// The name of the model version used by the inference schedular when running a
	// scheduled inference execution.
	ActiveModelVersion *int64

	// The Amazon Resource Name (ARN) of the model version used by the inference
	// scheduler when running a scheduled inference execution.
	ActiveModelVersionArn *string

	// Indicates the time and date at which the machine learning model was created.
	CreatedAt *time.Time

	// The configuration is the TargetSamplingRate , which is the sampling rate of the
	// data after post processing by Amazon Lookout for Equipment. For example, if you
	// provide data that has been collected at a 1 second level and you want the system
	// to resample the data at a 1 minute rate before training, the TargetSamplingRate
	// is 1 minute. When providing a value for the TargetSamplingRate , you must attach
	// the prefix "PT" to the rate you want. The value for a 1 second rate is therefore
	// PT1S, the value for a 15 minute rate is PT15M, and the value for a 1 hour rate
	// is PT1H
	DataPreProcessingConfiguration *types.DataPreProcessingConfiguration

	// The Amazon Resouce Name (ARN) of the dataset used to create the machine
	// learning model being described.
	DatasetArn *string

	// The name of the dataset being used by the machine learning being described.
	DatasetName *string

	// Indicates the time reference in the dataset that was used to end the subset of
	// evaluation data for the machine learning model.
	EvaluationDataEndTime *time.Time

	// Indicates the time reference in the dataset that was used to begin the subset
	// of evaluation data for the machine learning model.
	EvaluationDataStartTime *time.Time

	// If the training of the machine learning model failed, this indicates the reason
	// for that failure.
	FailedReason *string

	// The date and time when the import job was completed. This field appears if the
	// active model version was imported.
	ImportJobEndTime *time.Time

	// The date and time when the import job was started. This field appears if the
	// active model version was imported.
	ImportJobStartTime *time.Time

	// Specifies configuration information about the labels input, including its S3
	// location.
	LabelsInputConfiguration *types.LabelsInputConfiguration

	// Indicates the last time the machine learning model was updated. The type of
	// update is not specified.
	LastUpdatedTime *time.Time

	// Indicates the number of days of data used in the most recent scheduled
	// retraining run.
	LatestScheduledRetrainingAvailableDataInDays *int32

	// If the model version was generated by retraining and the training failed, this
	// indicates the reason for that failure.
	LatestScheduledRetrainingFailedReason *string

	// Indicates the most recent model version that was generated by retraining.
	LatestScheduledRetrainingModelVersion *int64

	// Indicates the start time of the most recent scheduled retraining run.
	LatestScheduledRetrainingStartTime *time.Time

	// Indicates the status of the most recent scheduled retraining run.
	LatestScheduledRetrainingStatus types.ModelVersionStatus

	// The Amazon Resource Name (ARN) of the machine learning model being described.
	ModelArn *string

	// The Model Metrics show an aggregated summary of the model's performance within
	// the evaluation time range. This is the JSON content of the metrics created when
	// evaluating the model.
	//
	// This value conforms to the media type: application/json
	ModelMetrics *string

	// The name of the machine learning model being described.
	ModelName *string

	// The date the active model version was activated.
	ModelVersionActivatedAt *time.Time

	// Indicates the date and time that the next scheduled retraining run will start
	// on. Lookout for Equipment truncates the time you provide to the nearest UTC day.
	NextScheduledRetrainingStartDate *time.Time

	// Indicates that the asset associated with this sensor has been shut off. As long
	// as this condition is met, Lookout for Equipment will not use data from this
	// asset for training, evaluation, or inference.
	OffCondition *string

	// The model version that was set as the active model version prior to the current
	// active model version.
	PreviousActiveModelVersion *int64

	// The ARN of the model version that was set as the active model version prior to
	// the current active model version.
	PreviousActiveModelVersionArn *string

	// The date and time when the previous active model version was activated.
	PreviousModelVersionActivatedAt *time.Time

	// If the model version was retrained, this field shows a summary of the
	// performance of the prior model on the new training range. You can use the
	// information in this JSON-formatted object to compare the new model version and
	// the prior model version.
	//
	// This value conforms to the media type: application/json
	PriorModelMetrics *string

	// Indicates the status of the retraining scheduler.
	RetrainingSchedulerStatus types.RetrainingSchedulerStatus

	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source for the machine learning model being described.
	RoleArn *string

	// A JSON description of the data that is in each time series dataset, including
	// names, column names, and data types.
	//
	// This value conforms to the media type: application/json
	Schema *string

	// Provides the identifier of the KMS key used to encrypt model data by Amazon
	// Lookout for Equipment.
	ServerSideKmsKeyId *string

	// The Amazon Resource Name (ARN) of the source model version. This field appears
	// if the active model version was imported.
	SourceModelVersionArn *string

	// Specifies the current status of the model being described. Status describes the
	// status of the most recent action of the model.
	Status types.ModelStatus

	// Indicates the time reference in the dataset that was used to end the subset of
	// training data for the machine learning model.
	TrainingDataEndTime *time.Time

	// Indicates the time reference in the dataset that was used to begin the subset
	// of training data for the machine learning model.
	TrainingDataStartTime *time.Time

	// Indicates the time at which the training of the machine learning model was
	// completed.
	TrainingExecutionEndTime *time.Time

	// Indicates the time at which the training of the machine learning model began.
	TrainingExecutionStartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeModel{}, middleware.After)
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
	if err = addDescribeModelResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "DescribeModel",
	}
}

type opDescribeModelResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeModelResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeModelResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "lookoutequipment"
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
				signingName = "lookoutequipment"
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
				v4aScheme.SigningName = aws.String("lookoutequipment")
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

func addDescribeModelResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeModelResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
