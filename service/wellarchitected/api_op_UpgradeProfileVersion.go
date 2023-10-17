// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Upgrade a profile.
func (c *Client) UpgradeProfileVersion(ctx context.Context, params *UpgradeProfileVersionInput, optFns ...func(*Options)) (*UpgradeProfileVersionOutput, error) {
	if params == nil {
		params = &UpgradeProfileVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpgradeProfileVersion", params, optFns, c.addOperationUpgradeProfileVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpgradeProfileVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpgradeProfileVersionInput struct {

	// The profile ARN.
	//
	// This member is required.
	ProfileArn *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once). You should not reuse the same token for other requests. If
	// you retry a request with the same client request token and the same parameters
	// after the original request has completed successfully, the result of the
	// original request is returned. This token is listed as required, however, if you
	// do not specify it, the Amazon Web Services SDKs automatically generate one for
	// you. If you are not using the Amazon Web Services SDK or the CLI, you must
	// provide this token or the request will fail.
	ClientRequestToken *string

	// The name of the milestone in a workload. Milestone names must be unique within
	// a workload.
	MilestoneName *string

	noSmithyDocumentSerde
}

type UpgradeProfileVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpgradeProfileVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpgradeProfileVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpgradeProfileVersion{}, middleware.After)
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
	if err = addIdempotencyToken_opUpgradeProfileVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpgradeProfileVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpgradeProfileVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpgradeProfileVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpgradeProfileVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpgradeProfileVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpgradeProfileVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpgradeProfileVersionInput ")
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
func addIdempotencyToken_opUpgradeProfileVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpgradeProfileVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpgradeProfileVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "UpgradeProfileVersion",
	}
}
