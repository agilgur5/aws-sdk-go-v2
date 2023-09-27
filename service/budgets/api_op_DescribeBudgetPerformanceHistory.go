// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the history for DAILY , MONTHLY , and QUARTERLY budgets. Budget
// history isn't available for ANNUAL budgets.
func (c *Client) DescribeBudgetPerformanceHistory(ctx context.Context, params *DescribeBudgetPerformanceHistoryInput, optFns ...func(*Options)) (*DescribeBudgetPerformanceHistoryOutput, error) {
	if params == nil {
		params = &DescribeBudgetPerformanceHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgetPerformanceHistory", params, optFns, c.addOperationDescribeBudgetPerformanceHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetPerformanceHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBudgetPerformanceHistoryInput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	//
	// This member is required.
	BudgetName *string

	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	MaxResults *int32

	// A generic string.
	NextToken *string

	// Retrieves how often the budget went into an ALARM state for the specified time
	// period.
	TimePeriod *types.TimePeriod

	noSmithyDocumentSerde
}

func (*DescribeBudgetPerformanceHistoryInput) operationName() string {
	return "DescribeBudgetPerformanceHistory"
}

type DescribeBudgetPerformanceHistoryOutput struct {

	// The history of how often the budget has gone into an ALARM state. For DAILY
	// budgets, the history saves the state of the budget for the last 60 days. For
	// MONTHLY budgets, the history saves the state of the budget for the current month
	// plus the last 12 months. For QUARTERLY budgets, the history saves the state of
	// the budget for the last four quarters.
	BudgetPerformanceHistory *types.BudgetPerformanceHistory

	// A generic string.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBudgetPerformanceHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudgetPerformanceHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudgetPerformanceHistory{}, middleware.After)
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
	if err = addDescribeBudgetPerformanceHistoryResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeBudgetPerformanceHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudgetPerformanceHistory(options.Region), middleware.Before); err != nil {
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

// DescribeBudgetPerformanceHistoryAPIClient is a client that implements the
// DescribeBudgetPerformanceHistory operation.
type DescribeBudgetPerformanceHistoryAPIClient interface {
	DescribeBudgetPerformanceHistory(context.Context, *DescribeBudgetPerformanceHistoryInput, ...func(*Options)) (*DescribeBudgetPerformanceHistoryOutput, error)
}

var _ DescribeBudgetPerformanceHistoryAPIClient = (*Client)(nil)

// DescribeBudgetPerformanceHistoryPaginatorOptions is the paginator options for
// DescribeBudgetPerformanceHistory
type DescribeBudgetPerformanceHistoryPaginatorOptions struct {
	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBudgetPerformanceHistoryPaginator is a paginator for
// DescribeBudgetPerformanceHistory
type DescribeBudgetPerformanceHistoryPaginator struct {
	options   DescribeBudgetPerformanceHistoryPaginatorOptions
	client    DescribeBudgetPerformanceHistoryAPIClient
	params    *DescribeBudgetPerformanceHistoryInput
	nextToken *string
	firstPage bool
}

// NewDescribeBudgetPerformanceHistoryPaginator returns a new
// DescribeBudgetPerformanceHistoryPaginator
func NewDescribeBudgetPerformanceHistoryPaginator(client DescribeBudgetPerformanceHistoryAPIClient, params *DescribeBudgetPerformanceHistoryInput, optFns ...func(*DescribeBudgetPerformanceHistoryPaginatorOptions)) *DescribeBudgetPerformanceHistoryPaginator {
	if params == nil {
		params = &DescribeBudgetPerformanceHistoryInput{}
	}

	options := DescribeBudgetPerformanceHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBudgetPerformanceHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBudgetPerformanceHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBudgetPerformanceHistory page.
func (p *DescribeBudgetPerformanceHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBudgetPerformanceHistoryOutput, error) {
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

	result, err := p.client.DescribeBudgetPerformanceHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBudgetPerformanceHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudgetPerformanceHistory",
	}
}

type opDescribeBudgetPerformanceHistoryResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeBudgetPerformanceHistoryResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeBudgetPerformanceHistoryResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "budgets"
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
				signingName = "budgets"
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
				v4aScheme.SigningName = aws.String("budgets")
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

func addDescribeBudgetPerformanceHistoryResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeBudgetPerformanceHistoryResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
