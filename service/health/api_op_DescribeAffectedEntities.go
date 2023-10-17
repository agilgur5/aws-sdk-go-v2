// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of entities that have been affected by the specified events,
// based on the specified filter criteria. Entities can refer to individual
// customer resources, groups of customer resources, or any other construct,
// depending on the Amazon Web Service. Events that have impact beyond that of the
// affected entities, or where the extent of impact is unknown, include at least
// one entity indicating this. At least one event ARN is required.
//   - This API operation uses pagination. Specify the nextToken parameter in the
//     next request to return more results.
//   - This operation supports resource-level permissions. You can use this
//     operation to allow or deny access to specific Health events. For more
//     information, see Resource- and action-based conditions (https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions)
//     in the Health User Guide.
func (c *Client) DescribeAffectedEntities(ctx context.Context, params *DescribeAffectedEntitiesInput, optFns ...func(*Options)) (*DescribeAffectedEntitiesOutput, error) {
	if params == nil {
		params = &DescribeAffectedEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAffectedEntities", params, optFns, c.addOperationDescribeAffectedEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAffectedEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAffectedEntitiesInput struct {

	// Values to narrow the results returned. At least one event ARN is required.
	//
	// This member is required.
	Filter *types.EntityFilter

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	MaxResults *int32

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAffectedEntitiesOutput struct {

	// The entities that match the filter criteria.
	Entities []types.AffectedEntity

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAffectedEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAffectedEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAffectedEntities{}, middleware.After)
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
	if err = addOpDescribeAffectedEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAffectedEntities(options.Region), middleware.Before); err != nil {
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

// DescribeAffectedEntitiesAPIClient is a client that implements the
// DescribeAffectedEntities operation.
type DescribeAffectedEntitiesAPIClient interface {
	DescribeAffectedEntities(context.Context, *DescribeAffectedEntitiesInput, ...func(*Options)) (*DescribeAffectedEntitiesOutput, error)
}

var _ DescribeAffectedEntitiesAPIClient = (*Client)(nil)

// DescribeAffectedEntitiesPaginatorOptions is the paginator options for
// DescribeAffectedEntities
type DescribeAffectedEntitiesPaginatorOptions struct {
	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAffectedEntitiesPaginator is a paginator for DescribeAffectedEntities
type DescribeAffectedEntitiesPaginator struct {
	options   DescribeAffectedEntitiesPaginatorOptions
	client    DescribeAffectedEntitiesAPIClient
	params    *DescribeAffectedEntitiesInput
	nextToken *string
	firstPage bool
}

// NewDescribeAffectedEntitiesPaginator returns a new
// DescribeAffectedEntitiesPaginator
func NewDescribeAffectedEntitiesPaginator(client DescribeAffectedEntitiesAPIClient, params *DescribeAffectedEntitiesInput, optFns ...func(*DescribeAffectedEntitiesPaginatorOptions)) *DescribeAffectedEntitiesPaginator {
	if params == nil {
		params = &DescribeAffectedEntitiesInput{}
	}

	options := DescribeAffectedEntitiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAffectedEntitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAffectedEntitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAffectedEntities page.
func (p *DescribeAffectedEntitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAffectedEntitiesOutput, error) {
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

	result, err := p.client.DescribeAffectedEntities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAffectedEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeAffectedEntities",
	}
}
