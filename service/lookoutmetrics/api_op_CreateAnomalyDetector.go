// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an anomaly detector.
func (c *Client) CreateAnomalyDetector(ctx context.Context, params *CreateAnomalyDetectorInput, optFns ...func(*Options)) (*CreateAnomalyDetectorOutput, error) {
	if params == nil {
		params = &CreateAnomalyDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAnomalyDetector", params, optFns, c.addOperationCreateAnomalyDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAnomalyDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAnomalyDetectorInput struct {

	// Contains information about the configuration of the anomaly detector.
	//
	// This member is required.
	AnomalyDetectorConfig *types.AnomalyDetectorConfig

	// The name of the detector.
	//
	// This member is required.
	AnomalyDetectorName *string

	// A description of the detector.
	AnomalyDetectorDescription *string

	// The ARN of the KMS key to use to encrypt your data.
	KmsKeyArn *string

	// A list of tags (https://docs.aws.amazon.com/lookoutmetrics/latest/dev/detectors-tags.html)
	// to apply to the anomaly detector.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAnomalyDetectorOutput struct {

	// The ARN of the detector.
	AnomalyDetectorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAnomalyDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAnomalyDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAnomalyDetector{}, middleware.After)
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
	if err = addOpCreateAnomalyDetectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAnomalyDetector(options.Region), middleware.Before); err != nil {
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
		operation: "CreateAnomalyDetector",
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

func newServiceMetadataMiddleware_opCreateAnomalyDetector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutmetrics",
		OperationName: "CreateAnomalyDetector",
	}
}
