// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an email or SMS text message contact method. A contact method is used
// to send you notifications about your Amazon Lightsail resources. You can add one
// email address and one mobile phone number contact method in each Amazon Web
// Services Region. However, SMS text messaging is not supported in some Amazon Web
// Services Regions, and SMS text messages cannot be sent to some
// countries/regions. For more information, see Notifications in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-notifications)
// .
func (c *Client) CreateContactMethod(ctx context.Context, params *CreateContactMethodInput, optFns ...func(*Options)) (*CreateContactMethodOutput, error) {
	if params == nil {
		params = &CreateContactMethodInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContactMethod", params, optFns, c.addOperationCreateContactMethodMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContactMethodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContactMethodInput struct {

	// The destination of the contact method, such as an email address or a mobile
	// phone number. Use the E.164 format when specifying a mobile phone number. E.164
	// is a standard for the phone number structure used for international
	// telecommunication. Phone numbers that follow this format can have a maximum of
	// 15 digits, and they are prefixed with the plus character (+) and the country
	// code. For example, a U.S. phone number in E.164 format would be specified as
	// +1XXX5550100. For more information, see E.164 (https://en.wikipedia.org/wiki/E.164)
	// on Wikipedia.
	//
	// This member is required.
	ContactEndpoint *string

	// The protocol of the contact method, such as Email or SMS (text messaging). The
	// SMS protocol is supported only in the following Amazon Web Services Regions.
	//   - US East (N. Virginia) ( us-east-1 )
	//   - US West (Oregon) ( us-west-2 )
	//   - Europe (Ireland) ( eu-west-1 )
	//   - Asia Pacific (Tokyo) ( ap-northeast-1 )
	//   - Asia Pacific (Singapore) ( ap-southeast-1 )
	//   - Asia Pacific (Sydney) ( ap-southeast-2 )
	// For a list of countries/regions where SMS text messages can be sent, and the
	// latest Amazon Web Services Regions where SMS text messaging is supported, see
	// Supported Regions and Countries (https://docs.aws.amazon.com/sns/latest/dg/sns-supported-regions-countries.html)
	// in the Amazon SNS Developer Guide. For more information about notifications in
	// Amazon Lightsail, see Notifications in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-notifications)
	// .
	//
	// This member is required.
	Protocol types.ContactProtocol

	noSmithyDocumentSerde
}

type CreateContactMethodOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContactMethodMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContactMethod{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContactMethod{}, middleware.After)
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
	if err = addOpCreateContactMethodValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContactMethod(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContactMethod(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateContactMethod",
	}
}
