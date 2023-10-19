// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the contents of the effective policy for specified policy type and
// account. The effective policy is the aggregation of any policies of the
// specified type that the account inherits, plus any policy of that type that is
// directly attached to the account. This operation applies only to policy types
// other than service control policies (SCPs). For more information about policy
// inheritance, see Understanding management policy inheritance (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_inheritance_mgmt.html)
// in the Organizations User Guide. This operation can be called from any account
// in the organization.
func (c *Client) DescribeEffectivePolicy(ctx context.Context, params *DescribeEffectivePolicyInput, optFns ...func(*Options)) (*DescribeEffectivePolicyOutput, error) {
	if params == nil {
		params = &DescribeEffectivePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEffectivePolicy", params, optFns, c.addOperationDescribeEffectivePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEffectivePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEffectivePolicyInput struct {

	// The type of policy that you want information about. You can specify one of the
	// following values:
	//   - AISERVICES_OPT_OUT_POLICY (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	//   - BACKUP_POLICY (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
	//   - TAG_POLICY (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
	//
	// This member is required.
	PolicyType types.EffectivePolicyType

	// When you're signed in as the management account, specify the ID of the account
	// that you want details about. Specifying an organization root or organizational
	// unit (OU) as the target is not supported.
	TargetId *string

	noSmithyDocumentSerde
}

type DescribeEffectivePolicyOutput struct {

	// The contents of the effective policy.
	EffectivePolicy *types.EffectivePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEffectivePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEffectivePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEffectivePolicy{}, middleware.After)
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
	if err = addOpDescribeEffectivePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEffectivePolicy(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeEffectivePolicy",
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

func newServiceMetadataMiddleware_opDescribeEffectivePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeEffectivePolicy",
	}
}
