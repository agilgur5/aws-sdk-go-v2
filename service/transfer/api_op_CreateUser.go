// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a user and associates them with an existing file transfer
// protocol-enabled server. You can only create and associate users with servers
// that have the IdentityProviderType set to SERVICE_MANAGED . Using parameters for
// CreateUser , you can specify the user name, set the home directory, store the
// user's public key, and assign the user's Identity and Access Management (IAM)
// role. You can also optionally add a session policy, and assign metadata with
// tags that can be used to group and search for users.
func (c *Client) CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) {
	if params == nil {
		params = &CreateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUser", params, optFns, c.addOperationCreateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserInput struct {

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that controls your users' access to your Amazon S3 bucket or Amazon EFS file
	// system. The policies attached to this role determine the level of access that
	// you want to provide your users when transferring files into and out of your
	// Amazon S3 bucket or Amazon EFS file system. The IAM role should also contain a
	// trust relationship that allows the server to access your resources when
	// servicing your users' transfer requests.
	//
	// This member is required.
	Role *string

	// A system-assigned unique identifier for a server instance. This is the specific
	// server that you added your user to.
	//
	// This member is required.
	ServerId *string

	// A unique string that identifies a user and is associated with a ServerId . This
	// user name must be a minimum of 3 and a maximum of 100 characters long. The
	// following are valid characters: a-z, A-Z, 0-9, underscore '_', hyphen '-',
	// period '.', and at sign '@'. The user name can't start with a hyphen, period, or
	// at sign.
	//
	// This member is required.
	UserName *string

	// The landing directory (folder) for a user when they log in to the server using
	// the client. A HomeDirectory example is /bucket_name/home/mydirectory . The
	// HomeDirectory parameter is only used if HomeDirectoryType is set to PATH .
	HomeDirectory *string

	// Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and
	// keys should be visible to your user and how you want to make them visible. You
	// must specify the Entry and Target pair, where Entry shows how the path is made
	// visible and Target is the actual Amazon S3 or Amazon EFS path. If you only
	// specify a target, it is displayed as is. You also must ensure that your Identity
	// and Access Management (IAM) role provides access to paths in Target . This value
	// can be set only when HomeDirectoryType is set to LOGICAL. The following is an
	// Entry and Target pair example. [ { "Entry": "/directory1", "Target":
	// "/bucket_name/home/mydirectory" } ] In most cases, you can use this value
	// instead of the session policy to lock your user down to the designated home
	// directory (" chroot "). To do this, you can set Entry to / and set Target to
	// the value the user should see for their home directory when they log in. The
	// following is an Entry and Target pair example for chroot . [ { "Entry": "/",
	// "Target": "/bucket_name/home/mydirectory" } ]
	HomeDirectoryMappings []types.HomeDirectoryMapEntry

	// The type of landing directory (folder) that you want your users' home directory
	// to be when they log in to the server. If you set it to PATH , the user will see
	// the absolute Amazon S3 bucket or Amazon EFS path as is in their file transfer
	// protocol clients. If you set it to LOGICAL , you need to provide mappings in the
	// HomeDirectoryMappings for how you want to make Amazon S3 or Amazon EFS paths
	// visible to your users. If HomeDirectoryType is LOGICAL , you must provide
	// mappings, using the HomeDirectoryMappings parameter. If, on the other hand,
	// HomeDirectoryType is PATH , you provide an absolute path using the HomeDirectory
	// parameter. You cannot have both HomeDirectory and HomeDirectoryMappings in your
	// template.
	HomeDirectoryType types.HomeDirectoryType

	// A session policy for your user so that you can use the same Identity and Access
	// Management (IAM) role across multiple users. This policy scopes down a user's
	// access to portions of their Amazon S3 bucket. Variables that you can use inside
	// this policy include ${Transfer:UserName} , ${Transfer:HomeDirectory} , and
	// ${Transfer:HomeBucket} . This policy applies only when the domain of ServerId
	// is Amazon S3. Amazon EFS does not use session policies. For session policies,
	// Transfer Family stores the policy as a JSON blob, instead of the Amazon Resource
	// Name (ARN) of the policy. You save the policy as a JSON blob and pass it in the
	// Policy argument. For an example of a session policy, see Example session policy (https://docs.aws.amazon.com/transfer/latest/userguide/session-policy.html)
	// . For more information, see AssumeRole (https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html)
	// in the Amazon Web Services Security Token Service API Reference.
	Policy *string

	// Specifies the full POSIX identity, including user ID ( Uid ), group ID ( Gid ),
	// and any secondary groups IDs ( SecondaryGids ), that controls your users' access
	// to your Amazon EFS file systems. The POSIX permissions that are set on files and
	// directories in Amazon EFS determine the level of access your users get when
	// transferring files into and out of your Amazon EFS file systems.
	PosixProfile *types.PosixProfile

	// The public portion of the Secure Shell (SSH) key used to authenticate the user
	// to the server. The three standard SSH public key format elements are , , and an
	// optional , with spaces between each element. Transfer Family accepts RSA,
	// ECDSA, and ED25519 keys.
	//   - For RSA keys, the key type is ssh-rsa .
	//   - For ED25519 keys, the key type is ssh-ed25519 .
	//   - For ECDSA keys, the key type is either ecdsa-sha2-nistp256 ,
	//   ecdsa-sha2-nistp384 , or ecdsa-sha2-nistp521 , depending on the size of the
	//   key you generated.
	SshPublicKeyBody *string

	// Key-value pairs that can be used to group and search for users. Tags are
	// metadata attached to users for any purpose.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateUserOutput struct {

	// The identifier of the server that the user is attached to.
	//
	// This member is required.
	ServerId *string

	// A unique string that identifies a Transfer Family user.
	//
	// This member is required.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUser{}, middleware.After)
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
	if err = addOpCreateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateUser",
	}
}
