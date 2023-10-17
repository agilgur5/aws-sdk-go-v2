// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of infrastructure configurations.
func (c *Client) ListInfrastructureConfigurations(ctx context.Context, params *ListInfrastructureConfigurationsInput, optFns ...func(*Options)) (*ListInfrastructureConfigurationsOutput, error) {
	if params == nil {
		params = &ListInfrastructureConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInfrastructureConfigurations", params, optFns, c.addOperationListInfrastructureConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInfrastructureConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInfrastructureConfigurationsInput struct {

	// You can filter on name to streamline results.
	Filters []types.Filter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListInfrastructureConfigurationsOutput struct {

	// The list of infrastructure configurations.
	InfrastructureConfigurationSummaryList []types.InfrastructureConfigurationSummary

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service has'ot included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInfrastructureConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInfrastructureConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInfrastructureConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInfrastructureConfigurations(options.Region), middleware.Before); err != nil {
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

// ListInfrastructureConfigurationsAPIClient is a client that implements the
// ListInfrastructureConfigurations operation.
type ListInfrastructureConfigurationsAPIClient interface {
	ListInfrastructureConfigurations(context.Context, *ListInfrastructureConfigurationsInput, ...func(*Options)) (*ListInfrastructureConfigurationsOutput, error)
}

var _ ListInfrastructureConfigurationsAPIClient = (*Client)(nil)

// ListInfrastructureConfigurationsPaginatorOptions is the paginator options for
// ListInfrastructureConfigurations
type ListInfrastructureConfigurationsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInfrastructureConfigurationsPaginator is a paginator for
// ListInfrastructureConfigurations
type ListInfrastructureConfigurationsPaginator struct {
	options   ListInfrastructureConfigurationsPaginatorOptions
	client    ListInfrastructureConfigurationsAPIClient
	params    *ListInfrastructureConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListInfrastructureConfigurationsPaginator returns a new
// ListInfrastructureConfigurationsPaginator
func NewListInfrastructureConfigurationsPaginator(client ListInfrastructureConfigurationsAPIClient, params *ListInfrastructureConfigurationsInput, optFns ...func(*ListInfrastructureConfigurationsPaginatorOptions)) *ListInfrastructureConfigurationsPaginator {
	if params == nil {
		params = &ListInfrastructureConfigurationsInput{}
	}

	options := ListInfrastructureConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInfrastructureConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInfrastructureConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInfrastructureConfigurations page.
func (p *ListInfrastructureConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInfrastructureConfigurationsOutput, error) {
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

	result, err := p.client.ListInfrastructureConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInfrastructureConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListInfrastructureConfigurations",
	}
}
