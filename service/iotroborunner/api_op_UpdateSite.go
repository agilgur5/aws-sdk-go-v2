// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Grants permission to update a site
func (c *Client) UpdateSite(ctx context.Context, params *UpdateSiteInput, optFns ...func(*Options)) (*UpdateSiteOutput, error) {
	if params == nil {
		params = &UpdateSiteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSite", params, optFns, c.addOperationUpdateSiteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSiteInput struct {

	// Site ARN.
	//
	// This member is required.
	Id *string

	// A valid ISO 3166-1 alpha-2 code for the country in which the site resides.
	// e.g., US.
	CountryCode *string

	// A high-level description of the site.
	Description *string

	// Human friendly name of the resource.
	Name *string

	noSmithyDocumentSerde
}

type UpdateSiteOutput struct {

	// Site ARN.
	//
	// This member is required.
	Arn *string

	// Filters access by the site's identifier
	//
	// This member is required.
	Id *string

	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string

	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// A valid ISO 3166-1 alpha-2 code for the country in which the site resides.
	// e.g., US.
	CountryCode *string

	// A high-level description of the site.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSiteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSite{}, middleware.After)
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
	if err = addOpUpdateSiteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSite(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateSite",
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

func newServiceMetadataMiddleware_opUpdateSite(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotroborunner",
		OperationName: "UpdateSite",
	}
}
