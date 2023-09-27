// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns comments made on a pull request. Reaction counts might include numbers
// from user identities who were deleted after the reaction was made. For a count
// of reactions from active identities, use GetCommentReactions.
func (c *Client) GetCommentsForPullRequest(ctx context.Context, params *GetCommentsForPullRequestInput, optFns ...func(*Options)) (*GetCommentsForPullRequestOutput, error) {
	if params == nil {
		params = &GetCommentsForPullRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCommentsForPullRequest", params, optFns, c.addOperationGetCommentsForPullRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCommentsForPullRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommentsForPullRequestInput struct {

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests .
	//
	// This member is required.
	PullRequestId *string

	// The full commit ID of the commit in the source branch that was the tip of the
	// branch at the time the comment was made. Requirement is conditional:
	// afterCommitId must be specified when repositoryName is included.
	AfterCommitId *string

	// The full commit ID of the commit in the destination branch that was the tip of
	// the branch at the time the pull request was created. Requirement is conditional:
	// beforeCommitId must be specified when repositoryName is included.
	BeforeCommitId *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments. You can return up to 500 comments with a single
	// request.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	// The name of the repository that contains the pull request. Requirement is
	// conditional: repositoryName must be specified when beforeCommitId and
	// afterCommitId are included.
	RepositoryName *string

	noSmithyDocumentSerde
}

func (*GetCommentsForPullRequestInput) operationName() string {
	return "GetCommentsForPullRequest"
}

type GetCommentsForPullRequestOutput struct {

	// An array of comment objects on the pull request.
	CommentsForPullRequestData []types.CommentsForPullRequest

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCommentsForPullRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCommentsForPullRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCommentsForPullRequest{}, middleware.After)
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
	if err = addGetCommentsForPullRequestResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetCommentsForPullRequestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCommentsForPullRequest(options.Region), middleware.Before); err != nil {
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

// GetCommentsForPullRequestAPIClient is a client that implements the
// GetCommentsForPullRequest operation.
type GetCommentsForPullRequestAPIClient interface {
	GetCommentsForPullRequest(context.Context, *GetCommentsForPullRequestInput, ...func(*Options)) (*GetCommentsForPullRequestOutput, error)
}

var _ GetCommentsForPullRequestAPIClient = (*Client)(nil)

// GetCommentsForPullRequestPaginatorOptions is the paginator options for
// GetCommentsForPullRequest
type GetCommentsForPullRequestPaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments. You can return up to 500 comments with a single
	// request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCommentsForPullRequestPaginator is a paginator for GetCommentsForPullRequest
type GetCommentsForPullRequestPaginator struct {
	options   GetCommentsForPullRequestPaginatorOptions
	client    GetCommentsForPullRequestAPIClient
	params    *GetCommentsForPullRequestInput
	nextToken *string
	firstPage bool
}

// NewGetCommentsForPullRequestPaginator returns a new
// GetCommentsForPullRequestPaginator
func NewGetCommentsForPullRequestPaginator(client GetCommentsForPullRequestAPIClient, params *GetCommentsForPullRequestInput, optFns ...func(*GetCommentsForPullRequestPaginatorOptions)) *GetCommentsForPullRequestPaginator {
	if params == nil {
		params = &GetCommentsForPullRequestInput{}
	}

	options := GetCommentsForPullRequestPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCommentsForPullRequestPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCommentsForPullRequestPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCommentsForPullRequest page.
func (p *GetCommentsForPullRequestPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCommentsForPullRequestOutput, error) {
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

	result, err := p.client.GetCommentsForPullRequest(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCommentsForPullRequest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetCommentsForPullRequest",
	}
}

type opGetCommentsForPullRequestResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetCommentsForPullRequestResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetCommentsForPullRequestResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "codecommit"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "codecommit"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("codecommit")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addGetCommentsForPullRequestResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetCommentsForPullRequestResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
