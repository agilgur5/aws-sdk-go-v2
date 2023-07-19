// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the metrics configurations for the bucket. The metrics configurations are
// only for the request metrics of the bucket and do not provide information on
// daily storage metrics. You can have up to 1,000 configurations per bucket. This
// action supports list pagination and does not return more than 100 configurations
// at a time. Always check the IsTruncated element in the response. If there are
// no more configurations to list, IsTruncated is set to false. If there are more
// configurations to list, IsTruncated is set to true, and there is a value in
// NextContinuationToken . You use the NextContinuationToken value to continue the
// pagination of the list by passing the value in continuation-token in the
// request to GET the next page. To use this operation, you must have permissions
// to perform the s3:GetMetricsConfiguration action. The bucket owner has this
// permission by default. The bucket owner can grant this permission to others. For
// more information about permissions, see Permissions Related to Bucket
// Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html)
// . For more information about metrics configurations and CloudWatch request
// metrics, see Monitoring Metrics with Amazon CloudWatch (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html)
// . The following operations are related to ListBucketMetricsConfigurations :
//   - PutBucketMetricsConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html)
//   - GetBucketMetricsConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html)
//   - DeleteBucketMetricsConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html)
func (c *Client) ListBucketMetricsConfigurations(ctx context.Context, params *ListBucketMetricsConfigurationsInput, optFns ...func(*Options)) (*ListBucketMetricsConfigurationsOutput, error) {
	if params == nil {
		params = &ListBucketMetricsConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBucketMetricsConfigurations", params, optFns, c.addOperationListBucketMetricsConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBucketMetricsConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketMetricsConfigurationsInput struct {

	// The name of the bucket containing the metrics configurations to retrieve.
	//
	// This member is required.
	Bucket *string

	// The marker that is used to continue a metrics configuration listing that has
	// been truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value that
	// Amazon S3 understands.
	ContinuationToken *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type ListBucketMetricsConfigurationsOutput struct {

	// The marker that is used as a starting point for this metrics configuration list
	// response. This value is present if it was sent in the request.
	ContinuationToken *string

	// Indicates whether the returned list of metrics configurations is complete. A
	// value of true indicates that the list is not complete and the
	// NextContinuationToken will be provided for a subsequent request.
	IsTruncated bool

	// The list of metrics configurations for a bucket.
	MetricsConfigurationList []types.MetricsConfiguration

	// The marker used to continue a metrics configuration listing that has been
	// truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value that
	// Amazon S3 understands.
	NextContinuationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBucketMetricsConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListBucketMetricsConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListBucketMetricsConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addLegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addListBucketMetricsConfigurationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListBucketMetricsConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBucketMetricsConfigurations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListBucketMetricsConfigurationsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func (v *ListBucketMetricsConfigurationsInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opListBucketMetricsConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListBucketMetricsConfigurations",
	}
}

// getListBucketMetricsConfigurationsBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getListBucketMetricsConfigurationsBucketMember(input interface{}) (*string, bool) {
	in := input.(*ListBucketMetricsConfigurationsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addListBucketMetricsConfigurationsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getListBucketMetricsConfigurationsBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}

type opListBucketMetricsConfigurationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opListBucketMetricsConfigurationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListBucketMetricsConfigurationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*ListBucketMetricsConfigurationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.Bucket = input.Bucket

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

	ctx = smithyhttp.DisableEndpointHostPrefix(ctx, true)

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "s3"
			signingRegion := m.BuiltInResolver.(*BuiltInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			ctx = s3cust.SetSignerVersion(ctx, internalauth.SigV4)
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
				signingName = "s3"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*BuiltInResolver).Region
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
			ctx = s3cust.SetSignerVersion(ctx, v4Scheme.Name)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("s3")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			ctx = s3cust.SetSignerVersion(ctx, v4aScheme.Name)
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListBucketMetricsConfigurationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListBucketMetricsConfigurationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:                         options.Region,
			UseFIPS:                        options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack:                   options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:                       options.BaseEndpoint,
			ForcePathStyle:                 options.UsePathStyle,
			Accelerate:                     options.UseAccelerate,
			DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
			UseArnRegion:                   options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
