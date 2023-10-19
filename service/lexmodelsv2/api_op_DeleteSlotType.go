// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a slot type from a bot locale. If a slot is using the slot type, Amazon
// Lex throws a ResourceInUseException exception. To avoid the exception, set the
// skipResourceInUseCheck parameter to true .
func (c *Client) DeleteSlotType(ctx context.Context, params *DeleteSlotTypeInput, optFns ...func(*Options)) (*DeleteSlotTypeOutput, error) {
	if params == nil {
		params = &DeleteSlotTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSlotType", params, optFns, c.addOperationDeleteSlotTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSlotTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSlotTypeInput struct {

	// The identifier of the bot associated with the slot type.
	//
	// This member is required.
	BotId *string

	// The version of the bot associated with the slot type.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale that the slot type will be deleted
	// from. The string must match one of the supported locales. For more information,
	// see Supported languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	// .
	//
	// This member is required.
	LocaleId *string

	// The identifier of the slot type to delete.
	//
	// This member is required.
	SlotTypeId *string

	// By default, the DeleteSlotType operations throws a ResourceInUseException
	// exception if you try to delete a slot type used by a slot. Set the
	// skipResourceInUseCheck parameter to true to skip this check and remove the slot
	// type even if a slot uses it.
	SkipResourceInUseCheck bool

	noSmithyDocumentSerde
}

type DeleteSlotTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSlotTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSlotType{}, middleware.After)
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
	if err = addOpDeleteSlotTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSlotType(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteSlotType",
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

func newServiceMetadataMiddleware_opDeleteSlotType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteSlotType",
	}
}
