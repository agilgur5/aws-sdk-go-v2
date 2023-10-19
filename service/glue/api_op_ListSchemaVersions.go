// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of schema versions that you have created, with minimal
// information. Schema versions in Deleted status will not be included in the
// results. Empty results will be returned if there are no schema versions
// available.
func (c *Client) ListSchemaVersions(ctx context.Context, params *ListSchemaVersionsInput, optFns ...func(*Options)) (*ListSchemaVersionsOutput, error) {
	if params == nil {
		params = &ListSchemaVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSchemaVersions", params, optFns, c.addOperationListSchemaVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSchemaVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchemaVersionsInput struct {

	// This is a wrapper structure to contain schema identity fields. The structure
	// contains:
	//   - SchemaId$SchemaArn: The Amazon Resource Name (ARN) of the schema. Either
	//   SchemaArn or SchemaName and RegistryName has to be provided.
	//   - SchemaId$SchemaName: The name of the schema. Either SchemaArn or SchemaName
	//   and RegistryName has to be provided.
	//
	// This member is required.
	SchemaId *types.SchemaId

	// Maximum number of results required per page. If the value is not supplied, this
	// will be defaulted to 25 per page.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSchemaVersionsOutput struct {

	// A continuation token for paginating the returned list of tokens, returned if
	// the current segment of the list is not the last.
	NextToken *string

	// An array of SchemaVersionList objects containing details of each schema version.
	Schemas []types.SchemaVersionListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSchemaVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSchemaVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSchemaVersions{}, middleware.After)
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
	if err = addOpListSchemaVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSchemaVersions(options.Region), middleware.Before); err != nil {
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
		operation: "ListSchemaVersions",
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

// ListSchemaVersionsAPIClient is a client that implements the ListSchemaVersions
// operation.
type ListSchemaVersionsAPIClient interface {
	ListSchemaVersions(context.Context, *ListSchemaVersionsInput, ...func(*Options)) (*ListSchemaVersionsOutput, error)
}

var _ ListSchemaVersionsAPIClient = (*Client)(nil)

// ListSchemaVersionsPaginatorOptions is the paginator options for
// ListSchemaVersions
type ListSchemaVersionsPaginatorOptions struct {
	// Maximum number of results required per page. If the value is not supplied, this
	// will be defaulted to 25 per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSchemaVersionsPaginator is a paginator for ListSchemaVersions
type ListSchemaVersionsPaginator struct {
	options   ListSchemaVersionsPaginatorOptions
	client    ListSchemaVersionsAPIClient
	params    *ListSchemaVersionsInput
	nextToken *string
	firstPage bool
}

// NewListSchemaVersionsPaginator returns a new ListSchemaVersionsPaginator
func NewListSchemaVersionsPaginator(client ListSchemaVersionsAPIClient, params *ListSchemaVersionsInput, optFns ...func(*ListSchemaVersionsPaginatorOptions)) *ListSchemaVersionsPaginator {
	if params == nil {
		params = &ListSchemaVersionsInput{}
	}

	options := ListSchemaVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSchemaVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSchemaVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSchemaVersions page.
func (p *ListSchemaVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSchemaVersionsOutput, error) {
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

	result, err := p.client.ListSchemaVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSchemaVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListSchemaVersions",
	}
}
