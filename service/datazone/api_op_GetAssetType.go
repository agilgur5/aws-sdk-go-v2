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

// Gets an Amazon DataZone asset type.
func (c *Client) GetAssetType(ctx context.Context, params *GetAssetTypeInput, optFns ...func(*Options)) (*GetAssetTypeOutput, error) {
	if params == nil {
		params = &GetAssetTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssetType", params, optFns, c.addOperationGetAssetTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssetTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssetTypeInput struct {

	// The ID of the Amazon DataZone domain in which the asset type exists.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the asset type.
	//
	// This member is required.
	Identifier *string

	// The revision of the asset type.
	Revision *string

	noSmithyDocumentSerde
}

type GetAssetTypeOutput struct {

	// The ID of the Amazon DataZone domain in which the asset type exists.
	//
	// This member is required.
	DomainId *string

	// The metadata forms attached to the asset type.
	//
	// This member is required.
	FormsOutput map[string]types.FormEntryOutput

	// The name of the asset type.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon DataZone project that owns the asset type.
	//
	// This member is required.
	OwningProjectId *string

	// The revision of the asset type.
	//
	// This member is required.
	Revision *string

	// The timestamp of when the asset type was created.
	CreatedAt *time.Time

	// The Amazon DataZone user who created the asset type.
	CreatedBy *string

	// The description of the asset type.
	Description *string

	// The ID of the Amazon DataZone domain in which the asset type was originally
	// created.
	OriginDomainId *string

	// The ID of the Amazon DataZone project in which the asset type was originally
	// created.
	OriginProjectId *string

	// The timestamp of when the asset type was updated.
	UpdatedAt *time.Time

	// The Amazon DataZone user that updated the asset type.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAssetTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAssetType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssetType{}, middleware.After)
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
	if err = addOpGetAssetTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssetType(options.Region), middleware.Before); err != nil {
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
		operation: "GetAssetType",
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

func newServiceMetadataMiddleware_opGetAssetType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "GetAssetType",
	}
}
