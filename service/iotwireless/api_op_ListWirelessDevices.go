// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the wireless devices registered to your AWS account.
func (c *Client) ListWirelessDevices(ctx context.Context, params *ListWirelessDevicesInput, optFns ...func(*Options)) (*ListWirelessDevicesOutput, error) {
	if params == nil {
		params = &ListWirelessDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWirelessDevices", params, optFns, c.addOperationListWirelessDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWirelessDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWirelessDevicesInput struct {

	// A filter to list only the wireless devices that use this destination.
	DestinationName *string

	// A filter to list only the wireless devices that use this device profile.
	DeviceProfileId *string

	// The ID of a FUOTA task.
	FuotaTaskId *string

	// The maximum number of results to return in this operation.
	MaxResults int32

	// The ID of the multicast group.
	MulticastGroupId *string

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// A filter to list only the wireless devices that use this service profile.
	ServiceProfileId *string

	// A filter to list only the wireless devices that use this wireless device type.
	WirelessDeviceType types.WirelessDeviceType

	noSmithyDocumentSerde
}

type ListWirelessDevicesOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The ID of the wireless device.
	WirelessDeviceList []types.WirelessDeviceStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWirelessDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWirelessDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWirelessDevices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWirelessDevices(options.Region), middleware.Before); err != nil {
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
		operation: "ListWirelessDevices",
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

// ListWirelessDevicesAPIClient is a client that implements the
// ListWirelessDevices operation.
type ListWirelessDevicesAPIClient interface {
	ListWirelessDevices(context.Context, *ListWirelessDevicesInput, ...func(*Options)) (*ListWirelessDevicesOutput, error)
}

var _ ListWirelessDevicesAPIClient = (*Client)(nil)

// ListWirelessDevicesPaginatorOptions is the paginator options for
// ListWirelessDevices
type ListWirelessDevicesPaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWirelessDevicesPaginator is a paginator for ListWirelessDevices
type ListWirelessDevicesPaginator struct {
	options   ListWirelessDevicesPaginatorOptions
	client    ListWirelessDevicesAPIClient
	params    *ListWirelessDevicesInput
	nextToken *string
	firstPage bool
}

// NewListWirelessDevicesPaginator returns a new ListWirelessDevicesPaginator
func NewListWirelessDevicesPaginator(client ListWirelessDevicesAPIClient, params *ListWirelessDevicesInput, optFns ...func(*ListWirelessDevicesPaginatorOptions)) *ListWirelessDevicesPaginator {
	if params == nil {
		params = &ListWirelessDevicesInput{}
	}

	options := ListWirelessDevicesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWirelessDevicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWirelessDevicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWirelessDevices page.
func (p *ListWirelessDevicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWirelessDevicesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListWirelessDevices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWirelessDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "ListWirelessDevices",
	}
}
