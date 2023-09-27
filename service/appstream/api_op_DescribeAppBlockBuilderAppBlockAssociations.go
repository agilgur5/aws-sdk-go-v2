// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list that describes one or more app block builder associations.
func (c *Client) DescribeAppBlockBuilderAppBlockAssociations(ctx context.Context, params *DescribeAppBlockBuilderAppBlockAssociationsInput, optFns ...func(*Options)) (*DescribeAppBlockBuilderAppBlockAssociationsOutput, error) {
	if params == nil {
		params = &DescribeAppBlockBuilderAppBlockAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppBlockBuilderAppBlockAssociations", params, optFns, c.addOperationDescribeAppBlockBuilderAppBlockAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppBlockBuilderAppBlockAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppBlockBuilderAppBlockAssociationsInput struct {

	// The ARN of the app block.
	AppBlockArn *string

	// The name of the app block builder.
	AppBlockBuilderName *string

	// The maximum size of each page of results.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

func (*DescribeAppBlockBuilderAppBlockAssociationsInput) operationName() string {
	return "DescribeAppBlockBuilderAppBlockAssociations"
}

type DescribeAppBlockBuilderAppBlockAssociationsOutput struct {

	// This list of app block builders associated with app blocks.
	AppBlockBuilderAppBlockAssociations []types.AppBlockBuilderAppBlockAssociation

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppBlockBuilderAppBlockAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAppBlockBuilderAppBlockAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAppBlockBuilderAppBlockAssociations{}, middleware.After)
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
	if err = addDescribeAppBlockBuilderAppBlockAssociationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppBlockBuilderAppBlockAssociations(options.Region), middleware.Before); err != nil {
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

// DescribeAppBlockBuilderAppBlockAssociationsAPIClient is a client that
// implements the DescribeAppBlockBuilderAppBlockAssociations operation.
type DescribeAppBlockBuilderAppBlockAssociationsAPIClient interface {
	DescribeAppBlockBuilderAppBlockAssociations(context.Context, *DescribeAppBlockBuilderAppBlockAssociationsInput, ...func(*Options)) (*DescribeAppBlockBuilderAppBlockAssociationsOutput, error)
}

var _ DescribeAppBlockBuilderAppBlockAssociationsAPIClient = (*Client)(nil)

// DescribeAppBlockBuilderAppBlockAssociationsPaginatorOptions is the paginator
// options for DescribeAppBlockBuilderAppBlockAssociations
type DescribeAppBlockBuilderAppBlockAssociationsPaginatorOptions struct {
	// The maximum size of each page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAppBlockBuilderAppBlockAssociationsPaginator is a paginator for
// DescribeAppBlockBuilderAppBlockAssociations
type DescribeAppBlockBuilderAppBlockAssociationsPaginator struct {
	options   DescribeAppBlockBuilderAppBlockAssociationsPaginatorOptions
	client    DescribeAppBlockBuilderAppBlockAssociationsAPIClient
	params    *DescribeAppBlockBuilderAppBlockAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAppBlockBuilderAppBlockAssociationsPaginator returns a new
// DescribeAppBlockBuilderAppBlockAssociationsPaginator
func NewDescribeAppBlockBuilderAppBlockAssociationsPaginator(client DescribeAppBlockBuilderAppBlockAssociationsAPIClient, params *DescribeAppBlockBuilderAppBlockAssociationsInput, optFns ...func(*DescribeAppBlockBuilderAppBlockAssociationsPaginatorOptions)) *DescribeAppBlockBuilderAppBlockAssociationsPaginator {
	if params == nil {
		params = &DescribeAppBlockBuilderAppBlockAssociationsInput{}
	}

	options := DescribeAppBlockBuilderAppBlockAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAppBlockBuilderAppBlockAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAppBlockBuilderAppBlockAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAppBlockBuilderAppBlockAssociations page.
func (p *DescribeAppBlockBuilderAppBlockAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAppBlockBuilderAppBlockAssociationsOutput, error) {
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

	result, err := p.client.DescribeAppBlockBuilderAppBlockAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAppBlockBuilderAppBlockAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DescribeAppBlockBuilderAppBlockAssociations",
	}
}

type opDescribeAppBlockBuilderAppBlockAssociationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeAppBlockBuilderAppBlockAssociationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeAppBlockBuilderAppBlockAssociationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "appstream"
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
				signingName = "appstream"
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
				v4aScheme.SigningName = aws.String("appstream")
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

func addDescribeAppBlockBuilderAppBlockAssociationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeAppBlockBuilderAppBlockAssociationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
