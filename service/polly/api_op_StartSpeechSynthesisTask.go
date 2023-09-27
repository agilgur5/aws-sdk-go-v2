// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows the creation of an asynchronous synthesis task, by starting a new
// SpeechSynthesisTask . This operation requires all the standard information
// needed for speech synthesis, plus the name of an Amazon S3 bucket for the
// service to store the output of the synthesis task and two optional parameters (
// OutputS3KeyPrefix and SnsTopicArn ). Once the synthesis task is created, this
// operation will return a SpeechSynthesisTask object, which will include an
// identifier of this task as well as the current status. The SpeechSynthesisTask
// object is available for 72 hours after starting the asynchronous synthesis task.
func (c *Client) StartSpeechSynthesisTask(ctx context.Context, params *StartSpeechSynthesisTaskInput, optFns ...func(*Options)) (*StartSpeechSynthesisTaskOutput, error) {
	if params == nil {
		params = &StartSpeechSynthesisTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSpeechSynthesisTask", params, optFns, c.addOperationStartSpeechSynthesisTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSpeechSynthesisTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSpeechSynthesisTaskInput struct {

	// The format in which the returned output will be encoded. For audio stream, this
	// will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json.
	//
	// This member is required.
	OutputFormat types.OutputFormat

	// Amazon S3 bucket name to which the output file will be saved.
	//
	// This member is required.
	OutputS3BucketName *string

	// The input text to synthesize. If you specify ssml as the TextType, follow the
	// SSML format for the input text.
	//
	// This member is required.
	Text *string

	// Voice ID to use for the synthesis.
	//
	// This member is required.
	VoiceId types.VoiceId

	// Specifies the engine ( standard or neural ) for Amazon Polly to use when
	// processing input text for speech synthesis. Using a voice that is not supported
	// for the engine selected will result in an error.
	Engine types.Engine

	// Optional language code for the Speech Synthesis request. This is only necessary
	// if using a bilingual voice, such as Aditi, which can be used for either Indian
	// English (en-IN) or Hindi (hi-IN). If a bilingual voice is used and no language
	// code is specified, Amazon Polly uses the default language of the bilingual
	// voice. The default language for any voice is the one returned by the
	// DescribeVoices (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html)
	// operation for the LanguageCode parameter. For example, if no language code is
	// specified, Aditi will use Indian English rather than Hindi.
	LanguageCode types.LanguageCode

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon is
	// the same as the language of the voice.
	LexiconNames []string

	// The Amazon S3 key prefix for the output speech file.
	OutputS3KeyPrefix *string

	// The audio frequency specified in Hz. The valid values for mp3 and ogg_vorbis
	// are "8000", "16000", "22050", and "24000". The default value for standard voices
	// is "22050". The default value for neural voices is "24000". Valid values for pcm
	// are "8000" and "16000" The default value is "16000".
	SampleRate *string

	// ARN for the SNS topic optionally used for providing status notification for a
	// speech synthesis task.
	SnsTopicArn *string

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []types.SpeechMarkType

	// Specifies whether the input text is plain text or SSML. The default value is
	// plain text.
	TextType types.TextType

	noSmithyDocumentSerde
}

func (*StartSpeechSynthesisTaskInput) operationName() string {
	return "StartSpeechSynthesisTask"
}

type StartSpeechSynthesisTaskOutput struct {

	// SynthesisTask object that provides information and attributes about a newly
	// submitted speech synthesis task.
	SynthesisTask *types.SynthesisTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSpeechSynthesisTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSpeechSynthesisTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSpeechSynthesisTask{}, middleware.After)
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
	if err = addStartSpeechSynthesisTaskResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSpeechSynthesisTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSpeechSynthesisTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartSpeechSynthesisTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "StartSpeechSynthesisTask",
	}
}

type opStartSpeechSynthesisTaskResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartSpeechSynthesisTaskResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartSpeechSynthesisTaskResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "polly"
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
				signingName = "polly"
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
				v4aScheme.SigningName = aws.String("polly")
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

func addStartSpeechSynthesisTaskResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartSpeechSynthesisTaskResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
