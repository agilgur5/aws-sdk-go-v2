// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve information about one or more parameters in a specific hierarchy.
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified. The
// number of items returned, however, can be between zero and the value of
// MaxResults . If the service reaches an internal limit while processing the
// results, it stops the operation and returns the matching values up to that point
// and a NextToken . You can specify the NextToken in a subsequent call to get the
// next set of results.
func (c *Client) GetParametersByPath(ctx context.Context, params *GetParametersByPathInput, optFns ...func(*Options)) (*GetParametersByPathOutput, error) {
	if params == nil {
		params = &GetParametersByPathInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetParametersByPath", params, optFns, c.addOperationGetParametersByPathMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetParametersByPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetParametersByPathInput struct {

	// The hierarchy for the parameter. Hierarchies start with a forward slash (/).
	// The hierarchy is the parameter name except the last part of the parameter. For
	// the API call to succeed, the last part of the parameter name can't be in the
	// path. A parameter name hierarchy can have a maximum of 15 levels. Here is an
	// example of a hierarchy: /Finance/Prod/IAD/WinServ2016/license33
	//
	// This member is required.
	Path *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// Filters to limit the request results. The following Key values are supported
	// for GetParametersByPath : Type , KeyId , and Label . The following Key values
	// aren't supported for GetParametersByPath : tag , DataType , Name , Path , and
	// Tier .
	ParameterFilters []types.ParameterStringFilter

	// Retrieve all parameters within a hierarchy. If a user has access to a path,
	// then the user can access all levels of that path. For example, if a user has
	// permission to access path /a , then the user can also access /a/b . Even if a
	// user has explicitly been denied access in IAM for parameter /a/b , they can
	// still call the GetParametersByPath API operation recursively for /a and view
	// /a/b .
	Recursive *bool

	// Retrieve all parameters in a hierarchy with their value decrypted.
	WithDecryption *bool

	noSmithyDocumentSerde
}

type GetParametersByPathOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A list of parameters found in the specified hierarchy.
	Parameters []types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetParametersByPathMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetParametersByPath{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetParametersByPath{}, middleware.After)
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
	if err = addOpGetParametersByPathValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetParametersByPath(options.Region), middleware.Before); err != nil {
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

// GetParametersByPathAPIClient is a client that implements the
// GetParametersByPath operation.
type GetParametersByPathAPIClient interface {
	GetParametersByPath(context.Context, *GetParametersByPathInput, ...func(*Options)) (*GetParametersByPathOutput, error)
}

var _ GetParametersByPathAPIClient = (*Client)(nil)

// GetParametersByPathPaginatorOptions is the paginator options for
// GetParametersByPath
type GetParametersByPathPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetParametersByPathPaginator is a paginator for GetParametersByPath
type GetParametersByPathPaginator struct {
	options   GetParametersByPathPaginatorOptions
	client    GetParametersByPathAPIClient
	params    *GetParametersByPathInput
	nextToken *string
	firstPage bool
}

// NewGetParametersByPathPaginator returns a new GetParametersByPathPaginator
func NewGetParametersByPathPaginator(client GetParametersByPathAPIClient, params *GetParametersByPathInput, optFns ...func(*GetParametersByPathPaginatorOptions)) *GetParametersByPathPaginator {
	if params == nil {
		params = &GetParametersByPathInput{}
	}

	options := GetParametersByPathPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetParametersByPathPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetParametersByPathPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetParametersByPath page.
func (p *GetParametersByPathPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetParametersByPathOutput, error) {
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

	result, err := p.client.GetParametersByPath(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetParametersByPath(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetParametersByPath",
	}
}
