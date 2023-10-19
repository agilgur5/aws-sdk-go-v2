// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Runs an SQL statement, which can be data manipulation language (DML) or data
// definition language (DDL). This statement must be a single SQL statement.
// Depending on the authorization method, use one of the following combinations of
// request parameters:
//   - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
//     secret stored in Secrets Manager which has username and password . The
//     specified secret contains credentials to connect to the database you specify.
//     When you are connecting to a cluster, you also supply the database name, If you
//     provide a cluster identifier ( dbClusterIdentifier ), it must match the
//     cluster identifier stored in the secret. When you are connecting to a serverless
//     workgroup, you also supply the database name.
//   - Temporary credentials - when connecting to your data warehouse, choose one
//     of the following options:
//   - When connecting to a serverless workgroup, specify the workgroup name and
//     database name. The database user name is derived from the IAM identity. For
//     example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
//     Also, permission to call the redshift-serverless:GetCredentials operation is
//     required.
//   - When connecting to a cluster as an IAM identity, specify the cluster
//     identifier and the database name. The database user name is derived from the IAM
//     identity. For example, arn:iam::123456789012:user:foo has the database user
//     name IAM:foo . Also, permission to call the
//     redshift:GetClusterCredentialsWithIAM operation is required.
//   - When connecting to a cluster as a database user, specify the cluster
//     identifier, the database name, and the database user name. Also, permission to
//     call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see Using the Amazon Redshift Data API (https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html)
// in the Amazon Redshift Management Guide.
func (c *Client) ExecuteStatement(ctx context.Context, params *ExecuteStatementInput, optFns ...func(*Options)) (*ExecuteStatementOutput, error) {
	if params == nil {
		params = &ExecuteStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteStatement", params, optFns, c.addOperationExecuteStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteStatementInput struct {

	// The name of the database. This parameter is required when authenticating using
	// either Secrets Manager or temporary credentials.
	//
	// This member is required.
	Database *string

	// The SQL statement text to run.
	//
	// This member is required.
	Sql *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The cluster identifier. This parameter is required when connecting to a cluster
	// and authenticating using either Secrets Manager or temporary credentials.
	ClusterIdentifier *string

	// The database user name. This parameter is required when connecting to a cluster
	// as a database user and authenticating using temporary credentials.
	DbUser *string

	// The parameters for the SQL statement.
	Parameters []types.SqlParameter

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using Secrets Manager.
	SecretArn *string

	// The name of the SQL statement. You can name the SQL statement when you create
	// it to identify the query.
	StatementName *string

	// A value that indicates whether to send an event to the Amazon EventBridge event
	// bus after the SQL statement runs.
	WithEvent *bool

	// The serverless workgroup name or Amazon Resource Name (ARN). This parameter is
	// required when connecting to a serverless workgroup and authenticating using
	// either Secrets Manager or temporary credentials.
	WorkgroupName *string

	noSmithyDocumentSerde
}

type ExecuteStatementOutput struct {

	// The cluster identifier. This element is not returned when connecting to a
	// serverless workgroup.
	ClusterIdentifier *string

	// The date and time (UTC) the statement was created.
	CreatedAt *time.Time

	// The name of the database.
	Database *string

	// The database user name.
	DbUser *string

	// The identifier of the SQL statement whose results are to be fetched. This value
	// is a universally unique identifier (UUID) generated by Amazon Redshift Data API.
	Id *string

	// The name or ARN of the secret that enables access to the database.
	SecretArn *string

	// The serverless workgroup name or Amazon Resource Name (ARN). This element is
	// not returned when connecting to a provisioned cluster.
	WorkgroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExecuteStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExecuteStatement{}, middleware.After)
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
	if err = addIdempotencyToken_opExecuteStatementMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpExecuteStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteStatement(options.Region), middleware.Before); err != nil {
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
		operation: "ExecuteStatement",
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

type idempotencyToken_initializeOpExecuteStatement struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpExecuteStatement) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpExecuteStatement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ExecuteStatementInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ExecuteStatementInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opExecuteStatementMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpExecuteStatement{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opExecuteStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "ExecuteStatement",
	}
}
