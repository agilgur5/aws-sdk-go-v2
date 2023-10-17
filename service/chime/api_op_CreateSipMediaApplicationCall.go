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

// Creates an outbound call to a phone number from the phone number specified in
// the request, and it invokes the endpoint of the specified sipMediaApplicationId
// . This API is is no longer supported and will not be updated. We recommend using
// the latest version, CreateSipMediaApplicationCall (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_CreateSipMediaApplicationCall.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by CreateSipMediaApplicationCall in the Amazon Chime SDK
// Voice Namespace
func (c *Client) CreateSipMediaApplicationCall(ctx context.Context, params *CreateSipMediaApplicationCallInput, optFns ...func(*Options)) (*CreateSipMediaApplicationCallOutput, error) {
	if params == nil {
		params = &CreateSipMediaApplicationCallInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSipMediaApplicationCall", params, optFns, c.addOperationCreateSipMediaApplicationCallMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSipMediaApplicationCallOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSipMediaApplicationCallInput struct {

	// The phone number that a user calls from. This is a phone number in your Amazon
	// Chime phone number inventory.
	//
	// This member is required.
	FromPhoneNumber *string

	// The ID of the SIP media application.
	//
	// This member is required.
	SipMediaApplicationId *string

	// The phone number that the service should call.
	//
	// This member is required.
	ToPhoneNumber *string

	// The SIP headers added to an outbound call leg.
	SipHeaders map[string]string

	noSmithyDocumentSerde
}

type CreateSipMediaApplicationCallOutput struct {

	// The actual call.
	SipMediaApplicationCall *types.SipMediaApplicationCall

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSipMediaApplicationCallMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSipMediaApplicationCall{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSipMediaApplicationCall{}, middleware.After)
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
	if err = addOpCreateSipMediaApplicationCallValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSipMediaApplicationCall(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSipMediaApplicationCall(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateSipMediaApplicationCall",
	}
}
