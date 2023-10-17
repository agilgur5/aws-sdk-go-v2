// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about an alarm model. If you don't specify a value for
// the alarmModelVersion parameter, the latest version is returned.
func (c *Client) DescribeAlarmModel(ctx context.Context, params *DescribeAlarmModelInput, optFns ...func(*Options)) (*DescribeAlarmModelOutput, error) {
	if params == nil {
		params = &DescribeAlarmModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAlarmModel", params, optFns, c.addOperationDescribeAlarmModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAlarmModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAlarmModelInput struct {

	// The name of the alarm model.
	//
	// This member is required.
	AlarmModelName *string

	// The version of the alarm model.
	AlarmModelVersion *string

	noSmithyDocumentSerde
}

type DescribeAlarmModelOutput struct {

	// Contains the configuration information of alarm state changes.
	AlarmCapabilities *types.AlarmCapabilities

	// Contains information about one or more alarm actions.
	AlarmEventActions *types.AlarmEventActions

	// The ARN of the alarm model. For more information, see Amazon Resource Names
	// (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	AlarmModelArn *string

	// The description of the alarm model.
	AlarmModelDescription *string

	// The name of the alarm model.
	AlarmModelName *string

	// The version of the alarm model.
	AlarmModelVersion *string

	// Contains information about one or more notification actions.
	AlarmNotification *types.AlarmNotification

	// Defines when your alarm is invoked.
	AlarmRule *types.AlarmRule

	// The time the alarm model was created, in the Unix epoch format.
	CreationTime *time.Time

	// An input attribute used as a key to create an alarm. AWS IoT Events routes
	// inputs (https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html)
	// associated with this key to the alarm.
	Key *string

	// The time the alarm model was last updated, in the Unix epoch format.
	LastUpdateTime *time.Time

	// The ARN of the IAM role that allows the alarm to perform actions and access AWS
	// resources. For more information, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	RoleArn *string

	// A non-negative integer that reflects the severity level of the alarm.
	Severity *int32

	// The status of the alarm model. The status can be one of the following values:
	//   - ACTIVE - The alarm model is active and it's ready to evaluate data.
	//   - ACTIVATING - AWS IoT Events is activating your alarm model. Activating an
	//   alarm model can take up to a few minutes.
	//   - INACTIVE - The alarm model is inactive, so it isn't ready to evaluate data.
	//   Check your alarm model information and update the alarm model.
	//   - FAILED - You couldn't create or update the alarm model. Check your alarm
	//   model information and try again.
	Status types.AlarmModelVersionStatus

	// Contains information about the status of the alarm model.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAlarmModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAlarmModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAlarmModel{}, middleware.After)
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
	if err = addOpDescribeAlarmModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAlarmModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAlarmModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "DescribeAlarmModel",
	}
}
