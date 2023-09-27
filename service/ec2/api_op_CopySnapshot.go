// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies a point-in-time snapshot of an EBS volume and stores it in Amazon S3.
// You can copy a snapshot within the same Region, from one Region to another, or
// from a Region to an Outpost. You can't copy a snapshot from an Outpost to a
// Region, from one Outpost to another, or within the same Outpost. You can use the
// snapshot to create EBS volumes or Amazon Machine Images (AMIs). When copying
// snapshots to a Region, copies of encrypted EBS snapshots remain encrypted.
// Copies of unencrypted snapshots remain unencrypted, unless you enable encryption
// for the snapshot copy operation. By default, encrypted snapshot copies use the
// default Key Management Service (KMS) KMS key; however, you can specify a
// different KMS key. To copy an encrypted snapshot that has been shared from
// another account, you must have permissions for the KMS key used to encrypt the
// snapshot. Snapshots copied to an Outpost are encrypted by default using the
// default encryption key for the Region, or a different key that you specify in
// the request using KmsKeyId. Outposts do not support unencrypted snapshots. For
// more information, Amazon EBS local snapshots on Outposts (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshots-outposts.html#ami)
// in the Amazon Elastic Compute Cloud User Guide. Snapshots created by copying
// another snapshot have an arbitrary volume ID that should not be used for any
// purpose. For more information, see Copy an Amazon EBS snapshot (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-copy-snapshot.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*Options)) (*CopySnapshotOutput, error) {
	if params == nil {
		params = &CopySnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopySnapshot", params, optFns, c.addOperationCopySnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopySnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopySnapshotInput struct {

	// The ID of the Region that contains the snapshot to be copied.
	//
	// This member is required.
	SourceRegion *string

	// The ID of the EBS snapshot to copy.
	//
	// This member is required.
	SourceSnapshotId *string

	// A description for the EBS snapshot.
	Description *string

	// The Amazon Resource Name (ARN) of the Outpost to which to copy the snapshot.
	// Only specify this parameter when copying a snapshot from an Amazon Web Services
	// Region to an Outpost. The snapshot must be in the Region for the destination
	// Outpost. You cannot copy a snapshot from an Outpost to a Region, from one
	// Outpost to another, or within the same Outpost. For more information, see Copy
	// snapshots from an Amazon Web Services Region to an Outpost (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshots-outposts.html#copy-snapshots)
	// in the Amazon Elastic Compute Cloud User Guide.
	DestinationOutpostArn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// To encrypt a copy of an unencrypted snapshot if encryption by default is not
	// enabled, enable encryption using this parameter. Otherwise, omit this parameter.
	// Encrypted snapshots are encrypted, even if you omit this parameter and
	// encryption by default is not enabled. You cannot set this parameter to false.
	// For more information, see Amazon EBS encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool

	// The identifier of the Key Management Service (KMS) KMS key to use for Amazon
	// EBS encryption. If this parameter is not specified, your KMS key for Amazon EBS
	// is used. If KmsKeyId is specified, the encrypted state must be true . You can
	// specify the KMS key using any of the following:
	//   - Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
	//   - Key alias. For example, alias/ExampleAlias.
	//   - Key ARN. For example,
	//   arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//   - Alias ARN. For example,
	//   arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	// Amazon Web Services authenticates the KMS key asynchronously. Therefore, if you
	// specify an ID, alias, or ARN that is not valid, the action can appear to
	// complete, but eventually fails.
	KmsKeyId *string

	// When you copy an encrypted source snapshot using the Amazon EC2 Query API, you
	// must supply a pre-signed URL. This parameter is optional for unencrypted
	// snapshots. For more information, see Query requests (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html)
	// . The PresignedUrl should use the snapshot source endpoint, the CopySnapshot
	// action, and include the SourceRegion , SourceSnapshotId , and DestinationRegion
	// parameters. The PresignedUrl must be signed using Amazon Web Services Signature
	// Version 4. Because EBS snapshots are stored in Amazon S3, the signing algorithm
	// for this parameter uses the same logic that is described in Authenticating
	// Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// in the Amazon Simple Storage Service API Reference. An invalid or improperly
	// signed PresignedUrl will cause the copy operation to fail asynchronously, and
	// the snapshot will move to an error state.
	PresignedUrl *string

	// The tags to apply to the new snapshot.
	TagSpecifications []types.TagSpecification

	// Used by the SDK's PresignURL autofill customization to specify the region the
	// of the client's request.
	destinationRegion *string

	noSmithyDocumentSerde
}

func (*CopySnapshotInput) operationName() string {
	return "CopySnapshot"
}

type CopySnapshotOutput struct {

	// The ID of the new snapshot.
	SnapshotId *string

	// Any tags applied to the new snapshot.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopySnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCopySnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCopySnapshot{}, middleware.After)
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
	if err = addCopySnapshotPresignURLMiddleware(stack, options); err != nil {
		return err
	}
	if err = addCopySnapshotResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCopySnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopySnapshot(options.Region), middleware.Before); err != nil {
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

func copyCopySnapshotInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCopySnapshotPresignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	if input.PresignedUrl == nil || len(*input.PresignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PresignedUrl, true, nil
}
func getCopySnapshotSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCopySnapshotPresignedUrl(params interface{}, value string) error {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	input.PresignedUrl = &value
	return nil
}
func setCopySnapshotdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}
func addCopySnapshotPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL: getCopySnapshotPresignedUrl,

			GetSourceRegion: getCopySnapshotSourceRegion,

			CopyInput: copyCopySnapshotInputForPresign,

			SetDestinationRegion: setCopySnapshotdestinationRegion,

			SetPresignedURL: setCopySnapshotPresignedUrl,
		},
		Presigner: &presignAutoFillCopySnapshotClient{client: NewPresignClient(New(options))},
	})
}

type presignAutoFillCopySnapshotClient struct {
	client *PresignClient
}

// PresignURL is a middleware accessor that satisfies URLPresigner interface.
func (c *presignAutoFillCopySnapshotClient) PresignURL(ctx context.Context, srcRegion string, params interface{}) (*v4.PresignedHTTPRequest, error) {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = srcRegion
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	presignOptFn := WithPresignClientFromClientOptions(optFn)
	return c.client.PresignCopySnapshot(ctx, input, presignOptFn)
}

func newServiceMetadataMiddleware_opCopySnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CopySnapshot",
	}
}

// PresignCopySnapshot is used to generate a presigned HTTP Request which contains
// presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignCopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &CopySnapshotInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "CopySnapshot", params, clientOptFns,
		c.client.addOperationCopySnapshotMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}

type opCopySnapshotResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCopySnapshotResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCopySnapshotResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

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

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ec2"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

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
				signingName = "ec2"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
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
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ec2")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCopySnapshotResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCopySnapshotResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
