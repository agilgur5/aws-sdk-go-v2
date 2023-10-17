// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new key-signing key (KSK) associated with a hosted zone. You can only
// have two KSKs per hosted zone.
func (c *Client) CreateKeySigningKey(ctx context.Context, params *CreateKeySigningKeyInput, optFns ...func(*Options)) (*CreateKeySigningKeyOutput, error) {
	if params == nil {
		params = &CreateKeySigningKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKeySigningKey", params, optFns, c.addOperationCreateKeySigningKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeySigningKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeySigningKeyInput struct {

	// A unique string that identifies the request.
	//
	// This member is required.
	CallerReference *string

	// The unique string (ID) used to identify a hosted zone.
	//
	// This member is required.
	HostedZoneId *string

	// The Amazon resource name (ARN) for a customer managed key in Key Management
	// Service (KMS). The KeyManagementServiceArn must be unique for each key-signing
	// key (KSK) in a single hosted zone. To see an example of KeyManagementServiceArn
	// that grants the correct permissions for DNSSEC, scroll down to Example. You must
	// configure the customer managed customer managed key as follows: Status Enabled
	// Key spec ECC_NIST_P256 Key usage Sign and verify Key policy The key policy must
	// give permission for the following actions:
	//   - DescribeKey
	//   - GetPublicKey
	//   - Sign
	// The key policy must also include the Amazon Route 53 service in the principal
	// for your account. Specify the following:
	//   - "Service": "dnssec-route53.amazonaws.com"
	// For more information about working with a customer managed key in KMS, see Key
	// Management Service concepts (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html)
	// .
	//
	// This member is required.
	KeyManagementServiceArn *string

	// A string used to identify a key-signing key (KSK). Name can include numbers,
	// letters, and underscores (_). Name must be unique for each key-signing key in
	// the same hosted zone.
	//
	// This member is required.
	Name *string

	// A string specifying the initial status of the key-signing key (KSK). You can
	// set the value to ACTIVE or INACTIVE .
	//
	// This member is required.
	Status *string

	noSmithyDocumentSerde
}

type CreateKeySigningKeyOutput struct {

	// A complex type that describes change information about changes made to your
	// hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// The key-signing key (KSK) that the request creates.
	//
	// This member is required.
	KeySigningKey *types.KeySigningKey

	// The unique URL representing the new key-signing key (KSK).
	//
	// This member is required.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKeySigningKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateKeySigningKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateKeySigningKey{}, middleware.After)
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
	if err = addOpCreateKeySigningKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKeySigningKey(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
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

func newServiceMetadataMiddleware_opCreateKeySigningKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateKeySigningKey",
	}
}
