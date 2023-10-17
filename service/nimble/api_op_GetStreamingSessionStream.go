// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets a StreamingSessionStream for a streaming session. Invoke this operation to
// poll the resource after invoking CreateStreamingSessionStream . After the
// StreamingSessionStream changes to the READY state, the url property will
// contain a stream to be used with the DCV streaming client.
func (c *Client) GetStreamingSessionStream(ctx context.Context, params *GetStreamingSessionStreamInput, optFns ...func(*Options)) (*GetStreamingSessionStreamOutput, error) {
	if params == nil {
		params = &GetStreamingSessionStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStreamingSessionStream", params, optFns, c.addOperationGetStreamingSessionStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStreamingSessionStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStreamingSessionStreamInput struct {

	// The streaming session ID.
	//
	// This member is required.
	SessionId *string

	// The streaming session stream ID.
	//
	// This member is required.
	StreamId *string

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	noSmithyDocumentSerde
}

type GetStreamingSessionStreamOutput struct {

	// The stream.
	Stream *types.StreamingSessionStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStreamingSessionStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetStreamingSessionStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStreamingSessionStream{}, middleware.After)
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
	if err = addOpGetStreamingSessionStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStreamingSessionStream(options.Region), middleware.Before); err != nil {
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

// GetStreamingSessionStreamAPIClient is a client that implements the
// GetStreamingSessionStream operation.
type GetStreamingSessionStreamAPIClient interface {
	GetStreamingSessionStream(context.Context, *GetStreamingSessionStreamInput, ...func(*Options)) (*GetStreamingSessionStreamOutput, error)
}

var _ GetStreamingSessionStreamAPIClient = (*Client)(nil)

// StreamingSessionStreamReadyWaiterOptions are waiter options for
// StreamingSessionStreamReadyWaiter
type StreamingSessionStreamReadyWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// StreamingSessionStreamReadyWaiter will use default minimum delay of 5 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, StreamingSessionStreamReadyWaiter will use default max delay of 150
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetStreamingSessionStreamInput, *GetStreamingSessionStreamOutput, error) (bool, error)
}

// StreamingSessionStreamReadyWaiter defines the waiters for
// StreamingSessionStreamReady
type StreamingSessionStreamReadyWaiter struct {
	client GetStreamingSessionStreamAPIClient

	options StreamingSessionStreamReadyWaiterOptions
}

// NewStreamingSessionStreamReadyWaiter constructs a
// StreamingSessionStreamReadyWaiter.
func NewStreamingSessionStreamReadyWaiter(client GetStreamingSessionStreamAPIClient, optFns ...func(*StreamingSessionStreamReadyWaiterOptions)) *StreamingSessionStreamReadyWaiter {
	options := StreamingSessionStreamReadyWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 150 * time.Second
	options.Retryable = streamingSessionStreamReadyStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &StreamingSessionStreamReadyWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for StreamingSessionStreamReady waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *StreamingSessionStreamReadyWaiter) Wait(ctx context.Context, params *GetStreamingSessionStreamInput, maxWaitDur time.Duration, optFns ...func(*StreamingSessionStreamReadyWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for StreamingSessionStreamReady waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *StreamingSessionStreamReadyWaiter) WaitForOutput(ctx context.Context, params *GetStreamingSessionStreamInput, maxWaitDur time.Duration, optFns ...func(*StreamingSessionStreamReadyWaiterOptions)) (*GetStreamingSessionStreamOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 150 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetStreamingSessionStream(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for StreamingSessionStreamReady waiter")
}

func streamingSessionStreamReadyStateRetryable(ctx context.Context, input *GetStreamingSessionStreamInput, output *GetStreamingSessionStreamOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("stream.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "READY"
		value, ok := pathValue.(types.StreamingSessionStreamState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StreamingSessionStreamState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("stream.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED"
		value, ok := pathValue.(types.StreamingSessionStreamState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StreamingSessionStreamState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetStreamingSessionStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "GetStreamingSessionStream",
	}
}
