// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the text detection results of a Amazon Rekognition Video analysis started
// by StartTextDetection . Text detection with Amazon Rekognition Video is an
// asynchronous operation. You start text detection by calling StartTextDetection
// which returns a job identifier ( JobId ) When the text detection operation
// finishes, Amazon Rekognition publishes a completion status to the Amazon Simple
// Notification Service topic registered in the initial call to StartTextDetection
// . To get the results of the text detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . if so, call
// GetTextDetection and pass the job identifier ( JobId ) from the initial call of
// StartLabelDetection . GetTextDetection returns an array of detected text (
// TextDetections ) sorted by the time the text was detected, up to 100 words per
// frame of video. Each element of the array includes the detected text, the
// precentage confidence in the acuracy of the detected text, the time the text was
// detected, bounding box information for where the text was located, and unique
// identifiers for words and their lines. Use MaxResults parameter to limit the
// number of text detections returned. If there are more results than specified in
// MaxResults , the value of NextToken in the operation response contains a
// pagination token for getting the next set of results. To get the next page of
// results, call GetTextDetection and populate the NextToken request parameter
// with the token value returned from the previous call to GetTextDetection .
func (c *Client) GetTextDetection(ctx context.Context, params *GetTextDetectionInput, optFns ...func(*Options)) (*GetTextDetectionOutput, error) {
	if params == nil {
		params = &GetTextDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTextDetection", params, optFns, c.addOperationGetTextDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTextDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTextDetectionInput struct {

	// Job identifier for the text detection operation for which you want results
	// returned. You get the job identifer from an initial call to StartTextDetection .
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more labels to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of text.
	NextToken *string

	noSmithyDocumentSerde
}

type GetTextDetectionOutput struct {

	// Job identifier for the text detection operation for which you want to obtain
	// results. The job identifer is returned by an initial call to StartTextDetection.
	JobId *string

	// Current status of the text detection job.
	JobStatus types.VideoJobStatus

	// A job identifier specified in the call to StartTextDetection and returned in
	// the job completion notification sent to your Amazon Simple Notification Service
	// topic.
	JobTag *string

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of text.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// An array of text detected in the video. Each element contains the detected
	// text, the time in milliseconds from the start of the video that the text was
	// detected, and where it was detected on the screen.
	TextDetections []types.TextDetectionResult

	// Version number of the text detection model that was used to detect text.
	TextModelVersion *string

	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetection use Video to specify a video for
	// analysis. The supported file formats are .mp4, .mov and .avi.
	Video *types.Video

	// Information about a video that Amazon Rekognition analyzed. Videometadata is
	// returned in every page of paginated responses from a Amazon Rekognition video
	// operation.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTextDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTextDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTextDetection{}, middleware.After)
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
	if err = addOpGetTextDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTextDetection(options.Region), middleware.Before); err != nil {
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
		operation: "GetTextDetection",
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

// GetTextDetectionAPIClient is a client that implements the GetTextDetection
// operation.
type GetTextDetectionAPIClient interface {
	GetTextDetection(context.Context, *GetTextDetectionInput, ...func(*Options)) (*GetTextDetectionOutput, error)
}

var _ GetTextDetectionAPIClient = (*Client)(nil)

// GetTextDetectionPaginatorOptions is the paginator options for GetTextDetection
type GetTextDetectionPaginatorOptions struct {
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTextDetectionPaginator is a paginator for GetTextDetection
type GetTextDetectionPaginator struct {
	options   GetTextDetectionPaginatorOptions
	client    GetTextDetectionAPIClient
	params    *GetTextDetectionInput
	nextToken *string
	firstPage bool
}

// NewGetTextDetectionPaginator returns a new GetTextDetectionPaginator
func NewGetTextDetectionPaginator(client GetTextDetectionAPIClient, params *GetTextDetectionInput, optFns ...func(*GetTextDetectionPaginatorOptions)) *GetTextDetectionPaginator {
	if params == nil {
		params = &GetTextDetectionInput{}
	}

	options := GetTextDetectionPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTextDetectionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTextDetectionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTextDetection page.
func (p *GetTextDetectionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTextDetectionOutput, error) {
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

	result, err := p.client.GetTextDetection(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTextDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetTextDetection",
	}
}
