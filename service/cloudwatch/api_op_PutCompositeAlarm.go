// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a composite alarm. When you create a composite alarm, you
// specify a rule expression for the alarm that takes into account the alarm states
// of other alarms that you have created. The composite alarm goes into ALARM state
// only if all conditions of the rule are met. The alarms specified in a composite
// alarm's rule expression can include metric alarms and other composite alarms.
// The rule expression of a composite alarm can include as many as 100 underlying
// alarms. Any single alarm can be included in the rule expressions of as many as
// 150 composite alarms. Using composite alarms can reduce alarm noise. You can
// create multiple metric alarms, and also create a composite alarm and set up
// alerts only for the composite alarm. For example, you could create a composite
// alarm that goes into ALARM state only when more than one of the underlying
// metric alarms are in ALARM state. Currently, the only alarm actions that can be
// taken by composite alarms are notifying SNS topics. It is possible to create a
// loop or cycle of composite alarms, where composite alarm A depends on composite
// alarm B, and composite alarm B also depends on composite alarm A. In this
// scenario, you can't delete any composite alarm that is part of the cycle because
// there is always still a composite alarm that depends on that alarm that you want
// to delete. To get out of such a situation, you must break the cycle by changing
// the rule of one of the composite alarms in the cycle to remove a dependency that
// creates the cycle. The simplest change to make to break a cycle is to change the
// AlarmRule of one of the alarms to false . Additionally, the evaluation of
// composite alarms stops if CloudWatch detects a cycle in the evaluation path.
// When this operation creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA . The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed. For
// a composite alarm, this initial time after creation is the only time that the
// alarm can be in INSUFFICIENT_DATA state. When you update an existing alarm, its
// state is left unchanged, but the update completely overwrites the previous
// configuration of the alarm. To use this operation, you must be signed on with
// the cloudwatch:PutCompositeAlarm permission that is scoped to * . You can't
// create a composite alarms if your cloudwatch:PutCompositeAlarm permission has a
// narrower scope. If you are an IAM user, you must have
// iam:CreateServiceLinkedRole to create a composite alarm that has Systems Manager
// OpsItem actions.
func (c *Client) PutCompositeAlarm(ctx context.Context, params *PutCompositeAlarmInput, optFns ...func(*Options)) (*PutCompositeAlarmOutput, error) {
	if params == nil {
		params = &PutCompositeAlarmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutCompositeAlarm", params, optFns, c.addOperationPutCompositeAlarmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutCompositeAlarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutCompositeAlarmInput struct {

	// The name for the composite alarm. This name must be unique within the Region.
	//
	// This member is required.
	AlarmName *string

	// An expression that specifies which other alarms are to be evaluated to
	// determine this composite alarm's state. For each alarm that you reference, you
	// designate a function that specifies whether that alarm needs to be in ALARM
	// state, OK state, or INSUFFICIENT_DATA state. You can use operators (AND, OR and
	// NOT) to combine multiple functions in a single expression. You can use
	// parenthesis to logically group the functions in your expression. You can use
	// either alarm names or ARNs to reference the other alarms that are to be
	// evaluated. Functions can include the following:
	//   - ALARM("alarm-name or alarm-ARN") is TRUE if the named alarm is in ALARM
	//   state.
	//   - OK("alarm-name or alarm-ARN") is TRUE if the named alarm is in OK state.
	//   - INSUFFICIENT_DATA("alarm-name or alarm-ARN") is TRUE if the named alarm is
	//   in INSUFFICIENT_DATA state.
	//   - TRUE always evaluates to TRUE.
	//   - FALSE always evaluates to FALSE.
	// TRUE and FALSE are useful for testing a complex AlarmRule structure, and for
	// testing your alarm actions. Alarm names specified in AlarmRule can be
	// surrounded with double-quotes ("), but do not have to be. The following are some
	// examples of AlarmRule :
	//   - ALARM(CPUUtilizationTooHigh) AND ALARM(DiskReadOpsTooHigh) specifies that
	//   the composite alarm goes into ALARM state only if both CPUUtilizationTooHigh and
	//   DiskReadOpsTooHigh alarms are in ALARM state.
	//   - ALARM(CPUUtilizationTooHigh) AND NOT ALARM(DeploymentInProgress) specifies
	//   that the alarm goes to ALARM state if CPUUtilizationTooHigh is in ALARM state
	//   and DeploymentInProgress is not in ALARM state. This example reduces alarm noise
	//   during a known deployment window.
	//   - (ALARM(CPUUtilizationTooHigh) OR ALARM(DiskReadOpsTooHigh)) AND
	//   OK(NetworkOutTooHigh) goes into ALARM state if CPUUtilizationTooHigh OR
	//   DiskReadOpsTooHigh is in ALARM state, and if NetworkOutTooHigh is in OK state.
	//   This provides another example of using a composite alarm to prevent noise. This
	//   rule ensures that you are not notified with an alarm action on high CPU or disk
	//   usage if a known network problem is also occurring.
	// The AlarmRule can specify as many as 100 "children" alarms. The AlarmRule
	// expression can have as many as 500 elements. Elements are child alarms, TRUE or
	// FALSE statements, and parentheses.
	//
	// This member is required.
	AlarmRule *string

	// Indicates whether actions should be executed during any changes to the alarm
	// state of the composite alarm. The default is TRUE .
	ActionsEnabled *bool

	// Actions will be suppressed if the suppressor alarm is in the ALARM state.
	// ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an
	// existing alarm.
	ActionsSuppressor *string

	// The maximum time in seconds that the composite alarm waits after suppressor
	// alarm goes out of the ALARM state. After this time, the composite alarm
	// performs its actions. ExtensionPeriod is required only when ActionsSuppressor
	// is specified.
	ActionsSuppressorExtensionPeriod *int32

	// The maximum time in seconds that the composite alarm waits for the suppressor
	// alarm to go into the ALARM state. After this time, the composite alarm performs
	// its actions. WaitPeriod is required only when ActionsSuppressor is specified.
	ActionsSuppressorWaitPeriod *int32

	// The actions to execute when this alarm transitions to the ALARM state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN). Valid
	// Values: arn:aws:sns:region:account-id:sns-topic-name  |
	// arn:aws:ssm:region:account-id:opsitem:severity
	AlarmActions []string

	// The description for the composite alarm.
	AlarmDescription *string

	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource Name
	// (ARN). Valid Values: arn:aws:sns:region:account-id:sns-topic-name
	InsufficientDataActions []string

	// The actions to execute when this alarm transitions to an OK state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN). Valid
	// Values: arn:aws:sns:region:account-id:sns-topic-name
	OKActions []string

	// A list of key-value pairs to associate with the composite alarm. You can
	// associate as many as 50 tags with an alarm. Tags can help you organize and
	// categorize your resources. You can also use them to scope user permissions, by
	// granting a user permission to access or change only resources with certain tag
	// values.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutCompositeAlarmOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutCompositeAlarmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutCompositeAlarm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutCompositeAlarm{}, middleware.After)
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
	if err = addOpPutCompositeAlarmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutCompositeAlarm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutCompositeAlarm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "PutCompositeAlarm",
	}
}
