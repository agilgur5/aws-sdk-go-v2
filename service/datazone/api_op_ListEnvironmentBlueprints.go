// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists blueprints in an Amazon DataZone environment.
func (c *Client) ListEnvironmentBlueprints(ctx context.Context, params *ListEnvironmentBlueprintsInput, optFns ...func(*Options)) (*ListEnvironmentBlueprintsOutput, error) {
	if params == nil {
		params = &ListEnvironmentBlueprintsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironmentBlueprints", params, optFns, c.addOperationListEnvironmentBlueprintsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentBlueprintsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentBlueprintsInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// Specifies whether the environment blueprint is managed by Amazon DataZone.
	Managed *bool

	// The maximum number of blueprints to return in a single call to
	// ListEnvironmentBlueprints . When the number of blueprints to be listed is
	// greater than the value of MaxResults , the response contains a NextToken value
	// that you can use in a subsequent call to ListEnvironmentBlueprints to list the
	// next set of blueprints.
	MaxResults *int32

	// The name of the Amazon DataZone environment.
	Name *string

	// When the number of blueprints in the environment is greater than the default
	// value for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of blueprints in the environment, the
	// response includes a pagination token named NextToken . You can specify this
	// NextToken value in a subsequent call to ListEnvironmentBlueprints to list the
	// next set of blueprints.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnvironmentBlueprintsOutput struct {

	// The results of the ListEnvironmentBlueprints action.
	//
	// This member is required.
	Items []types.EnvironmentBlueprintSummary

	// When the number of blueprints in the environment is greater than the default
	// value for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of blueprints in the environment, the
	// response includes a pagination token named NextToken . You can specify this
	// NextToken value in a subsequent call to ListEnvironmentBlueprints to list the
	// next set of blueprints.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentBlueprintsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnvironmentBlueprints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnvironmentBlueprints{}, middleware.After)
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
	if err = addOpListEnvironmentBlueprintsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironmentBlueprints(options.Region), middleware.Before); err != nil {
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

// ListEnvironmentBlueprintsAPIClient is a client that implements the
// ListEnvironmentBlueprints operation.
type ListEnvironmentBlueprintsAPIClient interface {
	ListEnvironmentBlueprints(context.Context, *ListEnvironmentBlueprintsInput, ...func(*Options)) (*ListEnvironmentBlueprintsOutput, error)
}

var _ ListEnvironmentBlueprintsAPIClient = (*Client)(nil)

// ListEnvironmentBlueprintsPaginatorOptions is the paginator options for
// ListEnvironmentBlueprints
type ListEnvironmentBlueprintsPaginatorOptions struct {
	// The maximum number of blueprints to return in a single call to
	// ListEnvironmentBlueprints . When the number of blueprints to be listed is
	// greater than the value of MaxResults , the response contains a NextToken value
	// that you can use in a subsequent call to ListEnvironmentBlueprints to list the
	// next set of blueprints.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentBlueprintsPaginator is a paginator for ListEnvironmentBlueprints
type ListEnvironmentBlueprintsPaginator struct {
	options   ListEnvironmentBlueprintsPaginatorOptions
	client    ListEnvironmentBlueprintsAPIClient
	params    *ListEnvironmentBlueprintsInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentBlueprintsPaginator returns a new
// ListEnvironmentBlueprintsPaginator
func NewListEnvironmentBlueprintsPaginator(client ListEnvironmentBlueprintsAPIClient, params *ListEnvironmentBlueprintsInput, optFns ...func(*ListEnvironmentBlueprintsPaginatorOptions)) *ListEnvironmentBlueprintsPaginator {
	if params == nil {
		params = &ListEnvironmentBlueprintsInput{}
	}

	options := ListEnvironmentBlueprintsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentBlueprintsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentBlueprintsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironmentBlueprints page.
func (p *ListEnvironmentBlueprintsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentBlueprintsOutput, error) {
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

	result, err := p.client.ListEnvironmentBlueprints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnvironmentBlueprints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "ListEnvironmentBlueprints",
	}
}
