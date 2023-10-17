// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the assessment runs that correspond to the assessment templates that are
// specified by the ARNs of the assessment templates.
func (c *Client) ListAssessmentRuns(ctx context.Context, params *ListAssessmentRunsInput, optFns ...func(*Options)) (*ListAssessmentRunsOutput, error) {
	if params == nil {
		params = &ListAssessmentRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessmentRuns", params, optFns, c.addOperationListAssessmentRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentRunsInput struct {

	// The ARNs that specify the assessment templates whose assessment runs you want
	// to list.
	AssessmentTemplateArns []string

	// You can use this parameter to specify a subset of data to be included in the
	// action's response. For a record to match a filter, all specified filter
	// attributes must match. When multiple values are specified for a filter
	// attribute, any of the values can match.
	Filter *types.AssessmentRunFilter

	// You can use this parameter to indicate the maximum number of items that you
	// want in the response. The default value is 10. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListAssessmentRuns action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// NextToken from the previous response to continue listing data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssessmentRunsOutput struct {

	// A list of ARNs that specifies the assessment runs that are returned by the
	// action.
	//
	// This member is required.
	AssessmentRunArns []string

	// When a response is generated, if there is more data to be listed, this
	// parameter is present in the response and contains the value to use for the
	// nextToken parameter in a subsequent pagination request. If there is no more data
	// to be listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssessmentRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssessmentRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssessmentRuns{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessmentRuns(options.Region), middleware.Before); err != nil {
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

// ListAssessmentRunsAPIClient is a client that implements the ListAssessmentRuns
// operation.
type ListAssessmentRunsAPIClient interface {
	ListAssessmentRuns(context.Context, *ListAssessmentRunsInput, ...func(*Options)) (*ListAssessmentRunsOutput, error)
}

var _ ListAssessmentRunsAPIClient = (*Client)(nil)

// ListAssessmentRunsPaginatorOptions is the paginator options for
// ListAssessmentRuns
type ListAssessmentRunsPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items that you
	// want in the response. The default value is 10. The maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssessmentRunsPaginator is a paginator for ListAssessmentRuns
type ListAssessmentRunsPaginator struct {
	options   ListAssessmentRunsPaginatorOptions
	client    ListAssessmentRunsAPIClient
	params    *ListAssessmentRunsInput
	nextToken *string
	firstPage bool
}

// NewListAssessmentRunsPaginator returns a new ListAssessmentRunsPaginator
func NewListAssessmentRunsPaginator(client ListAssessmentRunsAPIClient, params *ListAssessmentRunsInput, optFns ...func(*ListAssessmentRunsPaginatorOptions)) *ListAssessmentRunsPaginator {
	if params == nil {
		params = &ListAssessmentRunsInput{}
	}

	options := ListAssessmentRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssessmentRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssessmentRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssessmentRuns page.
func (p *ListAssessmentRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssessmentRunsOutput, error) {
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

	result, err := p.client.ListAssessmentRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssessmentRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "ListAssessmentRuns",
	}
}
