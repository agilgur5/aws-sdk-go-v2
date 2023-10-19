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

// This API is in preview release for Amazon Connect and is subject to change. For
// the specified referenceTypes , returns a list of references associated with the
// contact.
func (c *Client) ListContactReferences(ctx context.Context, params *ListContactReferencesInput, optFns ...func(*Options)) (*ListContactReferencesOutput, error) {
	if params == nil {
		params = &ListContactReferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContactReferences", params, optFns, c.addOperationListContactReferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContactReferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContactReferencesInput struct {

	// The identifier of the initial contact.
	//
	// This member is required.
	ContactId *string

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The type of reference.
	//
	// This member is required.
	ReferenceTypes []types.ReferenceType

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results. This is not
	// expected to be set, because the value returned in the previous response is
	// always null.
	NextToken *string

	noSmithyDocumentSerde
}

type ListContactReferencesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	// This is always returned as null in the response.
	NextToken *string

	// Information about the flows.
	ReferenceSummaryList []types.ReferenceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContactReferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContactReferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContactReferences{}, middleware.After)
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
	if err = addOpListContactReferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContactReferences(options.Region), middleware.Before); err != nil {
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
		operation: "ListContactReferences",
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

// ListContactReferencesAPIClient is a client that implements the
// ListContactReferences operation.
type ListContactReferencesAPIClient interface {
	ListContactReferences(context.Context, *ListContactReferencesInput, ...func(*Options)) (*ListContactReferencesOutput, error)
}

var _ ListContactReferencesAPIClient = (*Client)(nil)

// ListContactReferencesPaginatorOptions is the paginator options for
// ListContactReferences
type ListContactReferencesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContactReferencesPaginator is a paginator for ListContactReferences
type ListContactReferencesPaginator struct {
	options   ListContactReferencesPaginatorOptions
	client    ListContactReferencesAPIClient
	params    *ListContactReferencesInput
	nextToken *string
	firstPage bool
}

// NewListContactReferencesPaginator returns a new ListContactReferencesPaginator
func NewListContactReferencesPaginator(client ListContactReferencesAPIClient, params *ListContactReferencesInput, optFns ...func(*ListContactReferencesPaginatorOptions)) *ListContactReferencesPaginator {
	if params == nil {
		params = &ListContactReferencesInput{}
	}

	options := ListContactReferencesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContactReferencesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContactReferencesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContactReferences page.
func (p *ListContactReferencesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContactReferencesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListContactReferences(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContactReferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListContactReferences",
	}
}
