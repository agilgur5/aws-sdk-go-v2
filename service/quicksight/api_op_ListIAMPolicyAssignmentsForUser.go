// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the IAM policy assignments, including the Amazon Resource Names
// (ARNs), for the IAM policies assigned to the specified user and group, or groups
// that the user belongs to.
func (c *Client) ListIAMPolicyAssignmentsForUser(ctx context.Context, params *ListIAMPolicyAssignmentsForUserInput, optFns ...func(*Options)) (*ListIAMPolicyAssignmentsForUserOutput, error) {
	if params == nil {
		params = &ListIAMPolicyAssignmentsForUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIAMPolicyAssignmentsForUser", params, optFns, c.addOperationListIAMPolicyAssignmentsForUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIAMPolicyAssignmentsForUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIAMPolicyAssignmentsForUserInput struct {

	// The ID of the Amazon Web Services account that contains the assignments.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace of the assignment.
	//
	// This member is required.
	Namespace *string

	// The name of the user.
	//
	// This member is required.
	UserName *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListIAMPolicyAssignmentsForUserOutput struct {

	// The active assignments for this user.
	ActiveAssignments []types.ActiveIAMPolicyAssignment

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIAMPolicyAssignmentsForUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIAMPolicyAssignmentsForUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIAMPolicyAssignmentsForUser{}, middleware.After)
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
	if err = addOpListIAMPolicyAssignmentsForUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIAMPolicyAssignmentsForUser(options.Region), middleware.Before); err != nil {
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

// ListIAMPolicyAssignmentsForUserAPIClient is a client that implements the
// ListIAMPolicyAssignmentsForUser operation.
type ListIAMPolicyAssignmentsForUserAPIClient interface {
	ListIAMPolicyAssignmentsForUser(context.Context, *ListIAMPolicyAssignmentsForUserInput, ...func(*Options)) (*ListIAMPolicyAssignmentsForUserOutput, error)
}

var _ ListIAMPolicyAssignmentsForUserAPIClient = (*Client)(nil)

// ListIAMPolicyAssignmentsForUserPaginatorOptions is the paginator options for
// ListIAMPolicyAssignmentsForUser
type ListIAMPolicyAssignmentsForUserPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIAMPolicyAssignmentsForUserPaginator is a paginator for
// ListIAMPolicyAssignmentsForUser
type ListIAMPolicyAssignmentsForUserPaginator struct {
	options   ListIAMPolicyAssignmentsForUserPaginatorOptions
	client    ListIAMPolicyAssignmentsForUserAPIClient
	params    *ListIAMPolicyAssignmentsForUserInput
	nextToken *string
	firstPage bool
}

// NewListIAMPolicyAssignmentsForUserPaginator returns a new
// ListIAMPolicyAssignmentsForUserPaginator
func NewListIAMPolicyAssignmentsForUserPaginator(client ListIAMPolicyAssignmentsForUserAPIClient, params *ListIAMPolicyAssignmentsForUserInput, optFns ...func(*ListIAMPolicyAssignmentsForUserPaginatorOptions)) *ListIAMPolicyAssignmentsForUserPaginator {
	if params == nil {
		params = &ListIAMPolicyAssignmentsForUserInput{}
	}

	options := ListIAMPolicyAssignmentsForUserPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIAMPolicyAssignmentsForUserPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIAMPolicyAssignmentsForUserPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIAMPolicyAssignmentsForUser page.
func (p *ListIAMPolicyAssignmentsForUserPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIAMPolicyAssignmentsForUserOutput, error) {
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

	result, err := p.client.ListIAMPolicyAssignmentsForUser(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIAMPolicyAssignmentsForUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListIAMPolicyAssignmentsForUser",
	}
}
