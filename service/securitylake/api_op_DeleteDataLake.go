// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// When you disable Amazon Security Lake from your account, Security Lake is
// disabled in all Amazon Web Services Regions and it stops collecting data from
// your sources. Also, this API automatically takes steps to remove the account
// from Security Lake. However, Security Lake retains all of your existing settings
// and the resources that it created in your Amazon Web Services account in the
// current Amazon Web Services Region. The DeleteDataLake operation does not
// delete the data that is stored in your Amazon S3 bucket, which is owned by your
// Amazon Web Services account. For more information, see the Amazon Security Lake
// User Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/disable-security-lake.html)
// .
func (c *Client) DeleteDataLake(ctx context.Context, params *DeleteDataLakeInput, optFns ...func(*Options)) (*DeleteDataLakeOutput, error) {
	if params == nil {
		params = &DeleteDataLakeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataLake", params, optFns, c.addOperationDeleteDataLakeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDataLakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDataLakeInput struct {

	// The list of Regions where Security Lake is enabled.
	//
	// This member is required.
	Regions []string

	noSmithyDocumentSerde
}

type DeleteDataLakeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDataLakeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataLake{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataLake{}, middleware.After)
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
	if err = addOpDeleteDataLakeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataLake(options.Region), middleware.Before); err != nil {
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
		operation: "DeleteDataLake",
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

func newServiceMetadataMiddleware_opDeleteDataLake(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "DeleteDataLake",
	}
}
