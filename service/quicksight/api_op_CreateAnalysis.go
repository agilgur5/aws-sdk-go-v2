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

// Creates an analysis in Amazon QuickSight. Analyses can be created either from a
// template or from an AnalysisDefinition .
func (c *Client) CreateAnalysis(ctx context.Context, params *CreateAnalysisInput, optFns ...func(*Options)) (*CreateAnalysisOutput, error) {
	if params == nil {
		params = &CreateAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAnalysis", params, optFns, c.addOperationCreateAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAnalysisInput struct {

	// The ID for the analysis that you're creating. This ID displays in the URL of
	// the analysis.
	//
	// This member is required.
	AnalysisId *string

	// The ID of the Amazon Web Services account where you are creating an analysis.
	//
	// This member is required.
	AwsAccountId *string

	// A descriptive name for the analysis that you're creating. This name displays
	// for the analysis in the Amazon QuickSight console.
	//
	// This member is required.
	Name *string

	// The definition of an analysis. A definition is the data model of all features
	// in a Dashboard, Template, or Analysis. Either a SourceEntity or a Definition
	// must be provided in order for the request to be valid.
	Definition *types.AnalysisDefinition

	// When you create the analysis, Amazon QuickSight adds the analysis to these
	// folders.
	FolderArns []string

	// The parameter names and override values that you want to use. An analysis can
	// have any parameter type, and some parameters might accept multiple values.
	Parameters *types.Parameters

	// A structure that describes the principals and the resource-level permissions on
	// an analysis. You can use the Permissions structure to grant permissions by
	// providing a list of Identity and Access Management (IAM) action information for
	// each principal listed by Amazon Resource Name (ARN). To specify no permissions,
	// omit Permissions .
	Permissions []types.ResourcePermission

	// A source entity to use for the analysis that you're creating. This metadata
	// structure contains details that describe a source template and one or more
	// datasets. Either a SourceEntity or a Definition must be provided in order for
	// the request to be valid.
	SourceEntity *types.AnalysisSourceEntity

	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the analysis.
	Tags []types.Tag

	// The ARN for the theme to apply to the analysis that you're creating. To see the
	// theme in the Amazon QuickSight console, make sure that you have access to it.
	ThemeArn *string

	// The option to relax the validation needed to create an analysis with definition
	// objects. This skips the validation step for specific errors.
	ValidationStrategy *types.ValidationStrategy

	noSmithyDocumentSerde
}

type CreateAnalysisOutput struct {

	// The ID of the analysis.
	AnalysisId *string

	// The ARN for the analysis.
	Arn *string

	// The status of the creation of the analysis.
	CreationStatus types.ResourceStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAnalysis{}, middleware.After)
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
	if err = addOpCreateAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateAnalysis",
	}
}
