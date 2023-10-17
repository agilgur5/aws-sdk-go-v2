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

// Provides information for all active Amazon EC2 instances and Amazon EC2
// instances terminated in the last 30 days, up to a maximum of 2,000. Amazon EC2
// instances in any of the following states are considered active:
// AWAITING_FULFILLMENT, PROVISIONING, BOOTSTRAPPING, RUNNING.
func (c *Client) ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) {
	if params == nil {
		params = &ListInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstances", params, optFns, c.addOperationListInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which instances to list.
type ListInstancesInput struct {

	// The identifier of the cluster for which to list the instances.
	//
	// This member is required.
	ClusterId *string

	// The unique identifier of the instance fleet.
	InstanceFleetId *string

	// The node type of the instance fleet. For example MASTER, CORE, or TASK.
	InstanceFleetType types.InstanceFleetType

	// The identifier of the instance group for which to list the instances.
	InstanceGroupId *string

	// The type of instance group for which to list the instances.
	InstanceGroupTypes []types.InstanceGroupType

	// A list of instance states that will filter the instances returned with this
	// request.
	InstanceStates []types.InstanceState

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	noSmithyDocumentSerde
}

// This output contains the list of instances.
type ListInstancesOutput struct {

	// The list of instances for the cluster and given filters.
	Instances []types.Instance

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInstances{}, middleware.After)
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
	if err = addOpListInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstances(options.Region), middleware.Before); err != nil {
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

// ListInstancesAPIClient is a client that implements the ListInstances operation.
type ListInstancesAPIClient interface {
	ListInstances(context.Context, *ListInstancesInput, ...func(*Options)) (*ListInstancesOutput, error)
}

var _ ListInstancesAPIClient = (*Client)(nil)

// ListInstancesPaginatorOptions is the paginator options for ListInstances
type ListInstancesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstancesPaginator is a paginator for ListInstances
type ListInstancesPaginator struct {
	options   ListInstancesPaginatorOptions
	client    ListInstancesAPIClient
	params    *ListInstancesInput
	nextToken *string
	firstPage bool
}

// NewListInstancesPaginator returns a new ListInstancesPaginator
func NewListInstancesPaginator(client ListInstancesAPIClient, params *ListInstancesInput, optFns ...func(*ListInstancesPaginatorOptions)) *ListInstancesPaginator {
	if params == nil {
		params = &ListInstancesInput{}
	}

	options := ListInstancesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInstances page.
func (p *ListInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListInstances",
	}
}
