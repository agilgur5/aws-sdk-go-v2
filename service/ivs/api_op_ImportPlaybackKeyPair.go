// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports the public portion of a new key pair and returns its arn and fingerprint
// . The privateKey can then be used to generate viewer authorization tokens, to
// grant viewers access to private channels. For more information, see Setting Up
// Private Channels (https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html)
// in the Amazon IVS User Guide.
func (c *Client) ImportPlaybackKeyPair(ctx context.Context, params *ImportPlaybackKeyPairInput, optFns ...func(*Options)) (*ImportPlaybackKeyPairOutput, error) {
	if params == nil {
		params = &ImportPlaybackKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportPlaybackKeyPair", params, optFns, c.addOperationImportPlaybackKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportPlaybackKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportPlaybackKeyPairInput struct {

	// The public portion of a customer-generated key pair.
	//
	// This member is required.
	PublicKeyMaterial *string

	// Playback-key-pair name. The value does not need to be unique.
	Name *string

	// Any tags provided with the request are added to the playback key pair tags. See
	// Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ImportPlaybackKeyPairOutput struct {

	//
	KeyPair *types.PlaybackKeyPair

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportPlaybackKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportPlaybackKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportPlaybackKeyPair{}, middleware.After)
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
	if err = addOpImportPlaybackKeyPairValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportPlaybackKeyPair(options.Region), middleware.Before); err != nil {
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
		operation: "ImportPlaybackKeyPair",
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

func newServiceMetadataMiddleware_opImportPlaybackKeyPair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "ImportPlaybackKeyPair",
	}
}
