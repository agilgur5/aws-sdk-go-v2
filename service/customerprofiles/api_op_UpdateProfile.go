// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the properties of a profile. The ProfileId is required for updating a
// customer profile. When calling the UpdateProfile API, specifying an empty string
// value means that any existing value will be removed. Not specifying a string
// value means that any value already there will be kept.
func (c *Client) UpdateProfile(ctx context.Context, params *UpdateProfileInput, optFns ...func(*Options)) (*UpdateProfileOutput, error) {
	if params == nil {
		params = &UpdateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProfile", params, optFns, c.addOperationUpdateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProfileInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The unique identifier of a customer profile.
	//
	// This member is required.
	ProfileId *string

	// A unique account number that you have given to the customer.
	AccountNumber *string

	// Any additional information relevant to the customer’s profile.
	AdditionalInformation *string

	// A generic address associated with the customer that is not mailing, shipping,
	// or billing.
	Address *types.UpdateAddress

	// A key value pair of attributes of a customer profile.
	Attributes map[string]string

	// The customer’s billing address.
	BillingAddress *types.UpdateAddress

	// The customer’s birth date.
	BirthDate *string

	// The customer’s business email address.
	BusinessEmailAddress *string

	// The name of the customer’s business.
	BusinessName *string

	// The customer’s business phone number.
	BusinessPhoneNumber *string

	// The customer’s email address, which has not been specified as a personal or
	// business address.
	EmailAddress *string

	// The customer’s first name.
	FirstName *string

	// The gender with which the customer identifies.
	//
	// Deprecated: This member has been deprecated.
	Gender types.Gender

	// An alternative to Gender which accepts any string as input.
	GenderString *string

	// The customer’s home phone number.
	HomePhoneNumber *string

	// The customer’s last name.
	LastName *string

	// The customer’s mailing address.
	MailingAddress *types.UpdateAddress

	// The customer’s middle name.
	MiddleName *string

	// The customer’s mobile phone number.
	MobilePhoneNumber *string

	// The type of profile used to describe the customer.
	//
	// Deprecated: This member has been deprecated.
	PartyType types.PartyType

	// An alternative to PartyType which accepts any string as input.
	PartyTypeString *string

	// The customer’s personal email address.
	PersonalEmailAddress *string

	// The customer’s phone number, which has not been specified as a mobile, home, or
	// business number.
	PhoneNumber *string

	// The customer’s shipping address.
	ShippingAddress *types.UpdateAddress

	noSmithyDocumentSerde
}

type UpdateProfileOutput struct {

	// The unique identifier of a customer profile.
	//
	// This member is required.
	ProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProfile{}, middleware.After)
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
	if err = addOpUpdateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "UpdateProfile",
	}
}
