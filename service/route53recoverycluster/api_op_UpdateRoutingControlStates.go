// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycluster

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set multiple routing control states. You can set the value for each state to be
// ON or OFF. When the state is ON, traffic flows to a cell. When it's OFF, traffic
// does not flow. With Route 53 ARC, you can add safety rules for routing controls,
// which are safeguards for routing control state updates that help prevent
// unexpected outcomes, like fail open traffic routing. However, there are
// scenarios when you might want to bypass the routing control safeguards that are
// enforced with safety rules that you've configured. For example, you might want
// to fail over quickly for disaster recovery, and one or more safety rules might
// be unexpectedly preventing you from updating a routing control state to reroute
// traffic. In a "break glass" scenario like this, you can override one or more
// safety rules to change a routing control state and fail over your application.
// The SafetyRulesToOverride property enables you override one or more safety
// rules and update routing control states. For more information, see Override
// safety rules to reroute traffic (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.override-safety-rule.html)
// in the Amazon Route 53 Application Recovery Controller Developer Guide. You must
// specify Regional endpoints when you work with API cluster operations to get or
// update routing control states in Route 53 ARC. To see a code example for getting
// a routing control state, including accessing Regional cluster endpoints in
// sequence, see API examples (https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples_actions.html)
// in the Amazon Route 53 Application Recovery Controller Developer Guide.
//   - Viewing and updating routing control states (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html)
//   - Working with routing controls overall (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html)
func (c *Client) UpdateRoutingControlStates(ctx context.Context, params *UpdateRoutingControlStatesInput, optFns ...func(*Options)) (*UpdateRoutingControlStatesOutput, error) {
	if params == nil {
		params = &UpdateRoutingControlStatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoutingControlStates", params, optFns, c.addOperationUpdateRoutingControlStatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoutingControlStatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoutingControlStatesInput struct {

	// A set of routing control entries that you want to update.
	//
	// This member is required.
	UpdateRoutingControlStateEntries []types.UpdateRoutingControlStateEntry

	// The Amazon Resource Names (ARNs) for the safety rules that you want to override
	// when you're updating routing control states. You can override one safety rule or
	// multiple safety rules by including one or more ARNs, separated by commas. For
	// more information, see Override safety rules to reroute traffic (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.override-safety-rule.html)
	// in the Amazon Route 53 Application Recovery Controller Developer Guide.
	SafetyRulesToOverride []string

	noSmithyDocumentSerde
}

type UpdateRoutingControlStatesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRoutingControlStatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateRoutingControlStates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateRoutingControlStates{}, middleware.After)
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
	if err = addOpUpdateRoutingControlStatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoutingControlStates(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateRoutingControlStates",
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

func newServiceMetadataMiddleware_opUpdateRoutingControlStates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-cluster",
		OperationName: "UpdateRoutingControlStates",
	}
}
