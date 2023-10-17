// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifybackend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a model schema for an existing backend API resource.
func (c *Client) GenerateBackendAPIModels(ctx context.Context, params *GenerateBackendAPIModelsInput, optFns ...func(*Options)) (*GenerateBackendAPIModelsOutput, error) {
	if params == nil {
		params = &GenerateBackendAPIModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateBackendAPIModels", params, optFns, c.addOperationGenerateBackendAPIModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateBackendAPIModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for GenerateBackendAPIModels.
type GenerateBackendAPIModelsInput struct {

	// The app ID.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment.
	//
	// This member is required.
	BackendEnvironmentName *string

	// The name of this resource.
	//
	// This member is required.
	ResourceName *string

	noSmithyDocumentSerde
}

type GenerateBackendAPIModelsOutput struct {

	// The app ID.
	AppId *string

	// The name of the backend environment.
	BackendEnvironmentName *string

	// If the request fails, this error is returned.
	Error *string

	// The ID for the job.
	JobId *string

	// The name of the operation.
	Operation *string

	// The current status of the request.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateBackendAPIModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGenerateBackendAPIModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGenerateBackendAPIModels{}, middleware.After)
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
	if err = addOpGenerateBackendAPIModelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateBackendAPIModels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateBackendAPIModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplifybackend",
		OperationName: "GenerateBackendAPIModels",
	}
}
