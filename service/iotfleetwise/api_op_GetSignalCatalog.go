// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a signal catalog.
func (c *Client) GetSignalCatalog(ctx context.Context, params *GetSignalCatalogInput, optFns ...func(*Options)) (*GetSignalCatalogOutput, error) {
	if params == nil {
		params = &GetSignalCatalogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSignalCatalog", params, optFns, c.addOperationGetSignalCatalogMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSignalCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSignalCatalogInput struct {

	// The name of the signal catalog to retrieve information about.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetSignalCatalogOutput struct {

	// The Amazon Resource Name (ARN) of the signal catalog.
	//
	// This member is required.
	Arn *string

	// The time the signal catalog was created in seconds since epoch (January 1, 1970
	// at midnight UTC time).
	//
	// This member is required.
	CreationTime *time.Time

	// The last time the signal catalog was modified.
	//
	// This member is required.
	LastModificationTime *time.Time

	// The name of the signal catalog.
	//
	// This member is required.
	Name *string

	// A brief description of the signal catalog.
	Description *string

	// The total number of network nodes specified in a signal catalog.
	NodeCounts *types.NodeCounts

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSignalCatalogMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetSignalCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetSignalCatalog{}, middleware.After)
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
	if err = addOpGetSignalCatalogValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSignalCatalog(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSignalCatalog(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "GetSignalCatalog",
	}
}
