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

// Provides feedback for an authentication event indicating if it was from a valid
// user. This feedback is used for improving the risk evaluation decision for the
// user pool as part of Amazon Cognito advanced security. Amazon Cognito evaluates
// Identity and Access Management (IAM) policies in requests for this API
// operation. For this operation, you must use IAM credentials to authorize
// requests, and you must grant yourself the corresponding IAM permission in a
// policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) AdminUpdateAuthEventFeedback(ctx context.Context, params *AdminUpdateAuthEventFeedbackInput, optFns ...func(*Options)) (*AdminUpdateAuthEventFeedbackOutput, error) {
	if params == nil {
		params = &AdminUpdateAuthEventFeedbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminUpdateAuthEventFeedback", params, optFns, c.addOperationAdminUpdateAuthEventFeedbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminUpdateAuthEventFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminUpdateAuthEventFeedbackInput struct {

	// The authentication event ID.
	//
	// This member is required.
	EventId *string

	// The authentication event feedback value. When you provide a FeedbackValue value
	// of valid , you tell Amazon Cognito that you trust a user session where Amazon
	// Cognito has evaluated some level of risk. When you provide a FeedbackValue
	// value of invalid , you tell Amazon Cognito that you don't trust a user session,
	// or you don't believe that Amazon Cognito evaluated a high-enough risk level.
	//
	// This member is required.
	FeedbackValue types.FeedbackValueType

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user pool username.
	//
	// This member is required.
	Username *string

	noSmithyDocumentSerde
}

type AdminUpdateAuthEventFeedbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminUpdateAuthEventFeedbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminUpdateAuthEventFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminUpdateAuthEventFeedback{}, middleware.After)
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
	if err = addOpAdminUpdateAuthEventFeedbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminUpdateAuthEventFeedback(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminUpdateAuthEventFeedback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminUpdateAuthEventFeedback",
	}
}
