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

// Puts the specified proxy configuration to the specified Amazon Chime Voice
// Connector. This API is is no longer supported and will not be updated. We
// recommend using the latest version, PutVoiceConnectorProxy (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_PutVoiceConnectorProxy.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by PutVoiceConnectorProxy in the Amazon Chime SDK Voice
// Namespace
func (c *Client) PutVoiceConnectorProxy(ctx context.Context, params *PutVoiceConnectorProxyInput, optFns ...func(*Options)) (*PutVoiceConnectorProxyOutput, error) {
	if params == nil {
		params = &PutVoiceConnectorProxyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutVoiceConnectorProxy", params, optFns, c.addOperationPutVoiceConnectorProxyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutVoiceConnectorProxyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorProxyInput struct {

	// The default number of minutes allowed for proxy sessions.
	//
	// This member is required.
	DefaultSessionExpiryMinutes *int32

	// The countries for proxy phone numbers to be selected from.
	//
	// This member is required.
	PhoneNumberPoolCountries []string

	// The Amazon Chime voice connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// When true, stops proxy sessions from being created on the specified Amazon
	// Chime Voice Connector.
	Disabled *bool

	// The phone number to route calls to after a proxy session expires.
	FallBackPhoneNumber *string

	noSmithyDocumentSerde
}

type PutVoiceConnectorProxyOutput struct {

	// The proxy configuration details.
	Proxy *types.Proxy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutVoiceConnectorProxyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutVoiceConnectorProxy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutVoiceConnectorProxy{}, middleware.After)
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
	if err = addOpPutVoiceConnectorProxyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutVoiceConnectorProxy(options.Region), middleware.Before); err != nil {
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
		operation: "PutVoiceConnectorProxy",
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

func newServiceMetadataMiddleware_opPutVoiceConnectorProxy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorProxy",
	}
}
