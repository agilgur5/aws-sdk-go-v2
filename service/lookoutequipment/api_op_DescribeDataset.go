// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides a JSON description of the data in each time series dataset, including
// names, column names, and data types.
func (c *Client) DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) {
	if params == nil {
		params = &DescribeDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataset", params, optFns, c.addOperationDescribeDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetInput struct {

	// The name of the dataset to be described.
	//
	// This member is required.
	DatasetName *string

	noSmithyDocumentSerde
}

type DescribeDatasetOutput struct {

	// Specifies the time the dataset was created in Lookout for Equipment.
	CreatedAt *time.Time

	// Indicates the latest timestamp corresponding to data that was successfully
	// ingested during the most recent ingestion of this particular dataset.
	DataEndTime *time.Time

	// Gives statistics associated with the given dataset for the latest successful
	// associated ingestion job id. These statistics primarily relate to quantifying
	// incorrect data such as MissingCompleteSensorData, MissingSensorData,
	// UnsupportedDateFormats, InsufficientSensorData, and DuplicateTimeStamps.
	DataQualitySummary *types.DataQualitySummary

	// Indicates the earliest timestamp corresponding to data that was successfully
	// ingested during the most recent ingestion of this particular dataset.
	DataStartTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset being described.
	DatasetArn *string

	// The name of the dataset being described.
	DatasetName *string

	// IngestedFilesSummary associated with the given dataset for the latest
	// successful associated ingestion job id.
	IngestedFilesSummary *types.IngestedFilesSummary

	// Specifies the S3 location configuration for the data input for the data
	// ingestion job.
	IngestionInputConfiguration *types.IngestionInputConfiguration

	// Specifies the time the dataset was last updated, if it was.
	LastUpdatedAt *time.Time

	// The Amazon Resource Name (ARN) of the IAM role that you are using for this the
	// data ingestion job.
	RoleArn *string

	// A JSON description of the data that is in each time series dataset, including
	// names, column names, and data types.
	//
	// This value conforms to the media type: application/json
	Schema *string

	// Provides the identifier of the KMS key used to encrypt dataset data by Amazon
	// Lookout for Equipment.
	ServerSideKmsKeyId *string

	// The Amazon Resource Name (ARN) of the source dataset from which the current
	// data being described was imported from.
	SourceDatasetArn *string

	// Indicates the status of the dataset.
	Status types.DatasetStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeDataset{}, middleware.After)
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
	if err = addOpDescribeDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataset(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeDataset",
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

func newServiceMetadataMiddleware_opDescribeDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "DescribeDataset",
	}
}
