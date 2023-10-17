// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a pricing rule can be associated to a pricing plan, or a set of pricing
// plans.
func (c *Client) CreatePricingRule(ctx context.Context, params *CreatePricingRuleInput, optFns ...func(*Options)) (*CreatePricingRuleOutput, error) {
	if params == nil {
		params = &CreatePricingRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePricingRule", params, optFns, c.addOperationCreatePricingRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePricingRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePricingRuleInput struct {

	// The pricing rule name. The names must be unique to each pricing rule.
	//
	// This member is required.
	Name *string

	// The scope of pricing rule that indicates if it's globally applicable, or it's
	// service-specific.
	//
	// This member is required.
	Scope types.PricingRuleScope

	// The type of pricing rule.
	//
	// This member is required.
	Type types.PricingRuleType

	// The seller of services provided by Amazon Web Services, their affiliates, or
	// third-party providers selling services via Amazon Web Services Marketplace.
	BillingEntity *string

	// The token that's needed to support idempotency. Idempotency isn't currently
	// supported, but will be implemented in a future update.
	ClientToken *string

	// The pricing rule description.
	Description *string

	// A percentage modifier that's applied on the public pricing rates.
	ModifierPercentage *float64

	// Operation is the specific Amazon Web Services action covered by this line item.
	// This describes the specific usage of the line item. If the Scope attribute is
	// set to SKU , this attribute indicates which operation the PricingRule is
	// modifying. For example, a value of RunInstances:0202 indicates the operation of
	// running an Amazon EC2 instance.
	Operation *string

	// If the Scope attribute is set to SERVICE or SKU , the attribute indicates which
	// service the PricingRule is applicable for.
	Service *string

	// A map that contains tag keys and tag values that are attached to a pricing rule.
	Tags map[string]string

	// The set of tiering configurations for the pricing rule.
	Tiering *types.CreateTieringInput

	// Usage type is the unit that each service uses to measure the usage of a
	// specific type of resource. If the Scope attribute is set to SKU , this attribute
	// indicates which usage type the PricingRule is modifying. For example,
	// USW2-BoxUsage:m2.2xlarge describes an M2 High Memory Double Extra Large
	// instance in the US West (Oregon) Region.
	UsageType *string

	noSmithyDocumentSerde
}

type CreatePricingRuleOutput struct {

	// The Amazon Resource Name (ARN) of the created pricing rule.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePricingRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePricingRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePricingRule{}, middleware.After)
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
	if err = addIdempotencyToken_opCreatePricingRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePricingRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePricingRule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePricingRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePricingRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePricingRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePricingRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePricingRuleInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePricingRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePricingRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePricingRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "CreatePricingRule",
	}
}
