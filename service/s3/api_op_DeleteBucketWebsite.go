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

// This action removes the website configuration for a bucket. Amazon S3 returns a
// 200 OK response upon successfully deleting a website configuration on the
// specified bucket. You will get a 200 OK response if the website configuration
// you are trying to delete does not exist on the bucket. Amazon S3 returns a 404
// response if the bucket specified in the request does not exist. This DELETE
// action requires the S3:DeleteBucketWebsite permission. By default, only the
// bucket owner can delete the website configuration attached to a bucket. However,
// bucket owners can grant other users permission to delete the website
// configuration by writing a bucket policy granting them the
// S3:DeleteBucketWebsite permission. For more information about hosting websites,
// see Hosting Websites on Amazon S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html)
// . The following operations are related to DeleteBucketWebsite :
//   - GetBucketWebsite (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketWebsite.html)
//   - PutBucketWebsite (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketWebsite.html)
func (c *Client) DeleteBucketWebsite(ctx context.Context, params *DeleteBucketWebsiteInput, optFns ...func(*Options)) (*DeleteBucketWebsiteOutput, error) {
	if params == nil {
		params = &DeleteBucketWebsiteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketWebsite", params, optFns, c.addOperationDeleteBucketWebsiteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketWebsiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketWebsiteInput struct {

	// The bucket name for which you want to remove the website configuration.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (*DeleteBucketWebsiteInput) operationName() string {
	return "DeleteBucketWebsite"
}

func (in *DeleteBucketWebsiteInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket

}

type DeleteBucketWebsiteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBucketWebsiteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketWebsite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketWebsite{}, middleware.After)
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
	if err = addOpDeleteBucketWebsiteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketWebsite(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addDeleteBucketWebsiteUpdateEndpoint(stack, options); err != nil {
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

func (v *DeleteBucketWebsiteInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opDeleteBucketWebsite(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketWebsite",
	}
}

// getDeleteBucketWebsiteBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getDeleteBucketWebsiteBucketMember(input interface{}) (*string, bool) {
	in := input.(*DeleteBucketWebsiteInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addDeleteBucketWebsiteUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getDeleteBucketWebsiteBucketMember,
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
