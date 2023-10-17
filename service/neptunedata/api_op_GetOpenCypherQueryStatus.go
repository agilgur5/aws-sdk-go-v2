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

// Retrieves the status of a specified openCypher query. When invoking this
// operation in a Neptune cluster that has IAM authentication enabled, the IAM user
// or role making the request must have a policy attached that allows the
// neptune-db:GetQueryStatus (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus)
// IAM action in that cluster. Note that the neptune-db:QueryLanguage:Opencypher (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys)
// IAM condition key can be used in the policy document to restrict the use of
// openCypher queries (see Condition keys available in Neptune IAM data-access
// policy statements (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html)
// ).
func (c *Client) GetOpenCypherQueryStatus(ctx context.Context, params *GetOpenCypherQueryStatusInput, optFns ...func(*Options)) (*GetOpenCypherQueryStatusOutput, error) {
	if params == nil {
		params = &GetOpenCypherQueryStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOpenCypherQueryStatus", params, optFns, c.addOperationGetOpenCypherQueryStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOpenCypherQueryStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOpenCypherQueryStatusInput struct {

	// The unique ID of the openCypher query for which to retrieve the query status.
	//
	// This member is required.
	QueryId *string

	noSmithyDocumentSerde
}

type GetOpenCypherQueryStatusOutput struct {

	// The openCypher query evaluation status.
	QueryEvalStats *types.QueryEvalStats

	// The unique ID of the query for which status is being returned.
	QueryId *string

	// The openCypher query string.
	QueryString *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOpenCypherQueryStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetOpenCypherQueryStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetOpenCypherQueryStatus{}, middleware.After)
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
	if err = addOpGetOpenCypherQueryStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpenCypherQueryStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOpenCypherQueryStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "neptune-db",
		OperationName: "GetOpenCypherQueryStatus",
	}
}
