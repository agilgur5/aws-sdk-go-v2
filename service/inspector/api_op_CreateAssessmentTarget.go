// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new assessment target using the ARN of the resource group that is
// generated by CreateResourceGroup . If resourceGroupArn is not specified, all EC2
// instances in the current AWS account and region are included in the assessment
// target. If the service-linked role (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html)
// isn’t already registered, this action also creates and registers a
// service-linked role to grant Amazon Inspector access to AWS Services needed to
// perform security assessments. You can create up to 50 assessment targets per AWS
// account. You can run up to 500 concurrent agents per AWS account. For more
// information, see Amazon Inspector Assessment Targets (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_applications.html)
// .
func (c *Client) CreateAssessmentTarget(ctx context.Context, params *CreateAssessmentTargetInput, optFns ...func(*Options)) (*CreateAssessmentTargetOutput, error) {
	if params == nil {
		params = &CreateAssessmentTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssessmentTarget", params, optFns, c.addOperationCreateAssessmentTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssessmentTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssessmentTargetInput struct {

	// The user-defined name that identifies the assessment target that you want to
	// create. The name must be unique within the AWS account.
	//
	// This member is required.
	AssessmentTargetName *string

	// The ARN that specifies the resource group that is used to create the assessment
	// target. If resourceGroupArn is not specified, all EC2 instances in the current
	// AWS account and region are included in the assessment target.
	ResourceGroupArn *string

	noSmithyDocumentSerde
}

type CreateAssessmentTargetOutput struct {

	// The ARN that specifies the assessment target that is created.
	//
	// This member is required.
	AssessmentTargetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssessmentTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAssessmentTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAssessmentTarget{}, middleware.After)
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
	if err = addOpCreateAssessmentTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssessmentTarget(options.Region), middleware.Before); err != nil {
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
		operation: "CreateAssessmentTarget",
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

func newServiceMetadataMiddleware_opCreateAssessmentTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "CreateAssessmentTarget",
	}
}
