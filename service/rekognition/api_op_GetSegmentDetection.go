// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the segment detection results of a Amazon Rekognition Video analysis
// started by StartSegmentDetection . Segment detection with Amazon Rekognition
// Video is an asynchronous operation. You start segment detection by calling
// StartSegmentDetection which returns a job identifier ( JobId ). When the segment
// detection operation finishes, Amazon Rekognition publishes a completion status
// to the Amazon Simple Notification Service topic registered in the initial call
// to StartSegmentDetection . To get the results of the segment detection
// operation, first check that the status value published to the Amazon SNS topic
// is SUCCEEDED . if so, call GetSegmentDetection and pass the job identifier (
// JobId ) from the initial call of StartSegmentDetection . GetSegmentDetection
// returns detected segments in an array ( Segments ) of SegmentDetection objects.
// Segments is sorted by the segment types specified in the SegmentTypes input
// parameter of StartSegmentDetection . Each element of the array includes the
// detected segment, the precentage confidence in the acuracy of the detected
// segment, the type of the segment, and the frame in which the segment was
// detected. Use SelectedSegmentTypes to find out the type of segment detection
// requested in the call to StartSegmentDetection . Use the MaxResults parameter
// to limit the number of segment detections returned. If there are more results
// than specified in MaxResults , the value of NextToken in the operation response
// contains a pagination token for getting the next set of results. To get the next
// page of results, call GetSegmentDetection and populate the NextToken request
// parameter with the token value returned from the previous call to
// GetSegmentDetection . For more information, see Detecting video segments in
// stored video in the Amazon Rekognition Developer Guide.
func (c *Client) GetSegmentDetection(ctx context.Context, params *GetSegmentDetectionInput, optFns ...func(*Options)) (*GetSegmentDetectionOutput, error) {
	if params == nil {
		params = &GetSegmentDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSegmentDetection", params, optFns, c.addOperationGetSegmentDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSegmentDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSegmentDetectionInput struct {

	// Job identifier for the text detection operation for which you want results
	// returned. You get the job identifer from an initial call to
	// StartSegmentDetection .
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000.
	MaxResults *int32

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of text.
	NextToken *string

	noSmithyDocumentSerde
}

func (*GetSegmentDetectionInput) operationName() string {
	return "GetSegmentDetection"
}

type GetSegmentDetectionOutput struct {

	// An array of objects. There can be multiple audio streams. Each AudioMetadata
	// object contains metadata for a single audio stream. Audio information in an
	// AudioMetadata objects includes the audio codec, the number of audio channels,
	// the duration of the audio stream, and the sample rate. Audio metadata is
	// returned in each page of information returned by GetSegmentDetection .
	AudioMetadata []types.AudioMetadata

	// Job identifier for the segment detection operation for which you want to obtain
	// results. The job identifer is returned by an initial call to
	// StartSegmentDetection.
	JobId *string

	// Current status of the segment detection job.
	JobStatus types.VideoJobStatus

	// A job identifier specified in the call to StartSegmentDetection and returned in
	// the job completion notification sent to your Amazon Simple Notification Service
	// topic.
	JobTag *string

	// If the previous response was incomplete (because there are more labels to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of text.
	NextToken *string

	// An array of segments detected in a video. The array is sorted by the segment
	// types (TECHNICAL_CUE or SHOT) specified in the SegmentTypes input parameter of
	// StartSegmentDetection . Within each segment type the array is sorted by
	// timestamp values.
	Segments []types.SegmentDetection

	// An array containing the segment types requested in the call to
	// StartSegmentDetection .
	SelectedSegmentTypes []types.SegmentTypeInfo

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetection use Video to specify a video for
	// analysis. The supported file formats are .mp4, .mov and .avi.
	Video *types.Video

	// Currently, Amazon Rekognition Video returns a single object in the VideoMetadata
	// array. The object contains information about the video stream in the input file
	// that Amazon Rekognition Video chose to analyze. The VideoMetadata object
	// includes the video codec, video format and other information. Video metadata is
	// returned in each page of information returned by GetSegmentDetection .
	VideoMetadata []types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSegmentDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSegmentDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSegmentDetection{}, middleware.After)
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
	if err = addGetSegmentDetectionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetSegmentDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSegmentDetection(options.Region), middleware.Before); err != nil {
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

// GetSegmentDetectionAPIClient is a client that implements the
// GetSegmentDetection operation.
type GetSegmentDetectionAPIClient interface {
	GetSegmentDetection(context.Context, *GetSegmentDetectionInput, ...func(*Options)) (*GetSegmentDetectionOutput, error)
}

var _ GetSegmentDetectionAPIClient = (*Client)(nil)

// GetSegmentDetectionPaginatorOptions is the paginator options for
// GetSegmentDetection
type GetSegmentDetectionPaginatorOptions struct {
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSegmentDetectionPaginator is a paginator for GetSegmentDetection
type GetSegmentDetectionPaginator struct {
	options   GetSegmentDetectionPaginatorOptions
	client    GetSegmentDetectionAPIClient
	params    *GetSegmentDetectionInput
	nextToken *string
	firstPage bool
}

// NewGetSegmentDetectionPaginator returns a new GetSegmentDetectionPaginator
func NewGetSegmentDetectionPaginator(client GetSegmentDetectionAPIClient, params *GetSegmentDetectionInput, optFns ...func(*GetSegmentDetectionPaginatorOptions)) *GetSegmentDetectionPaginator {
	if params == nil {
		params = &GetSegmentDetectionInput{}
	}

	options := GetSegmentDetectionPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSegmentDetectionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSegmentDetectionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetSegmentDetection page.
func (p *GetSegmentDetectionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSegmentDetectionOutput, error) {
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

	result, err := p.client.GetSegmentDetection(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSegmentDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetSegmentDetection",
	}
}

type opGetSegmentDetectionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetSegmentDetectionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetSegmentDetectionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "rekognition"
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
				signingName = "rekognition"
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
				v4aScheme.SigningName = aws.String("rekognition")
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

func addGetSegmentDetectionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetSegmentDetectionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
