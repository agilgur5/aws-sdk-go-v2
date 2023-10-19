// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a target group. For more information, see the following:
//   - Target groups for your Application Load Balancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html)
//   - Target groups for your Network Load Balancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html)
//   - Target groups for your Gateway Load Balancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/target-groups.html)
//
// This operation is idempotent, which means that it completes at most one time.
// If you attempt to create multiple target groups with the same settings, each
// call succeeds.
func (c *Client) CreateTargetGroup(ctx context.Context, params *CreateTargetGroupInput, optFns ...func(*Options)) (*CreateTargetGroupOutput, error) {
	if params == nil {
		params = &CreateTargetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTargetGroup", params, optFns, c.addOperationCreateTargetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTargetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTargetGroupInput struct {

	// The name of the target group. This name must be unique per region per account,
	// can have a maximum of 32 characters, must contain only alphanumeric characters
	// or hyphens, and must not begin or end with a hyphen.
	//
	// This member is required.
	Name *string

	// Indicates whether health checks are enabled. If the target type is lambda ,
	// health checks are disabled by default but can be enabled. If the target type is
	// instance , ip , or alb , health checks are always enabled and cannot be disabled.
	HealthCheckEnabled *bool

	// The approximate amount of time, in seconds, between health checks of an
	// individual target. The range is 5-300. If the target group protocol is TCP, TLS,
	// UDP, TCP_UDP, HTTP or HTTPS, the default is 30 seconds. If the target group
	// protocol is GENEVE, the default is 10 seconds. If the target type is lambda ,
	// the default is 35 seconds.
	HealthCheckIntervalSeconds *int32

	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC
	// protocol version] The path of a custom health check method with the format
	// /package.service/method. The default is /Amazon Web Services.ALB/healthcheck.
	HealthCheckPath *string

	// The port the load balancer uses when performing health checks on targets. If
	// the protocol is HTTP, HTTPS, TCP, TLS, UDP, or TCP_UDP, the default is
	// traffic-port , which is the port on which each target receives traffic from the
	// load balancer. If the protocol is GENEVE, the default is port 80.
	HealthCheckPort *string

	// The protocol the load balancer uses when performing health checks on targets.
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers
	// and Gateway Load Balancers, the default is TCP. The TCP protocol is not
	// supported for health checks if the protocol of the target group is HTTP or
	// HTTPS. The GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health
	// checks.
	HealthCheckProtocol types.ProtocolEnum

	// The amount of time, in seconds, during which no response from a target means a
	// failed health check. The range is 2–120 seconds. For target groups with a
	// protocol of HTTP, the default is 6 seconds. For target groups with a protocol of
	// TCP, TLS or HTTPS, the default is 10 seconds. For target groups with a protocol
	// of GENEVE, the default is 5 seconds. If the target type is lambda , the default
	// is 30 seconds.
	HealthCheckTimeoutSeconds *int32

	// The number of consecutive health check successes required before considering a
	// target healthy. The range is 2-10. If the target group protocol is TCP, TCP_UDP,
	// UDP, TLS, HTTP or HTTPS, the default is 5. For target groups with a protocol of
	// GENEVE, the default is 5. If the target type is lambda , the default is 5.
	HealthyThresholdCount *int32

	// The type of IP address used for this target group. The possible values are ipv4
	// and ipv6 . This is an optional parameter. If not specified, the IP address type
	// defaults to ipv4 .
	IpAddressType types.TargetGroupIpAddressTypeEnum

	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a
	// successful response from a target. For target groups with a protocol of TCP,
	// TCP_UDP, UDP or TLS the range is 200-599. For target groups with a protocol of
	// HTTP or HTTPS, the range is 200-499. For target groups with a protocol of
	// GENEVE, the range is 200-399.
	Matcher *types.Matcher

	// The port on which the targets receive traffic. This port is used unless you
	// specify a port override when registering the target. If the target is a Lambda
	// function, this parameter does not apply. If the protocol is GENEVE, the
	// supported port is 6081.
	Port *int32

	// The protocol to use for routing traffic to the targets. For Application Load
	// Balancers, the supported protocols are HTTP and HTTPS. For Network Load
	// Balancers, the supported protocols are TCP, TLS, UDP, or TCP_UDP. For Gateway
	// Load Balancers, the supported protocol is GENEVE. A TCP_UDP listener must be
	// associated with a TCP_UDP target group. If the target is a Lambda function, this
	// parameter does not apply.
	Protocol types.ProtocolEnum

	// [HTTP/HTTPS protocol] The protocol version. Specify GRPC to send requests to
	// targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The
	// default is HTTP1 , which sends requests to targets using HTTP/1.1.
	ProtocolVersion *string

	// The tags to assign to the target group.
	Tags []types.Tag

	// The type of target that you must specify when registering targets with this
	// target group. You can't specify targets for a target group using more than one
	// target type.
	//   - instance - Register targets by instance ID. This is the default value.
	//   - ip - Register targets by IP address. You can specify IP addresses from the
	//   subnets of the virtual private cloud (VPC) for the target group, the RFC 1918
	//   range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range
	//   (100.64.0.0/10). You can't specify publicly routable IP addresses.
	//   - lambda - Register a single Lambda function as a target.
	//   - alb - Register a single Application Load Balancer as a target.
	TargetType types.TargetTypeEnum

	// The number of consecutive health check failures required before considering a
	// target unhealthy. The range is 2-10. If the target group protocol is TCP,
	// TCP_UDP, UDP, TLS, HTTP or HTTPS, the default is 2. For target groups with a
	// protocol of GENEVE, the default is 2. If the target type is lambda , the default
	// is 5.
	UnhealthyThresholdCount *int32

	// The identifier of the virtual private cloud (VPC). If the target is a Lambda
	// function, this parameter does not apply. Otherwise, this parameter is required.
	VpcId *string

	noSmithyDocumentSerde
}

type CreateTargetGroupOutput struct {

	// Information about the target group.
	TargetGroups []types.TargetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTargetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateTargetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateTargetGroup{}, middleware.After)
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
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTargetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTargetGroup(options.Region), middleware.Before); err != nil {
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
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "CreateTargetGroup",
		options:   options,
	}, "ResolveEndpointV2", middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&signRequestMiddleware{}, middleware.Before)
	if err != nil {
		return err
	}

	err = stack.Finalize.Add(&getIdentityMiddleware{
		options: options,
	}, middleware.Before)
	if err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateTargetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateTargetGroup",
	}
}
