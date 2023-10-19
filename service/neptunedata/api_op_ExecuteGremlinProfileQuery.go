// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Executes a Gremlin Profile query, which runs a specified traversal, collects
// various metrics about the run, and produces a profile report as output. See
// Gremlin profile API in Neptune (https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-profile-api.html)
// for details. When invoking this operation in a Neptune cluster that has IAM
// authentication enabled, the IAM user or role making the request must have a
// policy attached that allows the neptune-db:ReadDataViaQuery (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery)
// IAM action in that cluster. Note that the neptune-db:QueryLanguage:Gremlin (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys)
// IAM condition key can be used in the policy document to restrict the use of
// Gremlin queries (see Condition keys available in Neptune IAM data-access policy
// statements (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html)
// ).
func (c *Client) ExecuteGremlinProfileQuery(ctx context.Context, params *ExecuteGremlinProfileQueryInput, optFns ...func(*Options)) (*ExecuteGremlinProfileQueryOutput, error) {
	if params == nil {
		params = &ExecuteGremlinProfileQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteGremlinProfileQuery", params, optFns, c.addOperationExecuteGremlinProfileQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteGremlinProfileQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteGremlinProfileQueryInput struct {

	// The Gremlin query string to profile.
	//
	// This member is required.
	GremlinQuery *string

	// If non-zero, causes the results string to be truncated at that number of
	// characters. If set to zero, the string contains all the results.
	Chop *int32

	// If this flag is set to TRUE , the results include a detailed report of all index
	// operations that took place during query execution and serialization.
	IndexOps *bool

	// If this flag is set to TRUE , the query results are gathered and displayed as
	// part of the profile report. If FALSE , only the result count is displayed.
	Results *bool

	// If non-null, the gathered results are returned in a serialized response message
	// in the format specified by this parameter. See Gremlin profile API in Neptune (https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-profile-api.html)
	// for more information.
	Serializer *string

	noSmithyDocumentSerde
}

type ExecuteGremlinProfileQueryOutput struct {

	// A text blob containing the Gremlin Profile result. See Gremlin profile API in
	// Neptune (https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-profile-api.html)
	// for details.
	//
	// This value conforms to the media type: text/plain
	Output []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteGremlinProfileQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExecuteGremlinProfileQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExecuteGremlinProfileQuery{}, middleware.After)
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
	if err = addOpExecuteGremlinProfileQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteGremlinProfileQuery(options.Region), middleware.Before); err != nil {
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
		operation: "ExecuteGremlinProfileQuery",
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

func newServiceMetadataMiddleware_opExecuteGremlinProfileQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "neptune-db",
		OperationName: "ExecuteGremlinProfileQuery",
	}
}
