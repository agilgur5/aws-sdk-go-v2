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

// Cancels the specified Spot Fleet requests. After you cancel a Spot Fleet
// request, the Spot Fleet launches no new instances. You must also specify whether
// a canceled Spot Fleet request should terminate its instances. If you choose to
// terminate the instances, the Spot Fleet request enters the cancelled_terminating
// state. Otherwise, the Spot Fleet request enters the cancelled_running state and
// the instances continue to run until they are interrupted or you terminate them
// manually.
func (c *Client) CancelSpotFleetRequests(ctx context.Context, params *CancelSpotFleetRequestsInput, optFns ...func(*Options)) (*CancelSpotFleetRequestsOutput, error) {
	if params == nil {
		params = &CancelSpotFleetRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelSpotFleetRequests", params, optFns, c.addOperationCancelSpotFleetRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelSpotFleetRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CancelSpotFleetRequests.
type CancelSpotFleetRequestsInput struct {

	// The IDs of the Spot Fleet requests.
	//
	// This member is required.
	SpotFleetRequestIds []string

	// Indicates whether to terminate the associated instances when the Spot Fleet
	// request is canceled. The default is to terminate the instances. To let the
	// instances continue to run after the Spot Fleet request is canceled, specify
	// no-terminate-instances .
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

// Contains the output of CancelSpotFleetRequests.
type CancelSpotFleetRequestsOutput struct {

	// Information about the Spot Fleet requests that are successfully canceled.
	SuccessfulFleetRequests []types.CancelSpotFleetRequestsSuccessItem

	// Information about the Spot Fleet requests that are not successfully canceled.
	UnsuccessfulFleetRequests []types.CancelSpotFleetRequestsErrorItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelSpotFleetRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCancelSpotFleetRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCancelSpotFleetRequests{}, middleware.After)
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
	if err = addOpCancelSpotFleetRequestsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelSpotFleetRequests(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelSpotFleetRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CancelSpotFleetRequests",
	}
}
