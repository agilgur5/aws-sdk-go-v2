// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation sets the versioning state for S3 on Outposts buckets only. To
// set the versioning state for an S3 bucket, see PutBucketVersioning (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketVersioning.html)
// in the Amazon S3 API Reference. Sets the versioning state for an S3 on Outposts
// bucket. With S3 Versioning, you can save multiple distinct copies of your
// objects and recover from unintended user actions and application failures. You
// can set the versioning state to one of the following:
//   - Enabled - Enables versioning for the objects in the bucket. All objects
//     added to the bucket receive a unique version ID.
//   - Suspended - Suspends versioning for the objects in the bucket. All objects
//     added to the bucket receive the version ID null .
//
// If you've never set versioning on your bucket, it has no versioning state. In
// that case, a GetBucketVersioning (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketVersioning.html)
// request does not return a versioning state value. When you enable S3 Versioning,
// for each object in your bucket, you have a current version and zero or more
// noncurrent versions. You can configure your bucket S3 Lifecycle rules to expire
// noncurrent versions after a specified time period. For more information, see
// Creating and managing a lifecycle configuration for your S3 on Outposts bucket (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsLifecycleManaging.html)
// in the Amazon S3 User Guide. If you have an object expiration lifecycle
// configuration in your non-versioned bucket and you want to maintain the same
// permanent delete behavior when you enable versioning, you must add a noncurrent
// expiration policy. The noncurrent expiration lifecycle configuration will manage
// the deletes of the noncurrent object versions in the version-enabled bucket. For
// more information, see Versioning (https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html)
// in the Amazon S3 User Guide. All Amazon S3 on Outposts REST API requests for
// this action require an additional parameter of x-amz-outpost-id to be passed
// with the request. In addition, you must use an S3 on Outposts endpoint hostname
// prefix instead of s3-control . For an example of the request syntax for Amazon
// S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// x-amz-outpost-id derived by using the access point ARN, see the Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketVersioning.html#API_control_PutBucketVersioning_Examples)
// section. The following operations are related to PutBucketVersioning for S3 on
// Outposts.
//   - GetBucketVersioning (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketVersioning.html)
//   - PutBucketLifecycleConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html)
//   - GetBucketLifecycleConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html)
func (c *Client) PutBucketVersioning(ctx context.Context, params *PutBucketVersioningInput, optFns ...func(*Options)) (*PutBucketVersioningOutput, error) {
	if params == nil {
		params = &PutBucketVersioningInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketVersioning", params, optFns, c.addOperationPutBucketVersioningMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketVersioningOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketVersioningInput struct {

	// The Amazon Web Services account ID of the S3 on Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The S3 on Outposts bucket to set the versioning state for.
	//
	// This member is required.
	Bucket *string

	// The root-level tag for the VersioningConfiguration parameters.
	//
	// This member is required.
	VersioningConfiguration *types.VersioningConfiguration

	// The concatenation of the authentication device's serial number, a space, and
	// the value that is displayed on your authentication device.
	MFA *string

	noSmithyDocumentSerde
}

func (in *PutBucketVersioningInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type PutBucketVersioningOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketVersioningMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketVersioning{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketVersioning{}, middleware.After)
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opPutBucketVersioningMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutBucketVersioningValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketVersioning(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketVersioningUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "PutBucketVersioning",
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

func (m *PutBucketVersioningInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *PutBucketVersioningInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opPutBucketVersioningMiddleware struct {
}

func (*endpointPrefix_opPutBucketVersioningMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutBucketVersioningMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutBucketVersioningInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opPutBucketVersioningMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPutBucketVersioningMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketVersioning(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketVersioning",
	}
}

func copyPutBucketVersioningInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*PutBucketVersioningInput)
	if !ok {
		return nil, fmt.Errorf("expect *PutBucketVersioningInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getPutBucketVersioningARNMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketVersioningInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setPutBucketVersioningARNMember(input interface{}, v string) error {
	in := input.(*PutBucketVersioningInput)
	in.Bucket = &v
	return nil
}
func backFillPutBucketVersioningAccountID(input interface{}, v string) error {
	in := input.(*PutBucketVersioningInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addPutBucketVersioningUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getPutBucketVersioningARNMember,
			BackfillAccountID: backFillPutBucketVersioningAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setPutBucketVersioningARNMember,
			CopyInput:         copyPutBucketVersioningInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
