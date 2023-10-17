// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new configuration in the AppConfig hosted configuration store.
func (c *Client) CreateHostedConfigurationVersion(ctx context.Context, params *CreateHostedConfigurationVersionInput, optFns ...func(*Options)) (*CreateHostedConfigurationVersionOutput, error) {
	if params == nil {
		params = &CreateHostedConfigurationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHostedConfigurationVersion", params, optFns, c.addOperationCreateHostedConfigurationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHostedConfigurationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHostedConfigurationVersionInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The configuration profile ID.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The content of the configuration or the configuration data.
	//
	// This member is required.
	Content []byte

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type (https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17)
	// .
	//
	// This member is required.
	ContentType *string

	// A description of the configuration.
	Description *string

	// An optional locking token used to prevent race conditions from overwriting
	// configuration updates when creating a new version. To ensure your data is not
	// overwritten when creating multiple hosted configuration versions in rapid
	// succession, specify the version number of the latest hosted configuration
	// version.
	LatestVersionNumber *int32

	// An optional, user-defined label for the AppConfig hosted configuration version.
	// This value must contain at least one non-numeric character. For example,
	// "v2.2.0".
	VersionLabel *string

	noSmithyDocumentSerde
}

type CreateHostedConfigurationVersionOutput struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile ID.
	ConfigurationProfileId *string

	// The content of the configuration or the configuration data.
	Content []byte

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type (https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17)
	// .
	ContentType *string

	// A description of the configuration.
	Description *string

	// The Amazon Resource Name of the Key Management Service key that was used to
	// encrypt this specific version of the configuration data in the AppConfig hosted
	// configuration store.
	KmsKeyArn *string

	// A user-defined label for an AppConfig hosted configuration version.
	VersionLabel *string

	// The configuration version.
	VersionNumber int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHostedConfigurationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateHostedConfigurationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateHostedConfigurationVersion{}, middleware.After)
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
	if err = addOpCreateHostedConfigurationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHostedConfigurationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHostedConfigurationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "CreateHostedConfigurationVersion",
	}
}
