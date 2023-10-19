// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a changeset for a kdb database. A changeset allows you to add and
// delete existing files by using an ordered list of change requests.
func (c *Client) CreateKxChangeset(ctx context.Context, params *CreateKxChangesetInput, optFns ...func(*Options)) (*CreateKxChangesetOutput, error) {
	if params == nil {
		params = &CreateKxChangesetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKxChangeset", params, optFns, c.addOperationCreateKxChangesetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKxChangesetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKxChangesetInput struct {

	// A list of change request objects that are run in order. A change request object
	// consists of changeType , s3Path, and a dbPath. A changeType can has the
	// following values:
	//   - PUT – Adds or updates files in a database.
	//   - DELETE – Deletes files in a database.
	// All the change requests require a mandatory dbPath attribute that defines the
	// path within the database directory. The s3Path attribute defines the s3 source
	// file path and is required for a PUT change type. Here is an example of how you
	// can use the change request object: [ { "changeType": "PUT",
	// "s3Path":"s3://bucket/db/2020.01.02/", "dbPath":"/2020.01.02/"}, { "changeType":
	// "PUT", "s3Path":"s3://bucket/db/sym", "dbPath":"/"}, { "changeType": "DELETE",
	// "dbPath": "/2020.01.01/"} ] In this example, the first request with PUT change
	// type allows you to add files in the given s3Path under the 2020.01.02 partition
	// of the database. The second request with PUT change type allows you to add a
	// single sym file at database root location. The last request with DELETE change
	// type allows you to delete the files under the 2020.01.01 partition of the
	// database.
	//
	// This member is required.
	ChangeRequests []types.ChangeRequest

	// A token that ensures idempotency. This token expires in 10 minutes.
	//
	// This member is required.
	ClientToken *string

	// The name of the kdb database.
	//
	// This member is required.
	DatabaseName *string

	// A unique identifier of the kdb environment.
	//
	// This member is required.
	EnvironmentId *string

	noSmithyDocumentSerde
}

type CreateKxChangesetOutput struct {

	// A list of change requests.
	ChangeRequests []types.ChangeRequest

	// A unique identifier for the changeset.
	ChangesetId *string

	// The timestamp at which the changeset was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	// The name of the kdb database.
	DatabaseName *string

	// A unique identifier for the kdb environment.
	EnvironmentId *string

	// The details of the error that you receive when creating a changeset. It
	// consists of the type of error and the error message.
	ErrorInfo *types.ErrorInfo

	// The timestamp at which the changeset was updated in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	// Status of the changeset creation process.
	//   - Pending – Changeset creation is pending.
	//   - Processing – Changeset creation is running.
	//   - Failed – Changeset creation has failed.
	//   - Complete – Changeset creation has succeeded.
	Status types.ChangesetStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKxChangesetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKxChangeset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKxChangeset{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addIdempotencyToken_opCreateKxChangesetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKxChangesetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKxChangeset(options.Region), middleware.Before); err != nil {
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
		operation: "CreateKxChangeset",
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

type idempotencyToken_initializeOpCreateKxChangeset struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKxChangeset) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKxChangeset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKxChangesetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKxChangesetInput ")
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
func addIdempotencyToken_opCreateKxChangesetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKxChangeset{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKxChangeset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace",
		OperationName: "CreateKxChangeset",
	}
}
