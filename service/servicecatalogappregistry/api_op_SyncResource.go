// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Syncs the resource with current AppRegistry records. Specifically, the
// resource’s AppRegistry system tags sync with its associated application. We
// remove the resource's AppRegistry system tags if it does not associate with the
// application. The caller must have permissions to read and update the resource.
func (c *Client) SyncResource(ctx context.Context, params *SyncResourceInput, optFns ...func(*Options)) (*SyncResourceOutput, error) {
	if params == nil {
		params = &SyncResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SyncResource", params, optFns, c.addOperationSyncResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SyncResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SyncResourceInput struct {

	// An entity you can work with and specify with a name or ID. Examples include an
	// Amazon EC2 instance, an Amazon Web Services CloudFormation stack, or an Amazon
	// S3 bucket.
	//
	// This member is required.
	Resource *string

	// The type of resource of which the application will be associated.
	//
	// This member is required.
	ResourceType types.ResourceType

	noSmithyDocumentSerde
}

type SyncResourceOutput struct {

	// The results of the output if an application is associated with an ARN value,
	// which could be syncStarted or None.
	ActionTaken types.SyncAction

	// The Amazon resource name (ARN) that specifies the application.
	ApplicationArn *string

	// The Amazon resource name (ARN) that specifies the resource.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSyncResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSyncResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSyncResource{}, middleware.After)
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
	if err = addOpSyncResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSyncResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSyncResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "SyncResource",
	}
}
