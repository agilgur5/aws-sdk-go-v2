// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the VPC peering connection options on one side of a VPC peering
// connection. If the peered VPCs are in the same Amazon Web Services account, you
// can enable DNS resolution for queries from the local VPC. This ensures that
// queries from the local VPC resolve to private IP addresses in the peer VPC. This
// option is not available if the peered VPCs are in different Amazon Web Services
// accounts or different Regions. For peered VPCs in different Amazon Web Services
// accounts, each Amazon Web Services account owner must initiate a separate
// request to modify the peering connection options. For inter-region peering
// connections, you must use the Region for the requester VPC to modify the
// requester VPC peering options and the Region for the accepter VPC to modify the
// accepter VPC peering options. To verify which VPCs are the accepter and the
// requester for a VPC peering connection, use the DescribeVpcPeeringConnections
// command.
func (c *Client) ModifyVpcPeeringConnectionOptions(ctx context.Context, params *ModifyVpcPeeringConnectionOptionsInput, optFns ...func(*Options)) (*ModifyVpcPeeringConnectionOptionsOutput, error) {
	if params == nil {
		params = &ModifyVpcPeeringConnectionOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcPeeringConnectionOptions", params, optFns, c.addOperationModifyVpcPeeringConnectionOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcPeeringConnectionOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcPeeringConnectionOptionsInput struct {

	// The ID of the VPC peering connection.
	//
	// This member is required.
	VpcPeeringConnectionId *string

	// The VPC peering connection options for the accepter VPC.
	AccepterPeeringConnectionOptions *types.PeeringConnectionOptionsRequest

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The VPC peering connection options for the requester VPC.
	RequesterPeeringConnectionOptions *types.PeeringConnectionOptionsRequest

	noSmithyDocumentSerde
}

type ModifyVpcPeeringConnectionOptionsOutput struct {

	// Information about the VPC peering connection options for the accepter VPC.
	AccepterPeeringConnectionOptions *types.PeeringConnectionOptions

	// Information about the VPC peering connection options for the requester VPC.
	RequesterPeeringConnectionOptions *types.PeeringConnectionOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVpcPeeringConnectionOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcPeeringConnectionOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcPeeringConnectionOptions{}, middleware.After)
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
	if err = addOpModifyVpcPeeringConnectionOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcPeeringConnectionOptions(options.Region), middleware.Before); err != nil {
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
		operation: "ModifyVpcPeeringConnectionOptions",
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

func newServiceMetadataMiddleware_opModifyVpcPeeringConnectionOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcPeeringConnectionOptions",
	}
}
