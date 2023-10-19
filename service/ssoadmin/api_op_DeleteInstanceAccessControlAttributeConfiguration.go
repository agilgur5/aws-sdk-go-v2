// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables the attributes-based access control (ABAC) feature for the specified
// IAM Identity Center instance and deletes all of the attribute mappings that have
// been configured. Once deleted, any attributes that are received from an identity
// source and any custom attributes you have previously configured will not be
// passed. For more information about ABAC, see Attribute-Based Access Control in
// the IAM Identity Center User Guide.
func (c *Client) DeleteInstanceAccessControlAttributeConfiguration(ctx context.Context, params *DeleteInstanceAccessControlAttributeConfigurationInput, optFns ...func(*Options)) (*DeleteInstanceAccessControlAttributeConfigurationOutput, error) {
	if params == nil {
		params = &DeleteInstanceAccessControlAttributeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteInstanceAccessControlAttributeConfiguration", params, optFns, c.addOperationDeleteInstanceAccessControlAttributeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteInstanceAccessControlAttributeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInstanceAccessControlAttributeConfigurationInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed.
	//
	// This member is required.
	InstanceArn *string

	noSmithyDocumentSerde
}

type DeleteInstanceAccessControlAttributeConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteInstanceAccessControlAttributeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteInstanceAccessControlAttributeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteInstanceAccessControlAttributeConfiguration{}, middleware.After)
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
	if err = addOpDeleteInstanceAccessControlAttributeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInstanceAccessControlAttributeConfiguration(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteInstanceAccessControlAttributeConfiguration",
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

func newServiceMetadataMiddleware_opDeleteInstanceAccessControlAttributeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "DeleteInstanceAccessControlAttributeConfiguration",
	}
}
