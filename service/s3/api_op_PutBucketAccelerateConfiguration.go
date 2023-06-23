// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the accelerate configuration of an existing bucket. Amazon S3 Transfer
// Acceleration is a bucket-level feature that enables you to perform faster data
// transfers to Amazon S3. To use this operation, you must have permission to
// perform the s3:PutAccelerateConfiguration action. The bucket owner has this
// permission by default. The bucket owner can grant this permission to others. For
// more information about permissions, see Permissions Related to Bucket
// Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html)
// . The Transfer Acceleration state of a bucket can be set to one of the following
// two values:
//   - Enabled – Enables accelerated data transfers to the bucket.
//   - Suspended – Disables accelerated data transfers to the bucket.
//
// The GetBucketAccelerateConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAccelerateConfiguration.html)
// action returns the transfer acceleration state of a bucket. After setting the
// Transfer Acceleration state of a bucket to Enabled, it might take up to thirty
// minutes before the data transfer rates to the bucket increase. The name of the
// bucket used for Transfer Acceleration must be DNS-compliant and must not contain
// periods ("."). For more information about transfer acceleration, see Transfer
// Acceleration (https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html)
// . The following operations are related to PutBucketAccelerateConfiguration :
//   - GetBucketAccelerateConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAccelerateConfiguration.html)
//   - CreateBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
func (c *Client) PutBucketAccelerateConfiguration(ctx context.Context, params *PutBucketAccelerateConfigurationInput, optFns ...func(*Options)) (*PutBucketAccelerateConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketAccelerateConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketAccelerateConfiguration", params, optFns, c.addOperationPutBucketAccelerateConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketAccelerateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketAccelerateConfigurationInput struct {

	// Container for setting the transfer acceleration state.
	//
	// This member is required.
	AccelerateConfiguration *types.AccelerateConfiguration

	// The name of the bucket for which the accelerate configuration is set.
	//
	// This member is required.
	Bucket *string

	// Indicates the algorithm used to create the checksum for the object when using
	// the SDK. This header will not provide any additional functionality if not using
	// the SDK. When sending this header, there must be a corresponding x-amz-checksum
	// or x-amz-trailer header sent. Otherwise, Amazon S3 fails the request with the
	// HTTP status code 400 Bad Request . For more information, see Checking object
	// integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide. If you provide an individual checksum, Amazon S3
	// ignores any provided ChecksumAlgorithm parameter.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type PutBucketAccelerateConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketAccelerateConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketAccelerateConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketAccelerateConfiguration{}, middleware.After)
	if err != nil {
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
	if err = addPutBucketAccelerateConfigurationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutBucketAccelerateConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketAccelerateConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketAccelerateConfigurationInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addPutBucketAccelerateConfigurationUpdateEndpoint(stack, options); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opPutBucketAccelerateConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketAccelerateConfiguration",
	}
}

// getPutBucketAccelerateConfigurationRequestAlgorithmMember gets the request
// checksum algorithm value provided as input.
func getPutBucketAccelerateConfigurationRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutBucketAccelerateConfigurationInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutBucketAccelerateConfigurationInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutBucketAccelerateConfigurationRequestAlgorithmMember,
		RequireChecksum:                  false,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutBucketAccelerateConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getPutBucketAccelerateConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketAccelerateConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketAccelerateConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketAccelerateConfigurationBucketMember,
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

type opPutBucketAccelerateConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opPutBucketAccelerateConfigurationResolveEndpointMiddleware) ID() string {
	return "opPutBucketAccelerateConfigurationResolveEndpointMiddleware"
}

func (m *opPutBucketAccelerateConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutBucketAccelerateConfigurationInput)
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

	return next.HandleSerialize(ctx, in)
}

func addPutBucketAccelerateConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutBucketAccelerateConfigurationResolveEndpointMiddleware{
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
