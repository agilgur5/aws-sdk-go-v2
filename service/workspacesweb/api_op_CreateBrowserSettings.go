// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a browser settings resource that can be associated with a web portal.
// Once associated with a web portal, browser settings control how the browser will
// behave once a user starts a streaming session for the web portal.
func (c *Client) CreateBrowserSettings(ctx context.Context, params *CreateBrowserSettingsInput, optFns ...func(*Options)) (*CreateBrowserSettingsOutput, error) {
	if params == nil {
		params = &CreateBrowserSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBrowserSettings", params, optFns, c.addOperationCreateBrowserSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBrowserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBrowserSettingsInput struct {

	// A JSON string containing Chrome Enterprise policies that will be applied to all
	// streaming sessions.
	//
	// This member is required.
	BrowserPolicy *string

	// Additional encryption context of the browser settings.
	AdditionalEncryptionContext map[string]string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. Idempotency ensures that an API request completes only once.
	// With an idempotent request, if the original request completes successfully,
	// subsequent retries with the same client token returns the result from the
	// original successful request. If you do not specify a client token, one is
	// automatically generated by the AWS SDK.
	ClientToken *string

	// The custom managed key of the browser settings.
	CustomerManagedKey *string

	// The tags to add to the browser settings resource. A tag is a key-value pair.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateBrowserSettingsOutput struct {

	// The ARN of the browser settings.
	//
	// This member is required.
	BrowserSettingsArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBrowserSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBrowserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBrowserSettings{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateBrowserSettingsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBrowserSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBrowserSettings(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateBrowserSettings struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateBrowserSettings) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateBrowserSettings) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateBrowserSettingsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateBrowserSettingsInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateBrowserSettingsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateBrowserSettings{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateBrowserSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces-web",
		OperationName: "CreateBrowserSettings",
	}
}
