// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifybackend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplifybackend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the details for a backend API.
func (c *Client) GetBackendAPI(ctx context.Context, params *GetBackendAPIInput, optFns ...func(*Options)) (*GetBackendAPIOutput, error) {
	if params == nil {
		params = &GetBackendAPIInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBackendAPI", params, optFns, c.addOperationGetBackendAPIMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBackendAPIOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for GetBackendAPI.
type GetBackendAPIInput struct {

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

	// Defines the resource configuration for the data model in your Amplify project.
	ResourceConfig *types.BackendAPIResourceConfig

	noSmithyDocumentSerde
}

type GetBackendAPIOutput struct {

	// The app ID.
	AppId *string

	// The name of the backend environment.
	BackendEnvironmentName *string

	// If the request fails, this error is returned.
	Error *string

	// The resource configuration for this response object.
	ResourceConfig *types.BackendAPIResourceConfig

	// The name of this resource.
	ResourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBackendAPIMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBackendAPI{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBackendAPI{}, middleware.After)
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
	if err = addOpGetBackendAPIValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBackendAPI(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBackendAPI(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplifybackend",
		OperationName: "GetBackendAPI",
	}
}
