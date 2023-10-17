// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an encrypted token that is used by a chat participant to establish an
// individual WebSocket chat connection to a room. When the token is used to
// connect to chat, the connection is valid for the session duration specified in
// the request. The token becomes invalid at the token-expiration timestamp
// included in the response. Use the capabilities field to permit an end user to
// send messages or moderate a room. The attributes field securely attaches
// structured data to the chat session; the data is included within each message
// sent by the end user and received by other participants in the room. Common use
// cases for attributes include passing end-user profile data like an icon, display
// name, colors, badges, and other display features. Encryption keys are owned by
// Amazon IVS Chat and never used directly by your application.
func (c *Client) CreateChatToken(ctx context.Context, params *CreateChatTokenInput, optFns ...func(*Options)) (*CreateChatTokenOutput, error) {
	if params == nil {
		params = &CreateChatTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChatToken", params, optFns, c.addOperationCreateChatTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChatTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChatTokenInput struct {

	// Identifier of the room that the client is trying to access. Currently this must
	// be an ARN.
	//
	// This member is required.
	RoomIdentifier *string

	// Application-provided ID that uniquely identifies the user associated with this
	// token. This can be any UTF-8 encoded text.
	//
	// This member is required.
	UserId *string

	// Application-provided attributes to encode into the token and attach to a chat
	// session. Map keys and values can contain UTF-8 encoded text. The maximum length
	// of this field is 1 KB total.
	Attributes map[string]string

	// Set of capabilities that the user is allowed to perform in the room. Default:
	// None (the capability to view messages is implicitly included in all requests).
	Capabilities []types.ChatTokenCapability

	// Session duration (in minutes), after which the session expires. Default: 60 (1
	// hour).
	SessionDurationInMinutes int32

	noSmithyDocumentSerde
}

type CreateChatTokenOutput struct {

	// Time after which an end user's session is no longer valid. This is an ISO 8601
	// timestamp; note that this is returned as a string.
	SessionExpirationTime *time.Time

	// The issued client token, encrypted.
	Token *string

	// Time after which the token is no longer valid and cannot be used to connect to
	// a room. This is an ISO 8601 timestamp; note that this is returned as a string.
	TokenExpirationTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChatTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChatToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChatToken{}, middleware.After)
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
	if err = addOpCreateChatTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChatToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChatToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivschat",
		OperationName: "CreateChatToken",
	}
}
