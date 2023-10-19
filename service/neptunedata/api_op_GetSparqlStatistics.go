// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets RDF statistics (SPARQL).
func (c *Client) GetSparqlStatistics(ctx context.Context, params *GetSparqlStatisticsInput, optFns ...func(*Options)) (*GetSparqlStatisticsOutput, error) {
	if params == nil {
		params = &GetSparqlStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSparqlStatistics", params, optFns, c.addOperationGetSparqlStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSparqlStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSparqlStatisticsInput struct {
	noSmithyDocumentSerde
}

type GetSparqlStatisticsOutput struct {

	// Statistics for RDF data.
	//
	// This member is required.
	Payload *types.Statistics

	// The HTTP return code of the request. If the request succeeded, the code is 200.
	// See Common error codes for DFE statistics request (https://docs.aws.amazon.com/neptune/latest/userguide/neptune-dfe-statistics.html#neptune-dfe-statistics-errors)
	// for a list of common errors. When invoking this operation in a Neptune cluster
	// that has IAM authentication enabled, the IAM user or role making the request
	// must have a policy attached that allows the neptune-db:GetStatisticsStatus (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getstatisticsstatus)
	// IAM action in that cluster.
	//
	// This member is required.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSparqlStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSparqlStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSparqlStatistics{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSparqlStatistics(options.Region), middleware.Before); err != nil {
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
		operation: "GetSparqlStatistics",
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

func newServiceMetadataMiddleware_opGetSparqlStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "neptune-db",
		OperationName: "GetSparqlStatistics",
	}
}
