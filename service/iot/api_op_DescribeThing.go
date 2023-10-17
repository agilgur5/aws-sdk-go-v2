// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the specified thing. Requires permission to access the
// DescribeThing (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DescribeThing(ctx context.Context, params *DescribeThingInput, optFns ...func(*Options)) (*DescribeThingOutput, error) {
	if params == nil {
		params = &DescribeThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeThing", params, optFns, c.addOperationDescribeThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeThing operation.
type DescribeThingInput struct {

	// The name of the thing.
	//
	// This member is required.
	ThingName *string

	noSmithyDocumentSerde
}

// The output from the DescribeThing operation.
type DescribeThingOutput struct {

	// The thing attributes.
	Attributes map[string]string

	// The name of the billing group the thing belongs to.
	BillingGroupName *string

	// The default MQTT client ID. For a typical device, the thing name is also used
	// as the default MQTT client ID. Although we don’t require a mapping between a
	// thing's registry name and its use of MQTT client IDs, certificates, or shadow
	// state, we recommend that you choose a thing name and use it as the MQTT client
	// ID for the registry and the Device Shadow service. This lets you better organize
	// your IoT fleet without removing the flexibility of the underlying device
	// certificate model or shadows.
	DefaultClientId *string

	// The ARN of the thing to describe.
	ThingArn *string

	// The ID of the thing to describe.
	ThingId *string

	// The name of the thing.
	ThingName *string

	// The thing type name.
	ThingTypeName *string

	// The current version of the thing record in the registry. To avoid unintentional
	// changes to the information in the registry, you can pass the version information
	// in the expectedVersion parameter of the UpdateThing and DeleteThing calls.
	Version int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeThing{}, middleware.After)
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
	if err = addOpDescribeThingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeThing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeThing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "DescribeThing",
	}
}
