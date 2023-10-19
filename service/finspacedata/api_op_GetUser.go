// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves details for a specific user.
func (c *Client) GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) {
	if params == nil {
		params = &GetUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUser", params, optFns, c.addOperationGetUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserInput struct {

	// The unique identifier of the user to get data for.
	//
	// This member is required.
	UserId *string

	noSmithyDocumentSerde
}

type GetUserOutput struct {

	// Indicates whether the user can use the GetProgrammaticAccessCredentials API to
	// obtain credentials that can then be used to access other FinSpace Data API
	// operations.
	//   - ENABLED – The user has permissions to use the APIs.
	//   - DISABLED – The user does not have permissions to use any APIs.
	ApiAccess types.ApiAccess

	// The ARN identifier of an AWS user or role that is allowed to call the
	// GetProgrammaticAccessCredentials API to obtain a credentials token for a
	// specific FinSpace user. This must be an IAM role within your FinSpace account.
	ApiAccessPrincipalArn *string

	// The timestamp at which the user was created in FinSpace. The value is
	// determined as epoch time in milliseconds.
	CreateTime int64

	// The email address that is associated with the user.
	EmailAddress *string

	// The first name of the user.
	FirstName *string

	// Describes the last time the user was deactivated. The value is determined as
	// epoch time in milliseconds.
	LastDisabledTime int64

	// Describes the last time the user was activated. The value is determined as
	// epoch time in milliseconds.
	LastEnabledTime int64

	// Describes the last time that the user logged into their account. The value is
	// determined as epoch time in milliseconds.
	LastLoginTime int64

	// Describes the last time the user details were updated. The value is determined
	// as epoch time in milliseconds.
	LastModifiedTime int64

	// The last name of the user.
	LastName *string

	// The current status of the user.
	//   - CREATING – The creation is in progress.
	//   - ENABLED – The user is created and is currently active.
	//   - DISABLED – The user is currently inactive.
	Status types.UserStatus

	// Indicates the type of user.
	//   - SUPER_USER – A user with permission to all the functionality and data in
	//   FinSpace.
	//
	//   - APP_USER – A user with specific permissions in FinSpace. The users are
	//   assigned permissions by adding them to a permission group.
	Type types.UserType

	// The unique identifier for the user that is retrieved.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUser{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpGetUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUser(options.Region), middleware.Before); err != nil {
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
		operation: "GetUser",
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

func newServiceMetadataMiddleware_opGetUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "GetUser",
	}
}
