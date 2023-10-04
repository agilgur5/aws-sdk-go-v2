// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the specified user by user name in a user pool as an administrator. Works
// on any user. Amazon Cognito evaluates Identity and Access Management (IAM)
// policies in requests for this API operation. For this operation, you must use
// IAM credentials to authorize requests, and you must grant yourself the
// corresponding IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) AdminGetUser(ctx context.Context, params *AdminGetUserInput, optFns ...func(*Options)) (*AdminGetUserOutput, error) {
	if params == nil {
		params = &AdminGetUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminGetUser", params, optFns, c.addOperationAdminGetUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminGetUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get the specified user as an administrator.
type AdminGetUserInput struct {

	// The user pool ID for the user pool where you want to get information about the
	// user.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user you want to retrieve.
	//
	// This member is required.
	Username *string

	noSmithyDocumentSerde
}

func (*AdminGetUserInput) operationName() string {
	return "AdminGetUser"
}

// Represents the response from the server from the request to get the specified
// user as an administrator.
type AdminGetUserOutput struct {

	// The username of the user that you requested.
	//
	// This member is required.
	Username *string

	// Indicates that the status is enabled .
	Enabled bool

	// This response parameter is no longer supported. It provides information only
	// about SMS MFA configurations. It doesn't provide information about time-based
	// one-time password (TOTP) software token MFA configurations. To look up
	// information about either type of MFA configuration, use UserMFASettingList
	// instead.
	MFAOptions []types.MFAOptionType

	// The user's preferred MFA setting.
	PreferredMfaSetting *string

	// An array of name-value pairs representing user attributes.
	UserAttributes []types.AttributeType

	// The date the user was created.
	UserCreateDate *time.Time

	// The date and time, in ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html)
	// format, when the item was modified.
	UserLastModifiedDate *time.Time

	// The MFA options that are activated for the user. The possible values in this
	// list are SMS_MFA and SOFTWARE_TOKEN_MFA .
	UserMFASettingList []string

	// The user status. Can be one of the following:
	//   - UNCONFIRMED - User has been created but not confirmed.
	//   - CONFIRMED - User has been confirmed.
	//   - UNKNOWN - User status isn't known.
	//   - RESET_REQUIRED - User is confirmed, but the user must request a code and
	//   reset their password before they can sign in.
	//   - FORCE_CHANGE_PASSWORD - The user is confirmed and the user can sign in
	//   using a temporary password, but on first sign-in, the user must change their
	//   password to a new value before doing anything else.
	UserStatus types.UserStatusType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminGetUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminGetUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminGetUser{}, middleware.After)
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
	if err = addOpAdminGetUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminGetUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminGetUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminGetUser",
	}
}
