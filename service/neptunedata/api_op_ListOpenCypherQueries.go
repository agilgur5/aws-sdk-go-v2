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

// Lists active openCypher queries. See Neptune openCypher status endpoint (https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher-status.html)
// for more information. When invoking this operation in a Neptune cluster that has
// IAM authentication enabled, the IAM user or role making the request must have a
// policy attached that allows the neptune-db:GetQueryStatus (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus)
// IAM action in that cluster. Note that the neptune-db:QueryLanguage:Opencypher (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys)
// IAM condition key can be used in the policy document to restrict the use of
// openCypher queries (see Condition keys available in Neptune IAM data-access
// policy statements (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html)
// ).
func (c *Client) ListOpenCypherQueries(ctx context.Context, params *ListOpenCypherQueriesInput, optFns ...func(*Options)) (*ListOpenCypherQueriesOutput, error) {
	if params == nil {
		params = &ListOpenCypherQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOpenCypherQueries", params, optFns, c.addOperationListOpenCypherQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOpenCypherQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOpenCypherQueriesInput struct {

	// When set to TRUE and other parameters are not present, causes status
	// information to be returned for waiting queries as well as for running queries.
	IncludeWaiting *bool

	noSmithyDocumentSerde
}

type ListOpenCypherQueriesOutput struct {

	// The number of queries that have been accepted but not yet completed, including
	// queries in the queue.
	AcceptedQueryCount *int32

	// A list of current openCypher queries.
	Queries []types.GremlinQueryStatus

	// The number of currently running openCypher queries.
	RunningQueryCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOpenCypherQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOpenCypherQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOpenCypherQueries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOpenCypherQueries(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOpenCypherQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "neptune-db",
		OperationName: "ListOpenCypherQueries",
	}
}
