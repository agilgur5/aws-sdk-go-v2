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

// This operation applies only to Amazon Rekognition Custom Labels. Distributes
// the entries (images) in a training dataset across the training dataset and the
// test dataset for a project. DistributeDatasetEntries moves 20% of the training
// dataset images to the test dataset. An entry is a JSON Line that describes an
// image. You supply the Amazon Resource Names (ARN) of a project's training
// dataset and test dataset. The training dataset must contain the images that you
// want to split. The test dataset must be empty. The datasets must belong to the
// same project. To create training and test datasets for a project, call
// CreateDataset . Distributing a dataset takes a while to complete. To check the
// status call DescribeDataset . The operation is complete when the Status field
// for the training dataset and the test dataset is UPDATE_COMPLETE . If the
// dataset split fails, the value of Status is UPDATE_FAILED . This operation
// requires permissions to perform the rekognition:DistributeDatasetEntries action.
func (c *Client) DistributeDatasetEntries(ctx context.Context, params *DistributeDatasetEntriesInput, optFns ...func(*Options)) (*DistributeDatasetEntriesOutput, error) {
	if params == nil {
		params = &DistributeDatasetEntriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DistributeDatasetEntries", params, optFns, c.addOperationDistributeDatasetEntriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DistributeDatasetEntriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DistributeDatasetEntriesInput struct {

	// The ARNS for the training dataset and test dataset that you want to use. The
	// datasets must belong to the same project. The test dataset must be empty.
	//
	// This member is required.
	Datasets []types.DistributeDataset

	noSmithyDocumentSerde
}

type DistributeDatasetEntriesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDistributeDatasetEntriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDistributeDatasetEntries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDistributeDatasetEntries{}, middleware.After)
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
	if err = addOpDistributeDatasetEntriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDistributeDatasetEntries(options.Region), middleware.Before); err != nil {
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
		operation: "DistributeDatasetEntries",
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

func newServiceMetadataMiddleware_opDistributeDatasetEntries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DistributeDatasetEntries",
	}
}
