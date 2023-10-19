// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Calculates the Spot placement score for a Region or Availability Zone based on
// the specified target capacity and compute requirements. You can specify your
// compute requirements either by using InstanceRequirementsWithMetadata and
// letting Amazon EC2 choose the optimal instance types to fulfill your Spot
// request, or you can specify the instance types by using InstanceTypes . For more
// information, see Spot placement score (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html)
// in the Amazon EC2 User Guide.
func (c *Client) GetSpotPlacementScores(ctx context.Context, params *GetSpotPlacementScoresInput, optFns ...func(*Options)) (*GetSpotPlacementScoresOutput, error) {
	if params == nil {
		params = &GetSpotPlacementScoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSpotPlacementScores", params, optFns, c.addOperationGetSpotPlacementScoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSpotPlacementScoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSpotPlacementScoresInput struct {

	// The target capacity.
	//
	// This member is required.
	TargetCapacity *int32

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The attributes for the instance types. When you specify instance attributes,
	// Amazon EC2 will identify instance types with those attributes. If you specify
	// InstanceRequirementsWithMetadata , you can't specify InstanceTypes .
	InstanceRequirementsWithMetadata *types.InstanceRequirementsWithMetadataRequest

	// The instance types. We recommend that you specify at least three instance
	// types. If you specify one or two instance types, or specify variations of a
	// single instance type (for example, an m3.xlarge with and without instance
	// storage), the returned placement score will always be low. If you specify
	// InstanceTypes , you can't specify InstanceRequirementsWithMetadata .
	InstanceTypes []string

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	// The Regions used to narrow down the list of Regions to be scored. Enter the
	// Region code, for example, us-east-1 .
	RegionNames []string

	// Specify true so that the response returns a list of scored Availability Zones.
	// Otherwise, the response returns a list of scored Regions. A list of scored
	// Availability Zones is useful if you want to launch all of your Spot capacity
	// into a single Availability Zone.
	SingleAvailabilityZone *bool

	// The unit for the target capacity. Default: units (translates to number of
	// instances)
	TargetCapacityUnitType types.TargetCapacityUnitType

	noSmithyDocumentSerde
}

type GetSpotPlacementScoresOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// The Spot placement score for the top 10 Regions or Availability Zones, scored
	// on a scale from 1 to 10. Each score  reflects how likely it is that each Region
	// or Availability Zone will succeed at fulfilling the specified target capacity
	// at the time of the Spot placement score request. A score of 10 means that your
	// Spot capacity request is highly likely to succeed in that Region or Availability
	// Zone. If you request a Spot placement score for Regions, a high score assumes
	// that your fleet request will be configured to use all Availability Zones and the
	// capacity-optimized allocation strategy. If you request a Spot placement score
	// for Availability Zones, a high score assumes that your fleet request will be
	// configured to use a single Availability Zone and the capacity-optimized
	// allocation strategy. Different  Regions or Availability Zones might return the
	// same score. The Spot placement score serves as a recommendation only. No score
	// guarantees that your Spot request will be fully or partially fulfilled.
	SpotPlacementScores []types.SpotPlacementScore

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSpotPlacementScoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetSpotPlacementScores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetSpotPlacementScores{}, middleware.After)
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
	if err = addOpGetSpotPlacementScoresValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSpotPlacementScores(options.Region), middleware.Before); err != nil {
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
		operation: "GetSpotPlacementScores",
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

// GetSpotPlacementScoresAPIClient is a client that implements the
// GetSpotPlacementScores operation.
type GetSpotPlacementScoresAPIClient interface {
	GetSpotPlacementScores(context.Context, *GetSpotPlacementScoresInput, ...func(*Options)) (*GetSpotPlacementScoresOutput, error)
}

var _ GetSpotPlacementScoresAPIClient = (*Client)(nil)

// GetSpotPlacementScoresPaginatorOptions is the paginator options for
// GetSpotPlacementScores
type GetSpotPlacementScoresPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSpotPlacementScoresPaginator is a paginator for GetSpotPlacementScores
type GetSpotPlacementScoresPaginator struct {
	options   GetSpotPlacementScoresPaginatorOptions
	client    GetSpotPlacementScoresAPIClient
	params    *GetSpotPlacementScoresInput
	nextToken *string
	firstPage bool
}

// NewGetSpotPlacementScoresPaginator returns a new GetSpotPlacementScoresPaginator
func NewGetSpotPlacementScoresPaginator(client GetSpotPlacementScoresAPIClient, params *GetSpotPlacementScoresInput, optFns ...func(*GetSpotPlacementScoresPaginatorOptions)) *GetSpotPlacementScoresPaginator {
	if params == nil {
		params = &GetSpotPlacementScoresInput{}
	}

	options := GetSpotPlacementScoresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSpotPlacementScoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSpotPlacementScoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetSpotPlacementScores page.
func (p *GetSpotPlacementScoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSpotPlacementScoresOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetSpotPlacementScores(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetSpotPlacementScores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetSpotPlacementScores",
	}
}
