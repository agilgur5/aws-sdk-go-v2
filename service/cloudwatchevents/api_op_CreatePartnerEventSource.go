// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Called by an SaaS partner to create a partner event source. This operation is
// not used by Amazon Web Services customers. Each partner event source can be used
// by one Amazon Web Services account to create a matching partner event bus in
// that Amazon Web Services account. A SaaS partner must create one partner event
// source for each Amazon Web Services account that wants to receive those event
// types. A partner event source creates events based on resources within the SaaS
// partner's service or application. An Amazon Web Services account that creates a
// partner event bus that matches the partner event source can use that event bus
// to receive events from the partner, and then process them using Amazon Web
// Services Events rules and targets. Partner event source names follow this
// format: partner_name/event_namespace/event_name  partner_name is determined
// during partner registration and identifies the partner to Amazon Web Services
// customers. event_namespace is determined by the partner and is a way for the
// partner to categorize their events. event_name is determined by the partner, and
// should uniquely identify an event-generating resource within the partner system.
// The combination of event_namespace and event_name should help Amazon Web
// Services customers decide whether to create an event bus to receive these
// events.
func (c *Client) CreatePartnerEventSource(ctx context.Context, params *CreatePartnerEventSourceInput, optFns ...func(*Options)) (*CreatePartnerEventSourceOutput, error) {
	if params == nil {
		params = &CreatePartnerEventSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePartnerEventSource", params, optFns, c.addOperationCreatePartnerEventSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePartnerEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePartnerEventSourceInput struct {

	// The Amazon Web Services account ID that is permitted to create a matching
	// partner event bus for this partner event source.
	//
	// This member is required.
	Account *string

	// The name of the partner event source. This name must be unique and must be in
	// the format partner_name/event_namespace/event_name . The Amazon Web Services
	// account that wants to use this partner event source must create a partner event
	// bus with a name that matches the name of the partner event source.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreatePartnerEventSourceOutput struct {

	// The ARN of the partner event source.
	EventSourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePartnerEventSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePartnerEventSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePartnerEventSource{}, middleware.After)
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
	if err = addOpCreatePartnerEventSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePartnerEventSource(options.Region), middleware.Before); err != nil {
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
		operation: "CreatePartnerEventSource",
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

func newServiceMetadataMiddleware_opCreatePartnerEventSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "CreatePartnerEventSource",
	}
}
