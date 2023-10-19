// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes a data source in Amazon DataZone.
func (c *Client) DeleteDataSource(ctx context.Context, params *DeleteDataSourceInput, optFns ...func(*Options)) (*DeleteDataSourceOutput, error) {
	if params == nil {
		params = &DeleteDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataSource", params, optFns, c.addOperationDeleteDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDataSourceInput struct {

	// The ID of the Amazon DataZone domain in which the data source is deleted.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the data source that is deleted.
	//
	// This member is required.
	Identifier *string

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	noSmithyDocumentSerde
}

type DeleteDataSourceOutput struct {

	// The ID of the Amazon DataZone domain in which the data source is deleted.
	//
	// This member is required.
	DomainId *string

	// The ID of the environemnt associated with this data source.
	//
	// This member is required.
	EnvironmentId *string

	// The ID of the data source that is deleted.
	//
	// This member is required.
	Id *string

	// The name of the data source that is deleted.
	//
	// This member is required.
	Name *string

	// The ID of the project in which this data source exists and from which it's
	// deleted.
	//
	// This member is required.
	ProjectId *string

	// The asset data forms associated with this data source.
	AssetFormsOutput []types.FormOutput

	// The configuration of the data source that is deleted.
	Configuration types.DataSourceConfigurationOutput

	// The timestamp of when this data source was created.
	CreatedAt *time.Time

	// The description of the data source that is deleted.
	Description *string

	// The enable setting of the data source that specifies whether the data source is
	// enabled or disabled.
	EnableSetting types.EnableSetting

	// Specifies the error message that is returned if the operation cannot be
	// successfully completed.
	ErrorMessage *types.DataSourceErrorMessage

	// The timestamp of when the data source was last run.
	LastRunAt *time.Time

	// Specifies the error message that is returned if the operation cannot be
	// successfully completed.
	LastRunErrorMessage *types.DataSourceErrorMessage

	// The status of the last run of this data source.
	LastRunStatus types.DataSourceRunStatus

	// Specifies whether the assets that this data source creates in the inventory are
	// to be also automatically published to the catalog.
	PublishOnImport *bool

	// The schedule of runs for this data source.
	Schedule *types.ScheduleConfiguration

	// The status of this data source.
	Status types.DataSourceStatus

	// The type of this data source.
	Type *string

	// The timestamp of when this data source was updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataSource{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteDataSourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataSource(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteDataSource",
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

type idempotencyToken_initializeOpDeleteDataSource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteDataSource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteDataSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteDataSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteDataSourceInput ")
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
func addIdempotencyToken_opDeleteDataSourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteDataSource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "DeleteDataSource",
	}
}
