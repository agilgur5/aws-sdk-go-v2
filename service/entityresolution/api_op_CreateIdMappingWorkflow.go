// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an IdMappingWorkflow object which stores the configuration of the data
// processing job to be run. Each IdMappingWorkflow must have a unique workflow
// name. To modify an existing workflow, use the UpdateIdMappingWorkflow API.
func (c *Client) CreateIdMappingWorkflow(ctx context.Context, params *CreateIdMappingWorkflowInput, optFns ...func(*Options)) (*CreateIdMappingWorkflowOutput, error) {
	if params == nil {
		params = &CreateIdMappingWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIdMappingWorkflow", params, optFns, c.addOperationCreateIdMappingWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIdMappingWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIdMappingWorkflowInput struct {

	// An object which defines the idMappingType and the providerProperties .
	//
	// This member is required.
	IdMappingTechniques *types.IdMappingTechniques

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.IdMappingWorkflowInputSource

	// A list of IdMappingWorkflowOutputSource objects, each of which contains fields
	// OutputS3Path and Output .
	//
	// This member is required.
	OutputSourceConfig []types.IdMappingWorkflowOutputSource

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to create resources on your behalf as part of workflow execution.
	//
	// This member is required.
	RoleArn *string

	// The name of the workflow. There can't be multiple IdMappingWorkflows with the
	// same name.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateIdMappingWorkflowOutput struct {

	// An object which defines the idMappingType and the providerProperties .
	//
	// This member is required.
	IdMappingTechniques *types.IdMappingTechniques

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.IdMappingWorkflowInputSource

	// A list of IdMappingWorkflowOutputSource objects, each of which contains fields
	// OutputS3Path and Output .
	//
	// This member is required.
	OutputSourceConfig []types.IdMappingWorkflowOutputSource

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to create resources on your behalf as part of workflow execution.
	//
	// This member is required.
	RoleArn *string

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// IDMappingWorkflow .
	//
	// This member is required.
	WorkflowArn *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIdMappingWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIdMappingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIdMappingWorkflow{}, middleware.After)
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
	if err = addCreateIdMappingWorkflowResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateIdMappingWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIdMappingWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIdMappingWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "entityresolution",
		OperationName: "CreateIdMappingWorkflow",
	}
}

type opCreateIdMappingWorkflowResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateIdMappingWorkflowResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateIdMappingWorkflowResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "entityresolution"
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
				signingName = "entityresolution"
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
				v4aScheme.SigningName = aws.String("entityresolution")
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

func addCreateIdMappingWorkflowResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateIdMappingWorkflowResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
