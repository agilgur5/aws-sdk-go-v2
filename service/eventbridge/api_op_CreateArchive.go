// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an archive of events with the specified settings. When you create an
// archive, incoming events might not immediately start being sent to the archive.
// Allow a short period of time for changes to take effect. If you do not specify a
// pattern to filter events sent to the archive, all events are sent to the archive
// except replayed events. Replayed events are not sent to an archive.
func (c *Client) CreateArchive(ctx context.Context, params *CreateArchiveInput, optFns ...func(*Options)) (*CreateArchiveOutput, error) {
	if params == nil {
		params = &CreateArchiveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateArchive", params, optFns, c.addOperationCreateArchiveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateArchiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateArchiveInput struct {

	// The name for the archive to create.
	//
	// This member is required.
	ArchiveName *string

	// The ARN of the event bus that sends events to the archive.
	//
	// This member is required.
	EventSourceArn *string

	// A description for the archive.
	Description *string

	// An event pattern to use to filter events sent to the archive.
	EventPattern *string

	// The number of days to retain events for. Default value is 0. If set to 0,
	// events are retained indefinitely
	RetentionDays *int32

	noSmithyDocumentSerde
}

type CreateArchiveOutput struct {

	// The ARN of the archive that was created.
	ArchiveArn *string

	// The time at which the archive was created.
	CreationTime *time.Time

	// The state of the archive that was created.
	State types.ArchiveState

	// The reason that the archive is in the state.
	StateReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateArchiveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateArchive{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateArchive{}, middleware.After)
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
	if err = addOpCreateArchiveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateArchive(options.Region), middleware.Before); err != nil {
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
		operation: "CreateArchive",
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

func newServiceMetadataMiddleware_opCreateArchive(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "CreateArchive",
	}
}
