// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an association between the source and the destination. A source can be
// associated with multiple destinations, and a destination can be associated with
// multiple sources. An association is a lineage tracking entity. For more
// information, see Amazon SageMaker ML Lineage Tracking (https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking.html)
// .
func (c *Client) AddAssociation(ctx context.Context, params *AddAssociationInput, optFns ...func(*Options)) (*AddAssociationOutput, error) {
	if params == nil {
		params = &AddAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddAssociation", params, optFns, c.addOperationAddAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddAssociationInput struct {

	// The Amazon Resource Name (ARN) of the destination.
	//
	// This member is required.
	DestinationArn *string

	// The ARN of the source.
	//
	// This member is required.
	SourceArn *string

	// The type of association. The following are suggested uses for each type. Amazon
	// SageMaker places no restrictions on their use.
	//   - ContributedTo - The source contributed to the destination or had a part in
	//   enabling the destination. For example, the training data contributed to the
	//   training job.
	//   - AssociatedWith - The source is connected to the destination. For example,
	//   an approval workflow is associated with a model deployment.
	//   - DerivedFrom - The destination is a modification of the source. For example,
	//   a digest output of a channel input for a processing job is derived from the
	//   original inputs.
	//   - Produced - The source generated the destination. For example, a training
	//   job produced a model artifact.
	AssociationType types.AssociationEdgeType

	noSmithyDocumentSerde
}

type AddAssociationOutput struct {

	// The Amazon Resource Name (ARN) of the destination.
	DestinationArn *string

	// The ARN of the source.
	SourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddAssociation{}, middleware.After)
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
	if err = addOpAddAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddAssociation(options.Region), middleware.Before); err != nil {
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
		operation: "AddAssociation",
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

func newServiceMetadataMiddleware_opAddAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "AddAssociation",
	}
}
