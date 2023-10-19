// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Writes multiple data records into a Kinesis data stream in a single call (also
// referred to as a PutRecords request). Use this operation to send data into the
// stream for data ingestion and processing. When invoking this API, it is
// recommended you use the StreamARN input parameter rather than the StreamName
// input parameter. Each PutRecords request can support up to 500 records. Each
// record in the request can be as large as 1 MiB, up to a limit of 5 MiB for the
// entire request, including partition keys. Each shard can support writes up to
// 1,000 records per second, up to a maximum data write total of 1 MiB per second.
// You must specify the name of the stream that captures, stores, and transports
// the data; and an array of request Records , with each record in the array
// requiring a partition key and data blob. The record size limit applies to the
// total size of the partition key and data blob. The data blob can be any type of
// data; for example, a segment from a log file, geographic/location data, website
// clickstream data, and so on. The partition key is used by Kinesis Data Streams
// as input to a hash function that maps the partition key and associated data to a
// specific shard. An MD5 hash function is used to map partition keys to 128-bit
// integer values and to map associated data records to shards. As a result of this
// hashing mechanism, all data records with the same partition key map to the same
// shard within the stream. For more information, see Adding Data to a Stream (https://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-add-data-to-stream)
// in the Amazon Kinesis Data Streams Developer Guide. Each record in the Records
// array may include an optional parameter, ExplicitHashKey , which overrides the
// partition key to shard mapping. This parameter allows a data producer to
// determine explicitly the shard where the record is stored. For more information,
// see Adding Multiple Records with PutRecords (https://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-putrecords)
// in the Amazon Kinesis Data Streams Developer Guide. The PutRecords response
// includes an array of response Records . Each record in the response array
// directly correlates with a record in the request array using natural ordering,
// from the top to the bottom of the request and response. The response Records
// array always includes the same number of records as the request array. The
// response Records array includes both successfully and unsuccessfully processed
// records. Kinesis Data Streams attempts to process all records in each PutRecords
// request. A single record failure does not stop the processing of subsequent
// records. As a result, PutRecords doesn't guarantee the ordering of records. If
// you need to read records in the same order they are written to the stream, use
// PutRecord instead of PutRecords , and write to the same shard. A successfully
// processed record includes ShardId and SequenceNumber values. The ShardId
// parameter identifies the shard in the stream where the record is stored. The
// SequenceNumber parameter is an identifier assigned to the put record, unique to
// all records in the stream. An unsuccessfully processed record includes ErrorCode
// and ErrorMessage values. ErrorCode reflects the type of error and can be one of
// the following values: ProvisionedThroughputExceededException or InternalFailure
// . ErrorMessage provides more detailed information about the
// ProvisionedThroughputExceededException exception including the account ID,
// stream name, and shard ID of the record that was throttled. For more information
// about partially successful responses, see Adding Multiple Records with
// PutRecords (https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-add-data-to-stream.html#kinesis-using-sdk-java-putrecords)
// in the Amazon Kinesis Data Streams Developer Guide. After you write a record to
// a stream, you cannot modify that record or its order within the stream. By
// default, data records are accessible for 24 hours from the time that they are
// added to a stream. You can use IncreaseStreamRetentionPeriod or
// DecreaseStreamRetentionPeriod to modify this retention period.
func (c *Client) PutRecords(ctx context.Context, params *PutRecordsInput, optFns ...func(*Options)) (*PutRecordsOutput, error) {
	if params == nil {
		params = &PutRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRecords", params, optFns, c.addOperationPutRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A PutRecords request.
type PutRecordsInput struct {

	// The records associated with the request.
	//
	// This member is required.
	Records []types.PutRecordsRequestEntry

	// The ARN of the stream.
	StreamARN *string

	// The stream name associated with the request.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *PutRecordsInput) bindEndpointParams(p *EndpointParameters) {
	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("data")
}

// PutRecords results.
type PutRecordsOutput struct {

	// An array of successfully and unsuccessfully processed record results. A record
	// that is successfully added to a stream includes SequenceNumber and ShardId in
	// the result. A record that fails to be added to a stream includes ErrorCode and
	// ErrorMessage in the result.
	//
	// This member is required.
	Records []types.PutRecordsResultEntry

	// The encryption type used on the records. This parameter can be one of the
	// following values:
	//   - NONE : Do not encrypt the records.
	//   - KMS : Use server-side encryption on the records using a customer-managed
	//   Amazon Web Services KMS key.
	EncryptionType types.EncryptionType

	// The number of unsuccessfully processed records in a PutRecords request.
	FailedRecordCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRecords{}, middleware.After)
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
	if err = addOpPutRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecords(options.Region), middleware.Before); err != nil {
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
		operation: "PutRecords",
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

func newServiceMetadataMiddleware_opPutRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "PutRecords",
	}
}
