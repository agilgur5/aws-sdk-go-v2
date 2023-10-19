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

// Sends an invitation to another account to join your organization as a member
// account. Organizations sends email on your behalf to the email address that is
// associated with the other account's owner. The invitation is implemented as a
// Handshake whose details are in the response.
//   - You can invite Amazon Web Services accounts only from the same seller as
//     the management account. For example, if your organization's management account
//     was created by Amazon Internet Services Pvt. Ltd (AISPL), an Amazon Web Services
//     seller in India, you can invite only other AISPL accounts to your organization.
//     You can't combine accounts from AISPL and Amazon Web Services or from any other
//     Amazon Web Services seller. For more information, see Consolidated billing in
//     India (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/useconsolidatedbilling-India.html)
//     .
//   - If you receive an exception that indicates that you exceeded your account
//     limits for the organization or that the operation failed because your
//     organization is still initializing, wait one hour and then try again. If the
//     error persists after an hour, contact Amazon Web Services Support (https://console.aws.amazon.com/support/home#/)
//     .
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission. This operation can be called only from the
// organization's management account.
func (c *Client) InviteAccountToOrganization(ctx context.Context, params *InviteAccountToOrganizationInput, optFns ...func(*Options)) (*InviteAccountToOrganizationOutput, error) {
	if params == nil {
		params = &InviteAccountToOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InviteAccountToOrganization", params, optFns, c.addOperationInviteAccountToOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InviteAccountToOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InviteAccountToOrganizationInput struct {

	// The identifier (ID) of the Amazon Web Services account that you want to invite
	// to join your organization. This is a JSON object that contains the following
	// elements: { "Type": "ACCOUNT", "Id": "< account id number >" } If you use the
	// CLI, you can submit this as a single string, similar to the following example:
	// --target Id=123456789012,Type=ACCOUNT If you specify "Type": "ACCOUNT" , you
	// must provide the Amazon Web Services account ID number as the Id . If you
	// specify "Type": "EMAIL" , you must specify the email address that is associated
	// with the account. --target Id=diego@example.com,Type=EMAIL
	//
	// This member is required.
	Target *types.HandshakeParty

	// Additional information that you want to include in the generated email to the
	// recipient account owner.
	Notes *string

	// A list of tags that you want to attach to the account when it becomes a member
	// of the organization. For each tag in the list, you must specify both a tag key
	// and a value. You can set the value to an empty string, but you can't set it to
	// null . For more information about tagging, see Tagging Organizations resources (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html)
	// in the Organizations User Guide. Any tags in the request are checked for
	// compliance with any applicable tag policies when the request is made. The
	// request is rejected if the tags in the request don't match the requirements of
	// the policy at that time. Tag policy compliance is not checked again when the
	// invitation is accepted and the tags are actually attached to the account. That
	// means that if the tag policy changes between the invitation and the acceptance,
	// then that tags could potentially be non-compliant. If any one of the tags is not
	// valid or if you exceed the allowed number of tags for an account, then the
	// entire request fails and invitations are not sent.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type InviteAccountToOrganizationOutput struct {

	// A structure that contains details about the handshake that is created to
	// support this invitation request.
	Handshake *types.Handshake

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInviteAccountToOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInviteAccountToOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInviteAccountToOrganization{}, middleware.After)
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
	if err = addOpInviteAccountToOrganizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInviteAccountToOrganization(options.Region), middleware.Before); err != nil {
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
		operation: "InviteAccountToOrganization",
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

func newServiceMetadataMiddleware_opInviteAccountToOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "InviteAccountToOrganization",
	}
}
