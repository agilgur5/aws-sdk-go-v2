// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an account setting. Account settings are set on a per-Region basis. If
// you change the root user account setting, the default settings are reset for
// users and roles that do not have specified individual account settings. For more
// information, see Account Settings (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html)
// in the Amazon Elastic Container Service Developer Guide. When you specify
// serviceLongArnFormat , taskLongArnFormat , or containerInstanceLongArnFormat ,
// the Amazon Resource Name (ARN) and resource ID format of the resource type for a
// specified user, role, or the root user for an account is affected. The opt-in
// and opt-out account setting must be set for each Amazon ECS resource separately.
// The ARN and resource ID format of a resource is defined by the opt-in status of
// the user or role that created the resource. You must turn on this setting to use
// Amazon ECS features such as resource tagging. When you specify awsvpcTrunking ,
// the elastic network interface (ENI) limit for any new container instances that
// support the feature is changed. If awsvpcTrunking is turned on, any new
// container instances that support the feature are launched have the increased ENI
// limits available to them. For more information, see Elastic Network Interface
// Trunking (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-eni.html)
// in the Amazon Elastic Container Service Developer Guide. When you specify
// containerInsights , the default setting indicating whether Amazon Web Services
// CloudWatch Container Insights is turned on for your clusters is changed. If
// containerInsights is turned on, any new clusters that are created will have
// Container Insights turned on unless you disable it during cluster creation. For
// more information, see CloudWatch Container Insights (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-container-insights.html)
// in the Amazon Elastic Container Service Developer Guide. Amazon ECS is
// introducing tagging authorization for resource creation. Users must have
// permissions for actions that create the resource, such as ecsCreateCluster . If
// tags are specified when you create a resource, Amazon Web Services performs
// additional authorization to verify if users or roles have permissions to create
// tags. Therefore, you must grant explicit permissions to use the ecs:TagResource
// action. For more information, see Grant permission to tag resources on creation (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/supported-iam-actions-tagging.html)
// in the Amazon ECS Developer Guide. When Amazon Web Services determines that a
// security or infrastructure update is needed for an Amazon ECS task hosted on
// Fargate, the tasks need to be stopped and new tasks launched to replace them.
// Use fargateTaskRetirementWaitPeriod to configure the wait time to retire a
// Fargate task. For information about the Fargate tasks maintenance, see Amazon
// Web Services Fargate task maintenance (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-maintenance.html)
// in the Amazon ECS Developer Guide.
func (c *Client) PutAccountSetting(ctx context.Context, params *PutAccountSettingInput, optFns ...func(*Options)) (*PutAccountSettingOutput, error) {
	if params == nil {
		params = &PutAccountSettingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountSetting", params, optFns, c.addOperationPutAccountSettingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountSettingInput struct {

	// The Amazon ECS resource name for which to modify the account setting. If you
	// specify serviceLongArnFormat , the ARN for your Amazon ECS services is affected.
	// If you specify taskLongArnFormat , the ARN and resource ID for your Amazon ECS
	// tasks is affected. If you specify containerInstanceLongArnFormat , the ARN and
	// resource ID for your Amazon ECS container instances is affected. If you specify
	// awsvpcTrunking , the elastic network interface (ENI) limit for your Amazon ECS
	// container instances is affected. If you specify containerInsights , the default
	// setting for Amazon Web Services CloudWatch Container Insights for your clusters
	// is affected. If you specify fargateFIPSMode , Fargate FIPS 140 compliance is
	// affected. If you specify tagResourceAuthorization , the opt-in option for
	// tagging resources on creation is affected. For information about the opt-in
	// timeline, see Tagging authorization timeline (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#tag-resources)
	// in the Amazon ECS Developer Guide. If you specify
	// fargateTaskRetirementWaitPeriod , the wait time to retire a Fargate task is
	// affected.
	//
	// This member is required.
	Name types.SettingName

	// The account setting value for the specified principal ARN. Accepted values are
	// enabled , disabled , on , and off . When you specify
	// fargateTaskRetirementWaitPeriod for the name , the following are the valid
	// values:
	//   - 0 - Amazon Web Services sends the notification, and immediately retires the
	//   affected tasks.
	//   - 7 - Amazon Web Services sends the notification, and waits 7 calendar days to
	//   retire the tasks.
	//   - 14 - Amazon Web Services sends the notification, and waits 14 calendar days
	//   to retire the tasks.
	//
	// This member is required.
	Value *string

	// The ARN of the principal, which can be a user, role, or the root user. If you
	// specify the root user, it modifies the account setting for all users, roles, and
	// the root user of the account unless a user or role explicitly overrides these
	// settings. If this field is omitted, the setting is changed only for the
	// authenticated user. You must use the root user when you set the Fargate wait
	// time ( fargateTaskRetirementWaitPeriod ). Federated users assume the account
	// setting of the root user and can't have explicit account settings set for them.
	PrincipalArn *string

	noSmithyDocumentSerde
}

type PutAccountSettingOutput struct {

	// The current account setting for a resource.
	Setting *types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccountSettingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccountSetting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccountSetting{}, middleware.After)
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
	if err = addOpPutAccountSettingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountSetting(options.Region), middleware.Before); err != nil {
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
		operation: "PutAccountSetting",
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

func newServiceMetadataMiddleware_opPutAccountSetting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "PutAccountSetting",
	}
}
