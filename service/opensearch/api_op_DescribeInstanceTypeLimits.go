// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the instance count, storage, and master node limits for a given
// OpenSearch or Elasticsearch version and instance type.
func (c *Client) DescribeInstanceTypeLimits(ctx context.Context, params *DescribeInstanceTypeLimitsInput, optFns ...func(*Options)) (*DescribeInstanceTypeLimitsOutput, error) {
	if params == nil {
		params = &DescribeInstanceTypeLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceTypeLimits", params, optFns, c.addOperationDescribeInstanceTypeLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceTypeLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeInstanceTypeLimits operation.
type DescribeInstanceTypeLimitsInput struct {

	// Version of OpenSearch or Elasticsearch, in the format Elasticsearch_X.Y or
	// OpenSearch_X.Y. Defaults to the latest version of OpenSearch.
	//
	// This member is required.
	EngineVersion *string

	// The OpenSearch Service instance type for which you need limit information.
	//
	// This member is required.
	InstanceType types.OpenSearchPartitionInstanceType

	// The name of the domain. Only specify if you need the limits for an existing
	// domain.
	DomainName *string

	noSmithyDocumentSerde
}

// Container for the parameters received from the DescribeInstanceTypeLimits
// operation.
type DescribeInstanceTypeLimitsOutput struct {

	// Map that contains all applicable instance type limits. data refers to data
	// nodes. master refers to dedicated master nodes.
	LimitsByRole map[string]types.Limits

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceTypeLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInstanceTypeLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInstanceTypeLimits{}, middleware.After)
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
	if err = addOpDescribeInstanceTypeLimitsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceTypeLimits(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstanceTypeLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeInstanceTypeLimits",
	}
}
