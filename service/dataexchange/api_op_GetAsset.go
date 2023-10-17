// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation returns information about an asset.
func (c *Client) GetAsset(ctx context.Context, params *GetAssetInput, optFns ...func(*Options)) (*GetAssetOutput, error) {
	if params == nil {
		params = &GetAssetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAsset", params, optFns, c.addOperationGetAssetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssetInput struct {

	// The unique identifier for an asset.
	//
	// This member is required.
	AssetId *string

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string

	// The unique identifier for a revision.
	//
	// This member is required.
	RevisionId *string

	noSmithyDocumentSerde
}

type GetAssetOutput struct {

	// The ARN for the asset.
	Arn *string

	// Details about the asset.
	AssetDetails *types.AssetDetails

	// The type of asset that is added to a data set.
	AssetType types.AssetType

	// The date and time that the asset was created, in ISO 8601 format.
	CreatedAt *time.Time

	// The unique identifier for the data set associated with this asset.
	DataSetId *string

	// The unique identifier for the asset.
	Id *string

	// The name of the asset. When importing from Amazon S3, the Amazon S3 object key
	// is used as the asset name. When exporting to Amazon S3, the asset name is used
	// as default target Amazon S3 object key. When importing from Amazon API Gateway
	// API, the API name is used as the asset name. When importing from Amazon
	// Redshift, the datashare name is used as the asset name. When importing from AWS
	// Lake Formation, the static values of "Database(s) included in the LF-tag policy"
	// or "Table(s) included in the LF-tag policy" are used as the asset name.
	Name *string

	// The unique identifier for the revision associated with this asset.
	RevisionId *string

	// The asset ID of the owned asset corresponding to the entitled asset being
	// viewed. This parameter is returned when an asset owner is viewing the entitled
	// copy of its owned asset.
	SourceId *string

	// The date and time that the asset was last updated, in ISO 8601 format.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAssetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAsset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAsset{}, middleware.After)
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
	if err = addOpGetAssetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAsset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAsset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "GetAsset",
	}
}
