// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a usage limit for a specified Amazon Redshift Serverless usage type.
// The usage limit is identified by the returned usage limit identifier.
func (c *Client) CreateUsageLimit(ctx context.Context, params *CreateUsageLimitInput, optFns ...func(*Options)) (*CreateUsageLimitOutput, error) {
	if params == nil {
		params = &CreateUsageLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUsageLimit", params, optFns, c.addOperationCreateUsageLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUsageLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUsageLimitInput struct {

	// The limit amount. If time-based, this amount is in Redshift Processing Units
	// (RPU) consumed per hour. If data-based, this amount is in terabytes (TB) of data
	// transferred between Regions in cross-account sharing. The value must be a
	// positive number.
	//
	// This member is required.
	Amount *int64

	// The Amazon Resource Name (ARN) of the Amazon Redshift Serverless resource to
	// create the usage limit for.
	//
	// This member is required.
	ResourceArn *string

	// The type of Amazon Redshift Serverless usage to create a usage limit for.
	//
	// This member is required.
	UsageType types.UsageLimitUsageType

	// The action that Amazon Redshift Serverless takes when the limit is reached. The
	// default is log.
	BreachAction types.UsageLimitBreachAction

	// The time period that the amount applies to. A weekly period begins on Sunday.
	// The default is monthly.
	Period types.UsageLimitPeriod

	noSmithyDocumentSerde
}

type CreateUsageLimitOutput struct {

	// The returned usage limit object.
	UsageLimit *types.UsageLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUsageLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUsageLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUsageLimit{}, middleware.After)
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
	if err = addOpCreateUsageLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUsageLimit(options.Region), middleware.Before); err != nil {
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
		operation: "CreateUsageLimit",
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

func newServiceMetadataMiddleware_opCreateUsageLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-serverless",
		OperationName: "CreateUsageLimit",
	}
}
