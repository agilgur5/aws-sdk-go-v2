// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get information about an import task and count of device onboarding summary
// information for the import task.
func (c *Client) GetWirelessDeviceImportTask(ctx context.Context, params *GetWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*GetWirelessDeviceImportTaskOutput, error) {
	if params == nil {
		params = &GetWirelessDeviceImportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWirelessDeviceImportTask", params, optFns, c.addOperationGetWirelessDeviceImportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWirelessDeviceImportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWirelessDeviceImportTaskInput struct {

	// The identifier of the import task for which information is requested.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetWirelessDeviceImportTaskOutput struct {

	// The ARN (Amazon Resource Name) of the import task.
	Arn *string

	// The time at which the import task was created.
	CreationTime *time.Time

	// The name of the destination that's assigned to the wireless devices in the
	// import task.
	DestinationName *string

	// The number of devices in the import task that failed to onboard to the import
	// task.
	FailedImportedDeviceCount *int64

	// The identifier of the import task for which information is retrieved.
	Id *string

	// The number of devices in the import task that are waiting for the control log
	// to start processing.
	InitializedImportedDeviceCount *int64

	// The number of devices in the import task that have been onboarded to the import
	// task.
	OnboardedImportedDeviceCount *int64

	// The number of devices in the import task that are waiting in the import task
	// queue to be onboarded.
	PendingImportedDeviceCount *int64

	// The Sidewalk-related information about an import task.
	Sidewalk *types.SidewalkGetStartImportInfo

	// The import task status.
	Status types.ImportTaskStatus

	// The reason for the provided status information, such as a validation error that
	// causes the import task to fail.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWirelessDeviceImportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWirelessDeviceImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWirelessDeviceImportTask{}, middleware.After)
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
	if err = addOpGetWirelessDeviceImportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWirelessDeviceImportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWirelessDeviceImportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetWirelessDeviceImportTask",
	}
}
