// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an OpenSearch Serverless-managed interface VPC endpoint. For more
// information, see Access Amazon OpenSearch Serverless using an interface endpoint (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html)
// .
func (c *Client) CreateVpcEndpoint(ctx context.Context, params *CreateVpcEndpointInput, optFns ...func(*Options)) (*CreateVpcEndpointOutput, error) {
	if params == nil {
		params = &CreateVpcEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcEndpoint", params, optFns, c.addOperationCreateVpcEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVpcEndpointInput struct {

	// The name of the interface endpoint.
	//
	// This member is required.
	Name *string

	// The ID of one or more subnets from which you'll access OpenSearch Serverless.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the VPC from which you'll access OpenSearch Serverless.
	//
	// This member is required.
	VpcId *string

	// Unique, case-sensitive identifier to ensure idempotency of the request.
	ClientToken *string

	// The unique identifiers of the security groups that define the ports, protocols,
	// and sources for inbound traffic that you are authorizing into your endpoint.
	SecurityGroupIds []string

	noSmithyDocumentSerde
}

type CreateVpcEndpointOutput struct {

	// Details about the created interface VPC endpoint.
	CreateVpcEndpointDetail *types.CreateVpcEndpointDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVpcEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateVpcEndpoint{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateVpcEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateVpcEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcEndpoint(options.Region), middleware.Before); err != nil {
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
		operation: "CreateVpcEndpoint",
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

type idempotencyToken_initializeOpCreateVpcEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateVpcEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateVpcEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateVpcEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateVpcEndpointInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateVpcEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateVpcEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateVpcEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aoss",
		OperationName: "CreateVpcEndpoint",
	}
}
