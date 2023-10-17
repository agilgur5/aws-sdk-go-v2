// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a previously created parallel data resource by importing a new input
// file from Amazon S3.
func (c *Client) UpdateParallelData(ctx context.Context, params *UpdateParallelDataInput, optFns ...func(*Options)) (*UpdateParallelDataOutput, error) {
	if params == nil {
		params = &UpdateParallelDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateParallelData", params, optFns, c.addOperationUpdateParallelDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateParallelDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateParallelDataInput struct {

	// A unique identifier for the request. This token is automatically generated when
	// you use Amazon Translate through an AWS SDK.
	//
	// This member is required.
	ClientToken *string

	// The name of the parallel data resource being updated.
	//
	// This member is required.
	Name *string

	// Specifies the format and S3 location of the parallel data input file.
	//
	// This member is required.
	ParallelDataConfig *types.ParallelDataConfig

	// A custom description for the parallel data resource in Amazon Translate.
	Description *string

	noSmithyDocumentSerde
}

type UpdateParallelDataOutput struct {

	// The time that the most recent update was attempted.
	LatestUpdateAttemptAt *time.Time

	// The status of the parallel data update attempt. When the updated parallel data
	// resource is ready for you to use, the status is ACTIVE .
	LatestUpdateAttemptStatus types.ParallelDataStatus

	// The name of the parallel data resource being updated.
	Name *string

	// The status of the parallel data resource that you are attempting to update.
	// Your update request is accepted only if this status is either ACTIVE or FAILED .
	Status types.ParallelDataStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateParallelDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateParallelData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateParallelData{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateParallelDataMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateParallelDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateParallelData(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateParallelData struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateParallelData) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateParallelData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateParallelDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateParallelDataInput ")
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
func addIdempotencyToken_opUpdateParallelDataMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateParallelData{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateParallelData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "UpdateParallelData",
	}
}
