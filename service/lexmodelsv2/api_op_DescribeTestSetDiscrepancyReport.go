// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets metadata information about the test set discrepancy report.
func (c *Client) DescribeTestSetDiscrepancyReport(ctx context.Context, params *DescribeTestSetDiscrepancyReportInput, optFns ...func(*Options)) (*DescribeTestSetDiscrepancyReportOutput, error) {
	if params == nil {
		params = &DescribeTestSetDiscrepancyReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTestSetDiscrepancyReport", params, optFns, c.addOperationDescribeTestSetDiscrepancyReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTestSetDiscrepancyReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTestSetDiscrepancyReportInput struct {

	// The unique identifier of the test set discrepancy report.
	//
	// This member is required.
	TestSetDiscrepancyReportId *string

	noSmithyDocumentSerde
}

type DescribeTestSetDiscrepancyReportOutput struct {

	// The time and date of creation for the test set discrepancy report.
	CreationDateTime *time.Time

	// The failure report for the test set discrepancy report generation action.
	FailureReasons []string

	// The date and time of the last update for the test set discrepancy report.
	LastUpdatedDataTime *time.Time

	// The target bot location for the test set discrepancy report.
	Target *types.TestSetDiscrepancyReportResourceTarget

	// Pre-signed Amazon S3 URL to download the test set discrepancy report.
	TestSetDiscrepancyRawOutputUrl *string

	// The unique identifier of the test set discrepancy report to describe.
	TestSetDiscrepancyReportId *string

	// The status for the test set discrepancy report.
	TestSetDiscrepancyReportStatus types.TestSetDiscrepancyReportStatus

	// The top 200 error results from the test set discrepancy report.
	TestSetDiscrepancyTopErrors *types.TestSetDiscrepancyErrors

	// The test set Id for the test set discrepancy report.
	TestSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTestSetDiscrepancyReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTestSetDiscrepancyReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTestSetDiscrepancyReport{}, middleware.After)
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
	if err = addOpDescribeTestSetDiscrepancyReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTestSetDiscrepancyReport(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeTestSetDiscrepancyReport",
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

func newServiceMetadataMiddleware_opDescribeTestSetDiscrepancyReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeTestSetDiscrepancyReport",
	}
}
