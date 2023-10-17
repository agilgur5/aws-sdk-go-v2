// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcampaigns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectcampaigns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Onboard the specific Amazon Connect instance to Connect Campaigns.
func (c *Client) StartInstanceOnboardingJob(ctx context.Context, params *StartInstanceOnboardingJobInput, optFns ...func(*Options)) (*StartInstanceOnboardingJobOutput, error) {
	if params == nil {
		params = &StartInstanceOnboardingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartInstanceOnboardingJob", params, optFns, c.addOperationStartInstanceOnboardingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartInstanceOnboardingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request for StartInstanceOnboardingJob API.
type StartInstanceOnboardingJobInput struct {

	// Amazon Connect Instance Id
	//
	// This member is required.
	ConnectInstanceId *string

	// Encryption config for Connect Instance. Note that sensitive data will always be
	// encrypted. If disabled, service will perform encryption with its own key. If
	// enabled, a KMS key id needs to be provided and KMS charges will apply. KMS is
	// only type supported
	//
	// This member is required.
	EncryptionConfig *types.EncryptionConfig

	noSmithyDocumentSerde
}

// The response for StartInstanceOnboardingJob API.
type StartInstanceOnboardingJobOutput struct {

	// Instance onboarding job status object
	ConnectInstanceOnboardingJobStatus *types.InstanceOnboardingJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartInstanceOnboardingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartInstanceOnboardingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartInstanceOnboardingJob{}, middleware.After)
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
	if err = addOpStartInstanceOnboardingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartInstanceOnboardingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartInstanceOnboardingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect-campaigns",
		OperationName: "StartInstanceOnboardingJob",
	}
}
