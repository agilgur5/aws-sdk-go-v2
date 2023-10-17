// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified Availability Zones from the set of Availability Zones for
// the specified load balancer in EC2-Classic or a default VPC. For load balancers
// in a non-default VPC, use DetachLoadBalancerFromSubnets . There must be at least
// one Availability Zone registered with a load balancer at all times. After an
// Availability Zone is removed, all instances registered with the load balancer
// that are in the removed Availability Zone go into the OutOfService state. Then,
// the load balancer attempts to equally balance the traffic among its remaining
// Availability Zones. For more information, see Add or Remove Availability Zones (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-az.html)
// in the Classic Load Balancers Guide.
func (c *Client) DisableAvailabilityZonesForLoadBalancer(ctx context.Context, params *DisableAvailabilityZonesForLoadBalancerInput, optFns ...func(*Options)) (*DisableAvailabilityZonesForLoadBalancerOutput, error) {
	if params == nil {
		params = &DisableAvailabilityZonesForLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableAvailabilityZonesForLoadBalancer", params, optFns, c.addOperationDisableAvailabilityZonesForLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableAvailabilityZonesForLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DisableAvailabilityZonesForLoadBalancer.
type DisableAvailabilityZonesForLoadBalancerInput struct {

	// The Availability Zones.
	//
	// This member is required.
	AvailabilityZones []string

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	noSmithyDocumentSerde
}

// Contains the output for DisableAvailabilityZonesForLoadBalancer.
type DisableAvailabilityZonesForLoadBalancerOutput struct {

	// The remaining Availability Zones for the load balancer.
	AvailabilityZones []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableAvailabilityZonesForLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableAvailabilityZonesForLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableAvailabilityZonesForLoadBalancer{}, middleware.After)
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
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpDisableAvailabilityZonesForLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableAvailabilityZonesForLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableAvailabilityZonesForLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DisableAvailabilityZonesForLoadBalancer",
	}
}
