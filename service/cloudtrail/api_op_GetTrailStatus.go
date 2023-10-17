// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a JSON-formatted list of information about the specified trail. Fields
// include information on delivery errors, Amazon SNS and Amazon S3 errors, and
// start and stop logging times for each trail. This operation returns trail status
// from a single Region. To return trail status from all Regions, you must call the
// operation on each Region.
func (c *Client) GetTrailStatus(ctx context.Context, params *GetTrailStatusInput, optFns ...func(*Options)) (*GetTrailStatusOutput, error) {
	if params == nil {
		params = &GetTrailStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTrailStatus", params, optFns, c.addOperationGetTrailStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTrailStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The name of a trail about which you want the current status.
type GetTrailStatusInput struct {

	// Specifies the name or the CloudTrail ARN of the trail for which you are
	// requesting status. To get the status of a shadow trail (a replication of the
	// trail in another Region), you must specify its ARN. The following is the format
	// of a trail ARN. arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Returns the objects or data listed below if successful. Otherwise, returns an
// error.
type GetTrailStatusOutput struct {

	// Whether the CloudTrail trail is currently logging Amazon Web Services API calls.
	IsLogging *bool

	// Displays any CloudWatch Logs error that CloudTrail encountered when attempting
	// to deliver logs to CloudWatch Logs.
	LatestCloudWatchLogsDeliveryError *string

	// Displays the most recent date and time when CloudTrail delivered logs to
	// CloudWatch Logs.
	LatestCloudWatchLogsDeliveryTime *time.Time

	// This field is no longer in use.
	LatestDeliveryAttemptSucceeded *string

	// This field is no longer in use.
	LatestDeliveryAttemptTime *string

	// Displays any Amazon S3 error that CloudTrail encountered when attempting to
	// deliver log files to the designated bucket. For more information, see Error
	// Responses (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html)
	// in the Amazon S3 API Reference. This error occurs only when there is a problem
	// with the destination S3 bucket, and does not occur for requests that time out.
	// To resolve the issue, create a new bucket, and then call UpdateTrail to specify
	// the new bucket; or fix the existing objects so that CloudTrail can again write
	// to the bucket.
	LatestDeliveryError *string

	// Specifies the date and time that CloudTrail last delivered log files to an
	// account's Amazon S3 bucket.
	LatestDeliveryTime *time.Time

	// Displays any Amazon S3 error that CloudTrail encountered when attempting to
	// deliver a digest file to the designated bucket. For more information, see Error
	// Responses (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html)
	// in the Amazon S3 API Reference. This error occurs only when there is a problem
	// with the destination S3 bucket, and does not occur for requests that time out.
	// To resolve the issue, create a new bucket, and then call UpdateTrail to specify
	// the new bucket; or fix the existing objects so that CloudTrail can again write
	// to the bucket.
	LatestDigestDeliveryError *string

	// Specifies the date and time that CloudTrail last delivered a digest file to an
	// account's Amazon S3 bucket.
	LatestDigestDeliveryTime *time.Time

	// This field is no longer in use.
	LatestNotificationAttemptSucceeded *string

	// This field is no longer in use.
	LatestNotificationAttemptTime *string

	// Displays any Amazon SNS error that CloudTrail encountered when attempting to
	// send a notification. For more information about Amazon SNS errors, see the
	// Amazon SNS Developer Guide (https://docs.aws.amazon.com/sns/latest/dg/welcome.html)
	// .
	LatestNotificationError *string

	// Specifies the date and time of the most recent Amazon SNS notification that
	// CloudTrail has written a new log file to an account's Amazon S3 bucket.
	LatestNotificationTime *time.Time

	// Specifies the most recent date and time when CloudTrail started recording API
	// calls for an Amazon Web Services account.
	StartLoggingTime *time.Time

	// Specifies the most recent date and time when CloudTrail stopped recording API
	// calls for an Amazon Web Services account.
	StopLoggingTime *time.Time

	// This field is no longer in use.
	TimeLoggingStarted *string

	// This field is no longer in use.
	TimeLoggingStopped *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTrailStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTrailStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTrailStatus{}, middleware.After)
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
	if err = addOpGetTrailStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTrailStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTrailStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "GetTrailStatus",
	}
}
