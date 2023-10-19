// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Evidently feature that you want to launch or test. You can define up
// to five variations of a feature, and use these variations in your launches and
// experiments. A feature must be created in a project. For information about
// creating a project, see CreateProject (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_CreateProject.html)
// . Don't use this operation to update an existing feature. Instead, use
// UpdateFeature (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_UpdateFeature.html)
// .
func (c *Client) CreateFeature(ctx context.Context, params *CreateFeatureInput, optFns ...func(*Options)) (*CreateFeatureOutput, error) {
	if params == nil {
		params = &CreateFeatureInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFeature", params, optFns, c.addOperationCreateFeatureMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFeatureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFeatureInput struct {

	// The name for the new feature.
	//
	// This member is required.
	Name *string

	// The name or ARN of the project that is to contain the new feature.
	//
	// This member is required.
	Project *string

	// An array of structures that contain the configuration of the feature's
	// different variations.
	//
	// This member is required.
	Variations []types.VariationConfig

	// The name of the variation to use as the default variation. The default
	// variation is served to users who are not allocated to any ongoing launches or
	// experiments of this feature. This variation must also be listed in the
	// variations structure. If you omit defaultVariation , the first variation listed
	// in the variations structure is used as the default variation.
	DefaultVariation *string

	// An optional description of the feature.
	Description *string

	// Specify users that should always be served a specific variation of a feature.
	// Each user is specified by a key-value pair . For each key, specify a user by
	// entering their user ID, account ID, or some other identifier. For the value,
	// specify the name of the variation that they are to be served. This parameter is
	// limited to 2500 overrides or a total of 40KB. The 40KB limit includes an
	// overhead of 6 bytes per override.
	EntityOverrides map[string]string

	// Specify ALL_RULES to activate the traffic allocation specified by any ongoing
	// launches or experiments. Specify DEFAULT_VARIATION to serve the default
	// variation to all users instead.
	EvaluationStrategy types.FeatureEvaluationStrategy

	// Assigns one or more tags (key-value pairs) to the feature. Tags can help you
	// organize and categorize your resources. You can also use them to scope user
	// permissions by granting a user permission to access or change only resources
	// with certain tag values. Tags don't have any semantic meaning to Amazon Web
	// Services and are interpreted strictly as strings of characters. You can
	// associate as many as 50 tags with a feature. For more information, see Tagging
	// Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateFeatureOutput struct {

	// A structure that contains information about the new feature.
	Feature *types.Feature

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFeatureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFeature{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFeature{}, middleware.After)
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
	if err = addOpCreateFeatureValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFeature(options.Region), middleware.Before); err != nil {
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
		operation: "CreateFeature",
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

func newServiceMetadataMiddleware_opCreateFeature(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "CreateFeature",
	}
}
