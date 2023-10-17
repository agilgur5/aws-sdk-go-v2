// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables you to provide additional labels (examples of truth) to be used to
// teach the machine learning transform and improve its quality. This API operation
// is generally used as part of the active learning workflow that starts with the
// StartMLLabelingSetGenerationTaskRun call and that ultimately results in
// improving the quality of your machine learning transform. After the
// StartMLLabelingSetGenerationTaskRun finishes, Glue machine learning will have
// generated a series of questions for humans to answer. (Answering these questions
// is often called 'labeling' in the machine learning workflows). In the case of
// the FindMatches transform, these questions are of the form, “What is the
// correct way to group these rows together into groups composed entirely of
// matching records?” After the labeling process is finished, users upload their
// answers/labels with a call to StartImportLabelsTaskRun . After
// StartImportLabelsTaskRun finishes, all future runs of the machine learning
// transform use the new and improved labels and perform a higher-quality
// transformation. By default, StartMLLabelingSetGenerationTaskRun continually
// learns from and combines all labels that you upload unless you set Replace to
// true. If you set Replace to true, StartImportLabelsTaskRun deletes and forgets
// all previously uploaded labels and learns only from the exact set that you
// upload. Replacing labels can be helpful if you realize that you previously
// uploaded incorrect labels, and you believe that they are having a negative
// effect on your transform quality. You can check on the status of your task run
// by calling the GetMLTaskRun operation.
func (c *Client) StartImportLabelsTaskRun(ctx context.Context, params *StartImportLabelsTaskRunInput, optFns ...func(*Options)) (*StartImportLabelsTaskRunOutput, error) {
	if params == nil {
		params = &StartImportLabelsTaskRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImportLabelsTaskRun", params, optFns, c.addOperationStartImportLabelsTaskRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImportLabelsTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImportLabelsTaskRunInput struct {

	// The Amazon Simple Storage Service (Amazon S3) path from where you import the
	// labels.
	//
	// This member is required.
	InputS3Path *string

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string

	// Indicates whether to overwrite your existing labels.
	ReplaceAllLabels bool

	noSmithyDocumentSerde
}

type StartImportLabelsTaskRunOutput struct {

	// The unique identifier for the task run.
	TaskRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImportLabelsTaskRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartImportLabelsTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartImportLabelsTaskRun{}, middleware.After)
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
	if err = addOpStartImportLabelsTaskRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImportLabelsTaskRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartImportLabelsTaskRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartImportLabelsTaskRun",
	}
}
