// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deploys multiple groups in one operation. This action starts the bulk
// deployment of a specified set of group versions. Each group version deployment
// will be triggered with an adaptive rate that has a fixed upper limit. We
// recommend that you include an ”X-Amzn-Client-Token” token in every
// ”StartBulkDeployment” request. These requests are idempotent with respect to
// the token and the request parameters.
func (c *Client) StartBulkDeployment(ctx context.Context, params *StartBulkDeploymentInput, optFns ...func(*Options)) (*StartBulkDeploymentOutput, error) {
	if params == nil {
		params = &StartBulkDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBulkDeployment", params, optFns, c.addOperationStartBulkDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBulkDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBulkDeploymentInput struct {

	// The ARN of the execution role to associate with the bulk deployment operation.
	// This IAM role must allow the ''greengrass:CreateDeployment'' action for all
	// group versions that are listed in the input file. This IAM role must have access
	// to the S3 bucket containing the input file.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The URI of the input file contained in the S3 bucket. The execution role must
	// have ''getObject'' permissions on this bucket to access the input file. The
	// input file is a JSON-serialized, line delimited file with UTF-8 encoding that
	// provides a list of group and version IDs and the deployment type. This file must
	// be less than 100 MB. Currently, AWS IoT Greengrass supports only
	// ''NewDeployment'' deployment types.
	//
	// This member is required.
	InputFileUri *string

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// Tag(s) to add to the new resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartBulkDeploymentOutput struct {

	// The ARN of the bulk deployment.
	BulkDeploymentArn *string

	// The ID of the bulk deployment.
	BulkDeploymentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartBulkDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartBulkDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBulkDeployment{}, middleware.After)
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
	if err = addOpStartBulkDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBulkDeployment(options.Region), middleware.Before); err != nil {
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
		operation: "StartBulkDeployment",
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

func newServiceMetadataMiddleware_opStartBulkDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "StartBulkDeployment",
	}
}
