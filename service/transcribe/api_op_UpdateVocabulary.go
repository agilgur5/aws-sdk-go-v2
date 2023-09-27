// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an existing custom vocabulary with new values. This operation
// overwrites all existing information with your new values; you cannot append new
// terms onto an existing custom vocabulary.
func (c *Client) UpdateVocabulary(ctx context.Context, params *UpdateVocabularyInput, optFns ...func(*Options)) (*UpdateVocabularyOutput, error) {
	if params == nil {
		params = &UpdateVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVocabulary", params, optFns, c.addOperationUpdateVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVocabularyInput struct {

	// The language code that represents the language of the entries in the custom
	// vocabulary you want to update. Each custom vocabulary must contain terms in only
	// one language. A custom vocabulary can only be used to transcribe files in the
	// same language as the custom vocabulary. For example, if you create a custom
	// vocabulary using US English ( en-US ), you can only apply this custom vocabulary
	// to files that contain English audio. For a list of supported languages and their
	// associated language codes, refer to the Supported languages (https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html)
	// table.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The name of the custom vocabulary you want to update. Custom vocabulary names
	// are case sensitive.
	//
	// This member is required.
	VocabularyName *string

	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access
	// the Amazon S3 bucket that contains your input files (in this case, your custom
	// vocabulary). If the role that you specify doesn’t have the appropriate
	// permissions to access the specified Amazon S3 location, your request fails. IAM
	// role ARNs have the format arn:partition:iam::account:role/role-name-with-path .
	// For example: arn:aws:iam::111122223333:role/Admin . For more information, see
	// IAM ARNs (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns)
	// .
	DataAccessRoleArn *string

	// Use this parameter if you want to update your custom vocabulary by including
	// all desired terms, as comma-separated values, within your request. The other
	// option for updating your custom vocabulary is to save your entries in a text
	// file and upload them to an Amazon S3 bucket, then specify the location of your
	// file using the VocabularyFileUri parameter. Note that if you include Phrases in
	// your request, you cannot use VocabularyFileUri ; you must choose one or the
	// other. Each language has a character set that contains all allowed characters
	// for that specific language. If you use unsupported characters, your custom
	// vocabulary filter request fails. Refer to Character Sets for Custom Vocabularies (https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html)
	// to get the character set for your language.
	Phrases []string

	// The Amazon S3 location of the text file that contains your custom vocabulary.
	// The URI must be located in the same Amazon Web Services Region as the resource
	// you're calling. Here's an example URI path:
	// s3://DOC-EXAMPLE-BUCKET/my-vocab-file.txt Note that if you include
	// VocabularyFileUri in your request, you cannot use the Phrases flag; you must
	// choose one or the other.
	VocabularyFileUri *string

	noSmithyDocumentSerde
}

func (*UpdateVocabularyInput) operationName() string {
	return "UpdateVocabulary"
}

type UpdateVocabularyOutput struct {

	// The language code you selected for your custom vocabulary.
	LanguageCode types.LanguageCode

	// The date and time the specified custom vocabulary was last updated. Timestamps
	// are in the format YYYY-MM-DD'T'HH:MM:SS.SSSSSS-UTC . For example,
	// 2022-05-04T12:32:58.761000-07:00 represents 12:32 PM UTC-7 on May 4, 2022.
	LastModifiedTime *time.Time

	// The name of the updated custom vocabulary.
	VocabularyName *string

	// The processing state of your custom vocabulary. If the state is READY , you can
	// use the custom vocabulary in a StartTranscriptionJob request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateVocabulary{}, middleware.After)
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
	if err = addUpdateVocabularyResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVocabulary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "UpdateVocabulary",
	}
}

type opUpdateVocabularyResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateVocabularyResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateVocabularyResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "transcribe"
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
				signingName = "transcribe"
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
				v4aScheme.SigningName = aws.String("transcribe")
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

func addUpdateVocabularyResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateVocabularyResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
