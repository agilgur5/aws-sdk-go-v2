// Code generated by smithy-go-codegen DO NOT EDIT.

package snowdevicemanagement

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the Amazon Web Services resources available for a device.
// Currently, Amazon EC2 instances are the only supported resource type.
func (c *Client) ListDeviceResources(ctx context.Context, params *ListDeviceResourcesInput, optFns ...func(*Options)) (*ListDeviceResourcesOutput, error) {
	if params == nil {
		params = &ListDeviceResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeviceResources", params, optFns, c.addOperationListDeviceResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeviceResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeviceResourcesInput struct {

	// The ID of the managed device that you are listing the resources of.
	//
	// This member is required.
	ManagedDeviceId *string

	// The maximum number of resources per page.
	MaxResults *int32

	// A pagination token to continue to the next page of results.
	NextToken *string

	// A structure used to filter the results by type of resource.
	Type *string

	noSmithyDocumentSerde
}

type ListDeviceResourcesOutput struct {

	// A pagination token to continue to the next page of results.
	NextToken *string

	// A structure defining the resource's type, Amazon Resource Name (ARN), and ID.
	Resources []types.ResourceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeviceResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeviceResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeviceResources{}, middleware.After)
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
	if err = addOpListDeviceResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeviceResources(options.Region), middleware.Before); err != nil {
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
		operation: "ListDeviceResources",
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

// ListDeviceResourcesAPIClient is a client that implements the
// ListDeviceResources operation.
type ListDeviceResourcesAPIClient interface {
	ListDeviceResources(context.Context, *ListDeviceResourcesInput, ...func(*Options)) (*ListDeviceResourcesOutput, error)
}

var _ ListDeviceResourcesAPIClient = (*Client)(nil)

// ListDeviceResourcesPaginatorOptions is the paginator options for
// ListDeviceResources
type ListDeviceResourcesPaginatorOptions struct {
	// The maximum number of resources per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeviceResourcesPaginator is a paginator for ListDeviceResources
type ListDeviceResourcesPaginator struct {
	options   ListDeviceResourcesPaginatorOptions
	client    ListDeviceResourcesAPIClient
	params    *ListDeviceResourcesInput
	nextToken *string
	firstPage bool
}

// NewListDeviceResourcesPaginator returns a new ListDeviceResourcesPaginator
func NewListDeviceResourcesPaginator(client ListDeviceResourcesAPIClient, params *ListDeviceResourcesInput, optFns ...func(*ListDeviceResourcesPaginatorOptions)) *ListDeviceResourcesPaginator {
	if params == nil {
		params = &ListDeviceResourcesInput{}
	}

	options := ListDeviceResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeviceResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeviceResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeviceResources page.
func (p *ListDeviceResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeviceResourcesOutput, error) {
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

	result, err := p.client.ListDeviceResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeviceResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snow-device-management",
		OperationName: "ListDeviceResources",
	}
}
