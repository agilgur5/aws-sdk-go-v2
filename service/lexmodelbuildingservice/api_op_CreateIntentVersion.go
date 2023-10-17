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

// Creates a new version of an intent based on the $LATEST version of the intent.
// If the $LATEST version of this intent hasn't changed since you last updated it,
// Amazon Lex doesn't create a new version. It returns the last version you
// created. You can update only the $LATEST version of the intent. You can't
// update the numbered versions that you create with the CreateIntentVersion
// operation. When you create a version of an intent, Amazon Lex sets the version
// to 1. Subsequent versions increment by 1. For more information, see
// versioning-intro . This operation requires permissions to perform the
// lex:CreateIntentVersion action.
func (c *Client) CreateIntentVersion(ctx context.Context, params *CreateIntentVersionInput, optFns ...func(*Options)) (*CreateIntentVersionOutput, error) {
	if params == nil {
		params = &CreateIntentVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIntentVersion", params, optFns, c.addOperationCreateIntentVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIntentVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIntentVersionInput struct {

	// The name of the intent that you want to create a new version of. The name is
	// case sensitive.
	//
	// This member is required.
	Name *string

	// Checksum of the $LATEST version of the intent that should be used to create the
	// new version. If you specify a checksum and the $LATEST version of the intent
	// has a different checksum, Amazon Lex returns a PreconditionFailedException
	// exception and doesn't publish a new version. If you don't specify a checksum,
	// Amazon Lex publishes the $LATEST version.
	Checksum *string

	noSmithyDocumentSerde
}

type CreateIntentVersionOutput struct {

	// Checksum of the intent version created.
	Checksum *string

	// After the Lambda function specified in the fulfillmentActivity field fulfills
	// the intent, Amazon Lex conveys this statement to the user.
	ConclusionStatement *types.Statement

	// If defined, the prompt that Amazon Lex uses to confirm the user's intent before
	// fulfilling it.
	ConfirmationPrompt *types.Prompt

	// The date that the intent was created.
	CreatedDate *time.Time

	// A description of the intent.
	Description *string

	// If defined, Amazon Lex invokes this Lambda function for each user input.
	DialogCodeHook *types.CodeHook

	// If defined, Amazon Lex uses this prompt to solicit additional user activity
	// after the intent is fulfilled.
	FollowUpPrompt *types.FollowUpPrompt

	// Describes how the intent is fulfilled.
	FulfillmentActivity *types.FulfillmentActivity

	// An array of InputContext objects that lists the contexts that must be active
	// for Amazon Lex to choose the intent in a conversation with the user.
	InputContexts []types.InputContext

	// Configuration information, if any, for connecting an Amazon Kendra index with
	// the AMAZON.KendraSearchIntent intent.
	KendraConfiguration *types.KendraConfiguration

	// The date that the intent was updated.
	LastUpdatedDate *time.Time

	// The name of the intent.
	Name *string

	// An array of OutputContext objects that lists the contexts that the intent
	// activates when the intent is fulfilled.
	OutputContexts []types.OutputContext

	// A unique identifier for a built-in intent.
	ParentIntentSignature *string

	// If the user answers "no" to the question defined in confirmationPrompt , Amazon
	// Lex responds with this statement to acknowledge that the intent was canceled.
	RejectionStatement *types.Statement

	// An array of sample utterances configured for the intent.
	SampleUtterances []string

	// An array of slot types that defines the information required to fulfill the
	// intent.
	Slots []types.Slot

	// The version number assigned to the new version of the intent.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIntentVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIntentVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIntentVersion{}, middleware.After)
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
	if err = addOpCreateIntentVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIntentVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIntentVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateIntentVersion",
	}
}
