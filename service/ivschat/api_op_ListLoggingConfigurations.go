// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summary information about all your logging configurations in the AWS
// region where the API request is processed.
func (c *Client) ListLoggingConfigurations(ctx context.Context, params *ListLoggingConfigurationsInput, optFns ...func(*Options)) (*ListLoggingConfigurationsOutput, error) {
	if params == nil {
		params = &ListLoggingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLoggingConfigurations", params, optFns, c.addOperationListLoggingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLoggingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLoggingConfigurationsInput struct {

	// Maximum number of logging configurations to return. Default: 50.
	MaxResults int32

	// The first logging configurations to retrieve. This is used for pagination; see
	// the nextToken response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLoggingConfigurationsOutput struct {

	// List of the matching logging configurations (summary information only). There
	// is only one type of destination ( cloudWatchLogs , firehose , or s3 ) in a
	// destinationConfiguration .
	//
	// This member is required.
	LoggingConfigurations []types.LoggingConfigurationSummary

	// If there are more logging configurations than maxResults , use nextToken in the
	// request to get the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLoggingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLoggingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLoggingConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLoggingConfigurations(options.Region), middleware.Before); err != nil {
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
		operation: "ListLoggingConfigurations",
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

// ListLoggingConfigurationsAPIClient is a client that implements the
// ListLoggingConfigurations operation.
type ListLoggingConfigurationsAPIClient interface {
	ListLoggingConfigurations(context.Context, *ListLoggingConfigurationsInput, ...func(*Options)) (*ListLoggingConfigurationsOutput, error)
}

var _ ListLoggingConfigurationsAPIClient = (*Client)(nil)

// ListLoggingConfigurationsPaginatorOptions is the paginator options for
// ListLoggingConfigurations
type ListLoggingConfigurationsPaginatorOptions struct {
	// Maximum number of logging configurations to return. Default: 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLoggingConfigurationsPaginator is a paginator for ListLoggingConfigurations
type ListLoggingConfigurationsPaginator struct {
	options   ListLoggingConfigurationsPaginatorOptions
	client    ListLoggingConfigurationsAPIClient
	params    *ListLoggingConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListLoggingConfigurationsPaginator returns a new
// ListLoggingConfigurationsPaginator
func NewListLoggingConfigurationsPaginator(client ListLoggingConfigurationsAPIClient, params *ListLoggingConfigurationsInput, optFns ...func(*ListLoggingConfigurationsPaginatorOptions)) *ListLoggingConfigurationsPaginator {
	if params == nil {
		params = &ListLoggingConfigurationsInput{}
	}

	options := ListLoggingConfigurationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLoggingConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLoggingConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLoggingConfigurations page.
func (p *ListLoggingConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLoggingConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListLoggingConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLoggingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivschat",
		OperationName: "ListLoggingConfigurations",
	}
}
