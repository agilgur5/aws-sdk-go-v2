// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This API is identical to GetTemporaryTableCredentials except that this is used
// when the target Data Catalog resource is of type Partition. Lake Formation
// restricts the permission of the vended credentials with the same scope down
// policy which restricts access to a single Amazon S3 prefix.
func (c *Client) GetTemporaryGluePartitionCredentials(ctx context.Context, params *GetTemporaryGluePartitionCredentialsInput, optFns ...func(*Options)) (*GetTemporaryGluePartitionCredentialsOutput, error) {
	if params == nil {
		params = &GetTemporaryGluePartitionCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTemporaryGluePartitionCredentials", params, optFns, c.addOperationGetTemporaryGluePartitionCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTemporaryGluePartitionCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTemporaryGluePartitionCredentialsInput struct {

	// A list of partition values identifying a single partition.
	//
	// This member is required.
	Partition *types.PartitionValueList

	// The ARN of the partitions' table.
	//
	// This member is required.
	TableArn *string

	// A structure representing context to access a resource (column names, query ID,
	// etc).
	AuditContext *types.AuditContext

	// The time period, between 900 and 21,600 seconds, for the timeout of the
	// temporary credentials.
	DurationSeconds *int32

	// Filters the request based on the user having been granted a list of specified
	// permissions on the requested resource(s).
	Permissions []types.Permission

	// A list of supported permission types for the partition. Valid values are
	// COLUMN_PERMISSION and CELL_FILTER_PERMISSION .
	SupportedPermissionTypes []types.PermissionType

	noSmithyDocumentSerde
}

type GetTemporaryGluePartitionCredentialsOutput struct {

	// The access key ID for the temporary credentials.
	AccessKeyId *string

	// The date and time when the temporary credentials expire.
	Expiration *time.Time

	// The secret key for the temporary credentials.
	SecretAccessKey *string

	// The session token for the temporary credentials.
	SessionToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTemporaryGluePartitionCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTemporaryGluePartitionCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTemporaryGluePartitionCredentials{}, middleware.After)
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
	if err = addOpGetTemporaryGluePartitionCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemporaryGluePartitionCredentials(options.Region), middleware.Before); err != nil {
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
		operation: "GetTemporaryGluePartitionCredentials",
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

func newServiceMetadataMiddleware_opGetTemporaryGluePartitionCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GetTemporaryGluePartitionCredentials",
	}
}
