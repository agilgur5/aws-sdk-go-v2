// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all retraining schedulers in your account, filtering by model name prefix
// and status.
func (c *Client) ListRetrainingSchedulers(ctx context.Context, params *ListRetrainingSchedulersInput, optFns ...func(*Options)) (*ListRetrainingSchedulersOutput, error) {
	if params == nil {
		params = &ListRetrainingSchedulersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRetrainingSchedulers", params, optFns, c.addOperationListRetrainingSchedulersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRetrainingSchedulersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRetrainingSchedulersInput struct {

	// Specifies the maximum number of retraining schedulers to list.
	MaxResults *int32

	// Specify this field to only list retraining schedulers whose machine learning
	// models begin with the value you specify.
	ModelNameBeginsWith *string

	// If the number of results exceeds the maximum, a pagination token is returned.
	// Use the token in the request to show the next page of retraining schedulers.
	NextToken *string

	// Specify this field to only list retraining schedulers whose status matches the
	// value you specify.
	Status types.RetrainingSchedulerStatus

	noSmithyDocumentSerde
}

type ListRetrainingSchedulersOutput struct {

	// If the number of results exceeds the maximum, this pagination token is
	// returned. Use this token in the request to show the next page of retraining
	// schedulers.
	NextToken *string

	// Provides information on the specified retraining scheduler, including the model
	// name, model ARN, status, and start date.
	RetrainingSchedulerSummaries []types.RetrainingSchedulerSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRetrainingSchedulersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRetrainingSchedulers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRetrainingSchedulers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRetrainingSchedulers(options.Region), middleware.Before); err != nil {
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
		operation: "ListRetrainingSchedulers",
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

// ListRetrainingSchedulersAPIClient is a client that implements the
// ListRetrainingSchedulers operation.
type ListRetrainingSchedulersAPIClient interface {
	ListRetrainingSchedulers(context.Context, *ListRetrainingSchedulersInput, ...func(*Options)) (*ListRetrainingSchedulersOutput, error)
}

var _ ListRetrainingSchedulersAPIClient = (*Client)(nil)

// ListRetrainingSchedulersPaginatorOptions is the paginator options for
// ListRetrainingSchedulers
type ListRetrainingSchedulersPaginatorOptions struct {
	// Specifies the maximum number of retraining schedulers to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRetrainingSchedulersPaginator is a paginator for ListRetrainingSchedulers
type ListRetrainingSchedulersPaginator struct {
	options   ListRetrainingSchedulersPaginatorOptions
	client    ListRetrainingSchedulersAPIClient
	params    *ListRetrainingSchedulersInput
	nextToken *string
	firstPage bool
}

// NewListRetrainingSchedulersPaginator returns a new
// ListRetrainingSchedulersPaginator
func NewListRetrainingSchedulersPaginator(client ListRetrainingSchedulersAPIClient, params *ListRetrainingSchedulersInput, optFns ...func(*ListRetrainingSchedulersPaginatorOptions)) *ListRetrainingSchedulersPaginator {
	if params == nil {
		params = &ListRetrainingSchedulersInput{}
	}

	options := ListRetrainingSchedulersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRetrainingSchedulersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRetrainingSchedulersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRetrainingSchedulers page.
func (p *ListRetrainingSchedulersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRetrainingSchedulersOutput, error) {
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

	result, err := p.client.ListRetrainingSchedulers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRetrainingSchedulers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "ListRetrainingSchedulers",
	}
}
