// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a journal stream for a given Amazon QLDB ledger. The stream captures
// every document revision that is committed to the ledger's journal and delivers
// the data to a specified Amazon Kinesis Data Streams resource.
func (c *Client) StreamJournalToKinesis(ctx context.Context, params *StreamJournalToKinesisInput, optFns ...func(*Options)) (*StreamJournalToKinesisOutput, error) {
	if params == nil {
		params = &StreamJournalToKinesisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StreamJournalToKinesis", params, optFns, c.addOperationStreamJournalToKinesisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StreamJournalToKinesisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StreamJournalToKinesisInput struct {

	// The inclusive start date and time from which to start streaming journal data.
	// This parameter must be in ISO 8601 date and time format and in Universal
	// Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z . The
	// InclusiveStartTime cannot be in the future and must be before ExclusiveEndTime .
	// If you provide an InclusiveStartTime that is before the ledger's
	// CreationDateTime , QLDB effectively defaults it to the ledger's CreationDateTime
	// .
	//
	// This member is required.
	InclusiveStartTime *time.Time

	// The configuration settings of the Kinesis Data Streams destination for your
	// stream request.
	//
	// This member is required.
	KinesisConfiguration *types.KinesisConfiguration

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for
	// a journal stream to write data records to a Kinesis Data Streams resource. To
	// pass a role to QLDB when requesting a journal stream, you must have permissions
	// to perform the iam:PassRole action on the IAM role resource. This is required
	// for all journal stream requests.
	//
	// This member is required.
	RoleArn *string

	// The name that you want to assign to the QLDB journal stream. User-defined names
	// can help identify and indicate the purpose of a stream. Your stream name must be
	// unique among other active streams for a given ledger. Stream names have the same
	// naming constraints as ledger names, as defined in Quotas in Amazon QLDB (https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming)
	// in the Amazon QLDB Developer Guide.
	//
	// This member is required.
	StreamName *string

	// The exclusive date and time that specifies when the stream ends. If you don't
	// define this parameter, the stream runs indefinitely until you cancel it. The
	// ExclusiveEndTime must be in ISO 8601 date and time format and in Universal
	// Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z .
	ExclusiveEndTime *time.Time

	// The key-value pairs to add as tags to the stream that you want to create. Tag
	// keys are case sensitive. Tag values are case sensitive and can be null.
	Tags map[string]*string

	noSmithyDocumentSerde
}

type StreamJournalToKinesisOutput struct {

	// The UUID (represented in Base62-encoded text) that QLDB assigns to each QLDB
	// journal stream.
	StreamId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStreamJournalToKinesisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStreamJournalToKinesis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStreamJournalToKinesis{}, middleware.After)
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
	if err = addOpStreamJournalToKinesisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStreamJournalToKinesis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStreamJournalToKinesis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "StreamJournalToKinesis",
	}
}
