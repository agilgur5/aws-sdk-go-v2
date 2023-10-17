// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes details about an Resilience Hub application.
func (c *Client) DescribeAppVersionTemplate(ctx context.Context, params *DescribeAppVersionTemplateInput, optFns ...func(*Options)) (*DescribeAppVersionTemplateOutput, error) {
	if params == nil {
		params = &DescribeAppVersionTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppVersionTemplate", params, optFns, c.addOperationDescribeAppVersionTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppVersionTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppVersionTemplateInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The version of the application.
	//
	// This member is required.
	AppVersion *string

	noSmithyDocumentSerde
}

type DescribeAppVersionTemplateOutput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// A JSON string that provides information about your application structure. To
	// learn more about the appTemplateBody template, see the sample template provided
	// in the Examples section. The appTemplateBody JSON string has the following
	// structure:
	//   - resources The list of logical resources that must be included in the
	//   Resilience Hub application. Type: Array Don't add the resources that you want to
	//   exclude. Each resources array item includes the following fields:
	//   - logicalResourceId Logical identifier of the resource. Type: Object Each
	//   logicalResourceId object includes the following fields:
	//   - identifier Identifier of the resource. Type: String
	//   - logicalStackName The name of the CloudFormation stack this resource belongs
	//   to. Type: String
	//   - resourceGroupName The name of the resource group this resource belongs to.
	//   Type: String
	//   - terraformSourceName The name of the Terraform S3 state file this resource
	//   belongs to. Type: String
	//   - eksSourceName Name of the Amazon Elastic Kubernetes Service cluster and
	//   namespace this resource belongs to. This parameter accepts values in
	//   "eks-cluster/namespace" format. Type: String
	//   - type The type of resource. Type: string
	//   - name The name of the resource. Type: String
	//   - additionalInfo Additional configuration parameters for an Resilience Hub
	//   application. If you want to implement additionalInfo through the Resilience
	//   Hub console rather than using an API call, see Configure the application
	//   configuration parameters (https://docs.aws.amazon.com/resilience-hub/latest/userguide/app-config-param.html)
	//   . Currently, this parameter accepts a key-value mapping (in a string format) of
	//   only one failover region and one associated account. Key: "failover-regions"
	//   Value: "[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"
	//   - appComponents List of Application Components that this resource belongs to.
	//   If an Application Component is not part of the Resilience Hub application, it
	//   will be added. Type: Array Each appComponents array item includes the
	//   following fields:
	//   - name Name of the Application Component. Type: String
	//   - type Type of Application Component. For more information about the types of
	//   Application Component, see Grouping resources in an AppComponent (https://docs.aws.amazon.com/resilience-hub/latest/userguide/AppComponent.grouping.html)
	//   . Type: String
	//   - resourceNames The list of included resources that are assigned to the
	//   Application Component. Type: Array of strings
	//   - additionalInfo Additional configuration parameters for an Resilience Hub
	//   application. If you want to implement additionalInfo through the Resilience
	//   Hub console rather than using an API call, see Configure the application
	//   configuration parameters (https://docs.aws.amazon.com/resilience-hub/latest/userguide/app-config-param.html)
	//   . Currently, this parameter accepts a key-value mapping (in a string format) of
	//   only one failover region and one associated account. Key: "failover-regions"
	//   Value: "[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"
	//   - excludedResources The list of logical resource identifiers to be excluded
	//   from the application. Type: Array Don't add the resources that you want to
	//   include. Each excludedResources array item includes the following fields:
	//   - logicalResourceIds Logical identifier of the resource. Type: Object You can
	//   configure only one of the following fields:
	//   - logicalStackName
	//   - resourceGroupName
	//   - terraformSourceName
	//   - eksSourceName Each logicalResourceIds object includes the following fields:
	//   - identifier Identifier of the resource. Type: String
	//   - logicalStackName The name of the CloudFormation stack this resource belongs
	//   to. Type: String
	//   - resourceGroupName The name of the resource group this resource belongs to.
	//   Type: String
	//   - terraformSourceName The name of the Terraform S3 state file this resource
	//   belongs to. Type: String
	//   - eksSourceName Name of the Amazon Elastic Kubernetes Service cluster and
	//   namespace this resource belongs to. This parameter accepts values in
	//   "eks-cluster/namespace" format. Type: String
	//   - version Resilience Hub application version.
	//   - additionalInfo Additional configuration parameters for an Resilience Hub
	//   application. If you want to implement additionalInfo through the Resilience
	//   Hub console rather than using an API call, see Configure the application
	//   configuration parameters (https://docs.aws.amazon.com/resilience-hub/latest/userguide/app-config-param.html)
	//   . Currently, this parameter accepts a key-value mapping (in a string format) of
	//   only one failover region and one associated account. Key: "failover-regions"
	//   Value: "[{"region":"<REGION>", "accounts":[{"id":"<ACCOUNT_ID>"}]}]"
	//
	// This member is required.
	AppTemplateBody *string

	// The version of the application.
	//
	// This member is required.
	AppVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppVersionTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAppVersionTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAppVersionTemplate{}, middleware.After)
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
	if err = addOpDescribeAppVersionTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppVersionTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAppVersionTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "DescribeAppVersionTemplate",
	}
}
