// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detaches a specified Elastic Load Balancing instance from its layer. Required
// Permissions: To use this action, an IAM user must have a Manage permissions
// level for the stack, or an attached policy that explicitly grants permissions.
// For more information on user permissions, see Managing User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html)
// .
func (c *Client) DetachElasticLoadBalancer(ctx context.Context, params *DetachElasticLoadBalancerInput, optFns ...func(*Options)) (*DetachElasticLoadBalancerOutput, error) {
	if params == nil {
		params = &DetachElasticLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachElasticLoadBalancer", params, optFns, c.addOperationDetachElasticLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachElasticLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachElasticLoadBalancerInput struct {

	// The Elastic Load Balancing instance's name.
	//
	// This member is required.
	ElasticLoadBalancerName *string

	// The ID of the layer that the Elastic Load Balancing instance is attached to.
	//
	// This member is required.
	LayerId *string

	noSmithyDocumentSerde
}

type DetachElasticLoadBalancerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachElasticLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetachElasticLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachElasticLoadBalancer{}, middleware.After)
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
	if err = addOpDetachElasticLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachElasticLoadBalancer(options.Region), middleware.Before); err != nil {
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
		operation: "DetachElasticLoadBalancer",
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

func newServiceMetadataMiddleware_opDetachElasticLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DetachElasticLoadBalancer",
	}
}
