// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a third-party custom source in Amazon Security Lake, from the Amazon Web
// Services Region where you want to create a custom source. Security Lake can
// collect logs and events from third-party custom sources. After creating the
// appropriate IAM role to invoke Glue crawler, use this API to add a custom source
// name in Security Lake. This operation creates a partition in the Amazon S3
// bucket for Security Lake as the target location for log files from the custom
// source. In addition, this operation also creates an associated Glue table and an
// Glue crawler.
func (c *Client) CreateCustomLogSource(ctx context.Context, params *CreateCustomLogSourceInput, optFns ...func(*Options)) (*CreateCustomLogSourceOutput, error) {
	if params == nil {
		params = &CreateCustomLogSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomLogSource", params, optFns, c.addOperationCreateCustomLogSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomLogSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomLogSourceInput struct {

	// Specify the name for a third-party custom source. This must be a Regionally
	// unique value.
	//
	// This member is required.
	SourceName *string

	// The configuration for the third-party custom source.
	Configuration *types.CustomLogSourceConfiguration

	// The Open Cybersecurity Schema Framework (OCSF) event classes which describes
	// the type of data that the custom source will send to Security Lake. The
	// supported event classes are:
	//   - ACCESS_ACTIVITY
	//   - FILE_ACTIVITY
	//   - KERNEL_ACTIVITY
	//   - KERNEL_EXTENSION
	//   - MEMORY_ACTIVITY
	//   - MODULE_ACTIVITY
	//   - PROCESS_ACTIVITY
	//   - REGISTRY_KEY_ACTIVITY
	//   - REGISTRY_VALUE_ACTIVITY
	//   - RESOURCE_ACTIVITY
	//   - SCHEDULED_JOB_ACTIVITY
	//   - SECURITY_FINDING
	//   - ACCOUNT_CHANGE
	//   - AUTHENTICATION
	//   - AUTHORIZATION
	//   - ENTITY_MANAGEMENT_AUDIT
	//   - DHCP_ACTIVITY
	//   - NETWORK_ACTIVITY
	//   - DNS_ACTIVITY
	//   - FTP_ACTIVITY
	//   - HTTP_ACTIVITY
	//   - RDP_ACTIVITY
	//   - SMB_ACTIVITY
	//   - SSH_ACTIVITY
	//   - CONFIG_STATE
	//   - INVENTORY_INFO
	//   - EMAIL_ACTIVITY
	//   - API_ACTIVITY
	//   - CLOUD_API
	EventClasses []string

	// Specify the source version for the third-party custom source, to limit log
	// collection to a specific version of custom data source.
	SourceVersion *string

	noSmithyDocumentSerde
}

type CreateCustomLogSourceOutput struct {

	// The created third-party custom source.
	Source *types.CustomLogSourceResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomLogSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCustomLogSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCustomLogSource{}, middleware.After)
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
	if err = addOpCreateCustomLogSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomLogSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCustomLogSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateCustomLogSource",
	}
}
