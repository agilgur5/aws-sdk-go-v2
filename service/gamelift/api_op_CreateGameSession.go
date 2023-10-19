// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a multiplayer game session for players in a specific fleet location.
// This operation prompts an available server process to start a game session and
// retrieves connection information for the new game session. As an alternative,
// consider using the Amazon GameLift game session placement feature with
// StartGameSessionPlacement (https://docs.aws.amazon.com/gamelift/latest/apireference/API_StartGameSessionPlacement.html)
// , which uses the FleetIQ algorithm and queues to optimize the placement process.
// When creating a game session, you specify exactly where you want to place it and
// provide a set of game session configuration settings. The target fleet must be
// in ACTIVE status. You can use this operation in the following ways:
//   - To create a game session on an instance in a fleet's home Region, provide a
//     fleet or alias ID along with your game session configuration.
//   - To create a game session on an instance in a fleet's remote location,
//     provide a fleet or alias ID and a location name, along with your game session
//     configuration.
//   - To create a game session on an instance in an Anywhere fleet, specify the
//     fleet's custom location.
//
// If successful, Amazon GameLift initiates a workflow to start a new game session
// and returns a GameSession object containing the game session configuration and
// status. When the game session status is ACTIVE , it is updated with connection
// information and you can create player sessions for the game session. By default,
// newly created game sessions are open to new players. You can restrict new player
// access by using UpdateGameSession (https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateGameSession.html)
// to change the game session's player session creation policy. Amazon GameLift
// retains logs for active for 14 days. To access the logs, call
// GetGameSessionLogUrl (https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetGameSessionLogUrl.html)
// to download the log files. Available in Amazon GameLift Local. Learn more Start
// a game session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)
// All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreateGameSession(ctx context.Context, params *CreateGameSessionInput, optFns ...func(*Options)) (*CreateGameSessionOutput, error) {
	if params == nil {
		params = &CreateGameSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGameSession", params, optFns, c.addOperationCreateGameSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGameSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGameSessionInput struct {

	// The maximum number of players that can be connected simultaneously to the game
	// session.
	//
	// This member is required.
	MaximumPlayerSessionCount *int32

	// A unique identifier for the alias associated with the fleet to create a game
	// session in. You can use either the alias ID or ARN value. Each request must
	// reference either a fleet ID or alias ID, but not both.
	AliasId *string

	// A unique identifier for a player or entity creating the game session. If you
	// add a resource creation limit policy to a fleet, the CreateGameSession
	// operation requires a CreatorId . Amazon GameLift limits the number of game
	// session creation requests with the same CreatorId in a specified time period.
	// If you your fleet doesn't have a resource creation limit policy and you provide
	// a CreatorId in your CreateGameSession requests, Amazon GameLift limits requests
	// to one request per CreatorId per second. To not limit CreateGameSession
	// requests with the same CreatorId , don't provide a CreatorId in your
	// CreateGameSession request.
	CreatorId *string

	// A unique identifier for the fleet to create a game session in. You can use
	// either the fleet ID or ARN value. Each request must reference either a fleet ID
	// or alias ID, but not both.
	FleetId *string

	// A set of custom properties for a game session, formatted as key:value pairs.
	// These properties are passed to a game server process with a request to start a
	// new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)
	// ).
	GameProperties []types.GameProperty

	// A set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process with a request to start a new game
	// session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)
	// ).
	GameSessionData *string

	// This parameter is deprecated. Use IdempotencyToken instead. Custom string that
	// uniquely identifies a request for a new game session. Maximum token length is 48
	// characters. If provided, this string is included in the new game session's ID.
	GameSessionId *string

	// Custom string that uniquely identifies the new game session request. This is
	// useful for ensuring that game session requests with the same idempotency token
	// are processed only once. Subsequent requests with the same string return the
	// original GameSession object, with an updated status. Maximum token length is 48
	// characters. If provided, this string is included in the new game session's ID. A
	// game session ARN has the following format: arn:aws:gamelift:::gamesession// .
	// Idempotency tokens remain in use for 30 days after a game session has ended;
	// game session objects are retained for this time period and then deleted.
	IdempotencyToken *string

	// A fleet's remote location to place the new game session in. If this parameter
	// is not set, the new game session is placed in the fleet's home Region. Specify a
	// remote location with an Amazon Web Services Region code such as us-west-2 . When
	// using an Anywhere fleet, this parameter is required and must be set to the
	// Anywhere fleet's custom location.
	Location *string

	// A descriptive label that is associated with a game session. Session names do
	// not need to be unique.
	Name *string

	noSmithyDocumentSerde
}

type CreateGameSessionOutput struct {

	// Object that describes the newly created game session record.
	GameSession *types.GameSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGameSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGameSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGameSession{}, middleware.After)
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
	if err = addOpCreateGameSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGameSession(options.Region), middleware.Before); err != nil {
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
		operation: "CreateGameSession",
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

func newServiceMetadataMiddleware_opCreateGameSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateGameSession",
	}
}
