// Code generated by smithy-go-codegen DO NOT EDIT.

package controltower

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API call turns off a control. It starts an asynchronous operation that
// deletes AWS resources on the specified organizational unit and the accounts it
// contains. The resources will vary according to the control that you specify. For
// usage examples, see the AWS Control Tower User Guide  (https://docs.aws.amazon.com/controltower/latest/userguide/control-api-examples-short.html)
// .
func (c *Client) DisableControl(ctx context.Context, params *DisableControlInput, optFns ...func(*Options)) (*DisableControlOutput, error) {
	if params == nil {
		params = &DisableControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableControl", params, optFns, c.addOperationDisableControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableControlInput struct {

	// The ARN of the control. Only Strongly recommended and Elective controls are
	// permitted, with the exception of the Region deny control. For information on how
	// to find the controlIdentifier , see the overview page (https://docs.aws.amazon.com/controltower/latest/APIReference/Welcome.html)
	// .
	//
	// This member is required.
	ControlIdentifier *string

	// The ARN of the organizational unit. For information on how to find the
	// targetIdentifier , see the overview page (https://docs.aws.amazon.com/controltower/latest/APIReference/Welcome.html)
	// .
	//
	// This member is required.
	TargetIdentifier *string

	noSmithyDocumentSerde
}

type DisableControlOutput struct {

	// The ID of the asynchronous operation, which is used to track status. The
	// operation is available for 90 days.
	//
	// This member is required.
	OperationIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisableControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisableControl{}, middleware.After)
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
	if err = addOpDisableControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "controltower",
		OperationName: "DisableControl",
	}
}
