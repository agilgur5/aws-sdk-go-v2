// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Task runners call EvaluateExpression to evaluate a string in the context of the
// specified object. For example, a task runner can evaluate SQL queries stored in
// Amazon S3. POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1
// X-Amz-Target: DataPipeline.DescribePipelines Content-Length: 164 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams {"pipelineId": "df-08785951KAKJEXAMPLE", "objectId":
// "Schedule", "expression": "Transform started at #{startDateTime} and finished at
// #{endDateTime}"} x-amzn-RequestId: 02870eb7-0736-11e2-af6f-6bc7a6be60d9
// Content-Type: application/x-amz-json-1.1 Content-Length: 103 Date: Mon, 12 Nov
// 2012 17:50:53 GMT {"evaluatedExpression": "Transform started at
// 2012-12-12T00:00:00 and finished at 2012-12-21T18:00:00"}
func (c *Client) EvaluateExpression(ctx context.Context, params *EvaluateExpressionInput, optFns ...func(*Options)) (*EvaluateExpressionOutput, error) {
	if params == nil {
		params = &EvaluateExpressionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EvaluateExpression", params, optFns, c.addOperationEvaluateExpressionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EvaluateExpressionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for EvaluateExpression.
type EvaluateExpressionInput struct {

	// The expression to evaluate.
	//
	// This member is required.
	Expression *string

	// The ID of the object.
	//
	// This member is required.
	ObjectId *string

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	noSmithyDocumentSerde
}

// Contains the output of EvaluateExpression.
type EvaluateExpressionOutput struct {

	// The evaluated expression.
	//
	// This member is required.
	EvaluatedExpression *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEvaluateExpressionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEvaluateExpression{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEvaluateExpression{}, middleware.After)
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
	if err = addOpEvaluateExpressionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEvaluateExpression(options.Region), middleware.Before); err != nil {
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
		operation: "EvaluateExpression",
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

func newServiceMetadataMiddleware_opEvaluateExpression(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "EvaluateExpression",
	}
}
