// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorad/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the group access control entries for a template.
func (c *Client) GetTemplateGroupAccessControlEntry(ctx context.Context, params *GetTemplateGroupAccessControlEntryInput, optFns ...func(*Options)) (*GetTemplateGroupAccessControlEntryOutput, error) {
	if params == nil {
		params = &GetTemplateGroupAccessControlEntryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTemplateGroupAccessControlEntry", params, optFns, c.addOperationGetTemplateGroupAccessControlEntryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTemplateGroupAccessControlEntryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTemplateGroupAccessControlEntryInput struct {

	// Security identifier (SID) of the group object from Active Directory. The SID
	// starts with "S-".
	//
	// This member is required.
	GroupSecurityIdentifier *string

	// The Amazon Resource Name (ARN) that was returned when you called CreateTemplate (https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html)
	// .
	//
	// This member is required.
	TemplateArn *string

	noSmithyDocumentSerde
}

type GetTemplateGroupAccessControlEntryOutput struct {

	// An access control entry allows or denies an Active Directory group from
	// enrolling and/or autoenrolling with a template.
	AccessControlEntry *types.AccessControlEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTemplateGroupAccessControlEntryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTemplateGroupAccessControlEntry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTemplateGroupAccessControlEntry{}, middleware.After)
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
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetTemplateGroupAccessControlEntryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplateGroupAccessControlEntry(options.Region), middleware.Before); err != nil {
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
	err = stack.Serialize.Insert(&resolveAuthSchemeMiddleware{
		operation: "GetTemplateGroupAccessControlEntry",
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

func newServiceMetadataMiddleware_opGetTemplateGroupAccessControlEntry(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pca-connector-ad",
		OperationName: "GetTemplateGroupAccessControlEntry",
	}
}
