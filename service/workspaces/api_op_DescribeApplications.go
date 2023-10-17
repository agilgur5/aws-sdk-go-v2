// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified applications by filtering based on their compute types,
// license availability, operating systems, and owners.
func (c *Client) DescribeApplications(ctx context.Context, params *DescribeApplicationsInput, optFns ...func(*Options)) (*DescribeApplicationsOutput, error) {
	if params == nil {
		params = &DescribeApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplications", params, optFns, c.addOperationDescribeApplicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeApplicationsInput struct {

	// The identifiers of one or more applications.
	ApplicationIds []string

	// The compute types supported by the applications.
	ComputeTypeNames []types.Compute

	// The license availability for the applications.
	LicenseType types.WorkSpaceApplicationLicenseType

	// The maximum number of applications to return.
	MaxResults *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string

	// The operating systems supported by the applications.
	OperatingSystemNames []types.OperatingSystemName

	// The owner of the applications.
	Owner *string

	noSmithyDocumentSerde
}

type DescribeApplicationsOutput struct {

	// List of information about the specified applications.
	Applications []types.WorkSpaceApplication

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApplications{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplications(options.Region), middleware.Before); err != nil {
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

// DescribeApplicationsAPIClient is a client that implements the
// DescribeApplications operation.
type DescribeApplicationsAPIClient interface {
	DescribeApplications(context.Context, *DescribeApplicationsInput, ...func(*Options)) (*DescribeApplicationsOutput, error)
}

var _ DescribeApplicationsAPIClient = (*Client)(nil)

// DescribeApplicationsPaginatorOptions is the paginator options for
// DescribeApplications
type DescribeApplicationsPaginatorOptions struct {
	// The maximum number of applications to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeApplicationsPaginator is a paginator for DescribeApplications
type DescribeApplicationsPaginator struct {
	options   DescribeApplicationsPaginatorOptions
	client    DescribeApplicationsAPIClient
	params    *DescribeApplicationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeApplicationsPaginator returns a new DescribeApplicationsPaginator
func NewDescribeApplicationsPaginator(client DescribeApplicationsAPIClient, params *DescribeApplicationsInput, optFns ...func(*DescribeApplicationsPaginatorOptions)) *DescribeApplicationsPaginator {
	if params == nil {
		params = &DescribeApplicationsInput{}
	}

	options := DescribeApplicationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeApplicationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeApplicationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeApplications page.
func (p *DescribeApplicationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeApplicationsOutput, error) {
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

	result, err := p.client.DescribeApplications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeApplications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeApplications",
	}
}
