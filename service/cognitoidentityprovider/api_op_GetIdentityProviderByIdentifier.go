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

// Gets the specified IdP.
func (c *Client) GetIdentityProviderByIdentifier(ctx context.Context, params *GetIdentityProviderByIdentifierInput, optFns ...func(*Options)) (*GetIdentityProviderByIdentifierOutput, error) {
	if params == nil {
		params = &GetIdentityProviderByIdentifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentityProviderByIdentifier", params, optFns, c.addOperationGetIdentityProviderByIdentifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentityProviderByIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIdentityProviderByIdentifierInput struct {

	// The IdP identifier.
	//
	// This member is required.
	IdpIdentifier *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	noSmithyDocumentSerde
}

func (*GetIdentityProviderByIdentifierInput) operationName() string {
	return "GetIdentityProviderByIdentifier"
}

type GetIdentityProviderByIdentifierOutput struct {

	// The identity provider details.
	//
	// This member is required.
	IdentityProvider *types.IdentityProviderType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdentityProviderByIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetIdentityProviderByIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetIdentityProviderByIdentifier{}, middleware.After)
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
	if err = addOpGetIdentityProviderByIdentifierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityProviderByIdentifier(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdentityProviderByIdentifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GetIdentityProviderByIdentifier",
	}
}
