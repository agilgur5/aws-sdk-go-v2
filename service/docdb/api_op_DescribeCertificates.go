// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of certificate authority (CA) certificates provided by Amazon
// DocumentDB for this Amazon Web Services account.
func (c *Client) DescribeCertificates(ctx context.Context, params *DescribeCertificatesInput, optFns ...func(*Options)) (*DescribeCertificatesOutput, error) {
	if params == nil {
		params = &DescribeCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCertificates", params, optFns, c.addOperationDescribeCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCertificatesInput struct {

	// The user-supplied certificate identifier. If this parameter is specified,
	// information for only the specified certificate is returned. If this parameter is
	// omitted, a list of up to MaxRecords certificates is returned. This parameter is
	// not case sensitive. Constraints
	//   - Must match an existing CertificateIdentifier .
	CertificateIdentifier *string

	// This parameter is not currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeCertificates
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints:
	//   - Minimum: 20
	//   - Maximum: 100
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeCertificatesOutput struct {

	// A list of certificates for this Amazon Web Services account.
	Certificates []types.Certificate

	// An optional pagination token provided if the number of records retrieved is
	// greater than MaxRecords . If this parameter is specified, the marker specifies
	// the next record in the list. Including the value of Marker in the next call to
	// DescribeCertificates results in the next page of certificates.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCertificates{}, middleware.After)
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
	if err = addOpDescribeCertificatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCertificates(options.Region), middleware.Before); err != nil {
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

// DescribeCertificatesAPIClient is a client that implements the
// DescribeCertificates operation.
type DescribeCertificatesAPIClient interface {
	DescribeCertificates(context.Context, *DescribeCertificatesInput, ...func(*Options)) (*DescribeCertificatesOutput, error)
}

var _ DescribeCertificatesAPIClient = (*Client)(nil)

// DescribeCertificatesPaginatorOptions is the paginator options for
// DescribeCertificates
type DescribeCertificatesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints:
	//   - Minimum: 20
	//   - Maximum: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCertificatesPaginator is a paginator for DescribeCertificates
type DescribeCertificatesPaginator struct {
	options   DescribeCertificatesPaginatorOptions
	client    DescribeCertificatesAPIClient
	params    *DescribeCertificatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeCertificatesPaginator returns a new DescribeCertificatesPaginator
func NewDescribeCertificatesPaginator(client DescribeCertificatesAPIClient, params *DescribeCertificatesInput, optFns ...func(*DescribeCertificatesPaginatorOptions)) *DescribeCertificatesPaginator {
	if params == nil {
		params = &DescribeCertificatesInput{}
	}

	options := DescribeCertificatesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCertificates page.
func (p *DescribeCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCertificatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeCertificates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeCertificates",
	}
}
