// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an EC2 Fleet that contains the configuration information for On-Demand
// Instances and Spot Instances. Instances are launched immediately if there is
// available capacity. A single EC2 Fleet can include multiple launch
// specifications that vary by instance type, AMI, Availability Zone, or subnet.
// For more information, see EC2 Fleet (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet.html)
// in the Amazon EC2 User Guide.
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	if params == nil {
		params = &CreateFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleet", params, optFns, c.addOperationCreateFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetInput struct {

	// The configuration for the EC2 Fleet.
	//
	// This member is required.
	LaunchTemplateConfigs []types.FleetLaunchTemplateConfigRequest

	// The number of units to request.
	//
	// This member is required.
	TargetCapacitySpecification *types.TargetCapacitySpecificationRequest

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see Ensuring idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// Reserved.
	Context *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Indicates whether running instances should be terminated if the total target
	// capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.
	// Supported only for fleets of type maintain .
	ExcessCapacityTerminationPolicy types.FleetExcessCapacityTerminationPolicy

	// Describes the configuration of On-Demand Instances in an EC2 Fleet.
	OnDemandOptions *types.OnDemandOptionsRequest

	// Indicates whether EC2 Fleet should replace unhealthy Spot Instances. Supported
	// only for fleets of type maintain . For more information, see EC2 Fleet health
	// checks (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#ec2-fleet-health-checks)
	// in the Amazon EC2 User Guide.
	ReplaceUnhealthyInstances *bool

	// Describes the configuration of Spot Instances in an EC2 Fleet.
	SpotOptions *types.SpotOptionsRequest

	// The key-value pair for tagging the EC2 Fleet request on creation. For more
	// information, see Tag your resources (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-resources)
	// . If the fleet type is instant , specify a resource type of fleet to tag the
	// fleet or instance to tag the instances at launch. If the fleet type is maintain
	// or request , specify a resource type of fleet to tag the fleet. You cannot
	// specify a resource type of instance . To tag instances at launch, specify the
	// tags in a launch template (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template)
	// .
	TagSpecifications []types.TagSpecification

	// Indicates whether running instances should be terminated when the EC2 Fleet
	// expires.
	TerminateInstancesWithExpiration *bool

	// The fleet type. The default value is maintain .
	//   - maintain - The EC2 Fleet places an asynchronous request for your desired
	//   capacity, and continues to maintain your desired Spot capacity by replenishing
	//   interrupted Spot Instances.
	//   - request - The EC2 Fleet places an asynchronous one-time request for your
	//   desired capacity, but does submit Spot requests in alternative capacity pools if
	//   Spot capacity is unavailable, and does not maintain Spot capacity if Spot
	//   Instances are interrupted.
	//   - instant - The EC2 Fleet places a synchronous one-time request for your
	//   desired capacity, and returns errors for any instances that could not be
	//   launched.
	// For more information, see EC2 Fleet request types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-request-type.html)
	// in the Amazon EC2 User Guide.
	Type types.FleetType

	// The start date and time of the request, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request
	// immediately.
	ValidFrom *time.Time

	// The end date and time of the request, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or
	// able to fulfill the request. If no value is specified, the request remains until
	// you cancel it.
	ValidUntil *time.Time

	noSmithyDocumentSerde
}

func (*CreateFleetInput) operationName() string {
	return "CreateFleet"
}

type CreateFleetOutput struct {

	// Information about the instances that could not be launched by the fleet.
	// Supported only for fleets of type instant .
	Errors []types.CreateFleetError

	// The ID of the EC2 Fleet.
	FleetId *string

	// Information about the instances that were launched by the fleet. Supported only
	// for fleets of type instant .
	Instances []types.CreateFleetInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateFleet{}, middleware.After)
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
	if err = addCreateFleetResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateFleet",
	}
}

type opCreateFleetResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateFleetResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateFleetResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ec2"
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
				signingName = "ec2"
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
				v4aScheme.SigningName = aws.String("ec2")
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

func addCreateFleetResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateFleetResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
