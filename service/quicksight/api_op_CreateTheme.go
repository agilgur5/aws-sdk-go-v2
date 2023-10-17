// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a theme. A theme is set of configuration options for color and layout.
// Themes apply to analyses and dashboards. For more information, see Using Themes
// in Amazon QuickSight (https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html)
// in the Amazon QuickSight User Guide.
func (c *Client) CreateTheme(ctx context.Context, params *CreateThemeInput, optFns ...func(*Options)) (*CreateThemeOutput, error) {
	if params == nil {
		params = &CreateThemeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTheme", params, optFns, c.addOperationCreateThemeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateThemeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateThemeInput struct {

	// The ID of the Amazon Web Services account where you want to store the new theme.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the theme that a custom theme will inherit from. All themes inherit
	// from one of the starting themes defined by Amazon QuickSight. For a list of the
	// starting themes, use ListThemes or choose Themes from within an analysis.
	//
	// This member is required.
	BaseThemeId *string

	// The theme configuration, which contains the theme display properties.
	//
	// This member is required.
	Configuration *types.ThemeConfiguration

	// A display name for the theme.
	//
	// This member is required.
	Name *string

	// An ID for the theme that you want to create. The theme ID is unique per Amazon
	// Web Services Region in each Amazon Web Services account.
	//
	// This member is required.
	ThemeId *string

	// A valid grouping of resource permissions to apply to the new theme.
	Permissions []types.ResourcePermission

	// A map of the key-value pairs for the resource tag or tags that you want to add
	// to the resource.
	Tags []types.Tag

	// A description of the first version of the theme that you're creating. Every
	// time UpdateTheme is called, a new version is created. Each version of the theme
	// has a description of the version in the VersionDescription field.
	VersionDescription *string

	noSmithyDocumentSerde
}

type CreateThemeOutput struct {

	// The Amazon Resource Name (ARN) for the theme.
	Arn *string

	// The theme creation status.
	CreationStatus types.ResourceStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ID of the theme.
	ThemeId *string

	// The Amazon Resource Name (ARN) for the new theme.
	VersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateThemeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTheme{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTheme{}, middleware.After)
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
	if err = addOpCreateThemeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTheme(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTheme(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateTheme",
	}
}
