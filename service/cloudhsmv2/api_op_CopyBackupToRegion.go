// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copy an AWS CloudHSM cluster backup to a different region.
func (c *Client) CopyBackupToRegion(ctx context.Context, params *CopyBackupToRegionInput, optFns ...func(*Options)) (*CopyBackupToRegionOutput, error) {
	if params == nil {
		params = &CopyBackupToRegionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyBackupToRegion", params, optFns, c.addOperationCopyBackupToRegionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyBackupToRegionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyBackupToRegionInput struct {

	// The ID of the backup that will be copied to the destination region.
	//
	// This member is required.
	BackupId *string

	// The AWS region that will contain your copied CloudHSM cluster backup.
	//
	// This member is required.
	DestinationRegion *string

	// Tags to apply to the destination backup during creation. If you specify tags,
	// only these tags will be applied to the destination backup. If you do not specify
	// tags, the service copies tags from the source backup to the destination backup.
	TagList []types.Tag

	noSmithyDocumentSerde
}

type CopyBackupToRegionOutput struct {

	// Information on the backup that will be copied to the destination region,
	// including CreateTimestamp, SourceBackup, SourceCluster, and Source Region.
	// CreateTimestamp of the destination backup will be the same as that of the source
	// backup. You will need to use the sourceBackupID returned in this operation to
	// use the DescribeBackups operation on the backup that will be copied to the
	// destination region.
	DestinationBackup *types.DestinationBackup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyBackupToRegionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCopyBackupToRegion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopyBackupToRegion{}, middleware.After)
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
	if err = addOpCopyBackupToRegionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyBackupToRegion(options.Region), middleware.Before); err != nil {
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
		operation: "CopyBackupToRegion",
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

func newServiceMetadataMiddleware_opCopyBackupToRegion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "CopyBackupToRegion",
	}
}
