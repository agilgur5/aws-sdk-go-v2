// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of third party applications in a specific security profile.
func (c *Client) ListSecurityProfileApplications(ctx context.Context, params *ListSecurityProfileApplicationsInput, optFns ...func(*Options)) (*ListSecurityProfileApplicationsOutput, error) {
	if params == nil {
		params = &ListSecurityProfileApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityProfileApplications", params, optFns, c.addOperationListSecurityProfileApplicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityProfileApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityProfileApplicationsInput struct {

	// The instance identifier.
	//
	// This member is required.
	InstanceId *string

	// The security profile identifier.
	//
	// This member is required.
	SecurityProfileId *string

	// The maximum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. The next set of results can be retrieved
	// by using the token value returned in the previous response when making the next
	// request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSecurityProfileApplicationsOutput struct {

	// This API is in preview release for Amazon Connect and is subject to change. A
	// list of the third party application's metadata.
	Applications []types.Application

	// The token for the next set of results. The next set of results can be retrieved
	// by using the token value returned in the previous response when making the next
	// request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecurityProfileApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityProfileApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityProfileApplications{}, middleware.After)
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
	if err = addOpListSecurityProfileApplicationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityProfileApplications(options.Region), middleware.Before); err != nil {
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
		operation: "ListSecurityProfileApplications",
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

// ListSecurityProfileApplicationsAPIClient is a client that implements the
// ListSecurityProfileApplications operation.
type ListSecurityProfileApplicationsAPIClient interface {
	ListSecurityProfileApplications(context.Context, *ListSecurityProfileApplicationsInput, ...func(*Options)) (*ListSecurityProfileApplicationsOutput, error)
}

var _ ListSecurityProfileApplicationsAPIClient = (*Client)(nil)

// ListSecurityProfileApplicationsPaginatorOptions is the paginator options for
// ListSecurityProfileApplications
type ListSecurityProfileApplicationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecurityProfileApplicationsPaginator is a paginator for
// ListSecurityProfileApplications
type ListSecurityProfileApplicationsPaginator struct {
	options   ListSecurityProfileApplicationsPaginatorOptions
	client    ListSecurityProfileApplicationsAPIClient
	params    *ListSecurityProfileApplicationsInput
	nextToken *string
	firstPage bool
}

// NewListSecurityProfileApplicationsPaginator returns a new
// ListSecurityProfileApplicationsPaginator
func NewListSecurityProfileApplicationsPaginator(client ListSecurityProfileApplicationsAPIClient, params *ListSecurityProfileApplicationsInput, optFns ...func(*ListSecurityProfileApplicationsPaginatorOptions)) *ListSecurityProfileApplicationsPaginator {
	if params == nil {
		params = &ListSecurityProfileApplicationsInput{}
	}

	options := ListSecurityProfileApplicationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecurityProfileApplicationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecurityProfileApplicationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecurityProfileApplications page.
func (p *ListSecurityProfileApplicationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecurityProfileApplicationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListSecurityProfileApplications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSecurityProfileApplications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListSecurityProfileApplications",
	}
}
