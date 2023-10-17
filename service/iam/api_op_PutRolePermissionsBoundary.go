// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates the policy that is specified as the IAM role's permissions
// boundary. You can use an Amazon Web Services managed policy or a customer
// managed policy to set the boundary for a role. Use the boundary to control the
// maximum permissions that the role can have. Setting a permissions boundary is an
// advanced feature that can affect the permissions for the role. You cannot set
// the boundary for a service-linked role. Policies used as permissions boundaries
// do not provide permissions. You must also attach a permissions policy to the
// role. To learn how the effective permissions for a role are evaluated, see IAM
// JSON policy evaluation logic (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
// in the IAM User Guide.
func (c *Client) PutRolePermissionsBoundary(ctx context.Context, params *PutRolePermissionsBoundaryInput, optFns ...func(*Options)) (*PutRolePermissionsBoundaryOutput, error) {
	if params == nil {
		params = &PutRolePermissionsBoundaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRolePermissionsBoundary", params, optFns, c.addOperationPutRolePermissionsBoundaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRolePermissionsBoundaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRolePermissionsBoundaryInput struct {

	// The ARN of the managed policy that is used to set the permissions boundary for
	// the role. A permissions boundary policy defines the maximum permissions that
	// identity-based policies can grant to an entity, but does not grant permissions.
	// Permissions boundaries do not define the maximum permissions that a
	// resource-based policy can grant to an entity. To learn more, see Permissions
	// boundaries for IAM entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide. For more information about policy types, see Policy
	// types  (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policy-types)
	// in the IAM User Guide.
	//
	// This member is required.
	PermissionsBoundary *string

	// The name (friendly name, not ARN) of the IAM role for which you want to set the
	// permissions boundary.
	//
	// This member is required.
	RoleName *string

	noSmithyDocumentSerde
}

type PutRolePermissionsBoundaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRolePermissionsBoundaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutRolePermissionsBoundary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutRolePermissionsBoundary{}, middleware.After)
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
	if err = addOpPutRolePermissionsBoundaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRolePermissionsBoundary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRolePermissionsBoundary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "PutRolePermissionsBoundary",
	}
}
