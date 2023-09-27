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

// Describes one or more blue/green deployments. For more information, see Using
// Amazon RDS Blue/Green Deployments for database updates (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/blue-green-deployments.html)
// in the Amazon RDS User Guide and Using Amazon RDS Blue/Green Deployments for
// database updates (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html)
// in the Amazon Aurora User Guide.
func (c *Client) DescribeBlueGreenDeployments(ctx context.Context, params *DescribeBlueGreenDeploymentsInput, optFns ...func(*Options)) (*DescribeBlueGreenDeploymentsOutput, error) {
	if params == nil {
		params = &DescribeBlueGreenDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBlueGreenDeployments", params, optFns, c.addOperationDescribeBlueGreenDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBlueGreenDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBlueGreenDeploymentsInput struct {

	// The blue/green deployment identifier. If you specify this parameter, the
	// response only includes information about the specific blue/green deployment.
	// This parameter isn't case-sensitive. Constraints:
	//   - Must match an existing blue/green deployment identifier.
	BlueGreenDeploymentIdentifier *string

	// A filter that specifies one or more blue/green deployments to describe. Valid
	// Values:
	//   - blue-green-deployment-identifier - Accepts system-generated identifiers for
	//   blue/green deployments. The results list only includes information about the
	//   blue/green deployments with the specified identifiers.
	//   - blue-green-deployment-name - Accepts user-supplied names for blue/green
	//   deployments. The results list only includes information about the blue/green
	//   deployments with the specified names.
	//   - source - Accepts source databases for a blue/green deployment. The results
	//   list only includes information about the blue/green deployments with the
	//   specified source databases.
	//   - target - Accepts target databases for a blue/green deployment. The results
	//   list only includes information about the blue/green deployments with the
	//   specified target databases.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeBlueGreenDeployments
	// request. If you specify this parameter, the response only includes records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints:
	//   - Must be a minimum of 20.
	//   - Can't exceed 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

func (*DescribeBlueGreenDeploymentsInput) operationName() string {
	return "DescribeBlueGreenDeployments"
}

type DescribeBlueGreenDeploymentsOutput struct {

	// A list of blue/green deployments in the current account and Amazon Web Services
	// Region.
	BlueGreenDeployments []types.BlueGreenDeployment

	// A pagination token that can be used in a later DescribeBlueGreenDeployments
	// request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBlueGreenDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeBlueGreenDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeBlueGreenDeployments{}, middleware.After)
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
	if err = addDescribeBlueGreenDeploymentsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeBlueGreenDeploymentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBlueGreenDeployments(options.Region), middleware.Before); err != nil {
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

// DescribeBlueGreenDeploymentsAPIClient is a client that implements the
// DescribeBlueGreenDeployments operation.
type DescribeBlueGreenDeploymentsAPIClient interface {
	DescribeBlueGreenDeployments(context.Context, *DescribeBlueGreenDeploymentsInput, ...func(*Options)) (*DescribeBlueGreenDeploymentsOutput, error)
}

var _ DescribeBlueGreenDeploymentsAPIClient = (*Client)(nil)

// DescribeBlueGreenDeploymentsPaginatorOptions is the paginator options for
// DescribeBlueGreenDeployments
type DescribeBlueGreenDeploymentsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints:
	//   - Must be a minimum of 20.
	//   - Can't exceed 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBlueGreenDeploymentsPaginator is a paginator for
// DescribeBlueGreenDeployments
type DescribeBlueGreenDeploymentsPaginator struct {
	options   DescribeBlueGreenDeploymentsPaginatorOptions
	client    DescribeBlueGreenDeploymentsAPIClient
	params    *DescribeBlueGreenDeploymentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeBlueGreenDeploymentsPaginator returns a new
// DescribeBlueGreenDeploymentsPaginator
func NewDescribeBlueGreenDeploymentsPaginator(client DescribeBlueGreenDeploymentsAPIClient, params *DescribeBlueGreenDeploymentsInput, optFns ...func(*DescribeBlueGreenDeploymentsPaginatorOptions)) *DescribeBlueGreenDeploymentsPaginator {
	if params == nil {
		params = &DescribeBlueGreenDeploymentsInput{}
	}

	options := DescribeBlueGreenDeploymentsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBlueGreenDeploymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBlueGreenDeploymentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBlueGreenDeployments page.
func (p *DescribeBlueGreenDeploymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBlueGreenDeploymentsOutput, error) {
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

	result, err := p.client.DescribeBlueGreenDeployments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBlueGreenDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeBlueGreenDeployments",
	}
}

type opDescribeBlueGreenDeploymentsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeBlueGreenDeploymentsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeBlueGreenDeploymentsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDescribeBlueGreenDeploymentsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeBlueGreenDeploymentsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
