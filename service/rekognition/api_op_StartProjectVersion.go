// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation applies only to Amazon Rekognition Custom Labels. Starts the
// running of the version of a model. Starting a model takes a while to complete.
// To check the current state of the model, use DescribeProjectVersions . Once the
// model is running, you can detect custom labels in new images by calling
// DetectCustomLabels . You are charged for the amount of time that the model is
// running. To stop a running model, call StopProjectVersion . This operation
// requires permissions to perform the rekognition:StartProjectVersion action.
func (c *Client) StartProjectVersion(ctx context.Context, params *StartProjectVersionInput, optFns ...func(*Options)) (*StartProjectVersionOutput, error) {
	if params == nil {
		params = &StartProjectVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartProjectVersion", params, optFns, c.addOperationStartProjectVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartProjectVersionInput struct {

	// The minimum number of inference units to use. A single inference unit
	// represents 1 hour of processing. Use a higher number to increase the TPS
	// throughput of your model. You are charged for the number of inference units that
	// you use.
	//
	// This member is required.
	MinInferenceUnits *int32

	// The Amazon Resource Name(ARN) of the model version that you want to start.
	//
	// This member is required.
	ProjectVersionArn *string

	// The maximum number of inference units to use for auto-scaling the model. If you
	// don't specify a value, Amazon Rekognition Custom Labels doesn't auto-scale the
	// model.
	MaxInferenceUnits *int32

	noSmithyDocumentSerde
}

type StartProjectVersionOutput struct {

	// The current running status of the model.
	Status types.ProjectVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartProjectVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartProjectVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartProjectVersion{}, middleware.After)
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
	if err = addOpStartProjectVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartProjectVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartProjectVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartProjectVersion",
	}
}
