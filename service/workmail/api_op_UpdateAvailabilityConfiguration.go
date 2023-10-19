// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing AvailabilityConfiguration for the given WorkMail
// organization and domain.
func (c *Client) UpdateAvailabilityConfiguration(ctx context.Context, params *UpdateAvailabilityConfigurationInput, optFns ...func(*Options)) (*UpdateAvailabilityConfigurationOutput, error) {
	if params == nil {
		params = &UpdateAvailabilityConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAvailabilityConfiguration", params, optFns, c.addOperationUpdateAvailabilityConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAvailabilityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAvailabilityConfigurationInput struct {

	// The domain to which the provider applies the availability configuration.
	//
	// This member is required.
	DomainName *string

	// The WorkMail organization for which the AvailabilityConfiguration will be
	// updated.
	//
	// This member is required.
	OrganizationId *string

	// The EWS availability provider definition. The request must contain exactly one
	// provider definition, either EwsProvider or LambdaProvider . The previously
	// stored provider will be overridden by the one provided.
	EwsProvider *types.EwsAvailabilityProvider

	// The Lambda availability provider definition. The request must contain exactly
	// one provider definition, either EwsProvider or LambdaProvider . The previously
	// stored provider will be overridden by the one provided.
	LambdaProvider *types.LambdaAvailabilityProvider

	noSmithyDocumentSerde
}

type UpdateAvailabilityConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAvailabilityConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAvailabilityConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAvailabilityConfiguration{}, middleware.After)
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
	if err = addOpUpdateAvailabilityConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAvailabilityConfiguration(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateAvailabilityConfiguration",
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

func newServiceMetadataMiddleware_opUpdateAvailabilityConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "UpdateAvailabilityConfiguration",
	}
}
