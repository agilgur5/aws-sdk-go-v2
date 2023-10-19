// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new routing profile.
func (c *Client) CreateRoutingProfile(ctx context.Context, params *CreateRoutingProfileInput, optFns ...func(*Options)) (*CreateRoutingProfileOutput, error) {
	if params == nil {
		params = &CreateRoutingProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRoutingProfile", params, optFns, c.addOperationCreateRoutingProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRoutingProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRoutingProfileInput struct {

	// The default outbound queue for the routing profile.
	//
	// This member is required.
	DefaultOutboundQueueId *string

	// Description of the routing profile. Must not be more than 250 characters.
	//
	// This member is required.
	Description *string

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The channels that agents can handle in the Contact Control Panel (CCP) for this
	// routing profile.
	//
	// This member is required.
	MediaConcurrencies []types.MediaConcurrency

	// The name of the routing profile. Must not be more than 127 characters.
	//
	// This member is required.
	Name *string

	// Whether agents with this routing profile will have their routing order
	// calculated based on longest idle time or time since their last inbound contact.
	AgentAvailabilityTimer types.AgentAvailabilityTimer

	// The inbound queues associated with the routing profile. If no queue is added,
	// the agent can make only outbound calls. The limit of 10 array members applies to
	// the maximum number of RoutingProfileQueueConfig objects that can be passed
	// during a CreateRoutingProfile API request. It is different from the quota of 50
	// queues per routing profile per instance that is listed in Amazon Connect
	// service quotas (https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html)
	// .
	QueueConfigs []types.RoutingProfileQueueConfig

	// The tags used to organize, track, or control access for this resource. For
	// example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateRoutingProfileOutput struct {

	// The Amazon Resource Name (ARN) of the routing profile.
	RoutingProfileArn *string

	// The identifier of the routing profile.
	RoutingProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRoutingProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoutingProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoutingProfile{}, middleware.After)
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
	if err = addOpCreateRoutingProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRoutingProfile(options.Region), middleware.Before); err != nil {
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
		operation: "CreateRoutingProfile",
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

func newServiceMetadataMiddleware_opCreateRoutingProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "CreateRoutingProfile",
	}
}
