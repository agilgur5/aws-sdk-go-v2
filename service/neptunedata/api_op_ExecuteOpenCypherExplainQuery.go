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

// Executes an openCypher explain request. See The openCypher explain feature (https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher-explain.html)
// for more information. When invoking this operation in a Neptune cluster that has
// IAM authentication enabled, the IAM user or role making the request must have a
// policy attached that allows the neptune-db:ReadDataViaQuery (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery)
// IAM action in that cluster. Note that the neptune-db:QueryLanguage:Opencypher (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys)
// IAM condition key can be used in the policy document to restrict the use of
// openCypher queries (see Condition keys available in Neptune IAM data-access
// policy statements (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html)
// ).
func (c *Client) ExecuteOpenCypherExplainQuery(ctx context.Context, params *ExecuteOpenCypherExplainQueryInput, optFns ...func(*Options)) (*ExecuteOpenCypherExplainQueryOutput, error) {
	if params == nil {
		params = &ExecuteOpenCypherExplainQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteOpenCypherExplainQuery", params, optFns, c.addOperationExecuteOpenCypherExplainQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteOpenCypherExplainQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteOpenCypherExplainQueryInput struct {

	// The openCypher explain mode. Can be one of: static , dynamic , or details .
	//
	// This member is required.
	ExplainMode types.OpenCypherExplainMode

	// The openCypher query string.
	//
	// This member is required.
	OpenCypherQuery *string

	// The openCypher query parameters.
	Parameters *string

	noSmithyDocumentSerde
}

type ExecuteOpenCypherExplainQueryOutput struct {

	// A text blob containing the openCypher explain results.
	//
	// This member is required.
	Results []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteOpenCypherExplainQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExecuteOpenCypherExplainQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExecuteOpenCypherExplainQuery{}, middleware.After)
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
	if err = addOpExecuteOpenCypherExplainQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteOpenCypherExplainQuery(options.Region), middleware.Before); err != nil {
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
		operation: "ExecuteOpenCypherExplainQuery",
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

func newServiceMetadataMiddleware_opExecuteOpenCypherExplainQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "neptune-db",
		OperationName: "ExecuteOpenCypherExplainQuery",
	}
}
