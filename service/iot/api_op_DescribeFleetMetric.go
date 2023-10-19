// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about the specified fleet metric. Requires permission to
// access the DescribeFleetMetric (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DescribeFleetMetric(ctx context.Context, params *DescribeFleetMetricInput, optFns ...func(*Options)) (*DescribeFleetMetricOutput, error) {
	if params == nil {
		params = &DescribeFleetMetricInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetMetric", params, optFns, c.addOperationDescribeFleetMetricMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetMetricOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetMetricInput struct {

	// The name of the fleet metric to describe.
	//
	// This member is required.
	MetricName *string

	noSmithyDocumentSerde
}

type DescribeFleetMetricOutput struct {

	// The field to aggregate.
	AggregationField *string

	// The type of the aggregation query.
	AggregationType *types.AggregationType

	// The date when the fleet metric is created.
	CreationDate *time.Time

	// The fleet metric description.
	Description *string

	// The name of the index to search.
	IndexName *string

	// The date when the fleet metric is last modified.
	LastModifiedDate *time.Time

	// The ARN of the fleet metric to describe.
	MetricArn *string

	// The name of the fleet metric to describe.
	MetricName *string

	// The time in seconds between fleet metric emissions. Range [60(1 min), 86400(1
	// day)] and must be multiple of 60.
	Period *int32

	// The search query string.
	QueryString *string

	// The query version.
	QueryVersion *string

	// Used to support unit transformation such as milliseconds to seconds. The unit
	// must be supported by CW metric (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html)
	// .
	Unit types.FleetMetricUnit

	// The version of the fleet metric.
	Version int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetMetricMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFleetMetric{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFleetMetric{}, middleware.After)
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
	if err = addOpDescribeFleetMetricValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetMetric(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeFleetMetric",
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

func newServiceMetadataMiddleware_opDescribeFleetMetric(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "DescribeFleetMetric",
	}
}
