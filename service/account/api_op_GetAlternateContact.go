// Code generated by smithy-go-codegen DO NOT EDIT.

package account

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the specified alternate contact attached to an Amazon Web Services
// account. For complete details about how to use the alternate contact operations,
// see Access or updating the alternate contacts (https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact.html)
// . Before you can update the alternate contact information for an Amazon Web
// Services account that is managed by Organizations, you must first enable
// integration between Amazon Web Services Account Management and Organizations.
// For more information, see Enabling trusted access for Amazon Web Services
// Account Management (https://docs.aws.amazon.com/accounts/latest/reference/using-orgs-trusted-access.html)
// .
func (c *Client) GetAlternateContact(ctx context.Context, params *GetAlternateContactInput, optFns ...func(*Options)) (*GetAlternateContactOutput, error) {
	if params == nil {
		params = &GetAlternateContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAlternateContact", params, optFns, c.addOperationGetAlternateContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAlternateContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAlternateContactInput struct {

	// Specifies which alternate contact you want to retrieve.
	//
	// This member is required.
	AlternateContactType types.AlternateContactType

	// Specifies the 12 digit account ID number of the Amazon Web Services account
	// that you want to access or modify with this operation. If you do not specify
	// this parameter, it defaults to the Amazon Web Services account of the identity
	// used to call the operation. To use this parameter, the caller must be an
	// identity in the organization's management account (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account)
	// or a delegated administrator account, and the specified account ID must be a
	// member account in the same organization. The organization must have all
	// features enabled (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
	// , and the organization must have trusted access (https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html)
	// enabled for the Account Management service, and optionally a delegated admin (https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html)
	// account assigned. The management account can't specify its own AccountId ; it
	// must call the operation in standalone context by not including the AccountId
	// parameter. To call this operation on an account that is not a member of an
	// organization, then don't specify this parameter, and call the operation using an
	// identity belonging to the account whose contacts you wish to retrieve or modify.
	AccountId *string

	noSmithyDocumentSerde
}

type GetAlternateContactOutput struct {

	// A structure that contains the details for the specified alternate contact.
	AlternateContact *types.AlternateContact

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAlternateContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAlternateContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAlternateContact{}, middleware.After)
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
	if err = addOpGetAlternateContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAlternateContact(options.Region), middleware.Before); err != nil {
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
		operation: "GetAlternateContact",
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

func newServiceMetadataMiddleware_opGetAlternateContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "account",
		OperationName: "GetAlternateContact",
	}
}
