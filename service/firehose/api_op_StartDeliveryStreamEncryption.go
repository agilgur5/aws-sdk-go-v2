// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables server-side encryption (SSE) for the delivery stream. This operation is
// asynchronous. It returns immediately. When you invoke it, Kinesis Data Firehose
// first sets the encryption status of the stream to ENABLING , and then to ENABLED
// . The encryption status of a delivery stream is the Status property in
// DeliveryStreamEncryptionConfiguration . If the operation fails, the encryption
// status changes to ENABLING_FAILED . You can continue to read and write data to
// your delivery stream while the encryption status is ENABLING , but the data is
// not encrypted. It can take up to 5 seconds after the encryption status changes
// to ENABLED before all records written to the delivery stream are encrypted. To
// find out whether a record or a batch of records was encrypted, check the
// response elements PutRecordOutput$Encrypted and PutRecordBatchOutput$Encrypted ,
// respectively. To check the encryption status of a delivery stream, use
// DescribeDeliveryStream . Even if encryption is currently enabled for a delivery
// stream, you can still invoke this operation on it to change the ARN of the CMK
// or both its type and ARN. If you invoke this method to change the CMK, and the
// old CMK is of type CUSTOMER_MANAGED_CMK , Kinesis Data Firehose schedules the
// grant it had on the old CMK for retirement. If the new CMK is of type
// CUSTOMER_MANAGED_CMK , Kinesis Data Firehose creates a grant that enables it to
// use the new CMK to encrypt and decrypt data and to manage the grant. For the KMS
// grant creation to be successful, Kinesis Data Firehose APIs
// StartDeliveryStreamEncryption and CreateDeliveryStream should not be called
// with session credentials that are more than 6 hours old. If a delivery stream
// already has encryption enabled and then you invoke this operation to change the
// ARN of the CMK or both its type and ARN and you get ENABLING_FAILED , this only
// means that the attempt to change the CMK failed. In this case, encryption
// remains enabled with the old CMK. If the encryption status of your delivery
// stream is ENABLING_FAILED , you can invoke this operation again with a valid
// CMK. The CMK must be enabled and the key policy mustn't explicitly deny the
// permission for Kinesis Data Firehose to invoke KMS encrypt and decrypt
// operations. You can enable SSE for a delivery stream only if it's a delivery
// stream that uses DirectPut as its source. The StartDeliveryStreamEncryption and
// StopDeliveryStreamEncryption operations have a combined limit of 25 calls per
// delivery stream per 24 hours. For example, you reach the limit if you call
// StartDeliveryStreamEncryption 13 times and StopDeliveryStreamEncryption 12
// times for the same delivery stream in a 24-hour period.
func (c *Client) StartDeliveryStreamEncryption(ctx context.Context, params *StartDeliveryStreamEncryptionInput, optFns ...func(*Options)) (*StartDeliveryStreamEncryptionOutput, error) {
	if params == nil {
		params = &StartDeliveryStreamEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDeliveryStreamEncryption", params, optFns, c.addOperationStartDeliveryStreamEncryptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDeliveryStreamEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeliveryStreamEncryptionInput struct {

	// The name of the delivery stream for which you want to enable server-side
	// encryption (SSE).
	//
	// This member is required.
	DeliveryStreamName *string

	// Used to specify the type and Amazon Resource Name (ARN) of the KMS key needed
	// for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput *types.DeliveryStreamEncryptionConfigurationInput

	noSmithyDocumentSerde
}

func (*StartDeliveryStreamEncryptionInput) operationName() string {
	return "StartDeliveryStreamEncryption"
}

type StartDeliveryStreamEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDeliveryStreamEncryptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDeliveryStreamEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDeliveryStreamEncryption{}, middleware.After)
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
	if err = addStartDeliveryStreamEncryptionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartDeliveryStreamEncryptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeliveryStreamEncryption(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDeliveryStreamEncryption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "StartDeliveryStreamEncryption",
	}
}

type opStartDeliveryStreamEncryptionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartDeliveryStreamEncryptionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartDeliveryStreamEncryptionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "firehose"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "firehose"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("firehose")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addStartDeliveryStreamEncryptionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartDeliveryStreamEncryptionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
