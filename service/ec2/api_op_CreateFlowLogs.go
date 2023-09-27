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
)

// Creates one or more flow logs to capture information about IP traffic for a
// specific network interface, subnet, or VPC. Flow log data for a monitored
// network interface is recorded as flow log records, which are log events
// consisting of fields that describe the traffic flow. For more information, see
// Flow log records (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records)
// in the Amazon Virtual Private Cloud User Guide. When publishing to CloudWatch
// Logs, flow log records are published to a log group, and each network interface
// has a unique log stream in the log group. When publishing to Amazon S3, flow log
// records for all of the monitored network interfaces are published to a single
// log file object that is stored in the specified bucket. For more information,
// see VPC Flow Logs (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateFlowLogs(ctx context.Context, params *CreateFlowLogsInput, optFns ...func(*Options)) (*CreateFlowLogsOutput, error) {
	if params == nil {
		params = &CreateFlowLogsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlowLogs", params, optFns, c.addOperationCreateFlowLogsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlowLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowLogsInput struct {

	// The IDs of the resources to monitor. For example, if the resource type is VPC ,
	// specify the IDs of the VPCs. Constraints: Maximum of 25 for transit gateway
	// resource types. Maximum of 1000 for the other resource types.
	//
	// This member is required.
	ResourceIds []string

	// The type of resource to monitor.
	//
	// This member is required.
	ResourceType types.FlowLogsResourceType

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs across
	// accounts.
	DeliverCrossAccountRole *string

	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs to a
	// CloudWatch Logs log group in your account. This parameter is required if the
	// destination type is cloud-watch-logs and unsupported otherwise.
	DeliverLogsPermissionArn *string

	// The destination options.
	DestinationOptions *types.DestinationOptionsRequest

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The destination for the flow log data. The meaning of this parameter depends on
	// the destination type.
	//   - If the destination type is cloud-watch-logs , specify the ARN of a
	//   CloudWatch Logs log group. For example:
	//   arn:aws:logs:region:account_id:log-group:my_group Alternatively, use the
	//   LogGroupName parameter.
	//   - If the destination type is s3 , specify the ARN of an S3 bucket. For
	//   example: arn:aws:s3:::my_bucket/my_subfolder/ The subfolder is optional. Note
	//   that you can't use AWSLogs as a subfolder name.
	//   - If the destination type is kinesis-data-firehose , specify the ARN of a
	//   Kinesis Data Firehose delivery stream. For example:
	//   arn:aws:firehose:region:account_id:deliverystream:my_stream
	LogDestination *string

	// The type of destination for the flow log data. Default: cloud-watch-logs
	LogDestinationType types.LogDestinationType

	// The fields to include in the flow log record. List the fields in the order in
	// which they should appear. If you omit this parameter, the flow log is created
	// using the default format. If you specify this parameter, you must include at
	// least one field. For more information about the available fields, see Flow log
	// records (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records)
	// in the Amazon VPC User Guide or Transit Gateway Flow Log records (https://docs.aws.amazon.com/vpc/latest/tgw/tgw-flow-logs.html#flow-log-records)
	// in the Amazon Web Services Transit Gateway Guide. Specify the fields using the
	// ${field-id} format, separated by spaces.
	LogFormat *string

	// The name of a new or existing CloudWatch Logs log group where Amazon EC2
	// publishes your flow logs. This parameter is valid only if the destination type
	// is cloud-watch-logs .
	LogGroupName *string

	// The maximum interval of time during which a flow of packets is captured and
	// aggregated into a flow log record. The possible values are 60 seconds (1 minute)
	// or 600 seconds (10 minutes). This parameter must be 60 seconds for transit
	// gateway resource types. When a network interface is attached to a Nitro-based
	// instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances)
	// , the aggregation interval is always 60 seconds or less, regardless of the value
	// that you specify. Default: 600
	MaxAggregationInterval *int32

	// The tags to apply to the flow logs.
	TagSpecifications []types.TagSpecification

	// The type of traffic to monitor (accepted traffic, rejected traffic, or all
	// traffic). This parameter is not supported for transit gateway resource types. It
	// is required for the other resource types.
	TrafficType types.TrafficType

	noSmithyDocumentSerde
}

func (*CreateFlowLogsInput) operationName() string {
	return "CreateFlowLogs"
}

type CreateFlowLogsOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// The IDs of the flow logs.
	FlowLogIds []string

	// Information about the flow logs that could not be created successfully.
	Unsuccessful []types.UnsuccessfulItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFlowLogsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateFlowLogs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateFlowLogs{}, middleware.After)
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
	if err = addCreateFlowLogsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFlowLogsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlowLogs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFlowLogs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateFlowLogs",
	}
}

type opCreateFlowLogsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateFlowLogsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateFlowLogsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addCreateFlowLogsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateFlowLogsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
