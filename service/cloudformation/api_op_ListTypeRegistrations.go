// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of registration tokens for the specified extension(s).
func (c *Client) ListTypeRegistrations(ctx context.Context, params *ListTypeRegistrationsInput, optFns ...func(*Options)) (*ListTypeRegistrationsOutput, error) {
	if params == nil {
		params = &ListTypeRegistrationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTypeRegistrations", params, optFns, c.addOperationListTypeRegistrationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTypeRegistrationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypeRegistrationsInput struct {

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int32

	// If the previous paginated request didn't return all the remaining results, the
	// response object's NextToken parameter value is set to a token. To retrieve the
	// next set of results, call this action again and assign that token to the request
	// object's NextToken parameter. If there are no remaining results, the previous
	// response object's NextToken parameter is set to null .
	NextToken *string

	// The current status of the extension registration request. The default is
	// IN_PROGRESS .
	RegistrationStatusFilter types.RegistrationStatus

	// The kind of extension. Conditional: You must specify either TypeName and Type ,
	// or Arn .
	Type types.RegistryType

	// The Amazon Resource Name (ARN) of the extension. Conditional: You must specify
	// either TypeName and Type , or Arn .
	TypeArn *string

	// The name of the extension. Conditional: You must specify either TypeName and
	// Type , or Arn .
	TypeName *string

	noSmithyDocumentSerde
}

type ListTypeRegistrationsOutput struct {

	// If the request doesn't return all the remaining results, NextToken is set to a
	// token. To retrieve the next set of results, call this action again and assign
	// that token to the request object's NextToken parameter. If the request returns
	// all results, NextToken is set to null .
	NextToken *string

	// A list of extension registration tokens. Use DescribeTypeRegistration to return
	// detailed information about a type registration request.
	RegistrationTokenList []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTypeRegistrationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListTypeRegistrations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListTypeRegistrations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTypeRegistrations(options.Region), middleware.Before); err != nil {
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
		operation: "ListTypeRegistrations",
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

// ListTypeRegistrationsAPIClient is a client that implements the
// ListTypeRegistrations operation.
type ListTypeRegistrationsAPIClient interface {
	ListTypeRegistrations(context.Context, *ListTypeRegistrationsInput, ...func(*Options)) (*ListTypeRegistrationsOutput, error)
}

var _ ListTypeRegistrationsAPIClient = (*Client)(nil)

// ListTypeRegistrationsPaginatorOptions is the paginator options for
// ListTypeRegistrations
type ListTypeRegistrationsPaginatorOptions struct {
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTypeRegistrationsPaginator is a paginator for ListTypeRegistrations
type ListTypeRegistrationsPaginator struct {
	options   ListTypeRegistrationsPaginatorOptions
	client    ListTypeRegistrationsAPIClient
	params    *ListTypeRegistrationsInput
	nextToken *string
	firstPage bool
}

// NewListTypeRegistrationsPaginator returns a new ListTypeRegistrationsPaginator
func NewListTypeRegistrationsPaginator(client ListTypeRegistrationsAPIClient, params *ListTypeRegistrationsInput, optFns ...func(*ListTypeRegistrationsPaginatorOptions)) *ListTypeRegistrationsPaginator {
	if params == nil {
		params = &ListTypeRegistrationsInput{}
	}

	options := ListTypeRegistrationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTypeRegistrationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTypeRegistrationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTypeRegistrations page.
func (p *ListTypeRegistrationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTypeRegistrationsOutput, error) {
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

	result, err := p.client.ListTypeRegistrations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTypeRegistrations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListTypeRegistrations",
	}
}
