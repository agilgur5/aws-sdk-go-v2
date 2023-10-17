// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets one or more outcomes. This is a paginated API. If you provide a null
// maxResults , this actions retrieves a maximum of 100 records per page. If you
// provide a maxResults , the value must be between 50 and 100. To get the next
// page results, provide the pagination token from the GetOutcomesResult as part
// of your request. A null pagination token fetches the records from the beginning.
func (c *Client) GetOutcomes(ctx context.Context, params *GetOutcomesInput, optFns ...func(*Options)) (*GetOutcomesOutput, error) {
	if params == nil {
		params = &GetOutcomesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOutcomes", params, optFns, c.addOperationGetOutcomesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOutcomesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOutcomesInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The name of the outcome or outcomes to get.
	Name *string

	// The next page token for the request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetOutcomesOutput struct {

	// The next page token for subsequent requests.
	NextToken *string

	// The outcomes.
	Outcomes []types.Outcome

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOutcomesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOutcomes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOutcomes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOutcomes(options.Region), middleware.Before); err != nil {
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

// GetOutcomesAPIClient is a client that implements the GetOutcomes operation.
type GetOutcomesAPIClient interface {
	GetOutcomes(context.Context, *GetOutcomesInput, ...func(*Options)) (*GetOutcomesOutput, error)
}

var _ GetOutcomesAPIClient = (*Client)(nil)

// GetOutcomesPaginatorOptions is the paginator options for GetOutcomes
type GetOutcomesPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetOutcomesPaginator is a paginator for GetOutcomes
type GetOutcomesPaginator struct {
	options   GetOutcomesPaginatorOptions
	client    GetOutcomesAPIClient
	params    *GetOutcomesInput
	nextToken *string
	firstPage bool
}

// NewGetOutcomesPaginator returns a new GetOutcomesPaginator
func NewGetOutcomesPaginator(client GetOutcomesAPIClient, params *GetOutcomesInput, optFns ...func(*GetOutcomesPaginatorOptions)) *GetOutcomesPaginator {
	if params == nil {
		params = &GetOutcomesInput{}
	}

	options := GetOutcomesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetOutcomesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetOutcomesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetOutcomes page.
func (p *GetOutcomesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetOutcomesOutput, error) {
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

	result, err := p.client.GetOutcomes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetOutcomes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetOutcomes",
	}
}
