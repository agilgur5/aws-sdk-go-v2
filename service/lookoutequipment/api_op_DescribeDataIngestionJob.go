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

// Provides information on a specific data ingestion job such as creation time,
// dataset ARN, and status.
func (c *Client) DescribeDataIngestionJob(ctx context.Context, params *DescribeDataIngestionJobInput, optFns ...func(*Options)) (*DescribeDataIngestionJobOutput, error) {
	if params == nil {
		params = &DescribeDataIngestionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataIngestionJob", params, optFns, c.addOperationDescribeDataIngestionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataIngestionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataIngestionJobInput struct {

	// The job ID of the data ingestion job.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeDataIngestionJobOutput struct {

	// The time at which the data ingestion job was created.
	CreatedAt *time.Time

	// Indicates the latest timestamp corresponding to data that was successfully
	// ingested during this specific ingestion job.
	DataEndTime *time.Time

	// Gives statistics about a completed ingestion job. These statistics primarily
	// relate to quantifying incorrect data such as MissingCompleteSensorData,
	// MissingSensorData, UnsupportedDateFormats, InsufficientSensorData, and
	// DuplicateTimeStamps.
	DataQualitySummary *types.DataQualitySummary

	// Indicates the earliest timestamp corresponding to data that was successfully
	// ingested during this specific ingestion job.
	DataStartTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset being used in the data ingestion
	// job.
	DatasetArn *string

	// Specifies the reason for failure when a data ingestion job has failed.
	FailedReason *string

	// Indicates the size of the ingested dataset.
	IngestedDataSize *int64

	// Gives statistics about how many files have been ingested, and which files have
	// not been ingested, for a particular ingestion job.
	IngestedFilesSummary *types.IngestedFilesSummary

	// Specifies the S3 location configuration for the data input for the data
	// ingestion job.
	IngestionInputConfiguration *types.IngestionInputConfiguration

	// Indicates the job ID of the data ingestion job.
	JobId *string

	// The Amazon Resource Name (ARN) of an IAM role with permission to access the
	// data source being ingested.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the source dataset from which the data used
	// for the data ingestion job was imported from.
	SourceDatasetArn *string

	// Indicates the status of the DataIngestionJob operation.
	Status types.IngestionJobStatus

	// Provides details about status of the ingestion job that is currently in
	// progress.
	StatusDetail *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDataIngestionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeDataIngestionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeDataIngestionJob{}, middleware.After)
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
	if err = addOpDescribeDataIngestionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataIngestionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDataIngestionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "DescribeDataIngestionJob",
	}
}
