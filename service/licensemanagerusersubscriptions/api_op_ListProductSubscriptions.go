// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the user-based subscription products available from an identity provider.
func (c *Client) ListProductSubscriptions(ctx context.Context, params *ListProductSubscriptionsInput, optFns ...func(*Options)) (*ListProductSubscriptionsOutput, error) {
	if params == nil {
		params = &ListProductSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProductSubscriptions", params, optFns, c.addOperationListProductSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProductSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProductSubscriptionsInput struct {

	// An object that specifies details for the identity provider.
	//
	// This member is required.
	IdentityProvider types.IdentityProvider

	// The name of the user-based subscription product.
	//
	// This member is required.
	Product *string

	// An array of structures that you can use to filter the results to those that
	// match one or more sets of key-value pairs that you specify.
	Filters []types.Filter

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProductSubscriptionsOutput struct {

	// Token for the next set of results.
	NextToken *string

	// Metadata that describes the list product subscriptions operation.
	ProductUserSummaries []types.ProductUserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProductSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProductSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProductSubscriptions{}, middleware.After)
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
	if err = addOpListProductSubscriptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProductSubscriptions(options.Region), middleware.Before); err != nil {
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
		operation: "ListProductSubscriptions",
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

// ListProductSubscriptionsAPIClient is a client that implements the
// ListProductSubscriptions operation.
type ListProductSubscriptionsAPIClient interface {
	ListProductSubscriptions(context.Context, *ListProductSubscriptionsInput, ...func(*Options)) (*ListProductSubscriptionsOutput, error)
}

var _ ListProductSubscriptionsAPIClient = (*Client)(nil)

// ListProductSubscriptionsPaginatorOptions is the paginator options for
// ListProductSubscriptions
type ListProductSubscriptionsPaginatorOptions struct {
	// Maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProductSubscriptionsPaginator is a paginator for ListProductSubscriptions
type ListProductSubscriptionsPaginator struct {
	options   ListProductSubscriptionsPaginatorOptions
	client    ListProductSubscriptionsAPIClient
	params    *ListProductSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewListProductSubscriptionsPaginator returns a new
// ListProductSubscriptionsPaginator
func NewListProductSubscriptionsPaginator(client ListProductSubscriptionsAPIClient, params *ListProductSubscriptionsInput, optFns ...func(*ListProductSubscriptionsPaginatorOptions)) *ListProductSubscriptionsPaginator {
	if params == nil {
		params = &ListProductSubscriptionsInput{}
	}

	options := ListProductSubscriptionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProductSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProductSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProductSubscriptions page.
func (p *ListProductSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProductSubscriptionsOutput, error) {
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

	result, err := p.client.ListProductSubscriptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProductSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager-user-subscriptions",
		OperationName: "ListProductSubscriptions",
	}
}
