// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the definition of a specific DataBrew recipe corresponding to a
// particular version.
func (c *Client) DescribeRecipe(ctx context.Context, params *DescribeRecipeInput, optFns ...func(*Options)) (*DescribeRecipeOutput, error) {
	if params == nil {
		params = &DescribeRecipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecipe", params, optFns, c.addOperationDescribeRecipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecipeInput struct {

	// The name of the recipe to be described.
	//
	// This member is required.
	Name *string

	// The recipe version identifier. If this parameter isn't specified, then the
	// latest published version is returned.
	RecipeVersion *string

	noSmithyDocumentSerde
}

type DescribeRecipeOutput struct {

	// The name of the recipe.
	//
	// This member is required.
	Name *string

	// The date and time that the recipe was created.
	CreateDate *time.Time

	// The identifier (user name) of the user who created the recipe.
	CreatedBy *string

	// The description of the recipe.
	Description *string

	// The identifier (user name) of the user who last modified the recipe.
	LastModifiedBy *string

	// The date and time that the recipe was last modified.
	LastModifiedDate *time.Time

	// The name of the project associated with this recipe.
	ProjectName *string

	// The identifier (user name) of the user who last published the recipe.
	PublishedBy *string

	// The date and time when the recipe was last published.
	PublishedDate *time.Time

	// The recipe version identifier.
	RecipeVersion *string

	// The ARN of the recipe.
	ResourceArn *string

	// One or more steps to be performed by the recipe. Each step consists of an
	// action, and the conditions under which the action should succeed.
	Steps []types.RecipeStep

	// Metadata tags associated with this project.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRecipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRecipe{}, middleware.After)
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
	if err = addOpDescribeRecipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecipe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRecipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "DescribeRecipe",
	}
}
