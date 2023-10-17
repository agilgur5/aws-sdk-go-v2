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

// Deletes the specified EC2 Fleets. After you delete an EC2 Fleet, it launches no
// new instances. You must also specify whether a deleted EC2 Fleet should
// terminate its instances. If you choose to terminate the instances, the EC2 Fleet
// enters the deleted_terminating state. Otherwise, the EC2 Fleet enters the
// deleted_running state, and the instances continue to run until they are
// interrupted or you terminate them manually. For instant fleets, EC2 Fleet must
// terminate the instances when the fleet is deleted. A deleted instant fleet with
// running instances is not supported. Restrictions
//   - You can delete up to 25 instant fleets in a single request. If you exceed
//     this number, no instant fleets are deleted and an error is returned. There is
//     no restriction on the number of fleets of type maintain or request that can be
//     deleted in a single request.
//   - Up to 1000 instances can be terminated in a single request to delete instant
//     fleets.
//
// For more information, see Delete an EC2 Fleet (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#delete-fleet)
// in the Amazon EC2 User Guide.
func (c *Client) DeleteFleets(ctx context.Context, params *DeleteFleetsInput, optFns ...func(*Options)) (*DeleteFleetsOutput, error) {
	if params == nil {
		params = &DeleteFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFleets", params, optFns, c.addOperationDeleteFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFleetsInput struct {

	// The IDs of the EC2 Fleets.
	//
	// This member is required.
	FleetIds []string

	// Indicates whether to terminate the associated instances when the EC2 Fleet is
	// deleted. The default is to terminate the instances. To let the instances
	// continue to run after the EC2 Fleet is deleted, specify no-terminate-instances .
	// Supported only for fleets of type maintain and request . For instant fleets,
	// you cannot specify NoTerminateInstances . A deleted instant fleet with running
	// instances is not supported.
	//
	// This member is required.
	TerminateInstances *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DeleteFleetsOutput struct {

	// Information about the EC2 Fleets that are successfully deleted.
	SuccessfulFleetDeletions []types.DeleteFleetSuccessItem

	// Information about the EC2 Fleets that are not successfully deleted.
	UnsuccessfulFleetDeletions []types.DeleteFleetErrorItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteFleets{}, middleware.After)
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
	if err = addOpDeleteFleetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFleets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteFleets",
	}
}
