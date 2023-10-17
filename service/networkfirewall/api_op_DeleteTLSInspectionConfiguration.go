// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified TLSInspectionConfiguration .
func (c *Client) DeleteTLSInspectionConfiguration(ctx context.Context, params *DeleteTLSInspectionConfigurationInput, optFns ...func(*Options)) (*DeleteTLSInspectionConfigurationOutput, error) {
	if params == nil {
		params = &DeleteTLSInspectionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTLSInspectionConfiguration", params, optFns, c.addOperationDeleteTLSInspectionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTLSInspectionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTLSInspectionConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the TLS inspection configuration. You must
	// specify the ARN or the name, and you can specify both.
	TLSInspectionConfigurationArn *string

	// The descriptive name of the TLS inspection configuration. You can't change the
	// name of a TLS inspection configuration after you create it. You must specify the
	// ARN or the name, and you can specify both.
	TLSInspectionConfigurationName *string

	noSmithyDocumentSerde
}

type DeleteTLSInspectionConfigurationOutput struct {

	// The high-level properties of a TLS inspection configuration. This, along with
	// the TLSInspectionConfiguration , define the TLS inspection configuration. You
	// can retrieve all objects for a TLS inspection configuration by calling
	// DescribeTLSInspectionConfiguration .
	//
	// This member is required.
	TLSInspectionConfigurationResponse *types.TLSInspectionConfigurationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTLSInspectionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteTLSInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteTLSInspectionConfiguration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTLSInspectionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTLSInspectionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "DeleteTLSInspectionConfiguration",
	}
}
