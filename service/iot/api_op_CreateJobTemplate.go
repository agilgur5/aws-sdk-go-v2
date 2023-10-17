// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a job template. Requires permission to access the CreateJobTemplate (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreateJobTemplate(ctx context.Context, params *CreateJobTemplateInput, optFns ...func(*Options)) (*CreateJobTemplateOutput, error) {
	if params == nil {
		params = &CreateJobTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJobTemplate", params, optFns, c.addOperationCreateJobTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobTemplateInput struct {

	// A description of the job document.
	//
	// This member is required.
	Description *string

	// A unique identifier for the job template. We recommend using a UUID.
	// Alpha-numeric characters, "-", and "_" are valid for use here.
	//
	// This member is required.
	JobTemplateId *string

	// The criteria that determine when and how a job abort takes place.
	AbortConfig *types.AbortConfig

	// The package version Amazon Resource Names (ARNs) that are installed on the
	// device when the job successfully completes. Note:The following Length
	// Constraints relates to a single string. Up to five strings are allowed.
	DestinationPackageVersions []string

	// The job document. Required if you don't specify a value for documentSource .
	Document *string

	// An S3 link, or S3 object URL, to the job document. The link is an Amazon S3
	// object URL and is required if you don't specify a value for document . For
	// example, --document-source
	// https://s3.region-code.amazonaws.com/example-firmware/device-firmware.1.0 For
	// more information, see Methods for accessing a bucket (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-bucket-intro.html)
	// .
	DocumentSource *string

	// The ARN of the job to use as the basis for the job template.
	JobArn *string

	// Allows you to create the criteria to retry a job.
	JobExecutionsRetryConfig *types.JobExecutionsRetryConfig

	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig *types.JobExecutionsRolloutConfig

	// Allows you to configure an optional maintenance window for the rollout of a job
	// document to all devices in the target group for a job.
	MaintenanceWindows []types.MaintenanceWindow

	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig *types.PresignedUrlConfig

	// Metadata that can be used to manage the job template.
	Tags []types.Tag

	// Specifies the amount of time each device has to finish its execution of the
	// job. A timer is started when the job execution status is set to IN_PROGRESS . If
	// the job execution status is not set to another terminal state before the timer
	// expires, it will be automatically set to TIMED_OUT .
	TimeoutConfig *types.TimeoutConfig

	noSmithyDocumentSerde
}

type CreateJobTemplateOutput struct {

	// The ARN of the job template.
	JobTemplateArn *string

	// The unique identifier of the job template.
	JobTemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateJobTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJobTemplate{}, middleware.After)
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
	if err = addOpCreateJobTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJobTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateJobTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "CreateJobTemplate",
	}
}
