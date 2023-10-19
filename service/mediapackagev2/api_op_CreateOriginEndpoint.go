// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The endpoint is attached to a channel, and represents the output of the live
// content. You can associate multiple endpoints to a single channel. Each endpoint
// gives players and downstream CDNs (such as Amazon CloudFront) access to the
// content for playback. Content can't be served from a channel until it has an
// endpoint. You can create only one endpoint with each request.
func (c *Client) CreateOriginEndpoint(ctx context.Context, params *CreateOriginEndpointInput, optFns ...func(*Options)) (*CreateOriginEndpointOutput, error) {
	if params == nil {
		params = &CreateOriginEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOriginEndpoint", params, optFns, c.addOperationCreateOriginEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOriginEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOriginEndpointInput struct {

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// The name that describes the channel. The name is the primary identifier for the
	// channel, and must be unique for your account in the AWS Region and channel
	// group.
	//
	// This member is required.
	ChannelName *string

	// The type of container to attach to this origin endpoint. A container type is a
	// file format that encapsulates one or more media streams, such as audio and
	// video, into a single file. You can't change the container type after you create
	// the endpoint.
	//
	// This member is required.
	ContainerType types.ContainerType

	// The name that describes the origin endpoint. The name is the primary identifier
	// for the origin endpoint, and must be unique for your account in the AWS Region
	// and channel. You can't use spaces in the name. You can't change the name after
	// you create the endpoint.
	//
	// This member is required.
	OriginEndpointName *string

	// A unique, case-sensitive token that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// Enter any descriptive text that helps you to identify the origin endpoint.
	Description *string

	// An HTTP live streaming (HLS) manifest configuration.
	HlsManifests []types.CreateHlsManifestConfiguration

	// A low-latency HLS manifest configuration.
	LowLatencyHlsManifests []types.CreateLowLatencyHlsManifestConfiguration

	// The segment configuration, including the segment name, duration, and other
	// configuration values.
	Segment *types.Segment

	// The size of the window (in seconds) to create a window of the live stream
	// that's available for on-demand viewing. Viewers can start-over or catch-up on
	// content that falls within the window. The maximum startover window is 1,209,600
	// seconds (14 days).
	StartoverWindowSeconds *int32

	// A comma-separated list of tag key:value pairs that you define. For example:
	// "Key1": "Value1",
	//     "Key2": "Value2"
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateOriginEndpointOutput struct {

	// The Amazon Resource Name (ARN) associated with the resource.
	//
	// This member is required.
	Arn *string

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// The name that describes the channel. The name is the primary identifier for the
	// channel, and must be unique for your account in the AWS Region and channel
	// group.
	//
	// This member is required.
	ChannelName *string

	// The type of container attached to this origin endpoint.
	//
	// This member is required.
	ContainerType types.ContainerType

	// The date and time the origin endpoint was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date and time the origin endpoint was modified.
	//
	// This member is required.
	ModifiedAt *time.Time

	// The name that describes the origin endpoint. The name is the primary identifier
	// for the origin endpoint, and and must be unique for your account in the AWS
	// Region and channel.
	//
	// This member is required.
	OriginEndpointName *string

	// The segment configuration, including the segment name, duration, and other
	// configuration values.
	//
	// This member is required.
	Segment *types.Segment

	// The description for your origin endpoint.
	Description *string

	// An HTTP live streaming (HLS) manifest configuration.
	HlsManifests []types.GetHlsManifestConfiguration

	// A low-latency HLS manifest configuration.
	LowLatencyHlsManifests []types.GetLowLatencyHlsManifestConfiguration

	// The size of the window (in seconds) to create a window of the live stream
	// that's available for on-demand viewing. Viewers can start-over or catch-up on
	// content that falls within the window.
	StartoverWindowSeconds *int32

	// The comma-separated list of tag key:value pairs assigned to the origin endpoint.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOriginEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateOriginEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateOriginEndpoint{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateOriginEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateOriginEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOriginEndpoint(options.Region), middleware.Before); err != nil {
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
		operation: "CreateOriginEndpoint",
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

type idempotencyToken_initializeOpCreateOriginEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateOriginEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateOriginEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateOriginEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateOriginEndpointInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateOriginEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateOriginEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateOriginEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackagev2",
		OperationName: "CreateOriginEndpoint",
	}
}
