// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this action to disassociate, or remove, one or more Domain Name Service
// (DNS) aliases from an Amazon FSx for Windows File Server file system. If you
// attempt to disassociate a DNS alias that is not associated with the file system,
// Amazon FSx responds with a 400 Bad Request. For more information, see Working
// with DNS Aliases (https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-dns-aliases.html)
// . The system generated response showing the DNS aliases that Amazon FSx is
// attempting to disassociate from the file system. Use the API operation to
// monitor the status of the aliases Amazon FSx is disassociating with the file
// system.
func (c *Client) DisassociateFileSystemAliases(ctx context.Context, params *DisassociateFileSystemAliasesInput, optFns ...func(*Options)) (*DisassociateFileSystemAliasesOutput, error) {
	if params == nil {
		params = &DisassociateFileSystemAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateFileSystemAliases", params, optFns, c.addOperationDisassociateFileSystemAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateFileSystemAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object of DNS aliases to disassociate from an Amazon FSx for
// Windows File Server file system.
type DisassociateFileSystemAliasesInput struct {

	// An array of one or more DNS alias names to disassociate, or remove, from the
	// file system.
	//
	// This member is required.
	Aliases []string

	// Specifies the file system from which to disassociate the DNS aliases.
	//
	// This member is required.
	FileSystemId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

// The system generated response showing the DNS aliases that Amazon FSx is
// attempting to disassociate from the file system. Use the API operation to
// monitor the status of the aliases Amazon FSx is removing from the file system.
type DisassociateFileSystemAliasesOutput struct {

	// An array of one or more DNS aliases that Amazon FSx is attempting to
	// disassociate from the file system.
	Aliases []types.Alias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateFileSystemAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateFileSystemAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateFileSystemAliases{}, middleware.After)
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
	if err = addIdempotencyToken_opDisassociateFileSystemAliasesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDisassociateFileSystemAliasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateFileSystemAliases(options.Region), middleware.Before); err != nil {
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
		operation: "DisassociateFileSystemAliases",
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

type idempotencyToken_initializeOpDisassociateFileSystemAliases struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDisassociateFileSystemAliases) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDisassociateFileSystemAliases) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DisassociateFileSystemAliasesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DisassociateFileSystemAliasesInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDisassociateFileSystemAliasesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDisassociateFileSystemAliases{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDisassociateFileSystemAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DisassociateFileSystemAliases",
	}
}
