// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes an existing model defined for a RestApi resource.
func (c *Client) GetModel(ctx context.Context, params *GetModelInput, optFns ...func(*Options)) (*GetModelOutput, error) {
	if params == nil {
		params = &GetModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModel", params, optFns, c.addOperationGetModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list information about a model in an existing RestApi resource.
type GetModelInput struct {

	// The name of the model as an identifier.
	//
	// This member is required.
	ModelName *string

	// The RestApi identifier under which the Model exists.
	//
	// This member is required.
	RestApiId *string

	// A query parameter of a Boolean value to resolve ( true ) all external model
	// references and returns a flattened model schema or not ( false ) The default is
	// false .
	Flatten bool

	noSmithyDocumentSerde
}

// Represents the data structure of a method's request or response payload.
type GetModelOutput struct {

	// The content-type for the model.
	ContentType *string

	// The description of the model.
	Description *string

	// The identifier for the model resource.
	Id *string

	// The name of the model. Must be an alphanumeric string.
	Name *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 model. Do not include "\*/" characters in the description of any
	// properties because such "\*/" characters may be interpreted as the closing
	// marker for comments in some languages, such as Java or JavaScript, causing the
	// installation of your API's SDK generated by API Gateway to fail.
	Schema *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetModel{}, middleware.After)
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
	if err = addOpGetModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetModel",
	}
}
