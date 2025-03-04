// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the phone numbers for the specified Amazon Connect
// instance. For more information about phone numbers, see Set Up Phone Numbers
// for Your Contact Center (https://docs.aws.amazon.com/connect/latest/adminguide/contact-center-phone-number.html)
// in the Amazon Connect Administrator Guide.
//   - We recommend using ListPhoneNumbersV2 (https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html)
//     to return phone number types. ListPhoneNumbers doesn't support number types
//     UIFN , SHARED , THIRD_PARTY_TF , and THIRD_PARTY_DID . While it returns
//     numbers of those types, it incorrectly lists them as TOLL_FREE or DID .
//   - The phone number Arn value that is returned from each of the items in the
//     PhoneNumberSummaryList (https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbers.html#connect-ListPhoneNumbers-response-PhoneNumberSummaryList)
//     cannot be used to tag phone number resources. It will fail with a
//     ResourceNotFoundException . Instead, use the ListPhoneNumbersV2 (https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html)
//     API. It returns the new phone number ARN that can be used to tag phone number
//     resources.
func (c *Client) ListPhoneNumbers(ctx context.Context, params *ListPhoneNumbersInput, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) {
	if params == nil {
		params = &ListPhoneNumbersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPhoneNumbers", params, optFns, c.addOperationListPhoneNumbersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPhoneNumbersInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The ISO country code.
	PhoneNumberCountryCodes []types.PhoneNumberCountryCode

	// The type of phone number. We recommend using ListPhoneNumbersV2 (https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html)
	// to return phone number types. While ListPhoneNumbers returns number types UIFN ,
	// SHARED , THIRD_PARTY_TF , and THIRD_PARTY_DID , it incorrectly lists them as
	// TOLL_FREE or DID .
	PhoneNumberTypes []types.PhoneNumberType

	noSmithyDocumentSerde
}

type ListPhoneNumbersOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the phone numbers.
	PhoneNumberSummaryList []types.PhoneNumberSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPhoneNumbersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPhoneNumbers{}, middleware.After)
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
	if err = addListPhoneNumbersResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListPhoneNumbersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPhoneNumbers(options.Region), middleware.Before); err != nil {
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

// ListPhoneNumbersAPIClient is a client that implements the ListPhoneNumbers
// operation.
type ListPhoneNumbersAPIClient interface {
	ListPhoneNumbers(context.Context, *ListPhoneNumbersInput, ...func(*Options)) (*ListPhoneNumbersOutput, error)
}

var _ ListPhoneNumbersAPIClient = (*Client)(nil)

// ListPhoneNumbersPaginatorOptions is the paginator options for ListPhoneNumbers
type ListPhoneNumbersPaginatorOptions struct {
	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPhoneNumbersPaginator is a paginator for ListPhoneNumbers
type ListPhoneNumbersPaginator struct {
	options   ListPhoneNumbersPaginatorOptions
	client    ListPhoneNumbersAPIClient
	params    *ListPhoneNumbersInput
	nextToken *string
	firstPage bool
}

// NewListPhoneNumbersPaginator returns a new ListPhoneNumbersPaginator
func NewListPhoneNumbersPaginator(client ListPhoneNumbersAPIClient, params *ListPhoneNumbersInput, optFns ...func(*ListPhoneNumbersPaginatorOptions)) *ListPhoneNumbersPaginator {
	if params == nil {
		params = &ListPhoneNumbersInput{}
	}

	options := ListPhoneNumbersPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPhoneNumbersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPhoneNumbersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPhoneNumbers page.
func (p *ListPhoneNumbersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPhoneNumbers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPhoneNumbers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListPhoneNumbers",
	}
}

type opListPhoneNumbersResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListPhoneNumbersResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListPhoneNumbersResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "connect"
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
				signingName = "connect"
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
				v4aScheme.SigningName = aws.String("connect")
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

func addListPhoneNumbersResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListPhoneNumbersResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
