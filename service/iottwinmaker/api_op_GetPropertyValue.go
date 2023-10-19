// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the property values for a component, component type, entity, or workspace.
// You must specify a value for either componentName , componentTypeId , entityId ,
// or workspaceId .
func (c *Client) GetPropertyValue(ctx context.Context, params *GetPropertyValueInput, optFns ...func(*Options)) (*GetPropertyValueOutput, error) {
	if params == nil {
		params = &GetPropertyValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPropertyValue", params, optFns, c.addOperationGetPropertyValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPropertyValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPropertyValueInput struct {

	// The properties whose values the operation returns.
	//
	// This member is required.
	SelectedProperties []string

	// The ID of the workspace whose values the operation returns.
	//
	// This member is required.
	WorkspaceId *string

	// The name of the component whose property values the operation returns.
	ComponentName *string

	// The ID of the component type whose property values the operation returns.
	ComponentTypeId *string

	// The ID of the entity whose property values the operation returns.
	EntityId *string

	// The maximum number of results to return at one time. The default is 25. Valid
	// Range: Minimum value of 1. Maximum value of 250.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	// The property group name.
	PropertyGroupName *string

	// The tabular conditions.
	TabularConditions *types.TabularConditions

	noSmithyDocumentSerde
}

type GetPropertyValueOutput struct {

	// The string that specifies the next page of results.
	NextToken *string

	// An object that maps strings to the properties and latest property values in the
	// response. Each string in the mapping must be unique to this object.
	PropertyValues map[string]types.PropertyLatestValue

	// A table of property values.
	TabularPropertyValues [][]map[string]types.DataValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPropertyValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPropertyValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPropertyValue{}, middleware.After)
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
	if err = addEndpointPrefix_opGetPropertyValueMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetPropertyValueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPropertyValue(options.Region), middleware.Before); err != nil {
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
		operation: "GetPropertyValue",
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

type endpointPrefix_opGetPropertyValueMiddleware struct {
}

func (*endpointPrefix_opGetPropertyValueMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetPropertyValueMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetPropertyValueMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetPropertyValueMiddleware{}, `OperationSerializer`, middleware.After)
}

// GetPropertyValueAPIClient is a client that implements the GetPropertyValue
// operation.
type GetPropertyValueAPIClient interface {
	GetPropertyValue(context.Context, *GetPropertyValueInput, ...func(*Options)) (*GetPropertyValueOutput, error)
}

var _ GetPropertyValueAPIClient = (*Client)(nil)

// GetPropertyValuePaginatorOptions is the paginator options for GetPropertyValue
type GetPropertyValuePaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25. Valid
	// Range: Minimum value of 1. Maximum value of 250.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetPropertyValuePaginator is a paginator for GetPropertyValue
type GetPropertyValuePaginator struct {
	options   GetPropertyValuePaginatorOptions
	client    GetPropertyValueAPIClient
	params    *GetPropertyValueInput
	nextToken *string
	firstPage bool
}

// NewGetPropertyValuePaginator returns a new GetPropertyValuePaginator
func NewGetPropertyValuePaginator(client GetPropertyValueAPIClient, params *GetPropertyValueInput, optFns ...func(*GetPropertyValuePaginatorOptions)) *GetPropertyValuePaginator {
	if params == nil {
		params = &GetPropertyValueInput{}
	}

	options := GetPropertyValuePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetPropertyValuePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetPropertyValuePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetPropertyValue page.
func (p *GetPropertyValuePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetPropertyValueOutput, error) {
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

	result, err := p.client.GetPropertyValue(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetPropertyValue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "GetPropertyValue",
	}
}
