// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new version of the bot based on the DRAFT version. If the DRAFT
// version of this resource hasn't changed since you created the last version,
// Amazon Lex doesn't create a new version, it returns the last created version.
// When you create the first version of a bot, Amazon Lex sets the version to 1.
// Subsequent versions increment by 1.
func (c *Client) CreateBotVersion(ctx context.Context, params *CreateBotVersionInput, optFns ...func(*Options)) (*CreateBotVersionOutput, error) {
	if params == nil {
		params = &CreateBotVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBotVersion", params, optFns, c.addOperationCreateBotVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBotVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBotVersionInput struct {

	// The identifier of the bot to create the version for.
	//
	// This member is required.
	BotId *string

	// Specifies the locales that Amazon Lex adds to this version. You can choose the
	// Draft version or any other previously published version for each locale. When
	// you specify a source version, the locale data is copied from the source version
	// to the new version.
	//
	// This member is required.
	BotVersionLocaleSpecification map[string]types.BotVersionLocaleDetails

	// A description of the version. Use the description to help identify the version
	// in lists.
	Description *string

	noSmithyDocumentSerde
}

type CreateBotVersionOutput struct {

	// The bot identifier specified in the request.
	BotId *string

	// When you send a request to create or update a bot, Amazon Lex sets the status
	// response element to Creating . After Amazon Lex builds the bot, it sets status
	// to Available . If Amazon Lex can't build the bot, it sets status to Failed .
	BotStatus types.BotStatus

	// The version number assigned to the version.
	BotVersion *string

	// The source versions used for each locale in the new version.
	BotVersionLocaleSpecification map[string]types.BotVersionLocaleDetails

	// A timestamp of the date and time that the version was created.
	CreationDateTime *time.Time

	// The description of the version specified in the request.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBotVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBotVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBotVersion{}, middleware.After)
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
	if err = addOpCreateBotVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBotVersion(options.Region), middleware.Before); err != nil {
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
		operation: "CreateBotVersion",
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

func newServiceMetadataMiddleware_opCreateBotVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateBotVersion",
	}
}
