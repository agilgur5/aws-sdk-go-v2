// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a temporary session URL and authorization code(bearer token) that you
// can use to embed an Amazon QuickSight read-only dashboard in your website or
// application. Before you use this command, make sure that you have configured the
// dashboards and permissions. Currently, you can use GetDashboardEmbedURL only
// from the server, not from the user's browser. The following rules apply to the
// generated URL:
//   - They must be used together.
//   - They can be used one time only.
//   - They are valid for 5 minutes after you run this command.
//   - You are charged only when the URL is used or there is interaction with
//     Amazon QuickSight.
//   - The resulting user session is valid for 15 minutes (default) up to 10 hours
//     (maximum). You can use the optional SessionLifetimeInMinutes parameter to
//     customize session duration.
//
// For more information, see Embedding Analytics Using GetDashboardEmbedUrl (https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics-deprecated.html)
// in the Amazon QuickSight User Guide. For more information about the high-level
// steps for embedding and for an interactive demo of the ways you can customize
// embedding, visit the Amazon QuickSight Developer Portal (https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-portal.html)
// .
func (c *Client) GetDashboardEmbedUrl(ctx context.Context, params *GetDashboardEmbedUrlInput, optFns ...func(*Options)) (*GetDashboardEmbedUrlOutput, error) {
	if params == nil {
		params = &GetDashboardEmbedUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDashboardEmbedUrl", params, optFns, c.addOperationGetDashboardEmbedUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDashboardEmbedUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDashboardEmbedUrlInput struct {

	// The ID for the Amazon Web Services account that contains the dashboard that
	// you're embedding.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard, also added to the Identity and Access Management
	// (IAM) policy.
	//
	// This member is required.
	DashboardId *string

	// The authentication method that the user uses to sign in.
	//
	// This member is required.
	IdentityType types.EmbeddingIdentityType

	// A list of one or more dashboard IDs that you want anonymous users to have
	// tempporary access to. Currently, the IdentityType parameter must be set to
	// ANONYMOUS because other identity types authenticate as Amazon QuickSight or IAM
	// users. For example, if you set " --dashboard-id dash_id1 --dashboard-id dash_id2
	// dash_id3 identity-type ANONYMOUS ", the session can access all three dashboards.
	AdditionalDashboardIds []string

	// The Amazon QuickSight namespace that contains the dashboard IDs in this
	// request. If you're not using a custom namespace, set Namespace = default .
	Namespace *string

	// Remove the reset button on the embedded dashboard. The default is FALSE, which
	// enables the reset button.
	ResetDisabled bool

	// How many minutes the session is valid. The session lifetime must be 15-600
	// minutes.
	SessionLifetimeInMinutes *int64

	// Adds persistence of state for the user session in an embedded dashboard.
	// Persistence applies to the sheet and the parameter settings. These are control
	// settings that the dashboard subscriber (Amazon QuickSight reader) chooses while
	// viewing the dashboard. If this is set to TRUE , the settings are the same when
	// the subscriber reopens the same dashboard URL. The state is stored in Amazon
	// QuickSight, not in a browser cookie. If this is set to FALSE, the state of the
	// user session is not persisted. The default is FALSE .
	StatePersistenceEnabled bool

	// Remove the undo/redo button on the embedded dashboard. The default is FALSE,
	// which enables the undo/redo button.
	UndoRedoDisabled bool

	// The Amazon QuickSight user's Amazon Resource Name (ARN), for use with QUICKSIGHT
	// identity type. You can use this for any Amazon QuickSight users in your account
	// (readers, authors, or admins) authenticated as one of the following:
	//   - Active Directory (AD) users or group members
	//   - Invited nonfederated users
	//   - IAM users and IAM role-based sessions authenticated through Federated
	//   Single Sign-On using SAML, OpenID Connect, or IAM federation.
	// Omit this parameter for users in the third group – IAM users and IAM role-based
	// sessions.
	UserArn *string

	noSmithyDocumentSerde
}

func (*GetDashboardEmbedUrlInput) operationName() string {
	return "GetDashboardEmbedUrl"
}

// Output returned from the GetDashboardEmbedUrl operation.
type GetDashboardEmbedUrlOutput struct {

	// A single-use URL that you can put into your server-side webpage to embed your
	// dashboard. This URL is valid for 5 minutes. The API operation provides the URL
	// with an auth_code value that enables one (and only one) sign-on to a user
	// session that is valid for 10 hours.
	EmbedUrl *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDashboardEmbedUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDashboardEmbedUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDashboardEmbedUrl{}, middleware.After)
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
	if err = addGetDashboardEmbedUrlResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetDashboardEmbedUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDashboardEmbedUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDashboardEmbedUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "GetDashboardEmbedUrl",
	}
}

type opGetDashboardEmbedUrlResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetDashboardEmbedUrlResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetDashboardEmbedUrlResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "quicksight"
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
				signingName = "quicksight"
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
				v4aScheme.SigningName = aws.String("quicksight")
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

func addGetDashboardEmbedUrlResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetDashboardEmbedUrlResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
