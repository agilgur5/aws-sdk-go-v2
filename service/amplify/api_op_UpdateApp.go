// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Amplify app.
func (c *Client) UpdateApp(ctx context.Context, params *UpdateAppInput, optFns ...func(*Options)) (*UpdateAppOutput, error) {
	if params == nil {
		params = &UpdateAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApp", params, optFns, c.addOperationUpdateAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the update app request.
type UpdateAppInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The personal access token for a GitHub repository for an Amplify app. The
	// personal access token is used to authorize access to a GitHub repository using
	// the Amplify GitHub App. The token is not stored. Use accessToken for GitHub
	// repositories only. To authorize access to a repository provider such as
	// Bitbucket or CodeCommit, use oauthToken . You must specify either accessToken
	// or oauthToken when you update an app. Existing Amplify apps deployed from a
	// GitHub repository using OAuth continue to work with CI/CD. However, we strongly
	// recommend that you migrate these apps to use the GitHub App. For more
	// information, see Migrating an existing OAuth app to the Amplify GitHub App (https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth)
	// in the Amplify User Guide .
	AccessToken *string

	// The automated branch creation configuration for an Amplify app.
	AutoBranchCreationConfig *types.AutoBranchCreationConfig

	// Describes the automated branch creation glob patterns for an Amplify app.
	AutoBranchCreationPatterns []string

	// The basic authorization credentials for an Amplify app. You must base64-encode
	// the authorization credentials and provide them in the format user:password .
	BasicAuthCredentials *string

	// The build specification (build spec) for an Amplify app.
	BuildSpec *string

	// The custom HTTP headers for an Amplify app.
	CustomHeaders *string

	// The custom redirect and rewrite rules for an Amplify app.
	CustomRules []types.CustomRule

	// The description for an Amplify app.
	Description *string

	// Enables automated branch creation for an Amplify app.
	EnableAutoBranchCreation *bool

	// Enables basic authorization for an Amplify app.
	EnableBasicAuth *bool

	// Enables branch auto-building for an Amplify app.
	EnableBranchAutoBuild *bool

	// Automatically disconnects a branch in the Amplify Console when you delete a
	// branch from your Git repository.
	EnableBranchAutoDeletion *bool

	// The environment variables for an Amplify app.
	EnvironmentVariables map[string]string

	// The AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn *string

	// The name for an Amplify app.
	Name *string

	// The OAuth token for a third-party source control system for an Amplify app. The
	// OAuth token is used to create a webhook and a read-only deploy key using SSH
	// cloning. The OAuth token is not stored. Use oauthToken for repository providers
	// other than GitHub, such as Bitbucket or CodeCommit. To authorize access to
	// GitHub as your repository provider, use accessToken . You must specify either
	// oauthToken or accessToken when you update an app. Existing Amplify apps
	// deployed from a GitHub repository using OAuth continue to work with CI/CD.
	// However, we strongly recommend that you migrate these apps to use the GitHub
	// App. For more information, see Migrating an existing OAuth app to the Amplify
	// GitHub App (https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth)
	// in the Amplify User Guide .
	OauthToken *string

	// The platform for the Amplify app. For a static app, set the platform type to WEB
	// . For a dynamic server-side rendered (SSR) app, set the platform type to
	// WEB_COMPUTE . For an app requiring Amplify Hosting's original SSR support only,
	// set the platform type to WEB_DYNAMIC .
	Platform types.Platform

	// The name of the repository for an Amplify app
	Repository *string

	noSmithyDocumentSerde
}

// The result structure for an Amplify app update request.
type UpdateAppOutput struct {

	// Represents the updated Amplify app.
	//
	// This member is required.
	App *types.App

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApp{}, middleware.After)
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
	if err = addOpUpdateAppValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "UpdateApp",
	}
}
