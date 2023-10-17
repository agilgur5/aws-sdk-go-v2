// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the backups of a streaming session in a studio.
func (c *Client) ListStreamingSessionBackups(ctx context.Context, params *ListStreamingSessionBackupsInput, optFns ...func(*Options)) (*ListStreamingSessionBackupsOutput, error) {
	if params == nil {
		params = &ListStreamingSessionBackupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamingSessionBackups", params, optFns, c.addOperationListStreamingSessionBackupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamingSessionBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamingSessionBackupsInput struct {

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The user ID of the user that owns the streaming session.
	OwnedBy *string

	noSmithyDocumentSerde
}

type ListStreamingSessionBackupsOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Information about the streaming session backups.
	StreamingSessionBackups []types.StreamingSessionBackup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamingSessionBackupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStreamingSessionBackups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStreamingSessionBackups{}, middleware.After)
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
	if err = addOpListStreamingSessionBackupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamingSessionBackups(options.Region), middleware.Before); err != nil {
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

// ListStreamingSessionBackupsAPIClient is a client that implements the
// ListStreamingSessionBackups operation.
type ListStreamingSessionBackupsAPIClient interface {
	ListStreamingSessionBackups(context.Context, *ListStreamingSessionBackupsInput, ...func(*Options)) (*ListStreamingSessionBackupsOutput, error)
}

var _ ListStreamingSessionBackupsAPIClient = (*Client)(nil)

// ListStreamingSessionBackupsPaginatorOptions is the paginator options for
// ListStreamingSessionBackups
type ListStreamingSessionBackupsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamingSessionBackupsPaginator is a paginator for
// ListStreamingSessionBackups
type ListStreamingSessionBackupsPaginator struct {
	options   ListStreamingSessionBackupsPaginatorOptions
	client    ListStreamingSessionBackupsAPIClient
	params    *ListStreamingSessionBackupsInput
	nextToken *string
	firstPage bool
}

// NewListStreamingSessionBackupsPaginator returns a new
// ListStreamingSessionBackupsPaginator
func NewListStreamingSessionBackupsPaginator(client ListStreamingSessionBackupsAPIClient, params *ListStreamingSessionBackupsInput, optFns ...func(*ListStreamingSessionBackupsPaginatorOptions)) *ListStreamingSessionBackupsPaginator {
	if params == nil {
		params = &ListStreamingSessionBackupsInput{}
	}

	options := ListStreamingSessionBackupsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamingSessionBackupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamingSessionBackupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreamingSessionBackups page.
func (p *ListStreamingSessionBackupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamingSessionBackupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListStreamingSessionBackups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStreamingSessionBackups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "ListStreamingSessionBackups",
	}
}
