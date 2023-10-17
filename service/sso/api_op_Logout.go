// Code generated by smithy-go-codegen DO NOT EDIT.

package sso

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the locally stored SSO tokens from the client-side cache and sends an
// API call to the IAM Identity Center service to invalidate the corresponding
// server-side IAM Identity Center sign in session. If a user uses IAM Identity
// Center to access the AWS CLI, the user’s IAM Identity Center sign in session is
// used to obtain an IAM session, as specified in the corresponding IAM Identity
// Center permission set. More specifically, IAM Identity Center assumes an IAM
// role in the target account on behalf of the user, and the corresponding
// temporary AWS credentials are returned to the client. After user logout, any
// existing IAM role sessions that were created by using IAM Identity Center
// permission sets continue based on the duration configured in the permission set.
// For more information, see User authentications (https://docs.aws.amazon.com/singlesignon/latest/userguide/authconcept.html)
// in the IAM Identity Center User Guide.
func (c *Client) Logout(ctx context.Context, params *LogoutInput, optFns ...func(*Options)) (*LogoutOutput, error) {
	if params == nil {
		params = &LogoutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Logout", params, optFns, c.addOperationLogoutMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*LogoutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type LogoutInput struct {

	// The token issued by the CreateToken API call. For more information, see
	// CreateToken (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the IAM Identity Center OIDC API Reference Guide.
	//
	// This member is required.
	AccessToken *string

	noSmithyDocumentSerde
}

type LogoutOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationLogoutMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpLogout{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpLogout{}, middleware.After)
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
	if err = addOpLogoutValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opLogout(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opLogout(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Logout",
	}
}
