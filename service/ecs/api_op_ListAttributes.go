// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the attributes for Amazon ECS resources within a specified target type
// and cluster. When you specify a target type and cluster, ListAttributes returns
// a list of attribute objects, one for each attribute on each resource. You can
// filter the list of results to a single attribute name to only return results
// that have that name. You can also filter the results by attribute name and
// value. You can do this, for example, to see which container instances in a
// cluster are running a Linux AMI ( ecs.os-type=linux ).
func (c *Client) ListAttributes(ctx context.Context, params *ListAttributesInput, optFns ...func(*Options)) (*ListAttributesOutput, error) {
	if params == nil {
		params = &ListAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttributes", params, optFns, c.addOperationListAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttributesInput struct {

	// The type of the target to list attributes with.
	//
	// This member is required.
	TargetType types.TargetType

	// The name of the attribute to filter the results with.
	AttributeName *string

	// The value of the attribute to filter results with. You must also specify an
	// attribute name to use this parameter.
	AttributeValue *string

	// The short name or full Amazon Resource Name (ARN) of the cluster to list
	// attributes. If you do not specify a cluster, the default cluster is assumed.
	Cluster *string

	// The maximum number of cluster results that ListAttributes returned in paginated
	// output. When this parameter is used, ListAttributes only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListAttributes
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter isn't used, then ListAttributes returns up to 100 results and
	// a nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a ListAttributes request indicating that more
	// results are available to fulfill the request and further calls are needed. If
	// maxResults was provided, it's possible the number of results to be fewer than
	// maxResults . This token should be treated as an opaque identifier that is only
	// used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAttributesOutput struct {

	// A list of attribute objects that meet the criteria of the request.
	Attributes []types.Attribute

	// The nextToken value to include in a future ListAttributes request. When the
	// results of a ListAttributes request exceed maxResults , this value can be used
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAttributes{}, middleware.After)
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
	if err = addOpListAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttributes(options.Region), middleware.Before); err != nil {
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

// ListAttributesAPIClient is a client that implements the ListAttributes
// operation.
type ListAttributesAPIClient interface {
	ListAttributes(context.Context, *ListAttributesInput, ...func(*Options)) (*ListAttributesOutput, error)
}

var _ ListAttributesAPIClient = (*Client)(nil)

// ListAttributesPaginatorOptions is the paginator options for ListAttributes
type ListAttributesPaginatorOptions struct {
	// The maximum number of cluster results that ListAttributes returned in paginated
	// output. When this parameter is used, ListAttributes only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListAttributes
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter isn't used, then ListAttributes returns up to 100 results and
	// a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttributesPaginator is a paginator for ListAttributes
type ListAttributesPaginator struct {
	options   ListAttributesPaginatorOptions
	client    ListAttributesAPIClient
	params    *ListAttributesInput
	nextToken *string
	firstPage bool
}

// NewListAttributesPaginator returns a new ListAttributesPaginator
func NewListAttributesPaginator(client ListAttributesAPIClient, params *ListAttributesInput, optFns ...func(*ListAttributesPaginatorOptions)) *ListAttributesPaginator {
	if params == nil {
		params = &ListAttributesInput{}
	}

	options := ListAttributesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttributesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttributesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttributes page.
func (p *ListAttributesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttributesOutput, error) {
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

	result, err := p.client.ListAttributes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListAttributes",
	}
}
