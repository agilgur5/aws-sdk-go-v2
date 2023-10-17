// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// . This operation does not affect any backing elements of the custom key store.
// It does not delete the CloudHSM cluster that is associated with an CloudHSM key
// store, or affect any users or keys in the cluster. For an external key store, it
// does not affect the external key store proxy, external key manager, or any
// external keys. This operation is part of the custom key stores (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in KMS, which combines the convenience and extensive integration of KMS
// with the isolation and control of a key store that you own and manage. The
// custom key store that you delete cannot contain any KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys)
// . Before deleting the key store, verify that you will never need to use any of
// the KMS keys in the key store for any cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
// . Then, use ScheduleKeyDeletion to delete the KMS keys from the key store.
// After the required waiting period expires and all KMS keys are deleted from the
// custom key store, use DisconnectCustomKeyStore to disconnect the key store from
// KMS. Then, you can delete the custom key store. For keys in an CloudHSM key
// store, the ScheduleKeyDeletion operation makes a best effort to delete the key
// material from the associated cluster. However, you might need to manually
// delete the orphaned key material (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key)
// from the cluster and its backups. KMS never creates, manages, or deletes
// cryptographic keys in the external key manager associated with an external key
// store. You must manage them using your external key manager tools. Instead of
// deleting the custom key store, consider using the DisconnectCustomKeyStore
// operation to disconnect the custom key store from its backing key store. While
// the key store is disconnected, you cannot create or use the KMS keys in the key
// store. But, you do not need to delete KMS keys and you can reconnect a
// disconnected custom key store at any time. If the operation succeeds, it returns
// a JSON object with no properties. Cross-account use: No. You cannot perform this
// operation on a custom key store in a different Amazon Web Services account.
// Required permissions: kms:DeleteCustomKeyStore (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (IAM policy) Related operations:
//   - ConnectCustomKeyStore
//   - CreateCustomKeyStore
//   - DescribeCustomKeyStores
//   - DisconnectCustomKeyStore
//   - UpdateCustomKeyStore
func (c *Client) DeleteCustomKeyStore(ctx context.Context, params *DeleteCustomKeyStoreInput, optFns ...func(*Options)) (*DeleteCustomKeyStoreOutput, error) {
	if params == nil {
		params = &DeleteCustomKeyStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomKeyStore", params, optFns, c.addOperationDeleteCustomKeyStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCustomKeyStoreInput struct {

	// Enter the ID of the custom key store you want to delete. To find the ID of a
	// custom key store, use the DescribeCustomKeyStores operation.
	//
	// This member is required.
	CustomKeyStoreId *string

	noSmithyDocumentSerde
}

type DeleteCustomKeyStoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCustomKeyStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCustomKeyStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCustomKeyStore{}, middleware.After)
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
	if err = addOpDeleteCustomKeyStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomKeyStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCustomKeyStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DeleteCustomKeyStore",
	}
}
