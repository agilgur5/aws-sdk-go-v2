// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a prefetch schedule for a playback configuration. A prefetch schedule
// allows you to tell MediaTailor to fetch and prepare certain ads before an ad
// break happens. For more information about ad prefetching, see Using ad
// prefetching (https://docs.aws.amazon.com/mediatailor/latest/ug/prefetching-ads.html)
// in the MediaTailor User Guide.
func (c *Client) GetPrefetchSchedule(ctx context.Context, params *GetPrefetchScheduleInput, optFns ...func(*Options)) (*GetPrefetchScheduleOutput, error) {
	if params == nil {
		params = &GetPrefetchScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPrefetchSchedule", params, optFns, c.addOperationGetPrefetchScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPrefetchScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPrefetchScheduleInput struct {

	// The name of the prefetch schedule. The name must be unique among all prefetch
	// schedules that are associated with the specified playback configuration.
	//
	// This member is required.
	Name *string

	// Returns information about the prefetch schedule for a specific playback
	// configuration. If you call GetPrefetchSchedule on an expired prefetch schedule,
	// MediaTailor returns an HTTP 404 status code.
	//
	// This member is required.
	PlaybackConfigurationName *string

	noSmithyDocumentSerde
}

type GetPrefetchScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the prefetch schedule.
	Arn *string

	// Consumption settings determine how, and when, MediaTailor places the prefetched
	// ads into ad breaks. Ad consumption occurs within a span of time that you define,
	// called a consumption window. You can designate which ad breaks that MediaTailor
	// fills with prefetch ads by setting avail matching criteria.
	Consumption *types.PrefetchConsumption

	// The name of the prefetch schedule. The name must be unique among all prefetch
	// schedules that are associated with the specified playback configuration.
	Name *string

	// The name of the playback configuration to create the prefetch schedule for.
	PlaybackConfigurationName *string

	// A complex type that contains settings for prefetch retrieval from the ad
	// decision server (ADS).
	Retrieval *types.PrefetchRetrieval

	// An optional stream identifier that you can specify in order to prefetch for
	// multiple streams that use the same playback configuration.
	StreamId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPrefetchScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPrefetchSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPrefetchSchedule{}, middleware.After)
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
	if err = addOpGetPrefetchScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPrefetchSchedule(options.Region), middleware.Before); err != nil {
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
		operation: "GetPrefetchSchedule",
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

func newServiceMetadataMiddleware_opGetPrefetchSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "GetPrefetchSchedule",
	}
}
