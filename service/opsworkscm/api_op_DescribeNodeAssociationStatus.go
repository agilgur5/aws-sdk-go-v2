// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns the current status of an existing association or disassociation
// request. A ResourceNotFoundException is thrown when no recent association or
// disassociation request with the specified token is found, or when the server
// does not exist. A ValidationException is raised when parameters of the request
// are not valid.
func (c *Client) DescribeNodeAssociationStatus(ctx context.Context, params *DescribeNodeAssociationStatusInput, optFns ...func(*Options)) (*DescribeNodeAssociationStatusOutput, error) {
	if params == nil {
		params = &DescribeNodeAssociationStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNodeAssociationStatus", params, optFns, c.addOperationDescribeNodeAssociationStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNodeAssociationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNodeAssociationStatusInput struct {

	// The token returned in either the AssociateNodeResponse or the
	// DisassociateNodeResponse.
	//
	// This member is required.
	NodeAssociationStatusToken *string

	// The name of the server from which to disassociate the node.
	//
	// This member is required.
	ServerName *string

	noSmithyDocumentSerde
}

func (*DescribeNodeAssociationStatusInput) operationName() string {
	return "DescribeNodeAssociationStatus"
}

type DescribeNodeAssociationStatusOutput struct {

	// Attributes specific to the node association. In Puppet, the attibute
	// PUPPET_NODE_CERT contains the signed certificate (the result of the CSR).
	EngineAttributes []types.EngineAttribute

	// The status of the association or disassociation request. Possible values:
	//   - SUCCESS : The association or disassociation succeeded.
	//   - FAILED : The association or disassociation failed.
	//   - IN_PROGRESS : The association or disassociation is still in progress.
	NodeAssociationStatus types.NodeAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNodeAssociationStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNodeAssociationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNodeAssociationStatus{}, middleware.After)
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
	if err = addDescribeNodeAssociationStatusResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeNodeAssociationStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNodeAssociationStatus(options.Region), middleware.Before); err != nil {
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

// DescribeNodeAssociationStatusAPIClient is a client that implements the
// DescribeNodeAssociationStatus operation.
type DescribeNodeAssociationStatusAPIClient interface {
	DescribeNodeAssociationStatus(context.Context, *DescribeNodeAssociationStatusInput, ...func(*Options)) (*DescribeNodeAssociationStatusOutput, error)
}

var _ DescribeNodeAssociationStatusAPIClient = (*Client)(nil)

// NodeAssociatedWaiterOptions are waiter options for NodeAssociatedWaiter
type NodeAssociatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// NodeAssociatedWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, NodeAssociatedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeNodeAssociationStatusInput, *DescribeNodeAssociationStatusOutput, error) (bool, error)
}

// NodeAssociatedWaiter defines the waiters for NodeAssociated
type NodeAssociatedWaiter struct {
	client DescribeNodeAssociationStatusAPIClient

	options NodeAssociatedWaiterOptions
}

// NewNodeAssociatedWaiter constructs a NodeAssociatedWaiter.
func NewNodeAssociatedWaiter(client DescribeNodeAssociationStatusAPIClient, optFns ...func(*NodeAssociatedWaiterOptions)) *NodeAssociatedWaiter {
	options := NodeAssociatedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = nodeAssociatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &NodeAssociatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for NodeAssociated waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *NodeAssociatedWaiter) Wait(ctx context.Context, params *DescribeNodeAssociationStatusInput, maxWaitDur time.Duration, optFns ...func(*NodeAssociatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for NodeAssociated waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *NodeAssociatedWaiter) WaitForOutput(ctx context.Context, params *DescribeNodeAssociationStatusInput, maxWaitDur time.Duration, optFns ...func(*NodeAssociatedWaiterOptions)) (*DescribeNodeAssociationStatusOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
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

		out, err := w.client.DescribeNodeAssociationStatus(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for NodeAssociated waiter")
}

func nodeAssociatedStateRetryable(ctx context.Context, input *DescribeNodeAssociationStatusInput, output *DescribeNodeAssociationStatusOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("NodeAssociationStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUCCESS"
		value, ok := pathValue.(types.NodeAssociationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.NodeAssociationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("NodeAssociationStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.NodeAssociationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.NodeAssociationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeNodeAssociationStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "DescribeNodeAssociationStatus",
	}
}

type opDescribeNodeAssociationStatusResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeNodeAssociationStatusResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeNodeAssociationStatusResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "opsworks-cm"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "opsworks-cm"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("opsworks-cm")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeNodeAssociationStatusResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeNodeAssociationStatusResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
