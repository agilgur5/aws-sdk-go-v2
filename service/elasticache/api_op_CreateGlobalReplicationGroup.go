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

// Global Datastore for Redis offers fully managed, fast, reliable and secure
// cross-region replication. Using Global Datastore for Redis, you can create
// cross-region read replica clusters for ElastiCache for Redis to enable
// low-latency reads and disaster recovery across regions. For more information,
// see Replication Across Regions Using Global Datastore (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Redis-Global-Datastore.html)
// .
//   - The GlobalReplicationGroupIdSuffix is the name of the Global datastore.
//   - The PrimaryReplicationGroupId represents the name of the primary cluster
//     that accepts writes and will replicate updates to the secondary cluster.
func (c *Client) CreateGlobalReplicationGroup(ctx context.Context, params *CreateGlobalReplicationGroupInput, optFns ...func(*Options)) (*CreateGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &CreateGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGlobalReplicationGroup", params, optFns, c.addOperationCreateGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGlobalReplicationGroupInput struct {

	// The suffix name of a Global datastore. Amazon ElastiCache automatically applies
	// a prefix to the Global datastore ID when it is created. Each Amazon Region has
	// its own prefix. For instance, a Global datastore ID created in the US-West-1
	// region will begin with "dsdfu" along with the suffix name you provide. The
	// suffix, combined with the auto-generated prefix, guarantees uniqueness of the
	// Global datastore name across multiple regions. For a full list of Amazon Regions
	// and their respective Global datastore iD prefixes, see Using the Amazon CLI
	// with Global datastores  (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Redis-Global-Datastores-CLI.html)
	// .
	//
	// This member is required.
	GlobalReplicationGroupIdSuffix *string

	// The name of the primary cluster that accepts writes and will replicate updates
	// to the secondary cluster.
	//
	// This member is required.
	PrimaryReplicationGroupId *string

	// Provides details of the Global datastore
	GlobalReplicationGroupDescription *string

	noSmithyDocumentSerde
}

type CreateGlobalReplicationGroupOutput struct {

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

func (c *Client) addOperationCreateGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateGlobalReplicationGroup{}, middleware.After)
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
	if err = addOpCreateGlobalReplicationGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGlobalReplicationGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGlobalReplicationGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateGlobalReplicationGroup",
	}
}
