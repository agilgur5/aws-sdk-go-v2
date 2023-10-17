// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables Shield Advanced for a specific Amazon Web Services resource. The
// resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone,
// Global Accelerator standard accelerator, Elastic IP Address, Application Load
// Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and
// Network Load Balancers by association with protected Amazon EC2 Elastic IP
// addresses. You can add protection to only a single resource with each
// CreateProtection request. You can add protection to multiple resources at once
// through the Shield Advanced console at
// https://console.aws.amazon.com/wafv2/shieldv2#/ (https://console.aws.amazon.com/wafv2/shieldv2#/)
// . For more information see Getting Started with Shield Advanced (https://docs.aws.amazon.com/waf/latest/developerguide/getting-started-ddos.html)
// and Adding Shield Advanced protection to Amazon Web Services resources (https://docs.aws.amazon.com/waf/latest/developerguide/configure-new-protection.html)
// .
func (c *Client) CreateProtection(ctx context.Context, params *CreateProtectionInput, optFns ...func(*Options)) (*CreateProtectionOutput, error) {
	if params == nil {
		params = &CreateProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProtection", params, optFns, c.addOperationCreateProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProtectionInput struct {

	// Friendly name for the Protection you are creating.
	//
	// This member is required.
	Name *string

	// The ARN (Amazon Resource Name) of the resource to be protected. The ARN should
	// be in one of the following formats:
	//   - For an Application Load Balancer:
	//   arn:aws:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	//   - For an Elastic Load Balancer (Classic Load Balancer):
	//   arn:aws:elasticloadbalancing:region:account-id:loadbalancer/load-balancer-name
	//   - For an Amazon CloudFront distribution:
	//   arn:aws:cloudfront::account-id:distribution/distribution-id
	//   - For an Global Accelerator standard accelerator:
	//   arn:aws:globalaccelerator::account-id:accelerator/accelerator-id
	//   - For Amazon Route 53: arn:aws:route53:::hostedzone/hosted-zone-id
	//   - For an Elastic IP address:
	//   arn:aws:ec2:region:account-id:eip-allocation/allocation-id
	//
	// This member is required.
	ResourceArn *string

	// One or more tag key-value pairs for the Protection object that is created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProtectionOutput struct {

	// The unique identifier (ID) for the Protection object that is created.
	ProtectionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProtection{}, middleware.After)
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
	if err = addOpCreateProtectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProtection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProtection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "CreateProtection",
	}
}
