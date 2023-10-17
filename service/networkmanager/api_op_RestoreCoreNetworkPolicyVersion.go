// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Restores a previous policy version as a new, immutable version of a core
// network policy. A subsequent change set is created showing the differences
// between the LIVE policy and restored policy.
func (c *Client) RestoreCoreNetworkPolicyVersion(ctx context.Context, params *RestoreCoreNetworkPolicyVersionInput, optFns ...func(*Options)) (*RestoreCoreNetworkPolicyVersionOutput, error) {
	if params == nil {
		params = &RestoreCoreNetworkPolicyVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreCoreNetworkPolicyVersion", params, optFns, c.addOperationRestoreCoreNetworkPolicyVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreCoreNetworkPolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreCoreNetworkPolicyVersionInput struct {

	// The ID of a core network.
	//
	// This member is required.
	CoreNetworkId *string

	// The ID of the policy version to restore.
	//
	// This member is required.
	PolicyVersionId *int32

	noSmithyDocumentSerde
}

type RestoreCoreNetworkPolicyVersionOutput struct {

	// Describes the restored core network policy.
	CoreNetworkPolicy *types.CoreNetworkPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreCoreNetworkPolicyVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRestoreCoreNetworkPolicyVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRestoreCoreNetworkPolicyVersion{}, middleware.After)
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
	if err = addOpRestoreCoreNetworkPolicyVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreCoreNetworkPolicyVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreCoreNetworkPolicyVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "RestoreCoreNetworkPolicyVersion",
	}
}
