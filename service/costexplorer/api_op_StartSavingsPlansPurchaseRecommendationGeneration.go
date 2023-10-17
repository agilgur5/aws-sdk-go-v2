// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests a Savings Plans recommendation generation. This enables you to
// calculate a fresh set of Savings Plans recommendations that takes your latest
// usage data and current Savings Plans inventory into account. You can refresh
// Savings Plans recommendations up to three times daily for a consolidated billing
// family. StartSavingsPlansPurchaseRecommendationGeneration has no request syntax
// because no input parameters are needed to support this operation.
func (c *Client) StartSavingsPlansPurchaseRecommendationGeneration(ctx context.Context, params *StartSavingsPlansPurchaseRecommendationGenerationInput, optFns ...func(*Options)) (*StartSavingsPlansPurchaseRecommendationGenerationOutput, error) {
	if params == nil {
		params = &StartSavingsPlansPurchaseRecommendationGenerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSavingsPlansPurchaseRecommendationGeneration", params, optFns, c.addOperationStartSavingsPlansPurchaseRecommendationGenerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSavingsPlansPurchaseRecommendationGenerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSavingsPlansPurchaseRecommendationGenerationInput struct {
	noSmithyDocumentSerde
}

type StartSavingsPlansPurchaseRecommendationGenerationOutput struct {

	// The estimated time for when the recommendation generation will complete.
	EstimatedCompletionTime *string

	// The start time of the recommendation generation.
	GenerationStartedTime *string

	// The ID for this specific recommendation.
	RecommendationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSavingsPlansPurchaseRecommendationGenerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartSavingsPlansPurchaseRecommendationGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSavingsPlansPurchaseRecommendationGeneration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSavingsPlansPurchaseRecommendationGeneration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartSavingsPlansPurchaseRecommendationGeneration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "StartSavingsPlansPurchaseRecommendationGeneration",
	}
}
