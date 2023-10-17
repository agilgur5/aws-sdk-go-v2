// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Maps a user or group to the Amazon EMR Studio specified by StudioId , and
// applies a session policy to refine Studio permissions for that user or group.
// Use CreateStudioSessionMapping to assign users to a Studio when you use IAM
// Identity Center authentication. For instructions on how to assign users to a
// Studio when you use IAM authentication, see Assign a user or group to your EMR
// Studio (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-manage-users.html#emr-studio-assign-users-groups)
// .
func (c *Client) CreateStudioSessionMapping(ctx context.Context, params *CreateStudioSessionMappingInput, optFns ...func(*Options)) (*CreateStudioSessionMappingOutput, error) {
	if params == nil {
		params = &CreateStudioSessionMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStudioSessionMapping", params, optFns, c.addOperationCreateStudioSessionMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStudioSessionMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStudioSessionMappingInput struct {

	// Specifies whether the identity to map to the Amazon EMR Studio is a user or a
	// group.
	//
	// This member is required.
	IdentityType types.IdentityType

	// The Amazon Resource Name (ARN) for the session policy that will be applied to
	// the user or group. You should specify the ARN for the session policy that you
	// want to apply, not the ARN of your user role. For more information, see Create
	// an Amazon EMR Studio User Role with Session Policies (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-user-role.html)
	// .
	//
	// This member is required.
	SessionPolicyArn *string

	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	//
	// This member is required.
	StudioId *string

	// The globally unique identifier (GUID) of the user or group from the IAM
	// Identity Center Identity Store. For more information, see UserId (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserId)
	// and GroupId (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-GroupId)
	// in the IAM Identity Center Identity Store API Reference. Either IdentityName or
	// IdentityId must be specified, but not both.
	IdentityId *string

	// The name of the user or group. For more information, see UserName (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserName)
	// and DisplayName (https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-DisplayName)
	// in the IAM Identity Center Identity Store API Reference. Either IdentityName or
	// IdentityId must be specified, but not both.
	IdentityName *string

	noSmithyDocumentSerde
}

type CreateStudioSessionMappingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStudioSessionMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStudioSessionMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStudioSessionMapping{}, middleware.After)
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
	if err = addOpCreateStudioSessionMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStudioSessionMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStudioSessionMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "CreateStudioSessionMapping",
	}
}
