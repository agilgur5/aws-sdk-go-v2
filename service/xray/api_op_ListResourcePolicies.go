// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of resource policies in the target Amazon Web Services account.
func (c *Client) ListResourcePolicies(ctx context.Context, params *ListResourcePoliciesInput, optFns ...func(*Options)) (*ListResourcePoliciesOutput, error) {
	if params == nil {
		params = &ListResourcePoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourcePolicies", params, optFns, c.addOperationListResourcePoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcePoliciesInput struct {

	// Not currently supported.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourcePoliciesOutput struct {

	// Pagination token. Not currently supported.
	NextToken *string

	// The list of resource policies in the target Amazon Web Services account.
	ResourcePolicies []types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourcePoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListResourcePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListResourcePolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourcePolicies(options.Region), middleware.Before); err != nil {
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

// ListResourcePoliciesAPIClient is a client that implements the
// ListResourcePolicies operation.
type ListResourcePoliciesAPIClient interface {
	ListResourcePolicies(context.Context, *ListResourcePoliciesInput, ...func(*Options)) (*ListResourcePoliciesOutput, error)
}

var _ ListResourcePoliciesAPIClient = (*Client)(nil)

// ListResourcePoliciesPaginatorOptions is the paginator options for
// ListResourcePolicies
type ListResourcePoliciesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourcePoliciesPaginator is a paginator for ListResourcePolicies
type ListResourcePoliciesPaginator struct {
	options   ListResourcePoliciesPaginatorOptions
	client    ListResourcePoliciesAPIClient
	params    *ListResourcePoliciesInput
	nextToken *string
	firstPage bool
}

// NewListResourcePoliciesPaginator returns a new ListResourcePoliciesPaginator
func NewListResourcePoliciesPaginator(client ListResourcePoliciesAPIClient, params *ListResourcePoliciesInput, optFns ...func(*ListResourcePoliciesPaginatorOptions)) *ListResourcePoliciesPaginator {
	if params == nil {
		params = &ListResourcePoliciesInput{}
	}

	options := ListResourcePoliciesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourcePoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourcePoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourcePolicies page.
func (p *ListResourcePoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourcePoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListResourcePolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourcePolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "ListResourcePolicies",
	}
}
