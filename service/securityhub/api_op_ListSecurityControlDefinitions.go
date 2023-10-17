// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the security controls that apply to a specified standard.
func (c *Client) ListSecurityControlDefinitions(ctx context.Context, params *ListSecurityControlDefinitionsInput, optFns ...func(*Options)) (*ListSecurityControlDefinitionsOutput, error) {
	if params == nil {
		params = &ListSecurityControlDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityControlDefinitions", params, optFns, c.addOperationListSecurityControlDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityControlDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityControlDefinitionsInput struct {

	// An optional parameter that limits the total results of the API response to the
	// specified number. If this parameter isn't provided in the request, the results
	// include the first 25 security controls that apply to the specified standard. The
	// results also include a NextToken parameter that you can use in a subsequent API
	// call to get the next 25 controls. This repeats until all controls for the
	// standard are returned.
	MaxResults int32

	// Optional pagination parameter.
	NextToken *string

	// The Amazon Resource Name (ARN) of the standard that you want to view controls
	// for.
	StandardsArn *string

	noSmithyDocumentSerde
}

type ListSecurityControlDefinitionsOutput struct {

	// An array of controls that apply to the specified standard.
	//
	// This member is required.
	SecurityControlDefinitions []types.SecurityControlDefinition

	// A pagination parameter that's included in the response only if it was included
	// in the request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecurityControlDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityControlDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityControlDefinitions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityControlDefinitions(options.Region), middleware.Before); err != nil {
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

// ListSecurityControlDefinitionsAPIClient is a client that implements the
// ListSecurityControlDefinitions operation.
type ListSecurityControlDefinitionsAPIClient interface {
	ListSecurityControlDefinitions(context.Context, *ListSecurityControlDefinitionsInput, ...func(*Options)) (*ListSecurityControlDefinitionsOutput, error)
}

var _ ListSecurityControlDefinitionsAPIClient = (*Client)(nil)

// ListSecurityControlDefinitionsPaginatorOptions is the paginator options for
// ListSecurityControlDefinitions
type ListSecurityControlDefinitionsPaginatorOptions struct {
	// An optional parameter that limits the total results of the API response to the
	// specified number. If this parameter isn't provided in the request, the results
	// include the first 25 security controls that apply to the specified standard. The
	// results also include a NextToken parameter that you can use in a subsequent API
	// call to get the next 25 controls. This repeats until all controls for the
	// standard are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecurityControlDefinitionsPaginator is a paginator for
// ListSecurityControlDefinitions
type ListSecurityControlDefinitionsPaginator struct {
	options   ListSecurityControlDefinitionsPaginatorOptions
	client    ListSecurityControlDefinitionsAPIClient
	params    *ListSecurityControlDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListSecurityControlDefinitionsPaginator returns a new
// ListSecurityControlDefinitionsPaginator
func NewListSecurityControlDefinitionsPaginator(client ListSecurityControlDefinitionsAPIClient, params *ListSecurityControlDefinitionsInput, optFns ...func(*ListSecurityControlDefinitionsPaginatorOptions)) *ListSecurityControlDefinitionsPaginator {
	if params == nil {
		params = &ListSecurityControlDefinitionsInput{}
	}

	options := ListSecurityControlDefinitionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecurityControlDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecurityControlDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecurityControlDefinitions page.
func (p *ListSecurityControlDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecurityControlDefinitionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListSecurityControlDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSecurityControlDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "ListSecurityControlDefinitions",
	}
}
