// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a set of frequently ask questions (FAQs) using a specified FAQ file
// stored in an Amazon S3 bucket. Adding FAQs to an index is an asynchronous
// operation. For an example of adding an FAQ to an index using Python and Java
// SDKs, see Using your FAQ file (https://docs.aws.amazon.com/kendra/latest/dg/in-creating-faq.html#using-faq-file)
// .
func (c *Client) CreateFaq(ctx context.Context, params *CreateFaqInput, optFns ...func(*Options)) (*CreateFaqOutput, error) {
	if params == nil {
		params = &CreateFaqInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFaq", params, optFns, c.addOperationCreateFaqMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFaqOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFaqInput struct {

	// The identifier of the index for the FAQ.
	//
	// This member is required.
	IndexId *string

	// A name for the FAQ.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of an IAM role with permission to access the S3
	// bucket that contains the FAQs. For more information, see IAM access roles for
	// Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html) .
	//
	// This member is required.
	RoleArn *string

	// The path to the FAQ file in S3.
	//
	// This member is required.
	S3Path *types.S3Path

	// A token that you provide to identify the request to create a FAQ. Multiple
	// calls to the CreateFaqRequest API with the same client token will create only
	// one FAQ.
	ClientToken *string

	// A description for the FAQ.
	Description *string

	// The format of the FAQ input file. You can choose between a basic CSV format, a
	// CSV format that includes customs attributes in a header, and a JSON format that
	// includes custom attributes. The default format is CSV. The format must match the
	// format of the file stored in the S3 bucket identified in the S3Path parameter.
	// For more information, see Adding questions and answers (https://docs.aws.amazon.com/kendra/latest/dg/in-creating-faq.html)
	// .
	FileFormat types.FaqFileFormat

	// The code for a language. This allows you to support a language for the FAQ
	// document. English is supported by default. For more information on supported
	// languages, including their codes, see Adding documents in languages other than
	// English (https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html) .
	LanguageCode *string

	// A list of key-value pairs that identify the FAQ. You can use the tags to
	// identify and organize your resources and to control access to resources.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFaqOutput struct {

	// The identifier of the FAQ.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFaqMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFaq{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFaq{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateFaqMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFaqValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFaq(options.Region), middleware.Before); err != nil {
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
		operation: "CreateFaq",
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

type idempotencyToken_initializeOpCreateFaq struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFaq) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFaq) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFaqInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFaqInput ")
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
func addIdempotencyToken_opCreateFaqMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFaq{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFaq(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateFaq",
	}
}
