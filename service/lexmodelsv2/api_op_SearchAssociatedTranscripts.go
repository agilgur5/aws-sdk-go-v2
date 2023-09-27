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
)

// Search for associated transcripts that meet the specified criteria.
func (c *Client) SearchAssociatedTranscripts(ctx context.Context, params *SearchAssociatedTranscriptsInput, optFns ...func(*Options)) (*SearchAssociatedTranscriptsOutput, error) {
	if params == nil {
		params = &SearchAssociatedTranscriptsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchAssociatedTranscripts", params, optFns, c.addOperationSearchAssociatedTranscriptsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchAssociatedTranscriptsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchAssociatedTranscriptsInput struct {

	// The unique identifier of the bot associated with the transcripts that you are
	// searching.
	//
	// This member is required.
	BotId *string

	// The unique identifier of the bot recommendation associated with the transcripts
	// to search.
	//
	// This member is required.
	BotRecommendationId *string

	// The version of the bot containing the transcripts that you are searching.
	//
	// This member is required.
	BotVersion *string

	// A list of filter objects.
	//
	// This member is required.
	Filters []types.AssociatedTranscriptFilter

	// The identifier of the language and locale of the transcripts to search. The
	// string must match one of the supported locales. For more information, see
	// Supported languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	//
	// This member is required.
	LocaleId *string

	// The maximum number of bot recommendations to return in each page of results. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	MaxResults *int32

	// If the response from the SearchAssociatedTranscriptsRequest operation contains
	// more results than specified in the maxResults parameter, an index is returned in
	// the response. Use that index in the nextIndex parameter to return the next page
	// of results.
	NextIndex *int32

	// How SearchResults are ordered. Valid values are Ascending or Descending. The
	// default is Descending.
	SearchOrder types.SearchOrder

	noSmithyDocumentSerde
}

func (*SearchAssociatedTranscriptsInput) operationName() string {
	return "SearchAssociatedTranscripts"
}

type SearchAssociatedTranscriptsOutput struct {

	// The object that contains the associated transcript that meet the criteria you
	// specified.
	AssociatedTranscripts []types.AssociatedTranscript

	// The unique identifier of the bot associated with the transcripts that you are
	// searching.
	BotId *string

	// The unique identifier of the bot recommendation associated with the transcripts
	// to search.
	BotRecommendationId *string

	// The version of the bot containing the transcripts that you are searching.
	BotVersion *string

	// The identifier of the language and locale of the transcripts to search. The
	// string must match one of the supported locales. For more information, see
	// Supported languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	LocaleId *string

	// A index that indicates whether there are more results to return in a response
	// to the SearchAssociatedTranscripts operation. If the nextIndex field is present,
	// you send the contents as the nextIndex parameter of a
	// SearchAssociatedTranscriptsRequest operation to get the next page of results.
	NextIndex *int32

	// The total number of transcripts returned by the search.
	TotalResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchAssociatedTranscriptsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchAssociatedTranscripts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchAssociatedTranscripts{}, middleware.After)
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
	if err = addSearchAssociatedTranscriptsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSearchAssociatedTranscriptsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchAssociatedTranscripts(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchAssociatedTranscripts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "SearchAssociatedTranscripts",
	}
}

type opSearchAssociatedTranscriptsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSearchAssociatedTranscriptsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSearchAssociatedTranscriptsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addSearchAssociatedTranscriptsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSearchAssociatedTranscriptsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
