// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action is no longer supported. You can use it to configure only SMS MFA.
// You can't use it to configure time-based one-time password (TOTP) software token
// MFA. To configure either type of MFA, use SetUserMFAPreference (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetUserMFAPreference.html)
// instead. Amazon Cognito doesn't evaluate Identity and Access Management (IAM)
// policies in requests for this API operation. For this operation, you can't use
// IAM credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// Using the Amazon Cognito native and OIDC APIs (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// .
func (c *Client) SetUserSettings(ctx context.Context, params *SetUserSettingsInput, optFns ...func(*Options)) (*SetUserSettingsOutput, error) {
	if params == nil {
		params = &SetUserSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUserSettings", params, optFns, c.addOperationSetUserSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to set user settings.
type SetUserSettingsInput struct {

	// A valid access token that Amazon Cognito issued to the user whose user settings
	// you want to configure.
	//
	// This member is required.
	AccessToken *string

	// You can use this parameter only to set an SMS configuration that uses SMS for
	// delivery.
	//
	// This member is required.
	MFAOptions []types.MFAOptionType

	noSmithyDocumentSerde
}

// The response from the server for a set user settings request.
type SetUserSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetUserSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetUserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUserSettings{}, middleware.After)
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
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpSetUserSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetUserSettings(options.Region), middleware.Before); err != nil {
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
		operation: "SetUserSettings",
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

func newServiceMetadataMiddleware_opSetUserSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetUserSettings",
	}
}
