// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves (queries) aggregated statistical data about all the S3 buckets that
// Amazon Macie monitors and analyzes for an account.
func (c *Client) GetBucketStatistics(ctx context.Context, params *GetBucketStatisticsInput, optFns ...func(*Options)) (*GetBucketStatisticsOutput, error) {
	if params == nil {
		params = &GetBucketStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketStatistics", params, optFns, c.addOperationGetBucketStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketStatisticsInput struct {

	// The unique identifier for the Amazon Web Services account.
	AccountId *string

	noSmithyDocumentSerde
}

type GetBucketStatisticsOutput struct {

	// The total number of buckets.
	BucketCount *int64

	// The total number of buckets that are publicly accessible due to a combination
	// of permissions settings for each bucket.
	BucketCountByEffectivePermission *types.BucketCountByEffectivePermission

	// The total number of buckets whose settings do or don't specify default
	// server-side encryption behavior for objects that are added to the buckets.
	BucketCountByEncryptionType *types.BucketCountByEncryptionType

	// The total number of buckets whose bucket policies do or don't require
	// server-side encryption of objects when objects are added to the buckets.
	BucketCountByObjectEncryptionRequirement *types.BucketCountPolicyAllowsUnencryptedObjectUploads

	// The total number of buckets that are or aren't shared with other Amazon Web
	// Services accounts, Amazon CloudFront origin access identities (OAIs), or
	// CloudFront origin access controls (OACs).
	BucketCountBySharedAccessType *types.BucketCountBySharedAccessType

	// The aggregated sensitive data discovery statistics for the buckets. If
	// automated sensitive data discovery is currently disabled for your account, the
	// value for each statistic is 0.
	BucketStatisticsBySensitivity *types.BucketStatisticsBySensitivity

	// The total number of objects that Amazon Macie can analyze in the buckets. These
	// objects use a supported storage class and have a file name extension for a
	// supported file or storage format.
	ClassifiableObjectCount *int64

	// The total storage size, in bytes, of all the objects that Amazon Macie can
	// analyze in the buckets. These objects use a supported storage class and have a
	// file name extension for a supported file or storage format. If versioning is
	// enabled for any of the buckets, this value is based on the size of the latest
	// version of each applicable object in the buckets. This value doesn't reflect the
	// storage size of all versions of all applicable objects in the buckets.
	ClassifiableSizeInBytes *int64

	// The date and time, in UTC and extended ISO 8601 format, when Amazon Macie most
	// recently retrieved bucket or object metadata from Amazon S3 for the buckets.
	LastUpdated *time.Time

	// The total number of objects in the buckets.
	ObjectCount *int64

	// The total storage size, in bytes, of the buckets. If versioning is enabled for
	// any of the buckets, this value is based on the size of the latest version of
	// each object in the buckets. This value doesn't reflect the storage size of all
	// versions of the objects in the buckets.
	SizeInBytes *int64

	// The total storage size, in bytes, of the objects that are compressed (.gz,
	// .gzip, .zip) files in the buckets. If versioning is enabled for any of the
	// buckets, this value is based on the size of the latest version of each
	// applicable object in the buckets. This value doesn't reflect the storage size of
	// all versions of the applicable objects in the buckets.
	SizeInBytesCompressed *int64

	// The total number of objects that Amazon Macie can't analyze in the buckets.
	// These objects don't use a supported storage class or don't have a file name
	// extension for a supported file or storage format.
	UnclassifiableObjectCount *types.ObjectLevelStatistics

	// The total storage size, in bytes, of the objects that Amazon Macie can't
	// analyze in the buckets. These objects don't use a supported storage class or
	// don't have a file name extension for a supported file or storage format.
	UnclassifiableObjectSizeInBytes *types.ObjectLevelStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBucketStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBucketStatistics{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBucketStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetBucketStatistics",
	}
}
