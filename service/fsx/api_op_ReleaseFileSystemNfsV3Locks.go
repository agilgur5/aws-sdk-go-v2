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

// Releases the file system lock from an Amazon FSx for OpenZFS file system.
func (c *Client) ReleaseFileSystemNfsV3Locks(ctx context.Context, params *ReleaseFileSystemNfsV3LocksInput, optFns ...func(*Options)) (*ReleaseFileSystemNfsV3LocksOutput, error) {
	if params == nil {
		params = &ReleaseFileSystemNfsV3LocksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReleaseFileSystemNfsV3Locks", params, optFns, c.addOperationReleaseFileSystemNfsV3LocksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReleaseFileSystemNfsV3LocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReleaseFileSystemNfsV3LocksInput struct {

	// The globally unique ID of the file system, assigned by Amazon FSx.
	//
	// This member is required.
	FileSystemId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type ReleaseFileSystemNfsV3LocksOutput struct {

	// A description of a specific Amazon FSx file system.
	FileSystem *types.FileSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReleaseFileSystemNfsV3LocksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReleaseFileSystemNfsV3Locks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReleaseFileSystemNfsV3Locks{}, middleware.After)
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
	if err = addIdempotencyToken_opReleaseFileSystemNfsV3LocksMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpReleaseFileSystemNfsV3LocksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReleaseFileSystemNfsV3Locks(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpReleaseFileSystemNfsV3Locks struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpReleaseFileSystemNfsV3Locks) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpReleaseFileSystemNfsV3Locks) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ReleaseFileSystemNfsV3LocksInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ReleaseFileSystemNfsV3LocksInput ")
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
func addIdempotencyToken_opReleaseFileSystemNfsV3LocksMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpReleaseFileSystemNfsV3Locks{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opReleaseFileSystemNfsV3Locks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "ReleaseFileSystemNfsV3Locks",
	}
}
