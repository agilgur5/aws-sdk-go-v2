// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the list of tags applied to an AWS Elastic Beanstalk resource. Two lists
// can be passed: TagsToAdd for tags to add or update, and TagsToRemove . Elastic
// Beanstalk supports tagging of all of its resources. For details about resource
// tagging, see Tagging Application Resources (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/applications-tagging-resources.html)
// . If you create a custom IAM user policy to control permission to this
// operation, specify one of the following two virtual actions (or both) instead of
// the API operation name: elasticbeanstalk:AddTags Controls permission to call
// UpdateTagsForResource and pass a list of tags to add in the TagsToAdd
// parameter. elasticbeanstalk:RemoveTags Controls permission to call
// UpdateTagsForResource and pass a list of tag keys to remove in the TagsToRemove
// parameter. For details about creating a custom user policy, see Creating a
// Custom User Policy (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/AWSHowTo.iam.managed-policies.html#AWSHowTo.iam.policies)
// .
func (c *Client) UpdateTagsForResource(ctx context.Context, params *UpdateTagsForResourceInput, optFns ...func(*Options)) (*UpdateTagsForResourceOutput, error) {
	if params == nil {
		params = &UpdateTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTagsForResource", params, optFns, c.addOperationUpdateTagsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTagsForResourceInput struct {

	// The Amazon Resource Name (ARN) of the resouce to be updated. Must be the ARN of
	// an Elastic Beanstalk resource.
	//
	// This member is required.
	ResourceArn *string

	// A list of tags to add or update. If a key of an existing tag is added, the
	// tag's value is updated. Specify at least one of these parameters: TagsToAdd ,
	// TagsToRemove .
	TagsToAdd []types.Tag

	// A list of tag keys to remove. If a tag key doesn't exist, it is silently
	// ignored. Specify at least one of these parameters: TagsToAdd , TagsToRemove .
	TagsToRemove []string

	noSmithyDocumentSerde
}

type UpdateTagsForResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateTagsForResource{}, middleware.After)
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
	if err = addOpUpdateTagsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTagsForResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTagsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "UpdateTagsForResource",
	}
}
