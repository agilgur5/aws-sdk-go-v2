// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a metadata form type in Amazon DataZone.
func (c *Client) GetFormType(ctx context.Context, params *GetFormTypeInput, optFns ...func(*Options)) (*GetFormTypeOutput, error) {
	if params == nil {
		params = &GetFormTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFormType", params, optFns, c.addOperationGetFormTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFormTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFormTypeInput struct {

	// The ID of the Amazon DataZone domain in which this metadata form type exists.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the metadata form type.
	//
	// This member is required.
	FormTypeIdentifier *string

	// The revision of this metadata form type.
	Revision *string

	noSmithyDocumentSerde
}

type GetFormTypeOutput struct {

	// The ID of the Amazon DataZone domain in which this metadata form type exists.
	//
	// This member is required.
	DomainId *string

	// The model of the metadata form type.
	//
	// This member is required.
	Model types.Model

	// The name of the metadata form type.
	//
	// This member is required.
	Name *string

	// The revision of the metadata form type.
	//
	// This member is required.
	Revision *string

	// The timestamp of when this metadata form type was created.
	CreatedAt *time.Time

	// The Amazon DataZone user who created this metadata form type.
	CreatedBy *string

	// The description of the metadata form type.
	Description *string

	// The imports of the metadata form type.
	Imports []types.Import

	// The ID of the Amazon DataZone domain in which the metadata form type was
	// originally created.
	OriginDomainId *string

	// The ID of the project in which this metadata form type was originally created.
	OriginProjectId *string

	// The ID of the project that owns this metadata form type.
	OwningProjectId *string

	// The status of the metadata form type.
	Status types.FormTypeStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFormTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFormType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFormType{}, middleware.After)
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
	if err = addOpGetFormTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFormType(options.Region), middleware.Before); err != nil {
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
		operation: "GetFormType",
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

func newServiceMetadataMiddleware_opGetFormType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "GetFormType",
	}
}
