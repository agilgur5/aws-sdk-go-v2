// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubrefactorspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets an Amazon Web Services Migration Hub Refactor Spaces application.
func (c *Client) GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) {
	if params == nil {
		params = &GetApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplication", params, optFns, c.addOperationGetApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationInput struct {

	// The ID of the application.
	//
	// This member is required.
	ApplicationIdentifier *string

	// The ID of the environment.
	//
	// This member is required.
	EnvironmentIdentifier *string

	noSmithyDocumentSerde
}

type GetApplicationOutput struct {

	// The endpoint URL of the API Gateway proxy.
	ApiGatewayProxy *types.ApiGatewayProxyConfig

	// The unique identifier of the application.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the application.
	Arn *string

	// The Amazon Web Services account ID of the application creator.
	CreatedByAccountId *string

	// A timestamp that indicates when the application is created.
	CreatedTime *time.Time

	// The unique identifier of the environment.
	EnvironmentId *string

	// Any error associated with the application resource.
	Error *types.ErrorResponse

	// A timestamp that indicates when the application was last updated.
	LastUpdatedTime *time.Time

	// The name of the application.
	Name *string

	// The Amazon Web Services account ID of the application owner (which is always
	// the same as the environment owner account ID).
	OwnerAccountId *string

	// The proxy type of the proxy created within the application.
	ProxyType types.ProxyType

	// The current state of the application.
	State types.ApplicationState

	// The tags assigned to the application. A tag is a label that you assign to an
	// Amazon Web Services resource. Each tag consists of a key-value pair.
	Tags map[string]string

	// The ID of the virtual private cloud (VPC).
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplication{}, middleware.After)
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
	if err = addOpGetApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplication(options.Region), middleware.Before); err != nil {
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
		operation: "GetApplication",
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

func newServiceMetadataMiddleware_opGetApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "refactor-spaces",
		OperationName: "GetApplication",
	}
}
