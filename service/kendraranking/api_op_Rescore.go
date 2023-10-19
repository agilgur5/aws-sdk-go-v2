// Code generated by smithy-go-codegen DO NOT EDIT.

package kendraranking

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendraranking/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Rescores or re-ranks search results from a search service such as OpenSearch
// (self managed). You use the semantic search capabilities of Amazon Kendra
// Intelligent Ranking to improve the search service's results.
func (c *Client) Rescore(ctx context.Context, params *RescoreInput, optFns ...func(*Options)) (*RescoreOutput, error) {
	if params == nil {
		params = &RescoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Rescore", params, optFns, c.addOperationRescoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RescoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RescoreInput struct {

	// The list of documents for Amazon Kendra Intelligent Ranking to rescore or rank
	// on.
	//
	// This member is required.
	Documents []types.Document

	// The identifier of the rescore execution plan. A rescore execution plan is an
	// Amazon Kendra Intelligent Ranking resource used for provisioning the Rescore
	// API.
	//
	// This member is required.
	RescoreExecutionPlanId *string

	// The input query from the search service.
	//
	// This member is required.
	SearchQuery *string

	noSmithyDocumentSerde
}

type RescoreOutput struct {

	// The identifier associated with the scores that Amazon Kendra Intelligent
	// Ranking gives to the results. Amazon Kendra Intelligent Ranking rescores or
	// re-ranks the results for the search service.
	RescoreId *string

	// A list of result items for documents with new relevancy scores. The results are
	// in descending order.
	ResultItems []types.RescoreResultItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRescoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRescore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRescore{}, middleware.After)
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
	if err = addOpRescoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRescore(options.Region), middleware.Before); err != nil {
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
		operation: "Rescore",
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

func newServiceMetadataMiddleware_opRescore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra-ranking",
		OperationName: "Rescore",
	}
}
