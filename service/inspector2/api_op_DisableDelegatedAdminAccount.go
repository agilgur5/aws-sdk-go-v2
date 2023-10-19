// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables the Amazon Inspector delegated administrator for your organization.
func (c *Client) DisableDelegatedAdminAccount(ctx context.Context, params *DisableDelegatedAdminAccountInput, optFns ...func(*Options)) (*DisableDelegatedAdminAccountOutput, error) {
	if params == nil {
		params = &DisableDelegatedAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableDelegatedAdminAccount", params, optFns, c.addOperationDisableDelegatedAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableDelegatedAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableDelegatedAdminAccountInput struct {

	// The Amazon Web Services account ID of the current Amazon Inspector delegated
	// administrator.
	//
	// This member is required.
	DelegatedAdminAccountId *string

	noSmithyDocumentSerde
}

type DisableDelegatedAdminAccountOutput struct {

	// The Amazon Web Services account ID of the successfully disabled delegated
	// administrator.
	//
	// This member is required.
	DelegatedAdminAccountId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableDelegatedAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisableDelegatedAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisableDelegatedAdminAccount{}, middleware.After)
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
	if err = addOpDisableDelegatedAdminAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableDelegatedAdminAccount(options.Region), middleware.Before); err != nil {
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
		operation: "DisableDelegatedAdminAccount",
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

func newServiceMetadataMiddleware_opDisableDelegatedAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "DisableDelegatedAdminAccount",
	}
}
