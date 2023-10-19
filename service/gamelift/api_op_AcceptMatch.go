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

// Registers a player's acceptance or rejection of a proposed FlexMatch match. A
// matchmaking configuration may require player acceptance; if so, then matches
// built with that configuration cannot be completed unless all players accept the
// proposed match within a specified time limit. When FlexMatch builds a match, all
// the matchmaking tickets involved in the proposed match are placed into status
// REQUIRES_ACCEPTANCE . This is a trigger for your game to get acceptance from all
// players in each ticket. Calls to this action are only valid for tickets that are
// in this status; calls for tickets not in this status result in an error. To
// register acceptance, specify the ticket ID, one or more players, and an
// acceptance response. When all players have accepted, Amazon GameLift advances
// the matchmaking tickets to status PLACING , and attempts to create a new game
// session for the match. If any player rejects the match, or if acceptances are
// not received before a specified timeout, the proposed match is dropped. Each
// matchmaking ticket in the failed match is handled as follows:
//   - If the ticket has one or more players who rejected the match or failed to
//     respond, the ticket status is set CANCELLED and processing is terminated.
//   - If all players in the ticket accepted the match, the ticket status is
//     returned to SEARCHING to find a new match.
//
// Learn more  Add FlexMatch to a game client (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-client.html)
// FlexMatch events (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html)
// (reference)
func (c *Client) AcceptMatch(ctx context.Context, params *AcceptMatchInput, optFns ...func(*Options)) (*AcceptMatchOutput, error) {
	if params == nil {
		params = &AcceptMatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptMatch", params, optFns, c.addOperationAcceptMatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptMatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptMatchInput struct {

	// Player response to the proposed match.
	//
	// This member is required.
	AcceptanceType types.AcceptanceType

	// A unique identifier for a player delivering the response. This parameter can
	// include one or multiple player IDs.
	//
	// This member is required.
	PlayerIds []string

	// A unique identifier for a matchmaking ticket. The ticket must be in status
	// REQUIRES_ACCEPTANCE ; otherwise this request will fail.
	//
	// This member is required.
	TicketId *string

	noSmithyDocumentSerde
}

type AcceptMatchOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptMatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptMatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptMatch{}, middleware.After)
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
	if err = addOpAcceptMatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptMatch(options.Region), middleware.Before); err != nil {
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
		operation: "AcceptMatch",
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

func newServiceMetadataMiddleware_opAcceptMatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "AcceptMatch",
	}
}
