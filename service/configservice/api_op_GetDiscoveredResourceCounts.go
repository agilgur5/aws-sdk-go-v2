// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the resource types, the number of each resource type, and the total
// number of resources that Config is recording in this region for your Amazon Web
// Services account. Example
//   - Config is recording three resource types in the US East (Ohio) Region for
//     your account: 25 EC2 instances, 20 IAM users, and 15 S3 buckets.
//   - You make a call to the GetDiscoveredResourceCounts action and specify that
//     you want all resource types.
//   - Config returns the following:
//   - The resource types (EC2 instances, IAM users, and S3 buckets).
//   - The number of each resource type (25, 20, and 15).
//   - The total number of all resources (60).
//
// The response is paginated. By default, Config lists 100 ResourceCount objects
// on each page. You can customize this number with the limit parameter. The
// response includes a nextToken string. To get the next page of results, run the
// request again and specify the string for the nextToken parameter. If you make a
// call to the GetDiscoveredResourceCounts action, you might not immediately
// receive resource counts in the following situations:
//   - You are a new Config customer.
//   - You just enabled resource recording.
//
// It might take a few minutes for Config to record and count your resources. Wait
// a few minutes and then retry the GetDiscoveredResourceCounts action.
func (c *Client) GetDiscoveredResourceCounts(ctx context.Context, params *GetDiscoveredResourceCountsInput, optFns ...func(*Options)) (*GetDiscoveredResourceCountsOutput, error) {
	if params == nil {
		params = &GetDiscoveredResourceCountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDiscoveredResourceCounts", params, optFns, c.addOperationGetDiscoveredResourceCountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDiscoveredResourceCountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDiscoveredResourceCountsInput struct {

	// The maximum number of ResourceCount objects returned on each page. The default
	// is 100. You cannot specify a number greater than 100. If you specify 0, Config
	// uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The comma-separated list that specifies the resource types that you want Config
	// to return (for example, "AWS::EC2::Instance" , "AWS::IAM::User" ). If a value
	// for resourceTypes is not specified, Config returns all resource types that
	// Config is recording in the region for your account. If the configuration
	// recorder is turned off, Config returns an empty list of ResourceCount objects.
	// If the configuration recorder is not recording a specific resource type (for
	// example, S3 buckets), that resource type is not returned in the list of
	// ResourceCount objects.
	ResourceTypes []string

	noSmithyDocumentSerde
}

func (*GetDiscoveredResourceCountsInput) operationName() string {
	return "GetDiscoveredResourceCounts"
}

type GetDiscoveredResourceCountsOutput struct {

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// The list of ResourceCount objects. Each object is listed in descending order by
	// the number of resources.
	ResourceCounts []types.ResourceCount

	// The total number of resources that Config is recording in the region for your
	// account. If you specify resource types in the request, Config returns only the
	// total number of resources for those resource types. Example
	//   - Config is recording three resource types in the US East (Ohio) Region for
	//   your account: 25 EC2 instances, 20 IAM users, and 15 S3 buckets, for a total of
	//   60 resources.
	//   - You make a call to the GetDiscoveredResourceCounts action and specify the
	//   resource type, "AWS::EC2::Instances" , in the request.
	//   - Config returns 25 for totalDiscoveredResources .
	TotalDiscoveredResources int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDiscoveredResourceCountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDiscoveredResourceCounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDiscoveredResourceCounts{}, middleware.After)
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
	if err = addGetDiscoveredResourceCountsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDiscoveredResourceCounts(options.Region), middleware.Before); err != nil {
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

// GetDiscoveredResourceCountsAPIClient is a client that implements the
// GetDiscoveredResourceCounts operation.
type GetDiscoveredResourceCountsAPIClient interface {
	GetDiscoveredResourceCounts(context.Context, *GetDiscoveredResourceCountsInput, ...func(*Options)) (*GetDiscoveredResourceCountsOutput, error)
}

var _ GetDiscoveredResourceCountsAPIClient = (*Client)(nil)

// GetDiscoveredResourceCountsPaginatorOptions is the paginator options for
// GetDiscoveredResourceCounts
type GetDiscoveredResourceCountsPaginatorOptions struct {
	// The maximum number of ResourceCount objects returned on each page. The default
	// is 100. You cannot specify a number greater than 100. If you specify 0, Config
	// uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDiscoveredResourceCountsPaginator is a paginator for
// GetDiscoveredResourceCounts
type GetDiscoveredResourceCountsPaginator struct {
	options   GetDiscoveredResourceCountsPaginatorOptions
	client    GetDiscoveredResourceCountsAPIClient
	params    *GetDiscoveredResourceCountsInput
	nextToken *string
	firstPage bool
}

// NewGetDiscoveredResourceCountsPaginator returns a new
// GetDiscoveredResourceCountsPaginator
func NewGetDiscoveredResourceCountsPaginator(client GetDiscoveredResourceCountsAPIClient, params *GetDiscoveredResourceCountsInput, optFns ...func(*GetDiscoveredResourceCountsPaginatorOptions)) *GetDiscoveredResourceCountsPaginator {
	if params == nil {
		params = &GetDiscoveredResourceCountsInput{}
	}

	options := GetDiscoveredResourceCountsPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetDiscoveredResourceCountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDiscoveredResourceCountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetDiscoveredResourceCounts page.
func (p *GetDiscoveredResourceCountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDiscoveredResourceCountsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.GetDiscoveredResourceCounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetDiscoveredResourceCounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetDiscoveredResourceCounts",
	}
}

type opGetDiscoveredResourceCountsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetDiscoveredResourceCountsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetDiscoveredResourceCountsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "config"
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
				signingName = "config"
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
				v4aScheme.SigningName = aws.String("config")
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

func addGetDiscoveredResourceCountsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetDiscoveredResourceCountsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
