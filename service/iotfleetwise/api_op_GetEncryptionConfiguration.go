// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the encryption configuration for resources and data in Amazon Web
// Services IoT FleetWise.
func (c *Client) GetEncryptionConfiguration(ctx context.Context, params *GetEncryptionConfigurationInput, optFns ...func(*Options)) (*GetEncryptionConfigurationOutput, error) {
	if params == nil {
		params = &GetEncryptionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEncryptionConfiguration", params, optFns, c.addOperationGetEncryptionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEncryptionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEncryptionConfigurationInput struct {
	noSmithyDocumentSerde
}

type GetEncryptionConfigurationOutput struct {

	// The encryption status.
	//
	// This member is required.
	EncryptionStatus types.EncryptionStatus

	// The type of encryption. Set to KMS_BASED_ENCRYPTION to use an KMS key that you
	// own and manage. Set to FLEETWISE_DEFAULT_ENCRYPTION to use an Amazon Web
	// Services managed key that is owned by the Amazon Web Services IoT FleetWise
	// service account.
	//
	// This member is required.
	EncryptionType types.EncryptionType

	// The time when encryption was configured in seconds since epoch (January 1, 1970
	// at midnight UTC time).
	CreationTime *time.Time

	// The error message that describes why encryption settings couldn't be
	// configured, if applicable.
	ErrorMessage *string

	// The ID of the KMS key that is used for encryption.
	KmsKeyId *string

	// The time when encryption was last updated in seconds since epoch (January 1,
	// 1970 at midnight UTC time).
	LastModificationTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEncryptionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEncryptionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEncryptionConfiguration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEncryptionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEncryptionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "GetEncryptionConfiguration",
	}
}
