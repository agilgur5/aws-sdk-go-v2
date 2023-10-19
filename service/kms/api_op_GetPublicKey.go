// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the public key of an asymmetric KMS key. Unlike the private key of a
// asymmetric KMS key, which never leaves KMS unencrypted, callers with
// kms:GetPublicKey permission can download the public key of an asymmetric KMS
// key. You can share the public key to allow others to encrypt messages and verify
// signatures outside of KMS. For information about asymmetric KMS keys, see
// Asymmetric KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the Key Management Service Developer Guide. You do not need to download the
// public key. Instead, you can use the public key within KMS by calling the
// Encrypt , ReEncrypt , or Verify operations with the identifier of an asymmetric
// KMS key. When you use the public key within KMS, you benefit from the
// authentication, authorization, and logging that are part of every KMS operation.
// You also reduce of risk of encrypting data that cannot be decrypted. These
// features are not effective outside of KMS. To help you use the public key safely
// outside of KMS, GetPublicKey returns important information about the public key
// in the response, including:
//   - KeySpec (https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeySpec)
//     : The type of key material in the public key, such as RSA_4096 or
//     ECC_NIST_P521 .
//   - KeyUsage (https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeyUsage)
//     : Whether the key is used for encryption or signing.
//   - EncryptionAlgorithms (https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-EncryptionAlgorithms)
//     or SigningAlgorithms (https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-SigningAlgorithms)
//     : A list of the encryption algorithms or the signing algorithms for the key.
//
// Although KMS cannot enforce these restrictions on external operations, it is
// crucial that you use this information to prevent the public key from being used
// improperly. For example, you can prevent a public signing key from being used
// encrypt data, or prevent a public key from being used with an encryption
// algorithm that is not supported by KMS. You can also avoid errors, such as using
// the wrong signing algorithm in a verification operation. To verify a signature
// outside of KMS with an SM2 public key (China Regions only), you must specify the
// distinguishing ID. By default, KMS uses 1234567812345678 as the distinguishing
// ID. For more information, see Offline verification with SM2 key pairs (https://docs.aws.amazon.com/kms/latest/developerguide/asymmetric-key-specs.html#key-spec-sm-offline-verification)
// . The KMS key that you use for this operation must be in a compatible key state.
// For details, see Key states of KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the Key Management Service Developer Guide. Cross-account use: Yes. To
// perform this operation with a KMS key in a different Amazon Web Services
// account, specify the key ARN or alias ARN in the value of the KeyId parameter.
// Required permissions: kms:GetPublicKey (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations: CreateKey
func (c *Client) GetPublicKey(ctx context.Context, params *GetPublicKeyInput, optFns ...func(*Options)) (*GetPublicKeyOutput, error) {
	if params == nil {
		params = &GetPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPublicKey", params, optFns, c.addOperationGetPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPublicKeyInput struct {

	// Identifies the asymmetric KMS key that includes the public key. To specify a
	// KMS key, use its key ID, key ARN, alias name, or alias ARN. When using an alias
	// name, prefix it with "alias/" . To specify a KMS key in a different Amazon Web
	// Services account, you must use the key ARN or alias ARN. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Alias name: alias/ExampleAlias
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey . To
	// get the alias name and alias ARN, use ListAliases .
	//
	// This member is required.
	KeyId *string

	// A list of grant tokens. Use a grant token when your permission to call this
	// operation comes from a new grant that has not yet achieved eventual consistency.
	// For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []string

	noSmithyDocumentSerde
}

type GetPublicKeyOutput struct {

	// Instead, use the KeySpec field in the GetPublicKey response. The KeySpec and
	// CustomerMasterKeySpec fields have the same value. We recommend that you use the
	// KeySpec field in your code. However, to avoid breaking changes, KMS supports
	// both fields.
	//
	// Deprecated: This field has been deprecated. Instead, use the KeySpec field.
	CustomerMasterKeySpec types.CustomerMasterKeySpec

	// The encryption algorithms that KMS supports for this key. This information is
	// critical. If a public key encrypts data outside of KMS by using an unsupported
	// encryption algorithm, the ciphertext cannot be decrypted. This field appears in
	// the response only when the KeyUsage of the public key is ENCRYPT_DECRYPT .
	EncryptionAlgorithms []types.EncryptionAlgorithmSpec

	// The Amazon Resource Name ( key ARN (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN)
	// ) of the asymmetric KMS key from which the public key was downloaded.
	KeyId *string

	// The type of the of the public key that was downloaded.
	KeySpec types.KeySpec

	// The permitted use of the public key. Valid values are ENCRYPT_DECRYPT or
	// SIGN_VERIFY . This information is critical. If a public key with SIGN_VERIFY
	// key usage encrypts data outside of KMS, the ciphertext cannot be decrypted.
	KeyUsage types.KeyUsageType

	// The exported public key. The value is a DER-encoded X.509 public key, also
	// known as SubjectPublicKeyInfo (SPKI), as defined in RFC 5280 (https://tools.ietf.org/html/rfc5280)
	// . When you use the HTTP API or the Amazon Web Services CLI, the value is
	// Base64-encoded. Otherwise, it is not Base64-encoded.
	PublicKey []byte

	// The signing algorithms that KMS supports for this key. This field appears in
	// the response only when the KeyUsage of the public key is SIGN_VERIFY .
	SigningAlgorithms []types.SigningAlgorithmSpec

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPublicKey{}, middleware.After)
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
	if err = addOpGetPublicKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPublicKey(options.Region), middleware.Before); err != nil {
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
		operation: "GetPublicKey",
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

func newServiceMetadataMiddleware_opGetPublicKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "GetPublicKey",
	}
}
