// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified IP address filter. For information about managing IP
// address filters, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-ip-filtering-console-walkthrough.html)
// . You can execute this operation no more than once per second.
func (c *Client) DeleteReceiptFilter(ctx context.Context, params *DeleteReceiptFilterInput, optFns ...func(*Options)) (*DeleteReceiptFilterOutput, error) {
	if params == nil {
		params = &DeleteReceiptFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReceiptFilter", params, optFns, c.addOperationDeleteReceiptFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReceiptFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete an IP address filter. You use IP address filters
// when you receive email with Amazon SES. For more information, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html)
// .
type DeleteReceiptFilterInput struct {

	// The name of the IP address filter to delete.
	//
	// This member is required.
	FilterName *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type DeleteReceiptFilterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteReceiptFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteReceiptFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteReceiptFilter{}, middleware.After)
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
	if err = addOpDeleteReceiptFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReceiptFilter(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteReceiptFilter",
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

func newServiceMetadataMiddleware_opDeleteReceiptFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteReceiptFilter",
	}
}
