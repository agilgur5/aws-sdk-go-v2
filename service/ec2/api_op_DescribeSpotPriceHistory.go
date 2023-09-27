// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the Spot price history. For more information, see Spot Instance
// pricing history (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances-history.html)
// in the Amazon EC2 User Guide for Linux Instances. When you specify a start and
// end time, the operation returns the prices of the instance types within that
// time range. It also returns the last price change before the start time, which
// is the effective price as of the start time.
func (c *Client) DescribeSpotPriceHistory(ctx context.Context, params *DescribeSpotPriceHistoryInput, optFns ...func(*Options)) (*DescribeSpotPriceHistoryOutput, error) {
	if params == nil {
		params = &DescribeSpotPriceHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpotPriceHistory", params, optFns, c.addOperationDescribeSpotPriceHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpotPriceHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotPriceHistory.
type DescribeSpotPriceHistoryInput struct {

	// Filters the results by the specified Availability Zone.
	AvailabilityZone *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The date and time, up to the current date, from which to stop retrieving the
	// price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *time.Time

	// The filters.
	//   - availability-zone - The Availability Zone for which prices should be
	//   returned.
	//   - instance-type - The type of instance (for example, m3.medium ).
	//   - product-description - The product description for the Spot price ( Linux/UNIX
	//   | Red Hat Enterprise Linux | SUSE Linux | Windows | Linux/UNIX (Amazon VPC) |
	//   Red Hat Enterprise Linux (Amazon VPC) | SUSE Linux (Amazon VPC) | Windows
	//   (Amazon VPC) ).
	//   - spot-price - The Spot price. The value must match exactly (or use wildcards;
	//   greater than or less than comparison is not supported).
	//   - timestamp - The time stamp of the Spot price history, in UTC format (for
	//   example, YYYY-MM-DDTHH:MM:SSZ). You can use wildcards (* and ?). Greater than or
	//   less than comparison is not supported.
	Filters []types.Filter

	// Filters the results by the specified instance types.
	InstanceTypes []types.InstanceType

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	// Filters the results by the specified basic product descriptions.
	ProductDescriptions []string

	// The date and time, up to the past 90 days, from which to start retrieving the
	// price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time

	noSmithyDocumentSerde
}

func (*DescribeSpotPriceHistoryInput) operationName() string {
	return "DescribeSpotPriceHistory"
}

// Contains the output of DescribeSpotPriceHistory.
type DescribeSpotPriceHistoryOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// The historical Spot prices.
	SpotPriceHistory []types.SpotPrice

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSpotPriceHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotPriceHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotPriceHistory{}, middleware.After)
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
	if err = addDescribeSpotPriceHistoryResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotPriceHistory(options.Region), middleware.Before); err != nil {
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

// DescribeSpotPriceHistoryAPIClient is a client that implements the
// DescribeSpotPriceHistory operation.
type DescribeSpotPriceHistoryAPIClient interface {
	DescribeSpotPriceHistory(context.Context, *DescribeSpotPriceHistoryInput, ...func(*Options)) (*DescribeSpotPriceHistoryOutput, error)
}

var _ DescribeSpotPriceHistoryAPIClient = (*Client)(nil)

// DescribeSpotPriceHistoryPaginatorOptions is the paginator options for
// DescribeSpotPriceHistory
type DescribeSpotPriceHistoryPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSpotPriceHistoryPaginator is a paginator for DescribeSpotPriceHistory
type DescribeSpotPriceHistoryPaginator struct {
	options   DescribeSpotPriceHistoryPaginatorOptions
	client    DescribeSpotPriceHistoryAPIClient
	params    *DescribeSpotPriceHistoryInput
	nextToken *string
	firstPage bool
}

// NewDescribeSpotPriceHistoryPaginator returns a new
// DescribeSpotPriceHistoryPaginator
func NewDescribeSpotPriceHistoryPaginator(client DescribeSpotPriceHistoryAPIClient, params *DescribeSpotPriceHistoryInput, optFns ...func(*DescribeSpotPriceHistoryPaginatorOptions)) *DescribeSpotPriceHistoryPaginator {
	if params == nil {
		params = &DescribeSpotPriceHistoryInput{}
	}

	options := DescribeSpotPriceHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSpotPriceHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSpotPriceHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSpotPriceHistory page.
func (p *DescribeSpotPriceHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSpotPriceHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeSpotPriceHistory(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeSpotPriceHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotPriceHistory",
	}
}

type opDescribeSpotPriceHistoryResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeSpotPriceHistoryResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeSpotPriceHistoryResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ec2"
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
				signingName = "ec2"
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
				v4aScheme.SigningName = aws.String("ec2")
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

func addDescribeSpotPriceHistoryResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeSpotPriceHistoryResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
