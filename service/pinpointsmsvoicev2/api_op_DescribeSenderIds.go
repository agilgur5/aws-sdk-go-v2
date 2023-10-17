// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified SenderIds or all SenderIds associated with your Amazon
// Web Services account. If you specify SenderIds, the output includes information
// for only the specified SenderIds. If you specify filters, the output includes
// information for only those SenderIds that meet the filter criteria. If you don't
// specify SenderIds or filters, the output includes information for all SenderIds.
// f you specify a sender ID that isn't valid, an Error is returned.
func (c *Client) DescribeSenderIds(ctx context.Context, params *DescribeSenderIdsInput, optFns ...func(*Options)) (*DescribeSenderIdsOutput, error) {
	if params == nil {
		params = &DescribeSenderIdsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSenderIds", params, optFns, c.addOperationDescribeSenderIdsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSenderIdsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSenderIdsInput struct {

	// An array of SenderIdFilter objects to filter the results.
	Filters []types.SenderIdFilter

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// An array of SenderIdAndCountry objects to search for.
	SenderIds []types.SenderIdAndCountry

	noSmithyDocumentSerde
}

type DescribeSenderIdsOutput struct {

	// The token to be used for the next set of paginated results. If this field is
	// empty then there are no more results.
	NextToken *string

	// An array of SernderIdInformation objects that contain the details for the
	// requested SenderIds.
	SenderIds []types.SenderIdInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSenderIdsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeSenderIds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeSenderIds{}, middleware.After)
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
	if err = addOpDescribeSenderIdsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSenderIds(options.Region), middleware.Before); err != nil {
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

// DescribeSenderIdsAPIClient is a client that implements the DescribeSenderIds
// operation.
type DescribeSenderIdsAPIClient interface {
	DescribeSenderIds(context.Context, *DescribeSenderIdsInput, ...func(*Options)) (*DescribeSenderIdsOutput, error)
}

var _ DescribeSenderIdsAPIClient = (*Client)(nil)

// DescribeSenderIdsPaginatorOptions is the paginator options for DescribeSenderIds
type DescribeSenderIdsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSenderIdsPaginator is a paginator for DescribeSenderIds
type DescribeSenderIdsPaginator struct {
	options   DescribeSenderIdsPaginatorOptions
	client    DescribeSenderIdsAPIClient
	params    *DescribeSenderIdsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSenderIdsPaginator returns a new DescribeSenderIdsPaginator
func NewDescribeSenderIdsPaginator(client DescribeSenderIdsAPIClient, params *DescribeSenderIdsInput, optFns ...func(*DescribeSenderIdsPaginatorOptions)) *DescribeSenderIdsPaginator {
	if params == nil {
		params = &DescribeSenderIdsInput{}
	}

	options := DescribeSenderIdsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSenderIdsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSenderIdsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSenderIds page.
func (p *DescribeSenderIdsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSenderIdsOutput, error) {
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

	result, err := p.client.DescribeSenderIds(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSenderIds(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DescribeSenderIds",
	}
}
