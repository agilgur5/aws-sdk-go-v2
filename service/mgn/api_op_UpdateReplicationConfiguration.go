// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to update multiple ReplicationConfigurations by Source Server ID.
func (c *Client) UpdateReplicationConfiguration(ctx context.Context, params *UpdateReplicationConfigurationInput, optFns ...func(*Options)) (*UpdateReplicationConfigurationOutput, error) {
	if params == nil {
		params = &UpdateReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateReplicationConfiguration", params, optFns, c.addOperationUpdateReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReplicationConfigurationInput struct {

	// Update replication configuration Source Server ID request.
	//
	// This member is required.
	SourceServerID *string

	// Update replication configuration Account ID request.
	AccountID *string

	// Update replication configuration associate default Application Migration
	// Service Security group request.
	AssociateDefaultSecurityGroup *bool

	// Update replication configuration bandwidth throttling request.
	BandwidthThrottling int64

	// Update replication configuration create Public IP request.
	CreatePublicIP *bool

	// Update replication configuration data plane routing request.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Update replication configuration use default large Staging Disk type request.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Update replication configuration EBS encryption request.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Update replication configuration EBS encryption key ARN request.
	EbsEncryptionKeyArn *string

	// Update replication configuration name request.
	Name *string

	// Update replication configuration replicated disks request.
	ReplicatedDisks []types.ReplicationConfigurationReplicatedDisk

	// Update replication configuration Replication Server instance type request.
	ReplicationServerInstanceType *string

	// Update replication configuration Replication Server Security Groups IDs request.
	ReplicationServersSecurityGroupsIDs []string

	// Update replication configuration Staging Area subnet request.
	StagingAreaSubnetId *string

	// Update replication configuration Staging Area Tags request.
	StagingAreaTags map[string]string

	// Update replication configuration use dedicated Replication Server request.
	UseDedicatedReplicationServer *bool

	// Update replication configuration use Fips Endpoint.
	UseFipsEndpoint *bool

	noSmithyDocumentSerde
}

type UpdateReplicationConfigurationOutput struct {

	// Replication Configuration associate default Application Migration Service
	// Security Group.
	AssociateDefaultSecurityGroup *bool

	// Replication Configuration set bandwidth throttling.
	BandwidthThrottling int64

	// Replication Configuration create Public IP.
	CreatePublicIP *bool

	// Replication Configuration data plane routing.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Replication Configuration use default large Staging Disks.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Replication Configuration EBS encryption.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Replication Configuration EBS encryption key ARN.
	EbsEncryptionKeyArn *string

	// Replication Configuration name.
	Name *string

	// Replication Configuration replicated disks.
	ReplicatedDisks []types.ReplicationConfigurationReplicatedDisk

	// Replication Configuration Replication Server instance type.
	ReplicationServerInstanceType *string

	// Replication Configuration Replication Server Security Group IDs.
	ReplicationServersSecurityGroupsIDs []string

	// Replication Configuration Source Server ID.
	SourceServerID *string

	// Replication Configuration Staging Area subnet ID.
	StagingAreaSubnetId *string

	// Replication Configuration Staging Area tags.
	StagingAreaTags map[string]string

	// Replication Configuration use Dedicated Replication Server.
	UseDedicatedReplicationServer *bool

	// Replication Configuration use Fips Endpoint.
	UseFipsEndpoint *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateReplicationConfiguration{}, middleware.After)
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
	if err = addOpUpdateReplicationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReplicationConfiguration(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateReplicationConfiguration",
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

func newServiceMetadataMiddleware_opUpdateReplicationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "UpdateReplicationConfiguration",
	}
}
