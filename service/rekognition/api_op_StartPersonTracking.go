// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the asynchronous tracking of a person's path in a stored video. Amazon
// Rekognition Video can track the path of people in a video stored in an Amazon S3
// bucket. Use Video to specify the bucket name and the filename of the video.
// StartPersonTracking returns a job identifier ( JobId ) which you use to get the
// results of the operation. When label detection is finished, Amazon Rekognition
// publishes a completion status to the Amazon Simple Notification Service topic
// that you specify in NotificationChannel . To get the results of the person
// detection operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED . If so, call GetPersonTracking and pass the job
// identifier ( JobId ) from the initial call to StartPersonTracking .
func (c *Client) StartPersonTracking(ctx context.Context, params *StartPersonTrackingInput, optFns ...func(*Options)) (*StartPersonTrackingOutput, error) {
	if params == nil {
		params = &StartPersonTrackingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartPersonTracking", params, optFns, c.addOperationStartPersonTrackingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartPersonTrackingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartPersonTrackingInput struct {

	// The video in which you want to detect people. The video must be stored in an
	// Amazon S3 bucket.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartPersonTracking requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// The Amazon SNS topic ARN you want Amazon Rekognition Video to publish the
	// completion status of the people detection operation to. The Amazon SNS topic
	// must have a topic name that begins with AmazonRekognition if you are using the
	// AmazonRekognitionServiceRole permissions policy.
	NotificationChannel *types.NotificationChannel

	noSmithyDocumentSerde
}

type StartPersonTrackingOutput struct {

	// The identifier for the person detection job. Use JobId to identify the job in a
	// subsequent call to GetPersonTracking .
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartPersonTrackingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartPersonTracking{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartPersonTracking{}, middleware.After)
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
	if err = addOpStartPersonTrackingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartPersonTracking(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartPersonTracking(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartPersonTracking",
	}
}
