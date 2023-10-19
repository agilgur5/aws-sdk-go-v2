// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an existing Amazon Kendra index.
func (c *Client) DescribeIndex(ctx context.Context, params *DescribeIndexInput, optFns ...func(*Options)) (*DescribeIndexOutput, error) {
	if params == nil {
		params = &DescribeIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIndex", params, optFns, c.addOperationDescribeIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIndexInput struct {

	// The identifier of the index you want to get information on.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DescribeIndexOutput struct {

	// For Enterprise Edition indexes, you can choose to use additional capacity to
	// meet the needs of your application. This contains the capacity units used for
	// the index. A query or document storage capacity of zero indicates that the index
	// is using the default capacity. For more information on the default capacity for
	// an index and adjusting this, see Adjusting capacity (https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html)
	// .
	CapacityUnits *types.CapacityUnitsConfiguration

	// The Unix timestamp when the index was created.
	CreatedAt *time.Time

	// The description for the index.
	Description *string

	// Configuration information for document metadata or fields. Document metadata
	// are fields or attributes associated with your documents. For example, the
	// company department name associated with each document.
	DocumentMetadataConfigurations []types.DocumentMetadataConfiguration

	// The Amazon Kendra edition used for the index. You decide the edition when you
	// create the index.
	Edition types.IndexEdition

	// When the Status field value is FAILED , the ErrorMessage field contains a
	// message that explains why.
	ErrorMessage *string

	// The identifier of the index.
	Id *string

	// Provides information about the number of FAQ questions and answers and the
	// number of text documents indexed.
	IndexStatistics *types.IndexStatistics

	// The name of the index.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role that gives Amazon Kendra
	// permission to write to your Amazon Cloudwatch logs.
	RoleArn *string

	// The identifier of the KMScustomer master key (CMK) that is used to encrypt your
	// data. Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The current status of the index. When the value is ACTIVE , the index is ready
	// for use. If the Status field value is FAILED , the ErrorMessage field contains
	// a message that explains why.
	Status types.IndexStatus

	// The Unix when the index was last updated.
	UpdatedAt *time.Time

	// The user context policy for the Amazon Kendra index.
	UserContextPolicy types.UserContextPolicy

	// Whether you have enabled the configuration for fetching access levels of groups
	// and users from an IAM Identity Center identity source.
	UserGroupResolutionConfiguration *types.UserGroupResolutionConfiguration

	// The user token configuration for the Amazon Kendra index.
	UserTokenConfigurations []types.UserTokenConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeIndex{}, middleware.After)
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
	if err = addOpDescribeIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIndex(options.Region), middleware.Before); err != nil {
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
		operation: "DescribeIndex",
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

func newServiceMetadataMiddleware_opDescribeIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeIndex",
	}
}
