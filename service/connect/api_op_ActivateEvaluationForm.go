// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates an evaluation form in the specified Amazon Connect instance. After
// the evaluation form is activated, it is available to start new evaluations based
// on the form.
func (c *Client) ActivateEvaluationForm(ctx context.Context, params *ActivateEvaluationFormInput, optFns ...func(*Options)) (*ActivateEvaluationFormOutput, error) {
	if params == nil {
		params = &ActivateEvaluationFormInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ActivateEvaluationForm", params, optFns, c.addOperationActivateEvaluationFormMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ActivateEvaluationFormOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ActivateEvaluationFormInput struct {

	// The unique identifier for the evaluation form.
	//
	// This member is required.
	EvaluationFormId *string

	// The version of the evaluation form to activate. If the version property is not
	// provided, the latest version of the evaluation form is activated.
	//
	// This member is required.
	EvaluationFormVersion *int32

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	noSmithyDocumentSerde
}

type ActivateEvaluationFormOutput struct {

	// The Amazon Resource Name (ARN) for the evaluation form resource.
	//
	// This member is required.
	EvaluationFormArn *string

	// The unique identifier for the evaluation form.
	//
	// This member is required.
	EvaluationFormId *string

	// A version of the evaluation form.
	//
	// This member is required.
	EvaluationFormVersion *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationActivateEvaluationFormMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpActivateEvaluationForm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpActivateEvaluationForm{}, middleware.After)
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
	if err = addOpActivateEvaluationFormValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opActivateEvaluationForm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opActivateEvaluationForm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ActivateEvaluationForm",
	}
}
