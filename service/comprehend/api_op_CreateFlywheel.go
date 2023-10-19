// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A flywheel is an Amazon Web Services resource that orchestrates the ongoing
// training of a model for custom classification or custom entity recognition. You
// can create a flywheel to start with an existing trained model, or Comprehend can
// create and train a new model. When you create the flywheel, Comprehend creates a
// data lake in your account. The data lake holds the training data and test data
// for all versions of the model. To use a flywheel with an existing trained model,
// you specify the active model version. Comprehend copies the model's training
// data and test data into the flywheel's data lake. To use the flywheel with a new
// model, you need to provide a dataset for training data (and optional test data)
// when you create the flywheel. For more information about flywheels, see
// Flywheel overview (https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html)
// in the Amazon Comprehend Developer Guide.
func (c *Client) CreateFlywheel(ctx context.Context, params *CreateFlywheelInput, optFns ...func(*Options)) (*CreateFlywheelOutput, error) {
	if params == nil {
		params = &CreateFlywheelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlywheel", params, optFns, c.addOperationCreateFlywheelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlywheelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlywheelInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend
	// the permissions required to access the flywheel data in the data lake.
	//
	// This member is required.
	DataAccessRoleArn *string

	// Enter the S3 location for the data lake. You can specify a new S3 bucket or a
	// new folder of an existing S3 bucket. The flywheel creates the data lake at this
	// location.
	//
	// This member is required.
	DataLakeS3Uri *string

	// Name for the flywheel.
	//
	// This member is required.
	FlywheelName *string

	// To associate an existing model with the flywheel, specify the Amazon Resource
	// Number (ARN) of the model version.
	ActiveModelArn *string

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string

	// Data security configurations.
	DataSecurityConfig *types.DataSecurityConfig

	// The model type.
	ModelType types.ModelType

	// The tags to associate with this flywheel.
	Tags []types.Tag

	// Configuration about the custom classifier associated with the flywheel.
	TaskConfig *types.TaskConfig

	noSmithyDocumentSerde
}

type CreateFlywheelOutput struct {

	// The Amazon Resource Number (ARN) of the active model version.
	ActiveModelArn *string

	// The Amazon Resource Number (ARN) of the flywheel.
	FlywheelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFlywheelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFlywheel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFlywheel{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateFlywheelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFlywheelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlywheel(options.Region), middleware.Before); err != nil {
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
		operation: "CreateFlywheel",
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

type idempotencyToken_initializeOpCreateFlywheel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFlywheel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFlywheel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFlywheelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFlywheelInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFlywheelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFlywheel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFlywheel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "CreateFlywheel",
	}
}
