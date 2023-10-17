// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a RequestValidator of a given RestApi.
func (c *Client) UpdateRequestValidator(ctx context.Context, params *UpdateRequestValidatorInput, optFns ...func(*Options)) (*UpdateRequestValidatorOutput, error) {
	if params == nil {
		params = &UpdateRequestValidatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRequestValidator", params, optFns, c.addOperationUpdateRequestValidatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRequestValidatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a RequestValidator of a given RestApi.
type UpdateRequestValidatorInput struct {

	// The identifier of RequestValidator to be updated.
	//
	// This member is required.
	RequestValidatorId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// For more information about supported patch operations, see Patch Operations (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html)
	// .
	PatchOperations []types.PatchOperation

	noSmithyDocumentSerde
}

// A set of validation rules for incoming Method requests.
type UpdateRequestValidatorOutput struct {

	// The identifier of this RequestValidator.
	Id *string

	// The name of this RequestValidator
	Name *string

	// A Boolean flag to indicate whether to validate a request body according to the
	// configured Model schema.
	ValidateRequestBody bool

	// A Boolean flag to indicate whether to validate request parameters ( true ) or
	// not ( false ).
	ValidateRequestParameters bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRequestValidatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRequestValidator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRequestValidator{}, middleware.After)
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
	if err = addOpUpdateRequestValidatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRequestValidator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRequestValidator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateRequestValidator",
	}
}
