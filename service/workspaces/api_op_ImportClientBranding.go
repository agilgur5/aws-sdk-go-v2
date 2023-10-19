// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports client branding. Client branding allows you to customize your
// WorkSpace's client login portal. You can tailor your login portal company logo,
// the support email address, support link, link to reset password, and a custom
// message for users trying to sign in. After you import client branding, the
// default branding experience for the specified platform type is replaced with the
// imported experience
//   - You must specify at least one platform type when importing client branding.
//   - You can import up to 6 MB of data with each request. If your request
//     exceeds this limit, you can import client branding for different platform types
//     using separate requests.
//   - In each platform type, the SupportEmail and SupportLink parameters are
//     mutually exclusive. You can specify only one parameter for each platform type,
//     but not both.
//   - Imported data can take up to a minute to appear in the WorkSpaces client.
func (c *Client) ImportClientBranding(ctx context.Context, params *ImportClientBrandingInput, optFns ...func(*Options)) (*ImportClientBrandingOutput, error) {
	if params == nil {
		params = &ImportClientBrandingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportClientBranding", params, optFns, c.addOperationImportClientBrandingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportClientBrandingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportClientBrandingInput struct {

	// The directory identifier of the WorkSpace for which you want to import client
	// branding.
	//
	// This member is required.
	ResourceId *string

	// The branding information to import for Android devices.
	DeviceTypeAndroid *types.DefaultImportClientBrandingAttributes

	// The branding information to import for iOS devices.
	DeviceTypeIos *types.IosImportClientBrandingAttributes

	// The branding information to import for Linux devices.
	DeviceTypeLinux *types.DefaultImportClientBrandingAttributes

	// The branding information to import for macOS devices.
	DeviceTypeOsx *types.DefaultImportClientBrandingAttributes

	// The branding information to import for web access.
	DeviceTypeWeb *types.DefaultImportClientBrandingAttributes

	// The branding information to import for Windows devices.
	DeviceTypeWindows *types.DefaultImportClientBrandingAttributes

	noSmithyDocumentSerde
}

type ImportClientBrandingOutput struct {

	// The branding information configured for Android devices.
	DeviceTypeAndroid *types.DefaultClientBrandingAttributes

	// The branding information configured for iOS devices.
	DeviceTypeIos *types.IosClientBrandingAttributes

	// The branding information configured for Linux devices.
	DeviceTypeLinux *types.DefaultClientBrandingAttributes

	// The branding information configured for macOS devices.
	DeviceTypeOsx *types.DefaultClientBrandingAttributes

	// The branding information configured for web access.
	DeviceTypeWeb *types.DefaultClientBrandingAttributes

	// The branding information configured for Windows devices.
	DeviceTypeWindows *types.DefaultClientBrandingAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportClientBrandingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportClientBranding{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportClientBranding{}, middleware.After)
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
	if err = addOpImportClientBrandingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportClientBranding(options.Region), middleware.Before); err != nil {
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
		operation: "ImportClientBranding",
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

func newServiceMetadataMiddleware_opImportClientBranding(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ImportClientBranding",
	}
}
