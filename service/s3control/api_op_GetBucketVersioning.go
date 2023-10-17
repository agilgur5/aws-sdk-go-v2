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

// This operation returns the versioning state for S3 on Outposts buckets only. To
// return the versioning state for an S3 bucket, see GetBucketVersioning (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html)
// in the Amazon S3 API Reference. Returns the versioning state for an S3 on
// Outposts bucket. With S3 Versioning, you can save multiple distinct copies of
// your objects and recover from unintended user actions and application failures.
// If you've never set versioning on your bucket, it has no versioning state. In
// that case, the GetBucketVersioning request does not return a versioning state
// value. For more information about versioning, see Versioning (https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html)
// in the Amazon S3 User Guide. All Amazon S3 on Outposts REST API requests for
// this action require an additional parameter of x-amz-outpost-id to be passed
// with the request. In addition, you must use an S3 on Outposts endpoint hostname
// prefix instead of s3-control . For an example of the request syntax for Amazon
// S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// x-amz-outpost-id derived by using the access point ARN, see the Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketVersioning.html#API_control_GetBucketVersioning_Examples)
// section. The following operations are related to GetBucketVersioning for S3 on
// Outposts.
//   - PutBucketVersioning (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketVersioning.html)
//   - PutBucketLifecycleConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html)
//   - GetBucketLifecycleConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html)
func (c *Client) GetBucketVersioning(ctx context.Context, params *GetBucketVersioningInput, optFns ...func(*Options)) (*GetBucketVersioningOutput, error) {
	if params == nil {
		params = &GetBucketVersioningInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketVersioning", params, optFns, c.addOperationGetBucketVersioningMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketVersioningOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketVersioningInput struct {

	// The Amazon Web Services account ID of the S3 on Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The S3 on Outposts bucket to return the versioning state for.
	//
	// This member is required.
	Bucket *string

	noSmithyDocumentSerde
}

func (in *GetBucketVersioningInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type GetBucketVersioningOutput struct {

	// Specifies whether MFA delete is enabled in the bucket versioning configuration.
	// This element is returned only if the bucket has been configured with MFA delete.
	// If MFA delete has never been configured for the bucket, this element is not
	// returned.
	MFADelete types.MFADeleteStatus

	// The versioning state of the S3 on Outposts bucket.
	Status types.BucketVersioningStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketVersioningMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketVersioning{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketVersioning{}, middleware.After)
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetBucketVersioningMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketVersioningValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketVersioning(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetBucketVersioningUpdateEndpoint(stack, options); err != nil {
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
	return nil
}

func (m *GetBucketVersioningInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *GetBucketVersioningInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opGetBucketVersioningMiddleware struct {
}

func (*endpointPrefix_opGetBucketVersioningMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetBucketVersioningMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetBucketVersioningInput)
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
func addEndpointPrefix_opGetBucketVersioningMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetBucketVersioningMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketVersioning(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketVersioning",
	}
}

func copyGetBucketVersioningInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetBucketVersioningInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetBucketVersioningInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getGetBucketVersioningARNMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketVersioningInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setGetBucketVersioningARNMember(input interface{}, v string) error {
	in := input.(*GetBucketVersioningInput)
	in.Bucket = &v
	return nil
}
func backFillGetBucketVersioningAccountID(input interface{}, v string) error {
	in := input.(*GetBucketVersioningInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetBucketVersioningUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getGetBucketVersioningARNMember,
			BackfillAccountID: backFillGetBucketVersioningAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setGetBucketVersioningARNMember,
			CopyInput:         copyGetBucketVersioningInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
