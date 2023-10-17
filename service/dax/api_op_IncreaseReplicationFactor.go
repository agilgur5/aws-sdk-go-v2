// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more nodes to a DAX cluster.
func (c *Client) IncreaseReplicationFactor(ctx context.Context, params *IncreaseReplicationFactorInput, optFns ...func(*Options)) (*IncreaseReplicationFactorOutput, error) {
	if params == nil {
		params = &IncreaseReplicationFactorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IncreaseReplicationFactor", params, optFns, c.addOperationIncreaseReplicationFactorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IncreaseReplicationFactorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IncreaseReplicationFactorInput struct {

	// The name of the DAX cluster that will receive additional nodes.
	//
	// This member is required.
	ClusterName *string

	// The new number of nodes for the DAX cluster.
	//
	// This member is required.
	NewReplicationFactor int32

	// The Availability Zones (AZs) in which the cluster nodes will be created. All
	// nodes belonging to the cluster are placed in these Availability Zones. Use this
	// parameter if you want to distribute the nodes across multiple AZs.
	AvailabilityZones []string

	noSmithyDocumentSerde
}

type IncreaseReplicationFactorOutput struct {

	// A description of the DAX cluster. with its new replication factor.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationIncreaseReplicationFactorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpIncreaseReplicationFactor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpIncreaseReplicationFactor{}, middleware.After)
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
	if err = addOpIncreaseReplicationFactorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIncreaseReplicationFactor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIncreaseReplicationFactor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dax",
		OperationName: "IncreaseReplicationFactor",
	}
}
