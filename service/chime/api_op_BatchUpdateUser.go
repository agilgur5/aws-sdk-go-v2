// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates user details within the UpdateUserRequestItem object for up to 20 users
// for the specified Amazon Chime account. Currently, only LicenseType updates are
// supported for this action.
func (c *Client) BatchUpdateUser(ctx context.Context, params *BatchUpdateUserInput, optFns ...func(*Options)) (*BatchUpdateUserOutput, error) {
	if params == nil {
		params = &BatchUpdateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateUser", params, optFns, c.addOperationBatchUpdateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateUserInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The request containing the user IDs and details to update.
	//
	// This member is required.
	UpdateUserRequestItems []types.UpdateUserRequestItem

	noSmithyDocumentSerde
}

type BatchUpdateUserOutput struct {

	// If the BatchUpdateUser action fails for one or more of the user IDs in the
	// request, a list of the user IDs is returned, along with error codes and error
	// messages.
	UserErrors []types.UserError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateUser{}, middleware.After)
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
	if err = addOpBatchUpdateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateUser(options.Region), middleware.Before); err != nil {
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
		operation: "BatchUpdateUser",
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

func newServiceMetadataMiddleware_opBatchUpdateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchUpdateUser",
	}
}
