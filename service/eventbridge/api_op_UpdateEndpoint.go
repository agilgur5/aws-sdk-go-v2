// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an existing endpoint. For more information about global endpoints, see
// Making applications Regional-fault tolerant with global endpoints and event
// replication (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-global-endpoints.html)
// in the Amazon EventBridge User Guide..
func (c *Client) UpdateEndpoint(ctx context.Context, params *UpdateEndpointInput, optFns ...func(*Options)) (*UpdateEndpointOutput, error) {
	if params == nil {
		params = &UpdateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEndpoint", params, optFns, c.addOperationUpdateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointInput struct {

	// The name of the endpoint you want to update.
	//
	// This member is required.
	Name *string

	// A description for the endpoint.
	Description *string

	// Define event buses used for replication.
	EventBuses []types.EndpointEventBus

	// Whether event replication was enabled or disabled by this request.
	ReplicationConfig *types.ReplicationConfig

	// The ARN of the role used by event replication for this request.
	RoleArn *string

	// Configure the routing policy, including the health check and secondary Region.
	RoutingConfig *types.RoutingConfig

	noSmithyDocumentSerde
}

func (*UpdateEndpointInput) operationName() string {
	return "UpdateEndpoint"
}

type UpdateEndpointOutput struct {

	// The ARN of the endpoint you updated in this request.
	Arn *string

	// The ID of the endpoint you updated in this request.
	EndpointId *string

	// The URL of the endpoint you updated in this request.
	EndpointUrl *string

	// The event buses used for replication for the endpoint you updated in this
	// request.
	EventBuses []types.EndpointEventBus

	// The name of the endpoint you updated in this request.
	Name *string

	// Whether event replication was enabled or disabled for the endpoint you updated
	// in this request.
	ReplicationConfig *types.ReplicationConfig

	// The ARN of the role used by event replication for the endpoint you updated in
	// this request.
	RoleArn *string

	// The routing configuration you updated in this request.
	RoutingConfig *types.RoutingConfig

	// The state of the endpoint you updated in this request.
	State types.EndpointState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpoint{}, middleware.After)
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
	if err = addOpUpdateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "UpdateEndpoint",
	}
}
