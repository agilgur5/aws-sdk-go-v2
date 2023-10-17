// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detects unsafe content in a specified JPEG or PNG format image. Use
// DetectModerationLabels to moderate images depending on your requirements. For
// example, you might want to filter images that contain nudity, but not images
// containing suggestive content. To filter images, use the labels returned by
// DetectModerationLabels to determine which types of content are appropriate. For
// information about moderation labels, see Detecting Unsafe Content in the Amazon
// Rekognition Developer Guide. You pass the input image either as base64-encoded
// image bytes or as a reference to an image in an Amazon S3 bucket. If you use the
// AWS CLI to call Amazon Rekognition operations, passing image bytes is not
// supported. The image must be either a PNG or JPEG formatted file. You can
// specify an adapter to use when retrieving label predictions by providing a
// ProjectVersionArn to the ProjectVersion argument.
func (c *Client) DetectModerationLabels(ctx context.Context, params *DetectModerationLabelsInput, optFns ...func(*Options)) (*DetectModerationLabelsOutput, error) {
	if params == nil {
		params = &DetectModerationLabelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectModerationLabels", params, optFns, c.addOperationDetectModerationLabelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectModerationLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectModerationLabelsInput struct {

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	Image *types.Image

	// Sets up the configuration for human evaluation, including the FlowDefinition
	// the image will be sent to.
	HumanLoopConfig *types.HumanLoopConfig

	// Specifies the minimum confidence level for the labels to return. Amazon
	// Rekognition doesn't return any labels with a confidence level lower than this
	// specified value. If you don't specify MinConfidence , the operation returns
	// labels with confidence values greater than or equal to 50 percent.
	MinConfidence *float32

	// Identifier for the custom adapter. Expects the ProjectVersionArn as a value.
	// Use the CreateProject or CreateProjectVersion APIs to create a custom adapter.
	ProjectVersion *string

	noSmithyDocumentSerde
}

type DetectModerationLabelsOutput struct {

	// Shows the results of the human in the loop evaluation.
	HumanLoopActivationOutput *types.HumanLoopActivationOutput

	// Array of detected Moderation labels and the time, in milliseconds from the
	// start of the video, they were detected.
	ModerationLabels []types.ModerationLabel

	// Version number of the base moderation detection model that was used to detect
	// unsafe content.
	ModerationModelVersion *string

	// Identifier of the custom adapter that was used during inference. If during
	// inference the adapter was EXPIRED, then the parameter will not be returned,
	// indicating that a base moderation detection project version was used.
	ProjectVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectModerationLabelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectModerationLabels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectModerationLabels{}, middleware.After)
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
	if err = addOpDetectModerationLabelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectModerationLabels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectModerationLabels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DetectModerationLabels",
	}
}
