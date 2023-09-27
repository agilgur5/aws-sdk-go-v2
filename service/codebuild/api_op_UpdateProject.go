// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the settings of a build project.
func (c *Client) UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) {
	if params == nil {
		params = &UpdateProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProject", params, optFns, c.addOperationUpdateProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProjectInput struct {

	// The name of the build project. You cannot change a build project's name.
	//
	// This member is required.
	Name *string

	// Information to be changed about the build output artifacts for the build
	// project.
	Artifacts *types.ProjectArtifacts

	// Set this to true to generate a publicly accessible URL for your project's build
	// badge.
	BadgeEnabled *bool

	// Contains configuration information about a batch build project.
	BuildBatchConfig *types.ProjectBuildBatchConfig

	// Stores recently used information so that it can be quickly accessed at a later
	// time.
	Cache *types.ProjectCache

	// The maximum number of concurrent builds that are allowed for this project. New
	// builds are only started if the current number of builds is less than or equal to
	// this limit. If the current build count meets this limit, new builds are
	// throttled and are not run. To remove this limit, set this value to -1.
	ConcurrentBuildLimit *int32

	// A new or replacement description of the build project.
	Description *string

	// The Key Management Service customer master key (CMK) to be used for encrypting
	// the build output artifacts. You can use a cross-account KMS key to encrypt the
	// build output artifacts if your service role has permission to that key. You can
	// specify either the Amazon Resource Name (ARN) of the CMK or, if available, the
	// CMK's alias (using the format alias/ ).
	EncryptionKey *string

	// Information to be changed about the build environment for the build project.
	Environment *types.ProjectEnvironment

	// An array of ProjectFileSystemLocation objects for a CodeBuild build project. A
	// ProjectFileSystemLocation object specifies the identifier , location ,
	// mountOptions , mountPoint , and type of a file system created using Amazon
	// Elastic File System.
	FileSystemLocations []types.ProjectFileSystemLocation

	// Information about logs for the build project. A project can create logs in
	// CloudWatch Logs, logs in an S3 bucket, or both.
	LogsConfig *types.LogsConfig

	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes *int32

	// An array of ProjectArtifact objects.
	SecondaryArtifacts []types.ProjectArtifacts

	// An array of ProjectSourceVersion objects. If secondarySourceVersions is
	// specified at the build level, then they take over these secondarySourceVersions
	// (at the project level).
	SecondarySourceVersions []types.ProjectSourceVersion

	// An array of ProjectSource objects.
	SecondarySources []types.ProjectSource

	// The replacement ARN of the IAM role that enables CodeBuild to interact with
	// dependent Amazon Web Services services on behalf of the Amazon Web Services
	// account.
	ServiceRole *string

	// Information to be changed about the build input source code for the build
	// project.
	Source *types.ProjectSource

	// A version of the build input to be built for this project. If not specified,
	// the latest version is used. If specified, it must be one of:
	//   - For CodeCommit: the commit ID, branch, or Git tag to use.
	//   - For GitHub: the commit ID, pull request ID, branch name, or tag name that
	//   corresponds to the version of the source code you want to build. If a pull
	//   request ID is specified, it must use the format pr/pull-request-ID (for
	//   example pr/25 ). If a branch name is specified, the branch's HEAD commit ID is
	//   used. If not specified, the default branch's HEAD commit ID is used.
	//   - For Bitbucket: the commit ID, branch name, or tag name that corresponds to
	//   the version of the source code you want to build. If a branch name is specified,
	//   the branch's HEAD commit ID is used. If not specified, the default branch's HEAD
	//   commit ID is used.
	//   - For Amazon S3: the version ID of the object that represents the build input
	//   ZIP file to use.
	// If sourceVersion is specified at the build level, then that version takes
	// precedence over this sourceVersion (at the project level). For more
	// information, see Source Version Sample with CodeBuild (https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html)
	// in the CodeBuild User Guide.
	SourceVersion *string

	// An updated list of tag key and value pairs associated with this build project.
	// These tags are available for use by Amazon Web Services services that support
	// CodeBuild build project tags.
	Tags []types.Tag

	// The replacement value in minutes, from 5 to 480 (8 hours), for CodeBuild to
	// wait before timing out any related build that did not get marked as completed.
	TimeoutInMinutes *int32

	// VpcConfig enables CodeBuild to access resources in an Amazon VPC.
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

func (*UpdateProjectInput) operationName() string {
	return "UpdateProject"
}

type UpdateProjectOutput struct {

	// Information about the build project that was changed.
	Project *types.Project

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProject{}, middleware.After)
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
	if err = addUpdateProjectResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "UpdateProject",
	}
}

type opUpdateProjectResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateProjectResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateProjectResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "codebuild"
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
				signingName = "codebuild"
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
				v4aScheme.SigningName = aws.String("codebuild")
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

func addUpdateProjectResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateProjectResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
