// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a block storage disk from a manual or automatic snapshot of a disk. The
// resulting disk can be attached to an Amazon Lightsail instance in the same
// Availability Zone (e.g., us-east-2a ). The create disk from snapshot operation
// supports tag-based access control via request tags and resource tags applied to
// the resource identified by disk snapshot name . For more information, see the
// Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags)
// .
func (c *Client) CreateDiskFromSnapshot(ctx context.Context, params *CreateDiskFromSnapshotInput, optFns ...func(*Options)) (*CreateDiskFromSnapshotOutput, error) {
	if params == nil {
		params = &CreateDiskFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDiskFromSnapshot", params, optFns, c.addOperationCreateDiskFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDiskFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDiskFromSnapshotInput struct {

	// The Availability Zone where you want to create the disk (e.g., us-east-2a ).
	// Choose the same Availability Zone as the Lightsail instance where you want to
	// create the disk. Use the GetRegions operation to list the Availability Zones
	// where Lightsail is currently available.
	//
	// This member is required.
	AvailabilityZone *string

	// The unique Lightsail disk name (e.g., my-disk ).
	//
	// This member is required.
	DiskName *string

	// The size of the disk in GB (e.g., 32 ).
	//
	// This member is required.
	SizeInGb *int32

	// An array of objects that represent the add-ons to enable for the new disk.
	AddOns []types.AddOnRequest

	// The name of the disk snapshot (e.g., my-snapshot ) from which to create the new
	// storage disk. Constraint:
	//   - This parameter cannot be defined together with the source disk name
	//   parameter. The disk snapshot name and source disk name parameters are mutually
	//   exclusive.
	DiskSnapshotName *string

	// The date of the automatic snapshot to use for the new disk. Use the get auto
	// snapshots operation to identify the dates of the available automatic snapshots.
	// Constraints:
	//   - Must be specified in YYYY-MM-DD format.
	//   - This parameter cannot be defined together with the use latest restorable
	//   auto snapshot parameter. The restore date and use latest restorable auto
	//   snapshot parameters are mutually exclusive.
	//   - Define this parameter only when creating a new disk from an automatic
	//   snapshot. For more information, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots)
	//   .
	RestoreDate *string

	// The name of the source disk from which the source automatic snapshot was
	// created. Constraints:
	//   - This parameter cannot be defined together with the disk snapshot name
	//   parameter. The source disk name and disk snapshot name parameters are mutually
	//   exclusive.
	//   - Define this parameter only when creating a new disk from an automatic
	//   snapshot. For more information, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots)
	//   .
	SourceDiskName *string

	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []types.Tag

	// A Boolean value to indicate whether to use the latest available automatic
	// snapshot. Constraints:
	//   - This parameter cannot be defined together with the restore date parameter.
	//   The use latest restorable auto snapshot and restore date parameters are
	//   mutually exclusive.
	//   - Define this parameter only when creating a new disk from an automatic
	//   snapshot. For more information, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots)
	//   .
	UseLatestRestorableAutoSnapshot *bool

	noSmithyDocumentSerde
}

type CreateDiskFromSnapshotOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDiskFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDiskFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDiskFromSnapshot{}, middleware.After)
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
	if err = addOpCreateDiskFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDiskFromSnapshot(options.Region), middleware.Before); err != nil {
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
		operation: "CreateDiskFromSnapshot",
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

func newServiceMetadataMiddleware_opCreateDiskFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateDiskFromSnapshot",
	}
}
