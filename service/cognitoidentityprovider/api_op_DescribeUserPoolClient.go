// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Client method for returning the configuration information and metadata of the
// specified user pool app client. Amazon Cognito evaluates Identity and Access
// Management (IAM) policies in requests for this API operation. For this
// operation, you must use IAM credentials to authorize requests, and you must
// grant yourself the corresponding IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) DescribeUserPoolClient(ctx context.Context, params *DescribeUserPoolClientInput, optFns ...func(*Options)) (*DescribeUserPoolClientOutput, error) {
	if params == nil {
		params = &DescribeUserPoolClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUserPoolClient", params, optFns, c.addOperationDescribeUserPoolClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserPoolClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to describe a user pool client.
type DescribeUserPoolClientInput struct {

	// The app client ID of the app associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The user pool ID for the user pool you want to describe.
	//
	// This member is required.
	UserPoolId *string

	noSmithyDocumentSerde
}

func (*DescribeUserPoolClientInput) operationName() string {
	return "DescribeUserPoolClient"
}

// Represents the response from the server from a request to describe the user
// pool client.
type DescribeUserPoolClientOutput struct {

	// The user pool client from a server response to describe the user pool client.
	UserPoolClient *types.UserPoolClientType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUserPoolClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUserPoolClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUserPoolClient{}, middleware.After)
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
	if err = addOpDescribeUserPoolClientValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUserPoolClient(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeUserPoolClient(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "DescribeUserPoolClient",
	}
}
