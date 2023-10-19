// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the attributes of the specified load balancer. You can modify the load
// balancer attributes, such as AccessLogs , ConnectionDraining , and
// CrossZoneLoadBalancing by either enabling or disabling them. Or, you can modify
// the load balancer attribute ConnectionSettings by specifying an idle connection
// timeout value for your load balancer. For more information, see the following in
// the Classic Load Balancers Guide:
//   - Cross-Zone Load Balancing (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html)
//   - Connection Draining (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html)
//   - Access Logs (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/access-log-collection.html)
//   - Idle Connection Timeout (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html)
func (c *Client) ModifyLoadBalancerAttributes(ctx context.Context, params *ModifyLoadBalancerAttributesInput, optFns ...func(*Options)) (*ModifyLoadBalancerAttributesOutput, error) {
	if params == nil {
		params = &ModifyLoadBalancerAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyLoadBalancerAttributes", params, optFns, c.addOperationModifyLoadBalancerAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyLoadBalancerAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyLoadBalancerAttributes.
type ModifyLoadBalancerAttributesInput struct {

	// The attributes for the load balancer.
	//
	// This member is required.
	LoadBalancerAttributes *types.LoadBalancerAttributes

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	noSmithyDocumentSerde
}

// Contains the output of ModifyLoadBalancerAttributes.
type ModifyLoadBalancerAttributesOutput struct {

	// Information about the load balancer attributes.
	LoadBalancerAttributes *types.LoadBalancerAttributes

	// The name of the load balancer.
	LoadBalancerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyLoadBalancerAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyLoadBalancerAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyLoadBalancerAttributes{}, middleware.After)
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
	if err = addOpModifyLoadBalancerAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyLoadBalancerAttributes(options.Region), middleware.Before); err != nil {
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
		operation: "ModifyLoadBalancerAttributes",
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

func newServiceMetadataMiddleware_opModifyLoadBalancerAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "ModifyLoadBalancerAttributes",
	}
}
