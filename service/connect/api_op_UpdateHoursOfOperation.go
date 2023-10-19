// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Updates the hours of operation.
func (c *Client) UpdateHoursOfOperation(ctx context.Context, params *UpdateHoursOfOperationInput, optFns ...func(*Options)) (*UpdateHoursOfOperationOutput, error) {
	if params == nil {
		params = &UpdateHoursOfOperationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateHoursOfOperation", params, optFns, c.addOperationUpdateHoursOfOperationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateHoursOfOperationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateHoursOfOperationInput struct {

	// The identifier of the hours of operation.
	//
	// This member is required.
	HoursOfOperationId *string

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// Configuration information of the hours of operation.
	Config []types.HoursOfOperationConfig

	// The description of the hours of operation.
	Description *string

	// The name of the hours of operation.
	Name *string

	// The time zone of the hours of operation.
	TimeZone *string

	noSmithyDocumentSerde
}

type UpdateHoursOfOperationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateHoursOfOperationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateHoursOfOperation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateHoursOfOperation{}, middleware.After)
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
	if err = addOpUpdateHoursOfOperationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateHoursOfOperation(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateHoursOfOperation",
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

func newServiceMetadataMiddleware_opUpdateHoursOfOperation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateHoursOfOperation",
	}
}
