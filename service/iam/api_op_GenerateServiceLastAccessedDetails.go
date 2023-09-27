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

// Generates a report that includes details about when an IAM resource (user,
// group, role, or policy) was last used in an attempt to access Amazon Web
// Services services. Recent activity usually appears within four hours. IAM
// reports activity for at least the last 400 days, or less if your Region began
// supporting this feature within the last year. For more information, see Regions
// where data is tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period)
// . The service last accessed data includes all attempts to access an Amazon Web
// Services API, not just the successful ones. This includes all attempts that were
// made using the Amazon Web Services Management Console, the Amazon Web Services
// API through any of the SDKs, or any of the command line tools. An unexpected
// entry in the service last accessed data does not mean that your account has been
// compromised, because the request might have been denied. Refer to your
// CloudTrail logs as the authoritative source for information about all API calls
// and whether they were successful or denied access. For more information, see
// Logging IAM events with CloudTrail (https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html)
// in the IAM User Guide. The GenerateServiceLastAccessedDetails operation returns
// a JobId . Use this parameter in the following operations to retrieve the
// following details from your report:
//   - GetServiceLastAccessedDetails – Use this operation for users, groups, roles,
//     or policies to list every Amazon Web Services service that the resource could
//     access using permissions policies. For each service, the response includes
//     information about the most recent access attempt. The JobId returned by
//     GenerateServiceLastAccessedDetail must be used by the same role within a
//     session, or by the same user when used to call GetServiceLastAccessedDetail .
//   - GetServiceLastAccessedDetailsWithEntities – Use this operation for groups
//     and policies to list information about the associated entities (users or roles)
//     that attempted to access a specific Amazon Web Services service.
//
// To check the status of the GenerateServiceLastAccessedDetails request, use the
// JobId parameter in the same operations and test the JobStatus response
// parameter. For additional information about the permissions policies that allow
// an identity (user, group, or role) to access specific services, use the
// ListPoliciesGrantingServiceAccess operation. Service last accessed data does not
// use other policy types when determining whether a resource could access a
// service. These other policy types include resource-based policies, access
// control lists, Organizations policies, IAM permissions boundaries, and STS
// assume role policies. It only applies permissions policy logic. For more about
// the evaluation of policy types, see Evaluating policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics)
// in the IAM User Guide. For more information about service and action last
// accessed data, see Reducing permissions using service last accessed data (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide.
func (c *Client) GenerateServiceLastAccessedDetails(ctx context.Context, params *GenerateServiceLastAccessedDetailsInput, optFns ...func(*Options)) (*GenerateServiceLastAccessedDetailsOutput, error) {
	if params == nil {
		params = &GenerateServiceLastAccessedDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateServiceLastAccessedDetails", params, optFns, c.addOperationGenerateServiceLastAccessedDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateServiceLastAccessedDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateServiceLastAccessedDetailsInput struct {

	// The ARN of the IAM resource (user, group, role, or managed policy) used to
	// generate information about when the resource was last used in an attempt to
	// access an Amazon Web Services service.
	//
	// This member is required.
	Arn *string

	// The level of detail that you want to generate. You can specify whether you want
	// to generate information about the last attempt to access services or actions. If
	// you specify service-level granularity, this operation generates only service
	// data. If you specify action-level granularity, it generates service and action
	// data. If you don't include this optional parameter, the operation generates
	// service data.
	Granularity types.AccessAdvisorUsageGranularityType

	noSmithyDocumentSerde
}

func (*GenerateServiceLastAccessedDetailsInput) operationName() string {
	return "GenerateServiceLastAccessedDetails"
}

type GenerateServiceLastAccessedDetailsOutput struct {

	// The JobId that you can use in the GetServiceLastAccessedDetails or
	// GetServiceLastAccessedDetailsWithEntities operations. The JobId returned by
	// GenerateServiceLastAccessedDetail must be used by the same role within a
	// session, or by the same user when used to call GetServiceLastAccessedDetail .
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateServiceLastAccessedDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGenerateServiceLastAccessedDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGenerateServiceLastAccessedDetails{}, middleware.After)
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
	if err = addGenerateServiceLastAccessedDetailsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGenerateServiceLastAccessedDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateServiceLastAccessedDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateServiceLastAccessedDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GenerateServiceLastAccessedDetails",
	}
}

type opGenerateServiceLastAccessedDetailsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGenerateServiceLastAccessedDetailsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGenerateServiceLastAccessedDetailsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addGenerateServiceLastAccessedDetailsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGenerateServiceLastAccessedDetailsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
