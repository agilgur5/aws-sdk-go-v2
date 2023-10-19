// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a notebook execution.
func (c *Client) StartNotebookExecution(ctx context.Context, params *StartNotebookExecutionInput, optFns ...func(*Options)) (*StartNotebookExecutionOutput, error) {
	if params == nil {
		params = &StartNotebookExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartNotebookExecution", params, optFns, c.addOperationStartNotebookExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartNotebookExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartNotebookExecutionInput struct {

	// Specifies the execution engine (cluster) that runs the notebook execution.
	//
	// This member is required.
	ExecutionEngine *types.ExecutionEngineConfig

	// The name or ARN of the IAM role that is used as the service role for Amazon EMR
	// (the Amazon EMR role) for the notebook execution.
	//
	// This member is required.
	ServiceRole *string

	// The unique identifier of the Amazon EMR Notebook to use for notebook execution.
	EditorId *string

	// The environment variables associated with the notebook execution.
	EnvironmentVariables map[string]string

	// An optional name for the notebook execution.
	NotebookExecutionName *string

	// The unique identifier of the Amazon EC2 security group to associate with the
	// Amazon EMR Notebook for this notebook execution.
	NotebookInstanceSecurityGroupId *string

	// Input parameters in JSON format passed to the Amazon EMR Notebook at runtime
	// for execution.
	NotebookParams *string

	// The Amazon S3 location for the notebook execution input.
	NotebookS3Location *types.NotebookS3LocationFromInput

	// The output format for the notebook execution.
	OutputNotebookFormat types.OutputNotebookFormat

	// The Amazon S3 location for the notebook execution output.
	OutputNotebookS3Location *types.OutputNotebookS3LocationFromInput

	// The path and file name of the notebook file for this execution, relative to the
	// path specified for the Amazon EMR Notebook. For example, if you specify a path
	// of s3://MyBucket/MyNotebooks when you create an Amazon EMR Notebook for a
	// notebook with an ID of e-ABCDEFGHIJK1234567890ABCD (the EditorID of this
	// request), and you specify a RelativePath of
	// my_notebook_executions/notebook_execution.ipynb , the location of the file for
	// the notebook execution is
	// s3://MyBucket/MyNotebooks/e-ABCDEFGHIJK1234567890ABCD/my_notebook_executions/notebook_execution.ipynb
	// .
	RelativePath *string

	// A list of tags associated with a notebook execution. Tags are user-defined
	// key-value pairs that consist of a required key string with a maximum of 128
	// characters and an optional value string with a maximum of 256 characters.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartNotebookExecutionOutput struct {

	// The unique identifier of the notebook execution.
	NotebookExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartNotebookExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartNotebookExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartNotebookExecution{}, middleware.After)
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
	if err = addOpStartNotebookExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartNotebookExecution(options.Region), middleware.Before); err != nil {
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
		operation: "StartNotebookExecution",
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

func newServiceMetadataMiddleware_opStartNotebookExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "StartNotebookExecution",
	}
}
