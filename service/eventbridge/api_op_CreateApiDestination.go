// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an API destination, which is an HTTP invocation endpoint configured as
// a target for events.
func (c *Client) CreateApiDestination(ctx context.Context, params *CreateApiDestinationInput, optFns ...func(*Options)) (*CreateApiDestinationOutput, error) {
	if params == nil {
		params = &CreateApiDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApiDestination", params, optFns, c.addOperationCreateApiDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApiDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApiDestinationInput struct {

	// The ARN of the connection to use for the API destination. The destination
	// endpoint must support the authorization type specified for the connection.
	//
	// This member is required.
	ConnectionArn *string

	// The method to use for the request to the HTTP invocation endpoint.
	//
	// This member is required.
	HttpMethod types.ApiDestinationHttpMethod

	// The URL to the HTTP invocation endpoint for the API destination.
	//
	// This member is required.
	InvocationEndpoint *string

	// The name for the API destination to create.
	//
	// This member is required.
	Name *string

	// A description for the API destination to create.
	Description *string

	// The maximum number of requests per second to send to the HTTP invocation
	// endpoint.
	InvocationRateLimitPerSecond *int32

	noSmithyDocumentSerde
}

func (*CreateApiDestinationInput) operationName() string {
	return "CreateApiDestination"
}

type CreateApiDestinationOutput struct {

	// The ARN of the API destination that was created by the request.
	ApiDestinationArn *string

	// The state of the API destination that was created by the request.
	ApiDestinationState types.ApiDestinationState

	// A time stamp indicating the time that the API destination was created.
	CreationTime *time.Time

	// A time stamp indicating the time that the API destination was last modified.
	LastModifiedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApiDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateApiDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateApiDestination{}, middleware.After)
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
	if err = addOpCreateApiDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApiDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApiDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "CreateApiDestination",
	}
}
