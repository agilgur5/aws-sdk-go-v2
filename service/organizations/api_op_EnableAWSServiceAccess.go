// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the integration of an Amazon Web Services service (the service that is
// specified by ServicePrincipal ) with Organizations. When you enable integration,
// you allow the specified service to create a service-linked role (https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in all the accounts in your organization. This allows the service to perform
// operations on your behalf in your organization and its accounts. We recommend
// that you enable integration between Organizations and the specified Amazon Web
// Services service by using the console or commands that are provided by the
// specified service. Doing so ensures that the service is aware that it can create
// the resources that are required for the integration. How the service creates
// those resources in the organization's accounts depends on that service. For more
// information, see the documentation for the other Amazon Web Services service.
// For more information about enabling services to integrate with Organizations,
// see Using Organizations with other Amazon Web Services services (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the Organizations User Guide. You can only call this operation from the
// organization's management account and only if the organization has enabled all
// features (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
// .
func (c *Client) EnableAWSServiceAccess(ctx context.Context, params *EnableAWSServiceAccessInput, optFns ...func(*Options)) (*EnableAWSServiceAccessOutput, error) {
	if params == nil {
		params = &EnableAWSServiceAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableAWSServiceAccess", params, optFns, c.addOperationEnableAWSServiceAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableAWSServiceAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableAWSServiceAccessInput struct {

	// The service principal name of the Amazon Web Services service for which you
	// want to enable integration with your organization. This is typically in the form
	// of a URL, such as service-abbreviation.amazonaws.com .
	//
	// This member is required.
	ServicePrincipal *string

	noSmithyDocumentSerde
}

type EnableAWSServiceAccessOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableAWSServiceAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableAWSServiceAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableAWSServiceAccess{}, middleware.After)
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
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointV2Middleware(stack, options); err != nil {
		return err
	}
	if err = addOpEnableAWSServiceAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableAWSServiceAccess(options.Region), middleware.Before); err != nil {
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
		operation: "EnableAWSServiceAccess",
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

func newServiceMetadataMiddleware_opEnableAWSServiceAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "EnableAWSServiceAccess",
	}
}
