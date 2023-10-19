// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the definition for a dimension. You cannot change the type of a
// dimension after it is created (you can delete it and recreate it). Requires
// permission to access the UpdateDimension (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) UpdateDimension(ctx context.Context, params *UpdateDimensionInput, optFns ...func(*Options)) (*UpdateDimensionOutput, error) {
	if params == nil {
		params = &UpdateDimensionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDimension", params, optFns, c.addOperationUpdateDimensionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDimensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDimensionInput struct {

	// A unique identifier for the dimension. Choose something that describes the type
	// and value to make it easy to remember what it does.
	//
	// This member is required.
	Name *string

	// Specifies the value or list of values for the dimension. For TOPIC_FILTER
	// dimensions, this is a pattern used to match the MQTT topic (for example,
	// "admin/#").
	//
	// This member is required.
	StringValues []string

	noSmithyDocumentSerde
}

type UpdateDimensionOutput struct {

	// The Amazon Resource Name (ARN)of the created dimension.
	Arn *string

	// The date and time, in milliseconds since epoch, when the dimension was
	// initially created.
	CreationDate *time.Time

	// The date and time, in milliseconds since epoch, when the dimension was most
	// recently updated.
	LastModifiedDate *time.Time

	// A unique identifier for the dimension.
	Name *string

	// The value or list of values used to scope the dimension. For example, for topic
	// filters, this is the pattern used to match the MQTT topic name.
	StringValues []string

	// The type of the dimension.
	Type types.DimensionType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDimensionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDimension{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDimension{}, middleware.After)
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
	if err = addOpUpdateDimensionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDimension(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateDimension",
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

func newServiceMetadataMiddleware_opUpdateDimension(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "UpdateDimension",
	}
}
