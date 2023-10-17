// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports a key from Amazon Web Services Payment Cryptography using either ANSI
// X9 TR-34 or TR-31 key export standard. Amazon Web Services Payment Cryptography
// simplifies main or root key exchange process by eliminating the need of a
// paper-based key exchange process. It takes a modern and secure approach based of
// the ANSI X9 TR-34 key exchange standard. You can use ExportKey to export main
// or root keys such as KEK (Key Encryption Key), using asymmetric key exchange
// technique following ANSI X9 TR-34 standard. The ANSI X9 TR-34 standard uses
// asymmetric keys to establishes bi-directional trust between the two parties
// exchanging keys. After which you can export working keys using the ANSI X9 TR-31
// symmetric key exchange standard as mandated by PCI PIN. Using this operation,
// you can share your Amazon Web Services Payment Cryptography generated keys with
// other service partners to perform cryptographic operations outside of Amazon Web
// Services Payment Cryptography TR-34 key export Amazon Web Services Payment
// Cryptography uses TR-34 asymmetric key exchange standard to export main keys
// such as KEK. In TR-34 terminology, the sending party of the key is called Key
// Distribution Host (KDH) and the receiving party of the key is called Key
// Receiving Host (KRH). In key export process, KDH is Amazon Web Services Payment
// Cryptography which initiates key export. KRH is the user receiving the key.
// Before you initiate TR-34 key export, you must obtain an export token by calling
// GetParametersForExport . This operation also returns the signing key certificate
// that KDH uses to sign the wrapped key to generate a TR-34 wrapped key block. The
// export token expires after 7 days. Set the following parameters:
// CertificateAuthorityPublicKeyIdentifier The KeyARN of the certificate chain
// that will sign the wrapping key certificate. This must exist within Amazon Web
// Services Payment Cryptography before you initiate TR-34 key export. If it does
// not exist, you can import it by calling ImportKey for RootCertificatePublicKey .
// ExportToken Obtained from KDH by calling GetParametersForExport .
// WrappingKeyCertificate Amazon Web Services Payment Cryptography uses this to
// wrap the key under export. When this operation is successful, Amazon Web
// Services Payment Cryptography returns the TR-34 wrapped key block. TR-31 key
// export Amazon Web Services Payment Cryptography uses TR-31 symmetric key
// exchange standard to export working keys. In TR-31, you must use a main key such
// as KEK to encrypt or wrap the key under export. To establish a KEK, you can use
// CreateKey or ImportKey . When this operation is successful, Amazon Web Services
// Payment Cryptography returns a TR-31 wrapped key block. Cross-account use: This
// operation can't be used across different Amazon Web Services accounts. Related
// operations:
//   - GetParametersForExport
//   - ImportKey
func (c *Client) ExportKey(ctx context.Context, params *ExportKeyInput, optFns ...func(*Options)) (*ExportKeyOutput, error) {
	if params == nil {
		params = &ExportKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportKey", params, optFns, c.addOperationExportKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportKeyInput struct {

	// The KeyARN of the key under export from Amazon Web Services Payment
	// Cryptography.
	//
	// This member is required.
	ExportKeyIdentifier *string

	// The key block format type, for example, TR-34 or TR-31, to use during key
	// material export.
	//
	// This member is required.
	KeyMaterial types.ExportKeyMaterial

	noSmithyDocumentSerde
}

type ExportKeyOutput struct {

	// The key material under export as a TR-34 or TR-31 wrapped key block.
	WrappedKey *types.WrappedKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpExportKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpExportKey{}, middleware.After)
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
	if err = addOpExportKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "payment-cryptography",
		OperationName: "ExportKey",
	}
}
