// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A complex type that contains settings for encryption of your firewall resources.
func (c *Client) UpdateFirewallEncryptionConfiguration(ctx context.Context, params *UpdateFirewallEncryptionConfigurationInput, optFns ...func(*Options)) (*UpdateFirewallEncryptionConfigurationOutput, error) {
	if params == nil {
		params = &UpdateFirewallEncryptionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFirewallEncryptionConfiguration", params, optFns, c.addOperationUpdateFirewallEncryptionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFirewallEncryptionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFirewallEncryptionConfigurationInput struct {

	// A complex type that contains optional Amazon Web Services Key Management
	// Service (KMS) encryption settings for your Network Firewall resources. Your data
	// is encrypted by default with an Amazon Web Services owned key that Amazon Web
	// Services owns and manages for you. You can use either the Amazon Web Services
	// owned key, or provide your own customer managed key. To learn more about KMS
	// encryption of your Network Firewall resources, see Encryption at rest with
	// Amazon Web Services Key Managment Service (https://docs.aws.amazon.com/kms/latest/developerguide/kms-encryption-at-rest.html)
	// in the Network Firewall Developer Guide.
	EncryptionConfiguration *types.EncryptionConfiguration

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

	noSmithyDocumentSerde
}

type UpdateFirewallEncryptionConfigurationOutput struct {

	// A complex type that contains optional Amazon Web Services Key Management
	// Service (KMS) encryption settings for your Network Firewall resources. Your data
	// is encrypted by default with an Amazon Web Services owned key that Amazon Web
	// Services owns and manages for you. You can use either the Amazon Web Services
	// owned key, or provide your own customer managed key. To learn more about KMS
	// encryption of your Network Firewall resources, see Encryption at rest with
	// Amazon Web Services Key Managment Service (https://docs.aws.amazon.com/kms/latest/developerguide/kms-encryption-at-rest.html)
	// in the Network Firewall Developer Guide.
	EncryptionConfiguration *types.EncryptionConfiguration

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

func (c *Client) addOperationUpdateFirewallEncryptionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateFirewallEncryptionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateFirewallEncryptionConfiguration{}, middleware.After)
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
	if err = addOpUpdateFirewallEncryptionConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFirewallEncryptionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFirewallEncryptionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "UpdateFirewallEncryptionConfiguration",
	}
}
