// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation has three possible return values:
//   - A successful response in the form of GreetingWithErrorsOutput
//   - An InvalidGreeting error.
//   - A BadRequest error.
//
// Implementations must be able to successfully take a response and properly
// (de)serialize successful and error responses based on the the presence of the
func (c *Client) GreetingWithErrors(ctx context.Context, params *GreetingWithErrorsInput, optFns ...func(*Options)) (*GreetingWithErrorsOutput, error) {
	if params == nil {
		params = &GreetingWithErrorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GreetingWithErrors", params, optFns, c.addOperationGreetingWithErrorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GreetingWithErrorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GreetingWithErrorsInput struct {
	noSmithyDocumentSerde
}

func (*GreetingWithErrorsInput) operationName() string {
	return "GreetingWithErrors"
}

type GreetingWithErrorsOutput struct {
	Greeting *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGreetingWithErrorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGreetingWithErrors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGreetingWithErrors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGreetingWithErrors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGreetingWithErrors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GreetingWithErrors",
	}
}
