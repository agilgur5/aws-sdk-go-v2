// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a DataflowEndpoint group containing the specified list of
// DataflowEndpoint objects. The name field in each endpoint is used in your
// mission profile DataflowEndpointConfig to specify which endpoints to use during
// a contact. When a contact uses multiple DataflowEndpointConfig objects, each
// Config must match a DataflowEndpoint in the same group.
func (c *Client) CreateDataflowEndpointGroup(ctx context.Context, params *CreateDataflowEndpointGroupInput, optFns ...func(*Options)) (*CreateDataflowEndpointGroupOutput, error) {
	if params == nil {
		params = &CreateDataflowEndpointGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataflowEndpointGroup", params, optFns, c.addOperationCreateDataflowEndpointGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataflowEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataflowEndpointGroupInput struct {

	// Endpoint details of each endpoint in the dataflow endpoint group.
	//
	// This member is required.
	EndpointDetails []types.EndpointDetails

	// Amount of time, in seconds, after a contact ends that the Ground Station
	// Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow
	// Endpoint Group State Change event will be emitted when the Dataflow Endpoint
	// Group enters and exits the POSTPASS state.
	ContactPostPassDurationSeconds *int32

	// Amount of time, in seconds, before a contact starts that the Ground Station
	// Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow
	// Endpoint Group State Change event will be emitted when the Dataflow Endpoint
	// Group enters and exits the PREPASS state.
	ContactPrePassDurationSeconds *int32

	// Tags of a dataflow endpoint group.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDataflowEndpointGroupOutput struct {

	// UUID of a dataflow endpoint group.
	DataflowEndpointGroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataflowEndpointGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataflowEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataflowEndpointGroup{}, middleware.After)
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
	if err = addOpCreateDataflowEndpointGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataflowEndpointGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataflowEndpointGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "groundstation",
		OperationName: "CreateDataflowEndpointGroup",
	}
}
