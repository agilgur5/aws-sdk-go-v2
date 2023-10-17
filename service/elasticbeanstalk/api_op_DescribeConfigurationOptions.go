// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the configuration options that are used in a particular configuration
// template or environment, or that a specified solution stack defines. The
// description includes the values the options, their default values, and an
// indication of the required action on a running environment if an option value is
// changed.
func (c *Client) DescribeConfigurationOptions(ctx context.Context, params *DescribeConfigurationOptionsInput, optFns ...func(*Options)) (*DescribeConfigurationOptionsOutput, error) {
	if params == nil {
		params = &DescribeConfigurationOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationOptions", params, optFns, c.addOperationDescribeConfigurationOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Result message containing a list of application version descriptions.
type DescribeConfigurationOptionsInput struct {

	// The name of the application associated with the configuration template or
	// environment. Only needed if you want to describe the configuration options
	// associated with either the configuration template or environment.
	ApplicationName *string

	// The name of the environment whose configuration options you want to describe.
	EnvironmentName *string

	// If specified, restricts the descriptions to only the specified options.
	Options []types.OptionSpecification

	// The ARN of the custom platform.
	PlatformArn *string

	// The name of the solution stack whose configuration options you want to describe.
	SolutionStackName *string

	// The name of the configuration template whose configuration options you want to
	// describe.
	TemplateName *string

	noSmithyDocumentSerde
}

// Describes the settings for a specified configuration set.
type DescribeConfigurationOptionsOutput struct {

	// A list of ConfigurationOptionDescription .
	Options []types.ConfigurationOptionDescription

	// The ARN of the platform version.
	PlatformArn *string

	// The name of the solution stack these configuration options belong to.
	SolutionStackName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeConfigurationOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeConfigurationOptions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfigurationOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeConfigurationOptions",
	}
}
