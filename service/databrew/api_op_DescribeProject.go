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

// Returns the definition of a specific DataBrew project.
func (c *Client) DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) {
	if params == nil {
		params = &DescribeProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProject", params, optFns, c.addOperationDescribeProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProjectInput struct {

	// The name of the project to be described.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeProjectOutput struct {

	// The name of the project.
	//
	// This member is required.
	Name *string

	// The date and time that the project was created.
	CreateDate *time.Time

	// The identifier (user name) of the user who created the project.
	CreatedBy *string

	// The dataset associated with the project.
	DatasetName *string

	// The identifier (user name) of the user who last modified the project.
	LastModifiedBy *string

	// The date and time that the project was last modified.
	LastModifiedDate *time.Time

	// The date and time when the project was opened.
	OpenDate *time.Time

	// The identifier (user name) of the user that opened the project for use.
	OpenedBy *string

	// The recipe associated with this job.
	RecipeName *string

	// The Amazon Resource Name (ARN) of the project.
	ResourceArn *string

	// The ARN of the Identity and Access Management (IAM) role to be assumed when
	// DataBrew runs the job.
	RoleArn *string

	// Represents the sample size and sampling type for DataBrew to use for
	// interactive data analysis.
	Sample *types.Sample

	// Describes the current state of the session:
	//   - PROVISIONING - allocating resources for the session.
	//   - INITIALIZING - getting the session ready for first use.
	//   - ASSIGNED - the session is ready for use.
	SessionStatus types.SessionStatus

	// Metadata tags associated with this project.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProject{}, middleware.After)
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
	if err = addOpDescribeProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProject(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeProject",
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

func newServiceMetadataMiddleware_opDescribeProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "DescribeProject",
	}
}
