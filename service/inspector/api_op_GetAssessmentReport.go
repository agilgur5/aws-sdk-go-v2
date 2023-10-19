// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Produces an assessment report that includes detailed and comprehensive results
// of a specified assessment run.
func (c *Client) GetAssessmentReport(ctx context.Context, params *GetAssessmentReportInput, optFns ...func(*Options)) (*GetAssessmentReportOutput, error) {
	if params == nil {
		params = &GetAssessmentReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssessmentReport", params, optFns, c.addOperationGetAssessmentReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssessmentReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssessmentReportInput struct {

	// The ARN that specifies the assessment run for which you want to generate a
	// report.
	//
	// This member is required.
	AssessmentRunArn *string

	// Specifies the file format (html or pdf) of the assessment report that you want
	// to generate.
	//
	// This member is required.
	ReportFileFormat types.ReportFileFormat

	// Specifies the type of the assessment report that you want to generate. There
	// are two types of assessment reports: a finding report and a full report. For
	// more information, see Assessment Reports (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_reports.html)
	// .
	//
	// This member is required.
	ReportType types.ReportType

	noSmithyDocumentSerde
}

type GetAssessmentReportOutput struct {

	// Specifies the status of the request to generate an assessment report.
	//
	// This member is required.
	Status types.ReportStatus

	// Specifies the URL where you can find the generated assessment report. This
	// parameter is only returned if the report is successfully generated.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAssessmentReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAssessmentReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAssessmentReport{}, middleware.After)
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
	if err = addOpGetAssessmentReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssessmentReport(options.Region), middleware.Before); err != nil {
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
		operation: "GetAssessmentReport",
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

func newServiceMetadataMiddleware_opGetAssessmentReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "GetAssessmentReport",
	}
}
