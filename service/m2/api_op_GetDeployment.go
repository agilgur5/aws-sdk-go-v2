// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets details of a specific deployment with a given deployment identifier.
func (c *Client) GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) {
	if params == nil {
		params = &GetDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeployment", params, optFns, c.addOperationGetDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeploymentInput struct {

	// The unique identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the deployment.
	//
	// This member is required.
	DeploymentId *string

	noSmithyDocumentSerde
}

type GetDeploymentOutput struct {

	// The unique identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The application version.
	//
	// This member is required.
	ApplicationVersion *int32

	// The timestamp when the deployment was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The unique identifier of the deployment.
	//
	// This member is required.
	DeploymentId *string

	// The unique identifier of the runtime environment.
	//
	// This member is required.
	EnvironmentId *string

	// The status of the deployment.
	//
	// This member is required.
	Status types.DeploymentLifecycle

	// The reason for the reported status.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeployment{}, middleware.After)
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
	if err = addOpGetDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeployment(options.Region), middleware.Before); err != nil {
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
		operation: "GetDeployment",
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

func newServiceMetadataMiddleware_opGetDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "GetDeployment",
	}
}
