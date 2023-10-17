// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssmincidents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified response plan.
func (c *Client) UpdateResponsePlan(ctx context.Context, params *UpdateResponsePlanInput, optFns ...func(*Options)) (*UpdateResponsePlanOutput, error) {
	if params == nil {
		params = &UpdateResponsePlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResponsePlan", params, optFns, c.addOperationUpdateResponsePlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResponsePlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResponsePlanInput struct {

	// The Amazon Resource Name (ARN) of the response plan.
	//
	// This member is required.
	Arn *string

	// The actions that this response plan takes at the beginning of an incident.
	Actions []types.Action

	// The Chatbot chat channel used for collaboration during an incident. Use the
	// empty structure to remove the chat channel from the response plan.
	ChatChannel types.ChatChannel

	// A token ensuring that the operation is called only once with the specified
	// details.
	ClientToken *string

	// The long format name of the response plan. The display name can't contain
	// spaces.
	DisplayName *string

	// The Amazon Resource Name (ARN) for the contacts and escalation plans that the
	// response plan engages during an incident.
	Engagements []string

	// The string Incident Manager uses to prevent duplicate incidents from being
	// created by the same incident in the same account.
	IncidentTemplateDedupeString *string

	// Defines the impact to the customers. Providing an impact overwrites the impact
	// provided by a response plan. Possible impacts:
	//   - 5 - Severe impact
	//   - 4 - High impact
	//   - 3 - Medium impact
	//   - 2 - Low impact
	//   - 1 - No impact
	IncidentTemplateImpact *int32

	// The Amazon SNS targets that are notified when updates are made to an incident.
	IncidentTemplateNotificationTargets []types.NotificationTargetItem

	// A brief summary of the incident. This typically contains what has happened,
	// what's currently happening, and next steps.
	IncidentTemplateSummary *string

	// Tags to assign to the template. When the StartIncident API action is called,
	// Incident Manager assigns the tags specified in the template to the incident. To
	// call this action, you must also have permission to call the TagResource API
	// action for the incident record resource.
	IncidentTemplateTags map[string]string

	// The short format name of the incident. The title can't contain spaces.
	IncidentTemplateTitle *string

	// Information about third-party services integrated into the response plan.
	Integrations []types.Integration

	noSmithyDocumentSerde
}

type UpdateResponsePlanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResponsePlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResponsePlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResponsePlan{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateResponsePlanMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateResponsePlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResponsePlan(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateResponsePlan struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateResponsePlan) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateResponsePlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateResponsePlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateResponsePlanInput ")
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
func addIdempotencyToken_opUpdateResponsePlanMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateResponsePlan{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateResponsePlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-incidents",
		OperationName: "UpdateResponsePlan",
	}
}
