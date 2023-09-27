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
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a snapshot or cluster export to Amazon S3. This API
// operation supports pagination.
func (c *Client) DescribeExportTasks(ctx context.Context, params *DescribeExportTasksInput, optFns ...func(*Options)) (*DescribeExportTasksOutput, error) {
	if params == nil {
		params = &DescribeExportTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExportTasks", params, optFns, c.addOperationDescribeExportTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExportTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportTasksInput struct {

	// The identifier of the snapshot or cluster export task to be described.
	ExportTaskIdentifier *string

	// Filters specify one or more snapshot or cluster exports to describe. The
	// filters are specified as name-value pairs that define what to include in the
	// output. Filter names and values are case-sensitive. Supported filters include
	// the following:
	//   - export-task-identifier - An identifier for the snapshot or cluster export
	//   task.
	//   - s3-bucket - The Amazon S3 bucket the data is exported to.
	//   - source-arn - The Amazon Resource Name (ARN) of the snapshot or cluster
	//   exported to Amazon S3.
	//   - status - The status of the export task. Must be lowercase. Valid statuses
	//   are the following:
	//   - canceled
	//   - canceling
	//   - complete
	//   - failed
	//   - in_progress
	//   - starting
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeExportTasks
	// request. If you specify this parameter, the response includes only records
	// beyond the marker, up to the value specified by the MaxRecords parameter.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified value, a pagination token called a marker is included in the
	// response. You can use the marker in a later DescribeExportTasks request to
	// retrieve the remaining results. Default: 100 Constraints: Minimum 20, maximum
	// 100.
	MaxRecords *int32

	// The Amazon Resource Name (ARN) of the snapshot or cluster exported to Amazon S3.
	SourceArn *string

	// The type of source for the export.
	SourceType types.ExportSourceType

	noSmithyDocumentSerde
}

func (*DescribeExportTasksInput) operationName() string {
	return "DescribeExportTasks"
}

type DescribeExportTasksOutput struct {

	// Information about an export of a snapshot or cluster to Amazon S3.
	ExportTasks []types.ExportTask

	// A pagination token that can be used in a later DescribeExportTasks request. A
	// marker is used for pagination to identify the location to begin output for the
	// next response of DescribeExportTasks .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExportTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeExportTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeExportTasks{}, middleware.After)
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
	if err = addDescribeExportTasksResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeExportTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExportTasks(options.Region), middleware.Before); err != nil {
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

// DescribeExportTasksAPIClient is a client that implements the
// DescribeExportTasks operation.
type DescribeExportTasksAPIClient interface {
	DescribeExportTasks(context.Context, *DescribeExportTasksInput, ...func(*Options)) (*DescribeExportTasksOutput, error)
}

var _ DescribeExportTasksAPIClient = (*Client)(nil)

// DescribeExportTasksPaginatorOptions is the paginator options for
// DescribeExportTasks
type DescribeExportTasksPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified value, a pagination token called a marker is included in the
	// response. You can use the marker in a later DescribeExportTasks request to
	// retrieve the remaining results. Default: 100 Constraints: Minimum 20, maximum
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeExportTasksPaginator is a paginator for DescribeExportTasks
type DescribeExportTasksPaginator struct {
	options   DescribeExportTasksPaginatorOptions
	client    DescribeExportTasksAPIClient
	params    *DescribeExportTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeExportTasksPaginator returns a new DescribeExportTasksPaginator
func NewDescribeExportTasksPaginator(client DescribeExportTasksAPIClient, params *DescribeExportTasksInput, optFns ...func(*DescribeExportTasksPaginatorOptions)) *DescribeExportTasksPaginator {
	if params == nil {
		params = &DescribeExportTasksInput{}
	}

	options := DescribeExportTasksPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeExportTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeExportTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeExportTasks page.
func (p *DescribeExportTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeExportTasksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeExportTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeExportTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeExportTasks",
	}
}

type opDescribeExportTasksResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeExportTasksResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeExportTasksResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDescribeExportTasksResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeExportTasksResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
