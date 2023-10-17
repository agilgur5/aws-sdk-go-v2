// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified attribute for a load balancer. You can only update one
// attribute at a time. The update load balancer attribute operation supports
// tag-based access control via resource tags applied to the resource identified by
// load balancer name . For more information, see the Amazon Lightsail Developer
// Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags)
// .
func (c *Client) UpdateLoadBalancerAttribute(ctx context.Context, params *UpdateLoadBalancerAttributeInput, optFns ...func(*Options)) (*UpdateLoadBalancerAttributeOutput, error) {
	if params == nil {
		params = &UpdateLoadBalancerAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLoadBalancerAttribute", params, optFns, c.addOperationUpdateLoadBalancerAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLoadBalancerAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLoadBalancerAttributeInput struct {

	// The name of the attribute you want to update.
	//
	// This member is required.
	AttributeName types.LoadBalancerAttributeName

	// The value that you want to specify for the attribute name. The following values
	// are supported depending on what you specify for the attributeName request
	// parameter:
	//   - If you specify HealthCheckPath for the attributeName request parameter, then
	//   the attributeValue request parameter must be the path to ping on the target
	//   (for example, /weather/us/wa/seattle ).
	//   - If you specify SessionStickinessEnabled for the attributeName request
	//   parameter, then the attributeValue request parameter must be true to activate
	//   session stickiness or false to deactivate session stickiness.
	//   - If you specify SessionStickiness_LB_CookieDurationSeconds for the
	//   attributeName request parameter, then the attributeValue request parameter
	//   must be an interger that represents the cookie duration in seconds.
	//   - If you specify HttpsRedirectionEnabled for the attributeName request
	//   parameter, then the attributeValue request parameter must be true to activate
	//   HTTP to HTTPS redirection or false to deactivate HTTP to HTTPS redirection.
	//   - If you specify TlsPolicyName for the attributeName request parameter, then
	//   the attributeValue request parameter must be the name of the TLS policy. Use
	//   the GetLoadBalancerTlsPolicies (https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetLoadBalancerTlsPolicies.html)
	//   action to get a list of TLS policy names that you can specify.
	//
	// This member is required.
	AttributeValue *string

	// The name of the load balancer that you want to modify (e.g., my-load-balancer .
	//
	// This member is required.
	LoadBalancerName *string

	noSmithyDocumentSerde
}

type UpdateLoadBalancerAttributeOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLoadBalancerAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLoadBalancerAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLoadBalancerAttribute{}, middleware.After)
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
	if err = addOpUpdateLoadBalancerAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLoadBalancerAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLoadBalancerAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "UpdateLoadBalancerAttribute",
	}
}
