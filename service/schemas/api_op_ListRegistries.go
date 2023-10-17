// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the registries.
func (c *Client) ListRegistries(ctx context.Context, params *ListRegistriesInput, optFns ...func(*Options)) (*ListRegistriesOutput, error) {
	if params == nil {
		params = &ListRegistriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRegistries", params, optFns, c.addOperationListRegistriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRegistriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegistriesInput struct {
	Limit *int32

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string

	// Specifying this limits the results to only those registry names that start with
	// the specified prefix.
	RegistryNamePrefix *string

	// Can be set to Local or AWS to limit responses to your custom registries, or the
	// ones provided by AWS.
	Scope *string

	noSmithyDocumentSerde
}

type ListRegistriesOutput struct {

	// The token that specifies the next page of results to return. To request the
	// first page, leave NextToken empty. The token will expire in 24 hours, and cannot
	// be shared with other accounts.
	NextToken *string

	// An array of registry summaries.
	Registries []types.RegistrySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRegistriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRegistries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRegistries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRegistries(options.Region), middleware.Before); err != nil {
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

// ListRegistriesAPIClient is a client that implements the ListRegistries
// operation.
type ListRegistriesAPIClient interface {
	ListRegistries(context.Context, *ListRegistriesInput, ...func(*Options)) (*ListRegistriesOutput, error)
}

var _ ListRegistriesAPIClient = (*Client)(nil)

// ListRegistriesPaginatorOptions is the paginator options for ListRegistries
type ListRegistriesPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRegistriesPaginator is a paginator for ListRegistries
type ListRegistriesPaginator struct {
	options   ListRegistriesPaginatorOptions
	client    ListRegistriesAPIClient
	params    *ListRegistriesInput
	nextToken *string
	firstPage bool
}

// NewListRegistriesPaginator returns a new ListRegistriesPaginator
func NewListRegistriesPaginator(client ListRegistriesAPIClient, params *ListRegistriesInput, optFns ...func(*ListRegistriesPaginatorOptions)) *ListRegistriesPaginator {
	if params == nil {
		params = &ListRegistriesInput{}
	}

	options := ListRegistriesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRegistriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRegistriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRegistries page.
func (p *ListRegistriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRegistriesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListRegistries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRegistries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "ListRegistries",
	}
}
