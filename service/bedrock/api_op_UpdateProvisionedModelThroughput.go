// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a provisioned throughput. For more information, see Provisioned
// throughput (https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html)
// in the Bedrock User Guide.
func (c *Client) UpdateProvisionedModelThroughput(ctx context.Context, params *UpdateProvisionedModelThroughputInput, optFns ...func(*Options)) (*UpdateProvisionedModelThroughputOutput, error) {
	if params == nil {
		params = &UpdateProvisionedModelThroughputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProvisionedModelThroughput", params, optFns, c.addOperationUpdateProvisionedModelThroughputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProvisionedModelThroughputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProvisionedModelThroughputInput struct {

	// The ARN or name of the provisioned throughput to update.
	//
	// This member is required.
	ProvisionedModelId *string

	// The ARN of the new model to associate with this provisioned throughput.
	DesiredModelId *string

	// The new name for this provisioned throughput.
	DesiredProvisionedModelName *string

	noSmithyDocumentSerde
}

type UpdateProvisionedModelThroughputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProvisionedModelThroughputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProvisionedModelThroughput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProvisionedModelThroughput{}, middleware.After)
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
	if err = addOpUpdateProvisionedModelThroughputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProvisionedModelThroughput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProvisionedModelThroughput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "bedrock",
		OperationName: "UpdateProvisionedModelThroughput",
	}
}
