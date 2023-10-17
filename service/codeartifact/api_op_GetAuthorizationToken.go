// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Generates a temporary authorization token for accessing repositories in the
// domain. This API requires the codeartifact:GetAuthorizationToken and
// sts:GetServiceBearerToken permissions. For more information about authorization
// tokens, see CodeArtifact authentication and tokens (https://docs.aws.amazon.com/codeartifact/latest/ug/tokens-authentication.html)
// . CodeArtifact authorization tokens are valid for a period of 12 hours when
// created with the login command. You can call login periodically to refresh the
// token. When you create an authorization token with the GetAuthorizationToken
// API, you can set a custom authorization period, up to a maximum of 12 hours,
// with the durationSeconds parameter. The authorization period begins after login
// or GetAuthorizationToken is called. If login or GetAuthorizationToken is called
// while assuming a role, the token lifetime is independent of the maximum session
// duration of the role. For example, if you call sts assume-role and specify a
// session duration of 15 minutes, then generate a CodeArtifact authorization
// token, the token will be valid for the full authorization period even though
// this is longer than the 15-minute session duration. See Using IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html)
// for more information on controlling session duration.
func (c *Client) GetAuthorizationToken(ctx context.Context, params *GetAuthorizationTokenInput, optFns ...func(*Options)) (*GetAuthorizationTokenOutput, error) {
	if params == nil {
		params = &GetAuthorizationTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAuthorizationToken", params, optFns, c.addOperationGetAuthorizationTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAuthorizationTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAuthorizationTokenInput struct {

	// The name of the domain that is in scope for the generated authorization token.
	//
	// This member is required.
	Domain *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The time, in seconds, that the generated authorization token is valid. Valid
	// values are 0 and any number between 900 (15 minutes) and 43200 (12 hours). A
	// value of 0 will set the expiration of the authorization token to the same
	// expiration of the user's role's temporary credentials.
	DurationSeconds *int64

	noSmithyDocumentSerde
}

type GetAuthorizationTokenOutput struct {

	// The returned authentication token.
	AuthorizationToken *string

	// A timestamp that specifies the date and time the authorization token expires.
	Expiration *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAuthorizationTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAuthorizationToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAuthorizationToken{}, middleware.After)
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
	if err = addOpGetAuthorizationTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAuthorizationToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAuthorizationToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "GetAuthorizationToken",
	}
}
