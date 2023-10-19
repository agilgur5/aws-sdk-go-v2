// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a canary. Canaries are scripts that monitor your endpoints and APIs
// from the outside-in. Canaries help you check the availability and latency of
// your web services and troubleshoot anomalies by investigating load time data,
// screenshots of the UI, logs, and metrics. You can set up a canary to run
// continuously or just once. Do not use CreateCanary to modify an existing
// canary. Use UpdateCanary (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_UpdateCanary.html)
// instead. To create canaries, you must have the CloudWatchSyntheticsFullAccess
// policy. If you are creating a new IAM role for the canary, you also need the
// iam:CreateRole , iam:CreatePolicy and iam:AttachRolePolicy permissions. For
// more information, see Necessary Roles and Permissions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Roles)
// . Do not include secrets or proprietary information in your canary names. The
// canary name makes up part of the Amazon Resource Name (ARN) for the canary, and
// the ARN is included in outbound calls over the internet. For more information,
// see Security Considerations for Synthetics Canaries (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html)
// .
func (c *Client) CreateCanary(ctx context.Context, params *CreateCanaryInput, optFns ...func(*Options)) (*CreateCanaryOutput, error) {
	if params == nil {
		params = &CreateCanaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCanary", params, optFns, c.addOperationCreateCanaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCanaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCanaryInput struct {

	// The location in Amazon S3 where Synthetics stores artifacts from the test runs
	// of this canary. Artifacts include the log file, screenshots, and HAR files. The
	// name of the S3 bucket can't include a period (.).
	//
	// This member is required.
	ArtifactS3Location *string

	// A structure that includes the entry point from which the canary should start
	// running your script. If the script is stored in an S3 bucket, the bucket name,
	// key, and version are also included.
	//
	// This member is required.
	Code *types.CanaryCodeInput

	// The ARN of the IAM role to be used to run the canary. This role must already
	// exist, and must include lambda.amazonaws.com as a principal in the trust
	// policy. The role must also have the following permissions:
	//   - s3:PutObject
	//   - s3:GetBucketLocation
	//   - s3:ListAllMyBuckets
	//   - cloudwatch:PutMetricData
	//   - logs:CreateLogGroup
	//   - logs:CreateLogStream
	//   - logs:PutLogEvents
	//
	// This member is required.
	ExecutionRoleArn *string

	// The name for this canary. Be sure to give it a descriptive name that
	// distinguishes it from other canaries in your account. Do not include secrets or
	// proprietary information in your canary names. The canary name makes up part of
	// the canary ARN, and the ARN is included in outbound calls over the internet. For
	// more information, see Security Considerations for Synthetics Canaries (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html)
	// .
	//
	// This member is required.
	Name *string

	// Specifies the runtime version to use for the canary. For a list of valid
	// runtime versions and more information about runtime versions, see Canary
	// Runtime Versions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html)
	// .
	//
	// This member is required.
	RuntimeVersion *string

	// A structure that contains information about how often the canary is to run and
	// when these test runs are to stop.
	//
	// This member is required.
	Schedule *types.CanaryScheduleInput

	// A structure that contains the configuration for canary artifacts, including the
	// encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
	ArtifactConfig *types.ArtifactConfigInput

	// The number of days to retain data about failed runs of this canary. If you omit
	// this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriodInDays *int32

	// A structure that contains the configuration for individual canary runs, such as
	// timeout value and environment variables. The environment variables keys and
	// values are not encrypted. Do not store sensitive information in this field.
	RunConfig *types.CanaryRunConfigInput

	// The number of days to retain data about successful runs of this canary. If you
	// omit this field, the default of 31 days is used. The valid range is 1 to 455
	// days.
	SuccessRetentionPeriodInDays *int32

	// A list of key-value pairs to associate with the canary. You can associate as
	// many as 50 tags with a canary. Tags can help you organize and categorize your
	// resources. You can also use them to scope user permissions, by granting a user
	// permission to access or change only the resources that have certain tag values.
	Tags map[string]string

	// If this canary is to test an endpoint in a VPC, this structure contains
	// information about the subnet and security groups of the VPC endpoint. For more
	// information, see Running a Canary in a VPC (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html)
	// .
	VpcConfig *types.VpcConfigInput

	noSmithyDocumentSerde
}

type CreateCanaryOutput struct {

	// The full details about the canary you have created.
	Canary *types.Canary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCanaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCanary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCanary{}, middleware.After)
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
	if err = addOpCreateCanaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCanary(options.Region), middleware.Before); err != nil {
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
		operation: "CreateCanary",
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

func newServiceMetadataMiddleware_opCreateCanary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "CreateCanary",
	}
}
