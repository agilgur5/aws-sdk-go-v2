// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a data source connector that you want to use with an Amazon Kendra
// index. You specify a name, data source connector type and description for your
// data source. You also specify configuration information for the data source
// connector. CreateDataSource is a synchronous operation. The operation returns
// 200 if the data source was successfully created. Otherwise, an exception is
// raised. For an example of creating an index and data source using the Python
// SDK, see Getting started with Python SDK (https://docs.aws.amazon.com/kendra/latest/dg/gs-python.html)
// . For an example of creating an index and data source using the Java SDK, see
// Getting started with Java SDK (https://docs.aws.amazon.com/kendra/latest/dg/gs-java.html)
// .
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	if params == nil {
		params = &CreateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSource", params, optFns, c.addOperationCreateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {

	// The identifier of the index you want to use with the data source connector.
	//
	// This member is required.
	IndexId *string

	// A name for the data source connector.
	//
	// This member is required.
	Name *string

	// The type of data source repository. For example, SHAREPOINT .
	//
	// This member is required.
	Type types.DataSourceType

	// A token that you provide to identify the request to create a data source
	// connector. Multiple calls to the CreateDataSource API with the same client
	// token will create only one data source connector.
	ClientToken *string

	// Configuration information to connect to your data source repository. You can't
	// specify the Configuration parameter when the Type parameter is set to CUSTOM .
	// If you do, you receive a ValidationException exception. The Configuration
	// parameter is required for all other data sources.
	Configuration *types.DataSourceConfiguration

	// Configuration information for altering document metadata and content during the
	// document ingestion process. For more information on how to create, modify and
	// delete document metadata, or make other content alterations when you ingest
	// documents into Amazon Kendra, see Customizing document metadata during the
	// ingestion process (https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html)
	// .
	CustomDocumentEnrichmentConfiguration *types.CustomDocumentEnrichmentConfiguration

	// A description for the data source connector.
	Description *string

	// The code for a language. This allows you to support a language for all
	// documents when creating the data source connector. English is supported by
	// default. For more information on supported languages, including their codes, see
	// Adding documents in languages other than English (https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html)
	// .
	LanguageCode *string

	// The Amazon Resource Name (ARN) of an IAM role with permission to access the
	// data source and required resources. For more information, see IAM access roles
	// for Amazon Kendra. (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html)
	// . You can't specify the RoleArn parameter when the Type parameter is set to
	// CUSTOM . If you do, you receive a ValidationException exception. The RoleArn
	// parameter is required for all other data sources.
	RoleArn *string

	// Sets the frequency for Amazon Kendra to check the documents in your data source
	// repository and update the index. If you don't set a schedule Amazon Kendra will
	// not periodically update the index. You can call the StartDataSourceSyncJob API
	// to update the index. Specify a cron- format schedule string or an empty string
	// to indicate that the index is updated on demand. You can't specify the Schedule
	// parameter when the Type parameter is set to CUSTOM . If you do, you receive a
	// ValidationException exception.
	Schedule *string

	// A list of key-value pairs that identify or categorize the data source
	// connector. You can also use tags to help control access to the data source
	// connector. Tag keys and values can consist of Unicode letters, digits, white
	// space, and any of the following symbols: _ . : / = + - @.
	Tags []types.Tag

	// Configuration information for an Amazon Virtual Private Cloud to connect to
	// your data source. For more information, see Configuring a VPC (https://docs.aws.amazon.com/kendra/latest/dg/vpc-configuration.html)
	// .
	VpcConfiguration *types.DataSourceVpcConfiguration

	noSmithyDocumentSerde
}

type CreateDataSourceOutput struct {

	// The identifier of the data source connector.
	//
	// This member is required.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataSource{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateDataSourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDataSource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataSource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataSourceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDataSourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataSource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateDataSource",
	}
}
