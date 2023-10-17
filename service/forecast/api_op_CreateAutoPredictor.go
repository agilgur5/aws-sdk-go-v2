// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Forecast predictor. Amazon Forecast creates predictors with
// AutoPredictor, which involves applying the optimal combination of algorithms to
// each time series in your datasets. You can use CreateAutoPredictor to create
// new predictors or upgrade/retrain existing predictors. Creating new predictors
// The following parameters are required when creating a new predictor:
//   - PredictorName - A unique name for the predictor.
//   - DatasetGroupArn - The ARN of the dataset group used to train the predictor.
//   - ForecastFrequency - The granularity of your forecasts (hourly, daily,
//     weekly, etc).
//   - ForecastHorizon - The number of time-steps that the model predicts. The
//     forecast horizon is also called the prediction length.
//
// When creating a new predictor, do not specify a value for ReferencePredictorArn
// . Upgrading and retraining predictors The following parameters are required when
// retraining or upgrading a predictor:
//   - PredictorName - A unique name for the predictor.
//   - ReferencePredictorArn - The ARN of the predictor to retrain or upgrade.
//
// When upgrading or retraining a predictor, only specify values for the
// ReferencePredictorArn and PredictorName .
func (c *Client) CreateAutoPredictor(ctx context.Context, params *CreateAutoPredictorInput, optFns ...func(*Options)) (*CreateAutoPredictorOutput, error) {
	if params == nil {
		params = &CreateAutoPredictorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoPredictor", params, optFns, c.addOperationCreateAutoPredictorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoPredictorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoPredictorInput struct {

	// A unique name for the predictor
	//
	// This member is required.
	PredictorName *string

	// The data configuration for your dataset group and any additional datasets.
	DataConfig *types.DataConfig

	// An Key Management Service (KMS) key and an Identity and Access Management (IAM)
	// role that Amazon Forecast can assume to access the key. You can specify this
	// optional object in the CreateDataset and CreatePredictor requests.
	EncryptionConfig *types.EncryptionConfig

	// Create an Explainability resource for the predictor.
	ExplainPredictor *bool

	// An array of dimension (field) names that specify how to group the generated
	// forecast. For example, if you are generating forecasts for item sales across all
	// your stores, and your dataset contains a store_id field, you would specify
	// store_id as a dimension to group sales forecasts for each store.
	ForecastDimensions []string

	// The frequency of predictions in a forecast. Valid intervals are an integer
	// followed by Y (Year), M (Month), W (Week), D (Day), H (Hour), and min (Minute).
	// For example, "1D" indicates every day and "15min" indicates every 15 minutes.
	// You cannot specify a value that would overlap with the next larger frequency.
	// That means, for example, you cannot specify a frequency of 60 minutes, because
	// that is equivalent to 1 hour. The valid values for each frequency are the
	// following:
	//   - Minute - 1-59
	//   - Hour - 1-23
	//   - Day - 1-6
	//   - Week - 1-4
	//   - Month - 1-11
	//   - Year - 1
	// Thus, if you want every other week forecasts, specify "2W". Or, if you want
	// quarterly forecasts, you specify "3M". The frequency must be greater than or
	// equal to the TARGET_TIME_SERIES dataset frequency. When a RELATED_TIME_SERIES
	// dataset is provided, the frequency must be equal to the RELATED_TIME_SERIES
	// dataset frequency.
	ForecastFrequency *string

	// The number of time-steps that the model predicts. The forecast horizon is also
	// called the prediction length. The maximum forecast horizon is the lesser of 500
	// time-steps or 1/4 of the TARGET_TIME_SERIES dataset length. If you are
	// retraining an existing AutoPredictor, then the maximum forecast horizon is the
	// lesser of 500 time-steps or 1/3 of the TARGET_TIME_SERIES dataset length. If you
	// are upgrading to an AutoPredictor or retraining an existing AutoPredictor, you
	// cannot update the forecast horizon parameter. You can meet this requirement by
	// providing longer time-series in the dataset.
	ForecastHorizon *int32

	// The forecast types used to train a predictor. You can specify up to five
	// forecast types. Forecast types can be quantiles from 0.01 to 0.99, by increments
	// of 0.01 or higher. You can also specify the mean forecast with mean .
	ForecastTypes []string

	// The configuration details for predictor monitoring. Provide a name for the
	// monitor resource to enable predictor monitoring. Predictor monitoring allows you
	// to see how your predictor's performance changes over time. For more information,
	// see Predictor Monitoring (https://docs.aws.amazon.com/forecast/latest/dg/predictor-monitoring.html)
	// .
	MonitorConfig *types.MonitorConfig

	// The accuracy metric used to optimize the predictor.
	OptimizationMetric types.OptimizationMetric

	// The ARN of the predictor to retrain or upgrade. This parameter is only used
	// when retraining or upgrading a predictor. When creating a new predictor, do not
	// specify a value for this parameter. When upgrading or retraining a predictor,
	// only specify values for the ReferencePredictorArn and PredictorName . The value
	// for PredictorName must be a unique predictor name.
	ReferencePredictorArn *string

	// Optional metadata to help you categorize and organize your predictors. Each tag
	// consists of a key and an optional value, both of which you define. Tag keys and
	// values are case sensitive. The following restrictions apply to tags:
	//   - For each resource, each tag key must be unique and each tag key must have
	//   one value.
	//   - Maximum number of tags per resource: 50.
	//   - Maximum key length: 128 Unicode characters in UTF-8.
	//   - Maximum value length: 256 Unicode characters in UTF-8.
	//   - Accepted characters: all letters and numbers, spaces representable in
	//   UTF-8, and + - = . _ : / @. If your tagging schema is used across other services
	//   and resources, the character restrictions of those services also apply.
	//   - Key prefixes cannot include any upper or lowercase combination of aws: or
	//   AWS: . Values can have this prefix. If a tag value has aws as its prefix but
	//   the key does not, Forecast considers it to be a user tag and will count against
	//   the limit of 50 tags. Tags with only the key prefix of aws do not count
	//   against your tags per resource limit. You cannot edit or delete tag keys with
	//   this prefix.
	Tags []types.Tag

	// The time boundary Forecast uses to align and aggregate any data that doesn't
	// align with your forecast frequency. Provide the unit of time and the time
	// boundary as a key value pair. For more information on specifying a time
	// boundary, see Specifying a Time Boundary (https://docs.aws.amazon.com/forecast/latest/dg/data-aggregation.html#specifying-time-boundary)
	// . If you don't provide a time boundary, Forecast uses a set of Default Time
	// Boundaries (https://docs.aws.amazon.com/forecast/latest/dg/data-aggregation.html#default-time-boundaries)
	// .
	TimeAlignmentBoundary *types.TimeAlignmentBoundary

	noSmithyDocumentSerde
}

type CreateAutoPredictorOutput struct {

	// The Amazon Resource Name (ARN) of the predictor.
	PredictorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAutoPredictorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAutoPredictor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAutoPredictor{}, middleware.After)
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
	if err = addOpCreateAutoPredictorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoPredictor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAutoPredictor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "CreateAutoPredictor",
	}
}
