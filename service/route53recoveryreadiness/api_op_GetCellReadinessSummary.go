// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets readiness for a cell. Aggregates the readiness of all the resources that
// are associated with the cell into a single value.
func (c *Client) GetCellReadinessSummary(ctx context.Context, params *GetCellReadinessSummaryInput, optFns ...func(*Options)) (*GetCellReadinessSummaryOutput, error) {
	if params == nil {
		params = &GetCellReadinessSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCellReadinessSummary", params, optFns, c.addOperationGetCellReadinessSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCellReadinessSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCellReadinessSummaryInput struct {

	// The name of the cell.
	//
	// This member is required.
	CellName *string

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCellReadinessSummaryOutput struct {

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// The readiness at a cell level.
	Readiness types.Readiness

	// Summaries for the readiness checks that make up the cell.
	ReadinessChecks []types.ReadinessCheckSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCellReadinessSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCellReadinessSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCellReadinessSummary{}, middleware.After)
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
	if err = addOpGetCellReadinessSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCellReadinessSummary(options.Region), middleware.Before); err != nil {
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

// GetCellReadinessSummaryAPIClient is a client that implements the
// GetCellReadinessSummary operation.
type GetCellReadinessSummaryAPIClient interface {
	GetCellReadinessSummary(context.Context, *GetCellReadinessSummaryInput, ...func(*Options)) (*GetCellReadinessSummaryOutput, error)
}

var _ GetCellReadinessSummaryAPIClient = (*Client)(nil)

// GetCellReadinessSummaryPaginatorOptions is the paginator options for
// GetCellReadinessSummary
type GetCellReadinessSummaryPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCellReadinessSummaryPaginator is a paginator for GetCellReadinessSummary
type GetCellReadinessSummaryPaginator struct {
	options   GetCellReadinessSummaryPaginatorOptions
	client    GetCellReadinessSummaryAPIClient
	params    *GetCellReadinessSummaryInput
	nextToken *string
	firstPage bool
}

// NewGetCellReadinessSummaryPaginator returns a new
// GetCellReadinessSummaryPaginator
func NewGetCellReadinessSummaryPaginator(client GetCellReadinessSummaryAPIClient, params *GetCellReadinessSummaryInput, optFns ...func(*GetCellReadinessSummaryPaginatorOptions)) *GetCellReadinessSummaryPaginator {
	if params == nil {
		params = &GetCellReadinessSummaryInput{}
	}

	options := GetCellReadinessSummaryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCellReadinessSummaryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCellReadinessSummaryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCellReadinessSummary page.
func (p *GetCellReadinessSummaryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCellReadinessSummaryOutput, error) {
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

	result, err := p.client.GetCellReadinessSummary(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCellReadinessSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "GetCellReadinessSummary",
	}
}
