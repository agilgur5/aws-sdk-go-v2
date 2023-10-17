// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified capacity provider. The FARGATE and FARGATE_SPOT capacity
// providers are reserved and can't be deleted. You can disassociate them from a
// cluster using either the PutClusterCapacityProviders API or by deleting the
// cluster. Prior to a capacity provider being deleted, the capacity provider must
// be removed from the capacity provider strategy from all services. The
// UpdateService API can be used to remove a capacity provider from a service's
// capacity provider strategy. When updating a service, the forceNewDeployment
// option can be used to ensure that any tasks using the Amazon EC2 instance
// capacity provided by the capacity provider are transitioned to use the capacity
// from the remaining capacity providers. Only capacity providers that aren't
// associated with a cluster can be deleted. To remove a capacity provider from a
// cluster, you can either use PutClusterCapacityProviders or delete the cluster.
func (c *Client) DeleteCapacityProvider(ctx context.Context, params *DeleteCapacityProviderInput, optFns ...func(*Options)) (*DeleteCapacityProviderOutput, error) {
	if params == nil {
		params = &DeleteCapacityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCapacityProvider", params, optFns, c.addOperationDeleteCapacityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCapacityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCapacityProviderInput struct {

	// The short name or full Amazon Resource Name (ARN) of the capacity provider to
	// delete.
	//
	// This member is required.
	CapacityProvider *string

	noSmithyDocumentSerde
}

type DeleteCapacityProviderOutput struct {

	// The details of the capacity provider.
	CapacityProvider *types.CapacityProvider

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCapacityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCapacityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCapacityProvider{}, middleware.After)
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
	if err = addOpDeleteCapacityProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCapacityProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCapacityProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeleteCapacityProvider",
	}
}
