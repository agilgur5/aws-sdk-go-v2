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

// Describes the specified Amazon Web Services Verified Access instances.
func (c *Client) DescribeVerifiedAccessInstanceLoggingConfigurations(ctx context.Context, params *DescribeVerifiedAccessInstanceLoggingConfigurationsInput, optFns ...func(*Options)) (*DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeVerifiedAccessInstanceLoggingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVerifiedAccessInstanceLoggingConfigurations", params, optFns, c.addOperationDescribeVerifiedAccessInstanceLoggingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVerifiedAccessInstanceLoggingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVerifiedAccessInstanceLoggingConfigurationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. Filter names and values are case-sensitive.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the Verified Access instances.
	VerifiedAccessInstanceIds []string

	noSmithyDocumentSerde
}

type DescribeVerifiedAccessInstanceLoggingConfigurationsOutput struct {

	// The current logging configuration for the Verified Access instances.
	LoggingConfigurations []types.VerifiedAccessInstanceLoggingConfiguration

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVerifiedAccessInstanceLoggingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVerifiedAccessInstanceLoggingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVerifiedAccessInstanceLoggingConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVerifiedAccessInstanceLoggingConfigurations(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeVerifiedAccessInstanceLoggingConfigurations",
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

// DescribeVerifiedAccessInstanceLoggingConfigurationsAPIClient is a client that
// implements the DescribeVerifiedAccessInstanceLoggingConfigurations operation.
type DescribeVerifiedAccessInstanceLoggingConfigurationsAPIClient interface {
	DescribeVerifiedAccessInstanceLoggingConfigurations(context.Context, *DescribeVerifiedAccessInstanceLoggingConfigurationsInput, ...func(*Options)) (*DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error)
}

var _ DescribeVerifiedAccessInstanceLoggingConfigurationsAPIClient = (*Client)(nil)

// DescribeVerifiedAccessInstanceLoggingConfigurationsPaginatorOptions is the
// paginator options for DescribeVerifiedAccessInstanceLoggingConfigurations
type DescribeVerifiedAccessInstanceLoggingConfigurationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVerifiedAccessInstanceLoggingConfigurationsPaginator is a paginator for
// DescribeVerifiedAccessInstanceLoggingConfigurations
type DescribeVerifiedAccessInstanceLoggingConfigurationsPaginator struct {
	options   DescribeVerifiedAccessInstanceLoggingConfigurationsPaginatorOptions
	client    DescribeVerifiedAccessInstanceLoggingConfigurationsAPIClient
	params    *DescribeVerifiedAccessInstanceLoggingConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVerifiedAccessInstanceLoggingConfigurationsPaginator returns a new
// DescribeVerifiedAccessInstanceLoggingConfigurationsPaginator
func NewDescribeVerifiedAccessInstanceLoggingConfigurationsPaginator(client DescribeVerifiedAccessInstanceLoggingConfigurationsAPIClient, params *DescribeVerifiedAccessInstanceLoggingConfigurationsInput, optFns ...func(*DescribeVerifiedAccessInstanceLoggingConfigurationsPaginatorOptions)) *DescribeVerifiedAccessInstanceLoggingConfigurationsPaginator {
	if params == nil {
		params = &DescribeVerifiedAccessInstanceLoggingConfigurationsInput{}
	}

	options := DescribeVerifiedAccessInstanceLoggingConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVerifiedAccessInstanceLoggingConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVerifiedAccessInstanceLoggingConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVerifiedAccessInstanceLoggingConfigurations
// page.
func (p *DescribeVerifiedAccessInstanceLoggingConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
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

	result, err := p.client.DescribeVerifiedAccessInstanceLoggingConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeVerifiedAccessInstanceLoggingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVerifiedAccessInstanceLoggingConfigurations",
	}
}
