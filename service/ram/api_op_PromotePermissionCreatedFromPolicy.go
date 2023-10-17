// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// When you attach a resource-based policy to a resource, RAM automatically
// creates a resource share of featureSet = CREATED_FROM_POLICY with a managed
// permission that has the same IAM permissions as the original resource-based
// policy. However, this type of managed permission is visible to only the resource
// share owner, and the associated resource share can't be modified by using RAM.
// This operation creates a separate, fully manageable customer managed permission
// that has the same IAM permissions as the original resource-based policy. You can
// associate this customer managed permission to any resource shares. Before you
// use PromoteResourceShareCreatedFromPolicy , you should first run this operation
// to ensure that you have an appropriate customer managed permission that can be
// associated with the promoted resource share.
//   - The original CREATED_FROM_POLICY policy isn't deleted, and resource shares
//     using that original policy aren't automatically updated.
//   - You can't modify a CREATED_FROM_POLICY resource share so you can't associate
//     the new customer managed permission by using ReplacePermsissionAssociations .
//     However, if you use PromoteResourceShareCreatedFromPolicy , that operation
//     automatically associates the fully manageable customer managed permission to the
//     newly promoted STANDARD resource share.
//   - After you promote a resource share, if the original CREATED_FROM_POLICY
//     managed permission has no other associations to A resource share, then RAM
//     automatically deletes it.
func (c *Client) PromotePermissionCreatedFromPolicy(ctx context.Context, params *PromotePermissionCreatedFromPolicyInput, optFns ...func(*Options)) (*PromotePermissionCreatedFromPolicyOutput, error) {
	if params == nil {
		params = &PromotePermissionCreatedFromPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PromotePermissionCreatedFromPolicy", params, optFns, c.addOperationPromotePermissionCreatedFromPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PromotePermissionCreatedFromPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PromotePermissionCreatedFromPolicyInput struct {

	// Specifies a name for the promoted customer managed permission.
	//
	// This member is required.
	Name *string

	// Specifies the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the CREATED_FROM_POLICY permission that you want to promote. You can get
	// this Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// by calling the ListResourceSharePermissions operation.
	//
	// This member is required.
	PermissionArn *string

	// Specifies a unique, case-sensitive identifier that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value. (https://wikipedia.org/wiki/Universally_unique_identifier)
	// . If you don't provide this value, then Amazon Web Services generates a random
	// one for you. If you retry the operation with the same ClientToken , but with
	// different parameters, the retry fails with an IdempotentParameterMismatch error.
	ClientToken *string

	noSmithyDocumentSerde
}

type PromotePermissionCreatedFromPolicyOutput struct {

	// The idempotency identifier associated with this request. If you want to repeat
	// the same operation in an idempotent manner then you must include this value in
	// the clientToken request parameter of that later call. All other parameters must
	// also have the same values that you used in the first call.
	ClientToken *string

	// Information about an RAM permission.
	Permission *types.ResourceSharePermissionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPromotePermissionCreatedFromPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPromotePermissionCreatedFromPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPromotePermissionCreatedFromPolicy{}, middleware.After)
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
	if err = addOpPromotePermissionCreatedFromPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPromotePermissionCreatedFromPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPromotePermissionCreatedFromPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "PromotePermissionCreatedFromPolicy",
	}
}
