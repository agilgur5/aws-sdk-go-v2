// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a machine learning (ML) project that can contain one or more templates
// that set up an ML pipeline from training to deploying an approved model.
func (c *Client) CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) {
	if params == nil {
		params = &CreateProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProject", params, optFns, c.addOperationCreateProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProjectInput struct {

	// The name of the project.
	//
	// This member is required.
	ProjectName *string

	// The product ID and provisioning artifact ID to provision a service catalog. The
	// provisioning artifact ID will default to the latest provisioning artifact ID of
	// the product, if you don't provide the provisioning artifact ID. For more
	// information, see What is Amazon Web Services Service Catalog (https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html)
	// .
	//
	// This member is required.
	ServiceCatalogProvisioningDetails *types.ServiceCatalogProvisioningDetails

	// A description for the project.
	ProjectDescription *string

	// An array of key-value pairs that you want to use to organize and track your
	// Amazon Web Services resource costs. For more information, see Tagging Amazon
	// Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the Amazon Web Services General Reference Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProjectOutput struct {

	// The Amazon Resource Name (ARN) of the project.
	//
	// This member is required.
	ProjectArn *string

	// The ID of the new project.
	//
	// This member is required.
	ProjectId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProject{}, middleware.After)
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
	if err = addOpCreateProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProject(options.Region), middleware.Before); err != nil {
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
		operation: "CreateProject",
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

func newServiceMetadataMiddleware_opCreateProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateProject",
	}
}
