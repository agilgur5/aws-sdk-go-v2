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

// Gets the subscription grant in Amazon DataZone.
func (c *Client) GetSubscriptionGrant(ctx context.Context, params *GetSubscriptionGrantInput, optFns ...func(*Options)) (*GetSubscriptionGrantOutput, error) {
	if params == nil {
		params = &GetSubscriptionGrantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSubscriptionGrant", params, optFns, c.addOperationGetSubscriptionGrantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSubscriptionGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSubscriptionGrantInput struct {

	// The ID of the Amazon DataZone domain in which the subscription grant exists.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the subscription grant.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetSubscriptionGrantOutput struct {

	// The timestamp of when the subscription grant is created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon DataZone user who created the subscription grant.
	//
	// This member is required.
	CreatedBy *string

	// The ID of the Amazon DataZone domain in which the subscription grant exists.
	//
	// This member is required.
	DomainId *string

	// The entity to which the subscription is granted.
	//
	// This member is required.
	GrantedEntity types.GrantedEntity

	// The ID of the subscription grant.
	//
	// This member is required.
	Id *string

	// The status of the subscription grant.
	//
	// This member is required.
	Status types.SubscriptionGrantOverallStatus

	// The subscription target ID associated with the subscription grant.
	//
	// This member is required.
	SubscriptionTargetId *string

	// The timestamp of when the subscription grant was upated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The assets for which the subscription grant is created.
	Assets []types.SubscribedAsset

	// The identifier of the subscription.
	SubscriptionId *string

	// The Amazon DataZone user who updated the subscription grant.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSubscriptionGrantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSubscriptionGrant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSubscriptionGrant{}, middleware.After)
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
	if err = addOpGetSubscriptionGrantValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSubscriptionGrant(options.Region), middleware.Before); err != nil {
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
		operation: "GetSubscriptionGrant",
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

func newServiceMetadataMiddleware_opGetSubscriptionGrant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "GetSubscriptionGrant",
	}
}
