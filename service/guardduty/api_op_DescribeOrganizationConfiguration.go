// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the account selected as the delegated administrator
// for GuardDuty. There might be regional differences because some data sources
// might not be available in all the Amazon Web Services Regions where GuardDuty is
// presently supported. For more information, see Regions and endpoints (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html)
// .
func (c *Client) DescribeOrganizationConfiguration(ctx context.Context, params *DescribeOrganizationConfigurationInput, optFns ...func(*Options)) (*DescribeOrganizationConfigurationOutput, error) {
	if params == nil {
		params = &DescribeOrganizationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationConfiguration", params, optFns, c.addOperationDescribeOrganizationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConfigurationInput struct {

	// The ID of the detector to retrieve information about the delegated
	// administrator from.
	//
	// This member is required.
	DetectorId *string

	// You can use this parameter to indicate the maximum number of items that you
	// want in the response.
	MaxResults int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls to
	// the action, fill nextToken in the request with the value of NextToken from the
	// previous response to continue listing data.
	NextToken *string

	noSmithyDocumentSerde
}

func (*DescribeOrganizationConfigurationInput) operationName() string {
	return "DescribeOrganizationConfiguration"
}

type DescribeOrganizationConfigurationOutput struct {

	// Indicates whether the maximum number of allowed member accounts are already
	// associated with the delegated administrator account for your organization.
	//
	// This member is required.
	MemberAccountLimitReached bool

	// Indicates whether GuardDuty is automatically enabled for accounts added to the
	// organization. Even though this is still supported, we recommend using
	// AutoEnableOrganizationMembers to achieve the similar results.
	//
	// Deprecated: This field is deprecated, use AutoEnableOrganizationMembers instead
	AutoEnable bool

	// Indicates the auto-enablement configuration of GuardDuty for the member
	// accounts in the organization.
	//   - NEW : Indicates that when a new account joins the organization, they will
	//   have GuardDuty enabled automatically.
	//   - ALL : Indicates that all accounts in the organization have GuardDuty enabled
	//   automatically. This includes NEW accounts that join the organization and
	//   accounts that may have been suspended or removed from the organization in
	//   GuardDuty.
	//   - NONE : Indicates that GuardDuty will not be automatically enabled for any
	//   account in the organization. The administrator must manage GuardDuty for each
	//   account in the organization individually.
	AutoEnableOrganizationMembers types.AutoEnableMembers

	// Describes which data sources are enabled automatically for member accounts.
	//
	// Deprecated: This parameter is deprecated, use Features instead
	DataSources *types.OrganizationDataSourceConfigurationsResult

	// A list of features that are configured for this organization.
	Features []types.OrganizationFeatureConfigurationResult

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOrganizationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOrganizationConfiguration{}, middleware.After)
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
	if err = addDescribeOrganizationConfigurationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeOrganizationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConfiguration(options.Region), middleware.Before); err != nil {
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

// DescribeOrganizationConfigurationAPIClient is a client that implements the
// DescribeOrganizationConfiguration operation.
type DescribeOrganizationConfigurationAPIClient interface {
	DescribeOrganizationConfiguration(context.Context, *DescribeOrganizationConfigurationInput, ...func(*Options)) (*DescribeOrganizationConfigurationOutput, error)
}

var _ DescribeOrganizationConfigurationAPIClient = (*Client)(nil)

// DescribeOrganizationConfigurationPaginatorOptions is the paginator options for
// DescribeOrganizationConfiguration
type DescribeOrganizationConfigurationPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items that you
	// want in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrganizationConfigurationPaginator is a paginator for
// DescribeOrganizationConfiguration
type DescribeOrganizationConfigurationPaginator struct {
	options   DescribeOrganizationConfigurationPaginatorOptions
	client    DescribeOrganizationConfigurationAPIClient
	params    *DescribeOrganizationConfigurationInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrganizationConfigurationPaginator returns a new
// DescribeOrganizationConfigurationPaginator
func NewDescribeOrganizationConfigurationPaginator(client DescribeOrganizationConfigurationAPIClient, params *DescribeOrganizationConfigurationInput, optFns ...func(*DescribeOrganizationConfigurationPaginatorOptions)) *DescribeOrganizationConfigurationPaginator {
	if params == nil {
		params = &DescribeOrganizationConfigurationInput{}
	}

	options := DescribeOrganizationConfigurationPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrganizationConfigurationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrganizationConfigurationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrganizationConfiguration page.
func (p *DescribeOrganizationConfigurationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrganizationConfigurationOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeOrganizationConfiguration(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOrganizationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "DescribeOrganizationConfiguration",
	}
}

type opDescribeOrganizationConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeOrganizationConfigurationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeOrganizationConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "guardduty"
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
				signingName = "guardduty"
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
				v4aScheme.SigningName = aws.String("guardduty")
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

func addDescribeOrganizationConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeOrganizationConfigurationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
