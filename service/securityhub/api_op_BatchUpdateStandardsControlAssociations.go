// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For a batch of security controls and standards, this operation updates the
// enablement status of a control in a standard.
func (c *Client) BatchUpdateStandardsControlAssociations(ctx context.Context, params *BatchUpdateStandardsControlAssociationsInput, optFns ...func(*Options)) (*BatchUpdateStandardsControlAssociationsOutput, error) {
	if params == nil {
		params = &BatchUpdateStandardsControlAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateStandardsControlAssociations", params, optFns, c.addOperationBatchUpdateStandardsControlAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateStandardsControlAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateStandardsControlAssociationsInput struct {

	// Updates the enablement status of a security control in a specified standard.
	//
	// This member is required.
	StandardsControlAssociationUpdates []types.StandardsControlAssociationUpdate

	noSmithyDocumentSerde
}

type BatchUpdateStandardsControlAssociationsOutput struct {

	// A security control (identified with SecurityControlId , SecurityControlArn , or
	// a mix of both parameters) whose enablement status in a specified standard
	// couldn't be updated.
	UnprocessedAssociationUpdates []types.UnprocessedStandardsControlAssociationUpdate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateStandardsControlAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateStandardsControlAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateStandardsControlAssociations{}, middleware.After)
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
	if err = addOpBatchUpdateStandardsControlAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateStandardsControlAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateStandardsControlAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchUpdateStandardsControlAssociations",
	}
}
