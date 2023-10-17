// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the specified Amazon Redshift HSM configuration. If
// no configuration ID is specified, returns information about all the HSM
// configurations owned by your Amazon Web Services account. If you specify both
// tag keys and tag values in the same request, Amazon Redshift returns all HSM
// connections that match any combination of the specified keys and values. For
// example, if you have owner and environment for tag keys, and admin and test for
// tag values, all HSM connections that have any combination of those values are
// returned. If both tag keys and values are omitted from the request, HSM
// connections are returned regardless of whether they have tag keys or values
// associated with them.
func (c *Client) DescribeHsmConfigurations(ctx context.Context, params *DescribeHsmConfigurationsInput, optFns ...func(*Options)) (*DescribeHsmConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeHsmConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHsmConfigurations", params, optFns, c.addOperationDescribeHsmConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHsmConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHsmConfigurationsInput struct {

	// The identifier of a specific Amazon Redshift HSM configuration to be described.
	// If no identifier is specified, information is returned for all HSM
	// configurations owned by your Amazon Web Services account.
	HsmConfigurationIdentifier *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeHsmConfigurations request
	// exceed the value specified in MaxRecords , Amazon Web Services returns a value
	// in the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// A tag key or keys for which you want to return all matching HSM configurations
	// that are associated with the specified key or keys. For example, suppose that
	// you have HSM configurations that are tagged with keys called owner and
	// environment . If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the HSM configurations that have either or both
	// of these tag keys associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching HSM
	// configurations that are associated with the specified tag value or values. For
	// example, suppose that you have HSM configurations that are tagged with values
	// called admin and test . If you specify both of these tag values in the request,
	// Amazon Redshift returns a response with the HSM configurations that have either
	// or both of these tag values associated with them.
	TagValues []string

	noSmithyDocumentSerde
}

type DescribeHsmConfigurationsOutput struct {

	// A list of HsmConfiguration objects.
	HsmConfigurations []types.HsmConfiguration

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHsmConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeHsmConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeHsmConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHsmConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeHsmConfigurationsAPIClient is a client that implements the
// DescribeHsmConfigurations operation.
type DescribeHsmConfigurationsAPIClient interface {
	DescribeHsmConfigurations(context.Context, *DescribeHsmConfigurationsInput, ...func(*Options)) (*DescribeHsmConfigurationsOutput, error)
}

var _ DescribeHsmConfigurationsAPIClient = (*Client)(nil)

// DescribeHsmConfigurationsPaginatorOptions is the paginator options for
// DescribeHsmConfigurations
type DescribeHsmConfigurationsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeHsmConfigurationsPaginator is a paginator for DescribeHsmConfigurations
type DescribeHsmConfigurationsPaginator struct {
	options   DescribeHsmConfigurationsPaginatorOptions
	client    DescribeHsmConfigurationsAPIClient
	params    *DescribeHsmConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeHsmConfigurationsPaginator returns a new
// DescribeHsmConfigurationsPaginator
func NewDescribeHsmConfigurationsPaginator(client DescribeHsmConfigurationsAPIClient, params *DescribeHsmConfigurationsInput, optFns ...func(*DescribeHsmConfigurationsPaginatorOptions)) *DescribeHsmConfigurationsPaginator {
	if params == nil {
		params = &DescribeHsmConfigurationsInput{}
	}

	options := DescribeHsmConfigurationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeHsmConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeHsmConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeHsmConfigurations page.
func (p *DescribeHsmConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeHsmConfigurationsOutput, error) {
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

	result, err := p.client.DescribeHsmConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeHsmConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeHsmConfigurations",
	}
}
