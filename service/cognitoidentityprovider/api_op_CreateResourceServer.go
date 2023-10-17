// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new OAuth2.0 resource server and defines custom scopes within it.
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) CreateResourceServer(ctx context.Context, params *CreateResourceServerInput, optFns ...func(*Options)) (*CreateResourceServerOutput, error) {
	if params == nil {
		params = &CreateResourceServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResourceServer", params, optFns, c.addOperationCreateResourceServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourceServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceServerInput struct {

	// A unique resource server identifier for the resource server. This could be an
	// HTTPS endpoint where the resource server is located, such as
	// https://my-weather-api.example.com .
	//
	// This member is required.
	Identifier *string

	// A friendly name for the resource server.
	//
	// This member is required.
	Name *string

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// A list of scopes. Each scope is a key-value map with the keys name and
	// description .
	Scopes []types.ResourceServerScopeType

	noSmithyDocumentSerde
}

type CreateResourceServerOutput struct {

	// The newly created resource server.
	//
	// This member is required.
	ResourceServer *types.ResourceServerType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResourceServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateResourceServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateResourceServer{}, middleware.After)
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
	if err = addOpCreateResourceServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResourceServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateResourceServer",
	}
}
