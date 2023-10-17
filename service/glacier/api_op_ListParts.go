// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation lists the parts of an archive that have been uploaded in a
// specific multipart upload. You can make this request at any time during an
// in-progress multipart upload before you complete the upload (see
// CompleteMultipartUpload . List Parts returns an error for completed uploads. The
// list returned in the List Parts response is sorted by part range. The List Parts
// operation supports pagination. By default, this operation returns up to 50
// uploaded parts in the response. You should always check the response for a
// marker at which to continue the list; if there are no more items the marker is
// null . To return a list of parts that begins at a specific part, set the marker
// request parameter to the value you obtained from a previous List Parts request.
// You can also limit the number of parts returned in the response by specifying
// the limit parameter in the request. An AWS account has full permission to
// perform all operations (actions). However, AWS Identity and Access Management
// (IAM) users don't have any permissions by default. You must grant them explicit
// permission to perform specific actions. For more information, see Access
// Control Using AWS Identity and Access Management (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html)
// . For conceptual information and the underlying REST API, see Working with
// Archives in Amazon S3 Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html)
// and List Parts (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-list-parts.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) ListParts(ctx context.Context, params *ListPartsInput, optFns ...func(*Options)) (*ListPartsOutput, error) {
	if params == nil {
		params = &ListPartsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListParts", params, optFns, c.addOperationListPartsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving a list of parts of an archive that have been
// uploaded in a specific multipart upload.
type ListPartsInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single ' - ' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The upload ID of the multipart upload.
	//
	// This member is required.
	UploadId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The maximum number of parts to be returned. The default limit is 50. The number
	// of parts returned might be fewer than the specified limit, but the number of
	// returned parts never exceeds the limit.
	Limit *int32

	// An opaque string used for pagination. This value specifies the part at which
	// the listing of parts should begin. Get the marker value from the response of a
	// previous List Parts response. You need only include the marker if you are
	// continuing the pagination of results started in a previous List Parts request.
	Marker *string

	noSmithyDocumentSerde
}

// Contains the Amazon S3 Glacier response to your request.
type ListPartsOutput struct {

	// The description of the archive that was specified in the Initiate Multipart
	// Upload request.
	ArchiveDescription *string

	// The UTC time at which the multipart upload was initiated.
	CreationDate *string

	// An opaque string that represents where to continue pagination of the results.
	// You use the marker in a new List Parts request to obtain more jobs in the list.
	// If there are no more parts, this value is null .
	Marker *string

	// The ID of the upload to which the parts are associated.
	MultipartUploadId *string

	// The part size in bytes. This is the same value that you specified in the
	// Initiate Multipart Upload request.
	PartSizeInBytes int64

	// A list of the part sizes of the multipart upload. Each object in the array
	// contains a RangeBytes and sha256-tree-hash name/value pair.
	Parts []types.PartListElement

	// The Amazon Resource Name (ARN) of the vault to which the multipart upload was
	// initiated.
	VaultARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPartsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListParts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListParts{}, middleware.After)
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
	if err = addOpListPartsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListParts(options.Region), middleware.Before); err != nil {
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
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
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

// ListPartsAPIClient is a client that implements the ListParts operation.
type ListPartsAPIClient interface {
	ListParts(context.Context, *ListPartsInput, ...func(*Options)) (*ListPartsOutput, error)
}

var _ ListPartsAPIClient = (*Client)(nil)

// ListPartsPaginatorOptions is the paginator options for ListParts
type ListPartsPaginatorOptions struct {
	// The maximum number of parts to be returned. The default limit is 50. The number
	// of parts returned might be fewer than the specified limit, but the number of
	// returned parts never exceeds the limit.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPartsPaginator is a paginator for ListParts
type ListPartsPaginator struct {
	options   ListPartsPaginatorOptions
	client    ListPartsAPIClient
	params    *ListPartsInput
	nextToken *string
	firstPage bool
}

// NewListPartsPaginator returns a new ListPartsPaginator
func NewListPartsPaginator(client ListPartsAPIClient, params *ListPartsInput, optFns ...func(*ListPartsPaginatorOptions)) *ListPartsPaginator {
	if params == nil {
		params = &ListPartsInput{}
	}

	options := ListPartsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPartsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPartsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListParts page.
func (p *ListPartsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPartsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListParts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListParts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListParts",
	}
}
