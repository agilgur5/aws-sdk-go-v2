// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// (Discontinued) Removes specified S3 resources from being monitored by Amazon
// Macie Classic. If memberAccountId isn't specified, the action removes specified
// S3 resources from Macie Classic for the current Macie Classic administrator
// account. If memberAccountId is specified, the action removes specified S3
// resources from Macie Classic for the specified member account.
func (c *Client) DisassociateS3Resources(ctx context.Context, params *DisassociateS3ResourcesInput, optFns ...func(*Options)) (*DisassociateS3ResourcesOutput, error) {
	if params == nil {
		params = &DisassociateS3ResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateS3Resources", params, optFns, c.addOperationDisassociateS3ResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateS3ResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateS3ResourcesInput struct {

	// (Discontinued) The S3 resources (buckets or prefixes) that you want to remove
	// from being monitored and classified by Amazon Macie Classic.
	//
	// This member is required.
	AssociatedS3Resources []types.S3Resource

	// (Discontinued) The ID of the Amazon Macie Classic member account whose
	// resources you want to remove from being monitored by Macie Classic.
	MemberAccountId *string

	noSmithyDocumentSerde
}

type DisassociateS3ResourcesOutput struct {

	// (Discontinued) S3 resources that couldn't be removed from being monitored and
	// classified by Amazon Macie Classic. An error code and an error message are
	// provided for each failed item.
	FailedS3Resources []types.FailedS3Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateS3ResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateS3Resources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateS3Resources{}, middleware.After)
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
	if err = addOpDisassociateS3ResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateS3Resources(options.Region), middleware.Before); err != nil {
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
		operation: "DisassociateS3Resources",
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

func newServiceMetadataMiddleware_opDisassociateS3Resources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie",
		OperationName: "DisassociateS3Resources",
	}
}
