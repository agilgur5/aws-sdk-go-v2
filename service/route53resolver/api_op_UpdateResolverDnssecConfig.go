// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing DNSSEC validation configuration. If there is no existing
// DNSSEC validation configuration, one is created.
func (c *Client) UpdateResolverDnssecConfig(ctx context.Context, params *UpdateResolverDnssecConfigInput, optFns ...func(*Options)) (*UpdateResolverDnssecConfigOutput, error) {
	if params == nil {
		params = &UpdateResolverDnssecConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResolverDnssecConfig", params, optFns, c.addOperationUpdateResolverDnssecConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResolverDnssecConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResolverDnssecConfigInput struct {

	// The ID of the virtual private cloud (VPC) that you're updating the DNSSEC
	// validation status for.
	//
	// This member is required.
	ResourceId *string

	// The new value that you are specifying for DNSSEC validation for the VPC. The
	// value can be ENABLE or DISABLE . Be aware that it can take time for a validation
	// status change to be completed.
	//
	// This member is required.
	Validation types.Validation

	noSmithyDocumentSerde
}

type UpdateResolverDnssecConfigOutput struct {

	// A complex type that contains settings for the specified DNSSEC configuration.
	ResolverDNSSECConfig *types.ResolverDnssecConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResolverDnssecConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateResolverDnssecConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateResolverDnssecConfig{}, middleware.After)
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
	if err = addOpUpdateResolverDnssecConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResolverDnssecConfig(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateResolverDnssecConfig",
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

func newServiceMetadataMiddleware_opUpdateResolverDnssecConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "UpdateResolverDnssecConfig",
	}
}
