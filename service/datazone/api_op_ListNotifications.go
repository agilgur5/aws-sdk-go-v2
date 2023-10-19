// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all Amazon DataZone notifications.
func (c *Client) ListNotifications(ctx context.Context, params *ListNotificationsInput, optFns ...func(*Options)) (*ListNotificationsOutput, error) {
	if params == nil {
		params = &ListNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotifications", params, optFns, c.addOperationListNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotificationsInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// The type of notifications.
	//
	// This member is required.
	Type types.NotificationType

	// The time after which you want to list notifications.
	AfterTimestamp *time.Time

	// The time before which you want to list notifications.
	BeforeTimestamp *time.Time

	// The maximum number of notifications to return in a single call to
	// ListNotifications . When the number of notifications to be listed is greater
	// than the value of MaxResults , the response contains a NextToken value that you
	// can use in a subsequent call to ListNotifications to list the next set of
	// notifications.
	MaxResults *int32

	// When the number of notifications is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of notifications, the response includes a pagination
	// token named NextToken . You can specify this NextToken value in a subsequent
	// call to ListNotifications to list the next set of notifications.
	NextToken *string

	// The subjects of notifications.
	Subjects []string

	// The task status of notifications.
	TaskStatus types.TaskStatus

	noSmithyDocumentSerde
}

type ListNotificationsOutput struct {

	// When the number of notifications is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of notifications, the response includes a pagination
	// token named NextToken . You can specify this NextToken value in a subsequent
	// call to ListNotifications to list the next set of notifications.
	NextToken *string

	// The results of the ListNotifications action.
	Notifications []types.NotificationOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNotifications{}, middleware.After)
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
	if err = addOpListNotificationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotifications(options.Region), middleware.Before); err != nil {
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
		operation: "ListNotifications",
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

// ListNotificationsAPIClient is a client that implements the ListNotifications
// operation.
type ListNotificationsAPIClient interface {
	ListNotifications(context.Context, *ListNotificationsInput, ...func(*Options)) (*ListNotificationsOutput, error)
}

var _ ListNotificationsAPIClient = (*Client)(nil)

// ListNotificationsPaginatorOptions is the paginator options for ListNotifications
type ListNotificationsPaginatorOptions struct {
	// The maximum number of notifications to return in a single call to
	// ListNotifications . When the number of notifications to be listed is greater
	// than the value of MaxResults , the response contains a NextToken value that you
	// can use in a subsequent call to ListNotifications to list the next set of
	// notifications.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNotificationsPaginator is a paginator for ListNotifications
type ListNotificationsPaginator struct {
	options   ListNotificationsPaginatorOptions
	client    ListNotificationsAPIClient
	params    *ListNotificationsInput
	nextToken *string
	firstPage bool
}

// NewListNotificationsPaginator returns a new ListNotificationsPaginator
func NewListNotificationsPaginator(client ListNotificationsAPIClient, params *ListNotificationsInput, optFns ...func(*ListNotificationsPaginatorOptions)) *ListNotificationsPaginator {
	if params == nil {
		params = &ListNotificationsInput{}
	}

	options := ListNotificationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNotificationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNotificationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNotifications page.
func (p *ListNotificationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNotificationsOutput, error) {
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

	result, err := p.client.ListNotifications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "ListNotifications",
	}
}
