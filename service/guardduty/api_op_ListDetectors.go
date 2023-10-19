// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists detectorIds of all the existing Amazon GuardDuty detector resources.
func (c *Client) ListDetectors(ctx context.Context, params *ListDetectorsInput, optFns ...func(*Options)) (*ListDetectorsOutput, error) {
	if params == nil {
		params = &ListDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDetectors", params, optFns, c.addOperationListDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDetectorsInput struct {

	// You can use this parameter to indicate the maximum number of items that you
	// want in the response. The default value is 50. The maximum value is 50.
	MaxResults int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls to
	// the action, fill nextToken in the request with the value of NextToken from the
	// previous response to continue listing data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDetectorsOutput struct {

	// A list of detector IDs.
	//
	// This member is required.
	DetectorIds []string

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDetectors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDetectors(options.Region), middleware.Before); err != nil {
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
		operation: "ListDetectors",
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

// ListDetectorsAPIClient is a client that implements the ListDetectors operation.
type ListDetectorsAPIClient interface {
	ListDetectors(context.Context, *ListDetectorsInput, ...func(*Options)) (*ListDetectorsOutput, error)
}

var _ ListDetectorsAPIClient = (*Client)(nil)

// ListDetectorsPaginatorOptions is the paginator options for ListDetectors
type ListDetectorsPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items that you
	// want in the response. The default value is 50. The maximum value is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDetectorsPaginator is a paginator for ListDetectors
type ListDetectorsPaginator struct {
	options   ListDetectorsPaginatorOptions
	client    ListDetectorsAPIClient
	params    *ListDetectorsInput
	nextToken *string
	firstPage bool
}

// NewListDetectorsPaginator returns a new ListDetectorsPaginator
func NewListDetectorsPaginator(client ListDetectorsAPIClient, params *ListDetectorsInput, optFns ...func(*ListDetectorsPaginatorOptions)) *ListDetectorsPaginator {
	if params == nil {
		params = &ListDetectorsInput{}
	}

	options := ListDetectorsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDetectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDetectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDetectors page.
func (p *ListDetectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDetectorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDetectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDetectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListDetectors",
	}
}
