// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a Function object.
func (c *Client) UpdateFunction(ctx context.Context, params *UpdateFunctionInput, optFns ...func(*Options)) (*UpdateFunctionOutput, error) {
	if params == nil {
		params = &UpdateFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFunction", params, optFns, c.addOperationUpdateFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFunctionInput struct {

	// The GraphQL API ID.
	//
	// This member is required.
	ApiId *string

	// The Function DataSource name.
	//
	// This member is required.
	DataSourceName *string

	// The function ID.
	//
	// This member is required.
	FunctionId *string

	// The Function name.
	//
	// This member is required.
	Name *string

	// The function code that contains the request and response functions. When code
	// is used, the runtime is required. The runtime value must be APPSYNC_JS .
	Code *string

	// The Function description.
	Description *string

	// The version of the request mapping template. Currently, the supported value is
	// 2018-05-29. Note that when using VTL and mapping templates, the functionVersion
	// is required.
	FunctionVersion *string

	// The maximum batching size for a resolver.
	MaxBatchSize int32

	// The Function request mapping template. Functions support only the 2018-05-29
	// version of the request mapping template.
	RequestMappingTemplate *string

	// The Function request mapping template.
	ResponseMappingTemplate *string

	// Describes a runtime used by an Amazon Web Services AppSync pipeline resolver or
	// Amazon Web Services AppSync function. Specifies the name and version of the
	// runtime to use. Note that if a runtime is specified, code must also be
	// specified.
	Runtime *types.AppSyncRuntime

	// Describes a Sync configuration for a resolver. Specifies which Conflict
	// Detection strategy and Resolution strategy to use when the resolver is invoked.
	SyncConfig *types.SyncConfig

	noSmithyDocumentSerde
}

type UpdateFunctionOutput struct {

	// The Function object.
	FunctionConfiguration *types.FunctionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFunction{}, middleware.After)
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
	if err = addOpUpdateFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFunction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "UpdateFunction",
	}
}
