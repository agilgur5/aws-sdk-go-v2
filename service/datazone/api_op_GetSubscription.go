// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a subscription in Amazon DataZone.
func (c *Client) GetSubscription(ctx context.Context, params *GetSubscriptionInput, optFns ...func(*Options)) (*GetSubscriptionOutput, error) {
	if params == nil {
		params = &GetSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSubscription", params, optFns, c.addOperationGetSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSubscriptionInput struct {

	// The ID of the Amazon DataZone domain in which the subscription exists.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the subscription.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetSubscriptionOutput struct {

	// The timestamp of when the subscription was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon DataZone user who created the subscription.
	//
	// This member is required.
	CreatedBy *string

	// The ID of the Amazon DataZone domain in which the subscription exists.
	//
	// This member is required.
	DomainId *string

	// The ID of the subscription.
	//
	// This member is required.
	Id *string

	// The status of the subscription.
	//
	// This member is required.
	Status types.SubscriptionStatus

	//
	//
	// This member is required.
	SubscribedListing *types.SubscribedListing

	// The principal that owns the subscription.
	//
	// This member is required.
	SubscribedPrincipal types.SubscribedPrincipal

	// The timestamp of when the subscription was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The retain permissions of the subscription.
	RetainPermissions *bool

	// The ID of the subscription request.
	SubscriptionRequestId *string

	// The Amazon DataZone user who updated the subscription.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSubscription{}, middleware.After)
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
	if err = addOpGetSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "GetSubscription",
	}
}
