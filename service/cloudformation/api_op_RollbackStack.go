// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// When specifying RollbackStack , you preserve the state of previously provisioned
// resources when an operation fails. You can check the status of the stack through
// the DescribeStacks operation. Rolls back the specified stack to the last known
// stable state from CREATE_FAILED or UPDATE_FAILED stack statuses. This operation
// will delete a stack if it doesn't contain a last known stable state. A last
// known stable state includes any status in a *_COMPLETE . This includes the
// following stack statuses.
//   - CREATE_COMPLETE
//   - UPDATE_COMPLETE
//   - UPDATE_ROLLBACK_COMPLETE
//   - IMPORT_COMPLETE
//   - IMPORT_ROLLBACK_COMPLETE
func (c *Client) RollbackStack(ctx context.Context, params *RollbackStackInput, optFns ...func(*Options)) (*RollbackStackOutput, error) {
	if params == nil {
		params = &RollbackStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RollbackStack", params, optFns, c.addOperationRollbackStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RollbackStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RollbackStackInput struct {

	// The name that's associated with the stack.
	//
	// This member is required.
	StackName *string

	// A unique identifier for this RollbackStack request.
	ClientRequestToken *string

	// When set to true , newly created resources are deleted when the operation rolls
	// back. This includes newly created resources marked with a deletion policy of
	// Retain . Default: false
	RetainExceptOnCreate *bool

	// The Amazon Resource Name (ARN) of an Identity and Access Management role that
	// CloudFormation assumes to rollback the stack.
	RoleARN *string

	noSmithyDocumentSerde
}

type RollbackStackOutput struct {

	// Unique identifier of the stack.
	StackId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRollbackStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRollbackStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRollbackStack{}, middleware.After)
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
	if err = addOpRollbackStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRollbackStack(options.Region), middleware.Before); err != nil {
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
		operation: "RollbackStack",
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

func newServiceMetadataMiddleware_opRollbackStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "RollbackStack",
	}
}
