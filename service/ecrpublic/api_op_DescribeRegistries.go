// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details for a public registry.
func (c *Client) DescribeRegistries(ctx context.Context, params *DescribeRegistriesInput, optFns ...func(*Options)) (*DescribeRegistriesOutput, error) {
	if params == nil {
		params = &DescribeRegistriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRegistries", params, optFns, c.addOperationDescribeRegistriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRegistriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegistriesInput struct {

	// The maximum number of repository results that's returned by DescribeRegistries
	// in paginated output. When this parameter is used, DescribeRegistries only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another DescribeRegistries request with the returned nextToken value. This
	// value can be between 1 and 1000. If this parameter isn't used, then
	// DescribeRegistries returns up to 100 results and a nextToken value, if
	// applicable.
	MaxResults *int32

	// The nextToken value that's returned from a previous paginated DescribeRegistries
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. If there are no more results to return, this
	// value is null . This token should be treated as an opaque identifier that is
	// only used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeRegistriesOutput struct {

	// An object that contains the details for a public registry.
	//
	// This member is required.
	Registries []types.Registry

	// The nextToken value to include in a future DescribeRepositories request. If the
	// results of a DescribeRepositories request exceed maxResults , you can use this
	// value to retrieve the next page of results. If there are no more results, this
	// value is null .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRegistriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRegistries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRegistries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegistries(options.Region), middleware.Before); err != nil {
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

// DescribeRegistriesAPIClient is a client that implements the DescribeRegistries
// operation.
type DescribeRegistriesAPIClient interface {
	DescribeRegistries(context.Context, *DescribeRegistriesInput, ...func(*Options)) (*DescribeRegistriesOutput, error)
}

var _ DescribeRegistriesAPIClient = (*Client)(nil)

// DescribeRegistriesPaginatorOptions is the paginator options for
// DescribeRegistries
type DescribeRegistriesPaginatorOptions struct {
	// The maximum number of repository results that's returned by DescribeRegistries
	// in paginated output. When this parameter is used, DescribeRegistries only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another DescribeRegistries request with the returned nextToken value. This
	// value can be between 1 and 1000. If this parameter isn't used, then
	// DescribeRegistries returns up to 100 results and a nextToken value, if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRegistriesPaginator is a paginator for DescribeRegistries
type DescribeRegistriesPaginator struct {
	options   DescribeRegistriesPaginatorOptions
	client    DescribeRegistriesAPIClient
	params    *DescribeRegistriesInput
	nextToken *string
	firstPage bool
}

// NewDescribeRegistriesPaginator returns a new DescribeRegistriesPaginator
func NewDescribeRegistriesPaginator(client DescribeRegistriesAPIClient, params *DescribeRegistriesInput, optFns ...func(*DescribeRegistriesPaginatorOptions)) *DescribeRegistriesPaginator {
	if params == nil {
		params = &DescribeRegistriesInput{}
	}

	options := DescribeRegistriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRegistriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRegistriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRegistries page.
func (p *DescribeRegistriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRegistriesOutput, error) {
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

	result, err := p.client.DescribeRegistries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRegistries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr-public",
		OperationName: "DescribeRegistries",
	}
}
