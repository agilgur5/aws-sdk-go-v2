// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a Virtual Private Cloud (VPC) configuration to the application.
// Applications can use VPCs to store and access resources securely. Note the
// following about VPC configurations for Kinesis Data Analytics applications:
//   - VPC configurations are not supported for SQL applications.
//   - When a VPC is added to a Kinesis Data Analytics application, the
//     application can no longer be accessed from the Internet directly. To enable
//     Internet access to the application, add an Internet gateway to your VPC.
func (c *Client) AddApplicationVpcConfiguration(ctx context.Context, params *AddApplicationVpcConfigurationInput, optFns ...func(*Options)) (*AddApplicationVpcConfigurationOutput, error) {
	if params == nil {
		params = &AddApplicationVpcConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationVpcConfiguration", params, optFns, c.addOperationAddApplicationVpcConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationVpcConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationVpcConfigurationInput struct {

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// Description of the VPC to add to the application.
	//
	// This member is required.
	VpcConfiguration *types.VpcConfiguration

	// A value you use to implement strong concurrency for application updates. You
	// must provide the ApplicationVersionID or the ConditionalToken . You get the
	// application's current ConditionalToken using DescribeApplication . For better
	// concurrency support, use the ConditionalToken parameter instead of
	// CurrentApplicationVersionId .
	ConditionalToken *string

	// The version of the application to which you want to add the VPC configuration.
	// You must provide the CurrentApplicationVersionId or the ConditionalToken . You
	// can use the DescribeApplication operation to get the current application
	// version. If the version specified is not the current version, the
	// ConcurrentModificationException is returned. For better concurrency support, use
	// the ConditionalToken parameter instead of CurrentApplicationVersionId .
	CurrentApplicationVersionId *int64

	noSmithyDocumentSerde
}

type AddApplicationVpcConfigurationOutput struct {

	// The ARN of the application.
	ApplicationARN *string

	// Provides the current application version. Kinesis Data Analytics updates the
	// ApplicationVersionId each time you update the application.
	ApplicationVersionId *int64

	// The parameters of the new VPC configuration.
	VpcConfigurationDescription *types.VpcConfigurationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddApplicationVpcConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationVpcConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationVpcConfiguration{}, middleware.After)
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
	if err = addOpAddApplicationVpcConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationVpcConfiguration(options.Region), middleware.Before); err != nil {
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
		operation: "AddApplicationVpcConfiguration",
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

func newServiceMetadataMiddleware_opAddApplicationVpcConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationVpcConfiguration",
	}
}
