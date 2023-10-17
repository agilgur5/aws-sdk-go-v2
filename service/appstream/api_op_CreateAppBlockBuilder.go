// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an app block builder.
func (c *Client) CreateAppBlockBuilder(ctx context.Context, params *CreateAppBlockBuilderInput, optFns ...func(*Options)) (*CreateAppBlockBuilderOutput, error) {
	if params == nil {
		params = &CreateAppBlockBuilderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppBlockBuilder", params, optFns, c.addOperationCreateAppBlockBuilderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppBlockBuilderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppBlockBuilderInput struct {

	// The instance type to use when launching the app block builder. The following
	// instance types are available:
	//   - stream.standard.small
	//   - stream.standard.medium
	//   - stream.standard.large
	//   - stream.standard.xlarge
	//   - stream.standard.2xlarge
	//
	// This member is required.
	InstanceType *string

	// The unique name for the app block builder.
	//
	// This member is required.
	Name *string

	// The platform of the app block builder. WINDOWS_SERVER_2019 is the only valid
	// value.
	//
	// This member is required.
	Platform types.AppBlockBuilderPlatformType

	// The VPC configuration for the app block builder. App block builders require
	// that you specify at least two subnets in different availability zones.
	//
	// This member is required.
	VpcConfig *types.VpcConfig

	// The list of interface VPC endpoint (interface endpoint) objects. Administrators
	// can connect to the app block builder only through the specified endpoints.
	AccessEndpoints []types.AccessEndpoint

	// The description of the app block builder.
	Description *string

	// The display name of the app block builder.
	DisplayName *string

	// Enables or disables default internet access for the app block builder.
	EnableDefaultInternetAccess *bool

	// The Amazon Resource Name (ARN) of the IAM role to apply to the app block
	// builder. To assume a role, the app block builder calls the AWS Security Token
	// Service (STS) AssumeRole API operation and passes the ARN of the role to use.
	// The operation creates a new session with temporary credentials. AppStream 2.0
	// retrieves the temporary credentials and creates the appstream_machine_role
	// credential profile on the instance. For more information, see Using an IAM Role
	// to Grant Permissions to Applications and Scripts Running on AppStream 2.0
	// Streaming Instances (https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	IamRoleArn *string

	// The tags to associate with the app block builder. A tag is a key-value pair,
	// and the value is optional. For example, Environment=Test. If you do not specify
	// a value, Environment=. If you do not specify a value, the value is set to an
	// empty string. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following special characters: _ . : / = + \ - @
	// For more information, see Tagging Your Resources (https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAppBlockBuilderOutput struct {

	// Describes an app block builder.
	AppBlockBuilder *types.AppBlockBuilder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppBlockBuilderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAppBlockBuilder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAppBlockBuilder{}, middleware.After)
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
	if err = addOpCreateAppBlockBuilderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppBlockBuilder(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAppBlockBuilder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateAppBlockBuilder",
	}
}
