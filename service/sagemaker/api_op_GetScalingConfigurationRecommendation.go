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

// Starts an Amazon SageMaker Inference Recommender autoscaling recommendation
// job. Returns recommendations for autoscaling policies that you can apply to your
// SageMaker endpoint.
func (c *Client) GetScalingConfigurationRecommendation(ctx context.Context, params *GetScalingConfigurationRecommendationInput, optFns ...func(*Options)) (*GetScalingConfigurationRecommendationOutput, error) {
	if params == nil {
		params = &GetScalingConfigurationRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetScalingConfigurationRecommendation", params, optFns, c.addOperationGetScalingConfigurationRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetScalingConfigurationRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetScalingConfigurationRecommendationInput struct {

	// The name of a previously completed Inference Recommender job.
	//
	// This member is required.
	InferenceRecommendationsJobName *string

	// The name of an endpoint benchmarked during a previously completed inference
	// recommendation job. This name should come from one of the recommendations
	// returned by the job specified in the InferenceRecommendationsJobName field.
	// Specify either this field or the RecommendationId field.
	EndpointName *string

	// The recommendation ID of a previously completed inference recommendation. This
	// ID should come from one of the recommendations returned by the job specified in
	// the InferenceRecommendationsJobName field. Specify either this field or the
	// EndpointName field.
	RecommendationId *string

	// An object where you specify the anticipated traffic pattern for an endpoint.
	ScalingPolicyObjective *types.ScalingPolicyObjective

	// The percentage of how much utilization you want an instance to use before
	// autoscaling. The default value is 50%.
	TargetCpuUtilizationPerCore *int32

	noSmithyDocumentSerde
}

type GetScalingConfigurationRecommendationOutput struct {

	// An object with the recommended values for you to specify when creating an
	// autoscaling policy.
	DynamicScalingConfiguration *types.DynamicScalingConfiguration

	// The name of an endpoint benchmarked during a previously completed Inference
	// Recommender job.
	EndpointName *string

	// The name of a previously completed Inference Recommender job.
	InferenceRecommendationsJobName *string

	// An object with a list of metrics that were benchmarked during the previously
	// completed Inference Recommender job.
	Metric *types.ScalingPolicyMetric

	// The recommendation ID of a previously completed inference recommendation.
	RecommendationId *string

	// An object representing the anticipated traffic pattern for an endpoint that you
	// specified in the request.
	ScalingPolicyObjective *types.ScalingPolicyObjective

	// The percentage of how much utilization you want an instance to use before
	// autoscaling, which you specified in the request. The default value is 50%.
	TargetCpuUtilizationPerCore *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetScalingConfigurationRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetScalingConfigurationRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetScalingConfigurationRecommendation{}, middleware.After)
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
	if err = addOpGetScalingConfigurationRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetScalingConfigurationRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetScalingConfigurationRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "GetScalingConfigurationRecommendation",
	}
}
