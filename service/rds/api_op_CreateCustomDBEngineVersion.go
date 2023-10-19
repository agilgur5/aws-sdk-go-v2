// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a custom DB engine version (CEV).
func (c *Client) CreateCustomDBEngineVersion(ctx context.Context, params *CreateCustomDBEngineVersionInput, optFns ...func(*Options)) (*CreateCustomDBEngineVersionOutput, error) {
	if params == nil {
		params = &CreateCustomDBEngineVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomDBEngineVersion", params, optFns, c.addOperationCreateCustomDBEngineVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomDBEngineVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomDBEngineVersionInput struct {

	// The database engine to use for your custom engine version (CEV). The only
	// supported value is custom-oracle-ee .
	//
	// This member is required.
	Engine *string

	// The name of your CEV. The name format is 19.customized_string. For example, a
	// valid CEV name is 19.my_cev1 . This setting is required for RDS Custom for
	// Oracle, but optional for Amazon RDS. The combination of Engine and EngineVersion
	// is unique per customer per Region.
	//
	// This member is required.
	EngineVersion *string

	// The name of an Amazon S3 bucket that contains database installation files for
	// your CEV. For example, a valid bucket name is my-custom-installation-files .
	DatabaseInstallationFilesS3BucketName *string

	// The Amazon S3 directory that contains the database installation files for your
	// CEV. For example, a valid bucket name is 123456789012/cev1 . If this setting
	// isn't specified, no prefix is assumed.
	DatabaseInstallationFilesS3Prefix *string

	// An optional description of your CEV.
	Description *string

	// The ID of the Amazon Machine Image (AMI). For RDS Custom for SQL Server, an AMI
	// ID is required to create a CEV. For RDS Custom for Oracle, the default is the
	// most recent AMI available, but you can specify an AMI ID that was used in a
	// different Oracle CEV. Find the AMIs used by your CEVs by calling the
	// DescribeDBEngineVersions (https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeDBEngineVersions.html)
	// operation.
	ImageId *string

	// The Amazon Web Services KMS key identifier for an encrypted CEV. A symmetric
	// encryption KMS key is required for RDS Custom, but optional for Amazon RDS. If
	// you have an existing symmetric encryption KMS key in your account, you can use
	// it with RDS Custom. No further action is necessary. If you don't already have a
	// symmetric encryption KMS key in your account, follow the instructions in
	// Creating a symmetric encryption KMS key (https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html#create-symmetric-cmk)
	// in the Amazon Web Services Key Management Service Developer Guide. You can
	// choose the same symmetric encryption key when you create a CEV and a DB
	// instance, or choose different keys.
	KMSKeyId *string

	// The CEV manifest, which is a JSON document that describes the installation .zip
	// files stored in Amazon S3. Specify the name/value pairs in a file or a quoted
	// string. RDS Custom applies the patches in the order in which they are listed.
	// The following JSON fields are valid: MediaImportTemplateVersion Version of the
	// CEV manifest. The date is in the format YYYY-MM-DD .
	// databaseInstallationFileNames Ordered list of installation files for the CEV.
	// opatchFileNames Ordered list of OPatch installers used for the Oracle DB engine.
	// psuRuPatchFileNames The PSU and RU patches for this CEV. OtherPatchFileNames The
	// patches that are not in the list of PSU and RU patches. Amazon RDS applies these
	// patches after applying the PSU and RU patches. For more information, see
	// Creating the CEV manifest (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.preparing.manifest)
	// in the Amazon RDS User Guide.
	Manifest *string

	// Reserved for future use.
	SourceCustomDbEngineVersionIdentifier *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []types.Tag

	// Reserved for future use.
	UseAwsProvidedLatestImage *bool

	noSmithyDocumentSerde
}

// This data type is used as a response element in the action
// DescribeDBEngineVersions .
type CreateCustomDBEngineVersionOutput struct {

	// The creation time of the DB engine version.
	CreateTime *time.Time

	// JSON string that lists the installation files and parameters that RDS Custom
	// uses to create a custom engine version (CEV). RDS Custom applies the patches in
	// the order in which they're listed in the manifest. You can set the Oracle home,
	// Oracle base, and UNIX/Linux user and group using the installation parameters.
	// For more information, see JSON fields in the CEV manifest (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.preparing.html#custom-cev.preparing.manifest.fields)
	// in the Amazon RDS User Guide.
	CustomDBEngineVersionManifest *string

	// The description of the database engine.
	DBEngineDescription *string

	// A value that indicates the source media provider of the AMI based on the usage
	// operation. Applicable for RDS Custom for SQL Server.
	DBEngineMediaType *string

	// The ARN of the custom engine version.
	DBEngineVersionArn *string

	// The description of the database engine version.
	DBEngineVersionDescription *string

	// The name of the DB parameter group family for the database engine.
	DBParameterGroupFamily *string

	// The name of the Amazon S3 bucket that contains your database installation files.
	DatabaseInstallationFilesS3BucketName *string

	// The Amazon S3 directory that contains the database installation files. If not
	// specified, then no prefix is assumed.
	DatabaseInstallationFilesS3Prefix *string

	// The default character set for new instances of this engine version, if the
	// CharacterSetName parameter of the CreateDBInstance API isn't specified.
	DefaultCharacterSet *types.CharacterSet

	// The name of the database engine.
	Engine *string

	// The version number of the database engine.
	EngineVersion *string

	// The types of logs that the database engine has available for export to
	// CloudWatch Logs.
	ExportableLogTypes []string

	// The EC2 image
	Image *types.CustomDBEngineVersionAMI

	// The Amazon Web Services KMS key identifier for an encrypted CEV. This parameter
	// is required for RDS Custom, but optional for Amazon RDS.
	KMSKeyId *string

	// The major engine version of the CEV.
	MajorEngineVersion *string

	// The status of the DB engine version, either available or deprecated .
	Status *string

	// A list of the supported CA certificate identifiers. For more information, see
	// Using SSL/TLS to encrypt a connection to a DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html)
	// in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection to a DB
	// cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html)
	// in the Amazon Aurora User Guide.
	SupportedCACertificateIdentifiers []string

	// A list of the character sets supported by this engine for the CharacterSetName
	// parameter of the CreateDBInstance operation.
	SupportedCharacterSets []types.CharacterSet

	// A list of the supported DB engine modes.
	SupportedEngineModes []string

	// A list of features supported by the DB engine. The supported features vary by
	// DB engine and DB engine version. To determine the supported features for a
	// specific DB engine and DB engine version using the CLI, use the following
	// command: aws rds describe-db-engine-versions --engine --engine-version  For
	// example, to determine the supported features for RDS for PostgreSQL version 13.3
	// using the CLI, use the following command: aws rds describe-db-engine-versions
	// --engine postgres --engine-version 13.3 The supported features are listed under
	// SupportedFeatureNames in the output.
	SupportedFeatureNames []string

	// A list of the character sets supported by the Oracle DB engine for the
	// NcharCharacterSetName parameter of the CreateDBInstance operation.
	SupportedNcharCharacterSets []types.CharacterSet

	// A list of the time zones supported by this engine for the Timezone parameter of
	// the CreateDBInstance action.
	SupportedTimezones []types.Timezone

	// Indicates whether the engine version supports Babelfish for Aurora PostgreSQL.
	SupportsBabelfish bool

	// Indicates whether the engine version supports rotating the server certificate
	// without rebooting the DB instance.
	SupportsCertificateRotationWithoutRestart *bool

	// Indicates whether you can use Aurora global databases with a specific DB engine
	// version.
	SupportsGlobalDatabases bool

	// Indicates whether the DB engine version supports forwarding write operations
	// from reader DB instances to the writer DB instance in the DB cluster. By
	// default, write operations aren't allowed on reader DB instances. Valid for:
	// Aurora DB clusters only
	SupportsLocalWriteForwarding *bool

	// Indicates whether the engine version supports exporting the log types specified
	// by ExportableLogTypes to CloudWatch Logs.
	SupportsLogExportsToCloudwatchLogs bool

	// Indicates whether you can use Aurora parallel query with a specific DB engine
	// version.
	SupportsParallelQuery bool

	// Indicates whether the database engine version supports read replicas.
	SupportsReadReplica bool

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	TagList []types.Tag

	// A list of engine versions that this database engine version can be upgraded to.
	ValidUpgradeTarget []types.UpgradeTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomDBEngineVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateCustomDBEngineVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCustomDBEngineVersion{}, middleware.After)
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
	if err = addOpCreateCustomDBEngineVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomDBEngineVersion(options.Region), middleware.Before); err != nil {
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
		operation: "CreateCustomDBEngineVersion",
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

func newServiceMetadataMiddleware_opCreateCustomDBEngineVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateCustomDBEngineVersion",
	}
}
