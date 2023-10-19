// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the recommendation templates for the Resilience Hub applications.
func (c *Client) ListRecommendationTemplates(ctx context.Context, params *ListRecommendationTemplatesInput, optFns ...func(*Options)) (*ListRecommendationTemplatesOutput, error) {
	if params == nil {
		params = &ListRecommendationTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecommendationTemplates", params, optFns, c.addOperationListRecommendationTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecommendationTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecommendationTemplatesInput struct {

	// Amazon Resource Name (ARN) of the assessment. The format for this ARN is: arn:
	// partition :resiliencehub: region : account :app-assessment/ app-id . For more
	// information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	AssessmentArn *string

	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	MaxResults *int32

	// The name for one of the listed recommendation templates.
	Name *string

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	// The Amazon Resource Name (ARN) for a recommendation template.
	RecommendationTemplateArn *string

	// The default is to sort by ascending startTime. To sort by descending startTime,
	// set reverseOrder to true .
	ReverseOrder *bool

	// Status of the action.
	Status []types.RecommendationTemplateStatus

	noSmithyDocumentSerde
}

type ListRecommendationTemplatesOutput struct {

	// Token for the next set of results, or null if there are no more results.
	NextToken *string

	// The recommendation templates for the Resilience Hub applications.
	RecommendationTemplates []types.RecommendationTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecommendationTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecommendationTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecommendationTemplates{}, middleware.After)
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
	if err = addOpListRecommendationTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecommendationTemplates(options.Region), middleware.Before); err != nil {
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
		operation: "ListRecommendationTemplates",
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

// ListRecommendationTemplatesAPIClient is a client that implements the
// ListRecommendationTemplates operation.
type ListRecommendationTemplatesAPIClient interface {
	ListRecommendationTemplates(context.Context, *ListRecommendationTemplatesInput, ...func(*Options)) (*ListRecommendationTemplatesOutput, error)
}

var _ ListRecommendationTemplatesAPIClient = (*Client)(nil)

// ListRecommendationTemplatesPaginatorOptions is the paginator options for
// ListRecommendationTemplates
type ListRecommendationTemplatesPaginatorOptions struct {
	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecommendationTemplatesPaginator is a paginator for
// ListRecommendationTemplates
type ListRecommendationTemplatesPaginator struct {
	options   ListRecommendationTemplatesPaginatorOptions
	client    ListRecommendationTemplatesAPIClient
	params    *ListRecommendationTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListRecommendationTemplatesPaginator returns a new
// ListRecommendationTemplatesPaginator
func NewListRecommendationTemplatesPaginator(client ListRecommendationTemplatesAPIClient, params *ListRecommendationTemplatesInput, optFns ...func(*ListRecommendationTemplatesPaginatorOptions)) *ListRecommendationTemplatesPaginator {
	if params == nil {
		params = &ListRecommendationTemplatesInput{}
	}

	options := ListRecommendationTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecommendationTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecommendationTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecommendationTemplates page.
func (p *ListRecommendationTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecommendationTemplatesOutput, error) {
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

	result, err := p.client.ListRecommendationTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecommendationTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "ListRecommendationTemplates",
	}
}
