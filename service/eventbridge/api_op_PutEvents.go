// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	ebcust "github.com/aws/aws-sdk-go-v2/service/eventbridge/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends custom events to Amazon EventBridge so that they can be matched to rules.
// PutEvents will only process nested JSON up to 1100 levels deep.
func (c *Client) PutEvents(ctx context.Context, params *PutEventsInput, optFns ...func(*Options)) (*PutEventsOutput, error) {
	if params == nil {
		params = &PutEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEvents", params, optFns, c.addOperationPutEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEventsInput struct {

	// The entry that defines an event in your system. You can specify several
	// parameters for the entry such as the source and type of the event, resources
	// associated with the event, and so on.
	//
	// This member is required.
	Entries []types.PutEventsRequestEntry

	// The URL subdomain of the endpoint. For example, if the URL for Endpoint is
	// https://abcde.veo.endpoints.event.amazonaws.com, then the EndpointId is
	// abcde.veo . When using Java, you must include auth-crt on the class path.
	EndpointId *string

	noSmithyDocumentSerde
}

func (*PutEventsInput) operationName() string {
	return "PutEvents"
}

func (in *PutEventsInput) bindEndpointParams(p *EndpointParameters) {
	p.EndpointId = in.EndpointId

}

type PutEventsOutput struct {

	// The successfully and unsuccessfully ingested events results. If the ingestion
	// was successful, the entry has the event ID in it. Otherwise, you can use the
	// error code and error message to identify the problem with the entry. For each
	// record, the index of the response element is the same as the index in the
	// request array.
	Entries []types.PutEventsResultEntry

	// The number of failed entries.
	FailedEntryCount int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEvents{}, middleware.After)
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
	if err = addPutEventsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEvents(options.Region), middleware.Before); err != nil {
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

// getPutEventsEndpointId returns a pointer to string denoting a provided member
// value and a boolean indicating if the value is not nil
func getPutEventsEndpointId(input interface{}) (*string, bool) {
	in := input.(*PutEventsInput)
	if in.EndpointId == nil {
		return nil, false
	}
	return in.EndpointId, true
}

func addPutEventsUpdateEndpoint(stack *middleware.Stack, o Options) error {
	return ebcust.UpdateEndpoint(stack, ebcust.UpdateEndpointOptions{
		GetEndpointIDFromInput:  getPutEventsEndpointId,
		EndpointResolver:        o.EndpointResolver,
		EndpointResolverOptions: o.EndpointOptions,
	})
}

func newServiceMetadataMiddleware_opPutEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "PutEvents",
	}
}
