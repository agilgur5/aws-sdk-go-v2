// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the logging configuration for the specified SIP media application. This
// API is is no longer supported and will not be updated. We recommend using the
// latest version, GetSipMediaApplicationLoggingConfiguration (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_GetSipMediaApplicationLoggingConfiguration.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by GetSipMediaApplicationLoggingConfiguration in the
// Amazon Chime SDK Voice Namespace
func (c *Client) GetSipMediaApplicationLoggingConfiguration(ctx context.Context, params *GetSipMediaApplicationLoggingConfigurationInput, optFns ...func(*Options)) (*GetSipMediaApplicationLoggingConfigurationOutput, error) {
	if params == nil {
		params = &GetSipMediaApplicationLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSipMediaApplicationLoggingConfiguration", params, optFns, c.addOperationGetSipMediaApplicationLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSipMediaApplicationLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSipMediaApplicationLoggingConfigurationInput struct {

	// The SIP media application ID.
	//
	// This member is required.
	SipMediaApplicationId *string

	noSmithyDocumentSerde
}

type GetSipMediaApplicationLoggingConfigurationOutput struct {

	// The actual logging configuration.
	SipMediaApplicationLoggingConfiguration *types.SipMediaApplicationLoggingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSipMediaApplicationLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSipMediaApplicationLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSipMediaApplicationLoggingConfiguration{}, middleware.After)
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
	if err = addOpGetSipMediaApplicationLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSipMediaApplicationLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSipMediaApplicationLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetSipMediaApplicationLoggingConfiguration",
	}
}
