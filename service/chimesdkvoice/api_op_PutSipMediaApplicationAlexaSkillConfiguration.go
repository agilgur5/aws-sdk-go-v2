// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Alexa Skill configuration for the SIP media application.
func (c *Client) PutSipMediaApplicationAlexaSkillConfiguration(ctx context.Context, params *PutSipMediaApplicationAlexaSkillConfigurationInput, optFns ...func(*Options)) (*PutSipMediaApplicationAlexaSkillConfigurationOutput, error) {
	if params == nil {
		params = &PutSipMediaApplicationAlexaSkillConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSipMediaApplicationAlexaSkillConfiguration", params, optFns, c.addOperationPutSipMediaApplicationAlexaSkillConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSipMediaApplicationAlexaSkillConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSipMediaApplicationAlexaSkillConfigurationInput struct {

	// The SIP media application ID.
	//
	// This member is required.
	SipMediaApplicationId *string

	// The Alexa Skill configuration.
	SipMediaApplicationAlexaSkillConfiguration *types.SipMediaApplicationAlexaSkillConfiguration

	noSmithyDocumentSerde
}

type PutSipMediaApplicationAlexaSkillConfigurationOutput struct {

	// Returns the Alexa Skill configuration.
	SipMediaApplicationAlexaSkillConfiguration *types.SipMediaApplicationAlexaSkillConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSipMediaApplicationAlexaSkillConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSipMediaApplicationAlexaSkillConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSipMediaApplicationAlexaSkillConfiguration{}, middleware.After)
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
	if err = addOpPutSipMediaApplicationAlexaSkillConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSipMediaApplicationAlexaSkillConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSipMediaApplicationAlexaSkillConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutSipMediaApplicationAlexaSkillConfiguration",
	}
}
