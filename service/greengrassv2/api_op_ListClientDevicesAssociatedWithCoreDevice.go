// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of client devices that are associated with a core
// device.
func (c *Client) ListClientDevicesAssociatedWithCoreDevice(ctx context.Context, params *ListClientDevicesAssociatedWithCoreDeviceInput, optFns ...func(*Options)) (*ListClientDevicesAssociatedWithCoreDeviceOutput, error) {
	if params == nil {
		params = &ListClientDevicesAssociatedWithCoreDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClientDevicesAssociatedWithCoreDevice", params, optFns, c.addOperationListClientDevicesAssociatedWithCoreDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClientDevicesAssociatedWithCoreDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClientDevicesAssociatedWithCoreDeviceInput struct {

	// The name of the core device. This is also the name of the IoT thing.
	//
	// This member is required.
	CoreDeviceThingName *string

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListClientDevicesAssociatedWithCoreDeviceOutput struct {

	// A list that describes the client devices that are associated with the core
	// device.
	AssociatedClientDevices []types.AssociatedClientDevice

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClientDevicesAssociatedWithCoreDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListClientDevicesAssociatedWithCoreDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListClientDevicesAssociatedWithCoreDevice{}, middleware.After)
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
	if err = addOpListClientDevicesAssociatedWithCoreDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClientDevicesAssociatedWithCoreDevice(options.Region), middleware.Before); err != nil {
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
		operation: "ListClientDevicesAssociatedWithCoreDevice",
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

// ListClientDevicesAssociatedWithCoreDeviceAPIClient is a client that implements
// the ListClientDevicesAssociatedWithCoreDevice operation.
type ListClientDevicesAssociatedWithCoreDeviceAPIClient interface {
	ListClientDevicesAssociatedWithCoreDevice(context.Context, *ListClientDevicesAssociatedWithCoreDeviceInput, ...func(*Options)) (*ListClientDevicesAssociatedWithCoreDeviceOutput, error)
}

var _ ListClientDevicesAssociatedWithCoreDeviceAPIClient = (*Client)(nil)

// ListClientDevicesAssociatedWithCoreDevicePaginatorOptions is the paginator
// options for ListClientDevicesAssociatedWithCoreDevice
type ListClientDevicesAssociatedWithCoreDevicePaginatorOptions struct {
	// The maximum number of results to be returned per paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClientDevicesAssociatedWithCoreDevicePaginator is a paginator for
// ListClientDevicesAssociatedWithCoreDevice
type ListClientDevicesAssociatedWithCoreDevicePaginator struct {
	options   ListClientDevicesAssociatedWithCoreDevicePaginatorOptions
	client    ListClientDevicesAssociatedWithCoreDeviceAPIClient
	params    *ListClientDevicesAssociatedWithCoreDeviceInput
	nextToken *string
	firstPage bool
}

// NewListClientDevicesAssociatedWithCoreDevicePaginator returns a new
// ListClientDevicesAssociatedWithCoreDevicePaginator
func NewListClientDevicesAssociatedWithCoreDevicePaginator(client ListClientDevicesAssociatedWithCoreDeviceAPIClient, params *ListClientDevicesAssociatedWithCoreDeviceInput, optFns ...func(*ListClientDevicesAssociatedWithCoreDevicePaginatorOptions)) *ListClientDevicesAssociatedWithCoreDevicePaginator {
	if params == nil {
		params = &ListClientDevicesAssociatedWithCoreDeviceInput{}
	}

	options := ListClientDevicesAssociatedWithCoreDevicePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClientDevicesAssociatedWithCoreDevicePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClientDevicesAssociatedWithCoreDevicePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListClientDevicesAssociatedWithCoreDevice page.
func (p *ListClientDevicesAssociatedWithCoreDevicePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClientDevicesAssociatedWithCoreDeviceOutput, error) {
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

	result, err := p.client.ListClientDevicesAssociatedWithCoreDevice(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListClientDevicesAssociatedWithCoreDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListClientDevicesAssociatedWithCoreDevice",
	}
}
