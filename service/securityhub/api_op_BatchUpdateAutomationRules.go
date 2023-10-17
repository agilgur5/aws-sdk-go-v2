// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates one or more automation rules based on rule Amazon Resource Names (ARNs)
// and input parameters.
func (c *Client) BatchUpdateAutomationRules(ctx context.Context, params *BatchUpdateAutomationRulesInput, optFns ...func(*Options)) (*BatchUpdateAutomationRulesOutput, error) {
	if params == nil {
		params = &BatchUpdateAutomationRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateAutomationRules", params, optFns, c.addOperationBatchUpdateAutomationRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateAutomationRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateAutomationRulesInput struct {

	// An array of ARNs for the rules that are to be updated. Optionally, you can also
	// include RuleStatus and RuleOrder .
	//
	// This member is required.
	UpdateAutomationRulesRequestItems []types.UpdateAutomationRulesRequestItem

	noSmithyDocumentSerde
}

type BatchUpdateAutomationRulesOutput struct {

	// A list of properly processed rule ARNs.
	ProcessedAutomationRules []string

	// A list of objects containing RuleArn , ErrorCode , and ErrorMessage . This
	// parameter tells you which automation rules the request didn't update and why.
	UnprocessedAutomationRules []types.UnprocessedAutomationRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateAutomationRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateAutomationRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateAutomationRules{}, middleware.After)
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
	if err = addOpBatchUpdateAutomationRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateAutomationRules(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateAutomationRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchUpdateAutomationRules",
	}
}
