// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new WorkSpace image from an existing WorkSpace.
func (c *Client) CreateWorkspaceImage(ctx context.Context, params *CreateWorkspaceImageInput, optFns ...func(*Options)) (*CreateWorkspaceImageOutput, error) {
	if params == nil {
		params = &CreateWorkspaceImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkspaceImage", params, optFns, c.addOperationCreateWorkspaceImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkspaceImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkspaceImageInput struct {

	// The description of the new WorkSpace image.
	//
	// This member is required.
	Description *string

	// The name of the new WorkSpace image.
	//
	// This member is required.
	Name *string

	// The identifier of the source WorkSpace
	//
	// This member is required.
	WorkspaceId *string

	// The tags that you want to add to the new WorkSpace image. To add tags when
	// you're creating the image, you must create an IAM policy that grants your IAM
	// user permission to use workspaces:CreateTags .
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateWorkspaceImageOutput struct {

	// The date when the image was created.
	Created *time.Time

	// The description of the image.
	Description *string

	// The identifier of the new WorkSpace image.
	ImageId *string

	// The name of the image.
	Name *string

	// The operating system that the image is running.
	OperatingSystem *types.OperatingSystem

	// The identifier of the Amazon Web Services account that owns the image.
	OwnerAccountId *string

	// Specifies whether the image is running on dedicated hardware. When Bring Your
	// Own License (BYOL) is enabled, this value is set to DEDICATED. For more
	// information, see Bring Your Own Windows Desktop Images. (https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.htm)
	// .
	RequiredTenancy types.WorkspaceImageRequiredTenancy

	// The availability status of the image.
	State types.WorkspaceImageState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkspaceImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkspaceImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkspaceImage{}, middleware.After)
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
	if err = addOpCreateWorkspaceImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkspaceImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkspaceImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "CreateWorkspaceImage",
	}
}
