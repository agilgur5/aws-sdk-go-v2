// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a version of an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// , with a link to download the layer archive that's valid for 10 minutes.
func (c *Client) GetLayerVersionByArn(ctx context.Context, params *GetLayerVersionByArnInput, optFns ...func(*Options)) (*GetLayerVersionByArnOutput, error) {
	if params == nil {
		params = &GetLayerVersionByArnInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLayerVersionByArn", params, optFns, c.addOperationGetLayerVersionByArnMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLayerVersionByArnOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLayerVersionByArnInput struct {

	// The ARN of the layer version.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type GetLayerVersionByArnOutput struct {

	// A list of compatible instruction set architectures (https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html)
	// .
	CompatibleArchitectures []types.Architecture

	// The layer's compatible runtimes. The following list includes deprecated
	// runtimes. For more information, see Runtime deprecation policy (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy)
	// .
	CompatibleRuntimes []types.Runtime

	// Details about the layer version.
	Content *types.LayerVersionContentOutput

	// The date that the layer version was created, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	CreatedDate *string

	// The description of the version.
	Description *string

	// The ARN of the layer.
	LayerArn *string

	// The ARN of the layer version.
	LayerVersionArn *string

	// The layer's software license.
	LicenseInfo *string

	// The version number.
	Version int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLayerVersionByArnMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLayerVersionByArn{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLayerVersionByArn{}, middleware.After)
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
	if err = addOpGetLayerVersionByArnValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLayerVersionByArn(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLayerVersionByArn(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetLayerVersionByArn",
	}
}
