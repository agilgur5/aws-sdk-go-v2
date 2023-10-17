// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates the specified KMS key from the specified log group or from all
// CloudWatch Logs Insights query results in the account. When you use
// DisassociateKmsKey , you specify either the logGroupName parameter or the
// resourceIdentifier parameter. You can't specify both of those parameters in the
// same operation.
//   - Specify the logGroupName parameter to stop using the KMS key to encrypt
//     future log events ingested and stored in the log group. Instead, they will be
//     encrypted with the default CloudWatch Logs method. The log events that were
//     ingested while the key was associated with the log group are still encrypted
//     with that key. Therefore, CloudWatch Logs will need permissions for the key
//     whenever that data is accessed.
//   - Specify the resourceIdentifier parameter with the query-result resource to
//     stop using the KMS key to encrypt the results of all future StartQuery (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html)
//     operations in the account. They will instead be encrypted with the default
//     CloudWatch Logs method. The results from queries that ran while the key was
//     associated with the account are still encrypted with that key. Therefore,
//     CloudWatch Logs will need permissions for the key whenever that data is
//     accessed.
//
// It can take up to 5 minutes for this operation to take effect.
func (c *Client) DisassociateKmsKey(ctx context.Context, params *DisassociateKmsKeyInput, optFns ...func(*Options)) (*DisassociateKmsKeyOutput, error) {
	if params == nil {
		params = &DisassociateKmsKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateKmsKey", params, optFns, c.addOperationDisassociateKmsKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateKmsKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateKmsKeyInput struct {

	// The name of the log group. In your DisassociateKmsKey operation, you must
	// specify either the resourceIdentifier parameter or the logGroup parameter, but
	// you can't specify both.
	LogGroupName *string

	// Specifies the target for this operation. You must specify one of the following:
	//   - Specify the ARN of a log group to stop having CloudWatch Logs use the KMS
	//   key to encrypt log events that are ingested and stored by that log group. After
	//   you run this operation, CloudWatch Logs encrypts ingested log events with the
	//   default CloudWatch Logs method. The log group ARN must be in the following
	//   format. Replace REGION and ACCOUNT_ID with your Region and account ID.
	//   arn:aws:logs:REGION:ACCOUNT_ID:log-group:LOG_GROUP_NAME
	//   - Specify the following ARN to stop using this key to encrypt the results of
	//   future StartQuery (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html)
	//   operations in this account. Replace REGION and ACCOUNT_ID with your Region and
	//   account ID. arn:aws:logs:REGION:ACCOUNT_ID:query-result:*
	// In your DisssociateKmsKey operation, you must specify either the
	// resourceIdentifier parameter or the logGroup parameter, but you can't specify
	// both.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type DisassociateKmsKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateKmsKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateKmsKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateKmsKey{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateKmsKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateKmsKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DisassociateKmsKey",
	}
}
