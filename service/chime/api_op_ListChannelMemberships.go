// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all channel memberships in a channel. The x-amz-chime-bearer request
// header is mandatory. Use the AppInstanceUserArn of the user that makes the API
// call as the value in the header. This API is is no longer supported and will not
// be updated. We recommend using the latest version, ListChannelMemberships (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_messaging-chime_ListChannelMemberships.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by ListChannelMemberships in the Amazon Chime SDK
// Messaging Namespace
func (c *Client) ListChannelMemberships(ctx context.Context, params *ListChannelMembershipsInput, optFns ...func(*Options)) (*ListChannelMembershipsOutput, error) {
	if params == nil {
		params = &ListChannelMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChannelMemberships", params, optFns, c.addOperationListChannelMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChannelMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChannelMembershipsInput struct {

	// The maximum number of channel memberships that you want returned.
	//
	// This member is required.
	ChannelArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	// The maximum number of channel memberships that you want returned.
	MaxResults *int32

	// The token passed by previous API calls until all requested channel memberships
	// are returned.
	NextToken *string

	// The membership type of a user, DEFAULT or HIDDEN . Default members are always
	// returned as part of ListChannelMemberships . Hidden members are only returned if
	// the type filter in ListChannelMemberships equals HIDDEN . Otherwise hidden
	// members are not returned.
	Type types.ChannelMembershipType

	noSmithyDocumentSerde
}

type ListChannelMembershipsOutput struct {

	// The ARN of the channel.
	ChannelArn *string

	// The information for the requested channel memberships.
	ChannelMemberships []types.ChannelMembershipSummary

	// The token passed by previous API calls until all requested channel memberships
	// are returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChannelMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChannelMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChannelMemberships{}, middleware.After)
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
	if err = addEndpointPrefix_opListChannelMembershipsMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpListChannelMembershipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChannelMemberships(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListChannelMembershipsMiddleware struct {
}

func (*endpointPrefix_opListChannelMembershipsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListChannelMembershipsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListChannelMembershipsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListChannelMembershipsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListChannelMembershipsAPIClient is a client that implements the
// ListChannelMemberships operation.
type ListChannelMembershipsAPIClient interface {
	ListChannelMemberships(context.Context, *ListChannelMembershipsInput, ...func(*Options)) (*ListChannelMembershipsOutput, error)
}

var _ ListChannelMembershipsAPIClient = (*Client)(nil)

// ListChannelMembershipsPaginatorOptions is the paginator options for
// ListChannelMemberships
type ListChannelMembershipsPaginatorOptions struct {
	// The maximum number of channel memberships that you want returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChannelMembershipsPaginator is a paginator for ListChannelMemberships
type ListChannelMembershipsPaginator struct {
	options   ListChannelMembershipsPaginatorOptions
	client    ListChannelMembershipsAPIClient
	params    *ListChannelMembershipsInput
	nextToken *string
	firstPage bool
}

// NewListChannelMembershipsPaginator returns a new ListChannelMembershipsPaginator
func NewListChannelMembershipsPaginator(client ListChannelMembershipsAPIClient, params *ListChannelMembershipsInput, optFns ...func(*ListChannelMembershipsPaginatorOptions)) *ListChannelMembershipsPaginator {
	if params == nil {
		params = &ListChannelMembershipsInput{}
	}

	options := ListChannelMembershipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChannelMembershipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChannelMembershipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChannelMemberships page.
func (p *ListChannelMembershipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChannelMembershipsOutput, error) {
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

	result, err := p.client.ListChannelMemberships(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListChannelMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListChannelMemberships",
	}
}
