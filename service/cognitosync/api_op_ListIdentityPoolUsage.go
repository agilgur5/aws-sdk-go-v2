// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of identity pools registered with Cognito. ListIdentityPoolUsage
// can only be called with developer credentials. You cannot make this API call
// with the temporary user credentials provided by Cognito Identity.
// ListIdentityPoolUsage The following examples have been edited for readability.
// POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 9be7c425-ef05-48c0-aef3-9f0ff2fe17d3 X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.ListIdentityPoolUsage
// HOST: cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T211414Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#ListIdentityPoolUsage", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "MaxResults": "2" } } 1.1 200 OK x-amzn-requestid:
// 9be7c425-ef05-48c0-aef3-9f0ff2fe17d3 content-type: application/json
// content-length: 519 date: Tue, 11 Nov 2014 21:14:14 GMT { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#ListIdentityPoolUsageResponse", "Count": 2,
// "IdentityPoolUsages": [ { "DataStorage": 0, "IdentityPoolId":
// "IDENTITY_POOL_ID", "LastModifiedDate": 1.413836234607E9, "SyncSessionsCount":
// null }, { "DataStorage": 0, "IdentityPoolId": "IDENTITY_POOL_ID",
// "LastModifiedDate": 1.410892165601E9, "SyncSessionsCount": null }],
// "MaxResults": 2, "NextToken":
// "dXMtZWFzdC0xOjBjMWJhMDUyLWUwOTgtNDFmYS1hNzZlLWVhYTJjMTI1Zjg2MQ==" }, "Version":
// "1.0" }
func (c *Client) ListIdentityPoolUsage(ctx context.Context, params *ListIdentityPoolUsageInput, optFns ...func(*Options)) (*ListIdentityPoolUsageOutput, error) {
	if params == nil {
		params = &ListIdentityPoolUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentityPoolUsage", params, optFns, c.addOperationListIdentityPoolUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentityPoolUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for usage information on an identity pool.
type ListIdentityPoolUsageInput struct {

	// The maximum number of results to be returned.
	MaxResults *int32

	// A pagination token for obtaining the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Returned for a successful ListIdentityPoolUsage request.
type ListIdentityPoolUsageOutput struct {

	// Total number of identities for the identity pool.
	Count int32

	// Usage information for the identity pools.
	IdentityPoolUsages []types.IdentityPoolUsage

	// The maximum number of results to be returned.
	MaxResults int32

	// A pagination token for obtaining the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIdentityPoolUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIdentityPoolUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIdentityPoolUsage{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentityPoolUsage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListIdentityPoolUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "ListIdentityPoolUsage",
	}
}
