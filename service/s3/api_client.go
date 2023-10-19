// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	internalauthsmithy "github.com/aws/aws-sdk-go-v2/internal/auth/smithy"
	internalConfig "github.com/aws/aws-sdk-go-v2/internal/configsources"
	"github.com/aws/aws-sdk-go-v2/internal/v4a"
	acceptencodingcust "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared"
	s3sharedconfig "github.com/aws/aws-sdk-go-v2/service/internal/s3shared/config"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	smithy "github.com/aws/smithy-go"
	smithyauth "github.com/aws/smithy-go/auth"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net"
	"net/http"
	"time"
)

const ServiceID = "S3"
const ServiceAPIVersion = "2006-03-01"

// Client provides the API client to make operations call for Amazon Simple
// Storage Service.
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveDefaultLogger(&options)

	setResolvedDefaultsMode(&options)

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveHTTPSignerV4a(&options)

	resolveAuthSchemeResolver(&options)

	resolveAuthSchemes(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	resolveCredentialProvider(&options)

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// The optional application specific identifier appended to the User-Agent header.
	AppID string

	// This endpoint will be given as input to an EndpointResolverV2. It is used for
	// providing a custom base endpoint that is subject to modifications by the
	// processing EndpointResolverV2.
	BaseEndpoint *string

	// Configures the events that will be sent to the configured logger.
	ClientLogMode aws.ClientLogMode

	// The threshold ContentLength in bytes for HTTP PUT request to receive {Expect:
	// 100-continue} header. Setting to -1 will disable adding the Expect header to
	// requests; setting to 0 will set the threshold to default 2MB
	ContinueHeaderThresholdBytes int64

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The configuration DefaultsMode that the SDK should use when constructing the
	// clients initial default settings.
	DefaultsMode aws.DefaultsMode

	// Allows you to disable S3 Multi-Region access points feature.
	DisableMultiRegionAccessPoints bool

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions EndpointResolverOptions

	// The service endpoint resolver.
	//
	// Deprecated: Deprecated: EndpointResolver and WithEndpointResolver. Providing a
	// value for this field will likely prevent you from using any endpoint-related
	// service features released after the introduction of EndpointResolverV2 and
	// BaseEndpoint. To migrate an EndpointResolver implementation that uses a custom
	// endpoint, set the client option BaseEndpoint instead.
	EndpointResolver EndpointResolver

	// Resolves the endpoint used for a particular service. This should be used over
	// the deprecated EndpointResolver
	EndpointResolverV2 EndpointResolverV2

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The logger writer interface to write logging messages to.
	Logger logging.Logger

	// The region to send requests to. (Required)
	Region string

	// RetryMaxAttempts specifies the maximum number attempts an API client will call
	// an operation that fails with a retryable error. A value of 0 is ignored, and
	// will not be used to configure the API client created default retryer, or modify
	// per operation call's retry max attempts. When creating a new API Clients this
	// member will only be used if the Retryer Options member is nil. This value will
	// be ignored if Retryer is not nil. If specified in an operation call's functional
	// options with a value that is different than the constructed client's Options,
	// the Client's Retryer will be wrapped to use the operation's specific
	// RetryMaxAttempts value.
	RetryMaxAttempts int

	// RetryMode specifies the retry mode the API client will be created with, if
	// Retryer option is not also specified. When creating a new API Clients this
	// member will only be used if the Retryer Options member is nil. This value will
	// be ignored if Retryer is not nil. Currently does not support per operation call
	// overrides, may in the future.
	RetryMode aws.RetryMode

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer. The kind of
	// default retry created by the API client can be changed with the RetryMode
	// option.
	Retryer aws.Retryer

	// The RuntimeEnvironment configuration, only populated if the DefaultsMode is set
	// to DefaultsModeAuto and is initialized using config.LoadDefaultConfig . You
	// should not populate this structure programmatically, or rely on the values here
	// within your applications.
	RuntimeEnvironment aws.RuntimeEnvironment

	// Allows you to enable arn region support for the service.
	UseARNRegion bool

	// Allows you to enable S3 Accelerate feature. All operations compatible with S3
	// Accelerate will use the accelerate endpoint for requests. Requests not
	// compatible will fall back to normal S3 requests. The bucket must be enabled for
	// accelerate to be used with S3 client with accelerate enabled. If the bucket is
	// not enabled for accelerate an error will be returned. The bucket name must be
	// DNS compatible to work with accelerate.
	UseAccelerate bool

	// Allows you to enable dual-stack endpoint support for the service.
	//
	// Deprecated: Set dual-stack by setting UseDualStackEndpoint on
	// EndpointResolverOptions. When EndpointResolverOptions' UseDualStackEndpoint
	// field is set it overrides this field value.
	UseDualstack bool

	// Allows you to enable the client to use path-style addressing, i.e.,
	// https://s3.amazonaws.com/BUCKET/KEY . By default, the S3 client will use virtual
	// hosted bucket addressing when possible( https://BUCKET.s3.amazonaws.com/KEY ).
	UsePathStyle bool

	// Signature Version 4a (SigV4a) Signer
	httpSignerV4a httpSignerV4a

	// The initial DefaultsMode used when the client options were constructed. If the
	// DefaultsMode was set to aws.DefaultsModeAuto this will store what the resolved
	// value was at that point in time. Currently does not support per operation call
	// overrides, may in the future.
	resolvedDefaultsMode aws.DefaultsMode

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient

	// The auth scheme resolver which determines how to authenticate for each
	// operation.
	AuthSchemeResolver AuthSchemeResolver

	// The list of auth schemes supported by the client.
	AuthSchemes []smithyhttp.AuthScheme
}

func (o Options) GetIdentityResolver(schemeID string) smithyauth.IdentityResolver {
	if schemeID == "aws.auth#sigv4" {
		return &internalauthsmithy.CredentialsProviderAdapter{Provider: o.Credentials}
	}
	return nil
}

// WithAPIOptions returns a functional option for setting the Client's APIOptions
// option.
func WithAPIOptions(optFns ...func(*middleware.Stack) error) func(*Options) {
	return func(o *Options) {
		o.APIOptions = append(o.APIOptions, optFns...)
	}
}

// Deprecated: EndpointResolver and WithEndpointResolver. Providing a value for
// this field will likely prevent you from using any endpoint-related service
// features released after the introduction of EndpointResolverV2 and BaseEndpoint.
// To migrate an EndpointResolver implementation that uses a custom endpoint, set
// the client option BaseEndpoint instead.
func WithEndpointResolver(v EndpointResolver) func(*Options) {
	return func(o *Options) {
		o.EndpointResolver = v
	}
}

// WithEndpointResolverV2 returns a functional option for setting the Client's
// EndpointResolverV2 option.
func WithEndpointResolverV2(v EndpointResolverV2) func(*Options) {
	return func(o *Options) {
		o.EndpointResolverV2 = v
	}
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)

	return to
}
func (c *Client) invokeOperation(ctx context.Context, opID string, params interface{}, optFns []func(*Options), stackFns ...func(*middleware.Stack, Options) error) (result interface{}, metadata middleware.Metadata, err error) {
	ctx = middleware.ClearStackValues(ctx)
	stack := middleware.NewStack(opID, smithyhttp.NewStackRequest)
	options := c.options.Copy()
	resolveEndpointResolverV2(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	setSafeEventStreamClientLogMode(&options, opID)

	finalizeRetryMaxAttemptOptions(&options, *c)

	finalizeClientEndpointResolverOptions(&options)

	resolveCredentialProvider(&options)

	for _, fn := range stackFns {
		if err := fn(stack, options); err != nil {
			return nil, metadata, err
		}
	}

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, metadata, err
		}
	}

	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err = handler.Handle(ctx, params)
	if err != nil {
		err = &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: opID,
			Err:           err,
		}
	}
	return result, metadata, err
}
func resolveAuthSchemeResolver(options *Options) {
	options.AuthSchemeResolver = &defaultAuthSchemeResolver{}
}

func resolveAuthSchemes(options *Options) {
	options.AuthSchemes = []smithyhttp.AuthScheme{
		internalauth.NewHTTPAuthScheme("aws.auth#sigv4", &internalauthsmithy.V4SignerAdapter{Signer: options.HTTPSignerV4}),
	}
}

type noSmithyDocumentSerde = smithydocument.NoSerde

type legacyEndpointContextSetter struct {
	LegacyResolver EndpointResolver
}

func (*legacyEndpointContextSetter) ID() string {
	return "legacyEndpointContextSetter"
}

func (m *legacyEndpointContextSetter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.LegacyResolver != nil {
		ctx = awsmiddleware.SetRequiresLegacyEndpoints(ctx, true)
	}

	return next.HandleInitialize(ctx, in)

}
func addlegacyEndpointContextSetter(stack *middleware.Stack, o Options) error {
	return stack.Initialize.Add(&legacyEndpointContextSetter{
		LegacyResolver: o.EndpointResolver,
	}, middleware.Before)
}

func resolveDefaultLogger(o *Options) {
	if o.Logger != nil {
		return
	}
	o.Logger = logging.Nop{}
}

func addSetLoggerMiddleware(stack *middleware.Stack, o Options) error {
	return middleware.AddSetLoggerMiddleware(stack, o.Logger)
}

func setResolvedDefaultsMode(o *Options) {
	if len(o.resolvedDefaultsMode) > 0 {
		return
	}

	var mode aws.DefaultsMode
	mode.SetFromString(string(o.DefaultsMode))

	if mode == aws.DefaultsModeAuto {
		mode = defaults.ResolveDefaultsModeAuto(o.Region, o.RuntimeEnvironment)
	}

	o.resolvedDefaultsMode = mode
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:             cfg.Region,
		DefaultsMode:       cfg.DefaultsMode,
		RuntimeEnvironment: cfg.RuntimeEnvironment,
		HTTPClient:         cfg.HTTPClient,
		Credentials:        cfg.Credentials,
		APIOptions:         cfg.APIOptions,
		Logger:             cfg.Logger,
		ClientLogMode:      cfg.ClientLogMode,
		AppID:              cfg.AppID,
	}
	resolveAWSRetryerProvider(cfg, &opts)
	resolveAWSRetryMaxAttempts(cfg, &opts)
	resolveAWSRetryMode(cfg, &opts)
	resolveAWSEndpointResolver(cfg, &opts)
	resolveUseARNRegion(cfg, &opts)
	resolveDisableMultiRegionAccessPoints(cfg, &opts)
	resolveUseDualStackEndpoint(cfg, &opts)
	resolveUseFIPSEndpoint(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	var buildable *awshttp.BuildableClient

	if o.HTTPClient != nil {
		var ok bool
		buildable, ok = o.HTTPClient.(*awshttp.BuildableClient)
		if !ok {
			return
		}
	} else {
		buildable = awshttp.NewBuildableClient()
	}

	modeConfig, err := defaults.GetModeConfiguration(o.resolvedDefaultsMode)
	if err == nil {
		buildable = buildable.WithDialerOptions(func(dialer *net.Dialer) {
			if dialerTimeout, ok := modeConfig.GetConnectTimeout(); ok {
				dialer.Timeout = dialerTimeout
			}
		})

		buildable = buildable.WithTransportOptions(func(transport *http.Transport) {
			if tlsHandshakeTimeout, ok := modeConfig.GetTLSNegotiationTimeout(); ok {
				transport.TLSHandshakeTimeout = tlsHandshakeTimeout
			}
		})
	}

	o.HTTPClient = buildable
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}

	if len(o.RetryMode) == 0 {
		modeConfig, err := defaults.GetModeConfiguration(o.resolvedDefaultsMode)
		if err == nil {
			o.RetryMode = modeConfig.RetryMode
		}
	}
	if len(o.RetryMode) == 0 {
		o.RetryMode = aws.RetryModeStandard
	}

	var standardOptions []func(*retry.StandardOptions)
	if v := o.RetryMaxAttempts; v != 0 {
		standardOptions = append(standardOptions, func(so *retry.StandardOptions) {
			so.MaxAttempts = v
		})
	}

	switch o.RetryMode {
	case aws.RetryModeAdaptive:
		var adaptiveOptions []func(*retry.AdaptiveModeOptions)
		if len(standardOptions) != 0 {
			adaptiveOptions = append(adaptiveOptions, func(ao *retry.AdaptiveModeOptions) {
				ao.StandardOptions = append(ao.StandardOptions, standardOptions...)
			})
		}
		o.Retryer = retry.NewAdaptiveMode(adaptiveOptions...)

	default:
		o.Retryer = retry.NewStandard(standardOptions...)
	}
}

func resolveAWSRetryerProvider(cfg aws.Config, o *Options) {
	if cfg.Retryer == nil {
		return
	}
	o.Retryer = cfg.Retryer()
}

func resolveAWSRetryMode(cfg aws.Config, o *Options) {
	if len(cfg.RetryMode) == 0 {
		return
	}
	o.RetryMode = cfg.RetryMode
}
func resolveAWSRetryMaxAttempts(cfg aws.Config, o *Options) {
	if cfg.RetryMaxAttempts == 0 {
		return
	}
	o.RetryMaxAttempts = cfg.RetryMaxAttempts
}

func finalizeRetryMaxAttemptOptions(o *Options, client Client) {
	if v := o.RetryMaxAttempts; v == 0 || v == client.options.RetryMaxAttempts {
		return
	}

	o.Retryer = retry.AddWithMaxAttempts(o.Retryer, o.RetryMaxAttempts)
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil && cfg.EndpointResolverWithOptions == nil {
		return
	}
	o.EndpointResolver = withEndpointResolver(cfg.EndpointResolver, cfg.EndpointResolverWithOptions)
}

func addClientUserAgent(stack *middleware.Stack, options Options) error {
	if err := awsmiddleware.AddSDKAgentKeyValue(awsmiddleware.APIMetadata, "s3", goModuleVersion)(stack); err != nil {
		return err
	}

	if len(options.AppID) > 0 {
		return awsmiddleware.AddSDKAgentKey(awsmiddleware.ApplicationIdentifier, options.AppID)(stack)
	}

	return nil
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time, optFns ...func(*v4.SignerOptions)) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = newDefaultV4Signer(*o)
}

func newDefaultV4Signer(o Options) *v4.Signer {
	return v4.NewSigner(func(so *v4.SignerOptions) {
		so.Logger = o.Logger
		so.LogSigning = o.ClientLogMode.IsSigning()
		so.DisableURIPathEscaping = true
	})
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer:          o.Retryer,
		LogRetryAttempts: o.ClientLogMode.IsRetries(),
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

// resolves UseARNRegion S3 configuration
func resolveUseARNRegion(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := s3sharedconfig.ResolveUseARNRegion(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.UseARNRegion = value
	}
	return nil
}

// resolves DisableMultiRegionAccessPoints S3 configuration
func resolveDisableMultiRegionAccessPoints(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := s3sharedconfig.ResolveDisableMultiRegionAccessPoints(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.DisableMultiRegionAccessPoints = value
	}
	return nil
}

// resolves dual-stack endpoint configuration
func resolveUseDualStackEndpoint(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := internalConfig.ResolveUseDualStackEndpoint(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.EndpointOptions.UseDualStackEndpoint = value
	}
	return nil
}

// resolves FIPS endpoint configuration
func resolveUseFIPSEndpoint(cfg aws.Config, o *Options) error {
	if len(cfg.ConfigSources) == 0 {
		return nil
	}
	value, found, err := internalConfig.ResolveUseFIPSEndpoint(context.Background(), cfg.ConfigSources)
	if err != nil {
		return err
	}
	if found {
		o.EndpointOptions.UseFIPSEndpoint = value
	}
	return nil
}

func resolveCredentialProvider(o *Options) {
	if o.Credentials == nil {
		return
	}

	if _, ok := o.Credentials.(v4a.CredentialsProvider); ok {
		return
	}

	if aws.IsCredentialsProvider(o.Credentials, (*aws.AnonymousCredentials)(nil)) {
		return
	}

	o.Credentials = &v4a.SymmetricCredentialAdaptor{SymmetricProvider: o.Credentials}
}

func swapWithCustomHTTPSignerMiddleware(stack *middleware.Stack, o Options) error {
	mw := s3cust.NewSignHTTPRequestMiddleware(s3cust.SignHTTPRequestMiddlewareOptions{
		CredentialsProvider: o.Credentials,
		V4Signer:            o.HTTPSignerV4,
		V4aSigner:           o.httpSignerV4a,
		LogSigning:          o.ClientLogMode.IsSigning(),
	})

	return s3cust.RegisterSigningMiddleware(stack, mw)
}

type httpSignerV4a interface {
	SignHTTP(ctx context.Context, credentials v4a.Credentials, r *http.Request, payloadHash,
		service string, regionSet []string, signingTime time.Time,
		optFns ...func(*v4a.SignerOptions)) error
}

func resolveHTTPSignerV4a(o *Options) {
	if o.httpSignerV4a != nil {
		return
	}
	o.httpSignerV4a = newDefaultV4aSigner(*o)
}

func newDefaultV4aSigner(o Options) *v4a.Signer {
	return v4a.NewSigner(func(so *v4a.SignerOptions) {
		so.Logger = o.Logger
		so.LogSigning = o.ClientLogMode.IsSigning()
		so.DisableURIPathEscaping = true
	})
}

func addMetadataRetrieverMiddleware(stack *middleware.Stack) error {
	return s3shared.AddMetadataRetrieverMiddleware(stack)
}

func add100Continue(stack *middleware.Stack, options Options) error {
	return s3shared.Add100Continue(stack, options.ContinueHeaderThresholdBytes)
}

// ComputedInputChecksumsMetadata provides information about the algorithms used
// to compute the checksum(s) of the input payload.
type ComputedInputChecksumsMetadata struct {
	// ComputedChecksums is a map of algorithm name to checksum value of the computed
	// input payload's checksums.
	ComputedChecksums map[string]string
}

// GetComputedInputChecksumsMetadata retrieves from the result metadata the map of
// algorithms and input payload checksums values.
func GetComputedInputChecksumsMetadata(m middleware.Metadata) (ComputedInputChecksumsMetadata, bool) {
	values, ok := internalChecksum.GetComputedInputChecksums(m)
	if !ok {
		return ComputedInputChecksumsMetadata{}, false
	}
	return ComputedInputChecksumsMetadata{
		ComputedChecksums: values,
	}, true

}

// ChecksumValidationMetadata contains metadata such as the checksum algorithm
// used for data integrity validation.
type ChecksumValidationMetadata struct {
	// AlgorithmsUsed is the set of the checksum algorithms used to validate the
	// response payload. The response payload must be completely read in order for the
	// checksum validation to be performed. An error is returned by the operation
	// output's response io.ReadCloser if the computed checksums are invalid.
	AlgorithmsUsed []string
}

// GetChecksumValidationMetadata returns the set of algorithms that will be used
// to validate the response payload with. The response payload must be completely
// read in order for the checksum validation to be performed. An error is returned
// by the operation output's response io.ReadCloser if the computed checksums are
// invalid. Returns false if no checksum algorithm used metadata was found.
func GetChecksumValidationMetadata(m middleware.Metadata) (ChecksumValidationMetadata, bool) {
	values, ok := internalChecksum.GetOutputValidationAlgorithmsUsed(m)
	if !ok {
		return ChecksumValidationMetadata{}, false
	}
	return ChecksumValidationMetadata{
		AlgorithmsUsed: append(make([]string, 0, len(values)), values...),
	}, true

}

// nopGetBucketAccessor is no-op accessor for operation that don't support bucket
// member as input
func nopGetBucketAccessor(input interface{}) (*string, bool) {
	return nil, false
}

func addResponseErrorMiddleware(stack *middleware.Stack) error {
	return s3shared.AddResponseErrorMiddleware(stack)
}

func disableAcceptEncodingGzip(stack *middleware.Stack) error {
	return acceptencodingcust.AddAcceptEncodingGzip(stack, acceptencodingcust.AddAcceptEncodingGzipOptions{})
}

// ResponseError provides the HTTP centric error type wrapping the underlying
// error with the HTTP response value and the deserialized RequestID.
type ResponseError interface {
	error

	ServiceHostID() string
	ServiceRequestID() string
}

var _ ResponseError = (*s3shared.ResponseError)(nil)

// GetHostIDMetadata retrieves the host id from middleware metadata returns host
// id as string along with a boolean indicating presence of hostId on middleware
// metadata.
func GetHostIDMetadata(metadata middleware.Metadata) (string, bool) {
	return s3shared.GetHostIDMetadata(metadata)
}

// HTTPPresignerV4 represents presigner interface used by presign url client
type HTTPPresignerV4 interface {
	PresignHTTP(
		ctx context.Context, credentials aws.Credentials, r *http.Request,
		payloadHash string, service string, region string, signingTime time.Time,
		optFns ...func(*v4.SignerOptions),
	) (url string, signedHeader http.Header, err error)
}

// httpPresignerV4a represents sigv4a presigner interface used by presign url
// client
type httpPresignerV4a interface {
	PresignHTTP(
		ctx context.Context, credentials v4a.Credentials, r *http.Request,
		payloadHash string, service string, regionSet []string, signingTime time.Time,
		optFns ...func(*v4a.SignerOptions),
	) (url string, signedHeader http.Header, err error)
}

// PresignOptions represents the presign client options
type PresignOptions struct {

	// ClientOptions are list of functional options to mutate client options used by
	// the presign client.
	ClientOptions []func(*Options)

	// Presigner is the presigner used by the presign url client
	Presigner HTTPPresignerV4

	// Expires sets the expiration duration for the generated presign url. This should
	// be the duration in seconds the presigned URL should be considered valid for. If
	// not set or set to zero, presign url would default to expire after 900 seconds.
	Expires time.Duration

	// presignerV4a is the presigner used by the presign url client
	presignerV4a httpPresignerV4a
}

func (o PresignOptions) copy() PresignOptions {
	clientOptions := make([]func(*Options), len(o.ClientOptions))
	copy(clientOptions, o.ClientOptions)
	o.ClientOptions = clientOptions
	return o
}

// WithPresignClientFromClientOptions is a helper utility to retrieve a function
// that takes PresignOption as input
func WithPresignClientFromClientOptions(optFns ...func(*Options)) func(*PresignOptions) {
	return withPresignClientFromClientOptions(optFns).options
}

type withPresignClientFromClientOptions []func(*Options)

func (w withPresignClientFromClientOptions) options(o *PresignOptions) {
	o.ClientOptions = append(o.ClientOptions, w...)
}

// WithPresignExpires is a helper utility to append Expires value on presign
// options optional function
func WithPresignExpires(dur time.Duration) func(*PresignOptions) {
	return withPresignExpires(dur).options
}

type withPresignExpires time.Duration

func (w withPresignExpires) options(o *PresignOptions) {
	o.Expires = time.Duration(w)
}

// PresignClient represents the presign url client
type PresignClient struct {
	client  *Client
	options PresignOptions
}

// NewPresignClient generates a presign client using provided API Client and
// presign options
func NewPresignClient(c *Client, optFns ...func(*PresignOptions)) *PresignClient {
	var options PresignOptions
	for _, fn := range optFns {
		fn(&options)
	}
	if len(options.ClientOptions) != 0 {
		c = New(c.options, options.ClientOptions...)
	}

	if options.Presigner == nil {
		options.Presigner = newDefaultV4Signer(c.options)
	}

	if options.presignerV4a == nil {
		options.presignerV4a = newDefaultV4aSigner(c.options)
	}

	return &PresignClient{
		client:  c,
		options: options,
	}
}

func withNopHTTPClientAPIOption(o *Options) {
	o.HTTPClient = smithyhttp.NopClient{}
}

type presignConverter PresignOptions

func (c presignConverter) convertToPresignMiddleware(stack *middleware.Stack, options Options) (err error) {
	stack.Finalize.Clear()
	stack.Deserialize.Clear()
	stack.Build.Remove((*awsmiddleware.ClientRequestID)(nil).ID())
	stack.Build.Remove("UserAgent")
	pmw := v4.NewPresignHTTPRequestMiddleware(v4.PresignHTTPRequestMiddlewareOptions{
		CredentialsProvider: options.Credentials,
		Presigner:           c.Presigner,
		LogSigning:          options.ClientLogMode.IsSigning(),
	})
	err = stack.Finalize.Add(pmw, middleware.After)
	if err != nil {
		return err
	}
	if err = smithyhttp.AddNoPayloadDefaultContentTypeRemover(stack); err != nil {
		return err
	}

	// add multi-region access point presigner
	signermv := s3cust.NewPresignHTTPRequestMiddleware(s3cust.PresignHTTPRequestMiddlewareOptions{
		CredentialsProvider: options.Credentials,
		V4Presigner:         c.Presigner,
		V4aPresigner:        c.presignerV4a,
		LogSigning:          options.ClientLogMode.IsSigning(),
	})
	err = s3cust.RegisterPreSigningMiddleware(stack, signermv)
	if err != nil {
		return err
	}

	if c.Expires < 0 {
		return fmt.Errorf("presign URL duration must be 0 or greater, %v", c.Expires)
	}
	// add middleware to set expiration for s3 presigned url, if expiration is set to
	// 0, this middleware sets a default expiration of 900 seconds
	err = stack.Build.Add(&s3cust.AddExpiresOnPresignedURL{Expires: c.Expires}, middleware.After)
	if err != nil {
		return err
	}
	err = presignedurlcust.AddAsIsPresigingMiddleware(stack)
	if err != nil {
		return err
	}
	return nil
}

func addRequestResponseLogging(stack *middleware.Stack, o Options) error {
	return stack.Deserialize.Add(&smithyhttp.RequestResponseLogger{
		LogRequest:          o.ClientLogMode.IsRequest(),
		LogRequestWithBody:  o.ClientLogMode.IsRequestWithBody(),
		LogResponse:         o.ClientLogMode.IsResponse(),
		LogResponseWithBody: o.ClientLogMode.IsResponseWithBody(),
	}, middleware.After)
}

type endpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*endpointDisableHTTPSMiddleware) ID() string {
	return "endpointDisableHTTPSMiddleware"
}

func (m *endpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointDisableHTTPS && !smithyhttp.GetHostnameImmutable(ctx) {
		req.URL.Scheme = "http"
	}

	return next.HandleSerialize(ctx, in)

}
func addendpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&endpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "OperationSerializer", middleware.Before)
}
