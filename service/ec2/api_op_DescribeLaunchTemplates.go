// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more launch templates.
func (c *Client) DescribeLaunchTemplates(ctx context.Context, params *DescribeLaunchTemplatesInput, optFns ...func(*Options)) (*DescribeLaunchTemplatesOutput, error) {
	if params == nil {
		params = &DescribeLaunchTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLaunchTemplates", params, optFns, c.addOperationDescribeLaunchTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLaunchTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLaunchTemplatesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//   - create-time - The time the launch template was created.
	//   - launch-template-name - The name of the launch template.
	//   - tag : - The key/value combination of a tag assigned to the resource. Use the
	//   tag key in the filter name and the tag value as the filter value. For example,
	//   to find all resources that have a tag with the key Owner and the value TeamA ,
	//   specify tag:Owner for the filter name and TeamA for the filter value.
	//   - tag-key - The key of a tag assigned to the resource. Use this filter to find
	//   all resources assigned a tag with a specific key, regardless of the tag value.
	Filters []types.Filter

	// One or more launch template IDs.
	LaunchTemplateIds []string

	// One or more launch template names.
	LaunchTemplateNames []string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 1 and 200.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeLaunchTemplatesOutput struct {

	// Information about the launch templates.
	LaunchTemplates []types.LaunchTemplate

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLaunchTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeLaunchTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLaunchTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLaunchTemplates(options.Region), middleware.Before); err != nil {
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

// DescribeLaunchTemplatesAPIClient is a client that implements the
// DescribeLaunchTemplates operation.
type DescribeLaunchTemplatesAPIClient interface {
	DescribeLaunchTemplates(context.Context, *DescribeLaunchTemplatesInput, ...func(*Options)) (*DescribeLaunchTemplatesOutput, error)
}

var _ DescribeLaunchTemplatesAPIClient = (*Client)(nil)

// DescribeLaunchTemplatesPaginatorOptions is the paginator options for
// DescribeLaunchTemplates
type DescribeLaunchTemplatesPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 1 and 200.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLaunchTemplatesPaginator is a paginator for DescribeLaunchTemplates
type DescribeLaunchTemplatesPaginator struct {
	options   DescribeLaunchTemplatesPaginatorOptions
	client    DescribeLaunchTemplatesAPIClient
	params    *DescribeLaunchTemplatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeLaunchTemplatesPaginator returns a new
// DescribeLaunchTemplatesPaginator
func NewDescribeLaunchTemplatesPaginator(client DescribeLaunchTemplatesAPIClient, params *DescribeLaunchTemplatesInput, optFns ...func(*DescribeLaunchTemplatesPaginatorOptions)) *DescribeLaunchTemplatesPaginator {
	if params == nil {
		params = &DescribeLaunchTemplatesInput{}
	}

	options := DescribeLaunchTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLaunchTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLaunchTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLaunchTemplates page.
func (p *DescribeLaunchTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLaunchTemplatesOutput, error) {
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

	result, err := p.client.DescribeLaunchTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeLaunchTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLaunchTemplates",
	}
}
