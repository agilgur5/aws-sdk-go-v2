// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a stack as specified in the template. After the call completes
// successfully, the stack creation starts. You can check the status of the stack
// through the DescribeStacks operation.
func (c *Client) CreateStack(ctx context.Context, params *CreateStackInput, optFns ...func(*Options)) (*CreateStackOutput, error) {
	if params == nil {
		params = &CreateStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStack", params, optFns, c.addOperationCreateStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for CreateStack action.
type CreateStackInput struct {

	// The name that's associated with the stack. The name must be unique in the
	// Region in which you are creating the stack. A stack name can contain only
	// alphanumeric characters (case sensitive) and hyphens. It must start with an
	// alphabetical character and can't be longer than 128 characters.
	//
	// This member is required.
	StackName *string

	// In some cases, you must explicitly acknowledge that your stack template
	// contains certain capabilities in order for CloudFormation to create the stack.
	//   - CAPABILITY_IAM and CAPABILITY_NAMED_IAM Some stack templates might include
	//   resources that can affect permissions in your Amazon Web Services account; for
	//   example, by creating new Identity and Access Management (IAM) users. For those
	//   stacks, you must explicitly acknowledge this by specifying one of these
	//   capabilities. The following IAM resources require you to specify either the
	//   CAPABILITY_IAM or CAPABILITY_NAMED_IAM capability.
	//   - If you have IAM resources, you can specify either capability.
	//   - If you have IAM resources with custom names, you must specify
	//   CAPABILITY_NAMED_IAM .
	//   - If you don't specify either of these capabilities, CloudFormation returns
	//   an InsufficientCapabilities error. If your stack template contains these
	//   resources, we recommend that you review all permissions associated with them and
	//   edit their permissions if necessary.
	//   - AWS::IAM::AccessKey (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html)
	//   - AWS::IAM::Group (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html)
	//   - AWS::IAM::InstanceProfile (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html)
	//   - AWS::IAM::Policy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html)
	//   - AWS::IAM::Role (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html)
	//   - AWS::IAM::User (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html)
	//   - AWS::IAM::UserToGroupAddition (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html)
	//   For more information, see Acknowledging IAM Resources in CloudFormation
	//   Templates (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities)
	//   .
	//   - CAPABILITY_AUTO_EXPAND Some template contain macros. Macros perform custom
	//   processing on templates; this can include simple actions like find-and-replace
	//   operations, all the way to extensive transformations of entire templates.
	//   Because of this, users typically create a change set from the processed
	//   template, so that they can review the changes resulting from the macros before
	//   actually creating the stack. If your stack template contains one or more macros,
	//   and you choose to create a stack directly from the processed template, without
	//   first reviewing the resulting changes in a change set, you must acknowledge this
	//   capability. This includes the AWS::Include (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html)
	//   and AWS::Serverless (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html)
	//   transforms, which are macros hosted by CloudFormation. If you want to create a
	//   stack from a stack template that contains macros and nested stacks, you must
	//   create the stack directly from the template using this capability. You should
	//   only create stacks directly from a stack template that contains macros if you
	//   know what processing the macro performs. Each macro relies on an underlying
	//   Lambda service function for processing stack templates. Be aware that the Lambda
	//   function owner can update the function operation without CloudFormation being
	//   notified. For more information, see Using CloudFormation macros to perform
	//   custom processing on templates (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html)
	//   .
	Capabilities []types.Capability

	// A unique identifier for this CreateStack request. Specify this token if you
	// plan to retry requests so that CloudFormation knows that you're not attempting
	// to create a stack with the same name. You might retry CreateStack requests to
	// ensure that CloudFormation successfully received them. All events initiated by a
	// given stack operation are assigned the same client request token, which you can
	// use to track operations. For example, if you execute a CreateStack operation
	// with the token token1 , then all the StackEvents generated by that operation
	// will have ClientRequestToken set as token1 . In the console, stack operations
	// display the client request token on the Events tab. Stack operations that are
	// initiated from the console use the token format Console-StackOperation-ID, which
	// helps you easily identify the stack operation . For example, if you create a
	// stack using the console, each stack event would be assigned the same token in
	// the following format: Console-CreateStack-7f59c3cf-00d2-40c7-b2ff-e75db0987002 .
	ClientRequestToken *string

	// Set to true to disable rollback of the stack if stack creation failed. You can
	// specify either DisableRollback or OnFailure , but not both. Default: false
	DisableRollback *bool

	// Whether to enable termination protection on the specified stack. If a user
	// attempts to delete a stack with termination protection enabled, the operation
	// fails and the stack remains unchanged. For more information, see Protecting a
	// Stack From Being Deleted (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html)
	// in the CloudFormation User Guide. Termination protection is deactivated on
	// stacks by default. For nested stacks (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html)
	// , termination protection is set on the root stack and can't be changed directly
	// on the nested stack.
	EnableTerminationProtection *bool

	// The Amazon Simple Notification Service (Amazon SNS) topic ARNs to publish stack
	// related events. You can find your Amazon SNS topic ARNs using the Amazon SNS
	// console or your Command Line Interface (CLI).
	NotificationARNs []string

	// Determines what action will be taken if stack creation fails. This must be one
	// of: DO_NOTHING , ROLLBACK , or DELETE . You can specify either OnFailure or
	// DisableRollback , but not both. Default: ROLLBACK
	OnFailure types.OnFailure

	// A list of Parameter structures that specify input parameters for the stack. For
	// more information, see the Parameter (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html)
	// data type.
	Parameters []types.Parameter

	// The template resource types that you have permissions to work with for this
	// create stack action, such as AWS::EC2::Instance , AWS::EC2::* , or
	// Custom::MyCustomInstance . Use the following syntax to describe template
	// resource types: AWS::* (for all Amazon Web Services resources), Custom::* (for
	// all custom resources), Custom::logical_ID  (for a specific custom resource),
	// AWS::service_name::* (for all resources of a particular Amazon Web Services
	// service), and AWS::service_name::resource_logical_ID  (for a specific Amazon
	// Web Services resource). If the list of resource types doesn't include a resource
	// that you're creating, the stack creation fails. By default, CloudFormation
	// grants permissions to all resource types. Identity and Access Management (IAM)
	// uses this parameter for CloudFormation-specific condition keys in IAM policies.
	// For more information, see Controlling Access with Identity and Access Management (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html)
	// .
	ResourceTypes []string

	// When set to true , newly created resources are deleted when the operation rolls
	// back. This includes newly created resources marked with a deletion policy of
	// Retain . Default: false
	RetainExceptOnCreate *bool

	// The Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role
	// that CloudFormation assumes to create the stack. CloudFormation uses the role's
	// credentials to make calls on your behalf. CloudFormation always uses this role
	// for all future operations on the stack. Provided that users have permission to
	// operate on the stack, CloudFormation uses this role even if the users don't have
	// permission to pass it. Ensure that the role grants least privilege. If you don't
	// specify a value, CloudFormation uses the role that was previously associated
	// with the stack. If no role is available, CloudFormation uses a temporary session
	// that's generated from your user credentials.
	RoleARN *string

	// The rollback triggers for CloudFormation to monitor during stack creation and
	// updating operations, and for the specified monitoring period afterwards.
	RollbackConfiguration *types.RollbackConfiguration

	// Structure containing the stack policy body. For more information, go to
	// Prevent Updates to Stack Resources (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html)
	// in the CloudFormation User Guide. You can specify either the StackPolicyBody or
	// the StackPolicyURL parameter, but not both.
	StackPolicyBody *string

	// Location of a file containing the stack policy. The URL must point to a policy
	// (maximum size: 16 KB) located in an S3 bucket in the same Region as the stack.
	// You can specify either the StackPolicyBody or the StackPolicyURL parameter, but
	// not both.
	StackPolicyURL *string

	// Key-value pairs to associate with this stack. CloudFormation also propagates
	// these tags to the resources created in the stack. A maximum number of 50 tags
	// can be specified.
	Tags []types.Tag

	// Structure containing the template body with a minimum length of 1 byte and a
	// maximum length of 51,200 bytes. For more information, go to Template anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the CloudFormation User Guide. Conditional: You must specify either the
	// TemplateBody or the TemplateURL parameter, but not both.
	TemplateBody *string

	// Location of file containing the template body. The URL must point to a template
	// (max size: 460,800 bytes) that's located in an Amazon S3 bucket or a Systems
	// Manager document. For more information, go to the Template anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the CloudFormation User Guide. Conditional: You must specify either the
	// TemplateBody or the TemplateURL parameter, but not both.
	TemplateURL *string

	// The amount of time that can pass before the stack status becomes CREATE_FAILED;
	// if DisableRollback is not set or is set to false , the stack will be rolled back.
	TimeoutInMinutes *int32

	noSmithyDocumentSerde
}

func (*CreateStackInput) operationName() string {
	return "CreateStack"
}

// The output for a CreateStack action.
type CreateStackOutput struct {

	// Unique identifier of the stack.
	StackId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateStack{}, middleware.After)
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
	if err = addCreateStackResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "CreateStack",
	}
}

type opCreateStackResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateStackResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateStackResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "cloudformation"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "cloudformation"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("cloudformation")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCreateStackResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateStackResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
