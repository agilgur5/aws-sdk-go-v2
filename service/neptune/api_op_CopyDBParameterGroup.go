// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies the specified DB parameter group.
func (c *Client) CopyDBParameterGroup(ctx context.Context, params *CopyDBParameterGroupInput, optFns ...func(*Options)) (*CopyDBParameterGroupOutput, error) {
	if params == nil {
		params = &CopyDBParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDBParameterGroup", params, optFns, c.addOperationCopyDBParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyDBParameterGroupInput struct {

	// The identifier or ARN for the source DB parameter group. For information about
	// creating an ARN, see Constructing an Amazon Resource Name (ARN) (https://docs.aws.amazon.com/neptune/latest/UserGuide/tagging.ARN.html#tagging.ARN.Constructing)
	// . Constraints:
	//   - Must specify a valid DB parameter group.
	//   - Must specify a valid DB parameter group identifier, for example
	//   my-db-param-group , or a valid ARN.
	//
	// This member is required.
	SourceDBParameterGroupIdentifier *string

	// A description for the copied DB parameter group.
	//
	// This member is required.
	TargetDBParameterGroupDescription *string

	// The identifier for the copied DB parameter group. Constraints:
	//   - Cannot be null, empty, or blank.
	//   - Must contain from 1 to 255 letters, numbers, or hyphens.
	//   - First character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	// Example: my-db-parameter-group
	//
	// This member is required.
	TargetDBParameterGroupIdentifier *string

	// The tags to be assigned to the copied DB parameter group.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CopyDBParameterGroupOutput struct {

	// Contains the details of an Amazon Neptune DB parameter group. This data type is
	// used as a response element in the DescribeDBParameterGroups action.
	DBParameterGroup *types.DBParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyDBParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBParameterGroup{}, middleware.After)
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
	if err = addOpCopyDBParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBParameterGroup(options.Region), middleware.Before); err != nil {
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
		operation: "CopyDBParameterGroup",
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

func newServiceMetadataMiddleware_opCopyDBParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBParameterGroup",
	}
}
