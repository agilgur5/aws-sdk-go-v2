// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a rule to control sampling behavior for instrumented applications.
// Services retrieve rules with GetSamplingRules (https://docs.aws.amazon.com/xray/latest/api/API_GetSamplingRules.html)
// , and evaluate each rule in ascending order of priority for each request. If a
// rule matches, the service records a trace, borrowing it from the reservoir size.
// After 10 seconds, the service reports back to X-Ray with GetSamplingTargets (https://docs.aws.amazon.com/xray/latest/api/API_GetSamplingTargets.html)
// to get updated versions of each in-use rule. The updated rule contains a trace
// quota that the service can use instead of borrowing from the reservoir.
func (c *Client) CreateSamplingRule(ctx context.Context, params *CreateSamplingRuleInput, optFns ...func(*Options)) (*CreateSamplingRuleOutput, error) {
	if params == nil {
		params = &CreateSamplingRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSamplingRule", params, optFns, c.addOperationCreateSamplingRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSamplingRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSamplingRuleInput struct {

	// The rule definition.
	//
	// This member is required.
	SamplingRule *types.SamplingRule

	// A map that contains one or more tag keys and tag values to attach to an X-Ray
	// sampling rule. For more information about ways to use tags, see Tagging Amazon
	// Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the Amazon Web Services General Reference. The following restrictions apply
	// to tags:
	//   - Maximum number of user-applied tags per resource: 50
	//   - Maximum tag key length: 128 Unicode characters
	//   - Maximum tag value length: 256 Unicode characters
	//   - Valid values for key and value: a-z, A-Z, 0-9, space, and the following
	//   characters: _ . : / = + - and @
	//   - Tag keys and values are case sensitive.
	//   - Don't use aws: as a prefix for keys; it's reserved for Amazon Web Services
	//   use.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateSamplingRuleOutput struct {

	// The saved rule definition and metadata.
	SamplingRuleRecord *types.SamplingRuleRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSamplingRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSamplingRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSamplingRule{}, middleware.After)
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
	if err = addOpCreateSamplingRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSamplingRule(options.Region), middleware.Before); err != nil {
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
		operation: "CreateSamplingRule",
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

func newServiceMetadataMiddleware_opCreateSamplingRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "CreateSamplingRule",
	}
}
