// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all of the email identities that are associated with your
// Amazon Web Services account. An identity can be either an email address or a
// domain. This operation returns identities that are verified as well as those
// that aren't. This operation returns identities that are associated with Amazon
// SES and Amazon Pinpoint.
func (c *Client) ListEmailIdentities(ctx context.Context, params *ListEmailIdentitiesInput, optFns ...func(*Options)) (*ListEmailIdentitiesOutput, error) {
	if params == nil {
		params = &ListEmailIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEmailIdentities", params, optFns, c.addOperationListEmailIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEmailIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to list all of the email identities associated with your Amazon Web
// Services account. This list includes identities that you've already verified,
// identities that are unverified, and identities that were verified in the past,
// but are no longer verified.
type ListEmailIdentitiesInput struct {

	// A token returned from a previous call to ListEmailIdentities to indicate the
	// position in the list of identities.
	NextToken *string

	// The number of results to show in a single call to ListEmailIdentities . If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results. The value you specify has to be at least 0, and can be no
	// more than 1000.
	PageSize *int32

	noSmithyDocumentSerde
}

// A list of all of the identities that you've attempted to verify, regardless of
// whether or not those identities were successfully verified.
type ListEmailIdentitiesOutput struct {

	// An array that includes all of the email identities associated with your Amazon
	// Web Services account.
	EmailIdentities []types.IdentityInfo

	// A token that indicates that there are additional configuration sets to list. To
	// view additional configuration sets, issue another request to ListEmailIdentities
	// , and pass this token in the NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEmailIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEmailIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEmailIdentities{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEmailIdentities(options.Region), middleware.Before); err != nil {
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
		operation: "ListEmailIdentities",
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

// ListEmailIdentitiesAPIClient is a client that implements the
// ListEmailIdentities operation.
type ListEmailIdentitiesAPIClient interface {
	ListEmailIdentities(context.Context, *ListEmailIdentitiesInput, ...func(*Options)) (*ListEmailIdentitiesOutput, error)
}

var _ ListEmailIdentitiesAPIClient = (*Client)(nil)

// ListEmailIdentitiesPaginatorOptions is the paginator options for
// ListEmailIdentities
type ListEmailIdentitiesPaginatorOptions struct {
	// The number of results to show in a single call to ListEmailIdentities . If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results. The value you specify has to be at least 0, and can be no
	// more than 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEmailIdentitiesPaginator is a paginator for ListEmailIdentities
type ListEmailIdentitiesPaginator struct {
	options   ListEmailIdentitiesPaginatorOptions
	client    ListEmailIdentitiesAPIClient
	params    *ListEmailIdentitiesInput
	nextToken *string
	firstPage bool
}

// NewListEmailIdentitiesPaginator returns a new ListEmailIdentitiesPaginator
func NewListEmailIdentitiesPaginator(client ListEmailIdentitiesAPIClient, params *ListEmailIdentitiesInput, optFns ...func(*ListEmailIdentitiesPaginatorOptions)) *ListEmailIdentitiesPaginator {
	if params == nil {
		params = &ListEmailIdentitiesInput{}
	}

	options := ListEmailIdentitiesPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEmailIdentitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEmailIdentitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEmailIdentities page.
func (p *ListEmailIdentitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEmailIdentitiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListEmailIdentities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEmailIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListEmailIdentities",
	}
}
