// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the results for an Amazon Textract asynchronous operation that analyzes
// text in a lending document. You start asynchronous text analysis by calling
// StartLendingAnalysis , which returns a job identifier ( JobId ). When the text
// analysis operation finishes, Amazon Textract publishes a completion status to
// the Amazon Simple Notification Service (Amazon SNS) topic that's registered in
// the initial call to StartLendingAnalysis . To get the results of the text
// analysis operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED. If so, call GetLendingAnalysis, and pass the job
// identifier ( JobId ) from the initial call to StartLendingAnalysis .
func (c *Client) GetLendingAnalysis(ctx context.Context, params *GetLendingAnalysisInput, optFns ...func(*Options)) (*GetLendingAnalysisOutput, error) {
	if params == nil {
		params = &GetLendingAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLendingAnalysis", params, optFns, c.addOperationGetLendingAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLendingAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLendingAnalysisInput struct {

	// A unique identifier for the lending or text-detection job. The JobId is
	// returned from StartLendingAnalysis . A JobId value is only valid for 7 days.
	//
	// This member is required.
	JobId *string

	// The maximum number of results to return per paginated call. The largest value
	// that you can specify is 30. If you specify a value greater than 30, a maximum of
	// 30 results is returned. The default value is 30.
	MaxResults *int32

	// If the previous response was incomplete, Amazon Textract returns a pagination
	// token in the response. You can use this pagination token to retrieve the next
	// set of lending results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetLendingAnalysisOutput struct {

	// The current model version of the Analyze Lending API.
	AnalyzeLendingModelVersion *string

	// Information about the input document.
	DocumentMetadata *types.DocumentMetadata

	// The current status of the lending analysis job.
	JobStatus types.JobStatus

	// If the response is truncated, Amazon Textract returns this token. You can use
	// this token in the subsequent request to retrieve the next set of lending
	// results.
	NextToken *string

	// Holds the information returned by one of AmazonTextract's document analysis
	// operations for the pinstripe.
	Results []types.LendingResult

	// Returns if the lending analysis job could not be completed. Contains
	// explanation for what error occurred.
	StatusMessage *string

	// A list of warnings that occurred during the lending analysis operation.
	Warnings []types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLendingAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLendingAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLendingAnalysis{}, middleware.After)
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
	if err = addOpGetLendingAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLendingAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLendingAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "GetLendingAnalysis",
	}
}
