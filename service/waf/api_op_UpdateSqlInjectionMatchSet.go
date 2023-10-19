// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Inserts or deletes SqlInjectionMatchTuple objects (filters) in
// a SqlInjectionMatchSet . For each SqlInjectionMatchTuple object, you specify
// the following values:
//   - Action : Whether to insert the object into or delete the object from the
//     array. To change a SqlInjectionMatchTuple , you delete the existing object and
//     add a new one.
//   - FieldToMatch : The part of web requests that you want AWS WAF to inspect
//     and, if you want AWS WAF to inspect a header or custom query parameter, the name
//     of the header or parameter.
//   - TextTransformation : Which text transformation, if any, to perform on the
//     web request before inspecting the request for snippets of malicious SQL code.
//     You can only specify a single type of TextTransformation.
//
// You use SqlInjectionMatchSet objects to specify which CloudFront requests that
// you want to allow, block, or count. For example, if you're receiving requests
// that contain snippets of SQL code in the query string and you want to block the
// requests, you can create a SqlInjectionMatchSet with the applicable settings,
// and then configure AWS WAF to block the requests. To create and configure a
// SqlInjectionMatchSet , perform the following steps:
//   - Submit a CreateSqlInjectionMatchSet request.
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of an UpdateIPSet request.
//   - Submit an UpdateSqlInjectionMatchSet request to specify the parts of web
//     requests that you want AWS WAF to inspect for snippets of SQL code.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/)
// .
func (c *Client) UpdateSqlInjectionMatchSet(ctx context.Context, params *UpdateSqlInjectionMatchSetInput, optFns ...func(*Options)) (*UpdateSqlInjectionMatchSetOutput, error) {
	if params == nil {
		params = &UpdateSqlInjectionMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSqlInjectionMatchSet", params, optFns, c.addOperationUpdateSqlInjectionMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSqlInjectionMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update a SqlInjectionMatchSet .
type UpdateSqlInjectionMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken .
	//
	// This member is required.
	ChangeToken *string

	// The SqlInjectionMatchSetId of the SqlInjectionMatchSet that you want to update.
	// SqlInjectionMatchSetId is returned by CreateSqlInjectionMatchSet and by
	// ListSqlInjectionMatchSets .
	//
	// This member is required.
	SqlInjectionMatchSetId *string

	// An array of SqlInjectionMatchSetUpdate objects that you want to insert into or
	// delete from a SqlInjectionMatchSet . For more information, see the applicable
	// data types:
	//   - SqlInjectionMatchSetUpdate : Contains Action and SqlInjectionMatchTuple
	//   - SqlInjectionMatchTuple : Contains FieldToMatch and TextTransformation
	//   - FieldToMatch : Contains Data and Type
	//
	// This member is required.
	Updates []types.SqlInjectionMatchSetUpdate

	noSmithyDocumentSerde
}

// The response to an UpdateSqlInjectionMatchSets request.
type UpdateSqlInjectionMatchSetOutput struct {

	// The ChangeToken that you used to submit the UpdateSqlInjectionMatchSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus .
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSqlInjectionMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSqlInjectionMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSqlInjectionMatchSet{}, middleware.After)
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
	if err = addOpUpdateSqlInjectionMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSqlInjectionMatchSet(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateSqlInjectionMatchSet",
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

func newServiceMetadataMiddleware_opUpdateSqlInjectionMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateSqlInjectionMatchSet",
	}
}
