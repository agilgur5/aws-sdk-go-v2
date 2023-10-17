// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deregisters the specified directory. This operation is asynchronous and returns
// before the WorkSpace directory is deregistered. If any WorkSpaces are registered
// to this directory, you must remove them before you can deregister the directory.
// Simple AD and AD Connector are made available to you free of charge to use with
// WorkSpaces. If there are no WorkSpaces being used with your Simple AD or AD
// Connector directory for 30 consecutive days, this directory will be
// automatically deregistered for use with Amazon WorkSpaces, and you will be
// charged for this directory as per the Directory Service pricing terms (http://aws.amazon.com/directoryservice/pricing/)
// . To delete empty directories, see Delete the Directory for Your WorkSpaces (https://docs.aws.amazon.com/workspaces/latest/adminguide/delete-workspaces-directory.html)
// . If you delete your Simple AD or AD Connector directory, you can always create
// a new one when you want to start using WorkSpaces again.
func (c *Client) DeregisterWorkspaceDirectory(ctx context.Context, params *DeregisterWorkspaceDirectoryInput, optFns ...func(*Options)) (*DeregisterWorkspaceDirectoryOutput, error) {
	if params == nil {
		params = &DeregisterWorkspaceDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterWorkspaceDirectory", params, optFns, c.addOperationDeregisterWorkspaceDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterWorkspaceDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterWorkspaceDirectoryInput struct {

	// The identifier of the directory. If any WorkSpaces are registered to this
	// directory, you must remove them before you deregister the directory, or you will
	// receive an OperationNotSupportedException error.
	//
	// This member is required.
	DirectoryId *string

	noSmithyDocumentSerde
}

type DeregisterWorkspaceDirectoryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterWorkspaceDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterWorkspaceDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterWorkspaceDirectory{}, middleware.After)
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
	if err = addOpDeregisterWorkspaceDirectoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterWorkspaceDirectory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterWorkspaceDirectory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DeregisterWorkspaceDirectory",
	}
}
