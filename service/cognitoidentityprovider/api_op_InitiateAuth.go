// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates sign-in for a user in the Amazon Cognito user directory. You can't
// sign in a user with a federated IdP with InitiateAuth . For more information,
// see Adding user pool sign-in through a third party (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation.html)
// . Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// Using the Amazon Cognito native and OIDC APIs (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// . This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with Amazon Pinpoint (https://console.aws.amazon.com/pinpoint/home/)
// . Amazon Cognito uses the registered number automatically. Otherwise, Amazon
// Cognito users who must receive SMS messages might not be able to sign up,
// activate their accounts, or sign in. If you have never used SMS text messages
// with Amazon Cognito or any other Amazon Web Service, Amazon Simple Notification
// Service might place your account in the SMS sandbox. In sandbox mode (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html)
// , you can send messages only to verified phone numbers. After you test your app
// while in the sandbox environment, you can move out of the sandbox and into
// production. For more information, see SMS message settings for Amazon Cognito
// user pools (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html)
// in the Amazon Cognito Developer Guide.
func (c *Client) InitiateAuth(ctx context.Context, params *InitiateAuthInput, optFns ...func(*Options)) (*InitiateAuthOutput, error) {
	if params == nil {
		params = &InitiateAuthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InitiateAuth", params, optFns, c.addOperationInitiateAuthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InitiateAuthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Initiates the authentication request.
type InitiateAuthInput struct {

	// The authentication flow for this call to run. The API action will depend on
	// this value. For example:
	//   - REFRESH_TOKEN_AUTH takes in a valid refresh token and returns new tokens.
	//   - USER_SRP_AUTH takes in USERNAME and SRP_A and returns the SRP variables to
	//   be used for next challenge execution.
	//   - USER_PASSWORD_AUTH takes in USERNAME and PASSWORD and returns the next
	//   challenge or tokens.
	// Valid values include:
	//   - USER_SRP_AUTH : Authentication flow for the Secure Remote Password (SRP)
	//   protocol.
	//   - REFRESH_TOKEN_AUTH / REFRESH_TOKEN : Authentication flow for refreshing the
	//   access token and ID token by supplying a valid refresh token.
	//   - CUSTOM_AUTH : Custom authentication flow.
	//   - USER_PASSWORD_AUTH : Non-SRP authentication flow; user name and password are
	//   passed directly. If a user migration Lambda trigger is set, this flow will
	//   invoke the user migration Lambda if it doesn't find the user name in the user
	//   pool.
	// ADMIN_NO_SRP_AUTH isn't a valid value.
	//
	// This member is required.
	AuthFlow types.AuthFlowType

	// The app client ID.
	//
	// This member is required.
	ClientId *string

	// The Amazon Pinpoint analytics metadata that contributes to your metrics for
	// InitiateAuth calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// The authentication parameters. These are inputs corresponding to the AuthFlow
	// that you're invoking. The required values depend on the value of AuthFlow :
	//   - For USER_SRP_AUTH : USERNAME (required), SRP_A (required), SECRET_HASH
	//   (required if the app client is configured with a client secret), DEVICE_KEY .
	//   - For USER_PASSWORD_AUTH : USERNAME (required), PASSWORD (required),
	//   SECRET_HASH (required if the app client is configured with a client secret),
	//   DEVICE_KEY .
	//   - For REFRESH_TOKEN_AUTH/REFRESH_TOKEN : REFRESH_TOKEN (required), SECRET_HASH
	//   (required if the app client is configured with a client secret), DEVICE_KEY .
	//   - For CUSTOM_AUTH : USERNAME (required), SECRET_HASH (if app client is
	//   configured with client secret), DEVICE_KEY . To start the authentication flow
	//   with password verification, include ChallengeName: SRP_A and SRP_A: (The
	//   SRP_A Value) .
	// For more information about SECRET_HASH , see Computing secret hash values (https://docs.aws.amazon.com/cognito/latest/developerguide/signing-up-users-in-your-app.html#cognito-user-pools-computing-secret-hash)
	// . For information about DEVICE_KEY , see Working with user devices in your user
	// pool (https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html)
	// .
	AuthParameters map[string]string

	// A map of custom key-value pairs that you can provide as input for certain
	// custom workflows that this action triggers. You create custom workflows by
	// assigning Lambda functions to user pool triggers. When you use the InitiateAuth
	// API action, Amazon Cognito invokes the Lambda functions that are specified for
	// various triggers. The ClientMetadata value is passed as input to the functions
	// for only the following triggers:
	//   - Pre signup
	//   - Pre authentication
	//   - User migration
	// When Amazon Cognito invokes the functions for these triggers, it passes a JSON
	// payload, which the function receives as input. This payload contains a
	// validationData attribute, which provides the data that you assigned to the
	// ClientMetadata parameter in your InitiateAuth request. In your function code in
	// Lambda, you can process the validationData value to enhance your workflow for
	// your specific needs. When you use the InitiateAuth API action, Amazon Cognito
	// also invokes the functions for the following triggers, but it doesn't provide
	// the ClientMetadata value as input:
	//   - Post authentication
	//   - Custom message
	//   - Pre token generation
	//   - Create auth challenge
	//   - Define auth challenge
	//   - Verify auth challenge
	// For more information, see  Customizing user pool Workflows with Lambda Triggers (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. When you use the ClientMetadata
	// parameter, remember that Amazon Cognito won't do the following:
	//   - Store the ClientMetadata value. This data is available only to Lambda
	//   triggers that are assigned to a user pool to support custom workflows. If your
	//   user pool configuration doesn't include triggers, the ClientMetadata parameter
	//   serves no purpose.
	//   - Validate the ClientMetadata value.
	//   - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide
	//   sensitive information.
	ClientMetadata map[string]string

	// Contextual data about your user session, such as the device fingerprint, IP
	// address, or location. Amazon Cognito advanced security evaluates the risk of an
	// authentication event based on the context that your app generates and passes to
	// Amazon Cognito when it makes API requests.
	UserContextData *types.UserContextDataType

	noSmithyDocumentSerde
}

func (*InitiateAuthInput) operationName() string {
	return "InitiateAuth"
}

// Initiates the authentication response.
type InitiateAuthOutput struct {

	// The result of the authentication response. This result is only returned if the
	// caller doesn't need to pass another challenge. If the caller does need to pass
	// another challenge before it gets tokens, ChallengeName , ChallengeParameters ,
	// and Session are returned.
	AuthenticationResult *types.AuthenticationResultType

	// The name of the challenge that you're responding to with this call. This name
	// is returned in the AdminInitiateAuth response if you must pass another
	// challenge. Valid values include the following: All of the following challenges
	// require USERNAME and SECRET_HASH (if applicable) in the parameters.
	//   - SMS_MFA : Next challenge is to supply an SMS_MFA_CODE , delivered via SMS.
	//   - PASSWORD_VERIFIER : Next challenge is to supply PASSWORD_CLAIM_SIGNATURE ,
	//   PASSWORD_CLAIM_SECRET_BLOCK , and TIMESTAMP after the client-side SRP
	//   calculations.
	//   - CUSTOM_CHALLENGE : This is returned if your custom authentication flow
	//   determines that the user should pass another challenge before tokens are issued.
	//
	//   - DEVICE_SRP_AUTH : If device tracking was activated on your user pool and the
	//   previous challenges were passed, this challenge is returned so that Amazon
	//   Cognito can start tracking this device.
	//   - DEVICE_PASSWORD_VERIFIER : Similar to PASSWORD_VERIFIER , but for devices
	//   only.
	//   - NEW_PASSWORD_REQUIRED : For users who are required to change their passwords
	//   after successful first login. Respond to this challenge with NEW_PASSWORD and
	//   any required attributes that Amazon Cognito returned in the requiredAttributes
	//   parameter. You can also set values for attributes that aren't required by your
	//   user pool and that your app client can write. For more information, see
	//   RespondToAuthChallenge (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_RespondToAuthChallenge.html)
	//   . In a NEW_PASSWORD_REQUIRED challenge response, you can't modify a required
	//   attribute that already has a value. In RespondToAuthChallenge , set a value
	//   for any keys that Amazon Cognito returned in the requiredAttributes parameter,
	//   then use the UpdateUserAttributes API operation to modify the value of any
	//   additional attributes.
	//   - MFA_SETUP : For users who are required to setup an MFA factor before they
	//   can sign in. The MFA types activated for the user pool will be listed in the
	//   challenge parameters MFA_CAN_SETUP value. To set up software token MFA, use
	//   the session returned here from InitiateAuth as an input to
	//   AssociateSoftwareToken . Use the session returned by VerifySoftwareToken as an
	//   input to RespondToAuthChallenge with challenge name MFA_SETUP to complete
	//   sign-in. To set up SMS MFA, an administrator should help the user to add a phone
	//   number to their account, and then the user should call InitiateAuth again to
	//   restart sign-in.
	ChallengeName types.ChallengeNameType

	// The challenge parameters. These are returned in the InitiateAuth response if
	// you must pass another challenge. The responses in this parameter should be used
	// to compute inputs to the next call ( RespondToAuthChallenge ). All challenges
	// require USERNAME and SECRET_HASH (if applicable).
	ChallengeParameters map[string]string

	// The session that should pass both ways in challenge-response calls to the
	// service. If the caller must pass another challenge, they return a session with
	// other challenge parameters. This session should be passed as it is to the next
	// RespondToAuthChallenge API call.
	Session *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInitiateAuthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInitiateAuth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInitiateAuth{}, middleware.After)
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
	if err = addInitiateAuthResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpInitiateAuthValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateAuth(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInitiateAuth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InitiateAuth",
	}
}

type opInitiateAuthResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opInitiateAuthResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opInitiateAuthResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "cognito-idp"
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
				signingName = "cognito-idp"
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
				v4aScheme.SigningName = aws.String("cognito-idp")
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

func addInitiateAuthResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opInitiateAuthResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
