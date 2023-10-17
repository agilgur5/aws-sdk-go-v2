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

// Returns a list of the subscriptions to a specific topic. Each call returns a
// limited list of subscriptions, up to 100. If there are more subscriptions, a
// NextToken is also returned. Use the NextToken parameter in a new
// ListSubscriptionsByTopic call to get further results. This action is throttled
// at 30 transactions per second (TPS).
func (c *Client) ListSubscriptionsByTopic(ctx context.Context, params *ListSubscriptionsByTopicInput, optFns ...func(*Options)) (*ListSubscriptionsByTopicOutput, error) {
	if params == nil {
		params = &ListSubscriptionsByTopicInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubscriptionsByTopic", params, optFns, c.addOperationListSubscriptionsByTopicMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubscriptionsByTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListSubscriptionsByTopic action.
type ListSubscriptionsByTopicInput struct {

	// The ARN of the topic for which you wish to find subscriptions.
	//
	// This member is required.
	TopicArn *string

	// Token returned by the previous ListSubscriptionsByTopic request.
	NextToken *string

	noSmithyDocumentSerde
}

// Response for ListSubscriptionsByTopic action.
type ListSubscriptionsByTopicOutput struct {

	// Token to pass along to the next ListSubscriptionsByTopic request. This element
	// is returned if there are more subscriptions to retrieve.
	NextToken *string

	// A list of subscriptions.
	Subscriptions []types.Subscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubscriptionsByTopicMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListSubscriptionsByTopic{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListSubscriptionsByTopic{}, middleware.After)
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
	if err = addOpListSubscriptionsByTopicValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptionsByTopic(options.Region), middleware.Before); err != nil {
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

// ListSubscriptionsByTopicAPIClient is a client that implements the
// ListSubscriptionsByTopic operation.
type ListSubscriptionsByTopicAPIClient interface {
	ListSubscriptionsByTopic(context.Context, *ListSubscriptionsByTopicInput, ...func(*Options)) (*ListSubscriptionsByTopicOutput, error)
}

var _ ListSubscriptionsByTopicAPIClient = (*Client)(nil)

// ListSubscriptionsByTopicPaginatorOptions is the paginator options for
// ListSubscriptionsByTopic
type ListSubscriptionsByTopicPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubscriptionsByTopicPaginator is a paginator for ListSubscriptionsByTopic
type ListSubscriptionsByTopicPaginator struct {
	options   ListSubscriptionsByTopicPaginatorOptions
	client    ListSubscriptionsByTopicAPIClient
	params    *ListSubscriptionsByTopicInput
	nextToken *string
	firstPage bool
}

// NewListSubscriptionsByTopicPaginator returns a new
// ListSubscriptionsByTopicPaginator
func NewListSubscriptionsByTopicPaginator(client ListSubscriptionsByTopicAPIClient, params *ListSubscriptionsByTopicInput, optFns ...func(*ListSubscriptionsByTopicPaginatorOptions)) *ListSubscriptionsByTopicPaginator {
	if params == nil {
		params = &ListSubscriptionsByTopicInput{}
	}

	options := ListSubscriptionsByTopicPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubscriptionsByTopicPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubscriptionsByTopicPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSubscriptionsByTopic page.
func (p *ListSubscriptionsByTopicPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubscriptionsByTopicOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListSubscriptionsByTopic(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSubscriptionsByTopic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListSubscriptionsByTopic",
	}
}
