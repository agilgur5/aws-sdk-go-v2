// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a Network File System (NFS) file share. This operation is only
// supported in S3 File Gateways. To leave a file share field unchanged, set the
// corresponding input field to null. Updates the following file share settings:
//   - Default storage class for your S3 bucket
//   - Metadata defaults for your S3 bucket
//   - Allowed NFS clients for your file share
//   - Squash settings
//   - Write status of your file share
func (c *Client) UpdateNFSFileShare(ctx context.Context, params *UpdateNFSFileShareInput, optFns ...func(*Options)) (*UpdateNFSFileShareOutput, error) {
	if params == nil {
		params = &UpdateNFSFileShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNFSFileShare", params, optFns, c.addOperationUpdateNFSFileShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNFSFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// UpdateNFSFileShareInput
type UpdateNFSFileShareInput struct {

	// The Amazon Resource Name (ARN) of the file share to be updated.
	//
	// This member is required.
	FileShareARN *string

	// The Amazon Resource Name (ARN) of the storage used for audit logs.
	AuditDestinationARN *string

	// Specifies refresh cache information for the file share.
	CacheAttributes *types.CacheAttributes

	// The list of clients that are allowed to access the S3 File Gateway. The list
	// must contain either valid IP addresses or valid CIDR blocks.
	ClientList []string

	// The default storage class for objects put into an Amazon S3 bucket by the S3
	// File Gateway. The default value is S3_STANDARD . Optional. Valid Values:
	// S3_STANDARD | S3_INTELLIGENT_TIERING | S3_STANDARD_IA | S3_ONEZONE_IA
	DefaultStorageClass *string

	// The name of the file share. Optional. FileShareName must be set if an S3 prefix
	// name is set in LocationARN , or if an access point or access point alias is used.
	FileShareName *string

	// A value that enables guessing of the MIME type for uploaded objects based on
	// file extensions. Set this value to true to enable MIME type guessing, otherwise
	// set to false . The default value is true . Valid Values: true | false
	GuessMIMETypeEnabled *bool

	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or
	// false to use a key managed by Amazon S3. Optional. Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used
	// for Amazon S3 server-side encryption. Storage Gateway does not support
	// asymmetric CMKs. This value can only be set when KMSEncrypted is true . Optional.
	KMSKey *string

	// The default values for the file share. Optional.
	NFSFileShareDefaults *types.NFSFileShareDefaults

	// The notification policy of the file share. SettlingTimeInSeconds controls the
	// number of seconds to wait after the last point in time a client wrote to a file
	// before generating an ObjectUploaded notification. Because clients can make many
	// small writes to files, it's best to set this parameter for as long as possible
	// to avoid generating multiple notifications for the same file in a small time
	// period. SettlingTimeInSeconds has no effect on the timing of the object
	// uploading to Amazon S3, only the timing of the notification. The following
	// example sets NotificationPolicy on with SettlingTimeInSeconds set to 60.
	// {\"Upload\": {\"SettlingTimeInSeconds\": 60}} The following example sets
	// NotificationPolicy off. {}
	NotificationPolicy *string

	// A value that sets the access control list (ACL) permission for objects in the
	// S3 bucket that a S3 File Gateway puts objects into. The default value is private
	// .
	ObjectACL types.ObjectACL

	// A value that sets the write status of a file share. Set this value to true to
	// set the write status to read-only, otherwise set to false . Valid Values: true
	// | false
	ReadOnly *bool

	// A value that sets who pays the cost of the request and the cost associated with
	// data download from the S3 bucket. If this value is set to true , the requester
	// pays the costs; otherwise, the S3 bucket owner pays. However, the S3 bucket
	// owner always pays the cost of storing data. RequesterPays is a configuration
	// for the S3 bucket that backs the file share, so make sure that the configuration
	// on the file share is the same as the S3 bucket configuration. Valid Values: true
	// | false
	RequesterPays *bool

	// The user mapped to anonymous user. Valid values are the following:
	//   - RootSquash : Only root is mapped to anonymous user.
	//   - NoSquash : No one is mapped to anonymous user.
	//   - AllSquash : Everyone is mapped to anonymous user.
	Squash *string

	noSmithyDocumentSerde
}

// UpdateNFSFileShareOutput
type UpdateNFSFileShareOutput struct {

	// The Amazon Resource Name (ARN) of the updated file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNFSFileShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNFSFileShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNFSFileShare{}, middleware.After)
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
	if err = addOpUpdateNFSFileShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNFSFileShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNFSFileShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateNFSFileShare",
	}
}
