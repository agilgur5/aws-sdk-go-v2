// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Dev Environment in Amazon CodeCatalyst, a cloud-based development
// environment that you can use to quickly work on the code stored in the source
// repositories of your project. When created in the Amazon CodeCatalyst console,
// by default a Dev Environment is configured to have a 2 core processor, 4GB of
// RAM, and 16GB of persistent storage. None of these defaults apply to a Dev
// Environment created programmatically.
func (c *Client) CreateDevEnvironment(ctx context.Context, params *CreateDevEnvironmentInput, optFns ...func(*Options)) (*CreateDevEnvironmentOutput, error) {
	if params == nil {
		params = &CreateDevEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDevEnvironment", params, optFns, c.addOperationCreateDevEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDevEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDevEnvironmentInput struct {

	// The Amazon EC2 instace type to use for the Dev Environment.
	//
	// This member is required.
	InstanceType types.InstanceType

	// Information about the amount of storage allocated to the Dev Environment. By
	// default, a Dev Environment is configured to have 16GB of persistent storage when
	// created from the Amazon CodeCatalyst console, but there is no default when
	// programmatically creating a Dev Environment. Valid values for persistent storage
	// are based on memory sizes in 16GB increments. Valid values are 16, 32, and 64.
	//
	// This member is required.
	PersistentStorage *types.PersistentStorageConfiguration

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The user-defined alias for a Dev Environment.
	Alias *string

	// A user-specified idempotency token. Idempotency ensures that an API request
	// completes only once. With an idempotent request, if the original request
	// completes successfully, the subsequent retries return the result from the
	// original successful request and have no additional effect.
	ClientToken *string

	// Information about the integrated development environment (IDE) configured for a
	// Dev Environment. An IDE is required to create a Dev Environment. For Dev
	// Environment creation, this field contains configuration information and must be
	// provided.
	Ides []types.IdeConfiguration

	// The amount of time the Dev Environment will run without any activity detected
	// before stopping, in minutes. Only whole integers are allowed. Dev Environments
	// consume compute minutes when running.
	InactivityTimeoutMinutes int32

	// The source repository that contains the branch to clone into the Dev
	// Environment.
	Repositories []types.RepositoryInput

	noSmithyDocumentSerde
}

type CreateDevEnvironmentOutput struct {

	// The system-generated unique ID of the Dev Environment.
	//
	// This member is required.
	Id *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDevEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDevEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDevEnvironment{}, middleware.After)
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
	if err = addBearerAuthSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDevEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDevEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDevEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDevEnvironment",
	}
}
