// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A list of the instance types that Amazon EMR supports. You can filter the list
// by Amazon Web Services Region and Amazon EMR release.
func (c *Client) ListSupportedInstanceTypes(ctx context.Context, params *ListSupportedInstanceTypesInput, optFns ...func(*Options)) (*ListSupportedInstanceTypesOutput, error) {
	if params == nil {
		params = &ListSupportedInstanceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSupportedInstanceTypes", params, optFns, c.addOperationListSupportedInstanceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSupportedInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSupportedInstanceTypesInput struct {

	// The Amazon EMR release label determines the versions of open-source application
	// packages (https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-release-app-versions-6.x.html)
	// that Amazon EMR has installed on the cluster. Release labels are in the format
	// emr-x.x.x , where x.x.x is an Amazon EMR release number such as emr-6.10.0 . For
	// more information about Amazon EMR releases and their included application
	// versions and features, see the Amazon EMR Release Guide (https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-release-components.html)
	// .
	//
	// This member is required.
	ReleaseLabel *string

	// The pagination token that marks the next set of results to retrieve.
	Marker *string

	noSmithyDocumentSerde
}

type ListSupportedInstanceTypesOutput struct {

	// The pagination token that marks the next set of results to retrieve.
	Marker *string

	// The list of instance types that the release specified in
	// ListSupportedInstanceTypesInput$ReleaseLabel supports, filtered by Amazon Web
	// Services Region.
	SupportedInstanceTypes []types.SupportedInstanceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSupportedInstanceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSupportedInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSupportedInstanceTypes{}, middleware.After)
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
	if err = addOpListSupportedInstanceTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSupportedInstanceTypes(options.Region), middleware.Before); err != nil {
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
		operation: "ListSupportedInstanceTypes",
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

// ListSupportedInstanceTypesAPIClient is a client that implements the
// ListSupportedInstanceTypes operation.
type ListSupportedInstanceTypesAPIClient interface {
	ListSupportedInstanceTypes(context.Context, *ListSupportedInstanceTypesInput, ...func(*Options)) (*ListSupportedInstanceTypesOutput, error)
}

var _ ListSupportedInstanceTypesAPIClient = (*Client)(nil)

// ListSupportedInstanceTypesPaginatorOptions is the paginator options for
// ListSupportedInstanceTypes
type ListSupportedInstanceTypesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSupportedInstanceTypesPaginator is a paginator for
// ListSupportedInstanceTypes
type ListSupportedInstanceTypesPaginator struct {
	options   ListSupportedInstanceTypesPaginatorOptions
	client    ListSupportedInstanceTypesAPIClient
	params    *ListSupportedInstanceTypesInput
	nextToken *string
	firstPage bool
}

// NewListSupportedInstanceTypesPaginator returns a new
// ListSupportedInstanceTypesPaginator
func NewListSupportedInstanceTypesPaginator(client ListSupportedInstanceTypesAPIClient, params *ListSupportedInstanceTypesInput, optFns ...func(*ListSupportedInstanceTypesPaginatorOptions)) *ListSupportedInstanceTypesPaginator {
	if params == nil {
		params = &ListSupportedInstanceTypesInput{}
	}

	options := ListSupportedInstanceTypesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSupportedInstanceTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSupportedInstanceTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSupportedInstanceTypes page.
func (p *ListSupportedInstanceTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSupportedInstanceTypesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListSupportedInstanceTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSupportedInstanceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListSupportedInstanceTypes",
	}
}
