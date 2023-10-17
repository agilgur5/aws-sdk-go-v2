// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Runs a SQL statement against a database. If a call isn't part of a transaction
// because it doesn't include the transactionID parameter, changes that result
// from the call are committed automatically. If the binary response data from the
// database is more than 1 MB, the call is terminated.
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

// The request parameters represent the input of a request to run a SQL statement
// against a database.
type ExecuteStatementInput struct {

	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	//
	// This member is required.
	ResourceArn *string

	// The ARN of the secret that enables access to the DB cluster. Enter the database
	// user name and password for the credentials in the secret. For information about
	// creating the secret, see Create a database secret (https://docs.aws.amazon.com/secretsmanager/latest/userguide/create_database_secret.html)
	// .
	//
	// This member is required.
	SecretArn *string

	// The SQL statement to run.
	//
	// This member is required.
	Sql *string

	// A value that indicates whether to continue running the statement after the call
	// times out. By default, the statement stops running when the call times out. For
	// DDL statements, we recommend continuing to run the statement after the call
	// times out. When a DDL statement terminates before it is finished running, it can
	// result in errors and possibly corrupted data structures.
	ContinueAfterTimeout bool

	// The name of the database.
	Database *string

	// A value that indicates whether to format the result set as a single JSON
	// string. This parameter only applies to SELECT statements and is ignored for
	// other types of statements. Allowed values are NONE and JSON . The default value
	// is NONE . The result is returned in the formattedRecords field. For usage
	// information about the JSON format for result sets, see Using the Data API (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html)
	// in the Amazon Aurora User Guide.
	FormatRecordsAs types.RecordsFormatType

	// A value that indicates whether to include metadata in the results.
	IncludeResultMetadata bool

	// The parameters for the SQL statement. Array parameters are not supported.
	Parameters []types.SqlParameter

	// Options that control how the result set is returned.
	ResultSetOptions *types.ResultSetOptions

	// The name of the database schema. Currently, the schema parameter isn't
	// supported.
	Schema *string

	// The identifier of a transaction that was started by using the BeginTransaction
	// operation. Specify the transaction ID of the transaction that you want to
	// include the SQL statement in. If the SQL statement is not part of a transaction,
	// don't set this parameter.
	TransactionId *string

	noSmithyDocumentSerde
}

// The response elements represent the output of a request to run a SQL statement
// against a database.
type ExecuteStatementOutput struct {

	// Metadata for the columns included in the results. This field is blank if the
	// formatRecordsAs parameter is set to JSON .
	ColumnMetadata []types.ColumnMetadata

	// A string value that represents the result set of a SELECT statement in JSON
	// format. This value is only present when the formatRecordsAs parameter is set to
	// JSON . The size limit for this field is currently 10 MB. If the JSON-formatted
	// string representing the result set requires more than 10 MB, the call returns an
	// error.
	FormattedRecords *string

	// Values for fields generated during a DML request. The generatedFields data
	// isn't supported by Aurora PostgreSQL. To get the values of generated fields, use
	// the RETURNING clause. For more information, see Returning Data From Modified
	// Rows (https://www.postgresql.org/docs/10/dml-returning.html) in the PostgreSQL
	// documentation.
	GeneratedFields []types.Field

	// The number of records updated by the request.
	NumberOfRecordsUpdated int64

	// The records returned by the SQL statement. This field is blank if the
	// formatRecordsAs parameter is set to JSON .
	Records [][]types.Field

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExecuteStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExecuteStatement{}, middleware.After)
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
	return nil
}

func newServiceMetadataMiddleware_opExecuteStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds-data",
		OperationName: "ExecuteStatement",
	}
}
