// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an association between a Merged API and source API using the source
// API's identifier.
func (c *Client) AssociateMergedGraphqlApi(ctx context.Context, params *AssociateMergedGraphqlApiInput, optFns ...func(*Options)) (*AssociateMergedGraphqlApiOutput, error) {
	if params == nil {
		params = &AssociateMergedGraphqlApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateMergedGraphqlApi", params, optFns, c.addOperationAssociateMergedGraphqlApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateMergedGraphqlApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateMergedGraphqlApiInput struct {

	// The identifier of the AppSync Merged API. This is generated by the AppSync
	// service. In most cases, Merged APIs (especially in your account) only require
	// the API ID value or ARN of the merged API. However, Merged APIs in other
	// accounts (cross-account use cases) strictly require the full resource ARN of the
	// merged API.
	//
	// This member is required.
	MergedApiIdentifier *string

	// The identifier of the AppSync Source API. This is generated by the AppSync
	// service. In most cases, source APIs (especially in your account) only require
	// the API ID value or ARN of the source API. However, source APIs from other
	// accounts (cross-account use cases) strictly require the full resource ARN of the
	// source API.
	//
	// This member is required.
	SourceApiIdentifier *string

	// The description field.
	Description *string

	// The SourceApiAssociationConfig object data.
	SourceApiAssociationConfig *types.SourceApiAssociationConfig

	noSmithyDocumentSerde
}

type AssociateMergedGraphqlApiOutput struct {

	// The SourceApiAssociation object data.
	SourceApiAssociation *types.SourceApiAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateMergedGraphqlApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateMergedGraphqlApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateMergedGraphqlApi{}, middleware.After)
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
	if err = addOpAssociateMergedGraphqlApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateMergedGraphqlApi(options.Region), middleware.Before); err != nil {
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
		operation: "AssociateMergedGraphqlApi",
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

func newServiceMetadataMiddleware_opAssociateMergedGraphqlApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "AssociateMergedGraphqlApi",
	}
}
