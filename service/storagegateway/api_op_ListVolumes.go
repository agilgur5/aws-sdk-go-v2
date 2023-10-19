// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the iSCSI stored volumes of a gateway. Results are sorted by volume ARN.
// The response includes only the volume ARNs. If you want additional volume
// information, use the DescribeStorediSCSIVolumes or the
// DescribeCachediSCSIVolumes API. The operation supports pagination. By default,
// the operation returns a maximum of up to 100 volumes. You can optionally specify
// the Limit field in the body to limit the number of volumes in the response. If
// the number of volumes returned in the response is truncated, the response
// includes a Marker field. You can use this Marker value in your subsequent
// request to retrieve the next set of volumes. This operation is only supported in
// the cached volume and stored volume gateway types.
func (c *Client) ListVolumes(ctx context.Context, params *ListVolumesInput, optFns ...func(*Options)) (*ListVolumesOutput, error) {
	if params == nil {
		params = &ListVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVolumes", params, optFns, c.addOperationListVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object that contains one or more of the following fields:
//   - ListVolumesInput$Limit
//   - ListVolumesInput$Marker
type ListVolumesInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Specifies that the list of volumes returned be limited to the specified number
	// of items.
	Limit *int32

	// A string that indicates the position at which to begin the returned list of
	// volumes. Obtain the marker from the response of a previous List iSCSI Volumes
	// request.
	Marker *string

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
//   - ListVolumesOutput$Marker
//   - ListVolumesOutput$VolumeInfos
type ListVolumesOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Use the marker in your next request to continue pagination of iSCSI volumes. If
	// there are no more volumes to list, this field does not appear in the response
	// body.
	Marker *string

	// An array of VolumeInfo objects, where each object describes an iSCSI volume. If
	// no volumes are defined for the gateway, then VolumeInfos is an empty array "[]".
	VolumeInfos []types.VolumeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVolumes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVolumes(options.Region), middleware.Before); err != nil {
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
		operation: "ListVolumes",
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

// ListVolumesAPIClient is a client that implements the ListVolumes operation.
type ListVolumesAPIClient interface {
	ListVolumes(context.Context, *ListVolumesInput, ...func(*Options)) (*ListVolumesOutput, error)
}

var _ ListVolumesAPIClient = (*Client)(nil)

// ListVolumesPaginatorOptions is the paginator options for ListVolumes
type ListVolumesPaginatorOptions struct {
	// Specifies that the list of volumes returned be limited to the specified number
	// of items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVolumesPaginator is a paginator for ListVolumes
type ListVolumesPaginator struct {
	options   ListVolumesPaginatorOptions
	client    ListVolumesAPIClient
	params    *ListVolumesInput
	nextToken *string
	firstPage bool
}

// NewListVolumesPaginator returns a new ListVolumesPaginator
func NewListVolumesPaginator(client ListVolumesAPIClient, params *ListVolumesInput, optFns ...func(*ListVolumesPaginatorOptions)) *ListVolumesPaginator {
	if params == nil {
		params = &ListVolumesInput{}
	}

	options := ListVolumesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVolumesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVolumesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVolumes page.
func (p *ListVolumesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVolumesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListVolumes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListVolumes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListVolumes",
	}
}
