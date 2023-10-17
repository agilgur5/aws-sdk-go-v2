// Code generated by smithy-go-codegen DO NOT EDIT.

package appfabric

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appfabric/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an ingestion destination, which specifies how an application's ingested
// data is processed by Amazon Web Services AppFabric and where it's delivered.
func (c *Client) UpdateIngestionDestination(ctx context.Context, params *UpdateIngestionDestinationInput, optFns ...func(*Options)) (*UpdateIngestionDestinationOutput, error) {
	if params == nil {
		params = &UpdateIngestionDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIngestionDestination", params, optFns, c.addOperationUpdateIngestionDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIngestionDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIngestionDestinationInput struct {

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app
	// bundle to use for the request.
	//
	// This member is required.
	AppBundleIdentifier *string

	// Contains information about the destination of ingested data.
	//
	// This member is required.
	DestinationConfiguration types.DestinationConfiguration

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the
	// ingestion destination to use for the request.
	//
	// This member is required.
	IngestionDestinationIdentifier *string

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the
	// ingestion to use for the request.
	//
	// This member is required.
	IngestionIdentifier *string

	noSmithyDocumentSerde
}

type UpdateIngestionDestinationOutput struct {

	// Contains information about an ingestion destination.
	//
	// This member is required.
	IngestionDestination *types.IngestionDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIngestionDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIngestionDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIngestionDestination{}, middleware.After)
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
	if err = addOpUpdateIngestionDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIngestionDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIngestionDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appfabric",
		OperationName: "UpdateIngestionDestination",
	}
}
