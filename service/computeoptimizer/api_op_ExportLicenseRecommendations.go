// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Export optimization recommendations for your licenses. Recommendations are
// exported in a comma-separated values (CSV) file, and its metadata in a
// JavaScript Object Notation (JSON) file, to an existing Amazon Simple Storage
// Service (Amazon S3) bucket that you specify. For more information, see
// Exporting Recommendations (https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html)
// in the Compute Optimizer User Guide. You can have only one license export job in
// progress per Amazon Web Services Region.
func (c *Client) ExportLicenseRecommendations(ctx context.Context, params *ExportLicenseRecommendationsInput, optFns ...func(*Options)) (*ExportLicenseRecommendationsOutput, error) {
	if params == nil {
		params = &ExportLicenseRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportLicenseRecommendations", params, optFns, c.addOperationExportLicenseRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportLicenseRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportLicenseRecommendationsInput struct {

	// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name
	// and key prefix for a recommendations export job. You must create the destination
	// Amazon S3 bucket for your recommendations export before you create the export
	// job. Compute Optimizer does not create the S3 bucket for you. After you create
	// the S3 bucket, ensure that it has the required permission policy to allow
	// Compute Optimizer to write the export file to it. If you plan to specify an
	// object prefix when you create the export job, you must include the object prefix
	// in the policy that you add to the S3 bucket. For more information, see Amazon
	// S3 Bucket Policy for Compute Optimizer (https://docs.aws.amazon.com/compute-optimizer/latest/ug/create-s3-bucket-policy-for-compute-optimizer.html)
	// in the Compute Optimizer User Guide.
	//
	// This member is required.
	S3DestinationConfig *types.S3DestinationConfig

	// The IDs of the Amazon Web Services accounts for which to export license
	// recommendations. If your account is the management account of an organization,
	// use this parameter to specify the member account for which you want to export
	// recommendations. This parameter can't be specified together with the include
	// member accounts parameter. The parameters are mutually exclusive. If this
	// parameter is omitted, recommendations for member accounts aren't included in the
	// export. You can specify multiple account IDs per request.
	AccountIds []string

	// The recommendations data to include in the export file. For more information
	// about the fields that can be exported, see Exported files (https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html#exported-files)
	// in the Compute Optimizer User Guide.
	FieldsToExport []types.ExportableLicenseField

	// The format of the export file. A CSV file is the only export format currently
	// supported.
	FileFormat types.FileFormat

	// An array of objects to specify a filter that exports a more specific set of
	// license recommendations.
	Filters []types.LicenseRecommendationFilter

	// Indicates whether to include recommendations for resources in all member
	// accounts of the organization if your account is the management account of an
	// organization. The member accounts must also be opted in to Compute Optimizer,
	// and trusted access for Compute Optimizer must be enabled in the organization
	// account. For more information, see Compute Optimizer and Amazon Web Services
	// Organizations trusted access (https://docs.aws.amazon.com/compute-optimizer/latest/ug/security-iam.html#trusted-service-access)
	// in the Compute Optimizer User Guide. If this parameter is omitted,
	// recommendations for member accounts of the organization aren't included in the
	// export file . This parameter cannot be specified together with the account IDs
	// parameter. The parameters are mutually exclusive.
	IncludeMemberAccounts bool

	noSmithyDocumentSerde
}

type ExportLicenseRecommendationsOutput struct {

	// The identification number of the export job. To view the status of an export
	// job, use the DescribeRecommendationExportJobs action and specify the job ID.
	JobId *string

	// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name
	// and object keys of a recommendations export file, and its associated metadata
	// file.
	S3Destination *types.S3Destination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportLicenseRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpExportLicenseRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpExportLicenseRecommendations{}, middleware.After)
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
	if err = addOpExportLicenseRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportLicenseRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportLicenseRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "ExportLicenseRecommendations",
	}
}
