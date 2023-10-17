// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the status, metrics, and errors (if there are any) that are associated
// with a job.
func (c *Client) GetMatchingJob(ctx context.Context, params *GetMatchingJobInput, optFns ...func(*Options)) (*GetMatchingJobOutput, error) {
	if params == nil {
		params = &GetMatchingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMatchingJob", params, optFns, c.addOperationGetMatchingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMatchingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMatchingJobInput struct {

	// The ID of the job.
	//
	// This member is required.
	JobId *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	noSmithyDocumentSerde
}

type GetMatchingJobOutput struct {

	// The ID of the job.
	//
	// This member is required.
	JobId *string

	// The time at which the job was started.
	//
	// This member is required.
	StartTime *time.Time

	// The current status of the job.
	//
	// This member is required.
	Status types.JobStatus

	// The time at which the job has finished.
	EndTime *time.Time

	// An object containing an error message, if there was an error.
	ErrorDetails *types.ErrorDetails

	// Metrics associated with the execution, specifically total records processed,
	// unique IDs generated, and records the execution skipped.
	Metrics *types.JobMetrics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMatchingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMatchingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMatchingJob{}, middleware.After)
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
	if err = addOpGetMatchingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMatchingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMatchingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "entityresolution",
		OperationName: "GetMatchingJob",
	}
}
