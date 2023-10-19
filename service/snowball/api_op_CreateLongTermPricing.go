// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a job with the long-term usage option for a device. The long-term usage
// is a 1-year or 3-year long-term pricing type for the device. You are billed
// upfront, and Amazon Web Services provides discounts for long-term pricing.
func (c *Client) CreateLongTermPricing(ctx context.Context, params *CreateLongTermPricingInput, optFns ...func(*Options)) (*CreateLongTermPricingOutput, error) {
	if params == nil {
		params = &CreateLongTermPricingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLongTermPricing", params, optFns, c.addOperationCreateLongTermPricingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLongTermPricingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLongTermPricingInput struct {

	// The type of long-term pricing option you want for the device, either 1-year or
	// 3-year long-term pricing.
	//
	// This member is required.
	LongTermPricingType types.LongTermPricingType

	// The type of Snow Family devices to use for the long-term pricing job.
	//
	// This member is required.
	SnowballType types.SnowballType

	// Specifies whether the current long-term pricing type for the device should be
	// renewed.
	IsLongTermPricingAutoRenew *bool

	noSmithyDocumentSerde
}

type CreateLongTermPricingOutput struct {

	// The ID of the long-term pricing type for the device.
	LongTermPricingId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLongTermPricingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLongTermPricing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLongTermPricing{}, middleware.After)
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
	if err = addOpCreateLongTermPricingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLongTermPricing(options.Region), middleware.Before); err != nil {
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
		operation: "CreateLongTermPricing",
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

func newServiceMetadataMiddleware_opCreateLongTermPricing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "CreateLongTermPricing",
	}
}
