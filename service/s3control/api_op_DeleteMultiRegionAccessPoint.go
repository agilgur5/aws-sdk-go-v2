// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Deletes a Multi-Region Access Point. This action does not delete the buckets
// associated with the Multi-Region Access Point, only the Multi-Region Access
// Point itself. This action will always be routed to the US West (Oregon) Region.
// For more information about the restrictions around managing Multi-Region Access
// Points, see Managing Multi-Region Access Points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/ManagingMultiRegionAccessPoints.html)
// in the Amazon S3 User Guide. This request is asynchronous, meaning that you
// might receive a response before the command has completed. When this request
// provides a response, it provides a token that you can use to monitor the status
// of the request with DescribeMultiRegionAccessPointOperation . The following
// actions are related to DeleteMultiRegionAccessPoint :
//   - CreateMultiRegionAccessPoint (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateMultiRegionAccessPoint.html)
//   - DescribeMultiRegionAccessPointOperation (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html)
//   - GetMultiRegionAccessPoint (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPoint.html)
//   - ListMultiRegionAccessPoints (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListMultiRegionAccessPoints.html)
func (c *Client) DeleteMultiRegionAccessPoint(ctx context.Context, params *DeleteMultiRegionAccessPointInput, optFns ...func(*Options)) (*DeleteMultiRegionAccessPointOutput, error) {
	if params == nil {
		params = &DeleteMultiRegionAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMultiRegionAccessPoint", params, optFns, c.addOperationDeleteMultiRegionAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMultiRegionAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMultiRegionAccessPointInput struct {

	// The Amazon Web Services account ID for the owner of the Multi-Region Access
	// Point.
	//
	// This member is required.
	AccountId *string

	// An idempotency token used to identify the request and guarantee that requests
	// are unique.
	//
	// This member is required.
	ClientToken *string

	// A container element containing details about the Multi-Region Access Point.
	//
	// This member is required.
	Details *types.DeleteMultiRegionAccessPointInput

	noSmithyDocumentSerde
}

func (*DeleteMultiRegionAccessPointInput) operationName() string {
	return "DeleteMultiRegionAccessPoint"
}

type DeleteMultiRegionAccessPointOutput struct {

	// The request token associated with the request. You can use this token with
	// DescribeMultiRegionAccessPointOperation (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html)
	// to determine the status of asynchronous requests.
	RequestTokenARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMultiRegionAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteMultiRegionAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteMultiRegionAccessPoint{}, middleware.After)
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opDeleteMultiRegionAccessPointMiddleware(stack); err != nil {
		return err
	}
	if err = addDeleteMultiRegionAccessPointResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opDeleteMultiRegionAccessPointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteMultiRegionAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMultiRegionAccessPoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addDeleteMultiRegionAccessPointUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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

type endpointPrefix_opDeleteMultiRegionAccessPointMiddleware struct {
}

func (*endpointPrefix_opDeleteMultiRegionAccessPointMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteMultiRegionAccessPointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteMultiRegionAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteMultiRegionAccessPointMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteMultiRegionAccessPointMiddleware{}, `OperationSerializer`, middleware.After)
}

type idempotencyToken_initializeOpDeleteMultiRegionAccessPoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteMultiRegionAccessPoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteMultiRegionAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteMultiRegionAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteMultiRegionAccessPointInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteMultiRegionAccessPointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteMultiRegionAccessPoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteMultiRegionAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteMultiRegionAccessPoint",
	}
}

func copyDeleteMultiRegionAccessPointInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteMultiRegionAccessPointInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteMultiRegionAccessPointInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillDeleteMultiRegionAccessPointAccountID(input interface{}, v string) error {
	in := input.(*DeleteMultiRegionAccessPointInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteMultiRegionAccessPointUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyDeleteMultiRegionAccessPointInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

type opDeleteMultiRegionAccessPointResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDeleteMultiRegionAccessPointResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDeleteMultiRegionAccessPointResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteMultiRegionAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.AccountId = input.AccountId

	params.RequiresAccountId = ptr.Bool(true)

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
			signingName := "s3"
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
				signingName = "s3"
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
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	ctx = smithyhttp.DisableEndpointHostPrefix(ctx, true)

	return next.HandleSerialize(ctx, in)
}

func addDeleteMultiRegionAccessPointResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDeleteMultiRegionAccessPointResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:     options.BaseEndpoint,
			UseArnRegion: options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
