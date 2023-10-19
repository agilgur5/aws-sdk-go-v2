// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Redshift event notification subscription. This action
// requires an ARN (Amazon Resource Name) of an Amazon SNS topic created by either
// the Amazon Redshift console, the Amazon SNS console, or the Amazon SNS API. To
// obtain an ARN with Amazon SNS, you must create a topic in Amazon SNS and
// subscribe to the topic. The ARN is displayed in the SNS console. You can specify
// the source type, and lists of Amazon Redshift source IDs, event categories, and
// event severities. Notifications will be sent for all events you want that match
// those criteria. For example, you can specify source type = cluster, source ID =
// my-cluster-1 and mycluster2, event categories = Availability, Backup, and
// severity = ERROR. The subscription will only send notifications for those ERROR
// events in the Availability and Backup categories for the specified clusters. If
// you specify both the source type and source IDs, such as source type = cluster
// and source identifier = my-cluster-1, notifications will be sent for all the
// cluster events for my-cluster-1. If you specify a source type but do not specify
// a source identifier, you will receive notice of the events for the objects of
// that type in your Amazon Web Services account. If you do not specify either the
// SourceType nor the SourceIdentifier, you will be notified of events generated
// from all Amazon Redshift sources belonging to your Amazon Web Services account.
// You must specify a source type if you specify a source ID.
func (c *Client) CreateEventSubscription(ctx context.Context, params *CreateEventSubscriptionInput, optFns ...func(*Options)) (*CreateEventSubscriptionOutput, error) {
	if params == nil {
		params = &CreateEventSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventSubscription", params, optFns, c.addOperationCreateEventSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventSubscriptionInput struct {

	// The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the
	// event notifications. The ARN is created by Amazon SNS when you create a topic
	// and subscribe to it.
	//
	// This member is required.
	SnsTopicArn *string

	// The name of the event subscription to be created. Constraints:
	//   - Cannot be null, empty, or blank.
	//   - Must contain from 1 to 255 alphanumeric characters or hyphens.
	//   - First character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	SubscriptionName *string

	// A boolean value; set to true to activate the subscription, and set to false to
	// create the subscription but not activate it.
	Enabled *bool

	// Specifies the Amazon Redshift event categories to be published by the event
	// notification subscription. Values: configuration, management, monitoring,
	// security, pending
	EventCategories []string

	// Specifies the Amazon Redshift event severity to be published by the event
	// notification subscription. Values: ERROR, INFO
	Severity *string

	// A list of one or more identifiers of Amazon Redshift source objects. All of the
	// objects must be of the same type as was specified in the source type parameter.
	// The event subscription will return only events generated by the specified
	// objects. If not specified, then events are returned for all objects within the
	// source type specified. Example: my-cluster-1, my-cluster-2 Example:
	// my-snapshot-20131010
	SourceIds []string

	// The type of source that will be generating the events. For example, if you want
	// to be notified of events generated by a cluster, you would set this parameter to
	// cluster. If this value is not specified, events are returned for all Amazon
	// Redshift objects in your Amazon Web Services account. You must specify a source
	// type in order to specify source IDs. Valid values: cluster,
	// cluster-parameter-group, cluster-security-group, cluster-snapshot, and
	// scheduled-action.
	SourceType *string

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEventSubscriptionOutput struct {

	// Describes event subscriptions.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateEventSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateEventSubscription{}, middleware.After)
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
	if err = addOpCreateEventSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventSubscription(options.Region), middleware.Before); err != nil {
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
		operation: "CreateEventSubscription",
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

func newServiceMetadataMiddleware_opCreateEventSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateEventSubscription",
	}
}
