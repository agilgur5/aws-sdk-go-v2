// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a MethodResponse to an existing Method resource.
func (c *Client) PutMethodResponse(ctx context.Context, params *PutMethodResponseInput, optFns ...func(*Options)) (*PutMethodResponseOutput, error) {
	if params == nil {
		params = &PutMethodResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMethodResponse", params, optFns, c.addOperationPutMethodResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMethodResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to add a MethodResponse to an existing Method resource.
type PutMethodResponseInput struct {

	// The HTTP verb of the Method resource.
	//
	// This member is required.
	HttpMethod *string

	// The Resource identifier for the Method resource.
	//
	// This member is required.
	ResourceId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The method response's status code.
	//
	// This member is required.
	StatusCode *string

	// Specifies the Model resources used for the response's content type. Response
	// models are represented as a key/value map, with a content type as the key and a
	// Model name as the value.
	ResponseModels map[string]string

	// A key-value map specifying required or optional response parameters that API
	// Gateway can send back to the caller. A key defines a method response header name
	// and the associated value is a Boolean flag indicating whether the method
	// response parameter is required or not. The method response header names must
	// match the pattern of method.response.header.{name} , where name is a valid and
	// unique header name. The response parameter names defined here are available in
	// the integration response to be mapped from an integration response header
	// expressed in integration.response.header.{name} , a static value enclosed within
	// a pair of single quotes (e.g., 'application/json' ), or a JSON expression from
	// the back-end response payload in the form of
	// integration.response.body.{JSON-expression} , where JSON-expression is a valid
	// JSON expression without the $ prefix.)
	ResponseParameters map[string]bool

	noSmithyDocumentSerde
}

// Represents a method response of a given HTTP status code returned to the
// client. The method response is passed from the back end through the associated
// integration response that can be transformed using a mapping template.
type PutMethodResponseOutput struct {

	// Specifies the Model resources used for the response's content-type. Response
	// models are represented as a key/value map, with a content-type as the key and a
	// Model name as the value.
	ResponseModels map[string]string

	// A key-value map specifying required or optional response parameters that API
	// Gateway can send back to the caller. A key defines a method response header and
	// the value specifies whether the associated method response header is required or
	// not. The expression of the key must match the pattern
	// method.response.header.{name} , where name is a valid and unique header name.
	// API Gateway passes certain integration response data to the method response
	// headers specified here according to the mapping you prescribe in the API's
	// IntegrationResponse. The integration response data that can be mapped include an
	// integration response header expressed in integration.response.header.{name} , a
	// static value enclosed within a pair of single quotes (e.g., 'application/json'
	// ), or a JSON expression from the back-end response payload in the form of
	// integration.response.body.{JSON-expression} , where JSON-expression is a valid
	// JSON expression without the $ prefix.)
	ResponseParameters map[string]bool

	// The method response's status code.
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMethodResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutMethodResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutMethodResponse{}, middleware.After)
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
	if err = addOpPutMethodResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMethodResponse(options.Region), middleware.Before); err != nil {
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
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "PutMethodResponse",
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

func newServiceMetadataMiddleware_opPutMethodResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "PutMethodResponse",
	}
}
