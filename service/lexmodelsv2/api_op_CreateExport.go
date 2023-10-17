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

// Creates a zip archive containing the contents of a bot or a bot locale. The
// archive contains a directory structure that contains JSON files that define the
// bot. You can create an archive that contains the complete definition of a bot,
// or you can specify that the archive contain only the definition of a single bot
// locale. For more information about exporting bots, and about the structure of
// the export archive, see Importing and exporting bots  (https://docs.aws.amazon.com/lexv2/latest/dg/importing-exporting.html)
func (c *Client) CreateExport(ctx context.Context, params *CreateExportInput, optFns ...func(*Options)) (*CreateExportOutput, error) {
	if params == nil {
		params = &CreateExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExport", params, optFns, c.addOperationCreateExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExportInput struct {

	// The file format of the bot or bot locale definition files.
	//
	// This member is required.
	FileFormat types.ImportExportFileFormat

	// Specifies the type of resource to export, either a bot or a bot locale. You can
	// only specify one type of resource to export.
	//
	// This member is required.
	ResourceSpecification *types.ExportResourceSpecification

	// An password to use to encrypt the exported archive. Using a password is
	// optional, but you should encrypt the archive to protect the data in transit
	// between Amazon Lex and your local computer.
	FilePassword *string

	noSmithyDocumentSerde
}

type CreateExportOutput struct {

	// The date and time that the request to export a bot was created.
	CreationDateTime *time.Time

	// An identifier for a specific request to create an export.
	ExportId *string

	// The status of the export. When the status is Completed , you can use the
	// DescribeExport (https://docs.aws.amazon.com/lexv2/latest/APIReference/API_DescribeExport.html)
	// operation to get the pre-signed S3 URL link to your exported bot or bot locale.
	ExportStatus types.ExportStatus

	// The file format used for the bot or bot locale definition files.
	FileFormat types.ImportExportFileFormat

	// A description of the type of resource that was exported, either a bot or a bot
	// locale.
	ResourceSpecification *types.ExportResourceSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateExport{}, middleware.After)
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
	if err = addOpCreateExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateExport",
	}
}
