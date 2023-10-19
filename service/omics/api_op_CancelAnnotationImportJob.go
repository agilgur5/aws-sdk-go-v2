// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels an annotation import job.
func (c *Client) CancelAnnotationImportJob(ctx context.Context, params *CancelAnnotationImportJobInput, optFns ...func(*Options)) (*CancelAnnotationImportJobOutput, error) {
	if params == nil {
		params = &CancelAnnotationImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelAnnotationImportJob", params, optFns, c.addOperationCancelAnnotationImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelAnnotationImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelAnnotationImportJobInput struct {

	// The job's ID.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type CancelAnnotationImportJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelAnnotationImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelAnnotationImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelAnnotationImportJob{}, middleware.After)
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
	if err = addEndpointPrefix_opCancelAnnotationImportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpCancelAnnotationImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelAnnotationImportJob(options.Region), middleware.Before); err != nil {
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
		operation: "CancelAnnotationImportJob",
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

type endpointPrefix_opCancelAnnotationImportJobMiddleware struct {
}

func (*endpointPrefix_opCancelAnnotationImportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCancelAnnotationImportJobMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCancelAnnotationImportJobMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCancelAnnotationImportJobMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCancelAnnotationImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "CancelAnnotationImportJob",
	}
}
