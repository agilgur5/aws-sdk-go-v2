// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You can use this operation to use default (username and clientID) attribute or
// custom attribute mappings.
func (c *Client) SetPrincipalTagAttributeMap(ctx context.Context, params *SetPrincipalTagAttributeMapInput, optFns ...func(*Options)) (*SetPrincipalTagAttributeMapOutput, error) {
	if params == nil {
		params = &SetPrincipalTagAttributeMapInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetPrincipalTagAttributeMap", params, optFns, c.addOperationSetPrincipalTagAttributeMapMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetPrincipalTagAttributeMapOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetPrincipalTagAttributeMapInput struct {

	// The ID of the Identity Pool you want to set attribute mappings for.
	//
	// This member is required.
	IdentityPoolId *string

	// The provider name you want to use for attribute mappings.
	//
	// This member is required.
	IdentityProviderName *string

	// You can use this operation to add principal tags.
	PrincipalTags map[string]string

	// You can use this operation to use default (username and clientID) attribute
	// mappings.
	UseDefaults *bool

	noSmithyDocumentSerde
}

type SetPrincipalTagAttributeMapOutput struct {

	// The ID of the Identity Pool you want to set attribute mappings for.
	IdentityPoolId *string

	// The provider name you want to use for attribute mappings.
	IdentityProviderName *string

	// You can use this operation to add principal tags. The PrincipalTags operation
	// enables you to reference user attributes in your IAM permissions policy.
	PrincipalTags map[string]string

	// You can use this operation to select default (username and clientID) attribute
	// mappings.
	UseDefaults *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetPrincipalTagAttributeMapMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetPrincipalTagAttributeMap{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetPrincipalTagAttributeMap{}, middleware.After)
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
	if err = addOpSetPrincipalTagAttributeMapValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetPrincipalTagAttributeMap(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetPrincipalTagAttributeMap(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "SetPrincipalTagAttributeMap",
	}
}
