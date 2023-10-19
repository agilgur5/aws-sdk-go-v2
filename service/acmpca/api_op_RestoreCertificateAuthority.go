// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Restores a certificate authority (CA) that is in the DELETED state. You can
// restore a CA during the period that you defined in the
// PermanentDeletionTimeInDays parameter of the DeleteCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeleteCertificateAuthority.html)
// action. Currently, you can specify 7 to 30 days. If you did not specify a
// PermanentDeletionTimeInDays value, by default you can restore the CA at any time
// in a 30 day period. You can check the time remaining in the restoration period
// of a private CA in the DELETED state by calling the DescribeCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_DescribeCertificateAuthority.html)
// or ListCertificateAuthorities (https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html)
// actions. The status of a restored CA is set to its pre-deletion status when the
// RestoreCertificateAuthority action returns. To change its status to ACTIVE ,
// call the UpdateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html)
// action. If the private CA was in the PENDING_CERTIFICATE state at deletion, you
// must use the ImportCertificateAuthorityCertificate (https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html)
// action to import a certificate authority into the private CA before it can be
// activated. You cannot restore a CA after the restoration period has ended.
func (c *Client) RestoreCertificateAuthority(ctx context.Context, params *RestoreCertificateAuthorityInput, optFns ...func(*Options)) (*RestoreCertificateAuthorityOutput, error) {
	if params == nil {
		params = &RestoreCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreCertificateAuthority", params, optFns, c.addOperationRestoreCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreCertificateAuthorityInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called the
	// CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
	// action. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	noSmithyDocumentSerde
}

type RestoreCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreCertificateAuthority{}, middleware.After)
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
	if err = addOpRestoreCertificateAuthorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreCertificateAuthority(options.Region), middleware.Before); err != nil {
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
		operation: "RestoreCertificateAuthority",
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

func newServiceMetadataMiddleware_opRestoreCertificateAuthority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "RestoreCertificateAuthority",
	}
}
