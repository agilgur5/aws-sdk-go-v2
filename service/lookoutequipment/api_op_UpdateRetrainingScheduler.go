// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a retraining scheduler.
func (c *Client) UpdateRetrainingScheduler(ctx context.Context, params *UpdateRetrainingSchedulerInput, optFns ...func(*Options)) (*UpdateRetrainingSchedulerOutput, error) {
	if params == nil {
		params = &UpdateRetrainingSchedulerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRetrainingScheduler", params, optFns, c.addOperationUpdateRetrainingSchedulerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRetrainingSchedulerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRetrainingSchedulerInput struct {

	// The name of the model whose retraining scheduler you want to update.
	//
	// This member is required.
	ModelName *string

	// The number of past days of data that will be used for retraining.
	LookbackWindow *string

	// Indicates how the service will use new models. In MANAGED mode, new models will
	// automatically be used for inference if they have better performance than the
	// current model. In MANUAL mode, the new models will not be used until they are
	// manually activated (https://docs.aws.amazon.com/lookout-for-equipment/latest/ug/versioning-model.html#model-activation)
	// .
	PromoteMode types.ModelPromoteMode

	// This parameter uses the ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Durations)
	// standard to set the frequency at which you want retraining to occur in terms of
	// Years, Months, and/or Days (note: other parameters like Time are not currently
	// supported). The minimum value is 30 days (P30D) and the maximum value is 1 year
	// (P1Y). For example, the following values are valid:
	//   - P3M15D – Every 3 months and 15 days
	//   - P2M – Every 2 months
	//   - P150D – Every 150 days
	RetrainingFrequency *string

	// The start date for the retraining scheduler. Lookout for Equipment truncates
	// the time you provide to the nearest UTC day.
	RetrainingStartDate *time.Time

	noSmithyDocumentSerde
}

type UpdateRetrainingSchedulerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRetrainingSchedulerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateRetrainingScheduler{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateRetrainingScheduler{}, middleware.After)
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
	if err = addOpUpdateRetrainingSchedulerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRetrainingScheduler(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateRetrainingScheduler",
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

func newServiceMetadataMiddleware_opUpdateRetrainingScheduler(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "UpdateRetrainingScheduler",
	}
}
