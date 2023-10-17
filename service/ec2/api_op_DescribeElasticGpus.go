// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the Elastic Graphics accelerator associated with your instances. For
// more information about Elastic Graphics, see Amazon Elastic Graphics (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html)
// .
func (c *Client) DescribeElasticGpus(ctx context.Context, params *DescribeElasticGpusInput, optFns ...func(*Options)) (*DescribeElasticGpusOutput, error) {
	if params == nil {
		params = &DescribeElasticGpusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeElasticGpus", params, optFns, c.addOperationDescribeElasticGpusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeElasticGpusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeElasticGpusInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The Elastic Graphics accelerator IDs.
	ElasticGpuIds []string

	// The filters.
	//   - availability-zone - The Availability Zone in which the Elastic Graphics
	//   accelerator resides.
	//   - elastic-gpu-health - The status of the Elastic Graphics accelerator ( OK |
	//   IMPAIRED ).
	//   - elastic-gpu-state - The state of the Elastic Graphics accelerator ( ATTACHED
	//   ).
	//   - elastic-gpu-type - The type of Elastic Graphics accelerator; for example,
	//   eg1.medium .
	//   - instance-id - The ID of the instance to which the Elastic Graphics
	//   accelerator is associated.
	Filters []types.Filter

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeElasticGpusOutput struct {

	// Information about the Elastic Graphics accelerators.
	ElasticGpuSet []types.ElasticGpus

	// The total number of items to return. If the total number of items available is
	// more than the value specified in max-items then a Next-Token will be provided in
	// the output that you can use to resume pagination.
	MaxResults *int32

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeElasticGpusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeElasticGpus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeElasticGpus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticGpus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeElasticGpus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeElasticGpus",
	}
}
