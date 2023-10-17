// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports the specified Windows 10 or 11 Bring Your Own License (BYOL) image into
// Amazon WorkSpaces. The image must be an already licensed Amazon EC2 image that
// is in your Amazon Web Services account, and you must own the image. For more
// information about creating BYOL images, see Bring Your Own Windows Desktop
// Licenses (https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html)
// .
func (c *Client) ImportWorkspaceImage(ctx context.Context, params *ImportWorkspaceImageInput, optFns ...func(*Options)) (*ImportWorkspaceImageOutput, error) {
	if params == nil {
		params = &ImportWorkspaceImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportWorkspaceImage", params, optFns, c.addOperationImportWorkspaceImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportWorkspaceImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportWorkspaceImageInput struct {

	// The identifier of the EC2 image.
	//
	// This member is required.
	Ec2ImageId *string

	// The description of the WorkSpace image.
	//
	// This member is required.
	ImageDescription *string

	// The name of the WorkSpace image.
	//
	// This member is required.
	ImageName *string

	// The ingestion process to be used when importing the image, depending on which
	// protocol you want to use for your BYOL Workspace image, either PCoIP, WorkSpaces
	// Streaming Protocol (WSP), or bring your own protocol (BYOP). To use WSP, specify
	// a value that ends in _WSP . To use PCoIP, specify a value that does not end in
	// _WSP . To use BYOP, specify a value that ends in _BYOP . For non-GPU-enabled
	// bundles (bundles other than Graphics or GraphicsPro), specify BYOL_REGULAR ,
	// BYOL_REGULAR_WSP , or BYOL_REGULAR_BYOP , depending on the protocol. The
	// BYOL_REGULAR_BYOP and BYOL_GRAPHICS_G4DN_BYOP values are only supported by
	// Amazon WorkSpaces Core. Contact your account team to be allow-listed to use
	// these values. For more information, see Amazon WorkSpaces Core (http://aws.amazon.com/workspaces/core/)
	// .
	//
	// This member is required.
	IngestionProcess types.WorkspaceImageIngestionProcess

	// If specified, the version of Microsoft Office to subscribe to. Valid only for
	// Windows 10 and 11 BYOL images. For more information about subscribing to Office
	// for BYOL images, see Bring Your Own Windows Desktop Licenses (https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html)
	// .
	//   - Although this parameter is an array, only one item is allowed at this time.
	//   - Windows 11 only supports Microsoft_Office_2019 .
	Applications []types.Application

	// The tags. Each WorkSpaces resource can have a maximum of 50 tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ImportWorkspaceImageOutput struct {

	// The identifier of the WorkSpace image.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportWorkspaceImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportWorkspaceImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportWorkspaceImage{}, middleware.After)
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
	if err = addOpImportWorkspaceImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportWorkspaceImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportWorkspaceImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ImportWorkspaceImage",
	}
}
