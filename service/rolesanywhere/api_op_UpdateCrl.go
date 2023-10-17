// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the certificate revocation list (CRL). A CRL is a list of certificates
// that have been revoked by the issuing certificate authority (CA). IAM Roles
// Anywhere validates against the CRL before issuing credentials. Required
// permissions: rolesanywhere:UpdateCrl .
func (c *Client) UpdateCrl(ctx context.Context, params *UpdateCrlInput, optFns ...func(*Options)) (*UpdateCrlOutput, error) {
	if params == nil {
		params = &UpdateCrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCrl", params, optFns, c.addOperationUpdateCrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCrlInput struct {

	// The unique identifier of the certificate revocation list (CRL).
	//
	// This member is required.
	CrlId *string

	// The x509 v3 specified certificate revocation list (CRL).
	CrlData []byte

	// The name of the Crl.
	Name *string

	noSmithyDocumentSerde
}

type UpdateCrlOutput struct {

	// The state of the certificate revocation list (CRL) after a read or write
	// operation.
	//
	// This member is required.
	Crl *types.CrlDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCrl{}, middleware.After)
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
	if err = addOpUpdateCrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rolesanywhere",
		OperationName: "UpdateCrl",
	}
}
