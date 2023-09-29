// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Associates an access policy with the specified access point. Each access point
// can have only one policy, so a request made to this API replaces any existing
// policy associated with the specified access point. All Amazon S3 on Outposts
// REST API requests for this action require an additional parameter of
// x-amz-outpost-id to be passed with the request. In addition, you must use an S3
// on Outposts endpoint hostname prefix instead of s3-control . For an example of
// the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts
// endpoint hostname prefix and the x-amz-outpost-id derived by using the access
// point ARN, see the Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html#API_control_PutAccessPointPolicy_Examples)
// section. The following actions are related to PutAccessPointPolicy :
//   - GetAccessPointPolicy (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html)
//   - DeleteAccessPointPolicy (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html)
func (c *Client) PutAccessPointPolicy(ctx context.Context, params *PutAccessPointPolicyInput, optFns ...func(*Options)) (*PutAccessPointPolicyOutput, error) {
	if params == nil {
		params = &PutAccessPointPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccessPointPolicy", params, optFns, c.addOperationPutAccessPointPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccessPointPolicyInput struct {

	// The Amazon Web Services account ID for owner of the bucket associated with the
	// specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point that you want to associate with the specified
	// policy. For using this parameter with Amazon S3 on Outposts with the REST API,
	// you must specify the name and the x-amz-outpost-id as well. For using this
	// parameter with S3 on Outposts with the Amazon Web Services SDK and CLI, you must
	// specify the ARN of the access point accessed in the format
	// arn:aws:s3-outposts:::outpost//accesspoint/ . For example, to access the access
	// point reports-ap through Outpost my-outpost owned by account 123456789012 in
	// Region us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap
	// . The value must be URL encoded.
	//
	// This member is required.
	Name *string

	// The policy that you want to apply to the specified access point. For more
	// information about access point policies, see Managing data access with Amazon
	// S3 access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Policy *string

	noSmithyDocumentSerde
}

func (*PutAccessPointPolicyInput) operationName() string {
	return "PutAccessPointPolicy"
}

func (in *PutAccessPointPolicyInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.AccessPointName = in.Name
	p.RequiresAccountId = ptr.Bool(true)
}

type PutAccessPointPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccessPointPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutAccessPointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutAccessPointPolicy{}, middleware.After)
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
	if err = addEndpointPrefix_opPutAccessPointPolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutAccessPointPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccessPointPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutAccessPointPolicyUpdateEndpoint(stack, options); err != nil {
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

func (m *PutAccessPointPolicyInput) GetARNMember() (*string, bool) {
	if m.Name == nil {
		return nil, false
	}
	return m.Name, true
}

func (m *PutAccessPointPolicyInput) SetARNMember(v string) error {
	m.Name = &v
	return nil
}

type endpointPrefix_opPutAccessPointPolicyMiddleware struct {
}

func (*endpointPrefix_opPutAccessPointPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutAccessPointPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutAccessPointPolicyInput)
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
func addEndpointPrefix_opPutAccessPointPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPutAccessPointPolicyMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opPutAccessPointPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutAccessPointPolicy",
	}
}

func copyPutAccessPointPolicyInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*PutAccessPointPolicyInput)
	if !ok {
		return nil, fmt.Errorf("expect *PutAccessPointPolicyInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getPutAccessPointPolicyARNMember(input interface{}) (*string, bool) {
	in := input.(*PutAccessPointPolicyInput)
	if in.Name == nil {
		return nil, false
	}
	return in.Name, true
}
func setPutAccessPointPolicyARNMember(input interface{}, v string) error {
	in := input.(*PutAccessPointPolicyInput)
	in.Name = &v
	return nil
}
func backFillPutAccessPointPolicyAccountID(input interface{}, v string) error {
	in := input.(*PutAccessPointPolicyInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addPutAccessPointPolicyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getPutAccessPointPolicyARNMember,
			BackfillAccountID: backFillPutAccessPointPolicyAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setPutAccessPointPolicyARNMember,
			CopyInput:         copyPutAccessPointPolicyInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
