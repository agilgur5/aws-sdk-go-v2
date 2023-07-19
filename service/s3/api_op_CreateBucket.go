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
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new S3 bucket. To create a bucket, you must register with Amazon S3
// and have a valid Amazon Web Services Access Key ID to authenticate requests.
// Anonymous requests are never allowed to create buckets. By creating the bucket,
// you become the bucket owner. Not every string is an acceptable bucket name. For
// information about bucket naming restrictions, see Bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html)
// . If you want to create an Amazon S3 on Outposts bucket, see Create Bucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html)
// . By default, the bucket is created in the US East (N. Virginia) Region. You can
// optionally specify a Region in the request body. You might choose a Region to
// optimize latency, minimize costs, or address regulatory requirements. For
// example, if you reside in Europe, you will probably find it advantageous to
// create buckets in the Europe (Ireland) Region. For more information, see
// Accessing a bucket (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
// . If you send your create bucket request to the s3.amazonaws.com endpoint, the
// request goes to the us-east-1 Region. Accordingly, the signature calculations
// in Signature Version 4 must use us-east-1 as the Region, even if the location
// constraint in the request specifies another Region where the bucket is to be
// created. If you create a bucket in a Region other than US East (N. Virginia),
// your application must be able to handle 307 redirect. For more information, see
// Virtual hosting of buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html)
// . Permissions In addition to s3:CreateBucket , the following permissions are
// required when your CreateBucket request includes specific headers:
//   - Access control lists (ACLs) - If your CreateBucket request specifies access
//     control list (ACL) permissions and the ACL is public-read, public-read-write,
//     authenticated-read, or if you specify access permissions explicitly through any
//     other ACL, both s3:CreateBucket and s3:PutBucketAcl permissions are needed. If
//     the ACL for the CreateBucket request is private or if the request doesn't
//     specify any ACLs, only s3:CreateBucket permission is needed.
//   - Object Lock - If ObjectLockEnabledForBucket is set to true in your
//     CreateBucket request, s3:PutBucketObjectLockConfiguration and
//     s3:PutBucketVersioning permissions are required.
//   - S3 Object Ownership - If your CreateBucket request includes the
//     x-amz-object-ownership header, then the s3:PutBucketOwnershipControls
//     permission is required. By default, ObjectOwnership is set to
//     BucketOWnerEnforced and ACLs are disabled. We recommend keeping ACLs disabled,
//     except in uncommon use cases where you must control access for each object
//     individually. If you want to change the ObjectOwnership setting, you can use
//     the x-amz-object-ownership header in your CreateBucket request to set the
//     ObjectOwnership setting of your choice. For more information about S3 Object
//     Ownership, see Controlling object ownership  (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
//     in the Amazon S3 User Guide.
//   - S3 Block Public Access - If your specific use case requires granting public
//     access to your S3 resources, you can disable Block Public Access. You can create
//     a new bucket with Block Public Access enabled, then separately call the
//     DeletePublicAccessBlock (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html)
//     API. To use this operation, you must have the s3:PutBucketPublicAccessBlock
//     permission. By default, all Block Public Access settings are enabled for new
//     buckets. To avoid inadvertent exposure of your resources, we recommend keeping
//     the S3 Block Public Access settings enabled. For more information about S3 Block
//     Public Access, see Blocking public access to your Amazon S3 storage  (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
//     in the Amazon S3 User Guide.
//
// If your CreateBucket request sets BucketOwnerEnforced for Amazon S3 Object
// Ownership and specifies a bucket ACL that provides access to an external Amazon
// Web Services account, your request fails with a 400 error and returns the
// InvalidBucketAcLWithObjectOwnership error code. For more information, see
// Setting Object Ownership on an existing bucket  (https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-ownership-existing-bucket.html)
// in the Amazon S3 User Guide. The following operations are related to
// CreateBucket :
//   - PutObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//   - DeleteBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
func (c *Client) CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) {
	if params == nil {
		params = &CreateBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBucket", params, optFns, c.addOperationCreateBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBucketInput struct {

	// The name of the bucket to create.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket.
	ACL types.BucketCannedACL

	// The configuration information for the bucket.
	CreateBucketConfiguration *types.CreateBucketConfiguration

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket.
	GrantRead *string

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string

	// Allows grantee to create new objects in the bucket. For the bucket and object
	// owners of existing objects, also allows deletions and overwrites of those
	// objects.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	ObjectLockEnabledForBucket bool

	// The container element for object ownership for a bucket's ownership controls.
	// BucketOwnerPreferred - Objects uploaded to the bucket change ownership to the
	// bucket owner if the objects are uploaded with the bucket-owner-full-control
	// canned ACL. ObjectWriter - The uploading account will own the object if the
	// object is uploaded with the bucket-owner-full-control canned ACL.
	// BucketOwnerEnforced - Access control lists (ACLs) are disabled and no longer
	// affect permissions. The bucket owner automatically owns and has full control
	// over every object in the bucket. The bucket only accepts PUT requests that don't
	// specify an ACL or bucket owner full control ACLs, such as the
	// bucket-owner-full-control canned ACL or an equivalent form of this ACL expressed
	// in the XML format.
	ObjectOwnership types.ObjectOwnership

	noSmithyDocumentSerde
}

type CreateBucketOutput struct {

	// A forward slash followed by the name of the bucket.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateBucket{}, middleware.After)
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
	if err = addCreateBucketResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBucket(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addCreateBucketUpdateEndpoint(stack, options); err != nil {
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

func (v *CreateBucketInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opCreateBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CreateBucket",
	}
}

// getCreateBucketBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getCreateBucketBucketMember(input interface{}) (*string, bool) {
	in := input.(*CreateBucketInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addCreateBucketUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getCreateBucketBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             false,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}

type opCreateBucketResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opCreateBucketResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateBucketResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*CreateBucketInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.Bucket = input.Bucket

	params.DisableAccessPoints = ptr.Bool(true)

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

func addCreateBucketResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateBucketResolveEndpointMiddleware{
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
