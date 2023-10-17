// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the description for the specified firewall. Use the description to
// help you identify the firewall when you're working with it.
func (c *Client) UpdateFirewallDescription(ctx context.Context, params *UpdateFirewallDescriptionInput, optFns ...func(*Options)) (*UpdateFirewallDescriptionOutput, error) {
	if params == nil {
		params = &UpdateFirewallDescriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFirewallDescription", params, optFns, c.addOperationUpdateFirewallDescriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFirewallDescriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFirewallDescriptionInput struct {

	// The new description for the firewall. If you omit this setting, Network
	// Firewall removes the description for the firewall.
	Description *string

	// The Amazon Resource Name (ARN) of the firewall. You must specify the ARN or the
	// name, and you can specify both.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it. You must specify the ARN or the name, and you can specify
	// both.
	FirewallName *string

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request. To make an
	// unconditional change to the firewall, omit the token in your update request.
	// Without the token, Network Firewall performs your updates regardless of whether
	// the firewall has changed since you last retrieved it. To make a conditional
	// change to the firewall, provide the token in your update request. Network
	// Firewall uses the token to ensure that the firewall hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an
	// InvalidTokenException . If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	noSmithyDocumentSerde
}

type UpdateFirewallDescriptionOutput struct {

	// A description of the firewall.
	Description *string

	// The Amazon Resource Name (ARN) of the firewall.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it.
	FirewallName *string

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request. To make an
	// unconditional change to the firewall, omit the token in your update request.
	// Without the token, Network Firewall performs your updates regardless of whether
	// the firewall has changed since you last retrieved it. To make a conditional
	// change to the firewall, provide the token in your update request. Network
	// Firewall uses the token to ensure that the firewall hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an
	// InvalidTokenException . If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFirewallDescriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateFirewallDescription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateFirewallDescription{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFirewallDescription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFirewallDescription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "UpdateFirewallDescription",
	}
}
