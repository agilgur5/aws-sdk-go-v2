// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a usage plan key for adding an existing API key to a usage plan.
func (c *Client) CreateUsagePlanKey(ctx context.Context, params *CreateUsagePlanKeyInput, optFns ...func(*Options)) (*CreateUsagePlanKeyOutput, error) {
	if params == nil {
		params = &CreateUsagePlanKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUsagePlanKey", params, optFns, c.addOperationCreateUsagePlanKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUsagePlanKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The POST request to create a usage plan key for adding an existing API key to a
// usage plan.
type CreateUsagePlanKeyInput struct {

	// The identifier of a UsagePlanKey resource for a plan customer.
	//
	// This member is required.
	KeyId *string

	// The type of a UsagePlanKey resource for a plan customer.
	//
	// This member is required.
	KeyType *string

	// The Id of the UsagePlan resource representing the usage plan containing the
	// to-be-created UsagePlanKey resource representing a plan customer.
	//
	// This member is required.
	UsagePlanId *string

	noSmithyDocumentSerde
}

// Represents a usage plan key to identify a plan customer.
type CreateUsagePlanKeyOutput struct {

	// The Id of a usage plan key.
	Id *string

	// The name of a usage plan key.
	Name *string

	// The type of a usage plan key. Currently, the valid key type is API_KEY .
	Type *string

	// The value of a usage plan key.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUsagePlanKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateUsagePlanKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUsagePlanKey{}, middleware.After)
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
	if err = addOpCreateUsagePlanKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUsagePlanKey(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opCreateUsagePlanKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateUsagePlanKey",
	}
}
