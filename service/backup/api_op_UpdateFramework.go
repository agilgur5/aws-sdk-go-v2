// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an existing framework identified by its FrameworkName with the input
// document in JSON format.
func (c *Client) UpdateFramework(ctx context.Context, params *UpdateFrameworkInput, optFns ...func(*Options)) (*UpdateFrameworkOutput, error) {
	if params == nil {
		params = &UpdateFrameworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFramework", params, optFns, c.addOperationUpdateFrameworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFrameworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFrameworkInput struct {

	// The unique name of a framework. This name is between 1 and 256 characters,
	// starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and
	// underscores (_).
	//
	// This member is required.
	FrameworkName *string

	// A list of the controls that make up the framework. Each control in the list has
	// a name, input parameters, and scope.
	FrameworkControls []types.FrameworkControl

	// An optional description of the framework with a maximum 1,024 characters.
	FrameworkDescription *string

	// A customer-chosen string that you can use to distinguish between otherwise
	// identical calls to UpdateFrameworkInput . Retrying a successful request with the
	// same idempotency token results in a success message with no action taken.
	IdempotencyToken *string

	noSmithyDocumentSerde
}

type UpdateFrameworkOutput struct {

	// The date and time that a framework is created, in ISO 8601 representation. The
	// value of CreationTime is accurate to milliseconds. For example,
	// 2020-07-10T15:00:00.000-08:00 represents the 10th of July 2020 at 3:00 PM 8
	// hours behind UTC.
	CreationTime *time.Time

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format
	// of the ARN depends on the resource type.
	FrameworkArn *string

	// The unique name of a framework. This name is between 1 and 256 characters,
	// starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and
	// underscores (_).
	FrameworkName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFrameworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFramework{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFramework{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateFrameworkMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateFrameworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFramework(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateFramework struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateFramework) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateFramework) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateFrameworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateFrameworkInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateFrameworkMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateFramework{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateFramework(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "UpdateFramework",
	}
}
