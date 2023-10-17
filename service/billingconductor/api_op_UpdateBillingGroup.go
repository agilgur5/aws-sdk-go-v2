// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This updates an existing billing group.
func (c *Client) UpdateBillingGroup(ctx context.Context, params *UpdateBillingGroupInput, optFns ...func(*Options)) (*UpdateBillingGroupOutput, error) {
	if params == nil {
		params = &UpdateBillingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBillingGroup", params, optFns, c.addOperationUpdateBillingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBillingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBillingGroupInput struct {

	// The Amazon Resource Name (ARN) of the billing group being updated.
	//
	// This member is required.
	Arn *string

	// Specifies if the billing group has automatic account association ( AutoAssociate
	// ) enabled.
	AccountGrouping *types.UpdateBillingGroupAccountGrouping

	// The preferences and settings that will be used to compute the Amazon Web
	// Services charges for a billing group.
	ComputationPreference *types.ComputationPreference

	// A description of the billing group.
	Description *string

	// The name of the billing group. The names must be unique to each billing group.
	Name *string

	// The status of the billing group. Only one of the valid values can be used.
	Status types.BillingGroupStatus

	noSmithyDocumentSerde
}

type UpdateBillingGroupOutput struct {

	// Specifies if the billing group has automatic account association ( AutoAssociate
	// ) enabled.
	AccountGrouping *types.UpdateBillingGroupAccountGrouping

	// The Amazon Resource Name (ARN) of the billing group that was updated.
	Arn *string

	// A description of the billing group.
	Description *string

	// The most recent time when the billing group was modified.
	LastModifiedTime int64

	// The name of the billing group. The names must be unique to each billing group.
	Name *string

	// The Amazon Resource Name (ARN) of the pricing plan to compute Amazon Web
	// Services charges for the billing group.
	PricingPlanArn *string

	// The account ID that serves as the main account in a billing group.
	PrimaryAccountId *string

	// The number of accounts in the particular billing group.
	Size int64

	// The status of the billing group. Only one of the valid values can be used.
	Status types.BillingGroupStatus

	// The reason why the billing group is in its current status.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBillingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBillingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBillingGroup{}, middleware.After)
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
	if err = addOpUpdateBillingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBillingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBillingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "UpdateBillingGroup",
	}
}
