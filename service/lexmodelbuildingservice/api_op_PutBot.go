// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an Amazon Lex conversational bot or replaces an existing bot. When you
// create or update a bot you are only required to specify a name, a locale, and
// whether the bot is directed toward children under age 13. You can use this to
// add intents later, or to remove intents from an existing bot. When you create a
// bot with the minimum information, the bot is created or updated but Amazon Lex
// returns the response FAILED . You can build the bot after you add one or more
// intents. For more information about Amazon Lex bots, see how-it-works . If you
// specify the name of an existing bot, the fields in the request replace the
// existing values in the $LATEST version of the bot. Amazon Lex removes any
// fields that you don't provide values for in the request, except for the
// idleTTLInSeconds and privacySettings fields, which are set to their default
// values. If you don't specify values for required fields, Amazon Lex throws an
// exception. This operation requires permissions for the lex:PutBot action. For
// more information, see security-iam .
func (c *Client) PutBot(ctx context.Context, params *PutBotInput, optFns ...func(*Options)) (*PutBotOutput, error) {
	if params == nil {
		params = &PutBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBot", params, optFns, c.addOperationPutBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBotInput struct {

	// For each Amazon Lex bot created with the Amazon Lex Model Building Service, you
	// must specify whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to the Children's Online Privacy Protection Act (COPPA)
	// by specifying true or false in the childDirected field. By specifying true in
	// the childDirected field, you confirm that your use of Amazon Lex is related to
	// a website, program, or other application that is directed or targeted, in whole
	// or in part, to children under age 13 and subject to COPPA. By specifying false
	// in the childDirected field, you confirm that your use of Amazon Lex is not
	// related to a website, program, or other application that is directed or
	// targeted, in whole or in part, to children under age 13 and subject to COPPA.
	// You may not specify a default value for the childDirected field that does not
	// accurately reflect whether your use of Amazon Lex is related to a website,
	// program, or other application that is directed or targeted, in whole or in part,
	// to children under age 13 and subject to COPPA. If your use of Amazon Lex relates
	// to a website, program, or other application that is directed in whole or in
	// part, to children under age 13, you must obtain any required verifiable parental
	// consent under COPPA. For information regarding the use of Amazon Lex in
	// connection with websites, programs, or other applications that are directed or
	// targeted, in whole or in part, to children under age 13, see the Amazon Lex FAQ. (https://aws.amazon.com/lex/faqs#data-security)
	//
	// This member is required.
	ChildDirected *bool

	// Specifies the target locale for the bot. Any intent used in the bot must be
	// compatible with the locale of the bot. The default is en-US .
	//
	// This member is required.
	Locale types.Locale

	// The name of the bot. The name is not case sensitive.
	//
	// This member is required.
	Name *string

	// When Amazon Lex can't understand the user's input in context, it tries to
	// elicit the information a few times. After that, Amazon Lex sends the message
	// defined in abortStatement to the user, and then cancels the conversation. To
	// set the number of retries, use the valueElicitationPrompt field for the slot
	// type. For example, in a pizza ordering bot, Amazon Lex might ask a user "What
	// type of crust would you like?" If the user's response is not one of the expected
	// responses (for example, "thin crust, "deep dish," etc.), Amazon Lex tries to
	// elicit a correct response a few more times. For example, in a pizza ordering
	// application, OrderPizza might be one of the intents. This intent might require
	// the CrustType slot. You specify the valueElicitationPrompt field when you
	// create the CrustType slot. If you have defined a fallback intent the cancel
	// statement will not be sent to the user, the fallback intent is used instead. For
	// more information, see AMAZON.FallbackIntent (https://docs.aws.amazon.com/lex/latest/dg/built-in-intent-fallback.html)
	// .
	AbortStatement *types.Statement

	// Identifies a specific revision of the $LATEST version. When you create a new
	// bot, leave the checksum field blank. If you specify a checksum you get a
	// BadRequestException exception. When you want to update a bot, set the checksum
	// field to the checksum of the most recent revision of the $LATEST version. If
	// you don't specify the checksum field, or if the checksum does not match the
	// $LATEST version, you get a PreconditionFailedException exception.
	Checksum *string

	// When Amazon Lex doesn't understand the user's intent, it uses this message to
	// get clarification. To specify how many times Amazon Lex should repeat the
	// clarification prompt, use the maxAttempts field. If Amazon Lex still doesn't
	// understand, it sends the message in the abortStatement field. When you create a
	// clarification prompt, make sure that it suggests the correct response from the
	// user. for example, for a bot that orders pizza and drinks, you might create this
	// clarification prompt: "What would you like to do? You can say 'Order a pizza' or
	// 'Order a drink.'" If you have defined a fallback intent, it will be invoked if
	// the clarification prompt is repeated the number of times defined in the
	// maxAttempts field. For more information, see  AMAZON.FallbackIntent (https://docs.aws.amazon.com/lex/latest/dg/built-in-intent-fallback.html)
	// . If you don't define a clarification prompt, at runtime Amazon Lex will return
	// a 400 Bad Request exception in three cases:
	//   - Follow-up prompt - When the user responds to a follow-up prompt but does
	//   not provide an intent. For example, in response to a follow-up prompt that says
	//   "Would you like anything else today?" the user says "Yes." Amazon Lex will
	//   return a 400 Bad Request exception because it does not have a clarification
	//   prompt to send to the user to get an intent.
	//   - Lambda function - When using a Lambda function, you return an ElicitIntent
	//   dialog type. Since Amazon Lex does not have a clarification prompt to get an
	//   intent from the user, it returns a 400 Bad Request exception.
	//   - PutSession operation - When using the PutSession operation, you send an
	//   ElicitIntent dialog type. Since Amazon Lex does not have a clarification
	//   prompt to get an intent from the user, it returns a 400 Bad Request exception.
	ClarificationPrompt *types.Prompt

	// When set to true a new numbered version of the bot is created. This is the same
	// as calling the CreateBotVersion operation. If you don't specify createVersion ,
	// the default is false .
	CreateVersion *bool

	// A description of the bot.
	Description *string

	// When set to true user utterances are sent to Amazon Comprehend for sentiment
	// analysis. If you don't specify detectSentiment , the default is false .
	DetectSentiment *bool

	// Set to true to enable access to natural language understanding improvements.
	// When you set the enableModelImprovements parameter to true you can use the
	// nluIntentConfidenceThreshold parameter to configure confidence scores. For more
	// information, see Confidence Scores (https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html)
	// . You can only set the enableModelImprovements parameter in certain Regions. If
	// you set the parameter to true , your bot has access to accuracy improvements.
	// The Regions where you can set the enableModelImprovements parameter to true
	// are:
	//   - US East (N. Virginia) (us-east-1)
	//   - US West (Oregon) (us-west-2)
	//   - Asia Pacific (Sydney) (ap-southeast-2)
	//   - EU (Ireland) (eu-west-1)
	// In other Regions, the enableModelImprovements parameter is set to true by
	// default. In these Regions setting the parameter to false throws a
	// ValidationException exception.
	EnableModelImprovements *bool

	// The maximum time in seconds that Amazon Lex retains the data gathered in a
	// conversation. A user interaction session remains active for the amount of time
	// specified. If no conversation occurs during this time, the session expires and
	// Amazon Lex deletes any data provided before the timeout. For example, suppose
	// that a user chooses the OrderPizza intent, but gets sidetracked halfway through
	// placing an order. If the user doesn't complete the order within the specified
	// time, Amazon Lex discards the slot information that it gathered, and the user
	// must start over. If you don't include the idleSessionTTLInSeconds element in a
	// PutBot operation request, Amazon Lex uses the default value. This is also true
	// if the request replaces an existing bot. The default is 300 seconds (5 minutes).
	IdleSessionTTLInSeconds *int32

	// An array of Intent objects. Each intent represents a command that a user can
	// express. For example, a pizza ordering bot might support an OrderPizza intent.
	// For more information, see how-it-works .
	Intents []types.Intent

	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent
	// , AMAZON.KendraSearchIntent , or both when returning alternative intents in a
	// PostContent (https://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostContent.html)
	// or PostText (https://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostText.html)
	// response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted
	// if they are configured for the bot. You must set the enableModelImprovements
	// parameter to true to use confidence scores in the following regions.
	//   - US East (N. Virginia) (us-east-1)
	//   - US West (Oregon) (us-west-2)
	//   - Asia Pacific (Sydney) (ap-southeast-2)
	//   - EU (Ireland) (eu-west-1)
	// In other Regions, the enableModelImprovements parameter is set to true by
	// default. For example, suppose a bot is configured with the confidence threshold
	// of 0.80 and the AMAZON.FallbackIntent . Amazon Lex returns three alternative
	// intents with the following confidence scores: IntentA (0.70), IntentB (0.60),
	// IntentC (0.50). The response from the PostText operation would be:
	//   - AMAZON.FallbackIntent
	//   - IntentA
	//   - IntentB
	//   - IntentC
	NluIntentConfidenceThreshold *float64

	// If you set the processBehavior element to BUILD , Amazon Lex builds the bot so
	// that it can be run. If you set the element to SAVE Amazon Lex saves the bot,
	// but doesn't build it. If you don't specify this value, the default value is
	// BUILD .
	ProcessBehavior types.ProcessBehavior

	// A list of tags to add to the bot. You can only add tags when you create a bot,
	// you can't use the PutBot operation to update the tags on a bot. To update tags,
	// use the TagResource operation.
	Tags []types.Tag

	// The Amazon Polly voice ID that you want Amazon Lex to use for voice
	// interactions with the user. The locale configured for the voice must match the
	// locale of the bot. For more information, see Voices in Amazon Polly (https://docs.aws.amazon.com/polly/latest/dg/voicelist.html)
	// in the Amazon Polly Developer Guide.
	VoiceId *string

	noSmithyDocumentSerde
}

type PutBotOutput struct {

	// The message that Amazon Lex uses to cancel a conversation. For more
	// information, see PutBot .
	AbortStatement *types.Statement

	// Checksum of the bot that you created.
	Checksum *string

	// For each Amazon Lex bot created with the Amazon Lex Model Building Service, you
	// must specify whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to the Children's Online Privacy Protection Act (COPPA)
	// by specifying true or false in the childDirected field. By specifying true in
	// the childDirected field, you confirm that your use of Amazon Lex is related to
	// a website, program, or other application that is directed or targeted, in whole
	// or in part, to children under age 13 and subject to COPPA. By specifying false
	// in the childDirected field, you confirm that your use of Amazon Lex is not
	// related to a website, program, or other application that is directed or
	// targeted, in whole or in part, to children under age 13 and subject to COPPA.
	// You may not specify a default value for the childDirected field that does not
	// accurately reflect whether your use of Amazon Lex is related to a website,
	// program, or other application that is directed or targeted, in whole or in part,
	// to children under age 13 and subject to COPPA. If your use of Amazon Lex relates
	// to a website, program, or other application that is directed in whole or in
	// part, to children under age 13, you must obtain any required verifiable parental
	// consent under COPPA. For information regarding the use of Amazon Lex in
	// connection with websites, programs, or other applications that are directed or
	// targeted, in whole or in part, to children under age 13, see the Amazon Lex FAQ. (https://aws.amazon.com/lex/faqs#data-security)
	ChildDirected *bool

	// The prompts that Amazon Lex uses when it doesn't understand the user's intent.
	// For more information, see PutBot .
	ClarificationPrompt *types.Prompt

	// True if a new version of the bot was created. If the createVersion field was
	// not specified in the request, the createVersion field is set to false in the
	// response.
	CreateVersion *bool

	// The date that the bot was created.
	CreatedDate *time.Time

	// A description of the bot.
	Description *string

	// true if the bot is configured to send user utterances to Amazon Comprehend for
	// sentiment analysis. If the detectSentiment field was not specified in the
	// request, the detectSentiment field is false in the response.
	DetectSentiment *bool

	// Indicates whether the bot uses accuracy improvements. true indicates that the
	// bot is using the improvements, otherwise, false .
	EnableModelImprovements *bool

	// If status is FAILED , Amazon Lex provides the reason that it failed to build the
	// bot.
	FailureReason *string

	// The maximum length of time that Amazon Lex retains the data gathered in a
	// conversation. For more information, see PutBot .
	IdleSessionTTLInSeconds *int32

	// An array of Intent objects. For more information, see PutBot .
	Intents []types.Intent

	// The date that the bot was updated. When you create a resource, the creation
	// date and last updated date are the same.
	LastUpdatedDate *time.Time

	// The target locale for the bot.
	Locale types.Locale

	// The name of the bot.
	Name *string

	// The score that determines where Amazon Lex inserts the AMAZON.FallbackIntent ,
	// AMAZON.KendraSearchIntent , or both when returning alternative intents in a
	// PostContent (https://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostContent.html)
	// or PostText (https://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostText.html)
	// response. AMAZON.FallbackIntent is inserted if the confidence score for all
	// intents is below this value. AMAZON.KendraSearchIntent is only inserted if it
	// is configured for the bot.
	NluIntentConfidenceThreshold *float64

	// When you send a request to create a bot with processBehavior set to BUILD ,
	// Amazon Lex sets the status response element to BUILDING . In the
	// READY_BASIC_TESTING state you can test the bot with user inputs that exactly
	// match the utterances configured for the bot's intents and values in the slot
	// types. If Amazon Lex can't build the bot, Amazon Lex sets status to FAILED .
	// Amazon Lex returns the reason for the failure in the failureReason response
	// element. When you set processBehavior to SAVE , Amazon Lex sets the status code
	// to NOT BUILT . When the bot is in the READY state you can test and publish the
	// bot.
	Status types.Status

	// A list of tags associated with the bot.
	Tags []types.Tag

	// The version of the bot. For a new bot, the version is always $LATEST .
	Version *string

	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the
	// user. For more information, see PutBot .
	VoiceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBot{}, middleware.After)
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
	if err = addOpPutBotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBot(options.Region), middleware.Before); err != nil {
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
		operation: "PutBot",
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

func newServiceMetadataMiddleware_opPutBot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "PutBot",
	}
}
