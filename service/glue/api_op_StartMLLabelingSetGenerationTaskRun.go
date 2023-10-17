// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the active learning workflow for your machine learning transform to
// improve the transform's quality by generating label sets and adding labels. When
// the StartMLLabelingSetGenerationTaskRun finishes, Glue will have generated a
// "labeling set" or a set of questions for humans to answer. In the case of the
// FindMatches transform, these questions are of the form, “What is the correct way
// to group these rows together into groups composed entirely of matching records?”
// After the labeling process is finished, you can upload your labels with a call
// to StartImportLabelsTaskRun . After StartImportLabelsTaskRun finishes, all
// future runs of the machine learning transform will use the new and improved
// labels and perform a higher-quality transformation.
func (c *Client) StartMLLabelingSetGenerationTaskRun(ctx context.Context, params *StartMLLabelingSetGenerationTaskRunInput, optFns ...func(*Options)) (*StartMLLabelingSetGenerationTaskRunOutput, error) {
	if params == nil {
		params = &StartMLLabelingSetGenerationTaskRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMLLabelingSetGenerationTaskRun", params, optFns, c.addOperationStartMLLabelingSetGenerationTaskRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMLLabelingSetGenerationTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMLLabelingSetGenerationTaskRunInput struct {

	// The Amazon Simple Storage Service (Amazon S3) path where you generate the
	// labeling set.
	//
	// This member is required.
	OutputS3Path *string

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string

	noSmithyDocumentSerde
}

type StartMLLabelingSetGenerationTaskRunOutput struct {

	// The unique run identifier that is associated with this task run.
	TaskRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMLLabelingSetGenerationTaskRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMLLabelingSetGenerationTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMLLabelingSetGenerationTaskRun{}, middleware.After)
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
	if err = addOpStartMLLabelingSetGenerationTaskRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMLLabelingSetGenerationTaskRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMLLabelingSetGenerationTaskRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartMLLabelingSetGenerationTaskRun",
	}
}
