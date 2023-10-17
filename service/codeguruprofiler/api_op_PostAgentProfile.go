// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits profiling data to an aggregated profile of a profiling group. To get an
// aggregated profile that is created with this profiling data, use GetProfile (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_GetProfile.html)
// .
func (c *Client) PostAgentProfile(ctx context.Context, params *PostAgentProfileInput, optFns ...func(*Options)) (*PostAgentProfileOutput, error) {
	if params == nil {
		params = &PostAgentProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PostAgentProfile", params, optFns, c.addOperationPostAgentProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PostAgentProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the postAgentProfileRequest.
type PostAgentProfileInput struct {

	// The submitted profiling data.
	//
	// This member is required.
	AgentProfile []byte

	// The format of the submitted profiling data. The format maps to the Accept and
	// Content-Type headers of the HTTP request. You can specify one of the following:
	// or the default .
	//   - application/json — standard JSON format
	//   - application/x-amzn-ion — the Amazon Ion data format. For more information,
	//   see Amazon Ion (http://amzn.github.io/ion-docs/) .
	//
	// This member is required.
	ContentType *string

	// The name of the profiling group with the aggregated profile that receives the
	// submitted profiling data.
	//
	// This member is required.
	ProfilingGroupName *string

	// Amazon CodeGuru Profiler uses this universally unique identifier (UUID) to
	// prevent the accidental submission of duplicate profiling data if there are
	// failures and retries.
	ProfileToken *string

	noSmithyDocumentSerde
}

// The structure representing the postAgentProfileResponse.
type PostAgentProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPostAgentProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPostAgentProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPostAgentProfile{}, middleware.After)
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
	if err = addIdempotencyToken_opPostAgentProfileMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPostAgentProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPostAgentProfile(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPostAgentProfile struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPostAgentProfile) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPostAgentProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PostAgentProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PostAgentProfileInput ")
	}

	if input.ProfileToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ProfileToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opPostAgentProfileMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPostAgentProfile{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPostAgentProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "PostAgentProfile",
	}
}
