// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies and starts a remote access session.
func (c *Client) CreateRemoteAccessSession(ctx context.Context, params *CreateRemoteAccessSessionInput, optFns ...func(*Options)) (*CreateRemoteAccessSessionOutput, error) {
	if params == nil {
		params = &CreateRemoteAccessSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRemoteAccessSession", params, optFns, c.addOperationCreateRemoteAccessSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRemoteAccessSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates and submits a request to start a remote access session.
type CreateRemoteAccessSessionInput struct {

	// The ARN of the device for which you want to create a remote access session.
	//
	// This member is required.
	DeviceArn *string

	// The Amazon Resource Name (ARN) of the project for which you want to create a
	// remote access session.
	//
	// This member is required.
	ProjectArn *string

	// Unique identifier for the client. If you want access to multiple devices on the
	// same client, you should pass the same clientId value in each call to
	// CreateRemoteAccessSession . This identifier is required only if
	// remoteDebugEnabled is set to true . Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html)
	// .
	ClientId *string

	// The configuration information for the remote access session request.
	Configuration *types.CreateRemoteAccessSessionConfiguration

	// The Amazon Resource Name (ARN) of the device instance for which you want to
	// create a remote access session.
	InstanceArn *string

	// The interaction mode of the remote access session. Valid values are:
	//   - INTERACTIVE: You can interact with the iOS device by viewing, touching, and
	//   rotating the screen. You cannot run XCUITest framework-based tests in this mode.
	//
	//   - NO_VIDEO: You are connected to the device, but cannot interact with it or
	//   view the screen. This mode has the fastest test execution speed. You can run
	//   XCUITest framework-based tests in this mode.
	//   - VIDEO_ONLY: You can view the screen, but cannot touch or rotate it. You can
	//   run XCUITest framework-based tests and watch the screen in this mode.
	InteractionMode types.InteractionMode

	// The name of the remote access session to create.
	Name *string

	// Set to true if you want to access devices remotely for debugging in your remote
	// access session. Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html)
	// .
	RemoteDebugEnabled *bool

	// The Amazon Resource Name (ARN) for the app to be recorded in the remote access
	// session.
	RemoteRecordAppArn *string

	// Set to true to enable remote recording for the remote access session.
	RemoteRecordEnabled *bool

	// When set to true , for private devices, Device Farm does not sign your app
	// again. For public devices, Device Farm always signs your apps again. For more
	// information on how Device Farm modifies your uploads during tests, see Do you
	// modify my app? (http://aws.amazon.com/device-farm/faqs/)
	SkipAppResign *bool

	// Ignored. The public key of the ssh key pair you want to use for connecting to
	// remote devices in your remote debugging session. This key is required only if
	// remoteDebugEnabled is set to true . Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html)
	// .
	SshPublicKey *string

	noSmithyDocumentSerde
}

// Represents the server response from a request to create a remote access session.
type CreateRemoteAccessSessionOutput struct {

	// A container that describes the remote access session when the request to create
	// a remote access session is sent.
	RemoteAccessSession *types.RemoteAccessSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRemoteAccessSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRemoteAccessSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRemoteAccessSession{}, middleware.After)
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
	if err = addOpCreateRemoteAccessSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRemoteAccessSession(options.Region), middleware.Before); err != nil {
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
		operation: "CreateRemoteAccessSession",
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

func newServiceMetadataMiddleware_opCreateRemoteAccessSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateRemoteAccessSession",
	}
}
