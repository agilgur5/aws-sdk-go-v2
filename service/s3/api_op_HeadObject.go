// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// The HEAD action retrieves metadata from an object without returning the object
// itself. This action is useful if you're only interested in an object's metadata.
// To use HEAD , you must have READ access to the object. A HEAD request has the
// same options as a GET action on an object. The response is identical to the GET
// response except that there is no response body. Because of this, if the HEAD
// request generates an error, it returns a generic 400 Bad Request , 403 Forbidden
// or 404 Not Found code. It is not possible to retrieve the exact exception
// beyond these error codes. If you encrypt an object by using server-side
// encryption with customer-provided encryption keys (SSE-C) when you store the
// object in Amazon S3, then when you retrieve the metadata from the object, you
// must use the following headers:
//   - x-amz-server-side-encryption-customer-algorithm
//   - x-amz-server-side-encryption-customer-key
//   - x-amz-server-side-encryption-customer-key-MD5
//
// For more information about SSE-C, see Server-Side Encryption (Using
// Customer-Provided Encryption Keys) (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html)
// .
//   - Encryption request headers, like x-amz-server-side-encryption , should not
//     be sent for GET requests if your object uses server-side encryption with Key
//     Management Service (KMS) keys (SSE-KMS), dual-layer server-side encryption with
//     Amazon Web Services KMS keys (DSSE-KMS), or server-side encryption with Amazon
//     S3 managed encryption keys (SSE-S3). If your object does use these types of
//     keys, you’ll get an HTTP 400 Bad Request error.
//   - The last modified property in this case is the creation date of the object.
//
// Request headers are limited to 8 KB in size. For more information, see Common
// Request Headers (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonRequestHeaders.html)
// . Consider the following when using request headers:
//   - Consideration 1 – If both of the If-Match and If-Unmodified-Since headers
//     are present in the request as follows:
//   - If-Match condition evaluates to true , and;
//   - If-Unmodified-Since condition evaluates to false ; Then Amazon S3 returns
//     200 OK and the data requested.
//   - Consideration 2 – If both of the If-None-Match and If-Modified-Since headers
//     are present in the request as follows:
//   - If-None-Match condition evaluates to false , and;
//   - If-Modified-Since condition evaluates to true ; Then Amazon S3 returns the
//     304 Not Modified response code.
//
// For more information about conditional requests, see RFC 7232 (https://tools.ietf.org/html/rfc7232)
// . Permissions You need the relevant read object (or version) permission for this
// operation. For more information, see Actions, resources, and condition keys for
// Amazon S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/list_amazons3.html) .
// If the object you request doesn't exist, the error that Amazon S3 returns
// depends on whether you also have the s3:ListBucket permission.
//   - If you have the s3:ListBucket permission on the bucket, Amazon S3 returns an
//     HTTP status code 404 error.
//   - If you don’t have the s3:ListBucket permission, Amazon S3 returns an HTTP
//     status code 403 error.
//
// The following actions are related to HeadObject :
//   - GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
//   - GetObjectAttributes (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html)
func (c *Client) HeadObject(ctx context.Context, params *HeadObjectInput, optFns ...func(*Options)) (*HeadObjectOutput, error) {
	if params == nil {
		params = &HeadObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HeadObject", params, optFns, c.addOperationHeadObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HeadObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HeadObjectInput struct {

	// The name of the bucket containing the object. When using this action with an
	// access point, you must direct requests to the access point hostname. The access
	// point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When you use this action with Amazon S3 on
	// Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on
	// Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts? (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// The object key.
	//
	// This member is required.
	Key *string

	// To retrieve the checksum, this parameter must be enabled. In addition, if you
	// enable ChecksumMode and the object is encrypted with Amazon Web Services Key
	// Management Service (Amazon Web Services KMS), you must have permission to use
	// the kms:Decrypt action for the request to succeed.
	ChecksumMode types.ChecksumMode

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	// Return the object only if its entity tag (ETag) is the same as the one
	// specified; otherwise, return a 412 (precondition failed) error.
	IfMatch *string

	// Return the object only if it has been modified since the specified time;
	// otherwise, return a 304 (not modified) error.
	IfModifiedSince *time.Time

	// Return the object only if its entity tag (ETag) is different from the one
	// specified; otherwise, return a 304 (not modified) error.
	IfNoneMatch *string

	// Return the object only if it has not been modified since the specified time;
	// otherwise, return a 412 (precondition failed) error.
	IfUnmodifiedSince *time.Time

	// Part number of the object being read. This is a positive integer between 1 and
	// 10,000. Effectively performs a 'ranged' HEAD request for the part specified.
	// Useful querying about the size of the part and the number of parts in this
	// object.
	PartNumber int32

	// HeadObject returns only the metadata for an object. If the Range is
	// satisfiable, only the ContentLength is affected in the response. If the Range
	// is not satisfiable, S3 returns a 416 - Requested Range Not Satisfiable error.
	Range *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from Requester Pays buckets, see Downloading Objects
	// in Requester Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 User Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// VersionId used to reference a specific version of the object.
	VersionId *string

	noSmithyDocumentSerde
}

type HeadObjectOutput struct {

	// Indicates that a range of bytes was specified.
	AcceptRanges *string

	// The archive state of the head object.
	ArchiveStatus types.ArchiveStatus

	// Indicates whether the object uses an S3 Bucket Key for server-side encryption
	// with Key Management Service (KMS) keys (SSE-KMS).
	BucketKeyEnabled bool

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// The base64-encoded, 32-bit CRC32 checksum of the object. This will only be
	// present if it was uploaded with the object. With multipart uploads, this may not
	// be a checksum value of the object. For more information about how checksums are
	// calculated with multipart uploads, see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html#large-object-checksums)
	// in the Amazon S3 User Guide.
	ChecksumCRC32 *string

	// The base64-encoded, 32-bit CRC32C checksum of the object. This will only be
	// present if it was uploaded with the object. With multipart uploads, this may not
	// be a checksum value of the object. For more information about how checksums are
	// calculated with multipart uploads, see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html#large-object-checksums)
	// in the Amazon S3 User Guide.
	ChecksumCRC32C *string

	// The base64-encoded, 160-bit SHA-1 digest of the object. This will only be
	// present if it was uploaded with the object. With multipart uploads, this may not
	// be a checksum value of the object. For more information about how checksums are
	// calculated with multipart uploads, see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html#large-object-checksums)
	// in the Amazon S3 User Guide.
	ChecksumSHA1 *string

	// The base64-encoded, 256-bit SHA-256 digest of the object. This will only be
	// present if it was uploaded with the object. With multipart uploads, this may not
	// be a checksum value of the object. For more information about how checksums are
	// calculated with multipart uploads, see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html#large-object-checksums)
	// in the Amazon S3 User Guide.
	ChecksumSHA256 *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// Size of the body in bytes.
	ContentLength int64

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker bool

	// An entity tag (ETag) is an opaque identifier assigned by a web server to a
	// specific version of a resource found at a URL.
	ETag *string

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key-value pairs
	// providing object expiration information. The value of the rule-id is
	// URL-encoded.
	Expiration *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// Creation date of the object.
	LastModified *time.Time

	// A map of metadata to store with the object in S3.
	//
	// Map keys will be normalized to lower-case.
	Metadata map[string]string

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP, you
	// can create metadata whose values are not legal HTTP headers.
	MissingMeta int32

	// Specifies whether a legal hold is in effect for this object. This header is
	// only returned if the requester has the s3:GetObjectLegalHold permission. This
	// header is not returned if the specified version of this object has never had a
	// legal hold applied. For more information about S3 Object Lock, see Object Lock (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html)
	// .
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// The Object Lock mode, if any, that's in effect for this object. This header is
	// only returned if the requester has the s3:GetObjectRetention permission. For
	// more information about S3 Object Lock, see Object Lock (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html)
	// .
	ObjectLockMode types.ObjectLockMode

	// The date and time when the Object Lock retention period expires. This header is
	// only returned if the requester has the s3:GetObjectRetention permission.
	ObjectLockRetainUntilDate *time.Time

	// The count of parts this object has. This value is only returned if you specify
	// partNumber in your request and the object was uploaded as a multipart upload.
	PartsCount int32

	// Amazon S3 can return this header if your request involves a bucket that is
	// either a source or a destination in a replication rule. In replication, you have
	// a source bucket on which you configure replication and destination bucket or
	// buckets where Amazon S3 stores object replicas. When you request an object (
	// GetObject ) or object metadata ( HeadObject ) from these buckets, Amazon S3 will
	// return the x-amz-replication-status header in the response as follows:
	//   - If requesting an object from the source bucket, Amazon S3 will return the
	//   x-amz-replication-status header if the object in your request is eligible for
	//   replication. For example, suppose that in your replication configuration, you
	//   specify object prefix TaxDocs requesting Amazon S3 to replicate objects with
	//   key prefix TaxDocs . Any objects you upload with this key name prefix, for
	//   example TaxDocs/document1.pdf , are eligible for replication. For any object
	//   request with this key name prefix, Amazon S3 will return the
	//   x-amz-replication-status header with value PENDING, COMPLETED or FAILED
	//   indicating object replication status.
	//   - If requesting an object from a destination bucket, Amazon S3 will return
	//   the x-amz-replication-status header with value REPLICA if the object in your
	//   request is a replica that Amazon S3 created and there is no replica modification
	//   replication in progress.
	//   - When replicating objects to multiple destination buckets, the
	//   x-amz-replication-status header acts differently. The header of the source
	//   object will only return a value of COMPLETED when replication is successful to
	//   all destinations. The header will remain at value PENDING until replication has
	//   completed for all destinations. If one or more destinations fails replication
	//   the header will return FAILED.
	// For more information, see Replication (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html)
	// .
	ReplicationStatus types.ReplicationStatus

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If the object is an archived object (an object whose storage class is GLACIER),
	// the response includes this header if either the archive restoration is in
	// progress (see RestoreObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html)
	// or an archive copy is already restored. If an archive copy is already restored,
	// the header value indicates when Amazon S3 is scheduled to delete the object
	// copy. For example: x-amz-restore: ongoing-request="false", expiry-date="Fri, 21
	// Dec 2012 00:00:00 GMT" If the object restoration is in progress, the header
	// returns the value ongoing-request="true" . For more information about archiving
	// objects, see Transitioning Objects: General Considerations (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-transition-general-considerations)
	// .
	Restore *string

	// If server-side encryption with a customer-provided encryption key was
	// requested, the response will include this header confirming the encryption
	// algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was
	// requested, the response will include this header to provide round-trip message
	// integrity verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the ID of the Key Management Service (KMS) symmetric
	// encryption customer managed key that was used for the object.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256 , aws:kms , aws:kms:dsse ).
	ServerSideEncryption types.ServerSideEncryption

	// Provides storage class information of the object. Amazon S3 returns this header
	// for all objects except for S3 Standard storage class objects. For more
	// information, see Storage Classes (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// .
	StorageClass types.StorageClass

	// Version of the object.
	VersionId *string

	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata.
	WebsiteRedirectLocation *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationHeadObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpHeadObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpHeadObject{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addLegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addHeadObjectResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpHeadObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHeadObject(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addHeadObjectUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func (v *HeadObjectInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

// HeadObjectAPIClient is a client that implements the HeadObject operation.
type HeadObjectAPIClient interface {
	HeadObject(context.Context, *HeadObjectInput, ...func(*Options)) (*HeadObjectOutput, error)
}

var _ HeadObjectAPIClient = (*Client)(nil)

// ObjectExistsWaiterOptions are waiter options for ObjectExistsWaiter
type ObjectExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ObjectExistsWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ObjectExistsWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *HeadObjectInput, *HeadObjectOutput, error) (bool, error)
}

// ObjectExistsWaiter defines the waiters for ObjectExists
type ObjectExistsWaiter struct {
	client HeadObjectAPIClient

	options ObjectExistsWaiterOptions
}

// NewObjectExistsWaiter constructs a ObjectExistsWaiter.
func NewObjectExistsWaiter(client HeadObjectAPIClient, optFns ...func(*ObjectExistsWaiterOptions)) *ObjectExistsWaiter {
	options := ObjectExistsWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = objectExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ObjectExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ObjectExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ObjectExistsWaiter) Wait(ctx context.Context, params *HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*ObjectExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ObjectExists waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *ObjectExistsWaiter) WaitForOutput(ctx context.Context, params *HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*ObjectExistsWaiterOptions)) (*HeadObjectOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.HeadObject(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ObjectExists waiter")
}

func objectExistsStateRetryable(ctx context.Context, input *HeadObjectInput, output *HeadObjectOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var errorType *types.NotFound
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// ObjectNotExistsWaiterOptions are waiter options for ObjectNotExistsWaiter
type ObjectNotExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ObjectNotExistsWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ObjectNotExistsWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *HeadObjectInput, *HeadObjectOutput, error) (bool, error)
}

// ObjectNotExistsWaiter defines the waiters for ObjectNotExists
type ObjectNotExistsWaiter struct {
	client HeadObjectAPIClient

	options ObjectNotExistsWaiterOptions
}

// NewObjectNotExistsWaiter constructs a ObjectNotExistsWaiter.
func NewObjectNotExistsWaiter(client HeadObjectAPIClient, optFns ...func(*ObjectNotExistsWaiterOptions)) *ObjectNotExistsWaiter {
	options := ObjectNotExistsWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = objectNotExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ObjectNotExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ObjectNotExists waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *ObjectNotExistsWaiter) Wait(ctx context.Context, params *HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*ObjectNotExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ObjectNotExists waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ObjectNotExistsWaiter) WaitForOutput(ctx context.Context, params *HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*ObjectNotExistsWaiterOptions)) (*HeadObjectOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.HeadObject(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ObjectNotExists waiter")
}

func objectNotExistsStateRetryable(ctx context.Context, input *HeadObjectInput, output *HeadObjectOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.NotFound
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opHeadObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "HeadObject",
	}
}

// getHeadObjectBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getHeadObjectBucketMember(input interface{}) (*string, bool) {
	in := input.(*HeadObjectInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addHeadObjectUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getHeadObjectBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}

// PresignHeadObject is used to generate a presigned HTTP Request which contains
// presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignHeadObject(ctx context.Context, params *HeadObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &HeadObjectInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "HeadObject", params, clientOptFns,
		c.client.addOperationHeadObjectMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
		addHeadObjectPayloadAsUnsigned,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}

func addHeadObjectPayloadAsUnsigned(stack *middleware.Stack, options Options) error {
	v4.RemoveContentSHA256HeaderMiddleware(stack)
	v4.RemoveComputePayloadSHA256Middleware(stack)
	return v4.AddUnsignedPayloadMiddleware(stack)
}

type opHeadObjectResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opHeadObjectResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opHeadObjectResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*HeadObjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.Bucket = input.Bucket

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	ctx = smithyhttp.DisableEndpointHostPrefix(ctx, true)

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "s3"
			signingRegion := m.BuiltInResolver.(*BuiltInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			ctx = s3cust.SetSignerVersion(ctx, internalauth.SigV4)
		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "s3"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*BuiltInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			ctx = s3cust.SetSignerVersion(ctx, v4Scheme.Name)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("s3")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			ctx = s3cust.SetSignerVersion(ctx, v4aScheme.Name)
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addHeadObjectResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opHeadObjectResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:                         options.Region,
			UseFIPS:                        options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack:                   options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:                       options.BaseEndpoint,
			ForcePathStyle:                 options.UsePathStyle,
			Accelerate:                     options.UseAccelerate,
			DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
			UseArnRegion:                   options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
