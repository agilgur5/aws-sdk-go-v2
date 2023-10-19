// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified environment in Amazon DataZone.
func (c *Client) UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironment", params, optFns, c.addOperationUpdateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentInput struct {

	// The identifier of the domain in which the environment is to be updated.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the environment that is to be updated.
	//
	// This member is required.
	Identifier *string

	// The description to be updated as part of the UpdateEnvironment action.
	Description *string

	// The glossary terms to be updated as part of the UpdateEnvironment action.
	GlossaryTerms []string

	// The name to be updated as part of the UpdateEnvironment action.
	Name *string

	noSmithyDocumentSerde
}

type UpdateEnvironmentOutput struct {

	// The Amazon DataZone user who created the environment.
	//
	// This member is required.
	CreatedBy *string

	// The identifier of the domain in which the environment is to be updated.
	//
	// This member is required.
	DomainId *string

	// The profile identifier of the environment.
	//
	// This member is required.
	EnvironmentProfileId *string

	// The name to be updated as part of the UpdateEnvironment action.
	//
	// This member is required.
	Name *string

	// The project identifier of the environment.
	//
	// This member is required.
	ProjectId *string

	// The provider identifier of the environment.
	//
	// This member is required.
	Provider *string

	// The identifier of the Amazon Web Services account in which the environment is
	// to be updated.
	AwsAccountId *string

	// The Amazon Web Services Region in which the environment is updated.
	AwsAccountRegion *string

	// The timestamp of when the environment was created.
	CreatedAt *time.Time

	// The deployment properties to be updated as part of the UpdateEnvironment action.
	DeploymentProperties *types.DeploymentProperties

	// The description to be updated as part of the UpdateEnvironment action.
	Description *string

	// The environment actions to be updated as part of the UpdateEnvironment action.
	EnvironmentActions []types.ConfigurableEnvironmentAction

	// The blueprint identifier of the environment.
	EnvironmentBlueprintId *string

	// The glossary terms to be updated as part of the UpdateEnvironment action.
	GlossaryTerms []string

	// The identifier of the environment that is to be updated.
	Id *string

	// The last deployment of the environment.
	LastDeployment *types.Deployment

	// The provisioned resources to be updated as part of the UpdateEnvironment action.
	ProvisionedResources []types.Resource

	// The provisioning properties to be updated as part of the UpdateEnvironment
	// action.
	ProvisioningProperties types.ProvisioningProperties

	// The status to be updated as part of the UpdateEnvironment action.
	Status types.EnvironmentStatus

	// The timestamp of when the environment was updated.
	UpdatedAt *time.Time

	// The user parameters to be updated as part of the UpdateEnvironment action.
	UserParameters []types.CustomParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEnvironment{}, middleware.After)
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
	if err = addOpUpdateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnvironment(options.Region), middleware.Before); err != nil {
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
		operation: "UpdateEnvironment",
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

func newServiceMetadataMiddleware_opUpdateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datazone",
		OperationName: "UpdateEnvironment",
	}
}
