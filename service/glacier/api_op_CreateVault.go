// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation creates a new vault with the specified name. The name of the
// vault must be unique within a region for an AWS account. You can create up to
// 1,000 vaults per account. If you need to create more vaults, contact Amazon S3
// Glacier. You must use the following guidelines when naming a vault.
//   - Names can be between 1 and 255 characters long.
//   - Allowed characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and
//     '.' (period).
//
// This operation is idempotent. An AWS account has full permission to perform all
// operations (actions). However, AWS Identity and Access Management (IAM) users
// don't have any permissions by default. You must grant them explicit permission
// to perform specific actions. For more information, see Access Control Using AWS
// Identity and Access Management (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html)
// . For conceptual information and underlying REST API, see Creating a Vault in
// Amazon Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/creating-vaults.html)
// and Create Vault  (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-put.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) CreateVault(ctx context.Context, params *CreateVaultInput, optFns ...func(*Options)) (*CreateVaultOutput, error) {
	if params == nil {
		params = &CreateVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVault", params, optFns, c.addOperationCreateVaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to create a vault.
type CreateVaultInput struct {

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You can
	// either specify an AWS account ID or optionally a single ' - ' (hyphen), in which
	// case Amazon S3 Glacier uses the AWS account ID associated with the credentials
	// used to sign the request. If you specify your account ID, do not include any
	// hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	noSmithyDocumentSerde
}

// Contains the Amazon S3 Glacier response to your request.
type CreateVaultOutput struct {

	// The URI of the vault that was created.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVault{}, middleware.After)
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
	if err = addOpCreateVaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVault(options.Region), middleware.Before); err != nil {
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
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "CreateVault",
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

func newServiceMetadataMiddleware_opCreateVault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "CreateVault",
	}
}
