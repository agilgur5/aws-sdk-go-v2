// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Restarts selected nodes of a previous partially completed workflow run and
// resumes the workflow run. The selected nodes and all nodes that are downstream
// from the selected nodes are run.
func (c *Client) ResumeWorkflowRun(ctx context.Context, params *ResumeWorkflowRunInput, optFns ...func(*Options)) (*ResumeWorkflowRunOutput, error) {
	if params == nil {
		params = &ResumeWorkflowRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResumeWorkflowRun", params, optFns, c.addOperationResumeWorkflowRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResumeWorkflowRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeWorkflowRunInput struct {

	// The name of the workflow to resume.
	//
	// This member is required.
	Name *string

	// A list of the node IDs for the nodes you want to restart. The nodes that are to
	// be restarted must have a run attempt in the original run.
	//
	// This member is required.
	NodeIds []string

	// The ID of the workflow run to resume.
	//
	// This member is required.
	RunId *string

	noSmithyDocumentSerde
}

type ResumeWorkflowRunOutput struct {

	// A list of the node IDs for the nodes that were actually restarted.
	NodeIds []string

	// The new ID assigned to the resumed workflow run. Each resume of a workflow run
	// will have a new run ID.
	RunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResumeWorkflowRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResumeWorkflowRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResumeWorkflowRun{}, middleware.After)
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
	if err = addOpResumeWorkflowRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResumeWorkflowRun(options.Region), middleware.Before); err != nil {
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
		operation: "ResumeWorkflowRun",
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

func newServiceMetadataMiddleware_opResumeWorkflowRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ResumeWorkflowRun",
	}
}
