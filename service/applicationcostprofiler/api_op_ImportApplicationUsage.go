// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationcostprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Ingests application usage data from Amazon Simple Storage Service (Amazon S3).
// The data must already exist in the S3 location. As part of the action, AWS
// Application Cost Profiler copies the object from your S3 bucket to an S3 bucket
// owned by Amazon for processing asynchronously.
func (c *Client) ImportApplicationUsage(ctx context.Context, params *ImportApplicationUsageInput, optFns ...func(*Options)) (*ImportApplicationUsageOutput, error) {
	if params == nil {
		params = &ImportApplicationUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportApplicationUsage", params, optFns, c.addOperationImportApplicationUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportApplicationUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportApplicationUsageInput struct {

	// Amazon S3 location to import application usage data from.
	//
	// This member is required.
	SourceS3Location *types.SourceS3Location

	noSmithyDocumentSerde
}

type ImportApplicationUsageOutput struct {

	// ID of the import request.
	//
	// This member is required.
	ImportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportApplicationUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportApplicationUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportApplicationUsage{}, middleware.After)
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
	if err = addOpImportApplicationUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportApplicationUsage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportApplicationUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "application-cost-profiler",
		OperationName: "ImportApplicationUsage",
	}
}
