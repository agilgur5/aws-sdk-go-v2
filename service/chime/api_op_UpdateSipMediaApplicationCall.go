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

// Invokes the AWS Lambda function associated with the SIP media application and
// transaction ID in an update request. The Lambda function can then return a new
// set of actions. This API is is no longer supported and will not be updated. We
// recommend using the latest version, UpdateSipMediaApplicationCall (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_UpdateSipMediaApplicationCall.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by UpdateSipMediaApplicationCall in the Amazon Chime SDK
// Voice Namespace
func (c *Client) UpdateSipMediaApplicationCall(ctx context.Context, params *UpdateSipMediaApplicationCallInput, optFns ...func(*Options)) (*UpdateSipMediaApplicationCallOutput, error) {
	if params == nil {
		params = &UpdateSipMediaApplicationCallInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSipMediaApplicationCall", params, optFns, c.addOperationUpdateSipMediaApplicationCallMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSipMediaApplicationCallOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSipMediaApplicationCallInput struct {

	// Arguments made available to the Lambda function as part of the
	// CALL_UPDATE_REQUESTED event. Can contain 0-20 key-value pairs.
	//
	// This member is required.
	Arguments map[string]string

	// The ID of the SIP media application handling the call.
	//
	// This member is required.
	SipMediaApplicationId *string

	// The ID of the call transaction.
	//
	// This member is required.
	TransactionId *string

	noSmithyDocumentSerde
}

type UpdateSipMediaApplicationCallOutput struct {

	// A Call instance for a SIP media application.
	SipMediaApplicationCall *types.SipMediaApplicationCall

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSipMediaApplicationCallMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSipMediaApplicationCall{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSipMediaApplicationCall{}, middleware.After)
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
	if err = addOpUpdateSipMediaApplicationCallValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSipMediaApplicationCall(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateSipMediaApplicationCall",
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

func newServiceMetadataMiddleware_opUpdateSipMediaApplicationCall(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateSipMediaApplicationCall",
	}
}
