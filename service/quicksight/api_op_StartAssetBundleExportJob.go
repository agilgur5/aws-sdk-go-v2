// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an Asset Bundle export job. An Asset Bundle export job exports specified
// Amazon QuickSight assets. You can also choose to export any asset dependencies
// in the same job. Export jobs run asynchronously and can be polled with a
// DescribeAssetBundleExportJob API call. When a job is successfully completed, a
// download URL that contains the exported assets is returned. The URL is valid for
// 5 minutes and can be refreshed with a DescribeAssetBundleExportJob API call.
// Each Amazon QuickSight account can run up to 5 export jobs concurrently. The API
// caller must have the necessary permissions in their IAM role to access each
// resource before the resources can be exported.
func (c *Client) StartAssetBundleExportJob(ctx context.Context, params *StartAssetBundleExportJobInput, optFns ...func(*Options)) (*StartAssetBundleExportJobOutput, error) {
	if params == nil {
		params = &StartAssetBundleExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAssetBundleExportJob", params, optFns, c.addOperationStartAssetBundleExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAssetBundleExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAssetBundleExportJobInput struct {

	// The ID of the job. This ID is unique while the job is running. After the job is
	// completed, you can reuse this ID for another job.
	//
	// This member is required.
	AssetBundleExportJobId *string

	// The ID of the Amazon Web Services account to export assets from.
	//
	// This member is required.
	AwsAccountId *string

	// The export data format.
	//
	// This member is required.
	ExportFormat types.AssetBundleExportFormat

	// An array of resource ARNs to export. The following resources are supported.
	//   - Analysis
	//   - Dashboard
	//   - DataSet
	//   - DataSource
	//   - RefreshSchedule
	//   - Theme
	//   - VPCConnection
	// The API caller must have the necessary permissions in their IAM role to access
	// each resource before the resources can be exported.
	//
	// This member is required.
	ResourceArns []string

	// An optional collection of structures that generate CloudFormation parameters to
	// override the existing resource property values when the resource is exported to
	// a new CloudFormation template. Use this field if the ExportFormat field of a
	// StartAssetBundleExportJobRequest API call is set to CLOUDFORMATION_JSON .
	CloudFormationOverridePropertyConfiguration *types.AssetBundleCloudFormationOverridePropertyConfiguration

	// A Boolean that determines whether all dependencies of each resource ARN are
	// recursively exported with the job. For example, say you provided a Dashboard ARN
	// to the ResourceArns parameter. If you set IncludeAllDependencies to TRUE , any
	// theme, dataset, and data source resource that is a dependency of the dashboard
	// is also exported.
	IncludeAllDependencies bool

	noSmithyDocumentSerde
}

type StartAssetBundleExportJobOutput struct {

	// The Amazon Resource Name (ARN) for the export job.
	Arn *string

	// The ID of the job. This ID is unique while the job is running. After the job is
	// completed, you can reuse this ID for another job.
	AssetBundleExportJobId *string

	// The Amazon Web Services response ID for this operation.
	RequestId *string

	// The HTTP status of the response.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAssetBundleExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartAssetBundleExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAssetBundleExportJob{}, middleware.After)
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
	if err = addOpStartAssetBundleExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAssetBundleExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartAssetBundleExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "StartAssetBundleExportJob",
	}
}
