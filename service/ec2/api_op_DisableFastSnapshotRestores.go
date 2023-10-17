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

// Disables fast snapshot restores for the specified snapshots in the specified
// Availability Zones.
func (c *Client) DisableFastSnapshotRestores(ctx context.Context, params *DisableFastSnapshotRestoresInput, optFns ...func(*Options)) (*DisableFastSnapshotRestoresOutput, error) {
	if params == nil {
		params = &DisableFastSnapshotRestoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableFastSnapshotRestores", params, optFns, c.addOperationDisableFastSnapshotRestoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableFastSnapshotRestoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableFastSnapshotRestoresInput struct {

	// One or more Availability Zones. For example, us-east-2a .
	//
	// This member is required.
	AvailabilityZones []string

	// The IDs of one or more snapshots. For example, snap-1234567890abcdef0 .
	//
	// This member is required.
	SourceSnapshotIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DisableFastSnapshotRestoresOutput struct {

	// Information about the snapshots for which fast snapshot restores were
	// successfully disabled.
	Successful []types.DisableFastSnapshotRestoreSuccessItem

	// Information about the snapshots for which fast snapshot restores could not be
	// disabled.
	Unsuccessful []types.DisableFastSnapshotRestoreErrorItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableFastSnapshotRestoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisableFastSnapshotRestores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisableFastSnapshotRestores{}, middleware.After)
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
	if err = addOpDisableFastSnapshotRestoresValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableFastSnapshotRestores(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableFastSnapshotRestores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableFastSnapshotRestores",
	}
}
