// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the password of the IAM user who is calling this operation. This
// operation can be performed using the CLI, the Amazon Web Services API, or the My
// Security Credentials page in the Amazon Web Services Management Console. The
// Amazon Web Services account root user password is not affected by this
// operation. Use UpdateLoginProfile to use the CLI, the Amazon Web Services API,
// or the Users page in the IAM console to change the password for any IAM user.
// For more information about modifying passwords, see Managing passwords (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html)
// in the IAM User Guide.
func (c *Client) ChangePassword(ctx context.Context, params *ChangePasswordInput, optFns ...func(*Options)) (*ChangePasswordOutput, error) {
	if params == nil {
		params = &ChangePasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangePassword", params, optFns, c.addOperationChangePasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangePasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ChangePasswordInput struct {

	// The new password. The new password must conform to the Amazon Web Services
	// account's password policy, if one exists. The regex pattern (http://wikipedia.org/wiki/regex)
	// that is used to validate this parameter is a string of characters. That string
	// can include almost any printable ASCII character from the space ( \u0020 )
	// through the end of the ASCII character range ( \u00FF ). You can also include
	// the tab ( \u0009 ), line feed ( \u000A ), and carriage return ( \u000D )
	// characters. Any of these characters are valid in a password. However, many
	// tools, such as the Amazon Web Services Management Console, might restrict the
	// ability to type certain characters because they have special meaning within that
	// tool.
	//
	// This member is required.
	NewPassword *string

	// The IAM user's current password.
	//
	// This member is required.
	OldPassword *string

	noSmithyDocumentSerde
}

type ChangePasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationChangePasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpChangePassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpChangePassword{}, middleware.After)
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
	if err = addOpChangePasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChangePassword(options.Region), middleware.Before); err != nil {
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
		operation: "ChangePassword",
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

func newServiceMetadataMiddleware_opChangePassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ChangePassword",
	}
}
