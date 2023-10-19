// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about uploads, given an AWS Device Farm project ARN.
func (c *Client) ListUploads(ctx context.Context, params *ListUploadsInput, optFns ...func(*Options)) (*ListUploadsOutput, error) {
	if params == nil {
		params = &ListUploadsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUploads", params, optFns, c.addOperationListUploadsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUploadsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the list uploads operation.
type ListUploadsInput struct {

	// The Amazon Resource Name (ARN) of the project for which you want to list
	// uploads.
	//
	// This member is required.
	Arn *string

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The type of upload. Must be one of the following values:
	//   - ANDROID_APP
	//   - IOS_APP
	//   - WEB_APP
	//   - EXTERNAL_DATA
	//   - APPIUM_JAVA_JUNIT_TEST_PACKAGE
	//   - APPIUM_JAVA_TESTNG_TEST_PACKAGE
	//   - APPIUM_PYTHON_TEST_PACKAGE
	//   - APPIUM_NODE_TEST_PACKAGE
	//   - APPIUM_RUBY_TEST_PACKAGE
	//   - APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE
	//   - APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE
	//   - APPIUM_WEB_PYTHON_TEST_PACKAGE
	//   - APPIUM_WEB_NODE_TEST_PACKAGE
	//   - APPIUM_WEB_RUBY_TEST_PACKAGE
	//   - CALABASH_TEST_PACKAGE
	//   - INSTRUMENTATION_TEST_PACKAGE
	//   - UIAUTOMATION_TEST_PACKAGE
	//   - UIAUTOMATOR_TEST_PACKAGE
	//   - XCTEST_TEST_PACKAGE
	//   - XCTEST_UI_TEST_PACKAGE
	//   - APPIUM_JAVA_JUNIT_TEST_SPEC
	//   - APPIUM_JAVA_TESTNG_TEST_SPEC
	//   - APPIUM_PYTHON_TEST_SPEC
	//   - APPIUM_NODE_TEST_SPEC
	//   - APPIUM_RUBY_TEST_SPEC
	//   - APPIUM_WEB_JAVA_JUNIT_TEST_SPEC
	//   - APPIUM_WEB_JAVA_TESTNG_TEST_SPEC
	//   - APPIUM_WEB_PYTHON_TEST_SPEC
	//   - APPIUM_WEB_NODE_TEST_SPEC
	//   - APPIUM_WEB_RUBY_TEST_SPEC
	//   - INSTRUMENTATION_TEST_SPEC
	//   - XCTEST_UI_TEST_SPEC
	Type types.UploadType

	noSmithyDocumentSerde
}

// Represents the result of a list uploads request.
type ListUploadsOutput struct {

	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Information about the uploads.
	Uploads []types.Upload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUploadsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUploads{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUploads{}, middleware.After)
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
	if err = addOpListUploadsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUploads(options.Region), middleware.Before); err != nil {
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
		operation: "ListUploads",
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

// ListUploadsAPIClient is a client that implements the ListUploads operation.
type ListUploadsAPIClient interface {
	ListUploads(context.Context, *ListUploadsInput, ...func(*Options)) (*ListUploadsOutput, error)
}

var _ ListUploadsAPIClient = (*Client)(nil)

// ListUploadsPaginatorOptions is the paginator options for ListUploads
type ListUploadsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUploadsPaginator is a paginator for ListUploads
type ListUploadsPaginator struct {
	options   ListUploadsPaginatorOptions
	client    ListUploadsAPIClient
	params    *ListUploadsInput
	nextToken *string
	firstPage bool
}

// NewListUploadsPaginator returns a new ListUploadsPaginator
func NewListUploadsPaginator(client ListUploadsAPIClient, params *ListUploadsInput, optFns ...func(*ListUploadsPaginatorOptions)) *ListUploadsPaginator {
	if params == nil {
		params = &ListUploadsInput{}
	}

	options := ListUploadsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUploadsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUploadsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUploads page.
func (p *ListUploadsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUploadsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListUploads(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUploads(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListUploads",
	}
}
