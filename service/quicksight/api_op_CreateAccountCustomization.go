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

// Creates Amazon QuickSight customizations for the current Amazon Web Services
// Region. Currently, you can add a custom default theme by using the
// CreateAccountCustomization or UpdateAccountCustomization API operation. To
// further customize Amazon QuickSight by removing Amazon QuickSight sample assets
// and videos for all new users, see Customizing Amazon QuickSight (https://docs.aws.amazon.com/quicksight/latest/user/customizing-quicksight.html)
// in the Amazon QuickSight User Guide. You can create customizations for your
// Amazon Web Services account or, if you specify a namespace, for a QuickSight
// namespace instead. Customizations that apply to a namespace always override
// customizations that apply to an Amazon Web Services account. To find out which
// customizations apply, use the DescribeAccountCustomization API operation.
// Before you use the CreateAccountCustomization API operation to add a theme as
// the namespace default, make sure that you first share the theme with the
// namespace. If you don't share it with the namespace, the theme isn't visible to
// your users even if you make it the default theme. To check if the theme is
// shared, view the current permissions by using the DescribeThemePermissions (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeThemePermissions.html)
// API operation. To share the theme, grant permissions by using the
// UpdateThemePermissions (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateThemePermissions.html)
// API operation.
func (c *Client) CreateAccountCustomization(ctx context.Context, params *CreateAccountCustomizationInput, optFns ...func(*Options)) (*CreateAccountCustomizationOutput, error) {
	if params == nil {
		params = &CreateAccountCustomizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccountCustomization", params, optFns, c.addOperationCreateAccountCustomizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccountCustomizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccountCustomizationInput struct {

	// The Amazon QuickSight customizations you're adding in the current Amazon Web
	// Services Region. You can add these to an Amazon Web Services account and a
	// QuickSight namespace. For example, you can add a default theme by setting
	// AccountCustomization to the midnight theme: "AccountCustomization": {
	// "DefaultTheme": "arn:aws:quicksight::aws:theme/MIDNIGHT" } . Or, you can add a
	// custom theme by specifying "AccountCustomization": { "DefaultTheme":
	// "arn:aws:quicksight:us-west-2:111122223333:theme/bdb844d0-0fe9-4d9d-b520-0fe602d93639"
	// } .
	//
	// This member is required.
	AccountCustomization *types.AccountCustomization

	// The ID for the Amazon Web Services account that you want to customize Amazon
	// QuickSight for.
	//
	// This member is required.
	AwsAccountId *string

	// The Amazon QuickSight namespace that you want to add customizations to.
	Namespace *string

	// A list of the tags that you want to attach to this resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAccountCustomizationOutput struct {

	// The Amazon QuickSight customizations you're adding in the current Amazon Web
	// Services Region.
	AccountCustomization *types.AccountCustomization

	// The Amazon Resource Name (ARN) for the customization that you created for this
	// Amazon Web Services account.
	Arn *string

	// The ID for the Amazon Web Services account that you want to customize Amazon
	// QuickSight for.
	AwsAccountId *string

	// The namespace associated with the customization you're creating.
	Namespace *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccountCustomizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccountCustomization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccountCustomization{}, middleware.After)
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
	if err = addOpCreateAccountCustomizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccountCustomization(options.Region), middleware.Before); err != nil {
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
		operation: "CreateAccountCustomization",
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

func newServiceMetadataMiddleware_opCreateAccountCustomization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateAccountCustomization",
	}
}
