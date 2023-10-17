// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used by an Lambda function to deliver evaluation results to Config. This action
// is required in every Lambda function that is invoked by an Config rule.
func (c *Client) PutEvaluations(ctx context.Context, params *PutEvaluationsInput, optFns ...func(*Options)) (*PutEvaluationsOutput, error) {
	if params == nil {
		params = &PutEvaluationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEvaluations", params, optFns, c.addOperationPutEvaluationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEvaluationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEvaluationsInput struct {

	// An encrypted token that associates an evaluation with an Config rule.
	// Identifies the rule and the event that triggered the evaluation.
	//
	// This member is required.
	ResultToken *string

	// The assessments that the Lambda function performs. Each evaluation identifies
	// an Amazon Web Services resource and indicates whether it complies with the
	// Config rule that invokes the Lambda function.
	Evaluations []types.Evaluation

	// Use this parameter to specify a test run for PutEvaluations . You can verify
	// whether your Lambda function will deliver evaluation results to Config. No
	// updates occur to your existing evaluations, and evaluation results are not sent
	// to Config. When TestMode is true , PutEvaluations doesn't require a valid value
	// for the ResultToken parameter, but the value cannot be null.
	TestMode bool

	noSmithyDocumentSerde
}

type PutEvaluationsOutput struct {

	// Requests that failed because of a client or server error.
	FailedEvaluations []types.Evaluation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEvaluationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutEvaluations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEvaluations{}, middleware.After)
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
	if err = addOpPutEvaluationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEvaluations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEvaluations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutEvaluations",
	}
}
