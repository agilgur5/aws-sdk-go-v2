// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports a certificate into Certificate Manager (ACM) to use with services that
// are integrated with ACM. Note that integrated services (https://docs.aws.amazon.com/acm/latest/userguide/acm-services.html)
// allow only certificate types and keys they support to be associated with their
// resources. Further, their support differs depending on whether the certificate
// is imported into IAM or into ACM. For more information, see the documentation
// for each service. For more information about importing certificates into ACM,
// see Importing Certificates (https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html)
// in the Certificate Manager User Guide. ACM does not provide managed renewal (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html)
// for certificates that you import. Note the following guidelines when importing
// third party certificates:
//   - You must enter the private key that matches the certificate you are
//     importing.
//   - The private key must be unencrypted. You cannot import a private key that
//     is protected by a password or a passphrase.
//   - The private key must be no larger than 5 KB (5,120 bytes).
//   - If the certificate you are importing is not self-signed, you must enter its
//     certificate chain.
//   - If a certificate chain is included, the issuer must be the subject of one
//     of the certificates in the chain.
//   - The certificate, private key, and certificate chain must be PEM-encoded.
//   - The current time must be between the Not Before and Not After certificate
//     fields.
//   - The Issuer field must not be empty.
//   - The OCSP authority URL, if present, must not exceed 1000 characters.
//   - To import a new certificate, omit the CertificateArn argument. Include this
//     argument only when you want to replace a previously imported certificate.
//   - When you import a certificate by using the CLI, you must specify the
//     certificate, the certificate chain, and the private key by their file names
//     preceded by fileb:// . For example, you can specify a certificate saved in the
//     C:\temp folder as fileb://C:\temp\certificate_to_import.pem . If you are
//     making an HTTP or HTTPS Query request, include these arguments as BLOBs.
//   - When you import a certificate by using an SDK, you must specify the
//     certificate, the certificate chain, and the private key files in the manner
//     required by the programming language you're using.
//   - The cryptographic algorithm of an imported certificate must match the
//     algorithm of the signing CA. For example, if the signing CA key type is RSA,
//     then the certificate key type must also be RSA.
//
// This operation returns the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
// of the imported certificate.
func (c *Client) ImportCertificate(ctx context.Context, params *ImportCertificateInput, optFns ...func(*Options)) (*ImportCertificateOutput, error) {
	if params == nil {
		params = &ImportCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportCertificate", params, optFns, c.addOperationImportCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportCertificateInput struct {

	// The certificate to import.
	//
	// This member is required.
	Certificate []byte

	// The private key that matches the public key in the certificate.
	//
	// This member is required.
	PrivateKey []byte

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of an imported certificate to replace. To import a new certificate, omit this
	// field.
	CertificateArn *string

	// The PEM encoded certificate chain.
	CertificateChain []byte

	// One or more resource tags to associate with the imported certificate. Note: You
	// cannot apply tags when reimporting a certificate.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ImportCertificateOutput struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the imported certificate.
	CertificateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportCertificate{}, middleware.After)
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
	if err = addOpImportCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportCertificate(options.Region), middleware.Before); err != nil {
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
		operation: "ImportCertificate",
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

func newServiceMetadataMiddleware_opImportCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm",
		OperationName: "ImportCertificate",
	}
}
