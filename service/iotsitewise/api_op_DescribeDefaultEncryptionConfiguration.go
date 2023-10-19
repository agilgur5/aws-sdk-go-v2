// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the default encryption configuration for the Amazon
// Web Services account in the default or specified Region. For more information,
// see Key management (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/key-management.html)
// in the IoT SiteWise User Guide.
func (c *Client) DescribeDefaultEncryptionConfiguration(ctx context.Context, params *DescribeDefaultEncryptionConfigurationInput, optFns ...func(*Options)) (*DescribeDefaultEncryptionConfigurationOutput, error) {
	if params == nil {
		params = &DescribeDefaultEncryptionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDefaultEncryptionConfiguration", params, optFns, c.addOperationDescribeDefaultEncryptionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDefaultEncryptionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDefaultEncryptionConfigurationInput struct {
	noSmithyDocumentSerde
}

type DescribeDefaultEncryptionConfigurationOutput struct {

	// The status of the account configuration. This contains the ConfigurationState .
	// If there's an error, it also contains the ErrorDetails .
	//
	// This member is required.
	ConfigurationStatus *types.ConfigurationStatus

	// The type of encryption used for the encryption configuration.
	//
	// This member is required.
	EncryptionType types.EncryptionType

	// The key ARN of the customer managed key used for KMS encryption if you use
	// KMS_BASED_ENCRYPTION .
	KmsKeyArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDefaultEncryptionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDefaultEncryptionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDefaultEncryptionConfiguration{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeDefaultEncryptionConfigurationMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDefaultEncryptionConfiguration(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeDefaultEncryptionConfiguration",
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

type endpointPrefix_opDescribeDefaultEncryptionConfigurationMiddleware struct {
}

func (*endpointPrefix_opDescribeDefaultEncryptionConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeDefaultEncryptionConfigurationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeDefaultEncryptionConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeDefaultEncryptionConfigurationMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDefaultEncryptionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeDefaultEncryptionConfiguration",
	}
}
