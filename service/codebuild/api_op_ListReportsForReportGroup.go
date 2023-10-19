// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of ARNs for the reports that belong to a ReportGroup .
func (c *Client) ListReportsForReportGroup(ctx context.Context, params *ListReportsForReportGroupInput, optFns ...func(*Options)) (*ListReportsForReportGroupOutput, error) {
	if params == nil {
		params = &ListReportsForReportGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReportsForReportGroup", params, optFns, c.addOperationListReportsForReportGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReportsForReportGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReportsForReportGroupInput struct {

	// The ARN of the report group for which you want to return report ARNs.
	//
	// This member is required.
	ReportGroupArn *string

	// A ReportFilter object used to filter the returned reports.
	Filter *types.ReportFilter

	// The maximum number of paginated reports in this report group returned per
	// response. Use nextToken to iterate pages in the list of returned Report
	// objects. The default value is 100.
	MaxResults *int32

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults . If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// Use to specify whether the results are returned in ascending or descending
	// order.
	SortOrder types.SortOrderType

	noSmithyDocumentSerde
}

type ListReportsForReportGroupOutput struct {

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults . If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// The list of report ARNs.
	Reports []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReportsForReportGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReportsForReportGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReportsForReportGroup{}, middleware.After)
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
	if err = addOpListReportsForReportGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReportsForReportGroup(options.Region), middleware.Before); err != nil {
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
		operation: "ListReportsForReportGroup",
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

// ListReportsForReportGroupAPIClient is a client that implements the
// ListReportsForReportGroup operation.
type ListReportsForReportGroupAPIClient interface {
	ListReportsForReportGroup(context.Context, *ListReportsForReportGroupInput, ...func(*Options)) (*ListReportsForReportGroupOutput, error)
}

var _ ListReportsForReportGroupAPIClient = (*Client)(nil)

// ListReportsForReportGroupPaginatorOptions is the paginator options for
// ListReportsForReportGroup
type ListReportsForReportGroupPaginatorOptions struct {
	// The maximum number of paginated reports in this report group returned per
	// response. Use nextToken to iterate pages in the list of returned Report
	// objects. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReportsForReportGroupPaginator is a paginator for ListReportsForReportGroup
type ListReportsForReportGroupPaginator struct {
	options   ListReportsForReportGroupPaginatorOptions
	client    ListReportsForReportGroupAPIClient
	params    *ListReportsForReportGroupInput
	nextToken *string
	firstPage bool
}

// NewListReportsForReportGroupPaginator returns a new
// ListReportsForReportGroupPaginator
func NewListReportsForReportGroupPaginator(client ListReportsForReportGroupAPIClient, params *ListReportsForReportGroupInput, optFns ...func(*ListReportsForReportGroupPaginatorOptions)) *ListReportsForReportGroupPaginator {
	if params == nil {
		params = &ListReportsForReportGroupInput{}
	}

	options := ListReportsForReportGroupPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReportsForReportGroupPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReportsForReportGroupPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReportsForReportGroup page.
func (p *ListReportsForReportGroupPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReportsForReportGroupOutput, error) {
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

	result, err := p.client.ListReportsForReportGroup(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReportsForReportGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListReportsForReportGroup",
	}
}
