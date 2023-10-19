// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get the properties associated with a Bedrock custom model that you have
// created.For more information, see Custom models (https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html)
// in the Bedrock User Guide.
func (c *Client) GetCustomModel(ctx context.Context, params *GetCustomModelInput, optFns ...func(*Options)) (*GetCustomModelOutput, error) {
	if params == nil {
		params = &GetCustomModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCustomModel", params, optFns, c.addOperationGetCustomModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCustomModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCustomModelInput struct {

	// Name or ARN of the custom model.
	//
	// This member is required.
	ModelIdentifier *string

	noSmithyDocumentSerde
}

type GetCustomModelOutput struct {

	// ARN of the base model.
	//
	// This member is required.
	BaseModelArn *string

	// Creation time of the model.
	//
	// This member is required.
	CreationTime *time.Time

	// Job ARN associated with this model.
	//
	// This member is required.
	JobArn *string

	// ARN associated with this model.
	//
	// This member is required.
	ModelArn *string

	// Model name associated with this model.
	//
	// This member is required.
	ModelName *string

	// Output data configuration associated with this custom model.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// Information about the training dataset.
	//
	// This member is required.
	TrainingDataConfig *types.TrainingDataConfig

	// Hyperparameter values associated with this model.
	HyperParameters map[string]string

	// Job name associated with this model.
	JobName *string

	// The custom model is encrypted at rest using this key.
	ModelKmsKeyArn *string

	// The training metrics from the job creation.
	TrainingMetrics *types.TrainingMetrics

	// Array of up to 10 validators.
	ValidationDataConfig *types.ValidationDataConfig

	// The validation metrics from the job creation.
	ValidationMetrics []types.ValidatorMetric

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCustomModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCustomModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCustomModel{}, middleware.After)
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
	if err = addOpGetCustomModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCustomModel(options.Region), middleware.Before); err != nil {
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
		operation: "GetCustomModel",
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

func newServiceMetadataMiddleware_opGetCustomModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "bedrock",
		OperationName: "GetCustomModel",
	}
}
