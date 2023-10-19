// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the platform versions available for your account in an AWS Region.
// Provides summary information about each platform version. Compare to
// DescribePlatformVersion , which provides full details about a single platform
// version. For definitions of platform version and other platform-related terms,
// see AWS Elastic Beanstalk Platforms Glossary (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html)
// .
func (c *Client) ListPlatformVersions(ctx context.Context, params *ListPlatformVersionsInput, optFns ...func(*Options)) (*ListPlatformVersionsOutput, error) {
	if params == nil {
		params = &ListPlatformVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPlatformVersions", params, optFns, c.addOperationListPlatformVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPlatformVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlatformVersionsInput struct {

	// Criteria for restricting the resulting list of platform versions. The filter is
	// interpreted as a logical conjunction (AND) of the separate PlatformFilter terms.
	Filters []types.PlatformFilter

	// The maximum number of platform version values returned in one call.
	MaxRecords *int32

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request. If no NextToken is specified, the
	// first page is retrieved.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPlatformVersionsOutput struct {

	// In a paginated request, if this value isn't null , it's the token that you can
	// pass in a subsequent request to get the next response page.
	NextToken *string

	// Summary information about the platform versions.
	PlatformSummaryList []types.PlatformSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPlatformVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListPlatformVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListPlatformVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPlatformVersions(options.Region), middleware.Before); err != nil {
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
		operation: "ListPlatformVersions",
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

// ListPlatformVersionsAPIClient is a client that implements the
// ListPlatformVersions operation.
type ListPlatformVersionsAPIClient interface {
	ListPlatformVersions(context.Context, *ListPlatformVersionsInput, ...func(*Options)) (*ListPlatformVersionsOutput, error)
}

var _ ListPlatformVersionsAPIClient = (*Client)(nil)

// ListPlatformVersionsPaginatorOptions is the paginator options for
// ListPlatformVersions
type ListPlatformVersionsPaginatorOptions struct {
	// The maximum number of platform version values returned in one call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPlatformVersionsPaginator is a paginator for ListPlatformVersions
type ListPlatformVersionsPaginator struct {
	options   ListPlatformVersionsPaginatorOptions
	client    ListPlatformVersionsAPIClient
	params    *ListPlatformVersionsInput
	nextToken *string
	firstPage bool
}

// NewListPlatformVersionsPaginator returns a new ListPlatformVersionsPaginator
func NewListPlatformVersionsPaginator(client ListPlatformVersionsAPIClient, params *ListPlatformVersionsInput, optFns ...func(*ListPlatformVersionsPaginatorOptions)) *ListPlatformVersionsPaginator {
	if params == nil {
		params = &ListPlatformVersionsInput{}
	}

	options := ListPlatformVersionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPlatformVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPlatformVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPlatformVersions page.
func (p *ListPlatformVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPlatformVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.ListPlatformVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPlatformVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ListPlatformVersions",
	}
}
