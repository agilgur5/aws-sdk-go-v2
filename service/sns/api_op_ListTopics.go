// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the requester's topics. Each call returns a limited list of
// topics, up to 100. If there are more topics, a NextToken is also returned. Use
// the NextToken parameter in a new ListTopics call to get further results. This
// action is throttled at 30 transactions per second (TPS).
func (c *Client) ListTopics(ctx context.Context, params *ListTopicsInput, optFns ...func(*Options)) (*ListTopicsOutput, error) {
	if params == nil {
		params = &ListTopicsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTopics", params, optFns, c.addOperationListTopicsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTopicsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTopicsInput struct {

	// Token returned by the previous ListTopics request.
	NextToken *string

	noSmithyDocumentSerde
}

// Response for ListTopics action.
type ListTopicsOutput struct {

	// Token to pass along to the next ListTopics request. This element is returned if
	// there are additional topics to retrieve.
	NextToken *string

	// A list of topic ARNs.
	Topics []types.Topic

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTopicsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListTopics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListTopics{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTopics(options.Region), middleware.Before); err != nil {
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

// ListTopicsAPIClient is a client that implements the ListTopics operation.
type ListTopicsAPIClient interface {
	ListTopics(context.Context, *ListTopicsInput, ...func(*Options)) (*ListTopicsOutput, error)
}

var _ ListTopicsAPIClient = (*Client)(nil)

// ListTopicsPaginatorOptions is the paginator options for ListTopics
type ListTopicsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTopicsPaginator is a paginator for ListTopics
type ListTopicsPaginator struct {
	options   ListTopicsPaginatorOptions
	client    ListTopicsAPIClient
	params    *ListTopicsInput
	nextToken *string
	firstPage bool
}

// NewListTopicsPaginator returns a new ListTopicsPaginator
func NewListTopicsPaginator(client ListTopicsAPIClient, params *ListTopicsInput, optFns ...func(*ListTopicsPaginatorOptions)) *ListTopicsPaginator {
	if params == nil {
		params = &ListTopicsInput{}
	}

	options := ListTopicsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTopicsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTopicsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTopics page.
func (p *ListTopicsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTopicsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListTopics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTopics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListTopics",
	}
}
