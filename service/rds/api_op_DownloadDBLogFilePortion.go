// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Downloads all or a portion of the specified log file, up to 1 MB in size. This
// command doesn't apply to RDS Custom.
func (c *Client) DownloadDBLogFilePortion(ctx context.Context, params *DownloadDBLogFilePortionInput, optFns ...func(*Options)) (*DownloadDBLogFilePortionOutput, error) {
	if params == nil {
		params = &DownloadDBLogFilePortionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DownloadDBLogFilePortion", params, optFns, c.addOperationDownloadDBLogFilePortionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DownloadDBLogFilePortionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DownloadDBLogFilePortionInput struct {

	// The customer-assigned name of the DB instance that contains the log files you
	// want to list. Constraints:
	//   - Must match the identifier of an existing DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The name of the log file to be downloaded.
	//
	// This member is required.
	LogFileName *string

	// The pagination token provided in the previous request or "0". If the Marker
	// parameter is specified the response includes only records beyond the marker
	// until the end of the file or up to NumberOfLines.
	Marker *string

	// The number of lines to download. If the number of lines specified results in a
	// file over 1 MB in size, the file is truncated at 1 MB in size. If the
	// NumberOfLines parameter is specified, then the block of lines returned can be
	// from the beginning or the end of the log file, depending on the value of the
	// Marker parameter.
	//   - If neither Marker or NumberOfLines are specified, the entire log file is
	//   returned up to a maximum of 10000 lines, starting with the most recent log
	//   entries first.
	//   - If NumberOfLines is specified and Marker isn't specified, then the most
	//   recent lines from the end of the log file are returned.
	//   - If Marker is specified as "0", then the specified number of lines from the
	//   beginning of the log file are returned.
	//   - You can download the log file in blocks of lines by specifying the size of
	//   the block using the NumberOfLines parameter, and by specifying a value of "0"
	//   for the Marker parameter in your first request. Include the Marker value
	//   returned in the response as the Marker value for the next request, continuing
	//   until the AdditionalDataPending response element returns false.
	NumberOfLines int32

	noSmithyDocumentSerde
}

func (*DownloadDBLogFilePortionInput) operationName() string {
	return "DownloadDBLogFilePortion"
}

// This data type is used as a response element to DownloadDBLogFilePortion .
type DownloadDBLogFilePortionOutput struct {

	// A Boolean value that, if true, indicates there is more data to be downloaded.
	AdditionalDataPending bool

	// Entries from the specified log file.
	LogFileData *string

	// A pagination token that can be used in a later DownloadDBLogFilePortion request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDownloadDBLogFilePortionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDownloadDBLogFilePortion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDownloadDBLogFilePortion{}, middleware.After)
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
	if err = addDownloadDBLogFilePortionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDownloadDBLogFilePortionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDownloadDBLogFilePortion(options.Region), middleware.Before); err != nil {
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

// DownloadDBLogFilePortionAPIClient is a client that implements the
// DownloadDBLogFilePortion operation.
type DownloadDBLogFilePortionAPIClient interface {
	DownloadDBLogFilePortion(context.Context, *DownloadDBLogFilePortionInput, ...func(*Options)) (*DownloadDBLogFilePortionOutput, error)
}

var _ DownloadDBLogFilePortionAPIClient = (*Client)(nil)

// DownloadDBLogFilePortionPaginatorOptions is the paginator options for
// DownloadDBLogFilePortion
type DownloadDBLogFilePortionPaginatorOptions struct {
	// The number of lines to download. If the number of lines specified results in a
	// file over 1 MB in size, the file is truncated at 1 MB in size. If the
	// NumberOfLines parameter is specified, then the block of lines returned can be
	// from the beginning or the end of the log file, depending on the value of the
	// Marker parameter.
	//   - If neither Marker or NumberOfLines are specified, the entire log file is
	//   returned up to a maximum of 10000 lines, starting with the most recent log
	//   entries first.
	//   - If NumberOfLines is specified and Marker isn't specified, then the most
	//   recent lines from the end of the log file are returned.
	//   - If Marker is specified as "0", then the specified number of lines from the
	//   beginning of the log file are returned.
	//   - You can download the log file in blocks of lines by specifying the size of
	//   the block using the NumberOfLines parameter, and by specifying a value of "0"
	//   for the Marker parameter in your first request. Include the Marker value
	//   returned in the response as the Marker value for the next request, continuing
	//   until the AdditionalDataPending response element returns false.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DownloadDBLogFilePortionPaginator is a paginator for DownloadDBLogFilePortion
type DownloadDBLogFilePortionPaginator struct {
	options   DownloadDBLogFilePortionPaginatorOptions
	client    DownloadDBLogFilePortionAPIClient
	params    *DownloadDBLogFilePortionInput
	nextToken *string
	firstPage bool
}

// NewDownloadDBLogFilePortionPaginator returns a new
// DownloadDBLogFilePortionPaginator
func NewDownloadDBLogFilePortionPaginator(client DownloadDBLogFilePortionAPIClient, params *DownloadDBLogFilePortionInput, optFns ...func(*DownloadDBLogFilePortionPaginatorOptions)) *DownloadDBLogFilePortionPaginator {
	if params == nil {
		params = &DownloadDBLogFilePortionInput{}
	}

	options := DownloadDBLogFilePortionPaginatorOptions{}
	if params.NumberOfLines != 0 {
		options.Limit = params.NumberOfLines
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DownloadDBLogFilePortionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DownloadDBLogFilePortionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DownloadDBLogFilePortion page.
func (p *DownloadDBLogFilePortionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DownloadDBLogFilePortionOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	params.NumberOfLines = p.options.Limit

	result, err := p.client.DownloadDBLogFilePortion(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDownloadDBLogFilePortion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DownloadDBLogFilePortion",
	}
}

type opDownloadDBLogFilePortionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDownloadDBLogFilePortionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDownloadDBLogFilePortionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "rds"
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
				signingName = "rds"
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
				v4aScheme.SigningName = aws.String("rds")
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

func addDownloadDBLogFilePortionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDownloadDBLogFilePortionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
