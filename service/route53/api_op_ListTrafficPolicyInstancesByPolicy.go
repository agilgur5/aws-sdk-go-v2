// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the traffic policy instances that you created by using a
// specify traffic policy version. After you submit a CreateTrafficPolicyInstance
// or an UpdateTrafficPolicyInstance request, there's a brief delay while Amazon
// Route 53 creates the resource record sets that are specified in the traffic
// policy definition. For more information, see the State response element. Route
// 53 returns a maximum of 100 items in each response. If you have a lot of traffic
// policy instances, you can use the MaxItems parameter to list them in groups of
// up to 100.
func (c *Client) ListTrafficPolicyInstancesByPolicy(ctx context.Context, params *ListTrafficPolicyInstancesByPolicyInput, optFns ...func(*Options)) (*ListTrafficPolicyInstancesByPolicyOutput, error) {
	if params == nil {
		params = &ListTrafficPolicyInstancesByPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrafficPolicyInstancesByPolicy", params, optFns, c.addOperationListTrafficPolicyInstancesByPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrafficPolicyInstancesByPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains the information about the request to list your
// traffic policy instances.
type ListTrafficPolicyInstancesByPolicyInput struct {

	// The ID of the traffic policy for which you want to list traffic policy
	// instances.
	//
	// This member is required.
	TrafficPolicyId *string

	// The version of the traffic policy for which you want to list traffic policy
	// instances. The version must be associated with the traffic policy that is
	// specified by TrafficPolicyId .
	//
	// This member is required.
	TrafficPolicyVersion *int32

	// If the value of IsTruncated in the previous response was true , you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request. For the value of hostedzoneid ,
	// specify the value of HostedZoneIdMarker from the previous response, which is
	// the hosted zone ID of the first traffic policy instance that Amazon Route 53
	// will return if you submit another request. If the value of IsTruncated in the
	// previous response was false , there are no more traffic policy instances to get.
	HostedZoneIdMarker *string

	// The maximum number of traffic policy instances to be included in the response
	// body for this request. If you have more than MaxItems traffic policy instances,
	// the value of the IsTruncated element in the response is true , and the values of
	// HostedZoneIdMarker , TrafficPolicyInstanceNameMarker , and
	// TrafficPolicyInstanceTypeMarker represent the first traffic policy instance that
	// Amazon Route 53 will return if you submit another request.
	MaxItems *int32

	// If the value of IsTruncated in the previous response was true , you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request. For the value of
	// trafficpolicyinstancename , specify the value of TrafficPolicyInstanceNameMarker
	// from the previous response, which is the name of the first traffic policy
	// instance that Amazon Route 53 will return if you submit another request. If the
	// value of IsTruncated in the previous response was false , there are no more
	// traffic policy instances to get.
	TrafficPolicyInstanceNameMarker *string

	// If the value of IsTruncated in the previous response was true , you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request. For the value of
	// trafficpolicyinstancetype , specify the value of TrafficPolicyInstanceTypeMarker
	// from the previous response, which is the name of the first traffic policy
	// instance that Amazon Route 53 will return if you submit another request. If the
	// value of IsTruncated in the previous response was false , there are no more
	// traffic policy instances to get.
	TrafficPolicyInstanceTypeMarker types.RRType

	noSmithyDocumentSerde
}

func (*ListTrafficPolicyInstancesByPolicyInput) operationName() string {
	return "ListTrafficPolicyInstancesByPolicy"
}

// A complex type that contains the response information for the request.
type ListTrafficPolicyInstancesByPolicyOutput struct {

	// A flag that indicates whether there are more traffic policy instances to be
	// listed. If the response was truncated, you can get the next group of traffic
	// policy instances by calling ListTrafficPolicyInstancesByPolicy again and
	// specifying the values of the HostedZoneIdMarker ,
	// TrafficPolicyInstanceNameMarker , and TrafficPolicyInstanceTypeMarker elements
	// in the corresponding request parameters.
	//
	// This member is required.
	IsTruncated bool

	// The value that you specified for the MaxItems parameter in the call to
	// ListTrafficPolicyInstancesByPolicy that produced the current response.
	//
	// This member is required.
	MaxItems *int32

	// A list that contains one TrafficPolicyInstance element for each traffic policy
	// instance that matches the elements in the request.
	//
	// This member is required.
	TrafficPolicyInstances []types.TrafficPolicyInstance

	// If IsTruncated is true , HostedZoneIdMarker is the ID of the hosted zone of the
	// first traffic policy instance in the next group of traffic policy instances.
	HostedZoneIdMarker *string

	// If IsTruncated is true , TrafficPolicyInstanceNameMarker is the name of the
	// first traffic policy instance in the next group of MaxItems traffic policy
	// instances.
	TrafficPolicyInstanceNameMarker *string

	// If IsTruncated is true , TrafficPolicyInstanceTypeMarker is the DNS type of the
	// resource record sets that are associated with the first traffic policy instance
	// in the next group of MaxItems traffic policy instances.
	TrafficPolicyInstanceTypeMarker types.RRType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrafficPolicyInstancesByPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTrafficPolicyInstancesByPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTrafficPolicyInstancesByPolicy{}, middleware.After)
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
	if err = addListTrafficPolicyInstancesByPolicyResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListTrafficPolicyInstancesByPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficPolicyInstancesByPolicy(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
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

func newServiceMetadataMiddleware_opListTrafficPolicyInstancesByPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListTrafficPolicyInstancesByPolicy",
	}
}

type opListTrafficPolicyInstancesByPolicyResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListTrafficPolicyInstancesByPolicyResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListTrafficPolicyInstancesByPolicyResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "route53"
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
				signingName = "route53"
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
				v4aScheme.SigningName = aws.String("route53")
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

func addListTrafficPolicyInstancesByPolicyResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListTrafficPolicyInstancesByPolicyResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
