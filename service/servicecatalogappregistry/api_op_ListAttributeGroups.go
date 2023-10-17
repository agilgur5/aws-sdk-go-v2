// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all attribute groups which you have access to. Results are paginated.
func (c *Client) ListAttributeGroups(ctx context.Context, params *ListAttributeGroupsInput, optFns ...func(*Options)) (*ListAttributeGroupsOutput, error) {
	if params == nil {
		params = &ListAttributeGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttributeGroups", params, optFns, c.addOperationListAttributeGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttributeGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttributeGroupsInput struct {

	// The upper bound of the number of results to return (cannot exceed 25). If this
	// parameter is omitted, it defaults to 25. This value is optional.
	MaxResults *int32

	// The token to use to get the next page of results after a previous API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAttributeGroupsOutput struct {

	// This list of attribute groups.
	AttributeGroups []types.AttributeGroupSummary

	// The token to use to get the next page of results after a previous API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttributeGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAttributeGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAttributeGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttributeGroups(options.Region), middleware.Before); err != nil {
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

// ListAttributeGroupsAPIClient is a client that implements the
// ListAttributeGroups operation.
type ListAttributeGroupsAPIClient interface {
	ListAttributeGroups(context.Context, *ListAttributeGroupsInput, ...func(*Options)) (*ListAttributeGroupsOutput, error)
}

var _ ListAttributeGroupsAPIClient = (*Client)(nil)

// ListAttributeGroupsPaginatorOptions is the paginator options for
// ListAttributeGroups
type ListAttributeGroupsPaginatorOptions struct {
	// The upper bound of the number of results to return (cannot exceed 25). If this
	// parameter is omitted, it defaults to 25. This value is optional.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttributeGroupsPaginator is a paginator for ListAttributeGroups
type ListAttributeGroupsPaginator struct {
	options   ListAttributeGroupsPaginatorOptions
	client    ListAttributeGroupsAPIClient
	params    *ListAttributeGroupsInput
	nextToken *string
	firstPage bool
}

// NewListAttributeGroupsPaginator returns a new ListAttributeGroupsPaginator
func NewListAttributeGroupsPaginator(client ListAttributeGroupsAPIClient, params *ListAttributeGroupsInput, optFns ...func(*ListAttributeGroupsPaginatorOptions)) *ListAttributeGroupsPaginator {
	if params == nil {
		params = &ListAttributeGroupsInput{}
	}

	options := ListAttributeGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttributeGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttributeGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttributeGroups page.
func (p *ListAttributeGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttributeGroupsOutput, error) {
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

	result, err := p.client.ListAttributeGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAttributeGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListAttributeGroups",
	}
}
