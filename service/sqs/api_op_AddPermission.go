// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a permission to a queue for a specific principal (https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P)
// . This allows sharing access to the queue. When you create a queue, you have
// full control access rights for the queue. Only you, the owner of the queue, can
// grant or deny permissions to the queue. For more information about these
// permissions, see Allow Developers to Write Messages to a Shared Queue (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
// in the Amazon SQS Developer Guide.
//   - AddPermission generates a policy for you. You can use SetQueueAttributes to
//     upload your policy. For more information, see Using Custom Policies with the
//     Amazon SQS Access Policy Language (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html)
//     in the Amazon SQS Developer Guide.
//   - An Amazon SQS policy can have a maximum of seven actions per statement.
//   - To remove the ability to change queue permissions, you must deny permission
//     to the AddPermission , RemovePermission , and SetQueueAttributes actions in
//     your IAM policy.
//   - Amazon SQS AddPermission does not support adding a non-account principal.
//
// Cross-account permissions don't apply to this action. For more information, see
// Grant cross-account permissions to a role and a username (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon SQS Developer Guide.
func (c *Client) AddPermission(ctx context.Context, params *AddPermissionInput, optFns ...func(*Options)) (*AddPermissionOutput, error) {
	if params == nil {
		params = &AddPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddPermission", params, optFns, c.addOperationAddPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddPermissionInput struct {

	// The Amazon Web Services account numbers of the principals (https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P)
	// who are to receive permission. For information about locating the Amazon Web
	// Services account identification, see Your Amazon Web Services Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-making-api-requests.html#sqs-api-request-authentication)
	// in the Amazon SQS Developer Guide.
	//
	// This member is required.
	AWSAccountIds []string

	// The action the client wants to allow for the specified principal. Valid values:
	// the name of any action or * . For more information about these actions, see
	// Overview of Managing Access Permissions to Your Amazon Simple Queue Service
	// Resource (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-overview-of-managing-access.html)
	// in the Amazon SQS Developer Guide. Specifying SendMessage , DeleteMessage , or
	// ChangeMessageVisibility for ActionName.n also grants permissions for the
	// corresponding batch versions of those actions: SendMessageBatch ,
	// DeleteMessageBatch , and ChangeMessageVisibilityBatch .
	//
	// This member is required.
	Actions []string

	// The unique identification of the permission you're setting (for example,
	// AliceSendMessage ). Maximum 80 characters. Allowed characters include
	// alphanumeric characters, hyphens ( - ), and underscores ( _ ).
	//
	// This member is required.
	Label *string

	// The URL of the Amazon SQS queue to which permissions are added. Queue URLs and
	// names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	noSmithyDocumentSerde
}

type AddPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddPermission{}, middleware.After)
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
	if err = addOpAddPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddPermission(options.Region), middleware.Before); err != nil {
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
		operation: "AddPermission",
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

func newServiceMetadataMiddleware_opAddPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "AddPermission",
	}
}
