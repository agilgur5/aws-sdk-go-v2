// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the virtual MFA devices defined in the Amazon Web Services account by
// assignment status. If you do not specify an assignment status, the operation
// returns a list of all virtual MFA devices. Assignment status can be Assigned ,
// Unassigned , or Any . IAM resource-listing operations return a subset of the
// available attributes for the resource. For example, this operation does not
// return tags, even though they are an attribute of the returned object. To view
// tag information for a virtual MFA device, see ListMFADeviceTags . You can
// paginate the results using the MaxItems and Marker parameters.
func (c *Client) ListVirtualMFADevices(ctx context.Context, params *ListVirtualMFADevicesInput, optFns ...func(*Options)) (*ListVirtualMFADevicesOutput, error) {
	if params == nil {
		params = &ListVirtualMFADevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualMFADevices", params, optFns, c.addOperationListVirtualMFADevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualMFADevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVirtualMFADevicesInput struct {

	// The status ( Unassigned or Assigned ) of the devices to list. If you do not
	// specify an AssignmentStatus , the operation defaults to Any , which lists both
	// assigned and unassigned virtual MFA devices.,
	AssignmentStatus types.AssignmentStatusType

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	noSmithyDocumentSerde
}

func (*ListVirtualMFADevicesInput) operationName() string {
	return "ListVirtualMFADevices"
}

// Contains the response to a successful ListVirtualMFADevices request.
type ListVirtualMFADevicesOutput struct {

	// The list of virtual MFA devices in the current account that match the
	// AssignmentStatus value that was passed in the request.
	//
	// This member is required.
	VirtualMFADevices []types.VirtualMFADevice

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVirtualMFADevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListVirtualMFADevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListVirtualMFADevices{}, middleware.After)
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
	if err = addListVirtualMFADevicesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVirtualMFADevices(options.Region), middleware.Before); err != nil {
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

// ListVirtualMFADevicesAPIClient is a client that implements the
// ListVirtualMFADevices operation.
type ListVirtualMFADevicesAPIClient interface {
	ListVirtualMFADevices(context.Context, *ListVirtualMFADevicesInput, ...func(*Options)) (*ListVirtualMFADevicesOutput, error)
}

var _ ListVirtualMFADevicesAPIClient = (*Client)(nil)

// ListVirtualMFADevicesPaginatorOptions is the paginator options for
// ListVirtualMFADevices
type ListVirtualMFADevicesPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVirtualMFADevicesPaginator is a paginator for ListVirtualMFADevices
type ListVirtualMFADevicesPaginator struct {
	options   ListVirtualMFADevicesPaginatorOptions
	client    ListVirtualMFADevicesAPIClient
	params    *ListVirtualMFADevicesInput
	nextToken *string
	firstPage bool
}

// NewListVirtualMFADevicesPaginator returns a new ListVirtualMFADevicesPaginator
func NewListVirtualMFADevicesPaginator(client ListVirtualMFADevicesAPIClient, params *ListVirtualMFADevicesInput, optFns ...func(*ListVirtualMFADevicesPaginatorOptions)) *ListVirtualMFADevicesPaginator {
	if params == nil {
		params = &ListVirtualMFADevicesInput{}
	}

	options := ListVirtualMFADevicesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVirtualMFADevicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVirtualMFADevicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVirtualMFADevices page.
func (p *ListVirtualMFADevicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVirtualMFADevicesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListVirtualMFADevices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVirtualMFADevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListVirtualMFADevices",
	}
}

type opListVirtualMFADevicesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListVirtualMFADevicesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListVirtualMFADevicesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "iam"
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
				signingName = "iam"
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
				v4aScheme.SigningName = aws.String("iam")
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

func addListVirtualMFADevicesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListVirtualMFADevicesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
