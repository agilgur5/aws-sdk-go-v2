// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified KMS key with either one log group in the account, or
// with all stored CloudWatch Logs query insights results in the account. When you
// use AssociateKmsKey , you specify either the logGroupName parameter or the
// resourceIdentifier parameter. You can't specify both of those parameters in the
// same operation.
//   - Specify the logGroupName parameter to cause all log events stored in the log
//     group to be encrypted with that key. Only the log events ingested after the key
//     is associated are encrypted with that key. Associating a KMS key with a log
//     group overrides any existing associations between the log group and a KMS key.
//     After a KMS key is associated with a log group, all newly ingested data for the
//     log group is encrypted using the KMS key. This association is stored as long as
//     the data encrypted with the KMS key is still within CloudWatch Logs. This
//     enables CloudWatch Logs to decrypt this data whenever it is requested.
//     Associating a key with a log group does not cause the results of queries of that
//     log group to be encrypted with that key. To have query results encrypted with a
//     KMS key, you must use an AssociateKmsKey operation with the resourceIdentifier
//     parameter that specifies a query-result resource.
//   - Specify the resourceIdentifier parameter with a query-result resource, to
//     use that key to encrypt the stored results of all future StartQuery (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html)
//     operations in the account. The response from a GetQueryResults (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetQueryResults.html)
//     operation will still return the query results in plain text. Even if you have
//     not associated a key with your query results, the query results are encrypted
//     when stored, using the default CloudWatch Logs method. If you run a query from a
//     monitoring account that queries logs in a source account, the query results key
//     from the monitoring account, if any, is used.
//
// If you delete the key that is used to encrypt log events or log group query
// results, then all the associated stored log events or query results that were
// encrypted with that key will be unencryptable and unusable. CloudWatch Logs
// supports only symmetric KMS keys. Do not use an associate an asymmetric KMS key
// with your log group or query results. For more information, see Using Symmetric
// and Asymmetric Keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// . It can take up to 5 minutes for this operation to take effect. If you attempt
// to associate a KMS key with a log group but the KMS key does not exist or the
// KMS key is disabled, you receive an InvalidParameterException error.
func (c *Client) AssociateKmsKey(ctx context.Context, params *AssociateKmsKeyInput, optFns ...func(*Options)) (*AssociateKmsKeyOutput, error) {
	if params == nil {
		params = &AssociateKmsKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateKmsKey", params, optFns, c.addOperationAssociateKmsKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateKmsKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateKmsKeyInput struct {

	// The Amazon Resource Name (ARN) of the KMS key to use when encrypting log data.
	// This must be a symmetric KMS key. For more information, see Amazon Resource
	// Names (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms)
	// and Using Symmetric and Asymmetric Keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
	// .
	//
	// This member is required.
	KmsKeyId *string

	// The name of the log group. In your AssociateKmsKey operation, you must specify
	// either the resourceIdentifier parameter or the logGroup parameter, but you
	// can't specify both.
	LogGroupName *string

	// Specifies the target for this operation. You must specify one of the following:
	//   - Specify the following ARN to have future GetQueryResults (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetQueryResults.html)
	//   operations in this account encrypt the results with the specified KMS key.
	//   Replace REGION and ACCOUNT_ID with your Region and account ID.
	//   arn:aws:logs:REGION:ACCOUNT_ID:query-result:*
	//   - Specify the ARN of a log group to have CloudWatch Logs use the KMS key to
	//   encrypt log events that are ingested and stored by that log group. The log group
	//   ARN must be in the following format. Replace REGION and ACCOUNT_ID with your
	//   Region and account ID.
	//   arn:aws:logs:REGION:ACCOUNT_ID:log-group:LOG_GROUP_NAME
	// In your AssociateKmsKey operation, you must specify either the
	// resourceIdentifier parameter or the logGroup parameter, but you can't specify
	// both.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type AssociateKmsKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateKmsKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateKmsKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateKmsKey{}, middleware.After)
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
	if err = addOpAssociateKmsKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateKmsKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateKmsKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "AssociateKmsKey",
	}
}
