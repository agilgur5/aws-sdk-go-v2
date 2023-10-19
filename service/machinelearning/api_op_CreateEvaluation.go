// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Evaluation of an MLModel . An MLModel is evaluated on a set of
// observations associated to a DataSource . Like a DataSource for an MLModel , the
// DataSource for an Evaluation contains values for the Target Variable . The
// Evaluation compares the predicted result for each observation to the actual
// outcome and provides a summary so that you know how effective the MLModel
// functions on the test data. Evaluation generates a relevant performance metric,
// such as BinaryAUC, RegressionRMSE or MulticlassAvgFScore based on the
// corresponding MLModelType : BINARY , REGRESSION or MULTICLASS . CreateEvaluation
// is an asynchronous operation. In response to CreateEvaluation , Amazon Machine
// Learning (Amazon ML) immediately returns and sets the evaluation status to
// PENDING . After the Evaluation is created and ready for use, Amazon ML sets the
// status to COMPLETED . You can use the GetEvaluation operation to check progress
// of the evaluation during the creation operation.
func (c *Client) CreateEvaluation(ctx context.Context, params *CreateEvaluationInput, optFns ...func(*Options)) (*CreateEvaluationOutput, error) {
	if params == nil {
		params = &CreateEvaluationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEvaluation", params, optFns, c.addOperationCreateEvaluationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEvaluationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEvaluationInput struct {

	// The ID of the DataSource for the evaluation. The schema of the DataSource must
	// match the schema used to create the MLModel .
	//
	// This member is required.
	EvaluationDataSourceId *string

	// A user-supplied ID that uniquely identifies the Evaluation .
	//
	// This member is required.
	EvaluationId *string

	// The ID of the MLModel to evaluate. The schema used in creating the MLModel must
	// match the schema of the DataSource used in the Evaluation .
	//
	// This member is required.
	MLModelId *string

	// A user-supplied name or description of the Evaluation .
	EvaluationName *string

	noSmithyDocumentSerde
}

// Represents the output of a CreateEvaluation operation, and is an
// acknowledgement that Amazon ML received the request. CreateEvaluation operation
// is asynchronous. You can poll for status updates by using the GetEvcaluation
// operation and checking the Status parameter.
type CreateEvaluationOutput struct {

	// The user-supplied ID that uniquely identifies the Evaluation . This value should
	// be identical to the value of the EvaluationId in the request.
	EvaluationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEvaluationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEvaluation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEvaluation{}, middleware.After)
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
	if err = addOpCreateEvaluationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEvaluation(options.Region), middleware.Before); err != nil {
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
		operation: "CreateEvaluation",
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

func newServiceMetadataMiddleware_opCreateEvaluation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "CreateEvaluation",
	}
}
