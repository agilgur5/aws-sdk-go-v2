// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new role for your Amazon Web Services account. For more information
// about roles, see IAM roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
// in the IAM User Guide. For information about quotas for role names and the
// number of roles you can create, see IAM and STS quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html)
// in the IAM User Guide.
func (c *Client) CreateRole(ctx context.Context, params *CreateRoleInput, optFns ...func(*Options)) (*CreateRoleOutput, error) {
	if params == nil {
		params = &CreateRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRole", params, optFns, c.addOperationCreateRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRoleInput struct {

	// The trust relationship policy document that grants an entity permission to
	// assume the role. In IAM, you must provide a JSON policy that has been converted
	// to a string. However, for CloudFormation templates formatted in YAML, you can
	// provide the policy in JSON or YAML format. CloudFormation always converts a YAML
	// policy to JSON format before submitting it to IAM. The regex pattern (http://wikipedia.org/wiki/regex)
	// used to validate this parameter is a string of characters consisting of the
	// following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	// Upon success, the response includes the same trust policy in JSON format.
	//
	// This member is required.
	AssumeRolePolicyDocument *string

	// The name of the role to create. IAM user, group, role, and policy names must be
	// unique within the account. Names are not distinguished by case. For example, you
	// cannot create resources named both "MyResource" and "myresource". This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex) ) a string
	// of characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	RoleName *string

	// A description of the role.
	Description *string

	// The maximum session duration (in seconds) that you want to set for the
	// specified role. If you do not specify a value for this setting, the default
	// value of one hour is applied. This setting can have a value from 1 hour to 12
	// hours. Anyone who assumes the role from the CLI or API can use the
	// DurationSeconds API parameter or the duration-seconds CLI parameter to request
	// a longer session. The MaxSessionDuration setting determines the maximum
	// duration that can be requested using the DurationSeconds parameter. If users
	// don't specify a value for the DurationSeconds parameter, their security
	// credentials are valid for one hour by default. This applies when you use the
	// AssumeRole* API operations or the assume-role* CLI operations but does not
	// apply when you use those operations to create a console URL. For more
	// information, see Using IAM roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html)
	// in the IAM User Guide.
	MaxSessionDuration *int32

	// The path to the role. For more information about paths, see IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide. This parameter is optional. If it is not included, it
	// defaults to a slash (/). This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)
	// ) a string of characters consisting of either a forward slash (/) by itself or a
	// string that must begin and end with forward slashes. In addition, it can contain
	// any ASCII character from the ! ( \u0021 ) through the DEL character ( \u007F ),
	// including most punctuation characters, digits, and upper and lowercased letters.
	Path *string

	// The ARN of the managed policy that is used to set the permissions boundary for
	// the role. A permissions boundary policy defines the maximum permissions that
	// identity-based policies can grant to an entity, but does not grant permissions.
	// Permissions boundaries do not define the maximum permissions that a
	// resource-based policy can grant to an entity. To learn more, see Permissions
	// boundaries for IAM entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide. For more information about policy types, see Policy
	// types  (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policy-types)
	// in the IAM User Guide.
	PermissionsBoundary *string

	// A list of tags that you want to attach to the new role. Each tag consists of a
	// key name and an associated value. For more information about tagging, see
	// Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
	// in the IAM User Guide. If any one of the tags is invalid or if you exceed the
	// allowed maximum number of tags, then the entire request fails and the resource
	// is not created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the response to a successful CreateRole request.
type CreateRoleOutput struct {

	// A structure containing details about the new role.
	//
	// This member is required.
	Role *types.Role

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateRole{}, middleware.After)
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
	if err = addOpCreateRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRole(options.Region), middleware.Before); err != nil {
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
		operation: "CreateRole",
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

func newServiceMetadataMiddleware_opCreateRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreateRole",
	}
}
