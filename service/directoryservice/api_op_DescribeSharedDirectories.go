// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the shared directories in your account.
func (c *Client) DescribeSharedDirectories(ctx context.Context, params *DescribeSharedDirectoriesInput, optFns ...func(*Options)) (*DescribeSharedDirectoriesOutput, error) {
	if params == nil {
		params = &DescribeSharedDirectoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSharedDirectories", params, optFns, c.addOperationDescribeSharedDirectoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSharedDirectoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSharedDirectoriesInput struct {

	// Returns the identifier of the directory in the directory owner account.
	//
	// This member is required.
	OwnerDirectoryId *string

	// The number of shared directories to return in the response object.
	Limit *int32

	// The DescribeSharedDirectoriesResult.NextToken value from a previous call to
	// DescribeSharedDirectories . Pass null if this is the first call.
	NextToken *string

	// A list of identifiers of all shared directories in your account.
	SharedDirectoryIds []string

	noSmithyDocumentSerde
}

type DescribeSharedDirectoriesOutput struct {

	// If not null, token that indicates that more results are available. Pass this
	// value for the NextToken parameter in a subsequent call to
	// DescribeSharedDirectories to retrieve the next set of items.
	NextToken *string

	// A list of all shared directories in your account.
	SharedDirectories []types.SharedDirectory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSharedDirectoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSharedDirectories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSharedDirectories{}, middleware.After)
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
	if err = addOpDescribeSharedDirectoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSharedDirectories(options.Region), middleware.Before); err != nil {
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

// DescribeSharedDirectoriesAPIClient is a client that implements the
// DescribeSharedDirectories operation.
type DescribeSharedDirectoriesAPIClient interface {
	DescribeSharedDirectories(context.Context, *DescribeSharedDirectoriesInput, ...func(*Options)) (*DescribeSharedDirectoriesOutput, error)
}

var _ DescribeSharedDirectoriesAPIClient = (*Client)(nil)

// DescribeSharedDirectoriesPaginatorOptions is the paginator options for
// DescribeSharedDirectories
type DescribeSharedDirectoriesPaginatorOptions struct {
	// The number of shared directories to return in the response object.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSharedDirectoriesPaginator is a paginator for DescribeSharedDirectories
type DescribeSharedDirectoriesPaginator struct {
	options   DescribeSharedDirectoriesPaginatorOptions
	client    DescribeSharedDirectoriesAPIClient
	params    *DescribeSharedDirectoriesInput
	nextToken *string
	firstPage bool
}

// NewDescribeSharedDirectoriesPaginator returns a new
// DescribeSharedDirectoriesPaginator
func NewDescribeSharedDirectoriesPaginator(client DescribeSharedDirectoriesAPIClient, params *DescribeSharedDirectoriesInput, optFns ...func(*DescribeSharedDirectoriesPaginatorOptions)) *DescribeSharedDirectoriesPaginator {
	if params == nil {
		params = &DescribeSharedDirectoriesInput{}
	}

	options := DescribeSharedDirectoriesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSharedDirectoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSharedDirectoriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSharedDirectories page.
func (p *DescribeSharedDirectoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSharedDirectoriesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeSharedDirectories(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSharedDirectories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeSharedDirectories",
	}
}
