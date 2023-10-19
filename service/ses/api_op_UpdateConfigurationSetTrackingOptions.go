// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an association between a configuration set and a custom domain for
// open and click event tracking. By default, images and links used for tracking
// open and click events are hosted on domains operated by Amazon SES. You can
// configure a subdomain of your own to handle these events. For information about
// using custom domains, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/configure-custom-open-click-domains.html)
// .
func (c *Client) UpdateConfigurationSetTrackingOptions(ctx context.Context, params *UpdateConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*UpdateConfigurationSetTrackingOptionsOutput, error) {
	if params == nil {
		params = &UpdateConfigurationSetTrackingOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfigurationSetTrackingOptions", params, optFns, c.addOperationUpdateConfigurationSetTrackingOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfigurationSetTrackingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to update the tracking options for a configuration set.
type UpdateConfigurationSetTrackingOptionsInput struct {

	// The name of the configuration set.
	//
	// This member is required.
	ConfigurationSetName *string

	// A domain that is used to redirect email recipients to an Amazon SES-operated
	// domain. This domain captures open and click events generated by Amazon SES
	// emails. For more information, see Configuring Custom Domains to Handle Open and
	// Click Tracking (https://docs.aws.amazon.com/ses/latest/dg/configure-custom-open-click-domains.html)
	// in the Amazon SES Developer Guide.
	//
	// This member is required.
	TrackingOptions *types.TrackingOptions

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type UpdateConfigurationSetTrackingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConfigurationSetTrackingOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateConfigurationSetTrackingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateConfigurationSetTrackingOptions{}, middleware.After)
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
	if err = addOpUpdateConfigurationSetTrackingOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfigurationSetTrackingOptions(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateConfigurationSetTrackingOptions",
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

func newServiceMetadataMiddleware_opUpdateConfigurationSetTrackingOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "UpdateConfigurationSetTrackingOptions",
	}
}
