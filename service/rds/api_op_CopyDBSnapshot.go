// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies the specified DB snapshot. The source DB snapshot must be in the
// available state. You can copy a snapshot from one Amazon Web Services Region to
// another. In that case, the Amazon Web Services Region where you call the
// CopyDBSnapshot operation is the destination Amazon Web Services Region for the
// DB snapshot copy. This command doesn't apply to RDS Custom. For more information
// about copying snapshots, see Copying a DB Snapshot (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CopySnapshot.html#USER_CopyDBSnapshot)
// in the Amazon RDS User Guide.
func (c *Client) CopyDBSnapshot(ctx context.Context, params *CopyDBSnapshotInput, optFns ...func(*Options)) (*CopyDBSnapshotOutput, error) {
	if params == nil {
		params = &CopyDBSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDBSnapshot", params, optFns, c.addOperationCopyDBSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDBSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyDBSnapshotInput struct {

	// The identifier for the source DB snapshot. If the source snapshot is in the
	// same Amazon Web Services Region as the copy, specify a valid DB snapshot
	// identifier. For example, you might specify rds:mysql-instance1-snapshot-20130805
	// . If the source snapshot is in a different Amazon Web Services Region than the
	// copy, specify a valid DB snapshot ARN. For example, you might specify
	// arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20130805 .
	// If you are copying from a shared manual DB snapshot, this parameter must be the
	// Amazon Resource Name (ARN) of the shared DB snapshot. If you are copying an
	// encrypted snapshot this parameter must be in the ARN format for the source
	// Amazon Web Services Region. Constraints:
	//   - Must specify a valid system snapshot in the "available" state.
	// Example: rds:mydb-2012-04-02-00-01 Example:
	// arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20130805
	//
	// This member is required.
	SourceDBSnapshotIdentifier *string

	// The identifier for the copy of the snapshot. Constraints:
	//   - Can't be null, empty, or blank
	//   - Must contain from 1 to 255 letters, numbers, or hyphens
	//   - First character must be a letter
	//   - Can't end with a hyphen or contain two consecutive hyphens
	// Example: my-db-snapshot
	//
	// This member is required.
	TargetDBSnapshotIdentifier *string

	// Specifies whether to copy the DB option group associated with the source DB
	// snapshot to the target Amazon Web Services account and associate with the target
	// DB snapshot. The associated option group can be copied only with cross-account
	// snapshot copy calls.
	CopyOptionGroup *bool

	// Specifies whether to copy all tags from the source DB snapshot to the target DB
	// snapshot. By default, tags aren't copied.
	CopyTags *bool

	// The Amazon Web Services KMS key identifier for an encrypted DB snapshot. The
	// Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or
	// alias name for the KMS key. If you copy an encrypted DB snapshot from your
	// Amazon Web Services account, you can specify a value for this parameter to
	// encrypt the copy with a new KMS key. If you don't specify a value for this
	// parameter, then the copy of the DB snapshot is encrypted with the same Amazon
	// Web Services KMS key as the source DB snapshot. If you copy an encrypted DB
	// snapshot that is shared from another Amazon Web Services account, then you must
	// specify a value for this parameter. If you specify this parameter when you copy
	// an unencrypted snapshot, the copy is encrypted. If you copy an encrypted
	// snapshot to a different Amazon Web Services Region, then you must specify an
	// Amazon Web Services KMS key identifier for the destination Amazon Web Services
	// Region. KMS keys are specific to the Amazon Web Services Region that they are
	// created in, and you can't use KMS keys from one Amazon Web Services Region in
	// another Amazon Web Services Region.
	KmsKeyId *string

	// The name of an option group to associate with the copy of the snapshot. Specify
	// this option if you are copying a snapshot from one Amazon Web Services Region to
	// another, and your DB instance uses a nondefault option group. If your source DB
	// instance uses Transparent Data Encryption for Oracle or Microsoft SQL Server,
	// you must specify this option when copying across Amazon Web Services Regions.
	// For more information, see Option group considerations (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CopySnapshot.html#USER_CopySnapshot.Options)
	// in the Amazon RDS User Guide.
	OptionGroupName *string

	// When you are copying a snapshot from one Amazon Web Services GovCloud (US)
	// Region to another, the URL that contains a Signature Version 4 signed request
	// for the CopyDBSnapshot API operation in the source Amazon Web Services Region
	// that contains the source DB snapshot to copy. This setting applies only to
	// Amazon Web Services GovCloud (US) Regions. It's ignored in other Amazon Web
	// Services Regions. You must specify this parameter when you copy an encrypted DB
	// snapshot from another Amazon Web Services Region by using the Amazon RDS API.
	// Don't specify PreSignedUrl when you are copying an encrypted DB snapshot in the
	// same Amazon Web Services Region. The presigned URL must be a valid request for
	// the CopyDBClusterSnapshot API operation that can run in the source Amazon Web
	// Services Region that contains the encrypted DB cluster snapshot to copy. The
	// presigned URL request must contain the following parameter values:
	//   - DestinationRegion - The Amazon Web Services Region that the encrypted DB
	//   snapshot is copied to. This Amazon Web Services Region is the same one where the
	//   CopyDBSnapshot operation is called that contains this presigned URL. For
	//   example, if you copy an encrypted DB snapshot from the us-west-2 Amazon Web
	//   Services Region to the us-east-1 Amazon Web Services Region, then you call the
	//   CopyDBSnapshot operation in the us-east-1 Amazon Web Services Region and
	//   provide a presigned URL that contains a call to the CopyDBSnapshot operation
	//   in the us-west-2 Amazon Web Services Region. For this example, the
	//   DestinationRegion in the presigned URL must be set to the us-east-1 Amazon Web
	//   Services Region.
	//   - KmsKeyId - The KMS key identifier for the KMS key to use to encrypt the copy
	//   of the DB snapshot in the destination Amazon Web Services Region. This is the
	//   same identifier for both the CopyDBSnapshot operation that is called in the
	//   destination Amazon Web Services Region, and the operation contained in the
	//   presigned URL.
	//   - SourceDBSnapshotIdentifier - The DB snapshot identifier for the encrypted
	//   snapshot to be copied. This identifier must be in the Amazon Resource Name (ARN)
	//   format for the source Amazon Web Services Region. For example, if you are
	//   copying an encrypted DB snapshot from the us-west-2 Amazon Web Services Region,
	//   then your SourceDBSnapshotIdentifier looks like the following example:
	//   arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20161115
	//   .
	// To learn how to generate a Signature Version 4 signed request, see
	// Authenticating Requests: Using Query Parameters (Amazon Web Services Signature
	// Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
	// . If you are using an Amazon Web Services SDK tool or the CLI, you can specify
	// SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl
	// manually. Specifying SourceRegion autogenerates a presigned URL that is a valid
	// request for the operation that can run in the source Amazon Web Services Region.
	PreSignedUrl *string

	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []types.Tag

	// The external custom Availability Zone (CAZ) identifier for the target CAZ.
	// Example: rds-caz-aiqhTgQv .
	TargetCustomAvailabilityZone *string

	// Used by the SDK's PresignURL autofill customization to specify the region the
	// of the client's request.
	destinationRegion *string

	noSmithyDocumentSerde
}

type CopyDBSnapshotOutput struct {

	// Contains the details of an Amazon RDS DB snapshot. This data type is used as a
	// response element in the DescribeDBSnapshots action.
	DBSnapshot *types.DBSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyDBSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBSnapshot{}, middleware.After)
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
	if err = addCopyDBSnapshotPresignURLMiddleware(stack, options); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpCopyDBSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBSnapshot(options.Region), middleware.Before); err != nil {
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

func copyCopyDBSnapshotInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCopyDBSnapshotPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCopyDBSnapshotSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCopyDBSnapshotPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCopyDBSnapshotdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}
func addCopyDBSnapshotPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL: getCopyDBSnapshotPreSignedUrl,

			GetSourceRegion: getCopyDBSnapshotSourceRegion,

			CopyInput: copyCopyDBSnapshotInputForPresign,

			SetDestinationRegion: setCopyDBSnapshotdestinationRegion,

			SetPresignedURL: setCopyDBSnapshotPreSignedUrl,
		},
		Presigner: &presignAutoFillCopyDBSnapshotClient{client: NewPresignClient(New(options))},
	})
}

type presignAutoFillCopyDBSnapshotClient struct {
	client *PresignClient
}

// PresignURL is a middleware accessor that satisfies URLPresigner interface.
func (c *presignAutoFillCopyDBSnapshotClient) PresignURL(ctx context.Context, srcRegion string, params interface{}) (*v4.PresignedHTTPRequest, error) {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = srcRegion
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	presignOptFn := WithPresignClientFromClientOptions(optFn)
	return c.client.PresignCopyDBSnapshot(ctx, input, presignOptFn)
}

func newServiceMetadataMiddleware_opCopyDBSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBSnapshot",
	}
}

// PresignCopyDBSnapshot is used to generate a presigned HTTP Request which
// contains presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignCopyDBSnapshot(ctx context.Context, params *CopyDBSnapshotInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &CopyDBSnapshotInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "CopyDBSnapshot", params, clientOptFns,
		c.client.addOperationCopyDBSnapshotMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}
