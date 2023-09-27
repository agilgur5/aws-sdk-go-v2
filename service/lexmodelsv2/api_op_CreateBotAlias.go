// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an alias for the specified version of a bot. Use an alias to enable you
// to change the version of a bot without updating applications that use the bot.
// For example, you can create an alias called "PROD" that your applications use to
// call the Amazon Lex bot.
func (c *Client) CreateBotAlias(ctx context.Context, params *CreateBotAliasInput, optFns ...func(*Options)) (*CreateBotAliasOutput, error) {
	if params == nil {
		params = &CreateBotAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBotAlias", params, optFns, c.addOperationCreateBotAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBotAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBotAliasInput struct {

	// The alias to create. The name must be unique for the bot.
	//
	// This member is required.
	BotAliasName *string

	// The unique identifier of the bot that the alias applies to.
	//
	// This member is required.
	BotId *string

	// Maps configuration information to a specific locale. You can use this parameter
	// to specify a specific Lambda function to run different functions in different
	// locales.
	BotAliasLocaleSettings map[string]types.BotAliasLocaleSettings

	// The version of the bot that this alias points to. You can use the UpdateBotAlias (https://docs.aws.amazon.com/lexv2/latest/APIReference/API_UpdateBotAlias.html)
	// operation to change the bot version associated with the alias.
	BotVersion *string

	// Specifies whether Amazon Lex logs text and audio for a conversation with the
	// bot. When you enable conversation logs, text logs store text input, transcripts
	// of audio input, and associated metadata in Amazon CloudWatch Logs. Audio logs
	// store audio input in Amazon S3.
	ConversationLogSettings *types.ConversationLogSettings

	// A description of the alias. Use this description to help identify the alias.
	Description *string

	// Determines whether Amazon Lex will use Amazon Comprehend to detect the
	// sentiment of user utterances.
	SentimentAnalysisSettings *types.SentimentAnalysisSettings

	// A list of tags to add to the bot alias. You can only add tags when you create
	// an alias, you can't use the UpdateBotAlias operation to update the tags on a
	// bot alias. To update tags, use the TagResource operation.
	Tags map[string]string

	noSmithyDocumentSerde
}

func (*CreateBotAliasInput) operationName() string {
	return "CreateBotAlias"
}

type CreateBotAliasOutput struct {

	// The unique identifier of the bot alias.
	BotAliasId *string

	// Configuration information for a specific locale.
	BotAliasLocaleSettings map[string]types.BotAliasLocaleSettings

	// The name specified for the bot alias.
	BotAliasName *string

	// The current status of the alias. The alias is first put into the Creating
	// state. When the alias is ready to be used, it is put into the Available state.
	// You can use the DescribeBotAlias operation to get the current state of an alias.
	BotAliasStatus types.BotAliasStatus

	// The unique identifier of the bot that this alias applies to.
	BotId *string

	// The version of the bot associated with this alias.
	BotVersion *string

	// The conversation log settings specified for the alias.
	ConversationLogSettings *types.ConversationLogSettings

	// A Unix timestamp indicating the date and time that the bot alias was created.
	CreationDateTime *time.Time

	// The description specified for the bot alias.
	Description *string

	// Determines whether Amazon Lex will use Amazon Comprehend to detect the
	// sentiment of user utterances.
	SentimentAnalysisSettings *types.SentimentAnalysisSettings

	// A list of tags associated with the bot alias.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBotAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBotAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBotAlias{}, middleware.After)
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
	if err = addCreateBotAliasResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBotAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBotAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBotAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateBotAlias",
	}
}

type opCreateBotAliasResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateBotAliasResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateBotAliasResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "lex"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "lex"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("lex")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCreateBotAliasResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateBotAliasResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
