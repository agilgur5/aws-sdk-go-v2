// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a resource to the Resilience Hub application and assigns it to the
// specified Application Components. If you specify a new Application Component,
// Resilience Hub will automatically create the Application Component.
//   - This action has no effect outside Resilience Hub.
//   - This API updates the Resilience Hub application draft version. To use this
//     resource for running resiliency assessments, you must publish the Resilience Hub
//     application using the PublishAppVersion API.
//   - To update application version with new physicalResourceID , you must call
//     ResolveAppVersionResources API.
func (c *Client) CreateAppVersionResource(ctx context.Context, params *CreateAppVersionResourceInput, optFns ...func(*Options)) (*CreateAppVersionResourceOutput, error) {
	if params == nil {
		params = &CreateAppVersionResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppVersionResource", params, optFns, c.addOperationCreateAppVersionResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppVersionResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppVersionResourceInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// List of Application Components that this resource belongs to. If an Application
	// Component is not part of the Resilience Hub application, it will be added.
	//
	// This member is required.
	AppComponents []string

	// Logical identifier of the resource.
	//
	// This member is required.
	LogicalResourceId *types.LogicalResourceId

	// Physical identifier of the resource.
	//
	// This member is required.
	PhysicalResourceId *string

	// Type of resource.
	//
	// This member is required.
	ResourceType *string

	// Currently, there is no supported additional information for resources.
	AdditionalInfo map[string][]string

	// Amazon Web Services account that owns the physical resource.
	AwsAccountId *string

	// Amazon Web Services region that owns the physical resource.
	AwsRegion *string

	// Used for an idempotency token. A client token is a unique, case-sensitive
	// string of up to 64 ASCII characters. You should not reuse the same client token
	// for other API requests.
	ClientToken *string

	// Name of the resource.
	ResourceName *string

	noSmithyDocumentSerde
}

type CreateAppVersionResourceOutput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// Resilience Hub application version.
	//
	// This member is required.
	AppVersion *string

	// Defines a physical resource. A physical resource is a resource that exists in
	// your account. It can be identified using an Amazon Resource Name (ARN) or a
	// Resilience Hub-native identifier.
	PhysicalResource *types.PhysicalResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppVersionResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAppVersionResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAppVersionResource{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAppVersionResourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAppVersionResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppVersionResource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAppVersionResource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAppVersionResource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAppVersionResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAppVersionResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAppVersionResourceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAppVersionResourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAppVersionResource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAppVersionResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "CreateAppVersionResource",
	}
}
