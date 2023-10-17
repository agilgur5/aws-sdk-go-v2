// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Request a list of field-level encryption profiles that have been created in
// CloudFront for this account.
func (c *Client) ListFieldLevelEncryptionProfiles(ctx context.Context, params *ListFieldLevelEncryptionProfilesInput, optFns ...func(*Options)) (*ListFieldLevelEncryptionProfilesOutput, error) {
	if params == nil {
		params = &ListFieldLevelEncryptionProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFieldLevelEncryptionProfiles", params, optFns, c.addOperationListFieldLevelEncryptionProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFieldLevelEncryptionProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFieldLevelEncryptionProfilesInput struct {

	// Use this when paginating results to indicate where to begin in your list of
	// profiles. The results include profiles in the list that occur after the marker.
	// To get the next page of results, set the Marker to the value of the NextMarker
	// from the current page's response (which is also the ID of the last profile on
	// that page).
	Marker *string

	// The maximum number of field-level encryption profiles you want in the response
	// body.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListFieldLevelEncryptionProfilesOutput struct {

	// Returns a list of the field-level encryption profiles that have been created in
	// CloudFront for this account.
	FieldLevelEncryptionProfileList *types.FieldLevelEncryptionProfileList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFieldLevelEncryptionProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListFieldLevelEncryptionProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListFieldLevelEncryptionProfiles{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFieldLevelEncryptionProfiles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFieldLevelEncryptionProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListFieldLevelEncryptionProfiles",
	}
}
