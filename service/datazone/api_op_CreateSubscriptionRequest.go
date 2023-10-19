// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a subscription request in Amazon DataZone.
func (c *Client) CreateSubscriptionRequest(ctx context.Context, params *CreateSubscriptionRequestInput, optFns ...func(*Options)) (*CreateSubscriptionRequestOutput, error) {
	if params == nil {
		params = &CreateSubscriptionRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubscriptionRequest", params, optFns, c.addOperationCreateSubscriptionRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubscriptionRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubscriptionRequestInput struct {

	// The ID of the Amazon DataZone domain in which the subscription request is
	// created.
	//
	// This member is required.
	DomainIdentifier *string

	// The reason for the subscription request.
	//
	// This member is required.
	RequestReason *string

	//
	//
	// This member is required.
	SubscribedListings []types.SubscribedListingInput

	// The Amazon DataZone principals for whom the subscription request is created.
	//
	// This member is required.
	SubscribedPrincipals []types.SubscribedPrincipalInput

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateSubscriptionRequestOutput struct {

	// A timestamp of when the subscription request is created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon DataZone user who created the subscription request.
	//
	// This member is required.
	CreatedBy *string

	// The ID of the Amazon DataZone domain in whcih the subscription request is
	// created.
	//
	// This member is required.
	DomainId *string

	// The ID of the subscription request.
	//
	// This member is required.
	Id *string

	// The reason for the subscription request.
	//
	// This member is required.
	RequestReason *string

	// The status of the subscription request.
	//
	// This member is required.
	Status types.SubscriptionRequestStatus

	//
	//
	// This member is required.
	SubscribedListings []types.SubscribedListing

	// The subscribed principals of the subscription request.
	//
	// This member is required.
	SubscribedPrincipals []types.SubscribedPrincipal

	// The timestamp of when the subscription request was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The decision comment of the subscription request.
	DecisionComment *string

	// The ID of the reviewer of the subscription request.
	ReviewerId *string

	// The Amazon DataZone user who updated the subscription request.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubscriptionRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSubscriptionRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSubscriptionRequest{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateSubscriptionRequestMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSubscriptionRequestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubscriptionRequest(options.Region), middleware.Before); err != nil {
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
		operation: "CreateSubscriptionRequest",
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

type idempotencyToken_initializeOpCreateSubscriptionRequest struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSubscriptionRequest) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSubscriptionRequest) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSubscriptionRequestInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSubscriptionRequestInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateSubscriptionRequestMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSubscriptionRequest{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSubscriptionRequest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "CreateSubscriptionRequest",
	}
}
