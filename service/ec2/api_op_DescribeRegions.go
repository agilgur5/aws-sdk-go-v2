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

// Describes the Regions that are enabled for your account, or all Regions. For a
// list of the Regions supported by Amazon EC2, see Amazon Elastic Compute Cloud
// endpoints and quotas (https://docs.aws.amazon.com/general/latest/gr/ec2-service.html)
// . For information about enabling and disabling Regions for your account, see
// Managing Amazon Web Services Regions (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html)
// in the Amazon Web Services General Reference.
func (c *Client) DescribeRegions(ctx context.Context, params *DescribeRegionsInput, optFns ...func(*Options)) (*DescribeRegionsOutput, error) {
	if params == nil {
		params = &DescribeRegionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRegions", params, optFns, c.addOperationDescribeRegionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRegionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegionsInput struct {

	// Indicates whether to display all Regions, including Regions that are disabled
	// for your account.
	AllRegions *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - endpoint - The endpoint of the Region (for example,
	//   ec2.us-east-1.amazonaws.com ).
	//   - opt-in-status - The opt-in status of the Region ( opt-in-not-required |
	//   opted-in | not-opted-in ).
	//   - region-name - The name of the Region (for example, us-east-1 ).
	Filters []types.Filter

	// The names of the Regions. You can specify any Regions, whether they are enabled
	// and disabled for your account.
	RegionNames []string

	noSmithyDocumentSerde
}

type DescribeRegionsOutput struct {

	// Information about the Regions.
	Regions []types.Region

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRegionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeRegions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeRegions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRegions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeRegions",
	}
}
