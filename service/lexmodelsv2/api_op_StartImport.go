// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts importing a bot, bot locale, or custom vocabulary from a zip archive
// that you uploaded to an S3 bucket.
func (c *Client) StartImport(ctx context.Context, params *StartImportInput, optFns ...func(*Options)) (*StartImportOutput, error) {
	if params == nil {
		params = &StartImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImport", params, optFns, c.addOperationStartImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImportInput struct {

	// The unique identifier for the import. It is included in the response from the
	// CreateUploadUrl (https://docs.aws.amazon.com/lexv2/latest/APIReference/API_CreateUploadUrl.html)
	// operation.
	//
	// This member is required.
	ImportId *string

	// The strategy to use when there is a name conflict between the imported resource
	// and an existing resource. When the merge strategy is FailOnConflict existing
	// resources are not overwritten and the import fails.
	//
	// This member is required.
	MergeStrategy types.MergeStrategy

	// Parameters for creating the bot, bot locale or custom vocabulary.
	//
	// This member is required.
	ResourceSpecification *types.ImportResourceSpecification

	// The password used to encrypt the zip archive that contains the resource
	// definition. You should always encrypt the zip archive to protect it during
	// transit between your site and Amazon Lex.
	FilePassword *string

	noSmithyDocumentSerde
}

type StartImportOutput struct {

	// The date and time that the import request was created.
	CreationDateTime *time.Time

	// A unique identifier for the import.
	ImportId *string

	// The current status of the import. When the status is Complete the bot, bot
	// alias, or custom vocabulary is ready to use.
	ImportStatus types.ImportStatus

	// The strategy used when there was a name conflict between the imported resource
	// and an existing resource. When the merge strategy is FailOnConflict existing
	// resources are not overwritten and the import fails.
	MergeStrategy types.MergeStrategy

	// The parameters used when importing the resource.
	ResourceSpecification *types.ImportResourceSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartImport{}, middleware.After)
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
	if err = addOpStartImportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "StartImport",
	}
}
