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

// Detects faces within an image that is provided as input. DetectFaces detects
// the 100 largest faces in the image. For each face detected, the operation
// returns face details. These details include a bounding box of the face, a
// confidence value (that the bounding box contains a face), and a fixed set of
// attributes such as facial landmarks (for example, coordinates of eye and mouth),
// pose, presence of facial occlusion, and so on. The face-detection algorithm is
// most effective on frontal faces. For non-frontal or obscured faces, the
// algorithm might not detect the faces or might detect faces with lower
// confidence. You pass the input image either as base64-encoded image bytes or as
// a reference to an image in an Amazon S3 bucket. If you use the AWS CLI to call
// Amazon Rekognition operations, passing image bytes is not supported. The image
// must be either a PNG or JPEG formatted file. This is a stateless API operation.
// That is, the operation does not persist any data. This operation requires
// permissions to perform the rekognition:DetectFaces action.
func (c *Client) DetectFaces(ctx context.Context, params *DetectFacesInput, optFns ...func(*Options)) (*DetectFacesOutput, error) {
	if params == nil {
		params = &DetectFacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectFaces", params, optFns, c.addOperationDetectFacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectFacesInput struct {

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	Image *types.Image

	// An array of facial attributes you want to be returned. A DEFAULT subset of
	// facial attributes - BoundingBox , Confidence , Pose , Quality , and Landmarks -
	// will always be returned. You can request for specific facial attributes (in
	// addition to the default list) - by using [ "DEFAULT", "FACE_OCCLUDED" ] or just [
	// "FACE_OCCLUDED" ]. You can request for all facial attributes by using [ "ALL"] .
	// Requesting more attributes may increase response time. If you provide both,
	// ["ALL", "DEFAULT"] , the service uses a logical "AND" operator to determine
	// which attributes to return (in this case, all attributes). Note that while the
	// FaceOccluded and EyeDirection attributes are supported when using DetectFaces ,
	// they aren't supported when analyzing videos with StartFaceDetection and
	// GetFaceDetection .
	Attributes []types.Attribute

	noSmithyDocumentSerde
}

func (*DetectFacesInput) operationName() string {
	return "DetectFaces"
}

type DetectFacesOutput struct {

	// Details of each face found in the image.
	FaceDetails []types.FaceDetail

	// The value of OrientationCorrection is always null. If the input image is in
	// .jpeg format, it might contain exchangeable image file format (Exif) metadata
	// that includes the image's orientation. Amazon Rekognition uses this orientation
	// information to perform image correction. The bounding box coordinates are
	// translated to represent object locations after the orientation information in
	// the Exif metadata is used to correct the image orientation. Images in .png
	// format don't contain Exif metadata. Amazon Rekognition doesn’t perform image
	// correction for images in .png format and .jpeg images without orientation
	// information in the image Exif metadata. The bounding box coordinates aren't
	// translated and represent the object locations before the image is rotated.
	OrientationCorrection types.OrientationCorrection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectFacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectFaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectFaces{}, middleware.After)
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
	if err = addDetectFacesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDetectFacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectFaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectFaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DetectFaces",
	}
}

type opDetectFacesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDetectFacesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDetectFacesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDetectFacesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDetectFacesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
