// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a member account from its parent organization. This version of the
// operation is performed by the account that wants to leave. To remove a member
// account as a user in the management account, use RemoveAccountFromOrganization
// instead. This operation can be called only from a member account in the
// organization.
//   - The management account in an organization with all features enabled can set
//     service control policies (SCPs) that can restrict what administrators of member
//     accounts can do. This includes preventing them from successfully calling
//     LeaveOrganization and leaving the organization.
//   - You can leave an organization as a member account only if the account is
//     configured with the information required to operate as a standalone account.
//     When you create an account in an organization using the Organizations console,
//     API, or CLI commands, the information required of standalone accounts is not
//     automatically collected. For each account that you want to make standalone, you
//     must perform the following steps. If any of the steps are already completed for
//     this account, that step doesn't appear.
//   - Choose a support plan
//   - Provide and verify the required contact information
//   - Provide a current payment method Amazon Web Services uses the payment
//     method to charge for any billable (not free tier) Amazon Web Services activity
//     that occurs while the account isn't attached to an organization. For more
//     information, see Considerations before removing an account from an
//     organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html)
//     in the Organizations User Guide.
//   - The account that you want to leave must not be a delegated administrator
//     account for any Amazon Web Services service enabled for your organization. If
//     the account is a delegated administrator, you must first change the delegated
//     administrator account to another account that is remaining in the organization.
//   - You can leave an organization only after you enable IAM user access to
//     billing in your account. For more information, see About IAM access to the
//     Billing and Cost Management console (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html#ControllingAccessWebsite-Activate)
//     in the Amazon Web Services Billing and Cost Management User Guide.
//   - After the account leaves the organization, all tags that were attached to
//     the account object in the organization are deleted. Amazon Web Services accounts
//     outside of an organization do not support tags.
//   - A newly created account has a waiting period before it can be removed from
//     its organization. If you get an error that indicates that a wait period is
//     required, then try again in a few days.
//   - If you are using an organization principal to call LeaveOrganization across
//     multiple accounts, you can only do this up to 5 accounts per second in a single
//     organization.
func (c *Client) LeaveOrganization(ctx context.Context, params *LeaveOrganizationInput, optFns ...func(*Options)) (*LeaveOrganizationOutput, error) {
	if params == nil {
		params = &LeaveOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "LeaveOrganization", params, optFns, c.addOperationLeaveOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*LeaveOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type LeaveOrganizationInput struct {
	noSmithyDocumentSerde
}

type LeaveOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationLeaveOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpLeaveOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpLeaveOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opLeaveOrganization(options.Region), middleware.Before); err != nil {
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
		operation: "LeaveOrganization",
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

func newServiceMetadataMiddleware_opLeaveOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "LeaveOrganization",
	}
}
