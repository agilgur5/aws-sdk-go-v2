// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves all the information for the specified kdb environment.
func (c *Client) GetKxEnvironment(ctx context.Context, params *GetKxEnvironmentInput, optFns ...func(*Options)) (*GetKxEnvironmentOutput, error) {
	if params == nil {
		params = &GetKxEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetKxEnvironment", params, optFns, c.addOperationGetKxEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetKxEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKxEnvironmentInput struct {

	// A unique identifier for the kdb environment.
	//
	// This member is required.
	EnvironmentId *string

	noSmithyDocumentSerde
}

type GetKxEnvironmentOutput struct {

	// The identifier of the availability zones where subnets for the environment are
	// created.
	AvailabilityZoneIds []string

	// The unique identifier of the AWS account that is used to create the kdb
	// environment.
	AwsAccountId *string

	// The Amazon Resource Name (ARN) of the certificate authority of the kdb
	// environment.
	CertificateAuthorityArn *string

	// The timestamp at which the kdb environment was created in FinSpace.
	CreationTimestamp *time.Time

	// A list of DNS server name and server IP. This is used to set up Route-53
	// outbound resolvers.
	CustomDNSConfiguration []types.CustomDNSServer

	// A unique identifier for the AWS environment infrastructure account.
	DedicatedServiceAccountId *string

	// A description for the kdb environment.
	Description *string

	// The status of DNS configuration.
	DnsStatus types.DnsStatus

	// The ARN identifier of the environment.
	EnvironmentArn *string

	// A unique identifier for the kdb environment.
	EnvironmentId *string

	// Specifies the error message that appears if a flow fails.
	ErrorMessage *string

	// The KMS key ID to encrypt your data in the FinSpace environment.
	KmsKeyId *string

	// The name of the kdb environment.
	Name *string

	// The status of the kdb environment.
	Status types.EnvironmentStatus

	// The status of the network configuration.
	TgwStatus types.TgwStatus

	// The structure of the transit gateway and network configuration that is used to
	// connect the kdb environment to an internal network.
	TransitGatewayConfiguration *types.TransitGatewayConfiguration

	// The timestamp at which the kdb environment was updated.
	UpdateTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetKxEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetKxEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetKxEnvironment{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpGetKxEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetKxEnvironment(options.Region), middleware.Before); err != nil {
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
		operation: "GetKxEnvironment",
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

func newServiceMetadataMiddleware_opGetKxEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace",
		OperationName: "GetKxEnvironment",
	}
}
