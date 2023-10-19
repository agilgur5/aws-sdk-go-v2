// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of log groups that contain log anomalies.
func (c *Client) ListAnomalousLogGroups(ctx context.Context, params *ListAnomalousLogGroupsInput, optFns ...func(*Options)) (*ListAnomalousLogGroupsOutput, error) {
	if params == nil {
		params = &ListAnomalousLogGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnomalousLogGroups", params, optFns, c.addOperationListAnomalousLogGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnomalousLogGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnomalousLogGroupsInput struct {

	// The ID of the insight containing the log groups.
	//
	// This member is required.
	InsightId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAnomalousLogGroupsOutput struct {

	// The list of Amazon CloudWatch log groups that are related to an insight.
	//
	// This member is required.
	AnomalousLogGroups []types.AnomalousLogGroup

	// The ID of the insight containing the log groups.
	//
	// This member is required.
	InsightId *string

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnomalousLogGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnomalousLogGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnomalousLogGroups{}, middleware.After)
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
	if err = addOpListAnomalousLogGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnomalousLogGroups(options.Region), middleware.Before); err != nil {
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
		operation: "ListAnomalousLogGroups",
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

// ListAnomalousLogGroupsAPIClient is a client that implements the
// ListAnomalousLogGroups operation.
type ListAnomalousLogGroupsAPIClient interface {
	ListAnomalousLogGroups(context.Context, *ListAnomalousLogGroupsInput, ...func(*Options)) (*ListAnomalousLogGroupsOutput, error)
}

var _ ListAnomalousLogGroupsAPIClient = (*Client)(nil)

// ListAnomalousLogGroupsPaginatorOptions is the paginator options for
// ListAnomalousLogGroups
type ListAnomalousLogGroupsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnomalousLogGroupsPaginator is a paginator for ListAnomalousLogGroups
type ListAnomalousLogGroupsPaginator struct {
	options   ListAnomalousLogGroupsPaginatorOptions
	client    ListAnomalousLogGroupsAPIClient
	params    *ListAnomalousLogGroupsInput
	nextToken *string
	firstPage bool
}

// NewListAnomalousLogGroupsPaginator returns a new ListAnomalousLogGroupsPaginator
func NewListAnomalousLogGroupsPaginator(client ListAnomalousLogGroupsAPIClient, params *ListAnomalousLogGroupsInput, optFns ...func(*ListAnomalousLogGroupsPaginatorOptions)) *ListAnomalousLogGroupsPaginator {
	if params == nil {
		params = &ListAnomalousLogGroupsInput{}
	}

	options := ListAnomalousLogGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnomalousLogGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnomalousLogGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnomalousLogGroups page.
func (p *ListAnomalousLogGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnomalousLogGroupsOutput, error) {
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

	result, err := p.client.ListAnomalousLogGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAnomalousLogGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "ListAnomalousLogGroups",
	}
}
