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

// Provides information about any domain controllers in your directory.
func (c *Client) DescribeDomainControllers(ctx context.Context, params *DescribeDomainControllersInput, optFns ...func(*Options)) (*DescribeDomainControllersOutput, error) {
	if params == nil {
		params = &DescribeDomainControllersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomainControllers", params, optFns, c.addOperationDescribeDomainControllersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainControllersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDomainControllersInput struct {

	// Identifier of the directory for which to retrieve the domain controller
	// information.
	//
	// This member is required.
	DirectoryId *string

	// A list of identifiers for the domain controllers whose information will be
	// provided.
	DomainControllerIds []string

	// The maximum number of items to return.
	Limit *int32

	// The DescribeDomainControllers.NextToken value from a previous call to
	// DescribeDomainControllers . Pass null if this is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeDomainControllersOutput struct {

	// List of the DomainController objects that were retrieved.
	DomainControllers []types.DomainController

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to DescribeDomainControllers retrieve the next
	// set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDomainControllersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDomainControllers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDomainControllers{}, middleware.After)
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
	if err = addOpDescribeDomainControllersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomainControllers(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeDomainControllers",
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

// DescribeDomainControllersAPIClient is a client that implements the
// DescribeDomainControllers operation.
type DescribeDomainControllersAPIClient interface {
	DescribeDomainControllers(context.Context, *DescribeDomainControllersInput, ...func(*Options)) (*DescribeDomainControllersOutput, error)
}

var _ DescribeDomainControllersAPIClient = (*Client)(nil)

// DescribeDomainControllersPaginatorOptions is the paginator options for
// DescribeDomainControllers
type DescribeDomainControllersPaginatorOptions struct {
	// The maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDomainControllersPaginator is a paginator for DescribeDomainControllers
type DescribeDomainControllersPaginator struct {
	options   DescribeDomainControllersPaginatorOptions
	client    DescribeDomainControllersAPIClient
	params    *DescribeDomainControllersInput
	nextToken *string
	firstPage bool
}

// NewDescribeDomainControllersPaginator returns a new
// DescribeDomainControllersPaginator
func NewDescribeDomainControllersPaginator(client DescribeDomainControllersAPIClient, params *DescribeDomainControllersInput, optFns ...func(*DescribeDomainControllersPaginatorOptions)) *DescribeDomainControllersPaginator {
	if params == nil {
		params = &DescribeDomainControllersInput{}
	}

	options := DescribeDomainControllersPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDomainControllersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDomainControllersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDomainControllers page.
func (p *DescribeDomainControllersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDomainControllersOutput, error) {
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

	result, err := p.client.DescribeDomainControllers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDomainControllers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeDomainControllers",
	}
}
