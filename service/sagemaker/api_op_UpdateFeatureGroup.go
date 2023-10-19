// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the feature group by either adding features or updating the online
// store configuration. Use one of the following request parameters at a time while
// using the UpdateFeatureGroup API. You can add features for your feature group
// using the FeatureAdditions request parameter. Features cannot be removed from a
// feature group. You can update the online store configuration by using the
// OnlineStoreConfig request parameter. If a TtlDuration is specified, the default
// TtlDuration applies for all records added to the feature group after the feature
// group is updated. If a record level TtlDuration exists from using the PutRecord
// API, the record level TtlDuration applies to that record instead of the default
// TtlDuration .
func (c *Client) UpdateFeatureGroup(ctx context.Context, params *UpdateFeatureGroupInput, optFns ...func(*Options)) (*UpdateFeatureGroupOutput, error) {
	if params == nil {
		params = &UpdateFeatureGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFeatureGroup", params, optFns, c.addOperationUpdateFeatureGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFeatureGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFeatureGroupInput struct {

	// The name or Amazon Resource Name (ARN) of the feature group that you're
	// updating.
	//
	// This member is required.
	FeatureGroupName *string

	// Updates the feature group. Updating a feature group is an asynchronous
	// operation. When you get an HTTP 200 response, you've made a valid request. It
	// takes some time after you've made a valid request for Feature Store to update
	// the feature group.
	FeatureAdditions []types.FeatureDefinition

	// Updates the feature group online store configuration.
	OnlineStoreConfig *types.OnlineStoreConfigUpdate

	noSmithyDocumentSerde
}

type UpdateFeatureGroupOutput struct {

	// The Amazon Resource Number (ARN) of the feature group that you're updating.
	//
	// This member is required.
	FeatureGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFeatureGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFeatureGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFeatureGroup{}, middleware.After)
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
	if err = addOpUpdateFeatureGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFeatureGroup(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateFeatureGroup",
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

func newServiceMetadataMiddleware_opUpdateFeatureGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateFeatureGroup",
	}
}
