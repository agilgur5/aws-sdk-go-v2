// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a data store, which is a repository for messages.
func (c *Client) CreateDatastore(ctx context.Context, params *CreateDatastoreInput, optFns ...func(*Options)) (*CreateDatastoreOutput, error) {
	if params == nil {
		params = &CreateDatastoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatastore", params, optFns, c.addOperationCreateDatastoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatastoreInput struct {

	// The name of the data store.
	//
	// This member is required.
	DatastoreName *string

	// Contains information about the partition dimensions in a data store.
	DatastorePartitions *types.DatastorePartitions

	// Where data in a data store is stored.. You can choose serviceManagedS3 storage,
	// customerManagedS3 storage, or iotSiteWiseMultiLayerStorage storage. The default
	// is serviceManagedS3 . You can't change the choice of Amazon S3 storage after
	// your data store is created.
	DatastoreStorage types.DatastoreStorage

	// Contains the configuration information of file formats. IoT Analytics data
	// stores support JSON and Parquet (https://parquet.apache.org/) . The default file
	// format is JSON. You can specify only one format. You can't change the file
	// format after you create the data store.
	FileFormatConfiguration *types.FileFormatConfiguration

	// How long, in days, message data is kept for the data store. When
	// customerManagedS3 storage is selected, this parameter is ignored.
	RetentionPeriod *types.RetentionPeriod

	// Metadata which can be used to manage the data store.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDatastoreOutput struct {

	// The ARN of the data store.
	DatastoreArn *string

	// The name of the data store.
	DatastoreName *string

	// How long, in days, message data is kept for the data store.
	RetentionPeriod *types.RetentionPeriod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatastoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDatastore{}, middleware.After)
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
	if err = addOpCreateDatastoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatastore(options.Region), middleware.Before); err != nil {
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
		operation: "CreateDatastore",
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

func newServiceMetadataMiddleware_opCreateDatastore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "CreateDatastore",
	}
}
