// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a response headers policy configuration. To get a response headers policy
// configuration, you must provide the policy's identifier. If the response headers
// policy is attached to a distribution's cache behavior, you can get the policy's
// identifier using ListDistributions or GetDistribution . If the response headers
// policy is not attached to a cache behavior, you can get the identifier using
// ListResponseHeadersPolicies .
func (c *Client) GetResponseHeadersPolicyConfig(ctx context.Context, params *GetResponseHeadersPolicyConfigInput, optFns ...func(*Options)) (*GetResponseHeadersPolicyConfigOutput, error) {
	if params == nil {
		params = &GetResponseHeadersPolicyConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResponseHeadersPolicyConfig", params, optFns, c.addOperationGetResponseHeadersPolicyConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResponseHeadersPolicyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResponseHeadersPolicyConfigInput struct {

	// The identifier for the response headers policy. If the response headers policy
	// is attached to a distribution's cache behavior, you can get the policy's
	// identifier using ListDistributions or GetDistribution . If the response headers
	// policy is not attached to a cache behavior, you can get the identifier using
	// ListResponseHeadersPolicies .
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetResponseHeadersPolicyConfigOutput struct {

	// The version identifier for the current version of the response headers policy.
	ETag *string

	// Contains a response headers policy.
	ResponseHeadersPolicyConfig *types.ResponseHeadersPolicyConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResponseHeadersPolicyConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetResponseHeadersPolicyConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetResponseHeadersPolicyConfig{}, middleware.After)
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
	if err = addOpGetResponseHeadersPolicyConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResponseHeadersPolicyConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResponseHeadersPolicyConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetResponseHeadersPolicyConfig",
	}
}
