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

// Starts a SageMaker Edge Manager model packaging job. Edge Manager will use the
// model artifacts from the Amazon Simple Storage Service bucket that you specify.
// After the model has been packaged, Amazon SageMaker saves the resulting
// artifacts to an S3 bucket that you specify.
func (c *Client) CreateEdgePackagingJob(ctx context.Context, params *CreateEdgePackagingJobInput, optFns ...func(*Options)) (*CreateEdgePackagingJobOutput, error) {
	if params == nil {
		params = &CreateEdgePackagingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEdgePackagingJob", params, optFns, c.addOperationCreateEdgePackagingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEdgePackagingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEdgePackagingJobInput struct {

	// The name of the SageMaker Neo compilation job that will be used to locate model
	// artifacts for packaging.
	//
	// This member is required.
	CompilationJobName *string

	// The name of the edge packaging job.
	//
	// This member is required.
	EdgePackagingJobName *string

	// The name of the model.
	//
	// This member is required.
	ModelName *string

	// The version of the model.
	//
	// This member is required.
	ModelVersion *string

	// Provides information about the output location for the packaged model.
	//
	// This member is required.
	OutputConfig *types.EdgeOutputConfig

	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to
	// download and upload the model, and to contact SageMaker Neo.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Web Services KMS key to use when encrypting the EBS volume the edge
	// packaging job runs on.
	ResourceKey *string

	// Creates tags for the packaging job.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEdgePackagingJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEdgePackagingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEdgePackagingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEdgePackagingJob{}, middleware.After)
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
	if err = addOpCreateEdgePackagingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEdgePackagingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEdgePackagingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateEdgePackagingJob",
	}
}
