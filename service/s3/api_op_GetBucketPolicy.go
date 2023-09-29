// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the policy of a specified bucket. If you are using an identity other
// than the root user of the Amazon Web Services account that owns the bucket, the
// calling identity must have the GetBucketPolicy permissions on the specified
// bucket and belong to the bucket owner's account in order to use this operation.
// If you don't have GetBucketPolicy permissions, Amazon S3 returns a 403 Access
// Denied error. If you have the correct permissions, but you're not using an
// identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error. To ensure that bucket owners don't inadvertently lock
// themselves out of their own buckets, the root principal in a bucket owner's
// Amazon Web Services account can perform the GetBucketPolicy , PutBucketPolicy ,
// and DeleteBucketPolicy API actions, even if their bucket policy explicitly
// denies the root principal's access. Bucket owner root principals can only be
// blocked from performing these API actions by VPC endpoint policies and Amazon
// Web Services Organizations policies. To use this API operation against an access
// point, provide the alias of the access point in place of the bucket name. To use
// this API operation against an Object Lambda access point, provide the alias of
// the Object Lambda access point in place of the bucket name. If the Object Lambda
// access point alias in a request is not valid, the error code
// InvalidAccessPointAliasError is returned. For more information about
// InvalidAccessPointAliasError , see List of Error Codes (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList)
// . For more information about bucket policies, see Using Bucket Policies and
// User Policies (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html)
// . The following action is related to GetBucketPolicy :
//   - GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
func (c *Client) GetBucketPolicy(ctx context.Context, params *GetBucketPolicyInput, optFns ...func(*Options)) (*GetBucketPolicyOutput, error) {
	if params == nil {
		params = &GetBucketPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketPolicy", params, optFns, c.addOperationGetBucketPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketPolicyInput struct {

	// The bucket name for which to get the bucket policy. To use this API operation
	// against an access point, provide the alias of the access point in place of the
	// bucket name. To use this API operation against an Object Lambda access point,
	// provide the alias of the Object Lambda access point in place of the bucket name.
	// If the Object Lambda access point alias in a request is not valid, the error
	// code InvalidAccessPointAliasError is returned. For more information about
	// InvalidAccessPointAliasError , see List of Error Codes (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList)
	// .
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (*GetBucketPolicyInput) operationName() string {
	return "GetBucketPolicy"
}

func (in *GetBucketPolicyInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket

}

type GetBucketPolicyOutput struct {

	// The bucket policy as a JSON document.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketPolicy{}, middleware.After)
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetBucketPolicyUpdateEndpoint(stack, options); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func (v *GetBucketPolicyInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opGetBucketPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketPolicy",
	}
}

// getGetBucketPolicyBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getGetBucketPolicyBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketPolicyInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetBucketPolicyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetBucketPolicyBucketMember,
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
