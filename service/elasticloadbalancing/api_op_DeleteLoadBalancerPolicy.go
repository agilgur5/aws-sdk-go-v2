// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified policy from the specified load balancer. This policy must
// not be enabled for any listeners.
func (c *Client) DeleteLoadBalancerPolicy(ctx context.Context, params *DeleteLoadBalancerPolicyInput, optFns ...func(*Options)) (*DeleteLoadBalancerPolicyOutput, error) {
	if params == nil {
		params = &DeleteLoadBalancerPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLoadBalancerPolicy", params, optFns, c.addOperationDeleteLoadBalancerPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLoadBalancerPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeleteLoadBalancerPolicy.
type DeleteLoadBalancerPolicyInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The name of the policy.
	//
	// This member is required.
	PolicyName *string

	noSmithyDocumentSerde
}

// Contains the output of DeleteLoadBalancerPolicy.
type DeleteLoadBalancerPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteLoadBalancerPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteLoadBalancerPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteLoadBalancerPolicy{}, middleware.After)
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
	if err = addOpDeleteLoadBalancerPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLoadBalancerPolicy(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteLoadBalancerPolicy",
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

func newServiceMetadataMiddleware_opDeleteLoadBalancerPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DeleteLoadBalancerPolicy",
	}
}
