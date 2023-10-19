// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the what-if analysis created using the CreateWhatIfAnalysis
// operation. In addition to listing the properties provided in the
// CreateWhatIfAnalysis request, this operation lists the following properties:
//   - CreationTime
//   - LastModificationTime
//   - Message - If an error occurred, information about the error.
//   - Status
func (c *Client) DescribeWhatIfAnalysis(ctx context.Context, params *DescribeWhatIfAnalysisInput, optFns ...func(*Options)) (*DescribeWhatIfAnalysisOutput, error) {
	if params == nil {
		params = &DescribeWhatIfAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWhatIfAnalysis", params, optFns, c.addOperationDescribeWhatIfAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWhatIfAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWhatIfAnalysisInput struct {

	// The Amazon Resource Name (ARN) of the what-if analysis that you are interested
	// in.
	//
	// This member is required.
	WhatIfAnalysisArn *string

	noSmithyDocumentSerde
}

type DescribeWhatIfAnalysisOutput struct {

	// When the what-if analysis was created.
	CreationTime *time.Time

	// The approximate time remaining to complete the what-if analysis, in minutes.
	EstimatedTimeRemainingInMinutes *int64

	// The Amazon Resource Name (ARN) of the what-if forecast.
	ForecastArn *string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//   - CREATE_PENDING - The CreationTime .
	//   - CREATE_IN_PROGRESS - The current timestamp.
	//   - CREATE_STOPPING - The current timestamp.
	//   - CREATE_STOPPED - When the job stopped.
	//   - ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The status of the what-if analysis. States include:
	//   - ACTIVE
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//   - CREATE_STOPPING , CREATE_STOPPED
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	// The Status of the what-if analysis must be ACTIVE before you can access the
	// analysis.
	Status *string

	// Defines the set of time series that are used to create the forecasts in a
	// TimeSeriesIdentifiers object. The TimeSeriesIdentifiers object needs the
	// following information:
	//   - DataSource
	//   - Format
	//   - Schema
	TimeSeriesSelector *types.TimeSeriesSelector

	// The Amazon Resource Name (ARN) of the what-if analysis.
	WhatIfAnalysisArn *string

	// The name of the what-if analysis.
	WhatIfAnalysisName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWhatIfAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWhatIfAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWhatIfAnalysis{}, middleware.After)
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
	if err = addOpDescribeWhatIfAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWhatIfAnalysis(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeWhatIfAnalysis",
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

func newServiceMetadataMiddleware_opDescribeWhatIfAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeWhatIfAnalysis",
	}
}
