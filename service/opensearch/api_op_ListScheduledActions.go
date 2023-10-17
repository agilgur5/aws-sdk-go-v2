// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of configuration changes that are scheduled for a domain.
// These changes can be service software updates (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html)
// or blue/green Auto-Tune enhancements (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html#auto-tune-types)
// .
func (c *Client) ListScheduledActions(ctx context.Context, params *ListScheduledActionsInput, optFns ...func(*Options)) (*ListScheduledActionsOutput, error) {
	if params == nil {
		params = &ListScheduledActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListScheduledActions", params, optFns, c.addOperationListScheduledActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListScheduledActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListScheduledActionsInput struct {

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial ListScheduledActions operation returns a nextToken , you can
	// include the returned nextToken in subsequent ListScheduledActions operations,
	// which returns results in the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListScheduledActionsOutput struct {

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// A list of actions that are scheduled for the domain.
	ScheduledActions []types.ScheduledAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListScheduledActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListScheduledActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListScheduledActions{}, middleware.After)
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
	if err = addOpListScheduledActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListScheduledActions(options.Region), middleware.Before); err != nil {
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

// ListScheduledActionsAPIClient is a client that implements the
// ListScheduledActions operation.
type ListScheduledActionsAPIClient interface {
	ListScheduledActions(context.Context, *ListScheduledActionsInput, ...func(*Options)) (*ListScheduledActionsOutput, error)
}

var _ ListScheduledActionsAPIClient = (*Client)(nil)

// ListScheduledActionsPaginatorOptions is the paginator options for
// ListScheduledActions
type ListScheduledActionsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListScheduledActionsPaginator is a paginator for ListScheduledActions
type ListScheduledActionsPaginator struct {
	options   ListScheduledActionsPaginatorOptions
	client    ListScheduledActionsAPIClient
	params    *ListScheduledActionsInput
	nextToken *string
	firstPage bool
}

// NewListScheduledActionsPaginator returns a new ListScheduledActionsPaginator
func NewListScheduledActionsPaginator(client ListScheduledActionsAPIClient, params *ListScheduledActionsInput, optFns ...func(*ListScheduledActionsPaginatorOptions)) *ListScheduledActionsPaginator {
	if params == nil {
		params = &ListScheduledActionsInput{}
	}

	options := ListScheduledActionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListScheduledActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListScheduledActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListScheduledActions page.
func (p *ListScheduledActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListScheduledActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListScheduledActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListScheduledActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListScheduledActions",
	}
}
