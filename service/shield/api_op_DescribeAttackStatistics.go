// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the number and type of attacks Shield has detected
// in the last year for all resources that belong to your account, regardless of
// whether you've defined Shield protections for them. This operation is available
// to Shield customers as well as to Shield Advanced customers. The operation
// returns data for the time range of midnight UTC, one year ago, to midnight UTC,
// today. For example, if the current time is 2020-10-26 15:39:32 PDT , equal to
// 2020-10-26 22:39:32 UTC , then the time range for the attack data returned is
// from 2019-10-26 00:00:00 UTC to 2020-10-26 00:00:00 UTC . The time range
// indicates the period covered by the attack statistics data items.
func (c *Client) DescribeAttackStatistics(ctx context.Context, params *DescribeAttackStatisticsInput, optFns ...func(*Options)) (*DescribeAttackStatisticsOutput, error) {
	if params == nil {
		params = &DescribeAttackStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAttackStatistics", params, optFns, c.addOperationDescribeAttackStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAttackStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAttackStatisticsInput struct {
	noSmithyDocumentSerde
}

type DescribeAttackStatisticsOutput struct {

	// The data that describes the attacks detected during the time period.
	//
	// This member is required.
	DataItems []types.AttackStatisticsDataItem

	// The time range of the attack.
	//
	// This member is required.
	TimeRange *types.TimeRange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAttackStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAttackStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAttackStatistics{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAttackStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAttackStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DescribeAttackStatistics",
	}
}
