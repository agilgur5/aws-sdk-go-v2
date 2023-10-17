// Code generated by smithy-go-codegen DO NOT EDIT.

package kendraranking

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendraranking/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a rescore execution plan. A rescore execution plan is an Amazon Kendra
// Intelligent Ranking resource used for provisioning the Rescore API. You can
// update the number of capacity units you require for Amazon Kendra Intelligent
// Ranking to rescore or re-rank a search service's results.
func (c *Client) UpdateRescoreExecutionPlan(ctx context.Context, params *UpdateRescoreExecutionPlanInput, optFns ...func(*Options)) (*UpdateRescoreExecutionPlanOutput, error) {
	if params == nil {
		params = &UpdateRescoreExecutionPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRescoreExecutionPlan", params, optFns, c.addOperationUpdateRescoreExecutionPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRescoreExecutionPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRescoreExecutionPlanInput struct {

	// The identifier of the rescore execution plan that you want to update.
	//
	// This member is required.
	Id *string

	// You can set additional capacity units to meet the needs of your rescore
	// execution plan. You are given a single capacity unit by default. If you want to
	// use the default capacity, you don't set additional capacity units. For more
	// information on the default capacity and additional capacity units, see
	// Adjusting capacity (https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html)
	// .
	CapacityUnits *types.CapacityUnitsConfiguration

	// A new description for the rescore execution plan.
	Description *string

	// A new name for the rescore execution plan.
	Name *string

	noSmithyDocumentSerde
}

type UpdateRescoreExecutionPlanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRescoreExecutionPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateRescoreExecutionPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateRescoreExecutionPlan{}, middleware.After)
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
	if err = addOpUpdateRescoreExecutionPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRescoreExecutionPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRescoreExecutionPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra-ranking",
		OperationName: "UpdateRescoreExecutionPlan",
	}
}
