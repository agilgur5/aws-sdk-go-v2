// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/) , the AWS
// CloudHSM Classic User Guide (https://docs.aws.amazon.com/cloudhsm/classic/userguide/)
// , and the AWS CloudHSM Classic API Reference (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/)
// . For information about the current version of AWS CloudHSM, see AWS CloudHSM (http://aws.amazon.com/cloudhsm/)
// , the AWS CloudHSM User Guide (https://docs.aws.amazon.com/cloudhsm/latest/userguide/)
// , and the AWS CloudHSM API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/)
// . Modifies an HSM. This operation can result in the HSM being offline for up to
// 15 minutes while the AWS CloudHSM service is reconfigured. If you are modifying
// a production HSM, you should ensure that your AWS CloudHSM service is configured
// for high availability, and consider executing this operation during a
// maintenance window.
//
// Deprecated: This API is deprecated.
func (c *Client) ModifyHsm(ctx context.Context, params *ModifyHsmInput, optFns ...func(*Options)) (*ModifyHsmOutput, error) {
	if params == nil {
		params = &ModifyHsmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyHsm", params, optFns, c.addOperationModifyHsmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the ModifyHsm operation.
type ModifyHsmInput struct {

	// The ARN of the HSM to modify.
	//
	// This member is required.
	HsmArn *string

	// The new IP address for the elastic network interface (ENI) attached to the HSM.
	// If the HSM is moved to a different subnet, and an IP address is not specified,
	// an IP address will be randomly chosen from the CIDR range of the new subnet.
	EniIp *string

	// The new external ID.
	ExternalId *string

	// The new IAM role ARN.
	IamRoleArn *string

	// The new identifier of the subnet that the HSM is in. The new subnet must be in
	// the same Availability Zone as the current subnet.
	SubnetId *string

	// The new IP address for the syslog monitoring server. The AWS CloudHSM service
	// only supports one syslog monitoring server.
	SyslogIp *string

	noSmithyDocumentSerde
}

// Contains the output of the ModifyHsm operation.
type ModifyHsmOutput struct {

	// The ARN of the HSM.
	HsmArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyHsmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyHsm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyHsm{}, middleware.After)
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
	if err = addOpModifyHsmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyHsm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyHsm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "ModifyHsm",
	}
}
