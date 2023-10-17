// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information on a specific signing profile.
func (c *Client) GetSigningProfile(ctx context.Context, params *GetSigningProfileInput, optFns ...func(*Options)) (*GetSigningProfileOutput, error) {
	if params == nil {
		params = &GetSigningProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSigningProfile", params, optFns, c.addOperationGetSigningProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSigningProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSigningProfileInput struct {

	// The name of the target signing profile.
	//
	// This member is required.
	ProfileName *string

	// The AWS account ID of the profile owner.
	ProfileOwner *string

	noSmithyDocumentSerde
}

type GetSigningProfileOutput struct {

	// The Amazon Resource Name (ARN) for the signing profile.
	Arn *string

	// A list of overrides applied by the target signing profile for signing
	// operations.
	Overrides *types.SigningPlatformOverrides

	// A human-readable name for the signing platform associated with the signing
	// profile.
	PlatformDisplayName *string

	// The ID of the platform that is used by the target signing profile.
	PlatformId *string

	// The name of the target signing profile.
	ProfileName *string

	// The current version of the signing profile.
	ProfileVersion *string

	// The signing profile ARN, including the profile version.
	ProfileVersionArn *string

	// Revocation information for a signing profile.
	RevocationRecord *types.SigningProfileRevocationRecord

	// The validity period for a signing job.
	SignatureValidityPeriod *types.SignatureValidityPeriod

	// The ARN of the certificate that the target profile uses for signing operations.
	SigningMaterial *types.SigningMaterial

	// A map of key-value pairs for signing operations that is attached to the target
	// signing profile.
	SigningParameters map[string]string

	// The status of the target signing profile.
	Status types.SigningProfileStatus

	// Reason for the status of the target signing profile.
	StatusReason *string

	// A list of tags associated with the signing profile.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSigningProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSigningProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSigningProfile{}, middleware.After)
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
	if err = addOpGetSigningProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSigningProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSigningProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "GetSigningProfile",
	}
}
