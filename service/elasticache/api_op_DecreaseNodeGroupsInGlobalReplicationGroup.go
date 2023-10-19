// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Decreases the number of node groups in a Global datastore
func (c *Client) DecreaseNodeGroupsInGlobalReplicationGroup(ctx context.Context, params *DecreaseNodeGroupsInGlobalReplicationGroupInput, optFns ...func(*Options)) (*DecreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &DecreaseNodeGroupsInGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DecreaseNodeGroupsInGlobalReplicationGroup", params, optFns, c.addOperationDecreaseNodeGroupsInGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecreaseNodeGroupsInGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecreaseNodeGroupsInGlobalReplicationGroupInput struct {

	// Indicates that the shard reconfiguration process begins immediately. At
	// present, the only permitted value for this parameter is true.
	//
	// This member is required.
	ApplyImmediately bool

	// The name of the Global datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string

	// The number of node groups (shards) that results from the modification of the
	// shard configuration
	//
	// This member is required.
	NodeGroupCount int32

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// GlobalNodeGroupsToRemove is a list of NodeGroupIds to remove from the cluster.
	// ElastiCache for Redis will attempt to remove all node groups listed by
	// GlobalNodeGroupsToRemove from the cluster.
	GlobalNodeGroupsToRemove []string

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// GlobalNodeGroupsToRetain is a list of NodeGroupIds to retain from the cluster.
	// ElastiCache for Redis will attempt to retain all node groups listed by
	// GlobalNodeGroupsToRetain from the cluster.
	GlobalNodeGroupsToRetain []string

	noSmithyDocumentSerde
}

type DecreaseNodeGroupsInGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different Amazon region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.
	//   - The GlobalReplicationGroupIdSuffix represents the name of the Global
	//   datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDecreaseNodeGroupsInGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDecreaseNodeGroupsInGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDecreaseNodeGroupsInGlobalReplicationGroup{}, middleware.After)
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
	if err = addOpDecreaseNodeGroupsInGlobalReplicationGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDecreaseNodeGroupsInGlobalReplicationGroup(options.Region), middleware.Before); err != nil {
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
		operation: "DecreaseNodeGroupsInGlobalReplicationGroup",
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

func newServiceMetadataMiddleware_opDecreaseNodeGroupsInGlobalReplicationGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DecreaseNodeGroupsInGlobalReplicationGroup",
	}
}
